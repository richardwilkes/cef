// Code generated - DO NOT EDIT.

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
	d.Range.toNative(&native._range)
	native.color = C.cef_color_t(d.Color)
	native.background_color = C.cef_color_t(d.BackgroundColor)
	native.thick = C.int(d.Thick)
	return native
}

func (d *CompositionUnderline) fromNative(native *C.cef_composition_underline_t) *CompositionUnderline {
	d.Range.fromNative(&native._range)
	d.Color = Color(native.color)
	d.BackgroundColor = Color(native.background_color)
	d.Thick = int32(native.thick)
	return d
}
