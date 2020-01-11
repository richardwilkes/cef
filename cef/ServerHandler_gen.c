// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Code created from "callback.c.tmpl" - don't edit by hand

#include "ServerHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_server_handler_proxy(cef_server_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_server_created = (void *)&gocef_server_handler_on_server_created;
	self->on_server_destroyed = (void *)&gocef_server_handler_on_server_destroyed;
	self->on_client_connected = (void *)&gocef_server_handler_on_client_connected;
	self->on_client_disconnected = (void *)&gocef_server_handler_on_client_disconnected;
	self->on_http_request = (void *)&gocef_server_handler_on_http_request;
	self->on_web_socket_request = (void *)&gocef_server_handler_on_web_socket_request;
	self->on_web_socket_connected = (void *)&gocef_server_handler_on_web_socket_connected;
	self->on_web_socket_message = (void *)&gocef_server_handler_on_web_socket_message;
}
