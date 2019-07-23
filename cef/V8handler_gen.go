// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// int gocef_v8handler_execute(cef_v8handler_t * self, cef_string_t * name, cef_v8value_t * object, size_t argumentsCount, cef_v8value_t ** arguments, cef_v8value_t ** retval, cef_string_t * exception, int (CEF_CALLBACK *callback__)(cef_v8handler_t *, cef_string_t *, cef_v8value_t *, size_t, cef_v8value_t **, cef_v8value_t **, cef_string_t *)) { return callback__(self, name, object, argumentsCount, arguments, retval, exception); }
	"C"
)

// V8handler (cef_v8handler_t from include/capi/cef_v8_capi.h)
// Structure that should be implemented to handle V8 function calls. The
// functions of this structure will be called on the thread associated with the
// V8 function.
type V8handler C.cef_v8handler_t

func (d *V8handler) toNative() *C.cef_v8handler_t {
	return (*C.cef_v8handler_t)(d)
}

// Base (base)
// Base structure.
func (d *V8handler) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Execute (execute)
// Handle execution of the function identified by |name|. |object| is the
// receiver ('this' object) of the function. |arguments| is the list of
// arguments passed to the function. If execution succeeds set |retval| to the
// function return value. If execution fails set |exception| to the exception
// that will be thrown. Return true (1) if execution was handled.
func (d *V8handler) Execute(name string, object *V8value, argumentsCount uint64, arguments, retval **V8value, exception *string) int32 {
	name_ := C.cef_string_userfree_alloc()
	setCEFStr(name, name_)
	defer func() {
		C.cef_string_userfree_free(name_)
	}()
	arguments_ := (*arguments).toNative()
	retval_ := (*retval).toNative()
	exception_ := C.cef_string_userfree_alloc()
	setCEFStr(*exception, exception_)
	defer func() {
		*exception = cefstrToString(exception_)
		C.cef_string_userfree_free(exception_)
	}()
	return int32(C.gocef_v8handler_execute(d.toNative(), (*C.cef_string_t)(name_), object.toNative(), C.size_t(argumentsCount), &arguments_, &retval_, (*C.cef_string_t)(exception_), d.execute))
}
