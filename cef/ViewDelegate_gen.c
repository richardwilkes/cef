// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Code created from "callback.c.tmpl" - don't edit by hand

#include "ViewDelegate_gen.h"
#include "_cgo_export.h"

void gocef_set_view_delegate_proxy(cef_view_delegate_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->get_preferred_size = (void *)&gocef_view_delegate_get_preferred_size;
	self->get_minimum_size = (void *)&gocef_view_delegate_get_minimum_size;
	self->get_maximum_size = (void *)&gocef_view_delegate_get_maximum_size;
	self->get_height_for_width = (void *)&gocef_view_delegate_get_height_for_width;
	self->on_parent_view_changed = (void *)&gocef_view_delegate_on_parent_view_changed;
	self->on_child_view_changed = (void *)&gocef_view_delegate_on_child_view_changed;
	self->on_focus = (void *)&gocef_view_delegate_on_focus;
	self->on_blur = (void *)&gocef_view_delegate_on_blur;
}
