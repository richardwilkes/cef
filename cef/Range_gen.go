// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Code created from "struct.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	"C"
)

// Range (cef_range_t from include/internal/cef_types.h)
// Structure representing a range.
type Range struct {
	// From (from)
	From int32
	// To (to)
	To int32
}

// NewRange creates a new Range.
func NewRange() *Range {
	return &Range{}
}

func (d *Range) toNative(native *C.cef_range_t) *C.cef_range_t {
	if d == nil {
		return nil
	}
	native.from = C.int(d.From)
	native.to = C.int(d.To)
	return native
}

func (n *C.cef_range_t) toGo() *Range {
	if n == nil {
		return nil
	}
	var d Range
	n.intoGo(&d)
	return &d
}

func (n *C.cef_range_t) intoGo(d *Range) {
	d.From = int32(n.from)
	d.To = int32(n.to)
}
