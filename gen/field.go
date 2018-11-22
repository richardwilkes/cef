package main

import (
	"fmt"
	"strings"

	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/toolbox/txt"
)

var specialNames = map[string]bool{
	"range": true,
	"type":  true,
}

type field struct {
	Name         string
	GoName       string
	CReturnType  string
	GoReturnType string
	CParams      []string
	GoParams     []string
	FunctionPtr  bool
	NeedsUnsafe  bool
	Position     position
}

func newField(name, typeInfo string, pos position) *field {
	original := typeInfo
	f := &field{
		Name:     name,
		GoName:   txt.ToCamelCase(name),
		Position: pos,
	}
	if _, ok := specialNames[name]; ok {
		f.Name = "_" + name
	}
	fp := strings.Index(typeInfo, "(*)")
	if fp != -1 {
		f.FunctionPtr = true
		typeInfo = typeInfo[:fp]
	}
	f.CReturnType = filterCTypeName(typeInfo)
	f.GoReturnType = deriveGoTypeFromCType(f.CReturnType, &f.NeedsUnsafe)
	if f.FunctionPtr {
		params := original[fp+3:]
		if !strings.HasPrefix(params, "(") || !strings.HasSuffix(params, ")") {
			jot.Fatalf(1, "Can't handle params in type: %s", original)
		}
		f.CParams = strings.Split(params[1:len(params)-1], ",")
		for i := range f.CParams {
			f.CParams[i] = filterCTypeName(f.CParams[i])
			f.GoParams = append(f.GoParams, deriveGoTypeFromCType(f.CParams[i], &f.NeedsUnsafe))
		}
	}
	return f
}

func filterCTypeName(in string) string {
	in = strings.TrimSpace(in)
	if strings.HasPrefix(in, "const struct _cef_") {
		in = in[14:]
	} else if strings.HasPrefix(in, "struct _cef_") {
		in = in[8:]
	}
	if i := strings.Index(in, "':'"); i != -1 {
		in = in[:i]
	}
	return in
}

type cMapping struct {
	C  string
	Go string
}

var cMappings = []cMapping{
	// RAW: Look at the cef_base_* mappings...
	{C: "cef_base_scoped_t", Go: "uintptr"},
	{C: "cef_base_ref_counted_t", Go: "uintptr"},
	{C: "cef_string_t", Go: "string"},
	{C: "cef_string_userfree_t", Go: "string"},
	{C: "cef_string_userfree_utf8_t", Go: "string"},
	{C: "cef_string_userfree_utf16_t", Go: "string"},
	{C: "cef_string_userfree_wide_t", Go: "string"},
	{C: "cef_string_utf8_t", Go: "string"},
	{C: "cef_string_utf16_t", Go: "string"},
	{C: "cef_string_wide_t", Go: "string"},
	{C: "size_t", Go: "int64"},
	{C: "int", Go: "int32"},
	{C: "float", Go: "float32"},
	{C: "double", Go: "float64"},
	{C: "char", Go: "byte"},
	{C: "char16", Go: "int16"},
	{C: "wchar_t", Go: "int16"},
}

func deriveGoTypeFromCType(in string, needsUnsafe *bool) string {
	in = strings.Replace(in, "const ", "", -1)
	if in == "void" {
		return ""
	}
	if in == "void *" || in == "void **" {
		*needsUnsafe = true
		return "unsafe.Pointer"
	}
	prefix := ""
	if i := strings.Index(in, " "); i != -1 {
		prefix = in[i+1:]
		in = in[:i]
	}
	for _, one := range cMappings {
		if one.C == in {
			return prefix + one.Go
		}
	}
	return prefix + translateStructTypeName(in)
}

func (f *field) Skip() bool {
	return f.Name == "base" && (f.CReturnType == "cef_base_ref_counted_t" || f.CReturnType == "cef_base_scoped_t")
}

