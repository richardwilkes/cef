package main

import (
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/toolbox/txt"
)

var (
	defRegex   = regexp.MustCompile(` struct\s+_(cef_\S+)\s+definition$`)
	fieldRegex = regexp.MustCompile(`-FieldDecl .*>\s+\S+\s+(\S+)\s+'([^']+)'`)
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

func (s *structDef) NeedsSizeSet() bool {
	if !s.isClassEquivalent() {
		for _, f := range s.Fields {
			if f.Var.GoName == "Size" && f.Var.BaseType == "size_t" {
				return true
			}
		}
	}
	return false
}

func (s *structDef) NeedsUnsafeImport() bool {
	for _, f := range s.Fields {
		if f.Var.NeedUnsafe {
			return true
		}
	}
	return s.hasInheritance()
}

func (s *structDef) hasInheritance() bool {
	return !s.isClassEquivalent() && len(s.Fields) > 0 && s.Fields[0].Var.Name == "base"
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
	if result := defRegex.FindAllStringSubmatch(block[0].Line, -1); result != nil {
		name := result[0][1]
		if _, exclude := excludeMap[name]; !exclude {
			if _, exists := sdefsMap[name]; !exists {
				sdef := newStructDef(name, block[0].Position)
				for i := 1; i < len(block); i++ {
					if result = fieldRegex.FindAllStringSubmatch(block[i].Line, -1); result != nil {
						sdef.Fields = append(sdef.Fields, newField(sdef, result[0][1], result[0][2], block[i].Position))
					}
				}
				sdefsMap[name] = sdef
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
				if strings.HasSuffix(sdef.GoName, "Visitor") || strings.HasSuffix(sdef.GoName, "Callback") || strings.HasSuffix(sdef.GoName, "Handler") || strings.HasSuffix(sdef.GoName, "Delegate") || strings.HasSuffix(sdef.GoName, "Filter") || strings.HasSuffix(sdef.GoName, "Client") || sdef.GoName == "App" || sdef.GoName == "Task" {
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
