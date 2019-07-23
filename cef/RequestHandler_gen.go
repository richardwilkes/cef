// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

import (
	// #include "RequestHandler_gen.h"
	"C"
)

// RequestHandlerProxy defines methods required for using RequestHandler.
type RequestHandlerProxy interface {
	OnBeforeBrowse(self *RequestHandler, browser *Browser, frame *Frame, request *Request, user_gesture, is_redirect int32) int32
	OnOpenUrlfromTab(self *RequestHandler, browser *Browser, frame *Frame, target_url string, target_disposition WindowOpenDisposition, user_gesture int32) int32
	GetResourceRequestHandler(self *RequestHandler, browser *Browser, frame *Frame, request *Request, is_navigation, is_download int32, request_initiator string, disable_default_handling *int32) *ResourceRequestHandler
	GetAuthCredentials(self *RequestHandler, browser *Browser, origin_url string, isProxy int32, host string, port int32, realm, scheme string, callback *AuthCallback) int32
	OnQuotaRequest(self *RequestHandler, browser *Browser, origin_url string, new_size int64, callback *RequestCallback) int32
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

//nolint:gocritic
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

//nolint:gocritic
//export gocef_request_handler_on_open_urlfrom_tab
func gocef_request_handler_on_open_urlfrom_tab(self *C.cef_request_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, target_url *C.cef_string_t, target_disposition C.cef_window_open_disposition_t, user_gesture C.int) C.int {
	me__ := (*RequestHandler)(self)
	proxy__ := lookupRequestHandlerProxy(me__.Base())
	target_url_ := cefstrToString(target_url)
	return C.int(proxy__.OnOpenUrlfromTab(me__, (*Browser)(browser), (*Frame)(frame), target_url_, WindowOpenDisposition(target_disposition), int32(user_gesture)))
}

// GetResourceRequestHandler (get_resource_request_handler)
// Called on the browser process IO thread before a resource request is
// initiated. The |browser| and |frame| values represent the source of the
// request. |request| represents the request contents and cannot be modified
// in this callback. |is_navigation| will be true (1) if the resource request
// is a navigation. |is_download| will be true (1) if the resource request is
// a download. |request_initiator| is the origin (scheme + domain) of the page
// that initiated the request. Set |disable_default_handling| to true (1) to
// disable default handling of the request, in which case it will need to be
// handled via cef_resource_request_handler_t::GetResourceHandler or it will
// be canceled. To allow the resource load to proceed with default handling
// return NULL. To specify a handler for the resource return a
// cef_resource_request_handler_t object. If this callback returns NULL the
// same function will be called on the associated cef_request_tContextHandler,
// if any.
func (d *RequestHandler) GetResourceRequestHandler(browser *Browser, frame *Frame, request *Request, is_navigation, is_download int32, request_initiator string, disable_default_handling *int32) *ResourceRequestHandler {
	return lookupRequestHandlerProxy(d.Base()).GetResourceRequestHandler(d, browser, frame, request, is_navigation, is_download, request_initiator, disable_default_handling)
}

//nolint:gocritic
//export gocef_request_handler_get_resource_request_handler
func gocef_request_handler_get_resource_request_handler(self *C.cef_request_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, request *C.cef_request_t, is_navigation C.int, is_download C.int, request_initiator *C.cef_string_t, disable_default_handling *C.int) *C.cef_resource_request_handler_t {
	me__ := (*RequestHandler)(self)
	proxy__ := lookupRequestHandlerProxy(me__.Base())
	request_initiator_ := cefstrToString(request_initiator)
	return (proxy__.GetResourceRequestHandler(me__, (*Browser)(browser), (*Frame)(frame), (*Request)(request), int32(is_navigation), int32(is_download), request_initiator_, (*int32)(disable_default_handling))).toNative()
}

// GetAuthCredentials (get_auth_credentials)
// Called on the IO thread when the browser needs credentials from the user.
// |origin_url| is the origin making this authentication request. |isProxy|
// indicates whether the host is a proxy server. |host| contains the hostname
// and |port| contains the port number. |realm| is the realm of the challenge
// and may be NULL. |scheme| is the authentication scheme used, such as
// "basic" or "digest", and will be NULL if the source of the request is an
// FTP server. Return true (1) to continue the request and call
// cef_auth_callback_t::cont() either in this function or at a later time when
// the authentication information is available. Return false (0) to cancel the
// request immediately.
func (d *RequestHandler) GetAuthCredentials(browser *Browser, origin_url string, isProxy int32, host string, port int32, realm, scheme string, callback *AuthCallback) int32 {
	return lookupRequestHandlerProxy(d.Base()).GetAuthCredentials(d, browser, origin_url, isProxy, host, port, realm, scheme, callback)
}

//nolint:gocritic
//export gocef_request_handler_get_auth_credentials
func gocef_request_handler_get_auth_credentials(self *C.cef_request_handler_t, browser *C.cef_browser_t, origin_url *C.cef_string_t, isProxy C.int, host *C.cef_string_t, port C.int, realm *C.cef_string_t, scheme *C.cef_string_t, callback *C.cef_auth_callback_t) C.int {
	me__ := (*RequestHandler)(self)
	proxy__ := lookupRequestHandlerProxy(me__.Base())
	origin_url_ := cefstrToString(origin_url)
	host_ := cefstrToString(host)
	realm_ := cefstrToString(realm)
	scheme_ := cefstrToString(scheme)
	return C.int(proxy__.GetAuthCredentials(me__, (*Browser)(browser), origin_url_, int32(isProxy), host_, int32(port), realm_, scheme_, (*AuthCallback)(callback)))
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

//nolint:gocritic
//export gocef_request_handler_on_quota_request
func gocef_request_handler_on_quota_request(self *C.cef_request_handler_t, browser *C.cef_browser_t, origin_url *C.cef_string_t, new_size C.int64, callback *C.cef_request_callback_t) C.int {
	me__ := (*RequestHandler)(self)
	proxy__ := lookupRequestHandlerProxy(me__.Base())
	origin_url_ := cefstrToString(origin_url)
	return C.int(proxy__.OnQuotaRequest(me__, (*Browser)(browser), origin_url_, int64(new_size), (*RequestCallback)(callback)))
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

//nolint:gocritic
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

//nolint:gocritic
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

//nolint:gocritic
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

//nolint:gocritic
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

//nolint:gocritic
//export gocef_request_handler_on_render_process_terminated
func gocef_request_handler_on_render_process_terminated(self *C.cef_request_handler_t, browser *C.cef_browser_t, status C.cef_termination_status_t) {
	me__ := (*RequestHandler)(self)
	proxy__ := lookupRequestHandlerProxy(me__.Base())
	proxy__.OnRenderProcessTerminated(me__, (*Browser)(browser), TerminationStatus(status))
}
