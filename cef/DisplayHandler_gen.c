// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Code created from "callback.c.tmpl" - don't edit by hand

#include "DisplayHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_display_handler_proxy(cef_display_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_address_change = (void *)&gocef_display_handler_on_address_change;
	self->on_title_change = (void *)&gocef_display_handler_on_title_change;
	self->on_favicon_urlchange = (void *)&gocef_display_handler_on_favicon_urlchange;
	self->on_fullscreen_mode_change = (void *)&gocef_display_handler_on_fullscreen_mode_change;
	self->on_tooltip = (void *)&gocef_display_handler_on_tooltip;
	self->on_status_message = (void *)&gocef_display_handler_on_status_message;
	self->on_console_message = (void *)&gocef_display_handler_on_console_message;
	self->on_auto_resize = (void *)&gocef_display_handler_on_auto_resize;
	self->on_loading_progress_change = (void *)&gocef_display_handler_on_loading_progress_change;
}
