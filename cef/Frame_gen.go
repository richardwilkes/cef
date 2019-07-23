// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// int gocef_frame_is_valid(cef_frame_t * self, int (CEF_CALLBACK *callback__)(cef_frame_t *)) { return callback__(self); }
	// void gocef_frame_undo(cef_frame_t * self, void (CEF_CALLBACK *callback__)(cef_frame_t *)) { return callback__(self); }
	// void gocef_frame_redo(cef_frame_t * self, void (CEF_CALLBACK *callback__)(cef_frame_t *)) { return callback__(self); }
	// void gocef_frame_cut(cef_frame_t * self, void (CEF_CALLBACK *callback__)(cef_frame_t *)) { return callback__(self); }
	// void gocef_frame_copy(cef_frame_t * self, void (CEF_CALLBACK *callback__)(cef_frame_t *)) { return callback__(self); }
	// void gocef_frame_paste(cef_frame_t * self, void (CEF_CALLBACK *callback__)(cef_frame_t *)) { return callback__(self); }
	// void gocef_frame_del(cef_frame_t * self, void (CEF_CALLBACK *callback__)(cef_frame_t *)) { return callback__(self); }
	// void gocef_frame_select_all(cef_frame_t * self, void (CEF_CALLBACK *callback__)(cef_frame_t *)) { return callback__(self); }
	// void gocef_frame_view_source(cef_frame_t * self, void (CEF_CALLBACK *callback__)(cef_frame_t *)) { return callback__(self); }
	// void gocef_frame_get_source(cef_frame_t * self, cef_string_visitor_t * visitor, void (CEF_CALLBACK *callback__)(cef_frame_t *, cef_string_visitor_t *)) { return callback__(self, visitor); }
	// void gocef_frame_get_text(cef_frame_t * self, cef_string_visitor_t * visitor, void (CEF_CALLBACK *callback__)(cef_frame_t *, cef_string_visitor_t *)) { return callback__(self, visitor); }
	// void gocef_frame_load_request(cef_frame_t * self, cef_request_t * request, void (CEF_CALLBACK *callback__)(cef_frame_t *, cef_request_t *)) { return callback__(self, request); }
	// void gocef_frame_load_url(cef_frame_t * self, cef_string_t * url, void (CEF_CALLBACK *callback__)(cef_frame_t *, cef_string_t *)) { return callback__(self, url); }
	// void gocef_frame_load_string(cef_frame_t * self, cef_string_t * string_val, cef_string_t * url, void (CEF_CALLBACK *callback__)(cef_frame_t *, cef_string_t *, cef_string_t *)) { return callback__(self, string_val, url); }
	// void gocef_frame_execute_java_script(cef_frame_t * self, cef_string_t * code, cef_string_t * script_url, int start_line, void (CEF_CALLBACK *callback__)(cef_frame_t *, cef_string_t *, cef_string_t *, int)) { return callback__(self, code, script_url, start_line); }
	// int gocef_frame_is_main(cef_frame_t * self, int (CEF_CALLBACK *callback__)(cef_frame_t *)) { return callback__(self); }
	// int gocef_frame_is_focused(cef_frame_t * self, int (CEF_CALLBACK *callback__)(cef_frame_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_frame_get_name(cef_frame_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_frame_t *)) { return callback__(self); }
	// int64 gocef_frame_get_identifier(cef_frame_t * self, int64 (CEF_CALLBACK *callback__)(cef_frame_t *)) { return callback__(self); }
	// cef_frame_t * gocef_frame_get_parent(cef_frame_t * self, cef_frame_t * (CEF_CALLBACK *callback__)(cef_frame_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_frame_get_url(cef_frame_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_frame_t *)) { return callback__(self); }
	// cef_browser_t * gocef_frame_get_browser(cef_frame_t * self, cef_browser_t * (CEF_CALLBACK *callback__)(cef_frame_t *)) { return callback__(self); }
	// cef_v8context_t * gocef_frame_get_v8context(cef_frame_t * self, cef_v8context_t * (CEF_CALLBACK *callback__)(cef_frame_t *)) { return callback__(self); }
	// void gocef_frame_visit_dom(cef_frame_t * self, cef_domvisitor_t * visitor, void (CEF_CALLBACK *callback__)(cef_frame_t *, cef_domvisitor_t *)) { return callback__(self, visitor); }
	"C"
)

// Frame (cef_frame_t from include/capi/cef_frame_capi.h)
// Structure used to represent a frame in the browser window. When used in the
// browser process the functions of this structure may be called on any thread
// unless otherwise indicated in the comments. When used in the render process
// the functions of this structure may only be called on the main thread.
type Frame C.cef_frame_t

func (d *Frame) toNative() *C.cef_frame_t {
	return (*C.cef_frame_t)(d)
}

// Base (base)
// Base structure.
func (d *Frame) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// IsValid (is_valid)
// True if this object is currently attached to a valid frame.
func (d *Frame) IsValid() int32 {
	return int32(C.gocef_frame_is_valid(d.toNative(), d.is_valid))
}

// Undo (undo)
// Execute undo in this frame.
func (d *Frame) Undo() {
	C.gocef_frame_undo(d.toNative(), d.undo)
}

// Redo (redo)
// Execute redo in this frame.
func (d *Frame) Redo() {
	C.gocef_frame_redo(d.toNative(), d.redo)
}

// Cut (cut)
// Execute cut in this frame.
func (d *Frame) Cut() {
	C.gocef_frame_cut(d.toNative(), d.cut)
}

// Copy (copy)
// Execute copy in this frame.
func (d *Frame) Copy() {
	C.gocef_frame_copy(d.toNative(), d.copy)
}

// Paste (paste)
// Execute paste in this frame.
func (d *Frame) Paste() {
	C.gocef_frame_paste(d.toNative(), d.paste)
}

