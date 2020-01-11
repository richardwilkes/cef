// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Code created from "callback.c.tmpl" - don't edit by hand

#include "RequestHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_request_handler_proxy(cef_request_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_before_browse = (void *)&gocef_request_handler_on_before_browse;
	self->on_open_urlfrom_tab = (void *)&gocef_request_handler_on_open_urlfrom_tab;
	self->get_resource_request_handler = (void *)&gocef_request_handler_get_resource_request_handler;
	self->get_auth_credentials = (void *)&gocef_request_handler_get_auth_credentials;
	self->on_quota_request = (void *)&gocef_request_handler_on_quota_request;
	self->on_certificate_error = (void *)&gocef_request_handler_on_certificate_error;
	self->on_select_client_certificate = (void *)&gocef_request_handler_on_select_client_certificate;
	self->on_plugin_crashed = (void *)&gocef_request_handler_on_plugin_crashed;
	self->on_render_view_ready = (void *)&gocef_request_handler_on_render_view_ready;
	self->on_render_process_terminated = (void *)&gocef_request_handler_on_render_process_terminated;
}
