// Code generated - DO NOT EDIT.

package cef

import (
	// #include "RenderProcessHandler_gen.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

// RenderProcessHandlerProxy defines methods required for using RenderProcessHandler.
type RenderProcessHandlerProxy interface {
	OnRenderThreadCreated(self *RenderProcessHandler, extra_info *ListValue)
	OnWebKitInitialized(self *RenderProcessHandler)
	OnBrowserCreated(self *RenderProcessHandler, browser *Browser)
	OnBrowserDestroyed(self *RenderProcessHandler, browser *Browser)
	GetLoadHandler(self *RenderProcessHandler) *LoadHandler
	OnContextCreated(self *RenderProcessHandler, browser *Browser, frame *Frame, context *V8context)
	OnContextReleased(self *RenderProcessHandler, browser *Browser, frame *Frame, context *V8context)
	OnUncaughtException(self *RenderProcessHandler, browser *Browser, frame *Frame, context *V8context, exception *V8exception, stackTrace *V8stackTrace)
	OnFocusedNodeChanged(self *RenderProcessHandler, browser *Browser, frame *Frame, node *Domnode)
	OnProcessMessageReceived(self *RenderProcessHandler, browser *Browser, source_process ProcessID, message *ProcessMessage) int32
}

// RenderProcessHandler (cef_render_process_handler_t from include/capi/cef_render_process_handler_capi.h)
// Structure used to implement render process callbacks. The functions of this
// structure will be called on the render process main thread (TID_RENDERER)
// unless otherwise indicated.
type RenderProcessHandler C.cef_render_process_handler_t

// NewRenderProcessHandler creates a new RenderProcessHandler with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewRenderProcessHandler(proxy RenderProcessHandlerProxy) *RenderProcessHandler {
	result := (*RenderProcessHandler)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_render_process_handler_t, proxy)))
	if proxy != nil {
		C.gocef_set_render_process_handler_proxy(result.toNative())
	}
	return result
}

func (d *RenderProcessHandler) toNative() *C.cef_render_process_handler_t {
	return (*C.cef_render_process_handler_t)(d)
}

func lookupRenderProcessHandlerProxy(obj *BaseRefCounted) RenderProcessHandlerProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(RenderProcessHandlerProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type RenderProcessHandlerProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *RenderProcessHandler) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnRenderThreadCreated (on_render_thread_created)
// Called after the render process main thread has been created. |extra_info|
// is a read-only value originating from
// cef_browser_process_handler_t::on_render_process_thread_created(). Do not
// keep a reference to |extra_info| outside of this function.
func (d *RenderProcessHandler) OnRenderThreadCreated(extra_info *ListValue) {
	lookupRenderProcessHandlerProxy(d.Base()).OnRenderThreadCreated(d, extra_info)
}

//export gocef_render_process_handler_on_render_thread_created
func gocef_render_process_handler_on_render_thread_created(self *C.cef_render_process_handler_t, extra_info *C.cef_list_value_t) {
	me__ := (*RenderProcessHandler)(self)
	proxy__ := lookupRenderProcessHandlerProxy(me__.Base())
	proxy__.OnRenderThreadCreated(me__, (*ListValue)(extra_info))
}

// OnWebKitInitialized (on_web_kit_initialized)
// Called after WebKit has been initialized.
func (d *RenderProcessHandler) OnWebKitInitialized() {
	lookupRenderProcessHandlerProxy(d.Base()).OnWebKitInitialized(d)
}

//export gocef_render_process_handler_on_web_kit_initialized
func gocef_render_process_handler_on_web_kit_initialized(self *C.cef_render_process_handler_t) {
	me__ := (*RenderProcessHandler)(self)
	proxy__ := lookupRenderProcessHandlerProxy(me__.Base())
	proxy__.OnWebKitInitialized(me__)
}

// OnBrowserCreated (on_browser_created)
// Called after a browser has been created. When browsing cross-origin a new
// browser will be created before the old browser with the same identifier is
// destroyed.
func (d *RenderProcessHandler) OnBrowserCreated(browser *Browser) {
	lookupRenderProcessHandlerProxy(d.Base()).OnBrowserCreated(d, browser)
}

//export gocef_render_process_handler_on_browser_created
func gocef_render_process_handler_on_browser_created(self *C.cef_render_process_handler_t, browser *C.cef_browser_t) {
	me__ := (*RenderProcessHandler)(self)
	proxy__ := lookupRenderProcessHandlerProxy(me__.Base())
	proxy__.OnBrowserCreated(me__, (*Browser)(browser))
}

// OnBrowserDestroyed (on_browser_destroyed)
// Called before a browser is destroyed.
func (d *RenderProcessHandler) OnBrowserDestroyed(browser *Browser) {
	lookupRenderProcessHandlerProxy(d.Base()).OnBrowserDestroyed(d, browser)
}

//export gocef_render_process_handler_on_browser_destroyed
func gocef_render_process_handler_on_browser_destroyed(self *C.cef_render_process_handler_t, browser *C.cef_browser_t) {
	me__ := (*RenderProcessHandler)(self)
	proxy__ := lookupRenderProcessHandlerProxy(me__.Base())
	proxy__.OnBrowserDestroyed(me__, (*Browser)(browser))
}

// GetLoadHandler (get_load_handler)
// Return the handler for browser load status events.
func (d *RenderProcessHandler) GetLoadHandler() *LoadHandler {
	return lookupRenderProcessHandlerProxy(d.Base()).GetLoadHandler(d)
}

