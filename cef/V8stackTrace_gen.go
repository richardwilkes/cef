// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// int gocef_v8stack_trace_is_valid(cef_v8stack_trace_t * self, int (CEF_CALLBACK *callback__)(cef_v8stack_trace_t *)) { return callback__(self); }
	// int gocef_v8stack_trace_get_frame_count(cef_v8stack_trace_t * self, int (CEF_CALLBACK *callback__)(cef_v8stack_trace_t *)) { return callback__(self); }
	// cef_v8stack_frame_t * gocef_v8stack_trace_get_frame(cef_v8stack_trace_t * self, int index, cef_v8stack_frame_t * (CEF_CALLBACK *callback__)(cef_v8stack_trace_t *, int)) { return callback__(self, index); }
	"C"
)

// V8stackTrace (cef_v8stack_trace_t from include/capi/cef_v8_capi.h)
// Structure representing a V8 stack trace handle. V8 handles can only be
// accessed from the thread on which they are created. Valid threads for
// creating a V8 handle include the render process main thread (TID_RENDERER)
// and WebWorker threads. A task runner for posting tasks on the associated
// thread can be retrieved via the cef_v8context_t::get_task_runner() function.
type V8stackTrace C.cef_v8stack_trace_t

func (d *V8stackTrace) toNative() *C.cef_v8stack_trace_t {
	return (*C.cef_v8stack_trace_t)(d)
}

// Base (base)
// Base structure.
func (d *V8stackTrace) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// IsValid (is_valid)
// Returns true (1) if the underlying handle is valid and it can be accessed
// on the current thread. Do not call any other functions if this function
// returns false (0).
func (d *V8stackTrace) IsValid() int32 {
	return int32(C.gocef_v8stack_trace_is_valid(d.toNative(), d.is_valid))
}

// GetFrameCount (get_frame_count)
// Returns the number of stack frames.
func (d *V8stackTrace) GetFrameCount() int32 {
	return int32(C.gocef_v8stack_trace_get_frame_count(d.toNative(), d.get_frame_count))
}

// GetFrame (get_frame)
// Returns the stack frame at the specified 0-based index.
func (d *V8stackTrace) GetFrame(index int32) *V8stackFrame {
	return (*V8stackFrame)(C.gocef_v8stack_trace_get_frame(d.toNative(), C.int(index), d.get_frame))
}
