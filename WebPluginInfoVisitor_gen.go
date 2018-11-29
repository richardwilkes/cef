package cef

import (
	// #include "WebPluginInfoVisitor_gen.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

// WebPluginInfoVisitorProxy defines methods required for using WebPluginInfoVisitor.
type WebPluginInfoVisitorProxy interface {
	Visit(self *WebPluginInfoVisitor, info *WebPluginInfo, count int32, total int32) int32
}

// WebPluginInfoVisitor (cef_web_plugin_info_visitor_t from include/capi/cef_web_plugin_capi.h)
// Structure to implement for visiting web plugin information. The functions of
// this structure will be called on the browser process UI thread.
type WebPluginInfoVisitor C.cef_web_plugin_info_visitor_t

// NewWebPluginInfoVisitor creates a new WebPluginInfoVisitor with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewWebPluginInfoVisitor(proxy WebPluginInfoVisitorProxy) *WebPluginInfoVisitor {
	result := (*WebPluginInfoVisitor)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_web_plugin_info_visitor_t, proxy)))
	if proxy != nil {
		C.gocef_set_web_plugin_info_visitor_proxy(result.toNative())
	}
	return result
}

func (d *WebPluginInfoVisitor) toNative() *C.cef_web_plugin_info_visitor_t {
	return (*C.cef_web_plugin_info_visitor_t)(d)
}

func lookupWebPluginInfoVisitorProxy(obj *BaseRefCounted) WebPluginInfoVisitorProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(WebPluginInfoVisitorProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type WebPluginInfoVisitorProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *WebPluginInfoVisitor) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Visit (visit)
// Method that will be called once for each plugin. |count| is the 0-based
// index for the current plugin. |total| is the total number of plugins.
// Return false (0) to stop visiting plugins. This function may never be
// called if no plugins are found.
func (d *WebPluginInfoVisitor) Visit(info *WebPluginInfo, count, total int32) int32 {
	return lookupWebPluginInfoVisitorProxy(d.Base()).Visit(d, info, count, total)
}

//export gocef_web_plugin_info_visitor_visit
func gocef_web_plugin_info_visitor_visit(self *C.cef_web_plugin_info_visitor_t, info *C.cef_web_plugin_info_t, count C.int, total C.int) C.int {
	me__ := (*WebPluginInfoVisitor)(self)
	proxy__ := lookupWebPluginInfoVisitorProxy(me__.Base())
	return C.int(proxy__.Visit(me__, (*WebPluginInfo)(info), int32(count), int32(total)))
}
