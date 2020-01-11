// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Code created from "callback.c.tmpl" - don't edit by hand

#include "ExtensionHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_extension_handler_proxy(cef_extension_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_extension_load_failed = (void *)&gocef_extension_handler_on_extension_load_failed;
	self->on_extension_loaded = (void *)&gocef_extension_handler_on_extension_loaded;
	self->on_extension_unloaded = (void *)&gocef_extension_handler_on_extension_unloaded;
	self->on_before_background_browser = (void *)&gocef_extension_handler_on_before_background_browser;
	self->on_before_browser = (void *)&gocef_extension_handler_on_before_browser;
	self->get_active_browser = (void *)&gocef_extension_handler_get_active_browser;
	self->can_access_browser = (void *)&gocef_extension_handler_can_access_browser;
	self->get_extension_resource = (void *)&gocef_extension_handler_get_extension_resource;
}
