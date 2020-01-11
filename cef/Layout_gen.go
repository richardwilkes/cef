// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// cef_box_layout_t * gocef_layout_as_box_layout(cef_layout_t * self, cef_box_layout_t * (CEF_CALLBACK *callback__)(cef_layout_t *)) { return callback__(self); }
	// cef_fill_layout_t * gocef_layout_as_fill_layout(cef_layout_t * self, cef_fill_layout_t * (CEF_CALLBACK *callback__)(cef_layout_t *)) { return callback__(self); }
	// int gocef_layout_is_valid(cef_layout_t * self, int (CEF_CALLBACK *callback__)(cef_layout_t *)) { return callback__(self); }
	"C"
)

// Layout (cef_layout_t from include/capi/views/cef_layout_capi.h)
// A Layout handles the sizing of the children of a Panel according to
// implementation-specific heuristics. Methods must be called on the browser
// process UI thread unless otherwise indicated.
type Layout C.cef_layout_t

func (d *Layout) toNative() *C.cef_layout_t {
	return (*C.cef_layout_t)(d)
}

// Base (base)
// Base structure.
func (d *Layout) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// AsBoxLayout (as_box_layout)
// Returns this Layout as a BoxLayout or NULL if this is not a BoxLayout.
func (d *Layout) AsBoxLayout() *BoxLayout {
	return (*BoxLayout)(C.gocef_layout_as_box_layout(d.toNative(), d.as_box_layout))
}

// AsFillLayout (as_fill_layout)
// Returns this Layout as a FillLayout or NULL if this is not a FillLayout.
func (d *Layout) AsFillLayout() *FillLayout {
	return (C.gocef_layout_as_fill_layout(d.toNative(), d.as_fill_layout)).toGo()
}

// IsValid (is_valid)
// Returns true (1) if this Layout is valid.
func (d *Layout) IsValid() int32 {
	return int32(C.gocef_layout_is_valid(d.toNative(), d.is_valid))
}
