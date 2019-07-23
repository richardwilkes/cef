// Code created from "struct.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	"C"
)

// Point (cef_point_t from include/internal/cef_types.h)
// Structure representing a point.
type Point struct {
	// X (x)
	X int32
	// Y (y)
	Y int32
}

// NewPoint creates a new Point.
func NewPoint() *Point {
	return &Point{}
}

func (d *Point) toNative(native *C.cef_point_t) *C.cef_point_t {
	if d == nil {
		return nil
	}
	native.x = C.int(d.X)
	native.y = C.int(d.Y)
	return native
}

func (n *C.cef_point_t) toGo() *Point {
	if n == nil {
		return nil
	}
	var d Point
	n.intoGo(&d)
	return &d
}

func (n *C.cef_point_t) intoGo(d *Point) {
	d.X = int32(n.x)
	d.Y = int32(n.y)
}
