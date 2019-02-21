package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/toolbox/txt"
)

const (
	voidType         = "void"
	voidPtrType      = "void *"
	charType         = "char"
	cefStringType    = "cef_string_t"
	cefStringPtrType = "cef_string_t *"
	stringPtrGoType  = "*string"
)

var (
	cNamesToPrefixForAccess = []string{"range", "type"}
	paramRenames            = []string{"chan", "defer", "error", "fallthrough", "func", "go", "import", "interface", "map", "package", "range", "select", "string", "type", "var"}
)

type variable struct {
	Name        string
	BaseType    string
	Ptrs        string
	CType       string
	GoName      string
	GoType      string
	Params      []*variable
	FunctionPtr bool
	NeedUnsafe  bool
	HadConst    bool
}

func newCVar(name, typeInfo string, pos position) *variable {
	name = strings.TrimSpace(name)
	v := &variable{
		Name:   name,
		GoName: txt.ToCamelCase(name),
	}
	for _, one := range cNamesToPrefixForAccess {
		if one == name {
			v.Name = "_" + name
			break
		}
	}

	typeInfo = txt.CollapseSpaces(strings.TrimSpace(typeInfo))
	fp := strings.Index(typeInfo, "(*)")
	if fp != -1 {
		v.FunctionPtr = true
		params := strings.TrimSpace(typeInfo[fp+3:])
		if !strings.HasPrefix(params, "(") || !strings.HasSuffix(params, ")") {
			jot.Fatal(1, errs.Newf("Can't handle params in type: %s", typeInfo))
		}
		paramNames := extractParameterNames(pos)
		paramList := strings.Split(params[1:len(params)-1], ",")
		if len(paramNames) != len(paramList) {
			jot.Fatal(1, errs.Newf("Extracted param names (%v) don't match params (%v) for %s: %v", paramNames, paramList, name, pos))
		}
		typeInfo = strings.TrimSpace(typeInfo[:fp])
		for i, param := range paramList {
			one := newCVar(paramNames[i], param, pos)
			if one.NeedUnsafe {
				v.NeedUnsafe = true
			}
			v.Params = append(v.Params, one)
		}
	}

	if v.HadConst = strings.Contains(typeInfo, "const "); v.HadConst {
		typeInfo = strings.Replace(typeInfo, "const ", "", -1)
	}
	typeInfo = strings.TrimSpace(typeInfo)
	typeInfo = strings.TrimPrefix(typeInfo, "struct _")
	typeInfo = strings.Replace(typeInfo, "long long", "longlong", -1)
	if v.Name == baseFieldName {
		typeInfo += " *"
	}
	v.CType = typeInfo
	if space := strings.Index(typeInfo, " "); space != -1 {
		v.BaseType = typeInfo[:space]
		v.Ptrs = typeInfo[space+1:]
	} else {
		v.BaseType = typeInfo
	}
	switch v.BaseType {
	case voidType:
		switch v.Ptrs {
		case "":
		case "*":
			v.NeedUnsafe = true
			v.GoType = "unsafe.Pointer"
		case "**":
			v.NeedUnsafe = true
			v.GoType = "*unsafe.Pointer"
		default:
			jot.Fatal(1, errs.Newf("Unhandled void case: %s", v.Ptrs))
		}
	case charType:
		if v.Ptrs == "**" {
			v.NeedUnsafe = true
			v.GoType = "[]string"
		} else {
			jot.Fatal(1, errs.Newf("Unhandled char case: %s", v.Ptrs))
		}
	case "char16":
		v.GoType = v.Ptrs + "int16"
	case "int":
		v.GoType = v.Ptrs + "int32"
	case "int64_t", "time_t", "longlong":
		v.GoType = v.Ptrs + "int64"
	case "float":
		v.GoType = v.Ptrs + "float32"
	case "double":
		v.GoType = v.Ptrs + "float64"
	case "size_t":
		v.GoType = v.Ptrs + "uint64"
	case cefStringType, "cef_string_userfree_t", "cef_string_userfree_utf8_t",
		"cef_string_userfree_utf16_t", "cef_string_userfree_wide_t",
		"cef_string_utf8_t", "cef_string_utf16_t", "cef_string_wide_t":
		switch v.Ptrs {
		case "", "*":
			v.GoType = "string"
		case "**":
			v.GoType = stringPtrGoType
		default:
			jot.Fatal(1, errs.Newf("Unhandled string case: %s", v.Ptrs))
		}
	default:
		v.GoType = v.Ptrs + translateStructTypeName(v.BaseType)
	}
	return v
}

