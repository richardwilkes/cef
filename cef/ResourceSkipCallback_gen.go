// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

import (
	// #include "ResourceSkipCallback_gen.h"
	"C"
)

// ResourceSkipCallbackProxy defines methods required for using ResourceSkipCallback.
type ResourceSkipCallbackProxy interface {
	Cont(self *ResourceSkipCallback, bytes_skipped int64)
}

// ResourceSkipCallback (cef_resource_skip_callback_t from include/capi/cef_resource_handler_capi.h)
// Callback for asynchronous continuation of cef_resource_handler_t::skip().
type ResourceSkipCallback C.cef_resource_skip_callback_t

// NewResourceSkipCallback creates a new ResourceSkipCallback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewResourceSkipCallback(proxy ResourceSkipCallbackProxy) *ResourceSkipCallback {
	result := (*ResourceSkipCallback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_resource_skip_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_resource_skip_callback_proxy(result.toNative())
	}
	return result
}

func (d *ResourceSkipCallback) toNative() *C.cef_resource_skip_callback_t {
	return (*C.cef_resource_skip_callback_t)(d)
}

func lookupResourceSkipCallbackProxy(obj *BaseRefCounted) ResourceSkipCallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(ResourceSkipCallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type ResourceSkipCallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *ResourceSkipCallback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Cont (cont)
// Callback for asynchronous continuation of skip(). If |bytes_skipped| > 0
// then either skip() will be called again until the requested number of bytes
// have been skipped or the request will proceed. If |bytes_skipped| <= 0 the
// request will fail with ERR_REQUEST_RANGE_NOT_SATISFIABLE.
func (d *ResourceSkipCallback) Cont(bytes_skipped int64) {
	lookupResourceSkipCallbackProxy(d.Base()).Cont(d, bytes_skipped)
}

//nolint:gocritic
//export gocef_resource_skip_callback_cont
func gocef_resource_skip_callback_cont(self *C.cef_resource_skip_callback_t, bytes_skipped C.int64) {
	me__ := (*ResourceSkipCallback)(self)
	proxy__ := lookupResourceSkipCallbackProxy(me__.Base())
	proxy__.Cont(me__, int64(bytes_skipped))
}
