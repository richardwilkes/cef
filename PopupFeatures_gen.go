// Code generated - DO NOT EDIT.

package cef

import (
	// #include "capi_gen.h"
	"C"
)

// PopupFeatures (cef_popup_features_t from include/internal/cef_types.h)
// Popup window features.
type PopupFeatures struct {
	// X (x)
	X int32
	// XSet (xSet)
	XSet int32
	// Y (y)
	Y int32
	// YSet (ySet)
	YSet int32
	// Width (width)
	Width int32
	// WidthSet (widthSet)
	WidthSet int32
	// Height (height)
	Height int32
	// HeightSet (heightSet)
	HeightSet int32
	// MenuBarVisible (menuBarVisible)
	MenuBarVisible int32
	// StatusBarVisible (statusBarVisible)
	StatusBarVisible int32
	// ToolBarVisible (toolBarVisible)
	ToolBarVisible int32
	// ScrollbarsVisible (scrollbarsVisible)
	ScrollbarsVisible int32
}

// NewPopupFeatures creates a new PopupFeatures.
func NewPopupFeatures() *PopupFeatures {
	return &PopupFeatures{}
}

func (d *PopupFeatures) toNative(native *C.cef_popup_features_t) *C.cef_popup_features_t {
	native.x = C.int(d.X)
	native.xSet = C.int(d.XSet)
	native.y = C.int(d.Y)
	native.ySet = C.int(d.YSet)
	native.width = C.int(d.Width)
	native.widthSet = C.int(d.WidthSet)
	native.height = C.int(d.Height)
	native.heightSet = C.int(d.HeightSet)
	native.menuBarVisible = C.int(d.MenuBarVisible)
	native.statusBarVisible = C.int(d.StatusBarVisible)
	native.toolBarVisible = C.int(d.ToolBarVisible)
	native.scrollbarsVisible = C.int(d.ScrollbarsVisible)
	return native
}

func (d *PopupFeatures) fromNative(native *C.cef_popup_features_t) *PopupFeatures {
	d.X = int32(native.x)
	d.XSet = int32(native.xSet)
	d.Y = int32(native.y)
	d.YSet = int32(native.ySet)
	d.Width = int32(native.width)
	d.WidthSet = int32(native.widthSet)
	d.Height = int32(native.height)
	d.HeightSet = int32(native.heightSet)
	d.MenuBarVisible = int32(native.menuBarVisible)
	d.StatusBarVisible = int32(native.statusBarVisible)
	d.ToolBarVisible = int32(native.toolBarVisible)
	d.ScrollbarsVisible = int32(native.scrollbarsVisible)
	return d
}
