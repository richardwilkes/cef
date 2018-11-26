package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/toolbox/txt"
)

var typedefNamingRegex = regexp.MustCompile(`^.+ (.+) '([^']+)'`)

type typeDef struct {
	Name     string
	GoName   string
	RefersTo string
	Position position
}

func newTypeDef(name, refersTo string, pos position) *typeDef {
	return &typeDef{
		Name:     name,
		GoName:   translateStructTypeName(name),
		RefersTo: translateRefersToType(refersTo),
		Position: pos,
	}
}

var excludeMap = map[string]bool{
	"cef_string_t":                true,
	"cef_string_userfree_t":       true,
	"cef_string_userfree_utf8_t":  true,
	"cef_string_userfree_utf16_t": true,
	"cef_string_userfree_wide_t":  true,
	"cef_string_utf8_t":           true,
	"cef_string_utf16_t":          true,
	"cef_string_wide_t":           true,
}

func processTypedefDecl(block []lineInfo) {
	if result := typedefNamingRegex.FindAllStringSubmatch(block[0].Line, -1); result != nil {
		name := result[0][1]
		if strings.HasPrefix(name, "cef_") {
			if _, exclude := excludeMap[name]; !exclude {
				if _, exists := sdefsMap[name]; !exists {
					if _, exists = edefsMap[name]; !exists {
						if _, exists = tdefsMap[name]; !exists {
							tdefsMap[name] = newTypeDef(name, result[0][2], block[0].Position)
						}
					}
				}
			}
		}
	}
}

func dumpTypedefs() {
	type info struct {
		Headers []string
		Types   []*typeDef
	}
	var data info
	headerMap := make(map[string]bool)
	data.Types = make([]*typeDef, 0, len(tdefsMap))
	for _, tdef := range tdefsMap {
		data.Types = append(data.Types, tdef)
		headerMap[tdef.Position.Src] = true
	}
	sort.Slice(data.Types, func(i, j int) bool { return txt.NaturalLess(data.Types[i].Name, data.Types[j].Name, true) })
	data.Headers = make([]string, 0, len(headerMap))
	for h := range headerMap {
		data.Headers = append(data.Headers, h)
	}
	sort.Slice(data.Headers, func(i, j int) bool { return txt.NaturalLess(data.Headers[i], data.Headers[j], true) })
	const typedefTmplFile = "typedef.go.tmpl"
	tmpl, err := template.ParseFiles(typedefTmplFile)
	jot.FatalIfErr(err)
	genSourceFile(tmpl, typedefTmplFile, "typedefs_gen.go", &data)
}

func translateRefersToType(in string) string {
	switch in {
	case "void *":
		return "uintptr"
	case "uint32":
		return in
	case "char16":
		return "uint16"
	default:
		if i := strings.Index(in, " "); i != -1 {
			return fmt.Sprintf("%sC.%s", in[i+1:], in[:i])
		}
		return "C." + in
	}
}
