// Code generated - DO NOT EDIT.

package cef

import (
	// #include "RequestHandler_gen.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

// RequestHandlerProxy defines methods required for using RequestHandler.
type RequestHandlerProxy interface {
	OnBeforeBrowse(self *RequestHandler, browser *Browser, frame *Frame, request *Request, user_gesture, is_redirect int32) int32
	OnOpenUrlfromTab(self *RequestHandler, browser *Browser, frame *Frame, target_url string, target_disposition WindowOpenDisposition, user_gesture int32) int32
	OnBeforeResourceLoad(self *RequestHandler, browser *Browser, frame *Frame, request *Request, callback *RequestCallback) ReturnValue
	GetResourceHandler(self *RequestHandler, browser *Browser, frame *Frame, request *Request) *ResourceHandler
	OnResourceRedirect(self *RequestHandler, browser *Browser, frame *Frame, request *Request, response *Response, new_url string)
	OnResourceResponse(self *RequestHandler, browser *Browser, frame *Frame, request *Request, response *Response) int32
	GetResourceResponseFilter(self *RequestHandler, browser *Browser, frame *Frame, request *Request, response *Response) *ResponseFilter
	OnResourceLoadComplete(self *RequestHandler, browser *Browser, frame *Frame, request *Request, response *Response, status UrlrequestStatus, received_content_length int64)
	GetAuthCredentials(self *RequestHandler, browser *Browser, frame *Frame, isProxy int32, host string, port int32, realm, scheme string, callback *AuthCallback) int32
	CanGetCookies(self *RequestHandler, browser *Browser, frame *Frame, request *Request) int32
	CanSetCookie(self *RequestHandler, browser *Browser, frame *Frame, request *Request, cookie *Cookie) int32
	OnQuotaRequest(self *RequestHandler, browser *Browser, origin_url string, new_size int64, callback *RequestCallback) int32
	OnProtocolExecution(self *RequestHandler, browser *Browser, url string, allow_os_execution *int32)
	OnCertificateError(self *RequestHandler, browser *Browser, cert_error Errorcode, request_url string, ssl_info *Sslinfo, callback *RequestCallback) int32
	OnSelectClientCertificate(self *RequestHandler, browser *Browser, isProxy int32, host string, port int32, certificatesCount uint64, certificates **X509certificate, callback *SelectClientCertificateCallback) int32
	OnPluginCrashed(self *RequestHandler, browser *Browser, plugin_path string)
	OnRenderViewReady(self *RequestHandler, browser *Browser)
	OnRenderProcessTerminated(self *RequestHandler, browser *Browser, status TerminationStatus)
}

// RequestHandler (cef_request_handler_t from include/capi/cef_request_handler_capi.h)
// Implement this structure to handle events related to browser requests. The
// functions of this structure will be called on the thread indicated.
type RequestHandler C.cef_request_handler_t

// NewRequestHandler creates a new RequestHandler with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewRequestHandler(proxy RequestHandlerProxy) *RequestHandler {
	result := (*RequestHandler)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_request_handler_t, proxy)))
	if proxy != nil {
		C.gocef_set_request_handler_proxy(result.toNative())
	}
	return result
}

func (d *RequestHandler) toNative() *C.cef_request_handler_t {
	return (*C.cef_request_handler_t)(d)
}

func lookupRequestHandlerProxy(obj *BaseRefCounted) RequestHandlerProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(RequestHandlerProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type RequestHandlerProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *RequestHandler) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnBeforeBrowse (on_before_browse)
// Called on the UI thread before browser navigation. Return true (1) to
// cancel the navigation or false (0) to allow the navigation to proceed. The
// |request| object cannot be modified in this callback.
// cef_load_handler_t::OnLoadingStateChange will be called twice in all cases.
// If the navigation is allowed cef_load_handler_t::OnLoadStart and
// cef_load_handler_t::OnLoadEnd will be called. If the navigation is canceled
// cef_load_handler_t::OnLoadError will be called with an |errorCode| value of
// ERR_ABORTED. The |user_gesture| value will be true (1) if the browser
// navigated via explicit user gesture (e.g. clicking a link) or false (0) if
// it navigated automatically (e.g. via the DomContentLoaded event).
func (d *RequestHandler) OnBeforeBrowse(browser *Browser, frame *Frame, request *Request, user_gesture, is_redirect int32) int32 {
	return lookupRequestHandlerProxy(d.Base()).OnBeforeBrowse(d, browser, frame, request, user_gesture, is_redirect)
}

