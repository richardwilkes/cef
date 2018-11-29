package cef

import (
	// #include "CompletionCallback_gen.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

// CompletionCallbackProxy defines methods required for using CompletionCallback.
type CompletionCallbackProxy interface {
	OnComplete(self *CompletionCallback)
}

// CompletionCallback (cef_completion_callback_t from include/capi/cef_callback_capi.h)
// Generic callback structure used for asynchronous completion.
type CompletionCallback C.cef_completion_callback_t

// NewCompletionCallback creates a new CompletionCallback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewCompletionCallback(proxy CompletionCallbackProxy) *CompletionCallback {
	result := (*CompletionCallback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_completion_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_completion_callback_proxy(result.toNative())
	}
	return result
}

func (d *CompletionCallback) toNative() *C.cef_completion_callback_t {
	return (*C.cef_completion_callback_t)(d)
}

func lookupCompletionCallbackProxy(obj *BaseRefCounted) CompletionCallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(CompletionCallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type CompletionCallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *CompletionCallback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnComplete (on_complete)
// Method that will be called once the task is complete.
func (d *CompletionCallback) OnComplete() {
	lookupCompletionCallbackProxy(d.Base()).OnComplete(d)
}

//export gocef_completion_callback_on_complete
func gocef_completion_callback_on_complete(self *C.cef_completion_callback_t) {
	me__ := (*CompletionCallback)(self)
	proxy__ := lookupCompletionCallbackProxy(me__.Base())
	proxy__.OnComplete(me__)
}
