package cef

import (
	// #include "capi_gen.h"
	"C"
	"unsafe"
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
	d.Hotspot.toNative(&native.hotspot)
	native.image_scale_factor = C.float(d.ImageScaleFactor)
	native.buffer = d.Buffer
	d.Size.toNative(&native.size)
	return native
}

func (d *CursorInfo) fromNative(native *C.cef_cursor_info_t) *CursorInfo {
	d.Hotspot.fromNative(&native.hotspot)
	d.ImageScaleFactor = float32(native.image_scale_factor)
	d.Buffer = unsafe.Pointer(native.buffer)
	d.Size.fromNative(&native.size)
	return d
}
