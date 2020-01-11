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

// Point (cef_point_t from include/internal/cef_types.h)
// Structure representing a point.
type Point struct {
	// X (x)
	X int32
	// Y (y)
	Y int32
}

// NewPoint creates a new Point.
func NewPoint() *Point {
	return &Point{}
}

func (d *Point) toNative(native *C.cef_point_t) *C.cef_point_t {
	if d == nil {
		return nil
	}
	native.x = C.int(d.X)
	native.y = C.int(d.Y)
	return native
}

func (n *C.cef_point_t) toGo() *Point {
	if n == nil {
		return nil
	}
	var d Point
	n.intoGo(&d)
	return &d
}

func (n *C.cef_point_t) intoGo(d *Point) {
	d.X = int32(n.x)
	d.Y = int32(n.y)
}
