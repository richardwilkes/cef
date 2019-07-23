// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

import (
	// #include "ResourceRequestHandler_gen.h"
	"C"
)

// ResourceRequestHandlerProxy defines methods required for using ResourceRequestHandler.
type ResourceRequestHandlerProxy interface {
	GetCookieAccessFilter(self *ResourceRequestHandler, browser *Browser, frame *Frame, request *Request) *CookieAccessFilter
	OnBeforeResourceLoad(self *ResourceRequestHandler, browser *Browser, frame *Frame, request *Request, callback *RequestCallback) ReturnValue
	GetResourceHandler(self *ResourceRequestHandler, browser *Browser, frame *Frame, request *Request) *ResourceHandler
	OnResourceRedirect(self *ResourceRequestHandler, browser *Browser, frame *Frame, request *Request, response *Response, new_url *string)
	OnResourceResponse(self *ResourceRequestHandler, browser *Browser, frame *Frame, request *Request, response *Response) int32
	GetResourceResponseFilter(self *ResourceRequestHandler, browser *Browser, frame *Frame, request *Request, response *Response) *ResponseFilter
	OnResourceLoadComplete(self *ResourceRequestHandler, browser *Browser, frame *Frame, request *Request, response *Response, status UrlrequestStatus, received_content_length int64)
	OnProtocolExecution(self *ResourceRequestHandler, browser *Browser, frame *Frame, request *Request, allow_os_execution *int32)
}

// ResourceRequestHandler (cef_resource_request_handler_t from include/capi/cef_resource_request_handler_capi.h)
// Implement this structure to handle events related to browser requests. The
// functions of this structure will be called on the IO thread unless otherwise
// indicated.
type ResourceRequestHandler C.cef_resource_request_handler_t

// NewResourceRequestHandler creates a new ResourceRequestHandler with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewResourceRequestHandler(proxy ResourceRequestHandlerProxy) *ResourceRequestHandler {
	result := (*ResourceRequestHandler)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_resource_request_handler_t, proxy)))
	if proxy != nil {
		C.gocef_set_resource_request_handler_proxy(result.toNative())
	}
	return result
}

func (d *ResourceRequestHandler) toNative() *C.cef_resource_request_handler_t {
	return (*C.cef_resource_request_handler_t)(d)
}

func lookupResourceRequestHandlerProxy(obj *BaseRefCounted) ResourceRequestHandlerProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(ResourceRequestHandlerProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type ResourceRequestHandlerProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *ResourceRequestHandler) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// GetCookieAccessFilter (get_cookie_access_filter)
// Called on the IO thread before a resource request is loaded. The |browser|
// and |frame| values represent the source of the request, and may be NULL for
// requests originating from service workers or cef_urlrequest_t. To
// optionally filter cookies for the request return a
// cef_cookie_access_filter_t object. The |request| object cannot not be
// modified in this callback.
func (d *ResourceRequestHandler) GetCookieAccessFilter(browser *Browser, frame *Frame, request *Request) *CookieAccessFilter {
	return lookupResourceRequestHandlerProxy(d.Base()).GetCookieAccessFilter(d, browser, frame, request)
}

//nolint:gocritic
//export gocef_resource_request_handler_get_cookie_access_filter
func gocef_resource_request_handler_get_cookie_access_filter(self *C.cef_resource_request_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, request *C.cef_request_t) *C.cef_cookie_access_filter_t {
	me__ := (*ResourceRequestHandler)(self)
	proxy__ := lookupResourceRequestHandlerProxy(me__.Base())
	return (proxy__.GetCookieAccessFilter(me__, (*Browser)(browser), (*Frame)(frame), (*Request)(request))).toNative()
}

// OnBeforeResourceLoad (on_before_resource_load)
// Called on the IO thread before a resource request is loaded. The |browser|
// and |frame| values represent the source of the request, and may be NULL for
// requests originating from service workers or cef_urlrequest_t. To redirect
// or change the resource load optionally modify |request|. Modification of
// the request URL will be treated as a redirect. Return RV_CONTINUE to
// continue the request immediately. Return RV_CONTINUE_ASYNC and call
// cef_request_tCallback:: cont() at a later time to continue or cancel the
// request asynchronously. Return RV_CANCEL to cancel the request immediately.
//
func (d *ResourceRequestHandler) OnBeforeResourceLoad(browser *Browser, frame *Frame, request *Request, callback *RequestCallback) ReturnValue {
	return lookupResourceRequestHandlerProxy(d.Base()).OnBeforeResourceLoad(d, browser, frame, request, callback)
}

