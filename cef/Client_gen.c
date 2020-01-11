// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Code created from "callback.c.tmpl" - don't edit by hand

#include "Client_gen.h"
#include "_cgo_export.h"

void gocef_set_client_proxy(cef_client_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->get_audio_handler = (void *)&gocef_client_get_audio_handler;
	self->get_context_menu_handler = (void *)&gocef_client_get_context_menu_handler;
	self->get_dialog_handler = (void *)&gocef_client_get_dialog_handler;
	self->get_display_handler = (void *)&gocef_client_get_display_handler;
	self->get_download_handler = (void *)&gocef_client_get_download_handler;
	self->get_drag_handler = (void *)&gocef_client_get_drag_handler;
	self->get_find_handler = (void *)&gocef_client_get_find_handler;
	self->get_focus_handler = (void *)&gocef_client_get_focus_handler;
	self->get_jsdialog_handler = (void *)&gocef_client_get_jsdialog_handler;
	self->get_keyboard_handler = (void *)&gocef_client_get_keyboard_handler;
	self->get_life_span_handler = (void *)&gocef_client_get_life_span_handler;
	self->get_load_handler = (void *)&gocef_client_get_load_handler;
	self->get_render_handler = (void *)&gocef_client_get_render_handler;
	self->get_request_handler = (void *)&gocef_client_get_request_handler;
	self->on_process_message_received = (void *)&gocef_client_on_process_message_received;
}
