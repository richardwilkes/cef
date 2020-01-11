// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// int gocef_v8stack_frame_is_valid(cef_v8stack_frame_t * self, int (CEF_CALLBACK *callback__)(cef_v8stack_frame_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_v8stack_frame_get_script_name(cef_v8stack_frame_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_v8stack_frame_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_v8stack_frame_get_script_name_or_source_url(cef_v8stack_frame_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_v8stack_frame_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_v8stack_frame_get_function_name(cef_v8stack_frame_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_v8stack_frame_t *)) { return callback__(self); }
	// int gocef_v8stack_frame_get_line_number(cef_v8stack_frame_t * self, int (CEF_CALLBACK *callback__)(cef_v8stack_frame_t *)) { return callback__(self); }
	// int gocef_v8stack_frame_get_column(cef_v8stack_frame_t * self, int (CEF_CALLBACK *callback__)(cef_v8stack_frame_t *)) { return callback__(self); }
	// int gocef_v8stack_frame_is_eval(cef_v8stack_frame_t * self, int (CEF_CALLBACK *callback__)(cef_v8stack_frame_t *)) { return callback__(self); }
	// int gocef_v8stack_frame_is_constructor(cef_v8stack_frame_t * self, int (CEF_CALLBACK *callback__)(cef_v8stack_frame_t *)) { return callback__(self); }
	"C"
)

// V8stackFrame (cef_v8stack_frame_t from include/capi/cef_v8_capi.h)
// Structure representing a V8 stack frame handle. V8 handles can only be
// accessed from the thread on which they are created. Valid threads for
// creating a V8 handle include the render process main thread (TID_RENDERER)
// and WebWorker threads. A task runner for posting tasks on the associated
// thread can be retrieved via the cef_v8context_t::get_task_runner() function.
type V8stackFrame C.cef_v8stack_frame_t

func (d *V8stackFrame) toNative() *C.cef_v8stack_frame_t {
	return (*C.cef_v8stack_frame_t)(d)
}

// Base (base)
// Base structure.
func (d *V8stackFrame) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// IsValid (is_valid)
// Returns true (1) if the underlying handle is valid and it can be accessed
// on the current thread. Do not call any other functions if this function
// returns false (0).
func (d *V8stackFrame) IsValid() int32 {
	return int32(C.gocef_v8stack_frame_is_valid(d.toNative(), d.is_valid))
}

// GetScriptName (get_script_name)
// Returns the name of the resource script that contains the function.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *V8stackFrame) GetScriptName() string {
	return cefuserfreestrToString(C.gocef_v8stack_frame_get_script_name(d.toNative(), d.get_script_name))
}

// GetScriptNameOrSourceURL (get_script_name_or_source_url)
// Returns the name of the resource script that contains the function or the
// sourceURL value if the script name is undefined and its source ends with a
// "//@ sourceURL=..." string.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *V8stackFrame) GetScriptNameOrSourceURL() string {
	return cefuserfreestrToString(C.gocef_v8stack_frame_get_script_name_or_source_url(d.toNative(), d.get_script_name_or_source_url))
}

// GetFunctionName (get_function_name)
// Returns the name of the function.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *V8stackFrame) GetFunctionName() string {
	return cefuserfreestrToString(C.gocef_v8stack_frame_get_function_name(d.toNative(), d.get_function_name))
}

// GetLineNumber (get_line_number)
// Returns the 1-based line number for the function call or 0 if unknown.
func (d *V8stackFrame) GetLineNumber() int32 {
	return int32(C.gocef_v8stack_frame_get_line_number(d.toNative(), d.get_line_number))
}

// GetColumn (get_column)
// Returns the 1-based column offset on the line for the function call or 0 if
// unknown.
func (d *V8stackFrame) GetColumn() int32 {
	return int32(C.gocef_v8stack_frame_get_column(d.toNative(), d.get_column))
}

// IsEval (is_eval)
// Returns true (1) if the function was compiled using eval().
func (d *V8stackFrame) IsEval() int32 {
	return int32(C.gocef_v8stack_frame_is_eval(d.toNative(), d.is_eval))
}

// IsConstructor (is_constructor)
// Returns true (1) if the function was called as a constructor via "new".
func (d *V8stackFrame) IsConstructor() int32 {
	return int32(C.gocef_v8stack_frame_is_constructor(d.toNative(), d.is_constructor))
}
