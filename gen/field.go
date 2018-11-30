package main

import (
	"fmt"
	"strings"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

type field struct {
	Owner    *structDef
	Var      *variable
	Position position
}

func newField(owner *structDef, name, typeInfo string, pos position) *field {
	return &field{
		Owner:    owner,
		Var:      newCVar(name, typeInfo, pos),
		Position: pos,
	}
}

func (f *field) ParameterList() string {
	if f.Var.FunctionPtr {
		params := f.Var.Params
		if len(params) > 0 && params[0].GoType == "*"+f.Owner.GoName {
			params = params[1:]
		}
		return parameterList(params)
	}
	return ""
}

func (f *field) ProxyParameterList() string {
	if f.Var.FunctionPtr {
		return parameterList(f.Var.Params)
	}
	return ""
}

func (f *field) ParameterNames() string {
	var buffer strings.Builder
	for i, p := range f.Var.Params {
		if i == 0 {
			buffer.WriteString("d")
		} else {
			fmt.Fprintf(&buffer, ", %s", p.Name)
		}
	}
	return buffer.String()
}

type param struct {
	Type string
	Ptrs string
}

func (f *field) processedCParams(in []string) []param {
	params := make([]param, len(in))
	for i, p := range in {
		p = strings.Replace(p, "const ", "", -1)
		if space := strings.Index(p, " "); space != -1 {
			params[i] = param{
				Type: p[:space],
				Ptrs: p[space+1:],
			}
		} else {
			params[i] = param{Type: p}
		}
	}
	return params
}

func (f *field) CallFunctionPointer() string {
	prep, names := prepGoVarsForC(f.Var.Params)
	var buffer strings.Builder
	fmt.Fprintf(&buffer, "C.%s(d.toNative()", f.TrampolineName())
	emitParamsForCCall(&buffer, f.Var.Params[1:], names[1:], false)
	fmt.Fprintf(&buffer, ", d.%s)", f.Var.Name)
	expression := buffer.String()
	buffer.Reset()
	buffer.WriteString(prep)
	emitReturnForCCall(&buffer, expression, f.Var)
	return buffer.String()
}

func (f *field) ReturnField() string {
	var buffer strings.Builder
	buffer.WriteString("return ")
	if strings.HasPrefix(f.Var.GoType, "*") {
		fmt.Fprintf(&buffer, "(%s)", f.Var.GoType)
	} else {
		buffer.WriteString(f.Var.GoType)
	}
	buffer.WriteString("(")
	if f.Var.Name == "base" {
		buffer.WriteString("&")
	}
	fmt.Fprintf(&buffer, "d.%s)", f.Var.Name)
	return buffer.String()
}

func (f *field) ToNative() string {
	var buffer strings.Builder
	if sdef, exists := sdefsMap[f.Var.CType]; exists && !sdef.isClassEquivalent() {
		fmt.Fprintf(&buffer, "d.%s.toNative(&native.%s)", f.Var.GoName, f.Var.Name)
	} else {
		switch f.Var.CType {
		case "void *":
			fmt.Fprintf(&buffer, "native.%s = d.%s", f.Var.Name, f.Var.GoName)
		case "cef_string_t":
			fmt.Fprintf(&buffer, `setCEFStr(d.%s, &native.%s)`, f.Var.GoName, f.Var.Name)
		case "cef_string_t *":
			fmt.Fprintf(&buffer, `setCEFStr(d.%s, native.%s)`, f.Var.GoName, f.Var.Name)
		default:
			fmt.Fprintf(&buffer, "native.%s = ", f.Var.Name)
			if i := strings.Index(f.Var.CType, " "); i != -1 {
				fmt.Fprintf(&buffer, "(%sC.%s)", f.Var.CType[i+1:], f.Var.CType[:i])
			} else {
				fmt.Fprintf(&buffer, "C.%s", f.Var.CType)
			}
			fmt.Fprintf(&buffer, "(d.%s)", f.Var.GoName)
		}
	}
	return buffer.String()
}

func (f *field) FromNative() string {
	var buffer strings.Builder
	if sdef, exists := sdefsMap[f.Var.CType]; exists && !sdef.isClassEquivalent() {
		fmt.Fprintf(&buffer, "d.%s.fromNative(&native.%s)", f.Var.GoName, f.Var.Name)
	} else {
		fmt.Fprintf(&buffer, "d.%s = ", f.Var.GoName)
		switch f.Var.CType {
		case "void *":
			fmt.Fprintf(&buffer, "native.%s", f.Var.Name)
		case "cef_string_t":
			fmt.Fprintf(&buffer, "cefstrToString(&native.%s)", f.Var.Name)
		case "cef_string_t *":
			fmt.Fprintf(&buffer, "cefstrToString(native.%s)", f.Var.Name)
		default:
			if strings.HasPrefix(f.Var.GoType, "*") {
				fmt.Fprintf(&buffer, "(%s)", f.Var.GoType)
			} else {
				buffer.WriteString(f.Var.GoType)
			}
			fmt.Fprintf(&buffer, "(native.%s)", f.Var.Name)
		}
	}
	return buffer.String()
}

func (f *field) TrampolineName() string {
	return "gocef_" + f.TrampolineNameNoGoCEF()
}

func (f *field) TrampolineNameNoGoCEF() string {
	return fmt.Sprintf("%s_%s", strings.TrimSuffix(strings.TrimPrefix(f.Owner.Name, "cef_"), "_t"), f.Var.Name)
}

func (f *field) Trampoline() string {
	if !f.Var.FunctionPtr {
		return ""
	}
	var buffer strings.Builder
	fmt.Fprintf(&buffer, "%s %s(", f.Var.CType, f.TrampolineName())
	for i, p := range f.Var.Params {
		if i != 0 {
			buffer.WriteString(", ")
		}
		fmt.Fprintf(&buffer, "%s %s", p.CType, p.Name)
	}
	fmt.Fprintf(&buffer, ", %s (CEF_CALLBACK *callback__)(", f.Var.CType)
	for i, p := range f.Var.Params {
		if i != 0 {
			buffer.WriteString(", ")
		}
		fmt.Fprintf(&buffer, "%s", p.CType)
	}
	buffer.WriteString(")) { return callback__(")
	for i, p := range f.Var.Params {
		if i != 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString(p.Name)
	}
	buffer.WriteString("); }")
	return buffer.String()
}

func (f *field) Callback() string {
	var buffer strings.Builder
	pval := make([]string, len(f.Var.Params))
	for i, p := range f.Var.Params {
		if i != 0 {
			if sdef, exists := sdefsMap[p.BaseType]; exists {
				if sdef.isClassEquivalent() {
					if p.Ptrs == "**" {
						fmt.Fprintf(&buffer, "v%s := (%s)(*%s)\n", p.Name, p.GoType[1:], p.Name)
						pval[i] = fmt.Sprintf("&v%s", p.Name)
					} else if len(p.Ptrs) > 2 {
						jot.Fatal(1, errs.Newf("Unhandled param conversion: %s", p.Name))
					}
				} else {
					if p.Ptrs == "*" {
						fmt.Fprintf(&buffer, "var v%s %s\n", p.Name, p.GoType[1:])
						pval[i] = fmt.Sprintf("v%[1]s.fromNative(%[1]s)", p.Name)
					} else {
						jot.Fatal(1, errs.Newf("Unhandled param conversion: %s", p.Name))
					}
				}
			} else if _, exists := edefsMap[p.BaseType]; exists {
				if p.Ptrs == "*" {
					fmt.Fprintf(&buffer, "e%s := %s(*%s)\n", p.Name, p.GoType[1:], p.Name)
					pval[i] = fmt.Sprintf("&e%s", p.Name)
				} else if len(p.Ptrs) > 1 {
					jot.Fatal(1, errs.Newf("Unhandled param conversion: %s", p.Name))
				}
			}
		}
	}
	prefixLines := buffer.String()
	buffer.Reset()
	fmt.Fprintf(&buffer, "proxy__.%s(", f.Var.GoName)
	for i, p := range f.Var.Params {
		if i == 0 {
			buffer.WriteString("me__")
		} else {
			buffer.WriteString(", ")
			if pval[i] != "" {
				buffer.WriteString(pval[i])
			} else if p.BaseType == "cef_string_t" {
				fmt.Fprintf(&buffer, "cefstrToString(%s)", p.Name)
			} else {
				buffer.WriteString(p.GoCast(p.Name))
			}
		}
	}
	buffer.WriteString(")")
	call := buffer.String()
	buffer.Reset()
	buffer.WriteString(prefixLines)
	if f.Var.GoType == "" {
		buffer.WriteString(call)
	} else {
		if _, exists := sdefsMap[f.Var.BaseType]; exists && f.Var.Ptrs == "" {
			fmt.Fprintf(&buffer, "call__ := %s\n", call)
			fmt.Fprintf(&buffer, "var result__ C.%s\n", f.Var.CType)
			buffer.WriteString("call__.toNative(&result__)\n")
			buffer.WriteString("return result__")
		} else {
			fmt.Fprintf(&buffer, "return %s", f.Var.CGoCast(call))
		}
	}
	return buffer.String()
}

func (f *field) CallbackParams() string {
	var buffer strings.Builder
	for i, p := range f.Var.Params {
		if i != 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString(p.Name)
		if p.BaseType == "void" {
			switch p.Ptrs {
			case "":
			case "*":
				buffer.WriteString(" unsafe.Pointer")
			case "**":
				buffer.WriteString(" *unsafe.Pointer")
			default:
				jot.Fatal(1, errs.Newf("Unhandled void case: %s", p.Ptrs))
			}
		} else {
			fmt.Fprintf(&buffer, " %sC.%s", p.Ptrs, p.BaseType)
		}
	}
	return buffer.String()
}

func (f *field) CallbackReturnType() string {
	if f.Var.GoType == "" {
		return ""
	}
	rt := f.processedCParams([]string{f.Var.CType})[0]
	return fmt.Sprintf("%sC.%s", rt.Ptrs, rt.Type)
}
