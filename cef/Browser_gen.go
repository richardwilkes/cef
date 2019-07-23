// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// cef_browser_host_t * gocef_browser_get_host(cef_browser_t * self, cef_browser_host_t * (CEF_CALLBACK *callback__)(cef_browser_t *)) { return callback__(self); }
	// int gocef_browser_can_go_back(cef_browser_t * self, int (CEF_CALLBACK *callback__)(cef_browser_t *)) { return callback__(self); }
	// void gocef_browser_go_back(cef_browser_t * self, void (CEF_CALLBACK *callback__)(cef_browser_t *)) { return callback__(self); }
	// int gocef_browser_can_go_forward(cef_browser_t * self, int (CEF_CALLBACK *callback__)(cef_browser_t *)) { return callback__(self); }
	// void gocef_browser_go_forward(cef_browser_t * self, void (CEF_CALLBACK *callback__)(cef_browser_t *)) { return callback__(self); }
	// int gocef_browser_is_loading(cef_browser_t * self, int (CEF_CALLBACK *callback__)(cef_browser_t *)) { return callback__(self); }
	// void gocef_browser_reload(cef_browser_t * self, void (CEF_CALLBACK *callback__)(cef_browser_t *)) { return callback__(self); }
	// void gocef_browser_reload_ignore_cache(cef_browser_t * self, void (CEF_CALLBACK *callback__)(cef_browser_t *)) { return callback__(self); }
	// void gocef_browser_stop_load(cef_browser_t * self, void (CEF_CALLBACK *callback__)(cef_browser_t *)) { return callback__(self); }
	// int gocef_browser_get_identifier(cef_browser_t * self, int (CEF_CALLBACK *callback__)(cef_browser_t *)) { return callback__(self); }
	// int gocef_browser_is_same(cef_browser_t * self, cef_browser_t * that, int (CEF_CALLBACK *callback__)(cef_browser_t *, cef_browser_t *)) { return callback__(self, that); }
	// int gocef_browser_is_popup(cef_browser_t * self, int (CEF_CALLBACK *callback__)(cef_browser_t *)) { return callback__(self); }
	// int gocef_browser_has_document(cef_browser_t * self, int (CEF_CALLBACK *callback__)(cef_browser_t *)) { return callback__(self); }
	// cef_frame_t * gocef_browser_get_main_frame(cef_browser_t * self, cef_frame_t * (CEF_CALLBACK *callback__)(cef_browser_t *)) { return callback__(self); }
	// cef_frame_t * gocef_browser_get_focused_frame(cef_browser_t * self, cef_frame_t * (CEF_CALLBACK *callback__)(cef_browser_t *)) { return callback__(self); }
	// cef_frame_t * gocef_browser_get_frame_byident(cef_browser_t * self, int64 identifier, cef_frame_t * (CEF_CALLBACK *callback__)(cef_browser_t *, int64)) { return callback__(self, identifier); }
	// cef_frame_t * gocef_browser_get_frame(cef_browser_t * self, cef_string_t * name, cef_frame_t * (CEF_CALLBACK *callback__)(cef_browser_t *, cef_string_t *)) { return callback__(self, name); }
	// size_t gocef_browser_get_frame_count(cef_browser_t * self, size_t (CEF_CALLBACK *callback__)(cef_browser_t *)) { return callback__(self); }
	// void gocef_browser_get_frame_identifiers(cef_browser_t * self, size_t * identifiersCount, int64 * identifiers, void (CEF_CALLBACK *callback__)(cef_browser_t *, size_t *, int64 *)) { return callback__(self, identifiersCount, identifiers); }
	// void gocef_browser_get_frame_names(cef_browser_t * self, cef_string_list_t names, void (CEF_CALLBACK *callback__)(cef_browser_t *, cef_string_list_t)) { return callback__(self, names); }
	"C"
)

// Browser (cef_browser_t from include/capi/cef_browser_capi.h)
// Structure used to represent a browser window. When used in the browser
// process the functions of this structure may be called on any thread unless
// otherwise indicated in the comments. When used in the render process the
// functions of this structure may only be called on the main thread.
type Browser C.cef_browser_t

func (d *Browser) toNative() *C.cef_browser_t {
	return (*C.cef_browser_t)(d)
}

// Base (base)
// Base structure.
func (d *Browser) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// GetHost (get_host)
// Returns the browser host object. This function can only be called in the
// browser process.
func (d *Browser) GetHost() *BrowserHost {
	return (*BrowserHost)(C.gocef_browser_get_host(d.toNative(), d.get_host))
}

// CanGoBack (can_go_back)
// Returns true (1) if the browser can navigate backwards.
func (d *Browser) CanGoBack() int32 {
	return int32(C.gocef_browser_can_go_back(d.toNative(), d.can_go_back))
}

