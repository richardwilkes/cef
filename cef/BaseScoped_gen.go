// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// void gocef_base_scoped_del(cef_base_scoped_t * self, void (CEF_CALLBACK *callback__)(cef_base_scoped_t *)) { return callback__(self); }
	"C"
)

// BaseScoped (cef_base_scoped_t from include/capi/cef_base_capi.h)
// All scoped framework structures must include this structure first.
type BaseScoped C.cef_base_scoped_t

func (d *BaseScoped) toNative() *C.cef_base_scoped_t {
	return (*C.cef_base_scoped_t)(d)
}

// Size (size)
// Size of the data structure.
func (d *BaseScoped) Size() uint64 {
	return uint64(d.size)
}

// Del (del)
// Called to delete this object. May be NULL if the object is not owned.
func (d *BaseScoped) Del() {
	C.gocef_base_scoped_del(d.toNative(), d.del)
}
