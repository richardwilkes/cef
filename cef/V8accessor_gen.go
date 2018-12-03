// Code generated - DO NOT EDIT.

package cef

import (
	// #include "capi_gen.h"
	// int gocef_v8accessor_get(cef_v8accessor_t * self, cef_string_t * name, cef_v8value_t * object, cef_v8value_t ** retval, cef_string_t * exception, int (CEF_CALLBACK *callback__)(cef_v8accessor_t *, cef_string_t *, cef_v8value_t *, cef_v8value_t **, cef_string_t *)) { return callback__(self, name, object, retval, exception); }
	// int gocef_v8accessor_set(cef_v8accessor_t * self, cef_string_t * name, cef_v8value_t * object, cef_v8value_t * value, cef_string_t * exception, int (CEF_CALLBACK *callback__)(cef_v8accessor_t *, cef_string_t *, cef_v8value_t *, cef_v8value_t *, cef_string_t *)) { return callback__(self, name, object, value, exception); }
	"C"
)

// V8accessor (cef_v8accessor_t from include/capi/cef_v8_capi.h)
// Structure that should be implemented to handle V8 accessor calls. Accessor
// identifiers are registered by calling cef_v8value_t::set_value(). The
// functions of this structure will be called on the thread associated with the
// V8 accessor.
type V8accessor C.cef_v8accessor_t

func (d *V8accessor) toNative() *C.cef_v8accessor_t {
	return (*C.cef_v8accessor_t)(d)
}

// Base (base)
// Base structure.
func (d *V8accessor) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Get (get)
// Handle retrieval the accessor value identified by |name|. |object| is the
// receiver ('this' object) of the accessor. If retrieval succeeds set
// |retval| to the return value. If retrieval fails set |exception| to the
// exception that will be thrown. Return true (1) if accessor retrieval was
// handled.
func (d *V8accessor) Get(name string, object *V8value, retval **V8value, exception string) int32 {
	var name_ C.cef_string_t
	setCEFStr(name, &name_)
	retval_ := (*retval).toNative()
	var exception_ C.cef_string_t
	setCEFStr(exception, &exception_)
	return int32(C.gocef_v8accessor_get(d.toNative(), &name_, object.toNative(), &retval_, &exception_, d.get))
}

// Set (set)
// Handle assignment of the accessor value identified by |name|. |object| is
// the receiver ('this' object) of the accessor. |value| is the new value
// being assigned to the accessor. If assignment fails set |exception| to the
// exception that will be thrown. Return true (1) if accessor assignment was
// handled.
func (d *V8accessor) Set(name string, object, value *V8value, exception string) int32 {
	var name_ C.cef_string_t
	setCEFStr(name, &name_)
	var exception_ C.cef_string_t
	setCEFStr(exception, &exception_)
	return int32(C.gocef_v8accessor_set(d.toNative(), &name_, object.toNative(), value.toNative(), &exception_, d.set))
}
