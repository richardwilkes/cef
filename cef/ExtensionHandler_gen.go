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
	// #include "ExtensionHandler_gen.h"
	"C"
)

// ExtensionHandlerProxy defines methods required for using ExtensionHandler.
type ExtensionHandlerProxy interface {
	OnExtensionLoadFailed(self *ExtensionHandler, result Errorcode)
	OnExtensionLoaded(self *ExtensionHandler, extension *Extension)
	OnExtensionUnloaded(self *ExtensionHandler, extension *Extension)
	OnBeforeBackgroundBrowser(self *ExtensionHandler, extension *Extension, uRL string, client **Client, settings *BrowserSettings) int32
	OnBeforeBrowser(self *ExtensionHandler, extension *Extension, browser, activeBrowser *Browser, index int32, uRL string, active int32, windowInfo *WindowInfo, client **Client, settings *BrowserSettings) int32
	GetActiveBrowser(self *ExtensionHandler, extension *Extension, browser *Browser, includeIncognito int32) *Browser
	CanAccessBrowser(self *ExtensionHandler, extension *Extension, browser *Browser, includeIncognito int32, targetBrowser *Browser) int32
	GetExtensionResource(self *ExtensionHandler, extension *Extension, browser *Browser, file string, callback *GetExtensionResourceCallback) int32
}

// ExtensionHandler (cef_extension_handler_t from include/capi/cef_extension_handler_capi.h)
// Implement this structure to handle events related to browser extensions. The
// functions of this structure will be called on the UI thread. See
// cef_request_tContext::LoadExtension for information about extension loading.
type ExtensionHandler C.cef_extension_handler_t

// NewExtensionHandler creates a new ExtensionHandler with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewExtensionHandler(proxy ExtensionHandlerProxy) *ExtensionHandler {
	result := (*ExtensionHandler)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_extension_handler_t, proxy)))
	if proxy != nil {
		C.gocef_set_extension_handler_proxy(result.toNative())
	}
	return result
}

func (d *ExtensionHandler) toNative() *C.cef_extension_handler_t {
	return (*C.cef_extension_handler_t)(d)
}

func lookupExtensionHandlerProxy(obj *BaseRefCounted) ExtensionHandlerProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(ExtensionHandlerProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type ExtensionHandlerProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *ExtensionHandler) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnExtensionLoadFailed (on_extension_load_failed)
// Called if the cef_request_tContext::LoadExtension request fails. |result|
// will be the error code.
func (d *ExtensionHandler) OnExtensionLoadFailed(result Errorcode) {
	lookupExtensionHandlerProxy(d.Base()).OnExtensionLoadFailed(d, result)
}

//nolint:gocritic
//export gocef_extension_handler_on_extension_load_failed
func gocef_extension_handler_on_extension_load_failed(self *C.cef_extension_handler_t, result C.cef_errorcode_t) {
	me__ := (*ExtensionHandler)(self)
	proxy__ := lookupExtensionHandlerProxy(me__.Base())
	proxy__.OnExtensionLoadFailed(me__, Errorcode(result))
}

// OnExtensionLoaded (on_extension_loaded)
// Called if the cef_request_tContext::LoadExtension request succeeds.
// |extension| is the loaded extension.
func (d *ExtensionHandler) OnExtensionLoaded(extension *Extension) {
	lookupExtensionHandlerProxy(d.Base()).OnExtensionLoaded(d, extension)
}

//nolint:gocritic
//export gocef_extension_handler_on_extension_loaded
func gocef_extension_handler_on_extension_loaded(self *C.cef_extension_handler_t, extension *C.cef_extension_t) {
	me__ := (*ExtensionHandler)(self)
	proxy__ := lookupExtensionHandlerProxy(me__.Base())
	proxy__.OnExtensionLoaded(me__, (*Extension)(extension))
}

// OnExtensionUnloaded (on_extension_unloaded)
// Called after the cef_extension_t::Unload request has completed.
func (d *ExtensionHandler) OnExtensionUnloaded(extension *Extension) {
	lookupExtensionHandlerProxy(d.Base()).OnExtensionUnloaded(d, extension)
}

