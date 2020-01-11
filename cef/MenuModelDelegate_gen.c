// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Code created from "callback.c.tmpl" - don't edit by hand

#include "MenuModelDelegate_gen.h"
#include "_cgo_export.h"

void gocef_set_menu_model_delegate_proxy(cef_menu_model_delegate_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->execute_command = (void *)&gocef_menu_model_delegate_execute_command;
	self->mouse_outside_menu = (void *)&gocef_menu_model_delegate_mouse_outside_menu;
	self->unhandled_open_submenu = (void *)&gocef_menu_model_delegate_unhandled_open_submenu;
	self->unhandled_close_submenu = (void *)&gocef_menu_model_delegate_unhandled_close_submenu;
	self->menu_will_show = (void *)&gocef_menu_model_delegate_menu_will_show;
	self->menu_closed = (void *)&gocef_menu_model_delegate_menu_closed;
	self->format_label = (void *)&gocef_menu_model_delegate_format_label;
}
