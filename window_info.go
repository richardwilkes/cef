package cef

import (
	// #include <stdlib.h>
	// #include "include/internal/cef_types.h"
	"C"

	"github.com/richardwilkes/toolbox/xmath/geom"
)

type WindowInfo struct {
	native *C.cef_window_info_t
}

// NewWindowInfo creates a new default WindowInfo instance.
func NewWindowInfo(parent WindowHandle, bounds geom.Rect) *WindowInfo {
	info := &WindowInfo{native: (*C.cef_window_info_t)(C.calloc(1, C.sizeof_struct__cef_window_info_t))}
	info.native.x = C.int(bounds.X)
	info.native.y = C.int(bounds.Y)
	info.native.width = C.int(bounds.Width)
	info.native.height = C.int(bounds.Height)
	info.platformInit(parent)
	return info
}
