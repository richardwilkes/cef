// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// void gocef_base_ref_counted_add_ref(cef_base_ref_counted_t * self, void (CEF_CALLBACK *callback__)(cef_base_ref_counted_t *)) { return callback__(self); }
	// int gocef_base_ref_counted_release(cef_base_ref_counted_t * self, int (CEF_CALLBACK *callback__)(cef_base_ref_counted_t *)) { return callback__(self); }
	// int gocef_base_ref_counted_has_one_ref(cef_base_ref_counted_t * self, int (CEF_CALLBACK *callback__)(cef_base_ref_counted_t *)) { return callback__(self); }
	// int gocef_base_ref_counted_has_at_least_one_ref(cef_base_ref_counted_t * self, int (CEF_CALLBACK *callback__)(cef_base_ref_counted_t *)) { return callback__(self); }
	"C"
)

// BaseRefCounted (cef_base_ref_counted_t from include/capi/cef_base_capi.h)
// All ref-counted framework structures must include this structure first.
type BaseRefCounted C.cef_base_ref_counted_t

func (d *BaseRefCounted) toNative() *C.cef_base_ref_counted_t {
	return (*C.cef_base_ref_counted_t)(d)
}

// Size (size)
// Size of the data structure.
func (d *BaseRefCounted) Size() uint64 {
	return uint64(d.size)
}

// AddRef (add_ref)
// Called to increment the reference count for the object. Should be called
// for every new copy of a pointer to a given object.
func (d *BaseRefCounted) AddRef() {
	C.gocef_base_ref_counted_add_ref(d.toNative(), d.add_ref)
}

// Release (release)
// Called to decrement the reference count for the object. If the reference
// count falls to 0 the object should self-delete. Returns true (1) if the
// resulting reference count is 0.
func (d *BaseRefCounted) Release() int32 {
	return int32(C.gocef_base_ref_counted_release(d.toNative(), d.release))
}

// HasOneRef (has_one_ref)
// Returns true (1) if the current reference count is 1.
func (d *BaseRefCounted) HasOneRef() int32 {
	return int32(C.gocef_base_ref_counted_has_one_ref(d.toNative(), d.has_one_ref))
}

// HasAtLeastOneRef (has_at_least_one_ref)
// Returns true (1) if the current reference count is at least 1.
func (d *BaseRefCounted) HasAtLeastOneRef() int32 {
	return int32(C.gocef_base_ref_counted_has_at_least_one_ref(d.toNative(), d.has_at_least_one_ref))
}
