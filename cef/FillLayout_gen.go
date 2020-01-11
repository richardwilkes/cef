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
	"unsafe"
)

import (
	// #include "capi_gen.h"
	"C"
)

// FillLayout (cef_fill_layout_t from include/capi/views/cef_fill_layout_capi.h)
// A simple Layout that causes the associated Panel's one child to be sized to
// match the bounds of its parent. Methods must be called on the browser process
// UI thread unless otherwise indicated.
type FillLayout struct {
	// Base (base)
	// Base structure.
	Base *Layout
}

// NewFillLayout creates a new FillLayout.
func NewFillLayout() *FillLayout {
	return &FillLayout{}
}

func (d *FillLayout) toNative(native *C.cef_fill_layout_t) *C.cef_fill_layout_t {
	if d == nil {
		return nil
	}
	native.base = *(*C.cef_layout_t)(unsafe.Pointer(d.Base))
	return native
}

func (n *C.cef_fill_layout_t) toGo() *FillLayout {
	if n == nil {
		return nil
	}
	var d FillLayout
	n.intoGo(&d)
	return &d
}

func (n *C.cef_fill_layout_t) intoGo(d *FillLayout) {
	d.Base = (*Layout)(&n.base)
}