// GoBack (go_back)
// Navigate backwards.
func (d *Browser) GoBack() {
	C.gocef_browser_go_back(d.toNative(), d.go_back)
}

// CanGoForward (can_go_forward)
// Returns true (1) if the browser can navigate forwards.
func (d *Browser) CanGoForward() int32 {
	return int32(C.gocef_browser_can_go_forward(d.toNative(), d.can_go_forward))
}

// GoForward (go_forward)
// Navigate forwards.
func (d *Browser) GoForward() {
	C.gocef_browser_go_forward(d.toNative(), d.go_forward)
}

// IsLoading (is_loading)
// Returns true (1) if the browser is currently loading.
func (d *Browser) IsLoading() int32 {
	return int32(C.gocef_browser_is_loading(d.toNative(), d.is_loading))
}

// Reload (reload)
// Reload the current page.
func (d *Browser) Reload() {
	C.gocef_browser_reload(d.toNative(), d.reload)
}

// ReloadIgnoreCache (reload_ignore_cache)
// Reload the current page ignoring any cached data.
func (d *Browser) ReloadIgnoreCache() {
	C.gocef_browser_reload_ignore_cache(d.toNative(), d.reload_ignore_cache)
}

// StopLoad (stop_load)
// Stop loading the page.
func (d *Browser) StopLoad() {
	C.gocef_browser_stop_load(d.toNative(), d.stop_load)
}

// GetIdentifier (get_identifier)
// Returns the globally unique identifier for this browser. This value is also
// used as the tabId for extension APIs.
func (d *Browser) GetIdentifier() int32 {
	return int32(C.gocef_browser_get_identifier(d.toNative(), d.get_identifier))
}

// IsSame (is_same)
// Returns true (1) if this object is pointing to the same handle as |that|
// object.
func (d *Browser) IsSame(that *Browser) int32 {
	return int32(C.gocef_browser_is_same(d.toNative(), that.toNative(), d.is_same))
}

// IsPopup (is_popup)
// Returns true (1) if the window is a popup window.
func (d *Browser) IsPopup() int32 {
	return int32(C.gocef_browser_is_popup(d.toNative(), d.is_popup))
}

// HasDocument (has_document)
// Returns true (1) if a document has been loaded in the browser.
func (d *Browser) HasDocument() int32 {
	return int32(C.gocef_browser_has_document(d.toNative(), d.has_document))
}

// GetMainFrame (get_main_frame)
// Returns the main (top-level) frame for the browser window.
func (d *Browser) GetMainFrame() *Frame {
	return (*Frame)(C.gocef_browser_get_main_frame(d.toNative(), d.get_main_frame))
}

// GetFocusedFrame (get_focused_frame)
// Returns the focused frame for the browser window.
func (d *Browser) GetFocusedFrame() *Frame {
	return (*Frame)(C.gocef_browser_get_focused_frame(d.toNative(), d.get_focused_frame))
}

// GetFrameByident (get_frame_byident)
// Returns the frame with the specified identifier, or NULL if not found.
func (d *Browser) GetFrameByident(identifier int64) *Frame {
	return (*Frame)(C.gocef_browser_get_frame_byident(d.toNative(), C.int64(identifier), d.get_frame_byident))
}

// GetFrame (get_frame)
// Returns the frame with the specified name, or NULL if not found.
func (d *Browser) GetFrame(name string) *Frame {
	name_ := C.cef_string_userfree_alloc()
	setCEFStr(name, name_)
	defer func() {
		C.cef_string_userfree_free(name_)
	}()
	return (*Frame)(C.gocef_browser_get_frame(d.toNative(), (*C.cef_string_t)(name_), d.get_frame))
}

// GetFrameCount (get_frame_count)
// Returns the number of frames that currently exist.
func (d *Browser) GetFrameCount() uint64 {
	return uint64(C.gocef_browser_get_frame_count(d.toNative(), d.get_frame_count))
}

// GetFrameIdentifiers (get_frame_identifiers)
// Returns the identifiers of all existing frames.
func (d *Browser) GetFrameIdentifiers(identifiersCount *uint64, identifiers *int64) {
	C.gocef_browser_get_frame_identifiers(d.toNative(), (*C.size_t)(identifiersCount), (*C.int64)(identifiers), d.get_frame_identifiers)
}

// GetFrameNames (get_frame_names)
// Returns the names of all existing frames.
func (d *Browser) GetFrameNames(names StringList) {
	C.gocef_browser_get_frame_names(d.toNative(), C.cef_string_list_t(names), d.get_frame_names)
}