// Del (del)
// Execute delete in this frame.
func (d *Frame) Del() {
	C.gocef_frame_del(d.toNative(), d.del)
}

// SelectAll (select_all)
// Execute select all in this frame.
func (d *Frame) SelectAll() {
	C.gocef_frame_select_all(d.toNative(), d.select_all)
}

// ViewSource (view_source)
// Save this frame's HTML source to a temporary file and open it in the
// default text viewing application. This function can only be called from the
// browser process.
func (d *Frame) ViewSource() {
	C.gocef_frame_view_source(d.toNative(), d.view_source)
}

// GetSource (get_source)
// Retrieve this frame's HTML source as a string sent to the specified
// visitor.
func (d *Frame) GetSource(visitor *StringVisitor) {
	C.gocef_frame_get_source(d.toNative(), visitor.toNative(), d.get_source)
}

// GetText (get_text)
// Retrieve this frame's display text as a string sent to the specified
// visitor.
func (d *Frame) GetText(visitor *StringVisitor) {
	C.gocef_frame_get_text(d.toNative(), visitor.toNative(), d.get_text)
}

// LoadRequest (load_request)
// Load the request represented by the |request| object.
func (d *Frame) LoadRequest(request *Request) {
	C.gocef_frame_load_request(d.toNative(), request.toNative(), d.load_request)
}

// LoadUrl (load_url)
// Load the specified |url|.
func (d *Frame) LoadUrl(url string) {
	url_ := C.cef_string_userfree_alloc()
	setCEFStr(url, url_)
	defer func() {
		C.cef_string_userfree_free(url_)
	}()
	C.gocef_frame_load_url(d.toNative(), (*C.cef_string_t)(url_), d.load_url)
}

// LoadString (load_string)
// Load the contents of |string_val| with the specified dummy |url|. |url|
// should have a standard scheme (for example, http scheme) or behaviors like
// link clicks and web security restrictions may not behave as expected.
func (d *Frame) LoadString(string_val, url string) {
	string_val_ := C.cef_string_userfree_alloc()
	setCEFStr(string_val, string_val_)
	defer func() {
		C.cef_string_userfree_free(string_val_)
	}()
	url_ := C.cef_string_userfree_alloc()
	setCEFStr(url, url_)
	defer func() {
		C.cef_string_userfree_free(url_)
	}()
	C.gocef_frame_load_string(d.toNative(), (*C.cef_string_t)(string_val_), (*C.cef_string_t)(url_), d.load_string)
}

// ExecuteJavaScript (execute_java_script)
// Execute a string of JavaScript code in this frame. The |script_url|
// parameter is the URL where the script in question can be found, if any. The
// renderer may request this URL to show the developer the source of the
// error.  The |start_line| parameter is the base line number to use for error
// reporting.
func (d *Frame) ExecuteJavaScript(code, script_url string, start_line int32) {
	code_ := C.cef_string_userfree_alloc()
	setCEFStr(code, code_)
	defer func() {
		C.cef_string_userfree_free(code_)
	}()
	script_url_ := C.cef_string_userfree_alloc()
	setCEFStr(script_url, script_url_)
	defer func() {
		C.cef_string_userfree_free(script_url_)
	}()
	C.gocef_frame_execute_java_script(d.toNative(), (*C.cef_string_t)(code_), (*C.cef_string_t)(script_url_), C.int(start_line), d.execute_java_script)
}

// IsMain (is_main)
// Returns true (1) if this is the main (top-level) frame.
func (d *Frame) IsMain() int32 {
	return int32(C.gocef_frame_is_main(d.toNative(), d.is_main))
}

// IsFocused (is_focused)
// Returns true (1) if this is the focused frame.
func (d *Frame) IsFocused() int32 {
	return int32(C.gocef_frame_is_focused(d.toNative(), d.is_focused))
}

// GetName (get_name)
// Returns the name for this frame. If the frame has an assigned name (for
// example, set via the iframe "name" attribute) then that value will be
// returned. Otherwise a unique name will be constructed based on the frame
// parent hierarchy. The main (top-level) frame will always have an NULL name
// value.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Frame) GetName() string {
	return cefuserfreestrToString(C.gocef_frame_get_name(d.toNative(), d.get_name))
}

// GetIdentifier (get_identifier)
// Returns the globally unique identifier for this frame or < 0 if the
// underlying frame does not yet exist.
func (d *Frame) GetIdentifier() int64 {
	return int64(C.gocef_frame_get_identifier(d.toNative(), d.get_identifier))
}

// GetParent (get_parent)
// Returns the parent of this frame or NULL if this is the main (top-level)
// frame.
func (d *Frame) GetParent() *Frame {
	return (*Frame)(C.gocef_frame_get_parent(d.toNative(), d.get_parent))
}

// GetUrl (get_url)
// Returns the URL currently loaded in this frame.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Frame) GetUrl() string {
	return cefuserfreestrToString(C.gocef_frame_get_url(d.toNative(), d.get_url))
}

// GetBrowser (get_browser)
// Returns the browser that this frame belongs to.
func (d *Frame) GetBrowser() *Browser {
	return (*Browser)(C.gocef_frame_get_browser(d.toNative(), d.get_browser))
}

// GetV8context (get_v8context)
// Get the V8 context associated with the frame. This function can only be
// called from the render process.
func (d *Frame) GetV8context() *V8context {
	return (*V8context)(C.gocef_frame_get_v8context(d.toNative(), d.get_v8context))
}

// VisitDom (visit_dom)
// Visit the DOM document. This function can only be called from the render
// process.
func (d *Frame) VisitDom(visitor *Domvisitor) {
	C.gocef_frame_visit_dom(d.toNative(), visitor.toNative(), d.visit_dom)
}