//export gocef_request_handler_on_before_browse
func gocef_request_handler_on_before_browse(self *C.cef_request_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, request *C.cef_request_t, user_gesture C.int, is_redirect C.int) C.int {
	me__ := (*RequestHandler)(self)
	proxy__ := lookupRequestHandlerProxy(me__.Base())
	return C.int(proxy__.OnBeforeBrowse(me__, (*Browser)(browser), (*Frame)(frame), (*Request)(request), int32(user_gesture), int32(is_redirect)))
}

// OnOpenUrlfromTab (on_open_urlfrom_tab)
// Called on the UI thread before OnBeforeBrowse in certain limited cases
// where navigating a new or different browser might be desirable. This
// includes user-initiated navigation that might open in a special way (e.g.
// links clicked via middle-click or ctrl + left-click) and certain types of
// cross-origin navigation initiated from the renderer process (e.g.
// navigating the top-level frame to/from a file URL). The |browser| and
// |frame| values represent the source of the navigation. The
// |target_disposition| value indicates where the user intended to navigate
// the browser based on standard Chromium behaviors (e.g. current tab, new
// tab, etc). The |user_gesture| value will be true (1) if the browser
// navigated via explicit user gesture (e.g. clicking a link) or false (0) if
// it navigated automatically (e.g. via the DomContentLoaded event). Return
// true (1) to cancel the navigation or false (0) to allow the navigation to
// proceed in the source browser's top-level frame.
func (d *RequestHandler) OnOpenUrlfromTab(browser *Browser, frame *Frame, target_url string, target_disposition WindowOpenDisposition, user_gesture int32) int32 {
	return lookupRequestHandlerProxy(d.Base()).OnOpenUrlfromTab(d, browser, frame, target_url, target_disposition, user_gesture)
}

//export gocef_request_handler_on_open_urlfrom_tab
func gocef_request_handler_on_open_urlfrom_tab(self *C.cef_request_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, target_url *C.cef_string_t, target_disposition C.cef_window_open_disposition_t, user_gesture C.int) C.int {
	me__ := (*RequestHandler)(self)
	proxy__ := lookupRequestHandlerProxy(me__.Base())
	target_url_ := cefstrToString(target_url)
	return C.int(proxy__.OnOpenUrlfromTab(me__, (*Browser)(browser), (*Frame)(frame), target_url_, WindowOpenDisposition(target_disposition), int32(user_gesture)))
}

// OnBeforeResourceLoad (on_before_resource_load)
// Called on the IO thread before a resource request is loaded. The |request|
// object may be modified. Return RV_CONTINUE to continue the request
// immediately. Return RV_CONTINUE_ASYNC and call cef_request_tCallback::
// cont() at a later time to continue or cancel the request asynchronously.
// Return RV_CANCEL to cancel the request immediately.
//
func (d *RequestHandler) OnBeforeResourceLoad(browser *Browser, frame *Frame, request *Request, callback *RequestCallback) ReturnValue {
	return lookupRequestHandlerProxy(d.Base()).OnBeforeResourceLoad(d, browser, frame, request, callback)
}

//export gocef_request_handler_on_before_resource_load
func gocef_request_handler_on_before_resource_load(self *C.cef_request_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, request *C.cef_request_t, callback *C.cef_request_callback_t) C.cef_return_value_t {
	me__ := (*RequestHandler)(self)
	proxy__ := lookupRequestHandlerProxy(me__.Base())
	return C.cef_return_value_t(proxy__.OnBeforeResourceLoad(me__, (*Browser)(browser), (*Frame)(frame), (*Request)(request), (*RequestCallback)(callback)))
}

