// Code generated - DO NOT EDIT.

package cef

import (
	// #include "ContextMenuHandler_gen.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

// ContextMenuHandlerProxy defines methods required for using ContextMenuHandler.
type ContextMenuHandlerProxy interface {
	OnBeforeContextMenu(self *ContextMenuHandler, browser *Browser, frame *Frame, params *ContextMenuParams, model *MenuModel)
	RunContextMenu(self *ContextMenuHandler, browser *Browser, frame *Frame, params *ContextMenuParams, model *MenuModel, callback *RunContextMenuCallback) int32
	OnContextMenuCommand(self *ContextMenuHandler, browser *Browser, frame *Frame, params *ContextMenuParams, command_id int32, event_flags EventFlags) int32
	OnContextMenuDismissed(self *ContextMenuHandler, browser *Browser, frame *Frame)
}

// ContextMenuHandler (cef_context_menu_handler_t from include/capi/cef_context_menu_handler_capi.h)
// Implement this structure to handle context menu events. The functions of this
// structure will be called on the UI thread.
type ContextMenuHandler C.cef_context_menu_handler_t

// NewContextMenuHandler creates a new ContextMenuHandler with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewContextMenuHandler(proxy ContextMenuHandlerProxy) *ContextMenuHandler {
	result := (*ContextMenuHandler)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_context_menu_handler_t, proxy)))
	if proxy != nil {
		C.gocef_set_context_menu_handler_proxy(result.toNative())
	}
	return result
}

func (d *ContextMenuHandler) toNative() *C.cef_context_menu_handler_t {
	return (*C.cef_context_menu_handler_t)(d)
}

func lookupContextMenuHandlerProxy(obj *BaseRefCounted) ContextMenuHandlerProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(ContextMenuHandlerProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type ContextMenuHandlerProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *ContextMenuHandler) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnBeforeContextMenu (on_before_context_menu)
// Called before a context menu is displayed. |params| provides information
// about the context menu state. |model| initially contains the default
// context menu. The |model| can be cleared to show no context menu or
// modified to show a custom menu. Do not keep references to |params| or
// |model| outside of this callback.
func (d *ContextMenuHandler) OnBeforeContextMenu(browser *Browser, frame *Frame, params *ContextMenuParams, model *MenuModel) {
	lookupContextMenuHandlerProxy(d.Base()).OnBeforeContextMenu(d, browser, frame, params, model)
}

//export gocef_context_menu_handler_on_before_context_menu
func gocef_context_menu_handler_on_before_context_menu(self *C.cef_context_menu_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, params *C.cef_context_menu_params_t, model *C.cef_menu_model_t) {
	me__ := (*ContextMenuHandler)(self)
	proxy__ := lookupContextMenuHandlerProxy(me__.Base())
	proxy__.OnBeforeContextMenu(me__, (*Browser)(browser), (*Frame)(frame), (*ContextMenuParams)(params), (*MenuModel)(model))
}

// RunContextMenu (run_context_menu)
// Called to allow custom display of the context menu. |params| provides
// information about the context menu state. |model| contains the context menu
// model resulting from OnBeforeContextMenu. For custom display return true
// (1) and execute |callback| either synchronously or asynchronously with the
// selected command ID. For default display return false (0). Do not keep
// references to |params| or |model| outside of this callback.
func (d *ContextMenuHandler) RunContextMenu(browser *Browser, frame *Frame, params *ContextMenuParams, model *MenuModel, callback *RunContextMenuCallback) int32 {
	return lookupContextMenuHandlerProxy(d.Base()).RunContextMenu(d, browser, frame, params, model, callback)
}

//export gocef_context_menu_handler_run_context_menu
func gocef_context_menu_handler_run_context_menu(self *C.cef_context_menu_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, params *C.cef_context_menu_params_t, model *C.cef_menu_model_t, callback *C.cef_run_context_menu_callback_t) C.int {
	me__ := (*ContextMenuHandler)(self)
	proxy__ := lookupContextMenuHandlerProxy(me__.Base())
	return C.int(proxy__.RunContextMenu(me__, (*Browser)(browser), (*Frame)(frame), (*ContextMenuParams)(params), (*MenuModel)(model), (*RunContextMenuCallback)(callback)))
}

// OnContextMenuCommand (on_context_menu_command)
// Called to execute a command selected from the context menu. Return true (1)
// if the command was handled or false (0) for the default implementation. See
// cef_menu_id_t for the command ids that have default implementations. All
// user-defined command ids should be between MENU_ID_USER_FIRST and
// MENU_ID_USER_LAST. |params| will have the same values as what was passed to
// on_before_context_menu(). Do not keep a reference to |params| outside of
// this callback.
func (d *ContextMenuHandler) OnContextMenuCommand(browser *Browser, frame *Frame, params *ContextMenuParams, command_id int32, event_flags EventFlags) int32 {
	return lookupContextMenuHandlerProxy(d.Base()).OnContextMenuCommand(d, browser, frame, params, command_id, event_flags)
}

//export gocef_context_menu_handler_on_context_menu_command
func gocef_context_menu_handler_on_context_menu_command(self *C.cef_context_menu_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, params *C.cef_context_menu_params_t, command_id C.int, event_flags C.cef_event_flags_t) C.int {
	me__ := (*ContextMenuHandler)(self)
	proxy__ := lookupContextMenuHandlerProxy(me__.Base())
	return C.int(proxy__.OnContextMenuCommand(me__, (*Browser)(browser), (*Frame)(frame), (*ContextMenuParams)(params), int32(command_id), EventFlags(event_flags)))
}

// OnContextMenuDismissed (on_context_menu_dismissed)
// Called when the context menu is dismissed irregardless of whether the menu
// was NULL or a command was selected.
func (d *ContextMenuHandler) OnContextMenuDismissed(browser *Browser, frame *Frame) {
	lookupContextMenuHandlerProxy(d.Base()).OnContextMenuDismissed(d, browser, frame)
}

//export gocef_context_menu_handler_on_context_menu_dismissed
func gocef_context_menu_handler_on_context_menu_dismissed(self *C.cef_context_menu_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t) {
	me__ := (*ContextMenuHandler)(self)
	proxy__ := lookupContextMenuHandlerProxy(me__.Base())
	proxy__.OnContextMenuDismissed(me__, (*Browser)(browser), (*Frame)(frame))
}
