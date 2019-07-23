// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// int gocef_post_data_is_read_only(cef_post_data_t * self, int (CEF_CALLBACK *callback__)(cef_post_data_t *)) { return callback__(self); }
	// int gocef_post_data_has_excluded_elements(cef_post_data_t * self, int (CEF_CALLBACK *callback__)(cef_post_data_t *)) { return callback__(self); }
	// size_t gocef_post_data_get_element_count(cef_post_data_t * self, size_t (CEF_CALLBACK *callback__)(cef_post_data_t *)) { return callback__(self); }
	// void gocef_post_data_get_elements(cef_post_data_t * self, size_t * elementsCount, cef_post_data_element_t ** elements, void (CEF_CALLBACK *callback__)(cef_post_data_t *, size_t *, cef_post_data_element_t **)) { return callback__(self, elementsCount, elements); }
	// int gocef_post_data_remove_element(cef_post_data_t * self, cef_post_data_element_t * element, int (CEF_CALLBACK *callback__)(cef_post_data_t *, cef_post_data_element_t *)) { return callback__(self, element); }
	// int gocef_post_data_add_element(cef_post_data_t * self, cef_post_data_element_t * element, int (CEF_CALLBACK *callback__)(cef_post_data_t *, cef_post_data_element_t *)) { return callback__(self, element); }
	// void gocef_post_data_remove_elements(cef_post_data_t * self, void (CEF_CALLBACK *callback__)(cef_post_data_t *)) { return callback__(self); }
	"C"
)

// PostData (cef_post_data_t from include/capi/cef_request_capi.h)
// Structure used to represent post data for a web request. The functions of
// this structure may be called on any thread.
type PostData C.cef_post_data_t

func (d *PostData) toNative() *C.cef_post_data_t {
	return (*C.cef_post_data_t)(d)
}

// Base (base)
// Base structure.
func (d *PostData) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// IsReadOnly (is_read_only)
// Returns true (1) if this object is read-only.
func (d *PostData) IsReadOnly() int32 {
	return int32(C.gocef_post_data_is_read_only(d.toNative(), d.is_read_only))
}

// HasExcludedElements (has_excluded_elements)
// Returns true (1) if the underlying POST data includes elements that are not
// represented by this cef_post_data_t object (for example, multi-part file
// upload data). Modifying cef_post_data_t objects with excluded elements may
// result in the request failing.
func (d *PostData) HasExcludedElements() int32 {
	return int32(C.gocef_post_data_has_excluded_elements(d.toNative(), d.has_excluded_elements))
}

// GetElementCount (get_element_count)
// Returns the number of existing post data elements.
func (d *PostData) GetElementCount() uint64 {
	return uint64(C.gocef_post_data_get_element_count(d.toNative(), d.get_element_count))
}

// GetElements (get_elements)
// Retrieve the post data elements.
func (d *PostData) GetElements(elementsCount *uint64, elements **PostDataElement) {
	elements_ := (*elements).toNative()
	C.gocef_post_data_get_elements(d.toNative(), (*C.size_t)(elementsCount), &elements_, d.get_elements)
}

// RemoveElement (remove_element)
// Remove the specified post data element.  Returns true (1) if the removal
// succeeds.
func (d *PostData) RemoveElement(element *PostDataElement) int32 {
	return int32(C.gocef_post_data_remove_element(d.toNative(), element.toNative(), d.remove_element))
}

// AddElement (add_element)
// Add the specified post data element.  Returns true (1) if the add succeeds.
func (d *PostData) AddElement(element *PostDataElement) int32 {
	return int32(C.gocef_post_data_add_element(d.toNative(), element.toNative(), d.add_element))
}

// RemoveElements (remove_elements)
// Remove all existing post data elements.
func (d *PostData) RemoveElements() {
	C.gocef_post_data_remove_elements(d.toNative(), d.remove_elements)
}
