package cef

import (
	// #include "capi_gen.h"
	"C"
)

// Range (cef_range_t from include/internal/cef_types.h)
// Structure representing a range.
type Range struct {
	// From (from)
	From int32
	// To (to)
	To int32
}

// NewRange creates a new Range.
func NewRange() *Range {
	return &Range{}
}

func (d *Range) toNative(native *C.cef_range_t) *C.cef_range_t {
	native.from = C.int(d.From)
	native.to = C.int(d.To)
	return native
}

func (d *Range) fromNative(native *C.cef_range_t) *Range {
	d.From = int32(native.from)
	d.To = int32(native.to)
	return d
}
