package cef

import (
	// #include <stdlib.h>
	// #include "include/internal/cef_types.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/xmath/geom"
)

// NewWindowInfo creates a new default WindowInfo instance.
func NewWindowInfo(parent unsafe.Pointer, bounds geom.Rect) *WindowInfo {
	d := &WindowInfo{
		X:      int32(bounds.X),
		Y:      int32(bounds.Y),
		Width:  int32(bounds.Width),
		Height: int32(bounds.Height),
	}
	d.platformInit(parent)
	return d
}
