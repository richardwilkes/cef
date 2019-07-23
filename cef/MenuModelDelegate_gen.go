// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

import (
	// #include "MenuModelDelegate_gen.h"
	"C"
)

// MenuModelDelegateProxy defines methods required for using MenuModelDelegate.
type MenuModelDelegateProxy interface {
	ExecuteCommand(self *MenuModelDelegate, menu_model *MenuModel, command_id int32, event_flags EventFlags)
	MouseOutsideMenu(self *MenuModelDelegate, menu_model *MenuModel, screen_point *Point)
	UnhandledOpenSubmenu(self *MenuModelDelegate, menu_model *MenuModel, is_rtl int32)
	UnhandledCloseSubmenu(self *MenuModelDelegate, menu_model *MenuModel, is_rtl int32)
	MenuWillShow(self *MenuModelDelegate, menu_model *MenuModel)
	MenuClosed(self *MenuModelDelegate, menu_model *MenuModel)
	FormatLabel(self *MenuModelDelegate, menu_model *MenuModel, label *string) int32
}

// MenuModelDelegate (cef_menu_model_delegate_t from include/capi/cef_menu_model_delegate_capi.h)
// Implement this structure to handle menu model events. The functions of this
// structure will be called on the browser process UI thread unless otherwise
// indicated.
type MenuModelDelegate C.cef_menu_model_delegate_t

// NewMenuModelDelegate creates a new MenuModelDelegate with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewMenuModelDelegate(proxy MenuModelDelegateProxy) *MenuModelDelegate {
	result := (*MenuModelDelegate)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_menu_model_delegate_t, proxy)))
	if proxy != nil {
		C.gocef_set_menu_model_delegate_proxy(result.toNative())
	}
	return result
}

func (d *MenuModelDelegate) toNative() *C.cef_menu_model_delegate_t {
	return (*C.cef_menu_model_delegate_t)(d)
}

func lookupMenuModelDelegateProxy(obj *BaseRefCounted) MenuModelDelegateProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(MenuModelDelegateProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type MenuModelDelegateProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *MenuModelDelegate) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// ExecuteCommand (execute_command)
// Perform the action associated with the specified |command_id| and optional
// |event_flags|.
func (d *MenuModelDelegate) ExecuteCommand(menu_model *MenuModel, command_id int32, event_flags EventFlags) {
	lookupMenuModelDelegateProxy(d.Base()).ExecuteCommand(d, menu_model, command_id, event_flags)
}

//nolint:gocritic
//export gocef_menu_model_delegate_execute_command
func gocef_menu_model_delegate_execute_command(self *C.cef_menu_model_delegate_t, menu_model *C.cef_menu_model_t, command_id C.int, event_flags C.cef_event_flags_t) {
	me__ := (*MenuModelDelegate)(self)
	proxy__ := lookupMenuModelDelegateProxy(me__.Base())
	proxy__.ExecuteCommand(me__, (*MenuModel)(menu_model), int32(command_id), EventFlags(event_flags))
}

// MouseOutsideMenu (mouse_outside_menu)
// Called when the user moves the mouse outside the menu and over the owning
// window.
func (d *MenuModelDelegate) MouseOutsideMenu(menu_model *MenuModel, screen_point *Point) {
	lookupMenuModelDelegateProxy(d.Base()).MouseOutsideMenu(d, menu_model, screen_point)
}

//nolint:gocritic
//export gocef_menu_model_delegate_mouse_outside_menu
func gocef_menu_model_delegate_mouse_outside_menu(self *C.cef_menu_model_delegate_t, menu_model *C.cef_menu_model_t, screen_point *C.cef_point_t) {
	me__ := (*MenuModelDelegate)(self)
	proxy__ := lookupMenuModelDelegateProxy(me__.Base())
	screen_point_ := screen_point.toGo()
	proxy__.MouseOutsideMenu(me__, (*MenuModel)(menu_model), screen_point_)
}