//nolint:gocritic
//export gocef_resource_request_handler_on_before_resource_load
func gocef_resource_request_handler_on_before_resource_load(self *C.cef_resource_request_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, request *C.cef_request_t, callback *C.cef_request_callback_t) C.cef_return_value_t {
	me__ := (*ResourceRequestHandler)(self)
	proxy__ := lookupResourceRequestHandlerProxy(me__.Base())
	return C.cef_return_value_t(proxy__.OnBeforeResourceLoad(me__, (*Browser)(browser), (*Frame)(frame), (*Request)(request), (*RequestCallback)(callback)))
}

// GetResourceHandler (get_resource_handler)
// Called on the IO thread before a resource is loaded. The |browser| and
// |frame| values represent the source of the request, and may be NULL for
// requests originating from service workers or cef_urlrequest_t. To allow the
// resource to load using the default network loader return NULL. To specify a
// handler for the resource return a cef_resource_handler_t object. The
// |request| object cannot not be modified in this callback.
func (d *ResourceRequestHandler) GetResourceHandler(browser *Browser, frame *Frame, request *Request) *ResourceHandler {
	return lookupResourceRequestHandlerProxy(d.Base()).GetResourceHandler(d, browser, frame, request)
}

//nolint:gocritic
//export gocef_resource_request_handler_get_resource_handler
func gocef_resource_request_handler_get_resource_handler(self *C.cef_resource_request_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, request *C.cef_request_t) *C.cef_resource_handler_t {
	me__ := (*ResourceRequestHandler)(self)
	proxy__ := lookupResourceRequestHandlerProxy(me__.Base())
	return (proxy__.GetResourceHandler(me__, (*Browser)(browser), (*Frame)(frame), (*Request)(request))).toNative()
}

// OnResourceRedirect (on_resource_redirect)
// Called on the IO thread when a resource load is redirected. The |browser|
// and |frame| values represent the source of the request, and may be NULL for
// requests originating from service workers or cef_urlrequest_t. The
// |request| parameter will contain the old URL and other request-related
// information. The |response| parameter will contain the response that
// resulted in the redirect. The |new_url| parameter will contain the new URL
// and can be changed if desired. The |request| and |response| objects cannot
// be modified in this callback.
func (d *ResourceRequestHandler) OnResourceRedirect(browser *Browser, frame *Frame, request *Request, response *Response, new_url *string) {
	lookupResourceRequestHandlerProxy(d.Base()).OnResourceRedirect(d, browser, frame, request, response, new_url)
}

//nolint:gocritic
//export gocef_resource_request_handler_on_resource_redirect
func gocef_resource_request_handler_on_resource_redirect(self *C.cef_resource_request_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, request *C.cef_request_t, response *C.cef_response_t, new_url *C.cef_string_t) {
	me__ := (*ResourceRequestHandler)(self)
	proxy__ := lookupResourceRequestHandlerProxy(me__.Base())
	new_url_ := cefstrToString(new_url)
	proxy__.OnResourceRedirect(me__, (*Browser)(browser), (*Frame)(frame), (*Request)(request), (*Response)(response), &new_url_)
}

// OnResourceResponse (on_resource_response)
// Called on the IO thread when a resource response is received. The |browser|
// and |frame| values represent the source of the request, and may be NULL for
// requests originating from service workers or cef_urlrequest_t. To allow the
// resource load to proceed without modification return false (0). To redirect
// or retry the resource load optionally modify |request| and return true (1).
// Modification of the request URL will be treated as a redirect. Requests
// handled using the default network loader cannot be redirected in this
// callback. The |response| object cannot be modified in this callback.
//
// WARNING: Redirecting using this function is deprecated. Use
// OnBeforeResourceLoad or GetResourceHandler to perform redirects.
func (d *ResourceRequestHandler) OnResourceResponse(browser *Browser, frame *Frame, request *Request, response *Response) int32 {
	return lookupResourceRequestHandlerProxy(d.Base()).OnResourceResponse(d, browser, frame, request, response)
}

//nolint:gocritic
//export gocef_resource_request_handler_on_resource_response
func gocef_resource_request_handler_on_resource_response(self *C.cef_resource_request_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, request *C.cef_request_t, response *C.cef_response_t) C.int {
	me__ := (*ResourceRequestHandler)(self)
	proxy__ := lookupResourceRequestHandlerProxy(me__.Base())
	return C.int(proxy__.OnResourceResponse(me__, (*Browser)(browser), (*Frame)(frame), (*Request)(request), (*Response)(response)))
}