// GetResourceHandler (get_resource_handler)
// Called on the IO thread before a resource is loaded. To allow the resource
// to load normally return NULL. To specify a handler for the resource return
// a cef_resource_handler_t object. The |request| object should not be
// modified in this callback.
func (d *RequestHandler) GetResourceHandler(browser *Browser, frame *Frame, request *Request) *ResourceHandler {
	return lookupRequestHandlerProxy(d.Base()).GetResourceHandler(d, browser, frame, request)
}

//export gocef_request_handler_get_resource_handler
func gocef_request_handler_get_resource_handler(self *C.cef_request_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, request *C.cef_request_t) *C.cef_resource_handler_t {
	me__ := (*RequestHandler)(self)
	proxy__ := lookupRequestHandlerProxy(me__.Base())
	return (proxy__.GetResourceHandler(me__, (*Browser)(browser), (*Frame)(frame), (*Request)(request))).toNative()
}

// OnResourceRedirect (on_resource_redirect)
// Called on the IO thread when a resource load is redirected. The |request|
// parameter will contain the old URL and other request-related information.
// The |response| parameter will contain the response that resulted in the
// redirect. The |new_url| parameter will contain the new URL and can be
// changed if desired. The |request| object cannot be modified in this
// callback.
func (d *RequestHandler) OnResourceRedirect(browser *Browser, frame *Frame, request *Request, response *Response, new_url string) {
	lookupRequestHandlerProxy(d.Base()).OnResourceRedirect(d, browser, frame, request, response, new_url)
}

//export gocef_request_handler_on_resource_redirect
func gocef_request_handler_on_resource_redirect(self *C.cef_request_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, request *C.cef_request_t, response *C.cef_response_t, new_url *C.cef_string_t) {
	me__ := (*RequestHandler)(self)
	proxy__ := lookupRequestHandlerProxy(me__.Base())
	new_url_ := cefstrToString(new_url)
	proxy__.OnResourceRedirect(me__, (*Browser)(browser), (*Frame)(frame), (*Request)(request), (*Response)(response), new_url_)
}

// OnResourceResponse (on_resource_response)
// Called on the IO thread when a resource response is received. To allow the
// resource to load normally return false (0). To redirect or retry the
// resource modify |request| (url, headers or post body) and return true (1).
// The |response| object cannot be modified in this callback.
func (d *RequestHandler) OnResourceResponse(browser *Browser, frame *Frame, request *Request, response *Response) int32 {
	return lookupRequestHandlerProxy(d.Base()).OnResourceResponse(d, browser, frame, request, response)
}

//export gocef_request_handler_on_resource_response
func gocef_request_handler_on_resource_response(self *C.cef_request_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, request *C.cef_request_t, response *C.cef_response_t) C.int {
	me__ := (*RequestHandler)(self)
	proxy__ := lookupRequestHandlerProxy(me__.Base())
	return C.int(proxy__.OnResourceResponse(me__, (*Browser)(browser), (*Frame)(frame), (*Request)(request), (*Response)(response)))
}

// GetResourceResponseFilter (get_resource_response_filter)
// Called on the IO thread to optionally filter resource response content.
// |request| and |response| represent the request and response respectively
// and cannot be modified in this callback.
func (d *RequestHandler) GetResourceResponseFilter(browser *Browser, frame *Frame, request *Request, response *Response) *ResponseFilter {
	return lookupRequestHandlerProxy(d.Base()).GetResourceResponseFilter(d, browser, frame, request, response)
}

//export gocef_request_handler_get_resource_response_filter
func gocef_request_handler_get_resource_response_filter(self *C.cef_request_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, request *C.cef_request_t, response *C.cef_response_t) *C.cef_response_filter_t {
	me__ := (*RequestHandler)(self)
	proxy__ := lookupRequestHandlerProxy(me__.Base())
	return (proxy__.GetResourceResponseFilter(me__, (*Browser)(browser), (*Frame)(frame), (*Request)(request), (*Response)(response))).toNative()
}

// OnResourceLoadComplete (on_resource_load_complete)
// Called on the IO thread when a resource load has completed. |request| and
// |response| represent the request and response respectively and cannot be
// modified in this callback. |status| indicates the load completion status.
// |received_content_length| is the number of response bytes actually read.
func (d *RequestHandler) OnResourceLoadComplete(browser *Browser, frame *Frame, request *Request, response *Response, status UrlrequestStatus, received_content_length int64) {
	lookupRequestHandlerProxy(d.Base()).OnResourceLoadComplete(d, browser, frame, request, response, status, received_content_length)
}