func (f *field) ParameterList(d *structDef) string {
	var buffer strings.Builder
	if f.FunctionPtr {
		count := 0
		for i, p := range f.GoParams {
			if i != 0 || p != "*"+d.GoName {
				count++
				if count != 1 {
					buffer.WriteString(", ")
				}
				fmt.Fprintf(&buffer, "p%d", count)
				if i == len(f.GoParams)-1 || f.GoParams[i+1] != p {
					fmt.Fprintf(&buffer, " %s", p)
				}
			}
		}
	}
	return buffer.String()
}

func (f *field) GoReturn() string {
	var buffer strings.Builder
	if sdef, exists := sdefsMap[f.CReturnType]; exists && !sdef.isClassEquivalent() {
		fmt.Fprintf(&buffer, `var result C.%s // RAW: Need impl
	var goResult %s
	goResult.fromNative(&result)
	return goResult
`, f.CReturnType, sdef.GoName)
	} else {
		buffer.WriteString("return ")
		switch f.GoReturnType {
		case "unsafe.Pointer":
			buffer.WriteString("nil // RAW: Need impl")
		case "string":
			switch f.CReturnType {
			case "cef_string_t":
				fmt.Fprintf(&buffer, "cefstrToString(&d.native.%s)", f.Name)
			case "cef_string_t *":
				fmt.Fprintf(&buffer, "cefstrToString(d.native.%s)", f.Name)
			default:
				buffer.WriteString(`"" // RAW: Need impl`)
			}
		default:
			if f.FunctionPtr {
				if strings.HasPrefix(f.GoReturnType, "*") {
					buffer.WriteString("nil // RAW: Need impl")
				} else {
					buffer.WriteString("0 // RAW: Need impl")
				}
			} else {
				if strings.HasPrefix(f.GoReturnType, "*") {
					fmt.Fprintf(&buffer, "(%s)", f.GoReturnType)
				} else {
					buffer.WriteString(f.GoReturnType)
				}
				fmt.Fprintf(&buffer, "(d.native.%s)", f.Name)
			}
		}
	}
	return buffer.String()
}

func (f *field) ToNative() string {
	var buffer strings.Builder
	if sdef, exists := sdefsMap[f.CReturnType]; exists && !sdef.isClassEquivalent() {
		fmt.Fprintf(&buffer, "d.%s.toNative(&native.%s)", f.GoName, f.Name)
	} else {
		switch f.CReturnType {
		case "void *":
			fmt.Fprintf(&buffer, "native.%s = d.%s", f.Name, f.GoName)
		case "cef_string_t":
			fmt.Fprintf(&buffer, `setCEFStr(d.%s, &native.%s)`, f.GoName, f.Name)
		case "cef_string_t *":
			fmt.Fprintf(&buffer, `setCEFStr(d.%s, native.%s)`, f.GoName, f.Name)
		default:
			fmt.Fprintf(&buffer, "native.%s = ", f.Name)
			if i := strings.Index(f.CReturnType, " "); i != -1 {
				fmt.Fprintf(&buffer, "(%sC.%s)", f.CReturnType[i+1:], f.CReturnType[:i])
			} else {
				fmt.Fprintf(&buffer, "C.%s", f.CReturnType)
			}
			fmt.Fprintf(&buffer, "(d.%s)", f.GoName)
		}
	}
	return buffer.String()
}

func (f *field) FromNative() string {
	var buffer strings.Builder
	if sdef, exists := sdefsMap[f.CReturnType]; exists && !sdef.isClassEquivalent() {
		fmt.Fprintf(&buffer, "d.%s.fromNative(&native.%s)", f.GoName, f.Name)
	} else {
		fmt.Fprintf(&buffer, "d.%s = ", f.GoName)
		switch f.CReturnType {
		case "cef_string_t":
			fmt.Fprintf(&buffer, "cefstrToString(&native.%s)", f.Name)
		case "cef_string_t *":
			fmt.Fprintf(&buffer, "cefstrToString(native.%s)", f.Name)
		default:
			if strings.HasPrefix(f.GoReturnType, "*") {
				fmt.Fprintf(&buffer, "(%s)", f.GoReturnType)
			} else {
				buffer.WriteString(f.GoReturnType)
			}
			fmt.Fprintf(&buffer, "(native.%s)", f.Name)
		}
	}
	return buffer.String()
}
