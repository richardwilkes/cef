// Code created from "struct.go.tmpl" - don't edit by hand

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
	if d == nil {
		return nil
	}
	d.Bounds.toNative(&native.bounds)
	native.draggable = C.int(d.Draggable)
	return native
}

func (n *C.cef_draggable_region_t) toGo() *DraggableRegion {
	if n == nil {
		return nil
	}
	var d DraggableRegion
	n.intoGo(&d)
	return &d
}

func (n *C.cef_draggable_region_t) intoGo(d *DraggableRegion) {
	n.bounds.intoGo(&d.Bounds)
	d.Draggable = int32(n.draggable)
}
