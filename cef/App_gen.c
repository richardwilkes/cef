// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Code created from "callback.c.tmpl" - don't edit by hand

#include "App_gen.h"
#include "_cgo_export.h"

void gocef_set_app_proxy(cef_app_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_before_command_line_processing = (void *)&gocef_app_on_before_command_line_processing;
	self->on_register_custom_schemes = (void *)&gocef_app_on_register_custom_schemes;
	self->get_resource_bundle_handler = (void *)&gocef_app_get_resource_bundle_handler;
	self->get_browser_process_handler = (void *)&gocef_app_get_browser_process_handler;
	self->get_render_process_handler = (void *)&gocef_app_get_render_process_handler;
}
