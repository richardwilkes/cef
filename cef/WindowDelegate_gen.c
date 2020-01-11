// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Code created from "callback.c.tmpl" - don't edit by hand

#include "WindowDelegate_gen.h"
#include "_cgo_export.h"

void gocef_set_window_delegate_proxy(cef_window_delegate_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_window_created = (void *)&gocef_window_delegate_on_window_created;
	self->on_window_destroyed = (void *)&gocef_window_delegate_on_window_destroyed;
	self->get_parent_window = (void *)&gocef_window_delegate_get_parent_window;
	self->is_frameless = (void *)&gocef_window_delegate_is_frameless;
	self->can_resize = (void *)&gocef_window_delegate_can_resize;
	self->can_maximize = (void *)&gocef_window_delegate_can_maximize;
	self->can_minimize = (void *)&gocef_window_delegate_can_minimize;
	self->can_close = (void *)&gocef_window_delegate_can_close;
	self->on_accelerator = (void *)&gocef_window_delegate_on_accelerator;
	self->on_key_event = (void *)&gocef_window_delegate_on_key_event;
}
