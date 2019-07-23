// Code created from "struct.go.tmpl" - don't edit by hand

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
	if d == nil {
		return nil
	}
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

func (n *C.cef_popup_features_t) toGo() *PopupFeatures {
	if n == nil {
		return nil
	}
	var d PopupFeatures
	n.intoGo(&d)
	return &d
}

func (n *C.cef_popup_features_t) intoGo(d *PopupFeatures) {
	d.X = int32(n.x)
	d.XSet = int32(n.xSet)
	d.Y = int32(n.y)
	d.YSet = int32(n.ySet)
	d.Width = int32(n.width)
	d.WidthSet = int32(n.widthSet)
	d.Height = int32(n.height)
	d.HeightSet = int32(n.heightSet)
	d.MenuBarVisible = int32(n.menuBarVisible)
	d.StatusBarVisible = int32(n.statusBarVisible)
	d.ToolBarVisible = int32(n.toolBarVisible)
	d.ScrollbarsVisible = int32(n.scrollbarsVisible)
}
