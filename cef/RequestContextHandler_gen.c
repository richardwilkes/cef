// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Code created from "callback.c.tmpl" - don't edit by hand

#include "RequestContextHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_request_context_handler_proxy(cef_request_context_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_request_context_initialized = (void *)&gocef_request_context_handler_on_request_context_initialized;
	self->on_before_plugin_load = (void *)&gocef_request_context_handler_on_before_plugin_load;
	self->get_resource_request_handler = (void *)&gocef_request_context_handler_get_resource_request_handler;
}
