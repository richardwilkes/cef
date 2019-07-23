// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// int gocef_v8interceptor_get_byname(cef_v8interceptor_t * self, cef_string_t * name, cef_v8value_t * object, cef_v8value_t ** retval, cef_string_t * exception, int (CEF_CALLBACK *callback__)(cef_v8interceptor_t *, cef_string_t *, cef_v8value_t *, cef_v8value_t **, cef_string_t *)) { return callback__(self, name, object, retval, exception); }
	// int gocef_v8interceptor_get_byindex(cef_v8interceptor_t * self, int index, cef_v8value_t * object, cef_v8value_t ** retval, cef_string_t * exception, int (CEF_CALLBACK *callback__)(cef_v8interceptor_t *, int, cef_v8value_t *, cef_v8value_t **, cef_string_t *)) { return callback__(self, index, object, retval, exception); }
	// int gocef_v8interceptor_set_byname(cef_v8interceptor_t * self, cef_string_t * name, cef_v8value_t * object, cef_v8value_t * value, cef_string_t * exception, int (CEF_CALLBACK *callback__)(cef_v8interceptor_t *, cef_string_t *, cef_v8value_t *, cef_v8value_t *, cef_string_t *)) { return callback__(self, name, object, value, exception); }
	// int gocef_v8interceptor_set_byindex(cef_v8interceptor_t * self, int index, cef_v8value_t * object, cef_v8value_t * value, cef_string_t * exception, int (CEF_CALLBACK *callback__)(cef_v8interceptor_t *, int, cef_v8value_t *, cef_v8value_t *, cef_string_t *)) { return callback__(self, index, object, value, exception); }
	"C"
)

// V8interceptor (cef_v8interceptor_t from include/capi/cef_v8_capi.h)
// Structure that should be implemented to handle V8 interceptor calls. The
// functions of this structure will be called on the thread associated with the
// V8 interceptor. Interceptor's named property handlers (with first argument of
// type CefString) are called when object is indexed by string. Indexed property
// handlers (with first argument of type int) are called when object is indexed
// by integer.
type V8interceptor C.cef_v8interceptor_t

func (d *V8interceptor) toNative() *C.cef_v8interceptor_t {
	return (*C.cef_v8interceptor_t)(d)
}

// Base (base)
// Base structure.
func (d *V8interceptor) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// GetByname (get_byname)
// Handle retrieval of the interceptor value identified by |name|. |object| is
// the receiver ('this' object) of the interceptor. If retrieval succeeds, set
// |retval| to the return value. If the requested value does not exist, don't
// set either |retval| or |exception|. If retrieval fails, set |exception| to
// the exception that will be thrown. If the property has an associated
// accessor, it will be called only if you don't set |retval|. Return true (1)
// if interceptor retrieval was handled, false (0) otherwise.
func (d *V8interceptor) GetByname(name string, object *V8value, retval **V8value, exception *string) int32 {
	name_ := C.cef_string_userfree_alloc()
	setCEFStr(name, name_)
	defer func() {
		C.cef_string_userfree_free(name_)
	}()
	retval_ := (*retval).toNative()
	exception_ := C.cef_string_userfree_alloc()
	setCEFStr(*exception, exception_)
	defer func() {
		*exception = cefstrToString(exception_)
		C.cef_string_userfree_free(exception_)
	}()
	return int32(C.gocef_v8interceptor_get_byname(d.toNative(), (*C.cef_string_t)(name_), object.toNative(), &retval_, (*C.cef_string_t)(exception_), d.get_byname))
}

// GetByindex (get_byindex)
// Handle retrieval of the interceptor value identified by |index|. |object|
// is the receiver ('this' object) of the interceptor. If retrieval succeeds,
// set |retval| to the return value. If the requested value does not exist,
// don't set either |retval| or |exception|. If retrieval fails, set
// |exception| to the exception that will be thrown. Return true (1) if
// interceptor retrieval was handled, false (0) otherwise.
func (d *V8interceptor) GetByindex(index int32, object *V8value, retval **V8value, exception *string) int32 {
	retval_ := (*retval).toNative()
	exception_ := C.cef_string_userfree_alloc()
	setCEFStr(*exception, exception_)
	defer func() {
		*exception = cefstrToString(exception_)
		C.cef_string_userfree_free(exception_)
	}()
	return int32(C.gocef_v8interceptor_get_byindex(d.toNative(), C.int(index), object.toNative(), &retval_, (*C.cef_string_t)(exception_), d.get_byindex))
}

// SetByname (set_byname)
// Handle assignment of the interceptor value identified by |name|. |object|
// is the receiver ('this' object) of the interceptor. |value| is the new
// value being assigned to the interceptor. If assignment fails, set
// |exception| to the exception that will be thrown. This setter will always
// be called, even when the property has an associated accessor. Return true
// (1) if interceptor assignment was handled, false (0) otherwise.
func (d *V8interceptor) SetByname(name string, object, value *V8value, exception *string) int32 {
	name_ := C.cef_string_userfree_alloc()
	setCEFStr(name, name_)
	defer func() {
		C.cef_string_userfree_free(name_)
	}()
	exception_ := C.cef_string_userfree_alloc()
	setCEFStr(*exception, exception_)
	defer func() {
		*exception = cefstrToString(exception_)
		C.cef_string_userfree_free(exception_)
	}()
	return int32(C.gocef_v8interceptor_set_byname(d.toNative(), (*C.cef_string_t)(name_), object.toNative(), value.toNative(), (*C.cef_string_t)(exception_), d.set_byname))
}

// SetByindex (set_byindex)
// Handle assignment of the interceptor value identified by |index|. |object|
// is the receiver ('this' object) of the interceptor. |value| is the new
// value being assigned to the interceptor. If assignment fails, set
// |exception| to the exception that will be thrown. Return true (1) if
// interceptor assignment was handled, false (0) otherwise.
func (d *V8interceptor) SetByindex(index int32, object, value *V8value, exception *string) int32 {
	exception_ := C.cef_string_userfree_alloc()
	setCEFStr(*exception, exception_)
	defer func() {
		*exception = cefstrToString(exception_)
		C.cef_string_userfree_free(exception_)
	}()
	return int32(C.gocef_v8interceptor_set_byindex(d.toNative(), C.int(index), object.toNative(), value.toNative(), (*C.cef_string_t)(exception_), d.set_byindex))
}
