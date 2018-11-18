// +build !windows

package cef

import (
	// #include <stdlib.h>
	// #include "include/capi/cef_app_capi.h"
	"C"
	"os"
	"unsafe"
)

func mainArgs() (*C.cef_main_args_t, error) {
	args := (*C.cef_main_args_t)(C.calloc(1, C.sizeof_struct__cef_main_args_t))
	if len(os.Args) > 0 {
		cp := C.calloc(C.size_t(len(os.Args)), C.size_t(unsafe.Sizeof(uintptr(0))))
		p := (*[1<<30 - 1]*C.char)(cp)
		for i, arg := range os.Args {
			p[i] = C.CString(arg)
		}
		args.argc = C.int(len(os.Args))
		args.argv = (**C.char)(cp)
	}
	return args, nil
}
