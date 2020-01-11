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
