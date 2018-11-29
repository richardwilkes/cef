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
	native.x = C.int(d.X)
	native.y = C.int(d.Y)
	return native
}

func (d *Point) fromNative(native *C.cef_point_t) *Point {
	d.X = int32(native.x)
	d.Y = int32(native.y)
	return d
}
