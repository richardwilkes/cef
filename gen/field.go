package main

import (
	"fmt"
	"strings"
)

type field struct {
	Owner    *structDef
	Var      *variable
	Position position
}

func newField(owner *structDef, name, typeInfo string, pos position) *field {
	return &field{
		Owner:    owner,
		Var:      newCVar(name, typeInfo),
		Position: pos,
	}
}

func (f *field) ParameterList() string {
	var buffer strings.Builder
	if f.Var.FunctionPtr {
		for i, p := range f.Var.Params {
			if i != 0 || p.GoType != "*"+f.Owner.GoName {
				if i != 1 {
					buffer.WriteString(", ")
				}
				fmt.Fprintf(&buffer, "p%d", i)
				if i == len(f.Var.Params)-1 || f.Var.Params[i+1].GoType != p.GoType {
					fmt.Fprintf(&buffer, " %s", p.GoType)
				}
			}
		}
	}
	return buffer.String()
}

func (f *field) ProxyParameterList() string {
	var buffer strings.Builder
	if f.Var.FunctionPtr {
		for i, p := range f.Var.Params {
			if i == 0 {
				buffer.WriteString("obj")
			} else {
				fmt.Fprintf(&buffer, ", p%d", i)
			}
			fmt.Fprintf(&buffer, " %s", p.GoType)
		}
	}
	return buffer.String()
}

