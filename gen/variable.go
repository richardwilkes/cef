package main

import (
	"fmt"
	"strings"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/toolbox/txt"
)

var cNamesToPrefix = []string{"range", "select", "type"}

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
}

func newCVar(name, typeInfo string, pos position) *variable {
	name = strings.TrimSpace(name)
	typeInfo = strings.TrimPrefix(strings.Trim(strings.TrimSpace(strings.Replace(typeInfo, "const ", "", -1)), "'"), "struct _")
	if i := strings.Index(typeInfo, "':'"); i != -1 {
		typeInfo = typeInfo[:i]
	}
	v := &variable{
		Name:   name,
		GoName: txt.ToCamelCase(name),
	}
	for _, one := range cNamesToPrefix {
		if one == name {
			v.Name = "_" + name
			break
		}
	}
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
	if v.Name == "base" {
		switch typeInfo {
		case "cef_base_ref_counted_t":
			typeInfo = "cef_base_ref_counted_t *"
		case "cef_base_scoped_t":
			typeInfo = "cef_base_scoped_t *"
		default:
			jot.Fatal(1, errs.Newf("Unexpected base type: %s", typeInfo))
		}
	}
	v.CType = typeInfo
	if space := strings.Index(typeInfo, " "); space != -1 {
		v.BaseType = typeInfo[:space]
		v.Ptrs = typeInfo[space+1:]
	} else {
		v.BaseType = typeInfo
	}
	switch v.BaseType {
	case "void":
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
	case "char":
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
	case "float":
		v.GoType = v.Ptrs + "float32"
	case "double":
		v.GoType = v.Ptrs + "float64"
	case "size_t":
		v.GoType = v.Ptrs + "uint64"
	case "cef_string_t", "cef_string_userfree_t", "cef_string_userfree_utf8_t",
		"cef_string_userfree_utf16_t", "cef_string_userfree_wide_t",
		"cef_string_utf8_t", "cef_string_utf16_t", "cef_string_wide_t":
		switch v.Ptrs {
		case "", "*":
			v.GoType = "string"
		case "**":
			v.GoType = "*string"
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
					param := strings.TrimSpace(param)
					if i = strings.LastIndex(param, " "); i != -1 {
						params = append(params, param[i+1:])
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
	if v.BaseType == "void" {
		return expression
	}
	if strings.HasPrefix(v.GoType, "*") {
		return fmt.Sprintf("(%s)(%s)", v.GoType, expression)
	}
	return fmt.Sprintf("%s(%s)", v.GoType, expression)
}

func (v *variable) NameNoMangle() string {
	for _, one := range cNamesToPrefix {
		if "_"+one == v.Name {
			return one
		}
	}
	return v.Name
}