// UnhandledOpenSubmenu (unhandled_open_submenu)
// Called on unhandled open submenu keyboard commands. |is_rtl| will be true
// (1) if the menu is displaying a right-to-left language.
func (d *MenuModelDelegate) UnhandledOpenSubmenu(menu_model *MenuModel, is_rtl int32) {
	lookupMenuModelDelegateProxy(d.Base()).UnhandledOpenSubmenu(d, menu_model, is_rtl)
}

//nolint:gocritic
//export gocef_menu_model_delegate_unhandled_open_submenu
func gocef_menu_model_delegate_unhandled_open_submenu(self *C.cef_menu_model_delegate_t, menu_model *C.cef_menu_model_t, is_rtl C.int) {
	me__ := (*MenuModelDelegate)(self)
	proxy__ := lookupMenuModelDelegateProxy(me__.Base())
	proxy__.UnhandledOpenSubmenu(me__, (*MenuModel)(menu_model), int32(is_rtl))
}

// UnhandledCloseSubmenu (unhandled_close_submenu)
// Called on unhandled close submenu keyboard commands. |is_rtl| will be true
// (1) if the menu is displaying a right-to-left language.
func (d *MenuModelDelegate) UnhandledCloseSubmenu(menu_model *MenuModel, is_rtl int32) {
	lookupMenuModelDelegateProxy(d.Base()).UnhandledCloseSubmenu(d, menu_model, is_rtl)
}

//nolint:gocritic
//export gocef_menu_model_delegate_unhandled_close_submenu
func gocef_menu_model_delegate_unhandled_close_submenu(self *C.cef_menu_model_delegate_t, menu_model *C.cef_menu_model_t, is_rtl C.int) {
	me__ := (*MenuModelDelegate)(self)
	proxy__ := lookupMenuModelDelegateProxy(me__.Base())
	proxy__.UnhandledCloseSubmenu(me__, (*MenuModel)(menu_model), int32(is_rtl))
}

// MenuWillShow (menu_will_show)
// The menu is about to show.
func (d *MenuModelDelegate) MenuWillShow(menu_model *MenuModel) {
	lookupMenuModelDelegateProxy(d.Base()).MenuWillShow(d, menu_model)
}

//nolint:gocritic
//export gocef_menu_model_delegate_menu_will_show
func gocef_menu_model_delegate_menu_will_show(self *C.cef_menu_model_delegate_t, menu_model *C.cef_menu_model_t) {
	me__ := (*MenuModelDelegate)(self)
	proxy__ := lookupMenuModelDelegateProxy(me__.Base())
	proxy__.MenuWillShow(me__, (*MenuModel)(menu_model))
}

// MenuClosed (menu_closed)
// The menu has closed.
func (d *MenuModelDelegate) MenuClosed(menu_model *MenuModel) {
	lookupMenuModelDelegateProxy(d.Base()).MenuClosed(d, menu_model)
}

//nolint:gocritic
//export gocef_menu_model_delegate_menu_closed
func gocef_menu_model_delegate_menu_closed(self *C.cef_menu_model_delegate_t, menu_model *C.cef_menu_model_t) {
	me__ := (*MenuModelDelegate)(self)
	proxy__ := lookupMenuModelDelegateProxy(me__.Base())
	proxy__.MenuClosed(me__, (*MenuModel)(menu_model))
}

// FormatLabel (format_label)
// Optionally modify a menu item label. Return true (1) if |label| was
// modified.
func (d *MenuModelDelegate) FormatLabel(menu_model *MenuModel, label *string) int32 {
	return lookupMenuModelDelegateProxy(d.Base()).FormatLabel(d, menu_model, label)
}

//nolint:gocritic
//export gocef_menu_model_delegate_format_label
func gocef_menu_model_delegate_format_label(self *C.cef_menu_model_delegate_t, menu_model *C.cef_menu_model_t, label *C.cef_string_t) C.int {
	me__ := (*MenuModelDelegate)(self)
	proxy__ := lookupMenuModelDelegateProxy(me__.Base())
	label_ := cefstrToString(label)
	return C.int(proxy__.FormatLabel(me__, (*MenuModel)(menu_model), &label_))
}
