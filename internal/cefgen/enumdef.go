// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

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
						if strings.Contains(line, "-DeclRefExpr ") || (v[0] < '0' || v[0] > '9') {
							v = translateName(v)
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
