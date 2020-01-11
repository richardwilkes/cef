// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Code created from "callback.c.tmpl" - don't edit by hand

#include "RenderProcessHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_render_process_handler_proxy(cef_render_process_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_render_thread_created = (void *)&gocef_render_process_handler_on_render_thread_created;
	self->on_web_kit_initialized = (void *)&gocef_render_process_handler_on_web_kit_initialized;
	self->on_browser_created = (void *)&gocef_render_process_handler_on_browser_created;
	self->on_browser_destroyed = (void *)&gocef_render_process_handler_on_browser_destroyed;
	self->get_load_handler = (void *)&gocef_render_process_handler_get_load_handler;
	self->on_context_created = (void *)&gocef_render_process_handler_on_context_created;
	self->on_context_released = (void *)&gocef_render_process_handler_on_context_released;
	self->on_uncaught_exception = (void *)&gocef_render_process_handler_on_uncaught_exception;
	self->on_focused_node_changed = (void *)&gocef_render_process_handler_on_focused_node_changed;
	self->on_process_message_received = (void *)&gocef_render_process_handler_on_process_message_received;
}
