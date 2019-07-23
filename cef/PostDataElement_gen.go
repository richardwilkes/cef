// Code created from "class.go.tmpl" - don't edit by hand

package cef

import "unsafe"

import (
	// #include "capi_gen.h"
	// int gocef_post_data_element_is_read_only(cef_post_data_element_t * self, int (CEF_CALLBACK *callback__)(cef_post_data_element_t *)) { return callback__(self); }
	// void gocef_post_data_element_set_to_empty(cef_post_data_element_t * self, void (CEF_CALLBACK *callback__)(cef_post_data_element_t *)) { return callback__(self); }
	// void gocef_post_data_element_set_to_file(cef_post_data_element_t * self, cef_string_t * fileName, void (CEF_CALLBACK *callback__)(cef_post_data_element_t *, cef_string_t *)) { return callback__(self, fileName); }
	// void gocef_post_data_element_set_to_bytes(cef_post_data_element_t * self, size_t size, void * bytes, void (CEF_CALLBACK *callback__)(cef_post_data_element_t *, size_t, void *)) { return callback__(self, size, bytes); }
	// cef_postdataelement_type_t gocef_post_data_element_get_type(cef_post_data_element_t * self, cef_postdataelement_type_t (CEF_CALLBACK *callback__)(cef_post_data_element_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_post_data_element_get_file(cef_post_data_element_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_post_data_element_t *)) { return callback__(self); }
	// size_t gocef_post_data_element_get_bytes_count(cef_post_data_element_t * self, size_t (CEF_CALLBACK *callback__)(cef_post_data_element_t *)) { return callback__(self); }
	// size_t gocef_post_data_element_get_bytes(cef_post_data_element_t * self, size_t size, void * bytes, size_t (CEF_CALLBACK *callback__)(cef_post_data_element_t *, size_t, void *)) { return callback__(self, size, bytes); }
	"C"
)

// PostDataElement (cef_post_data_element_t from include/capi/cef_request_capi.h)
// Structure used to represent a single element in the request post data. The
// functions of this structure may be called on any thread.
type PostDataElement C.cef_post_data_element_t

func (d *PostDataElement) toNative() *C.cef_post_data_element_t {
	return (*C.cef_post_data_element_t)(d)
}

// Base (base)
// Base structure.
func (d *PostDataElement) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// IsReadOnly (is_read_only)
// Returns true (1) if this object is read-only.
func (d *PostDataElement) IsReadOnly() int32 {
	return int32(C.gocef_post_data_element_is_read_only(d.toNative(), d.is_read_only))
}

// SetToEmpty (set_to_empty)
// Remove all contents from the post data element.
func (d *PostDataElement) SetToEmpty() {
	C.gocef_post_data_element_set_to_empty(d.toNative(), d.set_to_empty)
}

// SetToFile (set_to_file)
// The post data element will represent a file.
func (d *PostDataElement) SetToFile(fileName string) {
	fileName_ := C.cef_string_userfree_alloc()
	setCEFStr(fileName, fileName_)
	defer func() {
		C.cef_string_userfree_free(fileName_)
	}()
	C.gocef_post_data_element_set_to_file(d.toNative(), (*C.cef_string_t)(fileName_), d.set_to_file)
}

// SetToBytes (set_to_bytes)
// The post data element will represent bytes.  The bytes passed in will be
// copied.
func (d *PostDataElement) SetToBytes(size uint64, bytes unsafe.Pointer) {
	C.gocef_post_data_element_set_to_bytes(d.toNative(), C.size_t(size), bytes, d.set_to_bytes)
}

// GetType (get_type)
// Return the type of this post data element.
func (d *PostDataElement) GetType() PostdataelementType {
	return PostdataelementType(C.gocef_post_data_element_get_type(d.toNative(), d.get_type))
}

// GetFile (get_file)
// Return the file name.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *PostDataElement) GetFile() string {
	return cefuserfreestrToString(C.gocef_post_data_element_get_file(d.toNative(), d.get_file))
}

// GetBytesCount (get_bytes_count)
// Return the number of bytes.
func (d *PostDataElement) GetBytesCount() uint64 {
	return uint64(C.gocef_post_data_element_get_bytes_count(d.toNative(), d.get_bytes_count))
}

// GetBytes (get_bytes)
// Read up to |size| bytes into |bytes| and return the number of bytes
// actually read.
func (d *PostDataElement) GetBytes(size uint64, bytes unsafe.Pointer) uint64 {
	return uint64(C.gocef_post_data_element_get_bytes(d.toNative(), C.size_t(size), bytes, d.get_bytes))
}
