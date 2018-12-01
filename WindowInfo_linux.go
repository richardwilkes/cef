package cef

import (
	// #include "include/internal/cef_types.h"
	"C"
	"unsafe"
)

// WindowInfo (cef_window_info_t) - Linux-specific variant
// Class representing window information.
type WindowInfo struct {
	// X (x)
	X int32
	// Y (y)
	Y int32
	// Width (width)
	Width int32
	// Height (height)
	Height int32
	// ParentWindow (parent_window)
	ParentWindow unsafe.Pointer
	// Menu (menu)
	Menu unsafe.Pointer
	// WindowlessRenderingEnabled (windowless_rendering_enabled)
	// Set to true (1) to create the browser using windowless (off-screen)
	// rendering. No window will be created for the browser and all rendering will
	// occur via the CefRenderHandler interface. The |parent_window| value will be
	// used to identify monitor info and to act as the parent window for dialogs,
	// context menus, etc. If |parent_window| is not provided then the main screen
	// monitor will be used and some functionality that requires a parent window
	// may not function correctly. In order to create windowless browsers the
	// CefSettings.windowless_rendering_enabled value must be set to true.
	// Transparent painting is enabled by default but can be disabled by setting
	// CefBrowserSettings.background_color to an opaque value.
	WindowlessRenderingEnabled int32
	// SharedTextEnabled (shared_texture_enabled)
	// Set to true (1) to enable shared textures for windowless rendering. Only
	// valid if windowless_rendering_enabled above is also set to true. Currently
	// only supported on Windows (D3D11).
	SharedTextureEnabled int32
	// ExternalBeginFrameEnabled (external_begin_frame_enabled)
	// Set to true (1) to enable the ability to issue BeginFrame requests from the
	// client application by calling CefBrowserHost::SendExternalBeginFrame.
	ExternalBeginFrameEnabled int32
	// Window (window)
	// Handle for the new browser window. Only used with windowed rendering.
	Window unsafe.Pointer
}

func (d *WindowInfo) toNative(native *C.cef_window_info_t) *C.cef_window_info_t {
	native.x = C.int(d.X)
	native.y = C.int(d.Y)
	native.width = C.int(d.Width)
	native.height = C.int(d.Height)
	native.parent_window = d.ParentWindow
	native.menu = d.Menu
	native.windowless_rendering_enabled = C.int(d.WindowlessRenderingEnabled)
	native.shared_texture_enabled = C.int(d.SharedTextureEnabled)
	native.external_begin_frame_enabled = C.int(d.ExternalBeginFrameEnabled)
	native.window = d.Window
	return native
}

func (n *C.cef_window_info_t) toGo() *WindowInfo {
	var d WindowInfo
	n.intoGo(&d)
	return &d
}

func (n *C.cef_window_info_t) intoGo(d *WindowInfo) {
	d.X = int32(n.x)
	d.Y = int32(n.y)
	d.Width = int32(n.width)
	d.Height = int32(n.height)
	d.ParentWindow = unsafe.Pointer(n.parent_window)
	d.Menu = unsafe.Pointer(n.menu)
	d.WindowlessRenderingEnabled = int32(n.windowless_rendering_enabled)
	d.SharedTextureEnabled = int32(n.shared_texture_enabled)
	d.ExternalBeginFrameEnabled = int32(n.external_begin_frame_enabled)
	d.Window = unsafe.Pointer(n.window)
}
