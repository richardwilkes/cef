package cef

import (
	// #include <stdlib.h>
	// #include "include/internal/cef_types.h"
	// #cgo CFLAGS: -I ${SRCDIR}/cef
	// #cgo darwin LDFLAGS: -framework Cocoa -F ${SRCDIR}/cef/lib -framework "Chromium Embedded Framework"
	// #cgo windows LDFLAGS: -L${SRCDIR}/cef/lib -lcef -Wl,--subsystem,windows
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

func (d *MainArgs) fromNative(native *C.cef_main_args_t) {
	d.native = *native
}
