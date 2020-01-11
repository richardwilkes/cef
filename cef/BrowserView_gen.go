// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// cef_browser_t * gocef_browser_view_get_browser(cef_browser_view_t * self, cef_browser_t * (CEF_CALLBACK *callback__)(cef_browser_view_t *)) { return callback__(self); }
	// void gocef_browser_view_set_prefer_accelerators(cef_browser_view_t * self, int preferAccelerators, void (CEF_CALLBACK *callback__)(cef_browser_view_t *, int)) { return callback__(self, preferAccelerators); }
	"C"
)

// BrowserView (cef_browser_view_t from include/capi/views/cef_browser_view_capi.h)
// A View hosting a cef_browser_t instance. Methods must be called on the
// browser process UI thread unless otherwise indicated.
type BrowserView C.cef_browser_view_t

func (d *BrowserView) toNative() *C.cef_browser_view_t {
	return (*C.cef_browser_view_t)(d)
}

// Base (base)
// Base structure.
func (d *BrowserView) Base() *View {
	return (*View)(&d.base)
}

// GetBrowser (get_browser)
// Returns the cef_browser_t hosted by this BrowserView. Will return NULL if
// the browser has not yet been created or has already been destroyed.
func (d *BrowserView) GetBrowser() *Browser {
	return (*Browser)(C.gocef_browser_view_get_browser(d.toNative(), d.get_browser))
}

// SetPreferAccelerators (set_prefer_accelerators)
// Sets whether accelerators registered with cef_window_t::SetAccelerator are
// triggered before or after the event is sent to the cef_browser_t. If
// |prefer_accelerators| is true (1) then the matching accelerator will be
// triggered immediately and the event will not be sent to the cef_browser_t.
// If |prefer_accelerators| is false (0) then the matching accelerator will
// only be triggered if the event is not handled by web content or by
// cef_keyboard_handler_t. The default value is false (0).
func (d *BrowserView) SetPreferAccelerators(preferAccelerators int32) {
	C.gocef_browser_view_set_prefer_accelerators(d.toNative(), C.int(preferAccelerators), d.set_prefer_accelerators)
}
