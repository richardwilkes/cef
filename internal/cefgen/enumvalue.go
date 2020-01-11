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
	"strconv"
)

type enumValue struct {
	Name     string
	GoName   string
	Value    string
	IsNum    bool
	Num      int
	Position position
}

func newEnumValue(name, value string, pos position) *enumValue {
	e := &enumValue{
		Name:     name,
		GoName:   translateName(name),
		Position: pos,
	}
	e.setValue(value)
	return e
}

func (e *enumValue) setValue(value string) {
	e.Value = value
	if v, err := strconv.Atoi(value); err == nil {
		e.IsNum = true
		e.Num = v
	} else {
		e.IsNum = false
	}
}
