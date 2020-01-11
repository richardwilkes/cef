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

// Insets (cef_insets_t from include/internal/cef_types.h)
// Structure representing insets.
type Insets struct {
	// Top (top)
	Top int32
	// Left (left)
	Left int32
	// Bottom (bottom)
	Bottom int32
	// Right (right)
	Right int32
}

// NewInsets creates a new Insets.
func NewInsets() *Insets {
	return &Insets{}
}

func (d *Insets) toNative(native *C.cef_insets_t) *C.cef_insets_t {
	if d == nil {
		return nil
	}
	native.top = C.int(d.Top)
	native.left = C.int(d.Left)
	native.bottom = C.int(d.Bottom)
	native.right = C.int(d.Right)
	return native
}

func (n *C.cef_insets_t) toGo() *Insets {
	if n == nil {
		return nil
	}
	var d Insets
	n.intoGo(&d)
	return &d
}

func (n *C.cef_insets_t) intoGo(d *Insets) {
	d.Top = int32(n.top)
	d.Left = int32(n.left)
	d.Bottom = int32(n.bottom)
	d.Right = int32(n.right)
}
