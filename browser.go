package cef

import (
	// #include "browser.h"
	"C"
)

// Browser represents a browser window. When used in the browser process the
// functions of this structure may be called on any thread unless otherwise
// indicated in the comments. When used in the render process the
// functions of this structure may only be called on the main thread.
//
// Defined in include/capi/cef_browser_capi.h
type Browser struct {
	native *C.cef_browser_t
}

// NewBrowser creates a new Browser instance. This function can only be called
// on the browser process UI thread.
func NewBrowser(info *WindowInfo, client *Client, url string, settings *BrowserSettings) *Browser {
	return &Browser{native: C.cef_browser_host_create_browser_sync(info.native, client.native, newCEFStr(url), settings.native, nil)}
}

// Host retrieves the BrowserHost.
func (b *Browser) Host() *BrowserHost {
	return &BrowserHost{native: C.gocef_get_browser_host(b.native)}
}

// FocusedFrame returns the currently focused frame.
func (b *Browser) FocusedFrame() *Frame {
	if f := C.gocef_get_focused_frame(b.native); f != nil {
		return &Frame{native: f}
	}
	return nil
}