//export gocef_request_handler_on_resource_load_complete
func gocef_request_handler_on_resource_load_complete(self *C.cef_request_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, request *C.cef_request_t, response *C.cef_response_t, status C.cef_urlrequest_status_t, received_content_length C.int64) {
	me__ := (*RequestHandler)(self)
	proxy__ := lookupRequestHandlerProxy(me__.Base())
	proxy__.OnResourceLoadComplete(me__, (*Browser)(browser), (*Frame)(frame), (*Request)(request), (*Response)(response), UrlrequestStatus(status), int64(received_content_length))
}

// GetAuthCredentials (get_auth_credentials)
// Called on the IO thread when the browser needs credentials from the user.
// |isProxy| indicates whether the host is a proxy server. |host| contains the
// hostname and |port| contains the port number. |realm| is the realm of the
// challenge and may be NULL. |scheme| is the authentication scheme used, such
// as "basic" or "digest", and will be NULL if the source of the request is an
// FTP server. Return true (1) to continue the request and call
// cef_auth_callback_t::cont() either in this function or at a later time when
// the authentication information is available. Return false (0) to cancel the
// request immediately.
func (d *RequestHandler) GetAuthCredentials(browser *Browser, frame *Frame, isProxy int32, host string, port int32, realm, scheme string, callback *AuthCallback) int32 {
	return lookupRequestHandlerProxy(d.Base()).GetAuthCredentials(d, browser, frame, isProxy, host, port, realm, scheme, callback)
}

//export gocef_request_handler_get_auth_credentials
func gocef_request_handler_get_auth_credentials(self *C.cef_request_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, isProxy C.int, host *C.cef_string_t, port C.int, realm *C.cef_string_t, scheme *C.cef_string_t, callback *C.cef_auth_callback_t) C.int {
	me__ := (*RequestHandler)(self)
	proxy__ := lookupRequestHandlerProxy(me__.Base())
	host_ := cefstrToString(host)
	realm_ := cefstrToString(realm)
	scheme_ := cefstrToString(scheme)
	return C.int(proxy__.GetAuthCredentials(me__, (*Browser)(browser), (*Frame)(frame), int32(isProxy), host_, int32(port), realm_, scheme_, (*AuthCallback)(callback)))
}

// CanGetCookies (can_get_cookies)
// Called on the IO thread before sending a network request with a "Cookie"
// request header. Return true (1) to allow cookies to be included in the
// network request or false (0) to block cookies. The |request| object should
// not be modified in this callback.
func (d *RequestHandler) CanGetCookies(browser *Browser, frame *Frame, request *Request) int32 {
	return lookupRequestHandlerProxy(d.Base()).CanGetCookies(d, browser, frame, request)
}

//export gocef_request_handler_can_get_cookies
func gocef_request_handler_can_get_cookies(self *C.cef_request_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, request *C.cef_request_t) C.int {
	me__ := (*RequestHandler)(self)
	proxy__ := lookupRequestHandlerProxy(me__.Base())
	return C.int(proxy__.CanGetCookies(me__, (*Browser)(browser), (*Frame)(frame), (*Request)(request)))
}

// CanSetCookie (can_set_cookie)
// Called on the IO thread when receiving a network request with a "Set-
// Cookie" response header value represented by |cookie|. Return true (1) to
// allow the cookie to be stored or false (0) to block the cookie. The
// |request| object should not be modified in this callback.
func (d *RequestHandler) CanSetCookie(browser *Browser, frame *Frame, request *Request, cookie *Cookie) int32 {
	return lookupRequestHandlerProxy(d.Base()).CanSetCookie(d, browser, frame, request, cookie)
}

//export gocef_request_handler_can_set_cookie
func gocef_request_handler_can_set_cookie(self *C.cef_request_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, request *C.cef_request_t, cookie *C.cef_cookie_t) C.int {
	me__ := (*RequestHandler)(self)
	proxy__ := lookupRequestHandlerProxy(me__.Base())
	cookie_ := cookie.toGo()
	return C.int(proxy__.CanSetCookie(me__, (*Browser)(browser), (*Frame)(frame), (*Request)(request), cookie_))
}

