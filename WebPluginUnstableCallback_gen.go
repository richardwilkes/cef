// Code generated - DO NOT EDIT.

package cef

import (
	// #include "WebPluginUnstableCallback_gen.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

// WebPluginUnstableCallbackProxy defines methods required for using WebPluginUnstableCallback.
type WebPluginUnstableCallbackProxy interface {
	IsUnstable(self *WebPluginUnstableCallback, path string, unstable int32)
}

// WebPluginUnstableCallback (cef_web_plugin_unstable_callback_t from include/capi/cef_web_plugin_capi.h)
// Structure to implement for receiving unstable plugin information. The
// functions of this structure will be called on the browser process IO thread.
type WebPluginUnstableCallback C.cef_web_plugin_unstable_callback_t

// NewWebPluginUnstableCallback creates a new WebPluginUnstableCallback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewWebPluginUnstableCallback(proxy WebPluginUnstableCallbackProxy) *WebPluginUnstableCallback {
	result := (*WebPluginUnstableCallback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_web_plugin_unstable_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_web_plugin_unstable_callback_proxy(result.toNative())
	}
	return result
}

func (d *WebPluginUnstableCallback) toNative() *C.cef_web_plugin_unstable_callback_t {
	return (*C.cef_web_plugin_unstable_callback_t)(d)
}

func lookupWebPluginUnstableCallbackProxy(obj *BaseRefCounted) WebPluginUnstableCallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(WebPluginUnstableCallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type WebPluginUnstableCallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *WebPluginUnstableCallback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// IsUnstable (is_unstable)
// Method that will be called for the requested plugin. |unstable| will be
// true (1) if the plugin has reached the crash count threshold of 3 times in
// 120 seconds.
func (d *WebPluginUnstableCallback) IsUnstable(path string, unstable int32) {
	lookupWebPluginUnstableCallbackProxy(d.Base()).IsUnstable(d, path, unstable)
}

//export gocef_web_plugin_unstable_callback_is_unstable
func gocef_web_plugin_unstable_callback_is_unstable(self *C.cef_web_plugin_unstable_callback_t, path *C.cef_string_t, unstable C.int) {
	me__ := (*WebPluginUnstableCallback)(self)
	proxy__ := lookupWebPluginUnstableCallbackProxy(me__.Base())
	proxy__.IsUnstable(me__, cefstrToString(path), int32(unstable))
}
