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