func (f *field) ParameterNames() string {
	var buffer strings.Builder
	for i := range f.Var.Params {
		if i == 0 {
			buffer.WriteString("d")
		} else {
			fmt.Fprintf(&buffer, ", p%d", i)
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
	var buffer strings.Builder
	for i, p := range f.Var.Params {
		if i != 0 {
			if p.BaseType == "cef_string_t" {
				fmt.Fprintf(&buffer, "var s%d C.cef_string_t\n", i)
				ptrs := p.Ptrs
				if len(ptrs) > 0 {
					ptrs = ptrs[1:]
				}
				fmt.Fprintf(&buffer, "setCEFStr(%sp%d, &s%d)\n", ptrs, i, i)
			} else if p.Ptrs == "**" {
				if sdef, exists := sdefsMap[p.BaseType]; exists {
					fmt.Fprintf(&buffer, "pd%d := (*p%d).toNative(", i, i)
					if !sdef.isClassEquivalent() {
						fmt.Fprintf(&buffer, "&C.%s{}", p.BaseType)
					}
					buffer.WriteString(")\n")
				} else if p.BaseType == "char" {
					fmt.Fprintf(&buffer, `cp%[1]d := C.calloc(C.size_t(len(p%[1]d)), C.size_t(unsafe.Sizeof(uintptr(0))))
tp%[1]d := (*[1<<30 - 1]*C.char)(cp%[1]d)
for i, one := range p%[1]d {
	tp%[1]d[i] = C.CString(one)
}
`, i)
				}
			} else if p.Ptrs == "*" {
				if _, exists := edefsMap[p.BaseType]; exists {
					fmt.Fprintf(&buffer, "e%d := C.%s(*p%d)\n", i, p.BaseType, i)
				}
			}
		}
	}
	prefixLines := buffer.String()
	buffer.Reset()
	fmt.Fprintf(&buffer, "C.%s(d.toNative()", f.TrampolineName())
	for i, p := range f.Var.Params {
		if i != 0 {
			buffer.WriteString(", ")
			if p.BaseType == "void" {
				fmt.Fprintf(&buffer, "p%d", i)
			} else if p.BaseType == "cef_string_t" && p.Ptrs == "*" {
				fmt.Fprintf(&buffer, "&s%d", i)
			} else if p.BaseType == "char" && p.Ptrs == "**" {
				fmt.Fprintf(&buffer, "(**C.char)(cp%d)", i)
			} else {
				if p.Ptrs == "*" {
					if _, exists := edefsMap[p.BaseType]; exists {
						fmt.Fprintf(&buffer, "&e%d", i)
						continue
					}
				}
				if sdef, exists := sdefsMap[p.BaseType]; exists {
					if len(p.Ptrs) > 1 {
						fmt.Fprintf(&buffer, "&pd%d", i)
					} else {
						fmt.Fprintf(&buffer, "p%d.toNative(", i)
						if !sdef.isClassEquivalent() {
							fmt.Fprintf(&buffer, "&C.%s{}", p.BaseType)
						}
						buffer.WriteString(")")
					}
				} else if len(p.Ptrs) > 0 {
					fmt.Fprintf(&buffer, "(%sC.%s)(p%d)", p.Ptrs, p.BaseType, i)
				} else {
					fmt.Fprintf(&buffer, "C.%s(p%d)", p.BaseType, i)
				}
			}
		}
	}
	fmt.Fprintf(&buffer, ", d.%s)", f.Var.Name)
	call := buffer.String()
	buffer.Reset()
	buffer.WriteString(prefixLines)
	if f.Var.GoType == "" {
		buffer.WriteString(call)
	} else {
		if sdef, exists := sdefsMap[f.Var.CType]; exists && !sdef.isClassEquivalent() {
			fmt.Fprintf(&buffer, `native := %s
var result %s
result.fromNative(&native)
return result`, call, f.Var.GoType)
		} else {
			switch f.Var.CType {
			case "cef_string_t":
				fmt.Fprintf(&buffer, "native := %s\nreturn cefstrToString(&native)", call)
			case "cef_string_t *":
				fmt.Fprintf(&buffer, "return cefstrToString(%s)", call)
			case "cef_string_userfree_t":
				fmt.Fprintf(&buffer, "return cefuserfreestrToString(%s)", call)
			default:
				buffer.WriteString("return ")
				if strings.HasPrefix(f.Var.GoType, "*") {
					fmt.Fprintf(&buffer, "(%s)", f.Var.GoType)
				} else {
					buffer.WriteString(f.Var.GoType)
				}
				fmt.Fprintf(&buffer, "(%s)", call)
			}
		}
	}
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
		if i == 0 {
			fmt.Fprintf(&buffer, "%s self", p.CType)
		} else {
			fmt.Fprintf(&buffer, ", %s p%d", p.CType, i)
		}
	}
	fmt.Fprintf(&buffer, ", %s (CEF_CALLBACK *callback)(", f.Var.CType)
	for i, p := range f.Var.Params {
		if i != 0 {
			buffer.WriteString(", ")
		}
		fmt.Fprintf(&buffer, "%s", p.CType)
	}
	buffer.WriteString(")) { return callback(")
	for i := range f.Var.Params {
		if i == 0 {
			buffer.WriteString("self")
		} else {
			fmt.Fprintf(&buffer, ", p%d", i)
		}
	}
	buffer.WriteString("); }")
	return buffer.String()
}

func (f *field) Callback() string {
	var buffer strings.Builder
	for i, p := range f.Var.Params {
		if i != 0 {
			if sdef, exists := sdefsMap[p.BaseType]; exists && !sdef.isClassEquivalent() {
				fmt.Fprintf(&buffer, "var v%s %s\n", p.Name, p.GoType)
			}
		}
	}
	prefixLines := buffer.String()
	buffer.Reset()
	fmt.Fprintf(&buffer, "proxy.%s(", f.Var.GoName)
	for i, p := range f.Var.Params {
		if i == 0 {
			buffer.WriteString("me")
		} else {
			buffer.WriteString(", ")
			if p.BaseType == "cef_string_t" {
				fmt.Fprintf(&buffer, "cefstrToString(%s)", p.Name)
			} else if sdef, exists := sdefsMap[p.BaseType]; exists && !sdef.isClassEquivalent() {
				fmt.Fprintf(&buffer, "v%[1]s.fromNative(%[1]s)", p.Name)
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
		buffer.WriteString("return ")
		buffer.WriteString(f.Var.CGoCast(call))
	}
	return buffer.String()
}

func (f *field) CallbackParams() string {
	var buffer strings.Builder
	for i, p := range f.Var.Params {
		if i == 0 {
			buffer.WriteString("self")
		} else {
			fmt.Fprintf(&buffer, ", p%d", i)
		}
		fmt.Fprintf(&buffer, " %sC.%s", p.Ptrs, p.BaseType)
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
