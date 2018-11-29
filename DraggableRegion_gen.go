package cef

import (
	// #include "capi_gen.h"
	"C"
)

// DraggableRegion (cef_draggable_region_t from include/internal/cef_types.h)
// Structure representing a draggable region.
type DraggableRegion struct {
	// Bounds (bounds)
	// Bounds of the region.
	Bounds Rect
	// Draggable (draggable)
	// True (1) this this region is draggable and false (0) otherwise.
	Draggable int32
}

// NewDraggableRegion creates a new DraggableRegion.
func NewDraggableRegion() *DraggableRegion {
	return &DraggableRegion{}
}

func (d *DraggableRegion) toNative(native *C.cef_draggable_region_t) *C.cef_draggable_region_t {
	d.Bounds.toNative(&native.bounds)
	native.draggable = C.int(d.Draggable)
	return native
}

func (d *DraggableRegion) fromNative(native *C.cef_draggable_region_t) *DraggableRegion {
	d.Bounds.fromNative(&native.bounds)
	d.Draggable = int32(native.draggable)
	return d
}
