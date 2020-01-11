// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Code created from "callback.c.tmpl" - don't edit by hand

#include "WriteHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_write_handler_proxy(cef_write_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->write = (void *)&gocef_write_handler_write;
	self->seek = (void *)&gocef_write_handler_seek;
	self->tell = (void *)&gocef_write_handler_tell;
	self->flush = (void *)&gocef_write_handler_flush;
	self->may_block = (void *)&gocef_write_handler_may_block;
}
