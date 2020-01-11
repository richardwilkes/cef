// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// int64 gocef_display_get_id(cef_display_t * self, int64 (CEF_CALLBACK *callback__)(cef_display_t *)) { return callback__(self); }
	// float gocef_display_get_device_scale_factor(cef_display_t * self, float (CEF_CALLBACK *callback__)(cef_display_t *)) { return callback__(self); }
	// void gocef_display_convert_point_to_pixels(cef_display_t * self, cef_point_t * point, void (CEF_CALLBACK *callback__)(cef_display_t *, cef_point_t *)) { return callback__(self, point); }
	// void gocef_display_convert_point_from_pixels(cef_display_t * self, cef_point_t * point, void (CEF_CALLBACK *callback__)(cef_display_t *, cef_point_t *)) { return callback__(self, point); }
	// cef_rect_t gocef_display_get_bounds(cef_display_t * self, cef_rect_t (CEF_CALLBACK *callback__)(cef_display_t *)) { return callback__(self); }
	// cef_rect_t gocef_display_get_work_area(cef_display_t * self, cef_rect_t (CEF_CALLBACK *callback__)(cef_display_t *)) { return callback__(self); }
	// int gocef_display_get_rotation(cef_display_t * self, int (CEF_CALLBACK *callback__)(cef_display_t *)) { return callback__(self); }
	"C"
)

// Display (cef_display_t from include/capi/views/cef_display_capi.h)
// This structure typically, but not always, corresponds to a physical display
// connected to the system. A fake Display may exist on a headless system, or a
// Display may correspond to a remote, virtual display. All size and position
// values are in density independent pixels (DIP) unless otherwise indicated.
// Methods must be called on the browser process UI thread unless otherwise
// indicated.
type Display C.cef_display_t

func (d *Display) toNative() *C.cef_display_t {
	return (*C.cef_display_t)(d)
}

// Base (base)
// Base structure.
func (d *Display) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// GetID (get_id)
// Returns the unique identifier for this Display.
func (d *Display) GetID() int64 {
	return int64(C.gocef_display_get_id(d.toNative(), d.get_id))
}

// GetDeviceScaleFactor (get_device_scale_factor)
// Returns this Display's device pixel scale factor. This specifies how much
// the UI should be scaled when the actual output has more pixels than
// standard displays (which is around 100~120dpi). The potential return values
// differ by platform.
func (d *Display) GetDeviceScaleFactor() float32 {
	return float32(C.gocef_display_get_device_scale_factor(d.toNative(), d.get_device_scale_factor))
}

// ConvertPointToPixels (convert_point_to_pixels)
// Convert |point| from density independent pixels (DIP) to pixel coordinates
// using this Display's device scale factor.
func (d *Display) ConvertPointToPixels(point *Point) {
	C.gocef_display_convert_point_to_pixels(d.toNative(), point.toNative(&C.cef_point_t{}), d.convert_point_to_pixels)
}

// ConvertPointFromPixels (convert_point_from_pixels)
// Convert |point| from pixel coordinates to density independent pixels (DIP)
// using this Display's device scale factor.
func (d *Display) ConvertPointFromPixels(point *Point) {
	C.gocef_display_convert_point_from_pixels(d.toNative(), point.toNative(&C.cef_point_t{}), d.convert_point_from_pixels)
}

// GetBounds (get_bounds)
// Returns this Display's bounds. This is the full size of the display.
func (d *Display) GetBounds() Rect {
	cresult__ := C.gocef_display_get_bounds(d.toNative(), d.get_bounds)
	var result__ Rect
	(&cresult__).intoGo(&result__)
	return result__
}

// GetWorkArea (get_work_area)
// Returns this Display's work area. This excludes areas of the display that
// are occupied for window manager toolbars, etc.
func (d *Display) GetWorkArea() Rect {
	cresult__ := C.gocef_display_get_work_area(d.toNative(), d.get_work_area)
	var result__ Rect
	(&cresult__).intoGo(&result__)
	return result__
}

// GetRotation (get_rotation)
// Returns this Display's rotation in degrees.
func (d *Display) GetRotation() int32 {
	return int32(C.gocef_display_get_rotation(d.toNative(), d.get_rotation))
}
