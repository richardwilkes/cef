package cef

import (
	// #include "frame.h"
	"C"

	"github.com/richardwilkes/toolbox/log/jot"
)

type Frame struct {
	native *C.cef_frame_t
}

// Valid returns true if this object is currently attached to a valid frame.
func (f *Frame) Valid() bool {
	return C.gocef_call_int_frame(f.native, f.native.is_valid) != 0
}

// Undo executes undo in this frame.
func (f *Frame) Undo() {
	C.gocef_call_void_frame(f.native, f.native.undo)
}

// Redo executes redo in this frame.
func (f *Frame) Redo() {
	C.gocef_call_void_frame(f.native, f.native.redo)
}

// Cut executes cut in this frame.
func (f *Frame) Cut() {
	C.gocef_call_void_frame(f.native, f.native.cut)
}

// Copy executes copy in this frame.
func (f *Frame) Copy() {
	C.gocef_call_void_frame(f.native, f.native.copy)
}

// Paste executes paste in this frame.
func (f *Frame) Paste() {
	C.gocef_call_void_frame(f.native, f.native.paste)
}

// Del executes delete in this frame.
func (f *Frame) Del() {
	C.gocef_call_void_frame(f.native, f.native.del)
}

// SelectAll executes select all in this frame.
func (f *Frame) SelectAll() {
	C.gocef_call_void_frame(f.native, f.native.select_all)
}

// ViewSource saves this frame's HTML source to a temporary file and opens it
// in the default text viewing application. This function can only be called
// from the browser process.
func (f *Frame) ViewSource() {
	C.gocef_call_void_frame(f.native, f.native.view_source)
}

// GetSource retrieves this frame's HTML source as a string.
func (f *Frame) GetSource(asyncCallback func(string)) {
	C.gocef_call_void_frame_string_visitor(f.native, newAsyncStringVisitor(asyncCallback), f.native.get_source)
}

// GetText retrieves this frame's display text as a string.
func (f *Frame) GetText(asyncCallback func(string)) {
	C.gocef_call_void_frame_string_visitor(f.native, newAsyncStringVisitor(asyncCallback), f.native.get_text)
}

// LoadRequest loads the request into the frame.
func (f *Frame) LoadRequest(request interface{}) {
	jot.Fatal(1, "Frame.LoadRequest has not been implemented yet")
	// RAW: Implement
	//
	// void(CEF_CALLBACK* load_request)(struct _cef_frame_t* self, struct _cef_request_t* request);
}

// LoadURL loads the url into the frame.
func (f *Frame) LoadURL(url string) {
	C.gocef_call_void_frame_string(f.native, newCEFStr(url), f.native.load_url)
}

// LoadString loads the contents of 'value' with the specified dummy 'url'.
// The URL should have a standard scheme (for example, the http scheme) or
// behaviors like link clicks and web security restrictions may not behave as
// expected.
func (f *Frame) LoadString(value, url string) {
	C.gocef_call_void_frame_string_string(f.native, newCEFStr(value), newCEFStr(url), f.native.load_string)
}

// ExecuteJavaScript executes a string of JavaScript code in this frame. The
// 'url' parameter is the URL where the script in question can be found, if
// any. The renderer may request this URL to show the developer the source of
// the error. The 'startLine' parameter is the base line number to use for
// error reporting.
func (f *Frame) ExecuteJavaScript(script, url string, startLine int) {
	C.gocef_call_void_frame_string_string_int(f.native, newCEFStr(script), newCEFStr(url), C.int(startLine), f.native.execute_java_script)
}

// Main returns true if this is the main (top-level) frame.
func (f *Frame) Main() bool {
	return C.gocef_call_int_frame(f.native, f.native.is_main) != 0
}

// Focused returns true if this is the focused frame.
func (f *Frame) Focused() bool {
	return C.gocef_call_int_frame(f.native, f.native.is_focused) != 0
}

// Name returns the name for this frame. If the frame has an assigned name
// (for example, set via the iframe "name" attribute), then that value will be
// returned, otherwise a unique name will be constructed based on the frame
// parent hierarchy. The main (top-level) frame will always have an empty name
// value.
func (f *Frame) Name() string {
	return cefuserfreestrToString(C.gocef_call_string_frame(f.native, f.native.get_name))
}

// ID returns the globally unique identifier for this frame or < 0 if the
// underlying frame does not yet exist.
func (f *Frame) ID() int64 {
	return int64(C.gocef_call_int_frame(f.native, f.native.get_identifier))
}

// Parent returns the parent of this frame or nil if this is the main
// (top-level) frame.
func (f *Frame) Parent() *Frame {
	if parent := C.gocef_call_frame_frame(f.native, f.native.get_parent); parent != nil {
		return &Frame{native: parent}
	}
	return nil
}

// URL returns the URL currently loaded in this frame.
func (f *Frame) URL() string {
	return cefuserfreestrToString(C.gocef_call_string_frame(f.native, f.native.get_url))
}

// Browser returns the browser that this frame belongs to.
func (f *Frame) Browser() *Browser {
	return &Browser{native: C.gocef_call_browser_frame(f.native, f.native.get_browser)}
}
