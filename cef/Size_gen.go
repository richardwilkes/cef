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
