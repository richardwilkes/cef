package cef

import (
	// #include "PrintJobCallback_gen.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

// PrintJobCallbackProxy defines methods required for using PrintJobCallback.
type PrintJobCallbackProxy interface {
	Cont(self *PrintJobCallback)
}

// PrintJobCallback (cef_print_job_callback_t from include/capi/cef_print_handler_capi.h)
// Callback structure for asynchronous continuation of print job requests.
type PrintJobCallback C.cef_print_job_callback_t

// NewPrintJobCallback creates a new PrintJobCallback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewPrintJobCallback(proxy PrintJobCallbackProxy) *PrintJobCallback {
	result := (*PrintJobCallback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_print_job_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_print_job_callback_proxy(result.toNative())
	}
	return result
}

func (d *PrintJobCallback) toNative() *C.cef_print_job_callback_t {
	return (*C.cef_print_job_callback_t)(d)
}

func lookupPrintJobCallbackProxy(obj *BaseRefCounted) PrintJobCallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(PrintJobCallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type PrintJobCallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *PrintJobCallback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Cont (cont)
// Indicate completion of the print job.
func (d *PrintJobCallback) Cont() {
	lookupPrintJobCallbackProxy(d.Base()).Cont(d)
}

//export gocef_print_job_callback_cont
func gocef_print_job_callback_cont(self *C.cef_print_job_callback_t) {
	me__ := (*PrintJobCallback)(self)
	proxy__ := lookupPrintJobCallbackProxy(me__.Base())
	proxy__.Cont(me__)
}
