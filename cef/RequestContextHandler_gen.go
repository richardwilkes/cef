// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

import (
	// #include "RequestContextHandler_gen.h"
	"C"
)

// RequestContextHandlerProxy defines methods required for using RequestContextHandler.
type RequestContextHandlerProxy interface {
	OnRequestContextInitialized(self *RequestContextHandler, request_context *RequestContext)
	OnBeforePluginLoad(self *RequestContextHandler, mime_type, plugin_url string, is_main_frame int32, top_origin_url string, plugin_info *WebPluginInfo, plugin_policy *PluginPolicy) int32
	GetResourceRequestHandler(self *RequestContextHandler, browser *Browser, frame *Frame, request *Request, is_navigation, is_download int32, request_initiator string, disable_default_handling *int32) *ResourceRequestHandler
}

// RequestContextHandler (cef_request_context_handler_t from include/capi/cef_request_context_handler_capi.h)
// Implement this structure to provide handler implementations. The handler
// instance will not be released until all objects related to the context have
// been destroyed.
type RequestContextHandler C.cef_request_context_handler_t

// NewRequestContextHandler creates a new RequestContextHandler with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewRequestContextHandler(proxy RequestContextHandlerProxy) *RequestContextHandler {
	result := (*RequestContextHandler)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_request_context_handler_t, proxy)))
	if proxy != nil {
		C.gocef_set_request_context_handler_proxy(result.toNative())
	}
	return result
}

func (d *RequestContextHandler) toNative() *C.cef_request_context_handler_t {
	return (*C.cef_request_context_handler_t)(d)
}

func lookupRequestContextHandlerProxy(obj *BaseRefCounted) RequestContextHandlerProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(RequestContextHandlerProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type RequestContextHandlerProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *RequestContextHandler) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnRequestContextInitialized (on_request_context_initialized)
// Called on the browser process UI thread immediately after the request
// context has been initialized.
func (d *RequestContextHandler) OnRequestContextInitialized(request_context *RequestContext) {
	lookupRequestContextHandlerProxy(d.Base()).OnRequestContextInitialized(d, request_context)
}

//nolint:gocritic
//export gocef_request_context_handler_on_request_context_initialized
func gocef_request_context_handler_on_request_context_initialized(self *C.cef_request_context_handler_t, request_context *C.cef_request_context_t) {
	me__ := (*RequestContextHandler)(self)
	proxy__ := lookupRequestContextHandlerProxy(me__.Base())
	proxy__.OnRequestContextInitialized(me__, (*RequestContext)(request_context))
}

// OnBeforePluginLoad (on_before_plugin_load)
// Called on multiple browser process threads before a plugin instance is
// loaded. |mime_type| is the mime type of the plugin that will be loaded.
// |plugin_url| is the content URL that the plugin will load and may be NULL.
// |is_main_frame| will be true (1) if the plugin is being loaded in the main
// (top-level) frame, |top_origin_url| is the URL for the top-level frame that
// contains the plugin when loading a specific plugin instance or NULL when
// building the initial list of enabled plugins for 'navigator.plugins'
// JavaScript state. |plugin_info| includes additional information about the
// plugin that will be loaded. |plugin_policy| is the recommended policy.
// Modify |plugin_policy| and return true (1) to change the policy. Return
// false (0) to use the recommended policy. The default plugin policy can be
// set at runtime using the `--plugin-policy=[allow|detect|block]` command-
// line flag. Decisions to mark a plugin as disabled by setting
// |plugin_policy| to PLUGIN_POLICY_DISABLED may be cached when
// |top_origin_url| is NULL. To purge the plugin list cache and potentially
// trigger new calls to this function call
// cef_request_tContext::PurgePluginListCache.
func (d *RequestContextHandler) OnBeforePluginLoad(mime_type, plugin_url string, is_main_frame int32, top_origin_url string, plugin_info *WebPluginInfo, plugin_policy *PluginPolicy) int32 {
	return lookupRequestContextHandlerProxy(d.Base()).OnBeforePluginLoad(d, mime_type, plugin_url, is_main_frame, top_origin_url, plugin_info, plugin_policy)
}

//nolint:gocritic
//export gocef_request_context_handler_on_before_plugin_load
func gocef_request_context_handler_on_before_plugin_load(self *C.cef_request_context_handler_t, mime_type *C.cef_string_t, plugin_url *C.cef_string_t, is_main_frame C.int, top_origin_url *C.cef_string_t, plugin_info *C.cef_web_plugin_info_t, plugin_policy *C.cef_plugin_policy_t) C.int {
	me__ := (*RequestContextHandler)(self)
	proxy__ := lookupRequestContextHandlerProxy(me__.Base())
	mime_type_ := cefstrToString(mime_type)
	plugin_url_ := cefstrToString(plugin_url)
	top_origin_url_ := cefstrToString(top_origin_url)
	plugin_policy_ := PluginPolicy(*plugin_policy)
	return C.int(proxy__.OnBeforePluginLoad(me__, mime_type_, plugin_url_, int32(is_main_frame), top_origin_url_, (*WebPluginInfo)(plugin_info), &plugin_policy_))
}

// GetResourceRequestHandler (get_resource_request_handler)
// Called on the browser process IO thread before a resource request is
// initiated. The |browser| and |frame| values represent the source of the
// request, and may be NULL for requests originating from service workers or
// cef_urlrequest_t. |request| represents the request contents and cannot be
// modified in this callback. |is_navigation| will be true (1) if the resource
// request is a navigation. |is_download| will be true (1) if the resource
// request is a download. |request_initiator| is the origin (scheme + domain)
// of the page that initiated the request. Set |disable_default_handling| to
// true (1) to disable default handling of the request, in which case it will
// need to be handled via cef_resource_request_handler_t::GetResourceHandler
// or it will be canceled. To allow the resource load to proceed with default
// handling return NULL. To specify a handler for the resource return a
// cef_resource_request_handler_t object. This function will not be called if
// the client associated with |browser| returns a non-NULL value from
// cef_request_tHandler::GetResourceRequestHandler for the same request
// (identified by cef_request_t::GetIdentifier).
func (d *RequestContextHandler) GetResourceRequestHandler(browser *Browser, frame *Frame, request *Request, is_navigation, is_download int32, request_initiator string, disable_default_handling *int32) *ResourceRequestHandler {
	return lookupRequestContextHandlerProxy(d.Base()).GetResourceRequestHandler(d, browser, frame, request, is_navigation, is_download, request_initiator, disable_default_handling)
}

//nolint:gocritic
//export gocef_request_context_handler_get_resource_request_handler
func gocef_request_context_handler_get_resource_request_handler(self *C.cef_request_context_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, request *C.cef_request_t, is_navigation C.int, is_download C.int, request_initiator *C.cef_string_t, disable_default_handling *C.int) *C.cef_resource_request_handler_t {
	me__ := (*RequestContextHandler)(self)
	proxy__ := lookupRequestContextHandlerProxy(me__.Base())
	request_initiator_ := cefstrToString(request_initiator)
	return (proxy__.GetResourceRequestHandler(me__, (*Browser)(browser), (*Frame)(frame), (*Request)(request), int32(is_navigation), int32(is_download), request_initiator_, (*int32)(disable_default_handling))).toNative()
}
