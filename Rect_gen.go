// Code generated - DO NOT EDIT.

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
	native.x = C.int(d.X)
	native.y = C.int(d.Y)
	native.width = C.int(d.Width)
	native.height = C.int(d.Height)
	return native
}

func (d *Rect) fromNative(native *C.cef_rect_t) *Rect {
	d.X = int32(native.x)
	d.Y = int32(native.y)
	d.Width = int32(native.width)
	d.Height = int32(native.height)
	return d
}
