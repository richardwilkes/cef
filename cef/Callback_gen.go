// Code generated - DO NOT EDIT.

package cef

import (
	// #include "Callback_gen.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

// CallbackProxy defines methods required for using Callback.
type CallbackProxy interface {
	Cont(self *Callback)
	Cancel(self *Callback)
}

// Callback (cef_callback_t from include/capi/cef_callback_capi.h)
// Generic callback structure used for asynchronous continuation.
type Callback C.cef_callback_t

// NewCallback creates a new Callback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewCallback(proxy CallbackProxy) *Callback {
	result := (*Callback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_callback_proxy(result.toNative())
	}
	return result
}

func (d *Callback) toNative() *C.cef_callback_t {
	return (*C.cef_callback_t)(d)
}

func lookupCallbackProxy(obj *BaseRefCounted) CallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(CallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type CallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *Callback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Cont (cont)
// Continue processing.
func (d *Callback) Cont() {
	lookupCallbackProxy(d.Base()).Cont(d)
}

//export gocef_callback_cont
func gocef_callback_cont(self *C.cef_callback_t) {
	me__ := (*Callback)(self)
	proxy__ := lookupCallbackProxy(me__.Base())
	proxy__.Cont(me__)
}

// Cancel (cancel)
// Cancel processing.
func (d *Callback) Cancel() {
	lookupCallbackProxy(d.Base()).Cancel(d)
}

//export gocef_callback_cancel
func gocef_callback_cancel(self *C.cef_callback_t) {
	me__ := (*Callback)(self)
	proxy__ := lookupCallbackProxy(me__.Base())
	proxy__.Cancel(me__)
}
