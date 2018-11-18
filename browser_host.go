package cef

import (
	// #include "browser_host.h"
	"C"
	"unsafe"
)

var x unsafe.Pointer // Just here to pull in unsafe, which the declaration for WindowHandle needs but doesn't pull in.

// WindowHandle holds a pointer to the parent native window or view.
type WindowHandle C.cef_window_handle_t

// BrowserHost is used to represent the browser process aspects of a browser
// window. The functions of this structure can only be called in the browser
// process. They may be called on any thread in that process unless otherwise
// indicated in the comments.
type BrowserHost struct {
	native *C.cef_browser_host_t
}

// WindowHandle retrieves the window handle for this browser. If this browser
// is wrapped in a BrowserView this function should be called on the browser
// process UI thread and it will return the handle for the top-level native
// window.
func (b *BrowserHost) WindowHandle() WindowHandle {
	return WindowHandle(C.gocef_windowhandle_browserhost(b.native, b.native.get_window_handle))
}
