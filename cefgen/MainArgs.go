package cef

import (
	// #include <stdlib.h>
	// #include "include/internal/cef_types.h"
	"C"
)

// MainArgs (cef_main_args_t)
// Structure representing CefExecuteProcess arguments.
type MainArgs struct {
	native C.cef_main_args_t
}

func newMainArgs() *MainArgs {
	args := &MainArgs{}
	args.platformInit()
	return args
}

func (d *MainArgs) toNative(native *C.cef_main_args_t) {
	*native = d.native
}

func (d *MainArgs) fromNative(native *C.cef_main_args_t) {
	d.native = *native
}
