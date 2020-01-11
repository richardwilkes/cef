// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// void gocef_menu_button_show_menu(cef_menu_button_t * self, cef_menu_model_t * menuModel, cef_point_t * screenPoint, cef_menu_anchor_position_t anchorPosition, void (CEF_CALLBACK *callback__)(cef_menu_button_t *, cef_menu_model_t *, cef_point_t *, cef_menu_anchor_position_t)) { return callback__(self, menuModel, screenPoint, anchorPosition); }
	// void gocef_menu_button_trigger_menu(cef_menu_button_t * self, void (CEF_CALLBACK *callback__)(cef_menu_button_t *)) { return callback__(self); }
	"C"
)

// MenuButton (cef_menu_button_t from include/capi/views/cef_menu_button_capi.h)
// MenuButton is a button with optional text, icon and/or menu marker that shows
// a menu when clicked with the left mouse button. All size and position values
// are in density independent pixels (DIP) unless otherwise indicated. Methods
// must be called on the browser process UI thread unless otherwise indicated.
type MenuButton C.cef_menu_button_t

func (d *MenuButton) toNative() *C.cef_menu_button_t {
	return (*C.cef_menu_button_t)(d)
}

// Base (base)
// Base structure.
func (d *MenuButton) Base() *LabelButton {
	return (*LabelButton)(&d.base)
}

// ShowMenu (show_menu)
// Show a menu with contents |menu_model|. |screen_point| specifies the menu
// position in screen coordinates. |anchor_position| specifies how the menu
// will be anchored relative to |screen_point|. This function should be called
// from cef_menu_button_delegate_t::on_menu_button_pressed().
func (d *MenuButton) ShowMenu(menuModel *MenuModel, screenPoint *Point, anchorPosition MenuAnchorPosition) {
	C.gocef_menu_button_show_menu(d.toNative(), menuModel.toNative(), screenPoint.toNative(&C.cef_point_t{}), C.cef_menu_anchor_position_t(anchorPosition), d.show_menu)
}

// TriggerMenu (trigger_menu)
// Show the menu for this button. Results in a call to
// cef_menu_button_delegate_t::on_menu_button_pressed().
func (d *MenuButton) TriggerMenu() {
	C.gocef_menu_button_trigger_menu(d.toNative(), d.trigger_menu)
}
