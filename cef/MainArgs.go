// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package cef

import (
	// #cgo darwin CFLAGS: -I/usr/local/cef
	// #cgo darwin LDFLAGS: -framework Cocoa -F/usr/local/cef/Release -framework "Chromium Embedded Framework"
	// #cgo linux CFLAGS: -I/usr/local/cef
	// #cgo linux LDFLAGS: -L/usr/local/cef/Release -lcef
	// #cgo windows pkg-config: cef
	// #cgo windows LDFLAGS: -Wl,--subsystem,windows
	// #include <stdlib.h>
	// #include "include/internal/cef_types.h"
	"C"
)

// MainArgs (cef_main_args_t)
// Structure representing CefExecuteProcess arguments.
type MainArgs struct {
	native C.cef_main_args_t
}

// NewMainArgs creates a new MainArgs value filled with the command line
// arguments.
func NewMainArgs() *MainArgs {
	args := &MainArgs{}
	args.platformInit()
	return args
}

func (d *MainArgs) toNative(native *C.cef_main_args_t) *C.cef_main_args_t {
	*native = d.native
	return native
}
