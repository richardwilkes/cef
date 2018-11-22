package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/toolbox/txt"
)

var colRegex = regexp.MustCompile(`<col:(\d+)(, col:(\d+))?>`)

type enumDef struct {
	Name     string
	GoName   string
	Values   []*enumValue
	Position position
	Unsigned bool
}

func newEnumDef(name string, pos position) *enumDef {
	return &enumDef{
		Name:     name,
		GoName:   translateStructTypeName(name),
		Position: pos,
	}
}

func (e *enumDef) Type() string {
	if e.Unsigned {
		return "uint"
	}
	return "int"
}

var nameExceptions = []string{
	"ascii", "com", "ct", "dom", "html", "id", "io", "json", "js", "pdf",
	"ssl", "ssl2", "ssl3", "st", "tid", "tls", "tls1", "tp", "ts", "tt",
	"ui", "uri", "url", "v8", "x509", "xml",
}

func translateConstantName(in string) string {
	in = strings.TrimSpace(in)
	if strings.HasPrefix(in, "CEF_") {
		in = in[4:]
	}
	in = strings.ToLower(in)
	for _, s := range nameExceptions {
		upper := strings.ToUpper(s)
		if strings.HasPrefix(in, s+"_") {
			in = upper + "_" + in[len(s)+1:]
		}
		if strings.HasSuffix(in, "_"+s) {
			in = in[:len(in)-(1+len(s))] + "_" + upper
		}
		for {
			if i := strings.Index(in, "_"+s+"_"); i == -1 {
				break
			} else {
				in = in[:i+1] + upper + in[i+1+len(s):]
			}
		}
	}
	in = txt.ToCamelCase(in)
	return in
}

func processTypedefDeclEnumDecl(curBlock, prevBlock []lineInfo) {
	if i := strings.Index(curBlock[0].Line, "'enum cef_"); i != -1 {
		name := curBlock[0].Line[i+5:]
		if i = strings.Index(name, "':'"); i != -1 {
			name = name[:i]
		}
		name = strings.TrimSpace(strings.TrimSuffix(name, "'"))
		if _, exists := edefsMap[name]; !exists {
			edef := newEnumDef(name, curBlock[0].Position)
			lookForValue := false
			for i = 1; i < len(prevBlock); i++ {
				line := prevBlock[i].Line
				if strings.Contains(line, "-EnumConstantDecl ") {
					lookForValue = true
					found := false
					if strings.HasSuffix(line, " 'unsigned int'") {
						edef.Unsigned = true
						line = line[:len(line)-15]
						found = true
					} else if strings.HasSuffix(line, " 'int'") {
						line = line[:len(line)-6]
						found = true
					}
					if found {
						found = false
						if space := strings.LastIndex(line, " "); space != -1 {
							found = true
							value := "0"
							if len(edef.Values) > 0 {
								prev := edef.Values[len(edef.Values)-1]
								if prev.IsNum {
									value = strconv.Itoa(prev.Num + 1)
								} else {
									value = fmt.Sprintf("(%s) + 1", prev.Value)
								}
							}
							edef.Values = append(edef.Values, newEnumValue(line[space+1:], value, prevBlock[i].Position))
						}
					}
					if !found {
						jot.Fatal(1, errs.New(prevBlock[i].Line))
					}
				} else if lookForValue {
					if result := colRegex.FindAllStringSubmatch(line, -1); len(result) > 0 {
						var ecol int
						scol, err := strconv.Atoi(result[0][1])
						jot.FatalIfErr(err)
						if result[0][3] == "" {
							ecol = scol
						} else {
							ecol, err = strconv.Atoi(result[0][3])
							jot.FatalIfErr(err)
						}
						v := prevBlock[i].Position.Text(prevBlock[i].Position.LineStart, scol, 1000)
						if comma := strings.IndexAny(v[ecol-scol:], ", "); comma != -1 {
							v = v[:comma+ecol-scol]
						}
						if strings.Contains(line, "-DeclRefExpr ") {
							v = translateConstantName(v)
						}
						edef.Values[len(edef.Values)-1].setValue(v)
					}
					lookForValue = false
				}
			}
			edefsMap[name] = edef
		}
	}
}

func dumpEnums() {
	edefs := make([]*enumDef, 0, len(edefsMap))
	for _, edef := range edefsMap {
		edefs = append(edefs, edef)
	}
	sort.Slice(edefs, func(i, j int) bool { return txt.NaturalLess(edefs[i].Name, edefs[j].Name, true) })
	const enumTmplFile = "enum.go.tmpl"
	tmpl, err := template.ParseFiles(enumTmplFile)
	jot.FatalIfErr(err)
	genSourceFile(tmpl, enumTmplFile, "enums_gen.go", edefs)
}