// OnQuotaRequest (on_quota_request)
// Called on the IO thread when JavaScript requests a specific storage quota
// size via the webkitStorageInfo.requestQuota function. |origin_url| is the
// origin of the page making the request. |new_size| is the requested quota
// size in bytes. Return true (1) to continue the request and call
// cef_request_tCallback::cont() either in this function or at a later time to
// grant or deny the request. Return false (0) to cancel the request
// immediately.
func (d *RequestHandler) OnQuotaRequest(browser *Browser, origin_url string, new_size int64, callback *RequestCallback) int32 {
	return lookupRequestHandlerProxy(d.Base()).OnQuotaRequest(d, browser, origin_url, new_size, callback)
}

//export gocef_request_handler_on_quota_request
func gocef_request_handler_on_quota_request(self *C.cef_request_handler_t, browser *C.cef_browser_t, origin_url *C.cef_string_t, new_size C.int64, callback *C.cef_request_callback_t) C.int {
	me__ := (*RequestHandler)(self)
	proxy__ := lookupRequestHandlerProxy(me__.Base())
	origin_url_ := cefstrToString(origin_url)
	return C.int(proxy__.OnQuotaRequest(me__, (*Browser)(browser), origin_url_, int64(new_size), (*RequestCallback)(callback)))
}

// OnProtocolExecution (on_protocol_execution)
// Called on the UI thread to handle requests for URLs with an unknown
// protocol component. Set |allow_os_execution| to true (1) to attempt
// execution via the registered OS protocol handler, if any. SECURITY WARNING:
// YOU SHOULD USE THIS METHOD TO ENFORCE RESTRICTIONS BASED ON SCHEME, HOST OR
// OTHER URL ANALYSIS BEFORE ALLOWING OS EXECUTION.
func (d *RequestHandler) OnProtocolExecution(browser *Browser, url string, allow_os_execution *int32) {
	lookupRequestHandlerProxy(d.Base()).OnProtocolExecution(d, browser, url, allow_os_execution)
}

//export gocef_request_handler_on_protocol_execution
func gocef_request_handler_on_protocol_execution(self *C.cef_request_handler_t, browser *C.cef_browser_t, url *C.cef_string_t, allow_os_execution *C.int) {
	me__ := (*RequestHandler)(self)
	proxy__ := lookupRequestHandlerProxy(me__.Base())
	url_ := cefstrToString(url)
	proxy__.OnProtocolExecution(me__, (*Browser)(browser), url_, (*int32)(allow_os_execution))
}

// OnCertificateError (on_certificate_error)
// Called on the UI thread to handle requests for URLs with an invalid SSL
// certificate. Return true (1) and call cef_request_tCallback::cont() either
// in this function or at a later time to continue or cancel the request.
// Return false (0) to cancel the request immediately. If
// CefSettings.ignore_certificate_errors is set all invalid certificates will
// be accepted without calling this function.
func (d *RequestHandler) OnCertificateError(browser *Browser, cert_error Errorcode, request_url string, ssl_info *Sslinfo, callback *RequestCallback) int32 {
	return lookupRequestHandlerProxy(d.Base()).OnCertificateError(d, browser, cert_error, request_url, ssl_info, callback)
}

//export gocef_request_handler_on_certificate_error
func gocef_request_handler_on_certificate_error(self *C.cef_request_handler_t, browser *C.cef_browser_t, cert_error C.cef_errorcode_t, request_url *C.cef_string_t, ssl_info *C.cef_sslinfo_t, callback *C.cef_request_callback_t) C.int {
	me__ := (*RequestHandler)(self)
	proxy__ := lookupRequestHandlerProxy(me__.Base())
	request_url_ := cefstrToString(request_url)
	return C.int(proxy__.OnCertificateError(me__, (*Browser)(browser), Errorcode(cert_error), request_url_, (*Sslinfo)(ssl_info), (*RequestCallback)(callback)))
}