// GetResourceResponseFilter (get_resource_response_filter)
// Called on the IO thread to optionally filter resource response content. The
// |browser| and |frame| values represent the source of the request, and may
// be NULL for requests originating from service workers or cef_urlrequest_t.
// |request| and |response| represent the request and response respectively
// and cannot be modified in this callback.
func (d *ResourceRequestHandler) GetResourceResponseFilter(browser *Browser, frame *Frame, request *Request, response *Response) *ResponseFilter {
	return lookupResourceRequestHandlerProxy(d.Base()).GetResourceResponseFilter(d, browser, frame, request, response)
}

//nolint:gocritic
//export gocef_resource_request_handler_get_resource_response_filter
func gocef_resource_request_handler_get_resource_response_filter(self *C.cef_resource_request_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, request *C.cef_request_t, response *C.cef_response_t) *C.cef_response_filter_t {
	me__ := (*ResourceRequestHandler)(self)
	proxy__ := lookupResourceRequestHandlerProxy(me__.Base())
	return (proxy__.GetResourceResponseFilter(me__, (*Browser)(browser), (*Frame)(frame), (*Request)(request), (*Response)(response))).toNative()
}

// OnResourceLoadComplete (on_resource_load_complete)
// Called on the IO thread when a resource load has completed. The |browser|
// and |frame| values represent the source of the request, and may be NULL for
// requests originating from service workers or cef_urlrequest_t. |request|
// and |response| represent the request and response respectively and cannot
// be modified in this callback. |status| indicates the load completion
// status. |received_content_length| is the number of response bytes actually
// read. This function will be called for all requests, including requests
// that are aborted due to CEF shutdown or destruction of the associated
// browser. In cases where the associated browser is destroyed this callback
// may arrive after the cef_life_span_handler_t::OnBeforeClose callback for
// that browser. The cef_frame_t::IsValid function can be used to test for
// this situation, and care should be taken not to call |browser| or |frame|
// functions that modify state (like LoadURL, SendProcessMessage, etc.) if the
// frame is invalid.
func (d *ResourceRequestHandler) OnResourceLoadComplete(browser *Browser, frame *Frame, request *Request, response *Response, status UrlrequestStatus, received_content_length int64) {
	lookupResourceRequestHandlerProxy(d.Base()).OnResourceLoadComplete(d, browser, frame, request, response, status, received_content_length)
}

//nolint:gocritic
//export gocef_resource_request_handler_on_resource_load_complete
func gocef_resource_request_handler_on_resource_load_complete(self *C.cef_resource_request_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, request *C.cef_request_t, response *C.cef_response_t, status C.cef_urlrequest_status_t, received_content_length C.int64) {
	me__ := (*ResourceRequestHandler)(self)
	proxy__ := lookupResourceRequestHandlerProxy(me__.Base())
	proxy__.OnResourceLoadComplete(me__, (*Browser)(browser), (*Frame)(frame), (*Request)(request), (*Response)(response), UrlrequestStatus(status), int64(received_content_length))
}

// OnProtocolExecution (on_protocol_execution)
// Called on the IO thread to handle requests for URLs with an unknown
// protocol component. The |browser| and |frame| values represent the source
// of the request, and may be NULL for requests originating from service
// workers or cef_urlrequest_t. |request| cannot be modified in this callback.
// Set |allow_os_execution| to true (1) to attempt execution via the
// registered OS protocol handler, if any. SECURITY WARNING: YOU SHOULD USE
// THIS METHOD TO ENFORCE RESTRICTIONS BASED ON SCHEME, HOST OR OTHER URL
// ANALYSIS BEFORE ALLOWING OS EXECUTION.
func (d *ResourceRequestHandler) OnProtocolExecution(browser *Browser, frame *Frame, request *Request, allow_os_execution *int32) {
	lookupResourceRequestHandlerProxy(d.Base()).OnProtocolExecution(d, browser, frame, request, allow_os_execution)
}

//nolint:gocritic
//export gocef_resource_request_handler_on_protocol_execution
func gocef_resource_request_handler_on_protocol_execution(self *C.cef_resource_request_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, request *C.cef_request_t, allow_os_execution *C.int) {
	me__ := (*ResourceRequestHandler)(self)
	proxy__ := lookupResourceRequestHandlerProxy(me__.Base())
	proxy__.OnProtocolExecution(me__, (*Browser)(browser), (*Frame)(frame), (*Request)(request), (*int32)(allow_os_execution))
}
