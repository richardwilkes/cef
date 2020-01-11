// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

import (
	// #include "ResourceHandler_gen.h"
	"C"
)

// ResourceHandlerProxy defines methods required for using ResourceHandler.
type ResourceHandlerProxy interface {
	Open(self *ResourceHandler, request *Request, handleRequest *int32, callback *Callback) int32
	ProcessRequest(self *ResourceHandler, request *Request, callback *Callback) int32
	GetResponseHeaders(self *ResourceHandler, response *Response, responseLength *int64, redirectURL *string)
	Skip(self *ResourceHandler, bytesToSkip int64, bytesSkipped *int64, callback *ResourceSkipCallback) int32
	Read(self *ResourceHandler, dataOut unsafe.Pointer, bytesToRead int32, bytesRead *int32, callback *ResourceReadCallback) int32
	ReadResponse(self *ResourceHandler, dataOut unsafe.Pointer, bytesToRead int32, bytesRead *int32, callback *Callback) int32
	Cancel(self *ResourceHandler)
}

// ResourceHandler (cef_resource_handler_t from include/capi/cef_resource_handler_capi.h)
// Structure used to implement a custom request handler structure. The functions
// of this structure will be called on the IO thread unless otherwise indicated.
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

// Open (open)
// Open the response stream. To handle the request immediately set
// |handle_request| to true (1) and return true (1). To decide at a later time
// set |handle_request| to false (0), return true (1), and execute |callback|
// to continue or cancel the request. To cancel the request immediately set
// |handle_request| to true (1) and return false (0). This function will be
// called in sequence but not from a dedicated thread. For backwards
// compatibility set |handle_request| to false (0) and return false (0) and
// the ProcessRequest function will be called.
func (d *ResourceHandler) Open(request *Request, handleRequest *int32, callback *Callback) int32 {
	return lookupResourceHandlerProxy(d.Base()).Open(d, request, handleRequest, callback)
}

//nolint:gocritic
//export gocef_resource_handler_open
func gocef_resource_handler_open(self *C.cef_resource_handler_t, request *C.cef_request_t, handleRequest *C.int, callback *C.cef_callback_t) C.int {
	me__ := (*ResourceHandler)(self)
	proxy__ := lookupResourceHandlerProxy(me__.Base())
	return C.int(proxy__.Open(me__, (*Request)(request), (*int32)(handleRequest), (*Callback)(callback)))
}

// ProcessRequest (process_request)
// Begin processing the request. To handle the request return true (1) and
// call cef_callback_t::cont() once the response header information is
// available (cef_callback_t::cont() can also be called from inside this
// function if header information is available immediately). To cancel the
// request return false (0).
//
// WARNING: This function is deprecated. Use Open instead.
func (d *ResourceHandler) ProcessRequest(request *Request, callback *Callback) int32 {
	return lookupResourceHandlerProxy(d.Base()).ProcessRequest(d, request, callback)
}

//nolint:gocritic
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
// URL. |redirectUrl| can be either a relative or fully qualified URL. It is
// also possible to set |response| to a redirect http status code and pass the
// new URL via a Location header. Likewise with |redirectUrl| it is valid to
// set a relative or fully qualified URL as the Location header value. If an
// error occured while setting up the request you can call set_error() on
// |response| to indicate the error condition.
func (d *ResourceHandler) GetResponseHeaders(response *Response, responseLength *int64, redirectURL *string) {
	lookupResourceHandlerProxy(d.Base()).GetResponseHeaders(d, response, responseLength, redirectURL)
}

//nolint:gocritic
//export gocef_resource_handler_get_response_headers
func gocef_resource_handler_get_response_headers(self *C.cef_resource_handler_t, response *C.cef_response_t, responseLength *C.int64, redirectURL *C.cef_string_t) {
	me__ := (*ResourceHandler)(self)
	proxy__ := lookupResourceHandlerProxy(me__.Base())
	redirectURL_ := cefstrToString(redirectURL)
	proxy__.GetResponseHeaders(me__, (*Response)(response), (*int64)(responseLength), &redirectURL_)
}

