// Code generated - DO NOT EDIT.

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
	native.width = C.int(d.Width)
	native.height = C.int(d.Height)
	return native
}

func (d *Size) fromNative(native *C.cef_size_t) *Size {
	d.Width = int32(native.width)
	d.Height = int32(native.height)
	return d
}
