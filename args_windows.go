package cef

import (
	// #include <stdlib.h>
	// #include "include/capi/cef_app_capi.h"
	"C"
	"syscall"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
)

func mainArgs() (*C.cef_main_args_t, error) {
	proc := syscall.NewLazyDLL("kernel32.dll").NewProc("GetModuleHandleW")
	h, _, err := proc.Call(0)
	if h == 0 {
		return nil, errs.NewWithCause(proc.Name, err)
	}
	args := (*C.cef_main_args_t)(C.calloc(1, C.sizeof_struct__cef_main_args_t))
	args.instance = C.HINSTANCE(unsafe.Pointer(h))
	return args, nil
}