// Skip (skip)
// Skip response data when requested by a Range header. Skip over and discard
// |bytes_to_skip| bytes of response data. If data is available immediately
// set |bytes_skipped| to the number of of bytes skipped and return true (1).
// To read the data at a later time set |bytes_skipped| to 0, return true (1)
// and execute |callback| when the data is available. To indicate failure set
// |bytes_skipped| to < 0 (e.g. -2 for ERR_FAILED) and return false (0). This
// function will be called in sequence but not from a dedicated thread.
func (d *ResourceHandler) Skip(bytesToSkip int64, bytesSkipped *int64, callback *ResourceSkipCallback) int32 {
	return lookupResourceHandlerProxy(d.Base()).Skip(d, bytesToSkip, bytesSkipped, callback)
}

//nolint:gocritic
//export gocef_resource_handler_skip
func gocef_resource_handler_skip(self *C.cef_resource_handler_t, bytesToSkip C.int64, bytesSkipped *C.int64, callback *C.cef_resource_skip_callback_t) C.int {
	me__ := (*ResourceHandler)(self)
	proxy__ := lookupResourceHandlerProxy(me__.Base())
	return C.int(proxy__.Skip(me__, int64(bytesToSkip), (*int64)(bytesSkipped), (*ResourceSkipCallback)(callback)))
}

// Read (read)
// Read response data. If data is available immediately copy up to
// |bytes_to_read| bytes into |data_out|, set |bytes_read| to the number of
// bytes copied, and return true (1). To read the data at a later time keep a
// pointer to |data_out|, set |bytes_read| to 0, return true (1) and execute
// |callback| when the data is available (|data_out| will remain valid until
// the callback is executed). To indicate response completion set |bytes_read|
// to 0 and return false (0). To indicate failure set |bytes_read| to < 0
// (e.g. -2 for ERR_FAILED) and return false (0). This function will be called
// in sequence but not from a dedicated thread. For backwards compatibility
// set |bytes_read| to -1 and return false (0) and the ReadResponse function
// will be called.
func (d *ResourceHandler) Read(dataOut unsafe.Pointer, bytesToRead int32, bytesRead *int32, callback *ResourceReadCallback) int32 {
	return lookupResourceHandlerProxy(d.Base()).Read(d, dataOut, bytesToRead, bytesRead, callback)
}

//nolint:gocritic
//export gocef_resource_handler_read
func gocef_resource_handler_read(self *C.cef_resource_handler_t, dataOut unsafe.Pointer, bytesToRead C.int, bytesRead *C.int, callback *C.cef_resource_read_callback_t) C.int {
	me__ := (*ResourceHandler)(self)
	proxy__ := lookupResourceHandlerProxy(me__.Base())
	return C.int(proxy__.Read(me__, dataOut, int32(bytesToRead), (*int32)(bytesRead), (*ResourceReadCallback)(callback)))
}

// ReadResponse (read_response)
// Read response data. If data is available immediately copy up to
// |bytes_to_read| bytes into |data_out|, set |bytes_read| to the number of
// bytes copied, and return true (1). To read the data at a later time set
// |bytes_read| to 0, return true (1) and call cef_callback_t::cont() when the
// data is available. To indicate response completion return false (0).
//
// WARNING: This function is deprecated. Use Skip and Read instead.
func (d *ResourceHandler) ReadResponse(dataOut unsafe.Pointer, bytesToRead int32, bytesRead *int32, callback *Callback) int32 {
	return lookupResourceHandlerProxy(d.Base()).ReadResponse(d, dataOut, bytesToRead, bytesRead, callback)
}

//nolint:gocritic
//export gocef_resource_handler_read_response
func gocef_resource_handler_read_response(self *C.cef_resource_handler_t, dataOut unsafe.Pointer, bytesToRead C.int, bytesRead *C.int, callback *C.cef_callback_t) C.int {
	me__ := (*ResourceHandler)(self)
	proxy__ := lookupResourceHandlerProxy(me__.Base())
	return C.int(proxy__.ReadResponse(me__, dataOut, int32(bytesToRead), (*int32)(bytesRead), (*Callback)(callback)))
}

// Cancel (cancel)
// Request processing has been canceled.
func (d *ResourceHandler) Cancel() {
	lookupResourceHandlerProxy(d.Base()).Cancel(d)
}

//nolint:gocritic
//export gocef_resource_handler_cancel
func gocef_resource_handler_cancel(self *C.cef_resource_handler_t) {
	me__ := (*ResourceHandler)(self)
	proxy__ := lookupResourceHandlerProxy(me__.Base())
	proxy__.Cancel(me__)
}
