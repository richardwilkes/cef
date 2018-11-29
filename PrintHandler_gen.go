// Code generated - DO NOT EDIT.

package cef

import (
	// #include "PrintHandler_gen.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

// PrintHandlerProxy defines methods required for using PrintHandler.
type PrintHandlerProxy interface {
	OnPrintStart(self *PrintHandler, browser *Browser)
	OnPrintSettings(self *PrintHandler, browser *Browser, settings *PrintSettings, get_defaults int32)
	OnPrintDialog(self *PrintHandler, browser *Browser, has_selection int32, callback *PrintDialogCallback) int32
	OnPrintJob(self *PrintHandler, browser *Browser, document_name string, pdf_file_path string, callback *PrintJobCallback) int32
	OnPrintReset(self *PrintHandler, browser *Browser)
	GetPdfPaperSize(self *PrintHandler, device_units_per_inch int32) Size
}

// PrintHandler (cef_print_handler_t from include/capi/cef_print_handler_capi.h)
// Implement this structure to handle printing on Linux. Each browser will have
// only one print job in progress at a time. The functions of this structure
// will be called on the browser process UI thread.
type PrintHandler C.cef_print_handler_t

// NewPrintHandler creates a new PrintHandler with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewPrintHandler(proxy PrintHandlerProxy) *PrintHandler {
	result := (*PrintHandler)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_print_handler_t, proxy)))
	if proxy != nil {
		C.gocef_set_print_handler_proxy(result.toNative())
	}
	return result
}

func (d *PrintHandler) toNative() *C.cef_print_handler_t {
	return (*C.cef_print_handler_t)(d)
}

func lookupPrintHandlerProxy(obj *BaseRefCounted) PrintHandlerProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(PrintHandlerProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type PrintHandlerProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *PrintHandler) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnPrintStart (on_print_start)
// Called when printing has started for the specified |browser|. This function
// will be called before the other OnPrint*() functions and irrespective of
// how printing was initiated (e.g. cef_browser_host_t::print(), JavaScript
// window.print() or PDF extension print button).
func (d *PrintHandler) OnPrintStart(browser *Browser) {
	lookupPrintHandlerProxy(d.Base()).OnPrintStart(d, browser)
}

//export gocef_print_handler_on_print_start
func gocef_print_handler_on_print_start(self *C.cef_print_handler_t, browser *C.cef_browser_t) {
	me__ := (*PrintHandler)(self)
	proxy__ := lookupPrintHandlerProxy(me__.Base())
	proxy__.OnPrintStart(me__, (*Browser)(browser))
}

// OnPrintSettings (on_print_settings)
// Synchronize |settings| with client state. If |get_defaults| is true (1)
// then populate |settings| with the default print settings. Do not keep a
// reference to |settings| outside of this callback.
func (d *PrintHandler) OnPrintSettings(browser *Browser, settings *PrintSettings, get_defaults int32) {
	lookupPrintHandlerProxy(d.Base()).OnPrintSettings(d, browser, settings, get_defaults)
}

//export gocef_print_handler_on_print_settings
func gocef_print_handler_on_print_settings(self *C.cef_print_handler_t, browser *C.cef_browser_t, settings *C.cef_print_settings_t, get_defaults C.int) {
	me__ := (*PrintHandler)(self)
	proxy__ := lookupPrintHandlerProxy(me__.Base())
	proxy__.OnPrintSettings(me__, (*Browser)(browser), (*PrintSettings)(settings), int32(get_defaults))
}

// OnPrintDialog (on_print_dialog)
// Show the print dialog. Execute |callback| once the dialog is dismissed.
// Return true (1) if the dialog will be displayed or false (0) to cancel the
// printing immediately.
func (d *PrintHandler) OnPrintDialog(browser *Browser, has_selection int32, callback *PrintDialogCallback) int32 {
	return lookupPrintHandlerProxy(d.Base()).OnPrintDialog(d, browser, has_selection, callback)
}

//export gocef_print_handler_on_print_dialog
func gocef_print_handler_on_print_dialog(self *C.cef_print_handler_t, browser *C.cef_browser_t, has_selection C.int, callback *C.cef_print_dialog_callback_t) C.int {
	me__ := (*PrintHandler)(self)
	proxy__ := lookupPrintHandlerProxy(me__.Base())
	return C.int(proxy__.OnPrintDialog(me__, (*Browser)(browser), int32(has_selection), (*PrintDialogCallback)(callback)))
}

// OnPrintJob (on_print_job)
// Send the print job to the printer. Execute |callback| once the job is
// completed. Return true (1) if the job will proceed or false (0) to cancel
// the job immediately.
func (d *PrintHandler) OnPrintJob(browser *Browser, document_name, pdf_file_path string, callback *PrintJobCallback) int32 {
	return lookupPrintHandlerProxy(d.Base()).OnPrintJob(d, browser, document_name, pdf_file_path, callback)
}

//export gocef_print_handler_on_print_job
func gocef_print_handler_on_print_job(self *C.cef_print_handler_t, browser *C.cef_browser_t, document_name *C.cef_string_t, pdf_file_path *C.cef_string_t, callback *C.cef_print_job_callback_t) C.int {
	me__ := (*PrintHandler)(self)
	proxy__ := lookupPrintHandlerProxy(me__.Base())
	return C.int(proxy__.OnPrintJob(me__, (*Browser)(browser), cefstrToString(document_name), cefstrToString(pdf_file_path), (*PrintJobCallback)(callback)))
}

// OnPrintReset (on_print_reset)
// Reset client state related to printing.
func (d *PrintHandler) OnPrintReset(browser *Browser) {
	lookupPrintHandlerProxy(d.Base()).OnPrintReset(d, browser)
}

//export gocef_print_handler_on_print_reset
func gocef_print_handler_on_print_reset(self *C.cef_print_handler_t, browser *C.cef_browser_t) {
	me__ := (*PrintHandler)(self)
	proxy__ := lookupPrintHandlerProxy(me__.Base())
	proxy__.OnPrintReset(me__, (*Browser)(browser))
}

// GetPdfPaperSize (get_pdf_paper_size)
// Return the PDF paper size in device units. Used in combination with
// cef_browser_host_t::print_to_pdf().
func (d *PrintHandler) GetPdfPaperSize(device_units_per_inch int32) Size {
	return lookupPrintHandlerProxy(d.Base()).GetPdfPaperSize(d, device_units_per_inch)
}

//export gocef_print_handler_get_pdf_paper_size
func gocef_print_handler_get_pdf_paper_size(self *C.cef_print_handler_t, device_units_per_inch C.int) C.cef_size_t {
	me__ := (*PrintHandler)(self)
	proxy__ := lookupPrintHandlerProxy(me__.Base())
	call__ := proxy__.GetPdfPaperSize(me__, int32(device_units_per_inch))
	var result__ C.cef_size_t
	call__.toNative(&result__)
	return result__
}
