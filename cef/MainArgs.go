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
