// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Code created from "callback.c.tmpl" - don't edit by hand

#include "ReadHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_read_handler_proxy(cef_read_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->read = (void *)&gocef_read_handler_read;
	self->seek = (void *)&gocef_read_handler_seek;
	self->tell = (void *)&gocef_read_handler_tell;
	self->eof = (void *)&gocef_read_handler_eof;
	self->may_block = (void *)&gocef_read_handler_may_block;
}
