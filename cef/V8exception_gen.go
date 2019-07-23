// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// cef_string_userfree_t gocef_v8exception_get_message(cef_v8exception_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_v8exception_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_v8exception_get_source_line(cef_v8exception_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_v8exception_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_v8exception_get_script_resource_name(cef_v8exception_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_v8exception_t *)) { return callback__(self); }
	// int gocef_v8exception_get_line_number(cef_v8exception_t * self, int (CEF_CALLBACK *callback__)(cef_v8exception_t *)) { return callback__(self); }
	// int gocef_v8exception_get_start_position(cef_v8exception_t * self, int (CEF_CALLBACK *callback__)(cef_v8exception_t *)) { return callback__(self); }
	// int gocef_v8exception_get_end_position(cef_v8exception_t * self, int (CEF_CALLBACK *callback__)(cef_v8exception_t *)) { return callback__(self); }
	// int gocef_v8exception_get_start_column(cef_v8exception_t * self, int (CEF_CALLBACK *callback__)(cef_v8exception_t *)) { return callback__(self); }
	// int gocef_v8exception_get_end_column(cef_v8exception_t * self, int (CEF_CALLBACK *callback__)(cef_v8exception_t *)) { return callback__(self); }
	"C"
)

// V8exception (cef_v8exception_t from include/capi/cef_v8_capi.h)
// Structure representing a V8 exception. The functions of this structure may be
// called on any render process thread.
type V8exception C.cef_v8exception_t

func (d *V8exception) toNative() *C.cef_v8exception_t {
	return (*C.cef_v8exception_t)(d)
}

// Base (base)
// Base structure.
func (d *V8exception) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// GetMessage (get_message)
// Returns the exception message.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *V8exception) GetMessage() string {
	return cefuserfreestrToString(C.gocef_v8exception_get_message(d.toNative(), d.get_message))
}

// GetSourceLine (get_source_line)
// Returns the line of source code that the exception occurred within.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *V8exception) GetSourceLine() string {
	return cefuserfreestrToString(C.gocef_v8exception_get_source_line(d.toNative(), d.get_source_line))
}

// GetScriptResourceName (get_script_resource_name)
// Returns the resource name for the script from where the function causing
// the error originates.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *V8exception) GetScriptResourceName() string {
	return cefuserfreestrToString(C.gocef_v8exception_get_script_resource_name(d.toNative(), d.get_script_resource_name))
}

// GetLineNumber (get_line_number)
// Returns the 1-based number of the line where the error occurred or 0 if the
// line number is unknown.
func (d *V8exception) GetLineNumber() int32 {
	return int32(C.gocef_v8exception_get_line_number(d.toNative(), d.get_line_number))
}

// GetStartPosition (get_start_position)
// Returns the index within the script of the first character where the error
// occurred.
func (d *V8exception) GetStartPosition() int32 {
	return int32(C.gocef_v8exception_get_start_position(d.toNative(), d.get_start_position))
}

// GetEndPosition (get_end_position)
// Returns the index within the script of the last character where the error
// occurred.
func (d *V8exception) GetEndPosition() int32 {
	return int32(C.gocef_v8exception_get_end_position(d.toNative(), d.get_end_position))
}

// GetStartColumn (get_start_column)
// Returns the index within the line of the first character where the error
// occurred.
func (d *V8exception) GetStartColumn() int32 {
	return int32(C.gocef_v8exception_get_start_column(d.toNative(), d.get_start_column))
}

// GetEndColumn (get_end_column)
// Returns the index within the line of the last character where the error
// occurred.
func (d *V8exception) GetEndColumn() int32 {
	return int32(C.gocef_v8exception_get_end_column(d.toNative(), d.get_end_column))
}
