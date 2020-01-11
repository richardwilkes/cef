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

// Size (cef_size_t from include/internal/cef_types.h)
// Structure representing a size.
type Size struct {
	// Width (width)
	Width int32
	// Height (height)
	Height int32
}

// NewSize creates a new Size.
func NewSize() *Size {
	return &Size{}
}

func (d *Size) toNative(native *C.cef_size_t) *C.cef_size_t {
	if d == nil {
		return nil
	}
	native.width = C.int(d.Width)
	native.height = C.int(d.Height)
	return native
}

func (n *C.cef_size_t) toGo() *Size {
	if n == nil {
		return nil
	}
	var d Size
	n.intoGo(&d)
	return &d
}

func (n *C.cef_size_t) intoGo(d *Size) {
	d.Width = int32(n.width)
	d.Height = int32(n.height)
}