//export gocef_render_process_handler_get_load_handler
func gocef_render_process_handler_get_load_handler(self *C.cef_render_process_handler_t) *C.cef_load_handler_t {
	me__ := (*RenderProcessHandler)(self)
	proxy__ := lookupRenderProcessHandlerProxy(me__.Base())
	return (proxy__.GetLoadHandler(me__)).toNative()
}

// OnContextCreated (on_context_created)
// Called immediately after the V8 context for a frame has been created. To
// retrieve the JavaScript 'window' object use the
// cef_v8context_t::get_global() function. V8 handles can only be accessed
// from the thread on which they are created. A task runner for posting tasks
// on the associated thread can be retrieved via the
// cef_v8context_t::get_task_runner() function.
func (d *RenderProcessHandler) OnContextCreated(browser *Browser, frame *Frame, context *V8context) {
	lookupRenderProcessHandlerProxy(d.Base()).OnContextCreated(d, browser, frame, context)
}

//export gocef_render_process_handler_on_context_created
func gocef_render_process_handler_on_context_created(self *C.cef_render_process_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, context *C.cef_v8context_t) {
	me__ := (*RenderProcessHandler)(self)
	proxy__ := lookupRenderProcessHandlerProxy(me__.Base())
	proxy__.OnContextCreated(me__, (*Browser)(browser), (*Frame)(frame), (*V8context)(context))
}

// OnContextReleased (on_context_released)
// Called immediately before the V8 context for a frame is released. No
// references to the context should be kept after this function is called.
func (d *RenderProcessHandler) OnContextReleased(browser *Browser, frame *Frame, context *V8context) {
	lookupRenderProcessHandlerProxy(d.Base()).OnContextReleased(d, browser, frame, context)
}

//export gocef_render_process_handler_on_context_released
func gocef_render_process_handler_on_context_released(self *C.cef_render_process_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, context *C.cef_v8context_t) {
	me__ := (*RenderProcessHandler)(self)
	proxy__ := lookupRenderProcessHandlerProxy(me__.Base())
	proxy__.OnContextReleased(me__, (*Browser)(browser), (*Frame)(frame), (*V8context)(context))
}

// OnUncaughtException (on_uncaught_exception)
// Called for global uncaught exceptions in a frame. Execution of this
// callback is disabled by default. To enable set
// CefSettings.uncaught_exception_stack_size > 0.
func (d *RenderProcessHandler) OnUncaughtException(browser *Browser, frame *Frame, context *V8context, exception *V8exception, stackTrace *V8stackTrace) {
	lookupRenderProcessHandlerProxy(d.Base()).OnUncaughtException(d, browser, frame, context, exception, stackTrace)
}

//export gocef_render_process_handler_on_uncaught_exception
func gocef_render_process_handler_on_uncaught_exception(self *C.cef_render_process_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, context *C.cef_v8context_t, exception *C.cef_v8exception_t, stackTrace *C.cef_v8stack_trace_t) {
	me__ := (*RenderProcessHandler)(self)
	proxy__ := lookupRenderProcessHandlerProxy(me__.Base())
	proxy__.OnUncaughtException(me__, (*Browser)(browser), (*Frame)(frame), (*V8context)(context), (*V8exception)(exception), (*V8stackTrace)(stackTrace))
}

// OnFocusedNodeChanged (on_focused_node_changed)
// Called when a new node in the the browser gets focus. The |node| value may
// be NULL if no specific node has gained focus. The node object passed to
// this function represents a snapshot of the DOM at the time this function is
// executed. DOM objects are only valid for the scope of this function. Do not
// keep references to or attempt to access any DOM objects outside the scope
// of this function.
func (d *RenderProcessHandler) OnFocusedNodeChanged(browser *Browser, frame *Frame, node *Domnode) {
	lookupRenderProcessHandlerProxy(d.Base()).OnFocusedNodeChanged(d, browser, frame, node)
}

//export gocef_render_process_handler_on_focused_node_changed
func gocef_render_process_handler_on_focused_node_changed(self *C.cef_render_process_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, node *C.cef_domnode_t) {
	me__ := (*RenderProcessHandler)(self)
	proxy__ := lookupRenderProcessHandlerProxy(me__.Base())
	proxy__.OnFocusedNodeChanged(me__, (*Browser)(browser), (*Frame)(frame), (*Domnode)(node))
}

// OnProcessMessageReceived (on_process_message_received)
// Called when a new message is received from a different process. Return true
// (1) if the message was handled or false (0) otherwise. Do not keep a
// reference to or attempt to access the message outside of this callback.
func (d *RenderProcessHandler) OnProcessMessageReceived(browser *Browser, source_process ProcessID, message *ProcessMessage) int32 {
	return lookupRenderProcessHandlerProxy(d.Base()).OnProcessMessageReceived(d, browser, source_process, message)
}

//export gocef_render_process_handler_on_process_message_received
func gocef_render_process_handler_on_process_message_received(self *C.cef_render_process_handler_t, browser *C.cef_browser_t, source_process C.cef_process_id_t, message *C.cef_process_message_t) C.int {
	me__ := (*RenderProcessHandler)(self)
	proxy__ := lookupRenderProcessHandlerProxy(me__.Base())
	return C.int(proxy__.OnProcessMessageReceived(me__, (*Browser)(browser), ProcessID(source_process), (*ProcessMessage)(message)))
}
