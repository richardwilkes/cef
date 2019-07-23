// +build !windows

package cef

import (
	"os"
	"unsafe"
)

import (
	// #include <stdlib.h>
	// #include "include/internal/cef_types.h"
	"C"
)

func (d *MainArgs) platformInit() {
	if len(os.Args) > 0 {
		cp := C.calloc(C.size_t(len(os.Args)), C.size_t(unsafe.Sizeof(uintptr(0))))
		p := (*[1<<30 - 1]*C.char)(cp)
		for i, arg := range os.Args {
			p[i] = C.CString(arg)
		}
		d.native.argc = C.int(len(os.Args))
		d.native.argv = (**C.char)(cp)
	}
}
