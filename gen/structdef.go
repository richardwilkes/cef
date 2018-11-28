package main

import (
	"sort"
	"strings"
	"text/template"

	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/toolbox/txt"
)

type structDef struct {
	Name     string
	GoName   string
	Position position
	Fields   []*field
}

func newStructDef(name string, pos position) *structDef {
	return &structDef{
		Name:     name,
		GoName:   translateStructTypeName(name),
		Position: pos,
	}
}

func (s *structDef) NeedsUnsafeImport() bool {
	for _, f := range s.Fields {
		if f.Var.NeedUnsafe {
			return true
		}
	}
	return false
}

func (s *structDef) isClassEquivalent() bool {
	for _, f := range s.Fields {
		if f.Var.FunctionPtr {
			return true
		}
	}
	return false
}

func (s *structDef) TrimmedName() string {
	return strings.TrimSuffix(strings.TrimPrefix(s.Name, "cef_"), "_t")
}

func (s *structDef) Trampolines() []string {
	var result []string
	for _, f := range s.Fields {
		if f.Var.FunctionPtr {
			result = append(result, f.Trampoline())
		}
	}
	return result
}

func translateStructTypeName(name string) string {
	if strings.HasPrefix(name, "cef_") && strings.HasSuffix(name, "_t") {
		name = translateConstantName(name[4 : len(name)-2])
	}
	return name
}

func processRecordDecl(block []lineInfo) {
	line := block[0].Line
	if strings.HasSuffix(line, " definition") {
		if i := strings.Index(line, " struct _cef_"); i != -1 {
			name := line[i+9 : len(line)-11]
			if _, exclude := excludeMap[name]; !exclude {
				if _, exists := sdefsMap[name]; !exists {
					sdef := newStructDef(name, block[0].Position)
					for i := 1; i < len(block); i++ {
						line = block[i].Line
						if len(line) > 3 && strings.HasPrefix(line[3:], "-FieldDecl ") {
							if start := strings.Index(line, "> "); start != -1 {
								line = line[start+2:]
								if space := strings.Index(line, " "); space != -1 {
									line = line[space+1:]
									if start = strings.Index(line, " "); space != -1 {
										sdef.Fields = append(sdef.Fields, newField(sdef, line[:start], strings.Trim(line[start+1:], "'"), block[i].Position))
									}
								}
							}
						}
					}
					sdefsMap[name] = sdef
				}
			}
		}
	}
}

func dumpStructs() {
	sdefs := make([]*structDef, 0, len(sdefsMap))
	for _, sdef := range sdefsMap {
		sdefs = append(sdefs, sdef)
	}
	sort.Slice(sdefs, func(i, j int) bool { return txt.NaturalLess(sdefs[i].Name, sdefs[j].Name, true) })
	const classTmplFile = "class.go.tmpl"
	const structTmplFile = "struct.go.tmpl"
	const callbackTmplFile = "callback.go.tmpl"
	const callbackHeaderTmplFile = "callback.h.tmpl"
	const callbackCTmplFile = "callback.c.tmpl"
	tmpl, err := template.ParseFiles(classTmplFile, structTmplFile, callbackTmplFile, callbackHeaderTmplFile, callbackCTmplFile)
	jot.FatalIfErr(err)

	for _, sdef := range sdefs {
		if sdef.GoName != "MainArgs" && sdef.GoName != "WindowInfo" {
			var tmplFile string
			if sdef.isClassEquivalent() {
				if strings.HasSuffix(sdef.GoName, "Visitor") || strings.HasSuffix(sdef.GoName, "Callback") || strings.HasSuffix(sdef.GoName, "Handler") || strings.HasSuffix(sdef.GoName, "Client") || sdef.GoName == "App" {
					genSourceFile(tmpl, callbackHeaderTmplFile, sdef.GoName+"_gen.h", sdef)
					genSourceFile(tmpl, callbackCTmplFile, sdef.GoName+"_gen.c", sdef)
					tmplFile = callbackTmplFile
				} else {
					tmplFile = classTmplFile
				}
			} else {
				tmplFile = structTmplFile
			}
			genSourceFile(tmpl, tmplFile, sdef.GoName+"_gen.go", sdef)
		}
	}
}
