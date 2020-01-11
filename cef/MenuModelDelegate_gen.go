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
	// #include "MenuModelDelegate_gen.h"
	"C"
)

// MenuModelDelegateProxy defines methods required for using MenuModelDelegate.
type MenuModelDelegateProxy interface {
	ExecuteCommand(self *MenuModelDelegate, menuModel *MenuModel, commandID int32, eventFlags EventFlags)
	MouseOutsideMenu(self *MenuModelDelegate, menuModel *MenuModel, screenPoint *Point)
	UnhandledOpenSubmenu(self *MenuModelDelegate, menuModel *MenuModel, isRtl int32)
	UnhandledCloseSubmenu(self *MenuModelDelegate, menuModel *MenuModel, isRtl int32)
	MenuWillShow(self *MenuModelDelegate, menuModel *MenuModel)
	MenuClosed(self *MenuModelDelegate, menuModel *MenuModel)
	FormatLabel(self *MenuModelDelegate, menuModel *MenuModel, label *string) int32
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
func (d *MenuModelDelegate) ExecuteCommand(menuModel *MenuModel, commandID int32, eventFlags EventFlags) {
	lookupMenuModelDelegateProxy(d.Base()).ExecuteCommand(d, menuModel, commandID, eventFlags)
}

//nolint:gocritic
//export gocef_menu_model_delegate_execute_command
func gocef_menu_model_delegate_execute_command(self *C.cef_menu_model_delegate_t, menuModel *C.cef_menu_model_t, commandID C.int, eventFlags C.cef_event_flags_t) {
	me__ := (*MenuModelDelegate)(self)
	proxy__ := lookupMenuModelDelegateProxy(me__.Base())
	proxy__.ExecuteCommand(me__, (*MenuModel)(menuModel), int32(commandID), EventFlags(eventFlags))
}

// MouseOutsideMenu (mouse_outside_menu)
// Called when the user moves the mouse outside the menu and over the owning
// window.
func (d *MenuModelDelegate) MouseOutsideMenu(menuModel *MenuModel, screenPoint *Point) {
	lookupMenuModelDelegateProxy(d.Base()).MouseOutsideMenu(d, menuModel, screenPoint)
}

//nolint:gocritic
//export gocef_menu_model_delegate_mouse_outside_menu
func gocef_menu_model_delegate_mouse_outside_menu(self *C.cef_menu_model_delegate_t, menuModel *C.cef_menu_model_t, screenPoint *C.cef_point_t) {
	me__ := (*MenuModelDelegate)(self)
	proxy__ := lookupMenuModelDelegateProxy(me__.Base())
	screenPoint_ := screenPoint.toGo()
	proxy__.MouseOutsideMenu(me__, (*MenuModel)(menuModel), screenPoint_)
}

// UnhandledOpenSubmenu (unhandled_open_submenu)
// Called on unhandled open submenu keyboard commands. |is_rtl| will be true
// (1) if the menu is displaying a right-to-left language.
func (d *MenuModelDelegate) UnhandledOpenSubmenu(menuModel *MenuModel, isRtl int32) {
	lookupMenuModelDelegateProxy(d.Base()).UnhandledOpenSubmenu(d, menuModel, isRtl)
}

//nolint:gocritic
//export gocef_menu_model_delegate_unhandled_open_submenu
func gocef_menu_model_delegate_unhandled_open_submenu(self *C.cef_menu_model_delegate_t, menuModel *C.cef_menu_model_t, isRtl C.int) {
	me__ := (*MenuModelDelegate)(self)
	proxy__ := lookupMenuModelDelegateProxy(me__.Base())
	proxy__.UnhandledOpenSubmenu(me__, (*MenuModel)(menuModel), int32(isRtl))
}

// UnhandledCloseSubmenu (unhandled_close_submenu)
// Called on unhandled close submenu keyboard commands. |is_rtl| will be true
// (1) if the menu is displaying a right-to-left language.
func (d *MenuModelDelegate) UnhandledCloseSubmenu(menuModel *MenuModel, isRtl int32) {
	lookupMenuModelDelegateProxy(d.Base()).UnhandledCloseSubmenu(d, menuModel, isRtl)
}

//nolint:gocritic
//export gocef_menu_model_delegate_unhandled_close_submenu
func gocef_menu_model_delegate_unhandled_close_submenu(self *C.cef_menu_model_delegate_t, menuModel *C.cef_menu_model_t, isRtl C.int) {
	me__ := (*MenuModelDelegate)(self)
	proxy__ := lookupMenuModelDelegateProxy(me__.Base())
	proxy__.UnhandledCloseSubmenu(me__, (*MenuModel)(menuModel), int32(isRtl))
}

// MenuWillShow (menu_will_show)
// The menu is about to show.
func (d *MenuModelDelegate) MenuWillShow(menuModel *MenuModel) {
	lookupMenuModelDelegateProxy(d.Base()).MenuWillShow(d, menuModel)
}

//nolint:gocritic
//export gocef_menu_model_delegate_menu_will_show
func gocef_menu_model_delegate_menu_will_show(self *C.cef_menu_model_delegate_t, menuModel *C.cef_menu_model_t) {
	me__ := (*MenuModelDelegate)(self)
	proxy__ := lookupMenuModelDelegateProxy(me__.Base())
	proxy__.MenuWillShow(me__, (*MenuModel)(menuModel))
}

// MenuClosed (menu_closed)
// The menu has closed.
func (d *MenuModelDelegate) MenuClosed(menuModel *MenuModel) {
	lookupMenuModelDelegateProxy(d.Base()).MenuClosed(d, menuModel)
}

//nolint:gocritic
//export gocef_menu_model_delegate_menu_closed
func gocef_menu_model_delegate_menu_closed(self *C.cef_menu_model_delegate_t, menuModel *C.cef_menu_model_t) {
	me__ := (*MenuModelDelegate)(self)
	proxy__ := lookupMenuModelDelegateProxy(me__.Base())
	proxy__.MenuClosed(me__, (*MenuModel)(menuModel))
}

// FormatLabel (format_label)
// Optionally modify a menu item label. Return true (1) if |label| was
// modified.
func (d *MenuModelDelegate) FormatLabel(menuModel *MenuModel, label *string) int32 {
	return lookupMenuModelDelegateProxy(d.Base()).FormatLabel(d, menuModel, label)
}

//nolint:gocritic
//export gocef_menu_model_delegate_format_label
func gocef_menu_model_delegate_format_label(self *C.cef_menu_model_delegate_t, menuModel *C.cef_menu_model_t, label *C.cef_string_t) C.int {
	me__ := (*MenuModelDelegate)(self)
	proxy__ := lookupMenuModelDelegateProxy(me__.Base())
	label_ := cefstrToString(label)
	return C.int(proxy__.FormatLabel(me__, (*MenuModel)(menuModel), &label_))
}
