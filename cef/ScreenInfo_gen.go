// Code created from "struct.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	"C"
)

// ScreenInfo (cef_screen_info_t from include/internal/cef_types.h)
// Screen information used when window rendering is disabled. This structure is
// passed as a parameter to CefRenderHandler::GetScreenInfo and should be filled
// in by the client.
type ScreenInfo struct {
	// DeviceScaleFactor (device_scale_factor)
	// Device scale factor. Specifies the ratio between physical and logical
	// pixels.
	DeviceScaleFactor float32
	// Depth (depth)
	// The screen depth in bits per pixel.
	Depth int32
	// DepthPerComponent (depth_per_component)
	// The bits per color component. This assumes that the colors are balanced
	// equally.
	DepthPerComponent int32
	// IsMonochrome (is_monochrome)
	// This can be true for black and white printers.
	IsMonochrome int32
	// Rect (rect)
	// This is set from the rcMonitor member of MONITORINFOEX, to whit:
	//   "A RECT structure that specifies the display monitor rectangle,
	//   expressed in virtual-screen coordinates. Note that if the monitor
	//   is not the primary display monitor, some of the rectangle's
	//   coordinates may be negative values."
	//
	// The |rect| and |available_rect| properties are used to determine the
	// available surface for rendering popup views.
	Rect Rect
	// AvailableRect (available_rect)
	// This is set from the rcWork member of MONITORINFOEX, to whit:
	//   "A RECT structure that specifies the work area rectangle of the
	//   display monitor that can be used by applications, expressed in
	//   virtual-screen coordinates. Windows uses this rectangle to
	//   maximize an application on the monitor. The rest of the area in
	//   rcMonitor contains system windows such as the task bar and side
	//   bars. Note that if the monitor is not the primary display monitor,
	//   some of the rectangle's coordinates may be negative values".
	//
	// The |rect| and |available_rect| properties are used to determine the
	// available surface for rendering popup views.
	AvailableRect Rect
}

// NewScreenInfo creates a new ScreenInfo.
func NewScreenInfo() *ScreenInfo {
	return &ScreenInfo{}
}

func (d *ScreenInfo) toNative(native *C.cef_screen_info_t) *C.cef_screen_info_t {
	if d == nil {
		return nil
	}
	native.device_scale_factor = C.float(d.DeviceScaleFactor)
	native.depth = C.int(d.Depth)
	native.depth_per_component = C.int(d.DepthPerComponent)
	native.is_monochrome = C.int(d.IsMonochrome)
	d.Rect.toNative(&native.rect)
	d.AvailableRect.toNative(&native.available_rect)
	return native
}

func (n *C.cef_screen_info_t) toGo() *ScreenInfo {
	if n == nil {
		return nil
	}
	var d ScreenInfo
	n.intoGo(&d)
	return &d
}

func (n *C.cef_screen_info_t) intoGo(d *ScreenInfo) {
	d.DeviceScaleFactor = float32(n.device_scale_factor)
	d.Depth = int32(n.depth)
	d.DepthPerComponent = int32(n.depth_per_component)
	d.IsMonochrome = int32(n.is_monochrome)
	n.rect.intoGo(&d.Rect)
	n.available_rect.intoGo(&d.AvailableRect)
}
