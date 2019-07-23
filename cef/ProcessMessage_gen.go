// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// int gocef_process_message_is_valid(cef_process_message_t * self, int (CEF_CALLBACK *callback__)(cef_process_message_t *)) { return callback__(self); }
	// int gocef_process_message_is_read_only(cef_process_message_t * self, int (CEF_CALLBACK *callback__)(cef_process_message_t *)) { return callback__(self); }
	// cef_process_message_t * gocef_process_message_copy(cef_process_message_t * self, cef_process_message_t * (CEF_CALLBACK *callback__)(cef_process_message_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_process_message_get_name(cef_process_message_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_process_message_t *)) { return callback__(self); }
	// cef_list_value_t * gocef_process_message_get_argument_list(cef_process_message_t * self, cef_list_value_t * (CEF_CALLBACK *callback__)(cef_process_message_t *)) { return callback__(self); }
	"C"
)

// ProcessMessage (cef_process_message_t from include/capi/cef_process_message_capi.h)
// Structure representing a message. Can be used on any process and thread.
type ProcessMessage C.cef_process_message_t

func (d *ProcessMessage) toNative() *C.cef_process_message_t {
	return (*C.cef_process_message_t)(d)
}

// Base (base)
// Base structure.
func (d *ProcessMessage) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// IsValid (is_valid)
// Returns true (1) if this object is valid. Do not call any other functions
// if this function returns false (0).
func (d *ProcessMessage) IsValid() int32 {
	return int32(C.gocef_process_message_is_valid(d.toNative(), d.is_valid))
}

// IsReadOnly (is_read_only)
// Returns true (1) if the values of this object are read-only. Some APIs may
// expose read-only objects.
func (d *ProcessMessage) IsReadOnly() int32 {
	return int32(C.gocef_process_message_is_read_only(d.toNative(), d.is_read_only))
}

// Copy (copy)
// Returns a writable copy of this object.
func (d *ProcessMessage) Copy() *ProcessMessage {
	return (*ProcessMessage)(C.gocef_process_message_copy(d.toNative(), d.copy))
}

// GetName (get_name)
// Returns the message name.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *ProcessMessage) GetName() string {
	return cefuserfreestrToString(C.gocef_process_message_get_name(d.toNative(), d.get_name))
}

// GetArgumentList (get_argument_list)
// Returns the list of arguments.
func (d *ProcessMessage) GetArgumentList() *ListValue {
	return (*ListValue)(C.gocef_process_message_get_argument_list(d.toNative(), d.get_argument_list))
}
