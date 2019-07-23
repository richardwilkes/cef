package cef

import "unsafe"

import (
	// #include "include/internal/cef_types.h"
	"C"
)

// WindowInfo (cef_window_info_t) - Mac-specific variant
// Class representing window information.
type WindowInfo struct { //nolint:maligned
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

func (n *C.cef_window_info_t) toGo() *WindowInfo {
	var d WindowInfo
	n.intoGo(&d)
	return &d
}

func (n *C.cef_window_info_t) intoGo(d *WindowInfo) {
	d.WindowName = cefstrToString(&n.window_name)
	d.X = int32(n.x)
	d.Y = int32(n.y)
	d.Width = int32(n.width)
	d.Height = int32(n.height)
	d.Hidden = int32(n.hidden)
	d.ParentView = n.parent_view
	d.WindowlessRenderingEnabled = int32(n.windowless_rendering_enabled)
	d.View = n.view
}
