// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Code created from "callback.c.tmpl" - don't edit by hand

#include "ResourceRequestHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_resource_request_handler_proxy(cef_resource_request_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->get_cookie_access_filter = (void *)&gocef_resource_request_handler_get_cookie_access_filter;
	self->on_before_resource_load = (void *)&gocef_resource_request_handler_on_before_resource_load;
	self->get_resource_handler = (void *)&gocef_resource_request_handler_get_resource_handler;
	self->on_resource_redirect = (void *)&gocef_resource_request_handler_on_resource_redirect;
	self->on_resource_response = (void *)&gocef_resource_request_handler_on_resource_response;
	self->get_resource_response_filter = (void *)&gocef_resource_request_handler_get_resource_response_filter;
	self->on_resource_load_complete = (void *)&gocef_resource_request_handler_on_resource_load_complete;
	self->on_protocol_execution = (void *)&gocef_resource_request_handler_on_protocol_execution;
}
