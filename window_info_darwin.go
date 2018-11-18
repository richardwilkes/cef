package cef

import (
	// #include "include/internal/cef_types.h"
	"C"
	"unsafe"
)

func (w *WindowInfo) platformInit(parent WindowHandle) {
	w.native.parent_view = unsafe.Pointer(parent)
}
