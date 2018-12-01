// Code generated - DO NOT EDIT.

package cef

import (
	// #include "ResourceHandler_gen.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

// ResourceHandlerProxy defines methods required for using ResourceHandler.
type ResourceHandlerProxy interface {
	ProcessRequest(self *ResourceHandler, request *Request, callback *Callback) int32
	GetResponseHeaders(self *ResourceHandler, response *Response, response_length *int64, redirectUrl string)
	ReadResponse(self *ResourceHandler, data_out unsafe.Pointer, bytes_to_read int32, bytes_read *int32, callback *Callback) int32
	CanGetCookie(self *ResourceHandler, cookie *Cookie) int32
	CanSetCookie(self *ResourceHandler, cookie *Cookie) int32
	Cancel(self *ResourceHandler)
}

// ResourceHandler (cef_resource_handler_t from include/capi/cef_resource_handler_capi.h)
// Structure used to implement a custom request handler structure. The functions
// of this structure will always be called on the IO thread.
type ResourceHandler C.cef_resource_handler_t

// NewResourceHandler creates a new ResourceHandler with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewResourceHandler(proxy ResourceHandlerProxy) *ResourceHandler {
	result := (*ResourceHandler)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_resource_handler_t, proxy)))
	if proxy != nil {
		C.gocef_set_resource_handler_proxy(result.toNative())
	}
	return result
}

func (d *ResourceHandler) toNative() *C.cef_resource_handler_t {
	return (*C.cef_resource_handler_t)(d)
}

func lookupResourceHandlerProxy(obj *BaseRefCounted) ResourceHandlerProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(ResourceHandlerProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type ResourceHandlerProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *ResourceHandler) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// ProcessRequest (process_request)
// Begin processing the request. To handle the request return true (1) and
// call cef_callback_t::cont() once the response header information is
// available (cef_callback_t::cont() can also be called from inside this
// function if header information is available immediately). To cancel the
// request return false (0).
func (d *ResourceHandler) ProcessRequest(request *Request, callback *Callback) int32 {
	return lookupResourceHandlerProxy(d.Base()).ProcessRequest(d, request, callback)
}

//export gocef_resource_handler_process_request
func gocef_resource_handler_process_request(self *C.cef_resource_handler_t, request *C.cef_request_t, callback *C.cef_callback_t) C.int {
	me__ := (*ResourceHandler)(self)
	proxy__ := lookupResourceHandlerProxy(me__.Base())
	return C.int(proxy__.ProcessRequest(me__, (*Request)(request), (*Callback)(callback)))
}

// GetResponseHeaders (get_response_headers)
// Retrieve response header information. If the response length is not known
// set |response_length| to -1 and read_response() will be called until it
// returns false (0). If the response length is known set |response_length| to
// a positive value and read_response() will be called until it returns false
// (0) or the specified number of bytes have been read. Use the |response|
// object to set the mime type, http status code and other optional header
// values. To redirect the request to a new URL set |redirectUrl| to the new
// URL. If an error occured while setting up the request you can call
// set_error() on |response| to indicate the error condition.
func (d *ResourceHandler) GetResponseHeaders(response *Response, response_length *int64, redirectUrl string) {
	lookupResourceHandlerProxy(d.Base()).GetResponseHeaders(d, response, response_length, redirectUrl)
}

//export gocef_resource_handler_get_response_headers
func gocef_resource_handler_get_response_headers(self *C.cef_resource_handler_t, response *C.cef_response_t, response_length *C.int64, redirectUrl *C.cef_string_t) {
	me__ := (*ResourceHandler)(self)
	proxy__ := lookupResourceHandlerProxy(me__.Base())
	redirectUrl_ := cefstrToString(redirectUrl)
	proxy__.GetResponseHeaders(me__, (*Response)(response), (*int64)(response_length), redirectUrl_)
}

// ReadResponse (read_response)
// Read response data. If data is available immediately copy up to
// |bytes_to_read| bytes into |data_out|, set |bytes_read| to the number of
// bytes copied, and return true (1). To read the data at a later time set
// |bytes_read| to 0, return true (1) and call cef_callback_t::cont() when the
// data is available. To indicate response completion return false (0).
func (d *ResourceHandler) ReadResponse(data_out unsafe.Pointer, bytes_to_read int32, bytes_read *int32, callback *Callback) int32 {
	return lookupResourceHandlerProxy(d.Base()).ReadResponse(d, data_out, bytes_to_read, bytes_read, callback)
}

//export gocef_resource_handler_read_response
func gocef_resource_handler_read_response(self *C.cef_resource_handler_t, data_out unsafe.Pointer, bytes_to_read C.int, bytes_read *C.int, callback *C.cef_callback_t) C.int {
	me__ := (*ResourceHandler)(self)
	proxy__ := lookupResourceHandlerProxy(me__.Base())
	return C.int(proxy__.ReadResponse(me__, data_out, int32(bytes_to_read), (*int32)(bytes_read), (*Callback)(callback)))
}

// CanGetCookie (can_get_cookie)
// Return true (1) if the specified cookie can be sent with the request or
// false (0) otherwise. If false (0) is returned for any cookie then no
// cookies will be sent with the request.
func (d *ResourceHandler) CanGetCookie(cookie *Cookie) int32 {
	return lookupResourceHandlerProxy(d.Base()).CanGetCookie(d, cookie)
}

//export gocef_resource_handler_can_get_cookie
func gocef_resource_handler_can_get_cookie(self *C.cef_resource_handler_t, cookie *C.cef_cookie_t) C.int {
	me__ := (*ResourceHandler)(self)
	proxy__ := lookupResourceHandlerProxy(me__.Base())
	cookie_ := cookie.toGo()
	return C.int(proxy__.CanGetCookie(me__, cookie_))
}

// CanSetCookie (can_set_cookie)
// Return true (1) if the specified cookie returned with the response can be
// set or false (0) otherwise.
func (d *ResourceHandler) CanSetCookie(cookie *Cookie) int32 {
	return lookupResourceHandlerProxy(d.Base()).CanSetCookie(d, cookie)
}

//export gocef_resource_handler_can_set_cookie
func gocef_resource_handler_can_set_cookie(self *C.cef_resource_handler_t, cookie *C.cef_cookie_t) C.int {
	me__ := (*ResourceHandler)(self)
	proxy__ := lookupResourceHandlerProxy(me__.Base())
	cookie_ := cookie.toGo()
	return C.int(proxy__.CanSetCookie(me__, cookie_))
}

// Cancel (cancel)
// Request processing has been canceled.
func (d *ResourceHandler) Cancel() {
	lookupResourceHandlerProxy(d.Base()).Cancel(d)
}

//export gocef_resource_handler_cancel
func gocef_resource_handler_cancel(self *C.cef_resource_handler_t) {
	me__ := (*ResourceHandler)(self)
	proxy__ := lookupResourceHandlerProxy(me__.Base())
	proxy__.Cancel(me__)
}
