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

var funcRegex = regexp.MustCompile(`.*>\s+\S+\s+(cef_\S+)\s+'([^(]+)`)
var paramRegex = regexp.MustCompile(`.*-ParmVarDecl\s+[^>]+>\s+\S+\s+(\S+)\s+'([^']+)`)

type funcDef struct {
	Name     string
	GoName   string
	Return   *variable
	Params   []*variable
	Position position
}

func (f *funcDef) ParameterList() string {
	return parameterList(f.Params)
}

func parameterList(params []*variable) string {
	var buffer strings.Builder
	for i, p := range params {
		if i == 0 {
			buffer.WriteString(p.Name)
		} else {
			fmt.Fprintf(&buffer, ", %s", p.Name)
		}
		if i == len(params)-1 || p.paramGoType() != params[i+1].paramGoType() {
			fmt.Fprintf(&buffer, " %s", p.paramGoType())
		}
	}
	return buffer.String()
}

func (f *funcDef) Body() string {
	prep, names := prepGoVarsForC(f.Params)
	var buffer strings.Builder
	fmt.Fprintf(&buffer, "C.%s(", f.Name)
	emitParamsForCCall(&buffer, f.Params, names, true)
	buffer.WriteString(")")
	expression := buffer.String()
	buffer.Reset()
	buffer.WriteString(prep)
	emitReturnForCCall(&buffer, expression, f.Return)
	return buffer.String()
}

func prepGoVarsForC(vars []*variable) (prep string, names []string) {
	var buffer strings.Builder
	names = make([]string, len(vars))
	for i, p := range vars {
		names[i] = p.Name
		switch {
		case p.BaseType == cefStringType:
			names[i] = p.Name + "_"
			fmt.Fprintf(&buffer, "%s_ := C.cef_string_userfree_alloc()\n", p.Name)
			ptrs := p.Ptrs
			if !p.HadConst && p.paramGoType() == stringPtrGoType {
				ptrs = "*"
			} else if len(ptrs) > 0 {
				ptrs = ptrs[1:]
			}
			fmt.Fprintf(&buffer, "setCEFStr(%[1]s%[2]s, %[2]s_)\n", ptrs, p.Name)
			buffer.WriteString("defer func() {\n")
			if !p.HadConst && p.paramGoType() == stringPtrGoType {
				fmt.Fprintf(&buffer, "*%[1]s = cefstrToString(%[1]s_)\n", p.Name)
			}
			fmt.Fprintf(&buffer, "C.cef_string_userfree_free(%s_)\n}()\n", p.Name)
		case p.Ptrs == "**":
			if sdef, exists := sdefsMap[p.BaseType]; exists {
				names[i] = p.Name + "_"
				fmt.Fprintf(&buffer, "%[1]s_ := (*%[1]s).toNative(", p.Name)
				if !sdef.isClassEquivalent() {
					fmt.Fprintf(&buffer, "&C.%s{}", p.BaseType)
				}
				buffer.WriteString(")\n")
			} else if p.BaseType == charType {
				names[i] = p.Name + "_"
				fmt.Fprintf(&buffer, `%[1]s_ := C.calloc(C.size_t(len(%[1]s)), C.size_t(unsafe.Sizeof(uintptr(0))))
%[1]s_p := (*[1<<30 - 1]*C.char)(%[1]s_)
for i, one := range %[1]s {
	%[1]s_p[i] = C.CString(one)
}
`, p.Name)
			}
		case p.Ptrs == "*":
			if _, exists := edefsMap[p.BaseType]; exists {
				names[i] = p.Name + "_"
				fmt.Fprintf(&buffer, "%[1]s_ := C.%[2]s(*%[1]s)\n", p.Name, p.BaseType)
			} else if p.Name == "delegate" {
				names[i] = p.Name + "_"
				fmt.Fprintf(&buffer, "var delegate_ %sC.%s\n", p.Ptrs, p.BaseType)
				buffer.WriteString("if delegate != nil {\n")
				buffer.WriteString("delegate_ = delegate.toNative(")
				if sdef, exists2 := sdefsMap[p.BaseType]; exists2 && !sdef.isClassEquivalent() {
					fmt.Fprintf(&buffer, "&C.%s{}", p.BaseType)
				}
				buffer.WriteString(")\n")
				buffer.WriteString("}\n")
			}
		default:
		}
	}
	return buffer.String(), names
}

