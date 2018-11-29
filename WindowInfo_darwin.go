package cef

import (
	// #include "include/internal/cef_types.h"
	"C"
	"unsafe"
)

// WindowInfo (cef_window_info_t) - Mac-specific variant
// Class representing window information.
type WindowInfo struct {
	// WindowName (window_name)
	WindowName string
	// X (x)
	X int32
	// Y (y)
	Y int32
	// Width (width)
	Width int32
	// Height (height)
	Height int32
	// Hidden (hidden)
	// Set to true (1) to create the view initially hidden.
	Hidden int32
	// ParentView (parent_view)
	ParentView unsafe.Pointer
	// WindowlessRenderingEnabled (windowless_rendering_enabled)
	// Set to true (1) to create the browser using windowless (off-screen)
	// rendering. No view will be created for the browser and all rendering will
	// occur via the CefRenderHandler interface. The |parent_view| value will be
	// used to identify monitor info and to act as the parent view for dialogs,
	// context menus, etc. If |parent_view| is not provided then the main screen
	// monitor will be used and some functionality that requires a parent view
	// may not function correctly. In order to create windowless browsers the
	// CefSettings.windowless_rendering_enabled value must be set to true.
	// Transparent painting is enabled by default but can be disabled by setting
	// CefBrowserSettings.background_color to an opaque value.
	WindowlessRenderingEnabled int32
	// View (view)
	View unsafe.Pointer
}

func (d *WindowInfo) platformInit(parent unsafe.Pointer) {
	d.ParentView = parent
}

func (d *WindowInfo) toNative(native *C.cef_window_info_t) *C.cef_window_info_t {
	setCEFStr(d.WindowName, &native.window_name)
	native.x = C.int(d.X)
	native.y = C.int(d.Y)
	native.width = C.int(d.Width)
	native.height = C.int(d.Height)
	native.hidden = C.int(d.Hidden)
	native.parent_view = d.ParentView
	native.windowless_rendering_enabled = C.int(d.WindowlessRenderingEnabled)
	native.view = d.View
	return native
}

func (d *WindowInfo) fromNative(native *C.cef_window_info_t) *WindowInfo {
	d.WindowName = cefstrToString(&native.window_name)
	d.X = int32(native.x)
	d.Y = int32(native.y)
	d.Width = int32(native.width)
	d.Height = int32(native.height)
	d.Hidden = int32(native.hidden)
	d.ParentView = unsafe.Pointer(native.parent_view)
	d.WindowlessRenderingEnabled = int32(native.windowless_rendering_enabled)
	d.View = unsafe.Pointer(native.view)
	return d
}