// OnSelectClientCertificate (on_select_client_certificate)
// Called on the UI thread when a client certificate is being requested for
// authentication. Return false (0) to use the default behavior and
// automatically select the first certificate available. Return true (1) and
// call cef_select_client_certificate_callback_t::Select either in this
// function or at a later time to select a certificate. Do not call Select or
// call it with NULL to continue without using any certificate. |isProxy|
// indicates whether the host is an HTTPS proxy or the origin server. |host|
// and |port| contains the hostname and port of the SSL server. |certificates|
// is the list of certificates to choose from; this list has already been
// pruned by Chromium so that it only contains certificates from issuers that
// the server trusts.
func (d *RequestHandler) OnSelectClientCertificate(browser *Browser, isProxy int32, host string, port int32, certificatesCount uint64, certificates **X509certificate, callback *SelectClientCertificateCallback) int32 {
	return lookupRequestHandlerProxy(d.Base()).OnSelectClientCertificate(d, browser, isProxy, host, port, certificatesCount, certificates, callback)
}

//export gocef_request_handler_on_select_client_certificate
func gocef_request_handler_on_select_client_certificate(self *C.cef_request_handler_t, browser *C.cef_browser_t, isProxy C.int, host *C.cef_string_t, port C.int, certificatesCount C.size_t, certificates **C.cef_x509certificate_t, callback *C.cef_select_client_certificate_callback_t) C.int {
	me__ := (*RequestHandler)(self)
	proxy__ := lookupRequestHandlerProxy(me__.Base())
	host_ := cefstrToString(host)
	certificates_ := (*X509certificate)(*certificates)
	certificates__p := &certificates_
	return C.int(proxy__.OnSelectClientCertificate(me__, (*Browser)(browser), int32(isProxy), host_, int32(port), uint64(certificatesCount), certificates__p, (*SelectClientCertificateCallback)(callback)))
}

// OnPluginCrashed (on_plugin_crashed)
// Called on the browser process UI thread when a plugin has crashed.
// |plugin_path| is the path of the plugin that crashed.
func (d *RequestHandler) OnPluginCrashed(browser *Browser, plugin_path string) {
	lookupRequestHandlerProxy(d.Base()).OnPluginCrashed(d, browser, plugin_path)
}

//export gocef_request_handler_on_plugin_crashed
func gocef_request_handler_on_plugin_crashed(self *C.cef_request_handler_t, browser *C.cef_browser_t, plugin_path *C.cef_string_t) {
	me__ := (*RequestHandler)(self)
	proxy__ := lookupRequestHandlerProxy(me__.Base())
	plugin_path_ := cefstrToString(plugin_path)
	proxy__.OnPluginCrashed(me__, (*Browser)(browser), plugin_path_)
}

// OnRenderViewReady (on_render_view_ready)
// Called on the browser process UI thread when the render view associated
// with |browser| is ready to receive/handle IPC messages in the render
// process.
func (d *RequestHandler) OnRenderViewReady(browser *Browser) {
	lookupRequestHandlerProxy(d.Base()).OnRenderViewReady(d, browser)
}

//export gocef_request_handler_on_render_view_ready
func gocef_request_handler_on_render_view_ready(self *C.cef_request_handler_t, browser *C.cef_browser_t) {
	me__ := (*RequestHandler)(self)
	proxy__ := lookupRequestHandlerProxy(me__.Base())
	proxy__.OnRenderViewReady(me__, (*Browser)(browser))
}

// OnRenderProcessTerminated (on_render_process_terminated)
// Called on the browser process UI thread when the render process terminates
// unexpectedly. |status| indicates how the process terminated.
func (d *RequestHandler) OnRenderProcessTerminated(browser *Browser, status TerminationStatus) {
	lookupRequestHandlerProxy(d.Base()).OnRenderProcessTerminated(d, browser, status)
}

//export gocef_request_handler_on_render_process_terminated
func gocef_request_handler_on_render_process_terminated(self *C.cef_request_handler_t, browser *C.cef_browser_t, status C.cef_termination_status_t) {
	me__ := (*RequestHandler)(self)
	proxy__ := lookupRequestHandlerProxy(me__.Base())
	proxy__.OnRenderProcessTerminated(me__, (*Browser)(browser), TerminationStatus(status))
}
