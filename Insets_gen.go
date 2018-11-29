package cef

import (
	// #include "capi_gen.h"
	"C"
)

// Insets (cef_insets_t from include/internal/cef_types.h)
// Structure representing insets.
type Insets struct {
	// Top (top)
	Top int32
	// Left (left)
	Left int32
	// Bottom (bottom)
	Bottom int32
	// Right (right)
	Right int32
}

// NewInsets creates a new Insets.
func NewInsets() *Insets {
	return &Insets{}
}

func (d *Insets) toNative(native *C.cef_insets_t) *C.cef_insets_t {
	native.top = C.int(d.Top)
	native.left = C.int(d.Left)
	native.bottom = C.int(d.Bottom)
	native.right = C.int(d.Right)
	return native
}

func (d *Insets) fromNative(native *C.cef_insets_t) *Insets {
	d.Top = int32(native.top)
	d.Left = int32(native.left)
	d.Bottom = int32(native.bottom)
	d.Right = int32(native.right)
	return d
}