func emitReturnForCCall(buffer *strings.Builder, expression string, retVar *variable) {
	if retVar.GoType == "" {
		buffer.WriteString(expression)
	} else if sdef, exists := sdefsMap[retVar.BaseType]; exists && !sdef.isClassEquivalent() {
		if retVar.Ptrs == "*" {
			fmt.Fprintf(buffer, "return (%s).toGo()", expression)
		} else {
			fmt.Fprintf(buffer, `cresult__ := %s
var result__ %s
(&cresult__).intoGo(&result__)
return result__`, expression, retVar.GoType)
		}
	} else {
		switch retVar.CType {
		case voidPtrType:
			fmt.Fprintf(buffer, "return %s", expression)
		case cefStringType:
			fmt.Fprintf(buffer, `native__ := %s
return cefstrToString(&native__)`, expression)
		case cefStringPtrType:
			fmt.Fprintf(buffer, "return cefstrToString(%s)", expression)
		case "cef_string_userfree_t":
			fmt.Fprintf(buffer, "return cefuserfreestrToString(%s)", expression)
		default:
			buffer.WriteString("return ")
			if strings.HasPrefix(retVar.GoType, "*") {
				fmt.Fprintf(buffer, "(%s)", retVar.GoType)
			} else {
				buffer.WriteString(retVar.GoType)
			}
			fmt.Fprintf(buffer, "(%s)", expression)
		}
	}
}

func emitParamsForCCall(buffer *strings.Builder, vars []*variable, names []string, omitFirstComma bool) {
	for i, p := range vars {
		if !omitFirstComma || i != 0 {
			buffer.WriteString(", ")
		}
		switch {
		case p.BaseType == voidType:
			buffer.WriteString(names[i])
		case p.BaseType == charType && p.Ptrs == "**":
			fmt.Fprintf(buffer, "(**C.char)(%s)", names[i])
		default:
			if p.Ptrs == "*" {
				if _, exists := edefsMap[p.BaseType]; exists {
					fmt.Fprintf(buffer, "&%s", names[i])
					continue
				}
			}
			sdef, exists := sdefsMap[p.BaseType]
			switch {
			case exists:
				if len(p.Ptrs) > 1 {
					fmt.Fprintf(buffer, "&%s", names[i])
				} else {
					if names[i] == "delegate_" {
						buffer.WriteString("delegate_")
					} else {
						fmt.Fprintf(buffer, "%s.toNative(", names[i])
						if !sdef.isClassEquivalent() {
							fmt.Fprintf(buffer, "&C.%s{}", p.BaseType)
						}
						buffer.WriteString(")")
					}
				}
			case len(p.Ptrs) > 0:
				fmt.Fprintf(buffer, "(%sC.%s)(%s)", p.Ptrs, p.BaseType, names[i])
			default:
				fmt.Fprintf(buffer, "C.%s(%s)", p.BaseType, names[i])
			}
		}
	}
}

func processFunctionDecl(block []lineInfo) {
	if block[0].Position.Src != "include/internal/cef_string_types.h" {
		if result := funcRegex.FindAllStringSubmatch(block[0].Line, -1); len(result) > 0 {
			name := result[0][1]
			if _, exclude := fdefsMap[name]; !exclude {
				fdef := &funcDef{
					Name:     name,
					GoName:   translateName(name[4:]),
					Return:   newCVar("result", result[0][2], block[0].Position),
					Position: block[0].Position,
				}
				for _, b := range block {
					if result = paramRegex.FindAllStringSubmatch(b.Line, -1); len(result) > 0 {
						fdef.Params = append(fdef.Params, newCVar(adjustedParamName(result[0][1]), result[0][2], b.Position))
					}
				}
				fdefsMap[name] = fdef
			}
		}
	}
}

func dumpFunctions() {
	funcs := make([]*funcDef, 0, len(fdefsMap))
	for _, fdef := range fdefsMap {
		funcs = append(funcs, fdef)
	}
	sort.Slice(funcs, func(i, j int) bool { return txt.NaturalLess(funcs[i].Name, funcs[j].Name, true) })
	const funcdefTmplFile = "funcdef.go.tmpl"
	tmpl, err := template.ParseFiles(funcdefTmplFile)
	jot.FatalIfErr(err)
	genSourceFile(tmpl, funcdefTmplFile, "functions_gen.go", funcs)
}