//nolint:gocritic
//export gocef_extension_handler_on_extension_unloaded
func gocef_extension_handler_on_extension_unloaded(self *C.cef_extension_handler_t, extension *C.cef_extension_t) {
	me__ := (*ExtensionHandler)(self)
	proxy__ := lookupExtensionHandlerProxy(me__.Base())
	proxy__.OnExtensionUnloaded(me__, (*Extension)(extension))
}

// OnBeforeBackgroundBrowser (on_before_background_browser)
// Called when an extension needs a browser to host a background script
// specified via the "background" manifest key. The browser will have no
// visible window and cannot be displayed. |extension| is the extension that
// is loading the background script. |url| is an internally generated
// reference to an HTML page that will be used to load the background script
// via a <script> src attribute. To allow creation of the browser optionally
// modify |client| and |settings| and return false (0). To cancel creation of
// the browser (and consequently cancel load of the background script) return
// true (1). Successful creation will be indicated by a call to
// cef_life_span_handler_t::OnAfterCreated, and
// cef_browser_host_t::IsBackgroundHost will return true (1) for the resulting
// browser. See https://developer.chrome.com/extensions/event_pages for more
// information about extension background script usage.
func (d *ExtensionHandler) OnBeforeBackgroundBrowser(extension *Extension, uRL string, client **Client, settings *BrowserSettings) int32 {
	return lookupExtensionHandlerProxy(d.Base()).OnBeforeBackgroundBrowser(d, extension, uRL, client, settings)
}

//nolint:gocritic
//export gocef_extension_handler_on_before_background_browser
func gocef_extension_handler_on_before_background_browser(self *C.cef_extension_handler_t, extension *C.cef_extension_t, uRL *C.cef_string_t, client **C.cef_client_t, settings *C.cef_browser_settings_t) C.int {
	me__ := (*ExtensionHandler)(self)
	proxy__ := lookupExtensionHandlerProxy(me__.Base())
	uRL_ := cefstrToString(uRL)
	client_ := (*Client)(*client)
	client__p := &client_
	settings_ := settings.toGo()
	return C.int(proxy__.OnBeforeBackgroundBrowser(me__, (*Extension)(extension), uRL_, client__p, settings_))
}

// OnBeforeBrowser (on_before_browser)
// Called when an extension API (e.g. chrome.tabs.create) requests creation of
// a new browser. |extension| and |browser| are the source of the API call.
// |active_browser| may optionally be specified via the windowId property or
// returned via the get_active_browser() callback and provides the default
// |client| and |settings| values for the new browser. |index| is the position
// value optionally specified via the index property. |url| is the URL that
// will be loaded in the browser. |active| is true (1) if the new browser
// should be active when opened.  To allow creation of the browser optionally
// modify |windowInfo|, |client| and |settings| and return false (0). To
// cancel creation of the browser return true (1). Successful creation will be
// indicated by a call to cef_life_span_handler_t::OnAfterCreated. Any
// modifications to |windowInfo| will be ignored if |active_browser| is
// wrapped in a cef_browser_view_t.
func (d *ExtensionHandler) OnBeforeBrowser(extension *Extension, browser, activeBrowser *Browser, index int32, uRL string, active int32, windowInfo *WindowInfo, client **Client, settings *BrowserSettings) int32 {
	return lookupExtensionHandlerProxy(d.Base()).OnBeforeBrowser(d, extension, browser, activeBrowser, index, uRL, active, windowInfo, client, settings)
}

//nolint:gocritic
//export gocef_extension_handler_on_before_browser
func gocef_extension_handler_on_before_browser(self *C.cef_extension_handler_t, extension *C.cef_extension_t, browser *C.cef_browser_t, activeBrowser *C.cef_browser_t, index C.int, uRL *C.cef_string_t, active C.int, windowInfo *C.cef_window_info_t, client **C.cef_client_t, settings *C.cef_browser_settings_t) C.int {
	me__ := (*ExtensionHandler)(self)
	proxy__ := lookupExtensionHandlerProxy(me__.Base())
	uRL_ := cefstrToString(uRL)
	windowInfo_ := windowInfo.toGo()
	client_ := (*Client)(*client)
	client__p := &client_
	settings_ := settings.toGo()
	return C.int(proxy__.OnBeforeBrowser(me__, (*Extension)(extension), (*Browser)(browser), (*Browser)(activeBrowser), int32(index), uRL_, int32(active), windowInfo_, client__p, settings_))
}

