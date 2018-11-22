package cef

import (
	// #include <stdlib.h>
	// #include "include/internal/cef_types.h"
	"C"
	"syscall"
	"unsafe"
)

func (d *MainArgs) platformInit() {
	proc := syscall.NewLazyDLL("kernel32.dll").NewProc("GetModuleHandleW")
	h, _, err := proc.Call(0)
	if h != 0 {
		d.native.instance = C.HINSTANCE(unsafe.Pointer(h))
	}
}
