package cef

import (
	// #include "PrintDialogCallback_gen.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

// PrintDialogCallbackProxy defines methods required for using PrintDialogCallback.
type PrintDialogCallbackProxy interface {
	Cont(self *PrintDialogCallback, settings *PrintSettings)
	Cancel(self *PrintDialogCallback)
}

// PrintDialogCallback (cef_print_dialog_callback_t from include/capi/cef_print_handler_capi.h)
// Callback structure for asynchronous continuation of print dialog requests.
type PrintDialogCallback C.cef_print_dialog_callback_t

// NewPrintDialogCallback creates a new PrintDialogCallback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewPrintDialogCallback(proxy PrintDialogCallbackProxy) *PrintDialogCallback {
	result := (*PrintDialogCallback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_print_dialog_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_print_dialog_callback_proxy(result.toNative())
	}
	return result
}

func (d *PrintDialogCallback) toNative() *C.cef_print_dialog_callback_t {
	return (*C.cef_print_dialog_callback_t)(d)
}

func lookupPrintDialogCallbackProxy(obj *BaseRefCounted) PrintDialogCallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(PrintDialogCallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type PrintDialogCallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *PrintDialogCallback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Cont (cont)
// Continue printing with the specified |settings|.
func (d *PrintDialogCallback) Cont(settings *PrintSettings) {
	lookupPrintDialogCallbackProxy(d.Base()).Cont(d, settings)
}

//export gocef_print_dialog_callback_cont
func gocef_print_dialog_callback_cont(self *C.cef_print_dialog_callback_t, settings *C.cef_print_settings_t) {
	me__ := (*PrintDialogCallback)(self)
	proxy__ := lookupPrintDialogCallbackProxy(me__.Base())
	proxy__.Cont(me__, (*PrintSettings)(settings))
}

// Cancel (cancel)
// Cancel the printing.
func (d *PrintDialogCallback) Cancel() {
	lookupPrintDialogCallbackProxy(d.Base()).Cancel(d)
}

//export gocef_print_dialog_callback_cancel
func gocef_print_dialog_callback_cancel(self *C.cef_print_dialog_callback_t) {
	me__ := (*PrintDialogCallback)(self)
	proxy__ := lookupPrintDialogCallbackProxy(me__.Base())
	proxy__.Cancel(me__)
}