// GetActiveBrowser (get_active_browser)
// Called when no tabId is specified to an extension API call that accepts a
// tabId parameter (e.g. chrome.tabs.*). |extension| and |browser| are the
// source of the API call. Return the browser that will be acted on by the API
// call or return NULL to act on |browser|. The returned browser must share
// the same cef_request_tContext as |browser|. Incognito browsers should not
// be considered unless the source extension has incognito access enabled, in
// which case |include_incognito| will be true (1).
func (d *ExtensionHandler) GetActiveBrowser(extension *Extension, browser *Browser, includeIncognito int32) *Browser {
	return lookupExtensionHandlerProxy(d.Base()).GetActiveBrowser(d, extension, browser, includeIncognito)
}

//nolint:gocritic
//export gocef_extension_handler_get_active_browser
func gocef_extension_handler_get_active_browser(self *C.cef_extension_handler_t, extension *C.cef_extension_t, browser *C.cef_browser_t, includeIncognito C.int) *C.cef_browser_t {
	me__ := (*ExtensionHandler)(self)
	proxy__ := lookupExtensionHandlerProxy(me__.Base())
	return (proxy__.GetActiveBrowser(me__, (*Extension)(extension), (*Browser)(browser), int32(includeIncognito))).toNative()
}

// CanAccessBrowser (can_access_browser)
// Called when the tabId associated with |target_browser| is specified to an
// extension API call that accepts a tabId parameter (e.g. chrome.tabs.*).
// |extension| and |browser| are the source of the API call. Return true (1)
// to allow access of false (0) to deny access. Access to incognito browsers
// should not be allowed unless the source extension has incognito access
// enabled, in which case |include_incognito| will be true (1).
func (d *ExtensionHandler) CanAccessBrowser(extension *Extension, browser *Browser, includeIncognito int32, targetBrowser *Browser) int32 {
	return lookupExtensionHandlerProxy(d.Base()).CanAccessBrowser(d, extension, browser, includeIncognito, targetBrowser)
}

//nolint:gocritic
//export gocef_extension_handler_can_access_browser
func gocef_extension_handler_can_access_browser(self *C.cef_extension_handler_t, extension *C.cef_extension_t, browser *C.cef_browser_t, includeIncognito C.int, targetBrowser *C.cef_browser_t) C.int {
	me__ := (*ExtensionHandler)(self)
	proxy__ := lookupExtensionHandlerProxy(me__.Base())
	return C.int(proxy__.CanAccessBrowser(me__, (*Extension)(extension), (*Browser)(browser), int32(includeIncognito), (*Browser)(targetBrowser)))
}

// GetExtensionResource (get_extension_resource)
// Called to retrieve an extension resource that would normally be loaded from
// disk (e.g. if a file parameter is specified to chrome.tabs.executeScript).
// |extension| and |browser| are the source of the resource request. |file| is
// the requested relative file path. To handle the resource request return
// true (1) and execute |callback| either synchronously or asynchronously. For
// the default behavior which reads the resource from the extension directory
// on disk return false (0). Localization substitutions will not be applied to
// resources handled via this function.
func (d *ExtensionHandler) GetExtensionResource(extension *Extension, browser *Browser, file string, callback *GetExtensionResourceCallback) int32 {
	return lookupExtensionHandlerProxy(d.Base()).GetExtensionResource(d, extension, browser, file, callback)
}

//nolint:gocritic
//export gocef_extension_handler_get_extension_resource
func gocef_extension_handler_get_extension_resource(self *C.cef_extension_handler_t, extension *C.cef_extension_t, browser *C.cef_browser_t, file *C.cef_string_t, callback *C.cef_get_extension_resource_callback_t) C.int {
	me__ := (*ExtensionHandler)(self)
	proxy__ := lookupExtensionHandlerProxy(me__.Base())
	file_ := cefstrToString(file)
	return C.int(proxy__.GetExtensionResource(me__, (*Extension)(extension), (*Browser)(browser), file_, (*GetExtensionResourceCallback)(callback)))
}
