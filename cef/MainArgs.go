package cef

import (
	// #cgo pkg-config: cef
	// #include <stdlib.h>
	// #include "include/internal/cef_types.h"
	// #cgo darwin LDFLAGS: -framework Cocoa
	// #cgo windows LDFLAGS: -Wl,--subsystem,windows
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
