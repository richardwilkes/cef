// Code created from "struct.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"
)

import (
	// #include "capi_gen.h"
	"C"
)

// CursorInfo (cef_cursor_info_t from include/internal/cef_types.h)
// Structure representing cursor information. |buffer| will be
// |size.width|*|size.height|*4 bytes in size and represents a BGRA image with
// an upper-left origin.
type CursorInfo struct {
	// Hotspot (hotspot)
	Hotspot Point
	// ImageScaleFactor (image_scale_factor)
	ImageScaleFactor float32
	// Buffer (buffer)
	Buffer unsafe.Pointer
	// Size (size)
	Size Size
}

// NewCursorInfo creates a new CursorInfo.
func NewCursorInfo() *CursorInfo {
	return &CursorInfo{}
}

func (d *CursorInfo) toNative(native *C.cef_cursor_info_t) *C.cef_cursor_info_t {
	if d == nil {
		return nil
	}
	d.Hotspot.toNative(&native.hotspot)
	native.image_scale_factor = C.float(d.ImageScaleFactor)
	native.buffer = d.Buffer
	d.Size.toNative(&native.size)
	return native
}

func (n *C.cef_cursor_info_t) toGo() *CursorInfo {
	if n == nil {
		return nil
	}
	var d CursorInfo
	n.intoGo(&d)
	return &d
}

func (n *C.cef_cursor_info_t) intoGo(d *CursorInfo) {
	n.hotspot.intoGo(&d.Hotspot)
	d.ImageScaleFactor = float32(n.image_scale_factor)
	d.Buffer = n.buffer
	n.size.intoGo(&d.Size)
}
