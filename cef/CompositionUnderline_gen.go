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

// CompositionUnderline (cef_composition_underline_t from include/internal/cef_types.h)
// Structure representing IME composition underline information. This is a thin
// wrapper around Blink's WebCompositionUnderline class and should be kept in
// sync with that.
type CompositionUnderline struct {
	// Range (_range)
	// Underline character range.
	Range Range
	// Color (color)
	// Text color.
	Color Color
	// BackgroundColor (background_color)
	// Background color.
	BackgroundColor Color
	// Thick (thick)
	// Set to true (1) for thick underline.
	Thick int32
}

// NewCompositionUnderline creates a new CompositionUnderline.
func NewCompositionUnderline() *CompositionUnderline {
	return &CompositionUnderline{}
}

func (d *CompositionUnderline) toNative(native *C.cef_composition_underline_t) *C.cef_composition_underline_t {
	if d == nil {
		return nil
	}
	d.Range.toNative(&native._range)
	native.color = C.cef_color_t(d.Color)
	native.background_color = C.cef_color_t(d.BackgroundColor)
	native.thick = C.int(d.Thick)
	return native
}

func (n *C.cef_composition_underline_t) toGo() *CompositionUnderline {
	if n == nil {
		return nil
	}
	var d CompositionUnderline
	n.intoGo(&d)
	return &d
}

func (n *C.cef_composition_underline_t) intoGo(d *CompositionUnderline) {
	n._range.intoGo(&d.Range)
	d.Color = Color(n.color)
	d.BackgroundColor = Color(n.background_color)
	d.Thick = int32(n.thick)
}