func extractParameterNames(pos position) []string {
	var buffer strings.Builder
	for i := pos.LineStart; i <= pos.LineEnd; i++ {
		if i != pos.LineStart {
			buffer.WriteString(" ")
		}
		buffer.WriteString(pos.Text(i, 1, 100000))
	}
	line := buffer.String()
	if i := strings.Index(line, "("); i != -1 {
		line = line[i+1:]
		if i = strings.Index(line, "("); i != -1 {
			line = line[i+1:]
			if i = strings.Index(line, ")"); i != -1 {
				var params []string
				line = txt.CollapseSpaces(strings.TrimSpace(line[:i]))
				for j, param := range strings.Split(line, ",") {
					param = strings.TrimSpace(param)
					if i = strings.LastIndex(param, " "); i != -1 {
						params = append(params, adjustedParamName(param[i+1:]))
					} else {
						jot.Fatal(1, errs.Newf("Unable to extract parameter name %d from: %s", j, line))
					}
				}
				return params
			}
		}
	}
	jot.Fatal(1, errs.Newf("Unable to extract parameter names from: %s", line))
	return nil
}

func adjustedParamName(name string) string {
	for _, one := range paramRenames {
		if name == one {
			return name + "_r"
		}
	}
	return name
}

func (v *variable) paramGoType() string {
	if !v.HadConst && v.GoType == "string" {
		return stringPtrGoType
	}
	return v.GoType
}

func (v *variable) transformCToGo(w io.Writer) string {
	if sdef, exists := sdefsMap[v.BaseType]; exists {
		if sdef.isClassEquivalent() {
			switch v.Ptrs {
			case "", "*":
				return v.GoCast(v.Name)
			case "**":
				fmt.Fprintf(w, "%[1]s_ := (%[2]s)(*%[1]s)\n", v.Name, v.GoType[1:])
				fmt.Fprintf(w, "%[1]s__p := &%[1]s_\n", v.Name)
				return fmt.Sprintf("%s__p", v.Name)
			}
		} else if v.Ptrs == "*" {
			fmt.Fprintf(w, "%[1]s_ := %[1]s.toGo()\n", v.Name)
			return fmt.Sprintf("%s_", v.Name)
		}
	} else {
		_, exists := edefsMap[v.BaseType]
		switch {
		case exists:
			switch v.Ptrs {
			case "":
				return v.GoCast(v.Name)
			case "*":
				fmt.Fprintf(w, "%[1]s_ := %[2]s(*%[1]s)\n", v.Name, v.GoType[1:])
				return fmt.Sprintf("&%s_", v.Name)
			}
		case v.BaseType == cefStringType:
			addressOf := ""
			if !v.HadConst {
				addressOf = "&"
			}
			fmt.Fprintf(w, "%[1]s_ := cefstrToString(%[1]s)\n", v.Name)
			return fmt.Sprintf("%s%s_", addressOf, v.Name)
		default:
			return v.GoCast(v.Name)
		}
	}
	jot.Fatal(1, errs.Newf("Unhandled param conversion: %s", v.Name))
	return ""
}

func (v *variable) CGoCast(expression string) string {
	if sdef, exists := sdefsMap[v.BaseType]; exists {
		if sdef.isClassEquivalent() {
			return fmt.Sprintf("(%s).toNative()", expression)
		}
		return fmt.Sprintf("(%s).toNative(&C.%s{})", expression, v.BaseType)
	}
	if strings.HasPrefix(v.Ptrs, "*") {
		return fmt.Sprintf("(%sC.%s)(%s)", v.Ptrs, v.BaseType, expression)
	}
	return fmt.Sprintf("C.%s(%s)", v.BaseType, expression)
}

func (v *variable) GoCast(expression string) string {
	if v.BaseType == voidType {
		return expression
	}
	if strings.HasPrefix(v.GoType, "*") {
		return fmt.Sprintf("(%s)(%s)", v.GoType, expression)
	}
	return fmt.Sprintf("%s(%s)", v.GoType, expression)
}

func (v *variable) NameNoMangle() string {
	for _, one := range cNamesToPrefixForAccess {
		if "_"+one == v.Name {
			return one
		}
	}
	return v.Name
}
