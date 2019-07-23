package cef

import "unsafe"

import (
	// #include "include/internal/cef_types.h"
	"C"
)

// WindowInfo (cef_window_info_t) - Windows-specific variant
// Class representing window information.
type WindowInfo struct {
	// ExStyle (ex_style)
	// Standard parameters required by CreateWindowEx()
	ExStyle uint32
	// WindowName (window_name)
	WindowName string
	// Style (style)
	Style uint32
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
	// Window (window)
	// Handle for the new browser window. Only used with windowed rendering.
	Window unsafe.Pointer
}

func (d *WindowInfo) platformInit(parent unsafe.Pointer) {
	d.Style = C.WS_CHILD | C.WS_CLIPCHILDREN | C.WS_CLIPSIBLINGS | C.WS_TABSTOP | C.WS_VISIBLE
	d.ParentWindow = parent
}

func (d *WindowInfo) toNative(native *C.cef_window_info_t) *C.cef_window_info_t {
	native.ex_style = C.DWORD(d.ExStyle)
	setCEFStr(d.WindowName, &native.window_name)
	native.style = C.DWORD(d.Style)
	native.x = C.int(d.X)
	native.y = C.int(d.Y)
	native.width = C.int(d.Width)
	native.height = C.int(d.Height)
	native.parent_window = C.HWND(d.ParentWindow)
	native.menu = C.HMENU(d.Menu)
	native.windowless_rendering_enabled = C.int(d.WindowlessRenderingEnabled)
	native.window = C.HWND(d.Window)
	return native
}

func (n *C.cef_window_info_t) toGo() *WindowInfo {
	var d WindowInfo
	n.intoGo(&d)
	return &d
}

func (n *C.cef_window_info_t) intoGo(d *WindowInfo) {
	d.ExStyle = uint32(n.ex_style)
	d.WindowName = cefstrToString(&n.window_name)
	d.Style = uint32(n.style)
	d.X = int32(n.x)
	d.Y = int32(n.y)
	d.Width = int32(n.width)
	d.Height = int32(n.height)
	d.ParentWindow = unsafe.Pointer(n.parent_window)
	d.Menu = unsafe.Pointer(n.menu)
	d.WindowlessRenderingEnabled = int32(n.windowless_rendering_enabled)
	d.Window = unsafe.Pointer(n.window)
}
