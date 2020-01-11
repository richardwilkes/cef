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

// Rect (cef_rect_t from include/internal/cef_types.h)
// Structure representing a rectangle.
type Rect struct {
	// X (x)
	X int32
	// Y (y)
	Y int32
	// Width (width)
	Width int32
	// Height (height)
	Height int32
}

// NewRect creates a new Rect.
func NewRect() *Rect {
	return &Rect{}
}

func (d *Rect) toNative(native *C.cef_rect_t) *C.cef_rect_t {
	if d == nil {
		return nil
	}
	native.x = C.int(d.X)
	native.y = C.int(d.Y)
	native.width = C.int(d.Width)
	native.height = C.int(d.Height)
	return native
}

func (n *C.cef_rect_t) toGo() *Rect {
	if n == nil {
		return nil
	}
	var d Rect
	n.intoGo(&d)
	return &d
}

func (n *C.cef_rect_t) intoGo(d *Rect) {
	d.X = int32(n.x)
	d.Y = int32(n.y)
	d.Width = int32(n.width)
	d.Height = int32(n.height)
}
