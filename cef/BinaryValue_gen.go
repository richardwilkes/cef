// Code created from "class.go.tmpl" - don't edit by hand

package cef

import "unsafe"

import (
	// #include "capi_gen.h"
	// int gocef_binary_value_is_valid(cef_binary_value_t * self, int (CEF_CALLBACK *callback__)(cef_binary_value_t *)) { return callback__(self); }
	// int gocef_binary_value_is_owned(cef_binary_value_t * self, int (CEF_CALLBACK *callback__)(cef_binary_value_t *)) { return callback__(self); }
	// int gocef_binary_value_is_same(cef_binary_value_t * self, cef_binary_value_t * that, int (CEF_CALLBACK *callback__)(cef_binary_value_t *, cef_binary_value_t *)) { return callback__(self, that); }
	// int gocef_binary_value_is_equal(cef_binary_value_t * self, cef_binary_value_t * that, int (CEF_CALLBACK *callback__)(cef_binary_value_t *, cef_binary_value_t *)) { return callback__(self, that); }
	// cef_binary_value_t * gocef_binary_value_copy(cef_binary_value_t * self, cef_binary_value_t * (CEF_CALLBACK *callback__)(cef_binary_value_t *)) { return callback__(self); }
	// size_t gocef_binary_value_get_size(cef_binary_value_t * self, size_t (CEF_CALLBACK *callback__)(cef_binary_value_t *)) { return callback__(self); }
	// size_t gocef_binary_value_get_data(cef_binary_value_t * self, void * buffer, size_t bufferSize, size_t dataOffset, size_t (CEF_CALLBACK *callback__)(cef_binary_value_t *, void *, size_t, size_t)) { return callback__(self, buffer, bufferSize, dataOffset); }
	"C"
)

// BinaryValue (cef_binary_value_t from include/capi/cef_values_capi.h)
// Structure representing a binary value. Can be used on any process and thread.
type BinaryValue C.cef_binary_value_t

func (d *BinaryValue) toNative() *C.cef_binary_value_t {
	return (*C.cef_binary_value_t)(d)
}

// Base (base)
// Base structure.
func (d *BinaryValue) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// IsValid (is_valid)
// Returns true (1) if this object is valid. This object may become invalid if
// the underlying data is owned by another object (e.g. list or dictionary)
// and that other object is then modified or destroyed. Do not call any other
// functions if this function returns false (0).
func (d *BinaryValue) IsValid() int32 {
	return int32(C.gocef_binary_value_is_valid(d.toNative(), d.is_valid))
}

// IsOwned (is_owned)
// Returns true (1) if this object is currently owned by another object.
func (d *BinaryValue) IsOwned() int32 {
	return int32(C.gocef_binary_value_is_owned(d.toNative(), d.is_owned))
}

// IsSame (is_same)
// Returns true (1) if this object and |that| object have the same underlying
// data.
func (d *BinaryValue) IsSame(that *BinaryValue) int32 {
	return int32(C.gocef_binary_value_is_same(d.toNative(), that.toNative(), d.is_same))
}

// IsEqual (is_equal)
// Returns true (1) if this object and |that| object have an equivalent
// underlying value but are not necessarily the same object.
func (d *BinaryValue) IsEqual(that *BinaryValue) int32 {
	return int32(C.gocef_binary_value_is_equal(d.toNative(), that.toNative(), d.is_equal))
}

// Copy (copy)
// Returns a copy of this object. The data in this object will also be copied.
func (d *BinaryValue) Copy() *BinaryValue {
	return (*BinaryValue)(C.gocef_binary_value_copy(d.toNative(), d.copy))
}

// GetSize (get_size)
// Returns the data size.
func (d *BinaryValue) GetSize() uint64 {
	return uint64(C.gocef_binary_value_get_size(d.toNative(), d.get_size))
}

// GetData (get_data)
// Read up to |buffer_size| number of bytes into |buffer|. Reading begins at
// the specified byte |data_offset|. Returns the number of bytes read.
func (d *BinaryValue) GetData(buffer unsafe.Pointer, bufferSize, dataOffset uint64) uint64 {
	return uint64(C.gocef_binary_value_get_data(d.toNative(), buffer, C.size_t(bufferSize), C.size_t(dataOffset), d.get_data))
}
