// Code created from "struct.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	"C"
)

// BoxLayoutSettings (cef_box_layout_settings_t from include/internal/cef_types.h)
// Settings used when initializing a CefBoxLayout.
type BoxLayoutSettings struct {
	// Horizontal (horizontal)
	// If true (1) the layout will be horizontal, otherwise the layout will be
	// vertical.
	Horizontal int32
	// InsideBorderHorizontalSpacing (inside_border_horizontal_spacing)
	// Adds additional horizontal space between the child view area and the host
	// view border.
	InsideBorderHorizontalSpacing int32
	// InsideBorderVerticalSpacing (inside_border_vertical_spacing)
	// Adds additional vertical space between the child view area and the host
	// view border.
	InsideBorderVerticalSpacing int32
	// InsideBorderInsets (inside_border_insets)
	// Adds additional space around the child view area.
	InsideBorderInsets Insets
	// BetweenChildSpacing (between_child_spacing)
	// Adds additional space between child views.
	BetweenChildSpacing int32
	// MainAxisAlignment (main_axis_alignment)
	// Specifies where along the main axis the child views should be laid out.
	MainAxisAlignment MainAxisAlignment
	// CrossAxisAlignment (cross_axis_alignment)
	// Specifies where along the cross axis the child views should be laid out.
	CrossAxisAlignment CrossAxisAlignment
	// MinimumCrossAxisSize (minimum_cross_axis_size)
	// Minimum cross axis size.
	MinimumCrossAxisSize int32
	// DefaultFlex (default_flex)
	// Default flex for views when none is specified via CefBoxLayout methods.
	// Using the preferred size as the basis, free space along the main axis is
	// distributed to views in the ratio of their flex weights. Similarly, if the
	// views will overflow the parent, space is subtracted in these ratios. A flex
	// of 0 means this view is not resized. Flex values must not be negative.
	DefaultFlex int32
}

// NewBoxLayoutSettings creates a new BoxLayoutSettings.
func NewBoxLayoutSettings() *BoxLayoutSettings {
	return &BoxLayoutSettings{}
}

func (d *BoxLayoutSettings) toNative(native *C.cef_box_layout_settings_t) *C.cef_box_layout_settings_t {
	if d == nil {
		return nil
	}
	native.horizontal = C.int(d.Horizontal)
	native.inside_border_horizontal_spacing = C.int(d.InsideBorderHorizontalSpacing)
	native.inside_border_vertical_spacing = C.int(d.InsideBorderVerticalSpacing)
	d.InsideBorderInsets.toNative(&native.inside_border_insets)
	native.between_child_spacing = C.int(d.BetweenChildSpacing)
	native.main_axis_alignment = C.cef_main_axis_alignment_t(d.MainAxisAlignment)
	native.cross_axis_alignment = C.cef_cross_axis_alignment_t(d.CrossAxisAlignment)
	native.minimum_cross_axis_size = C.int(d.MinimumCrossAxisSize)
	native.default_flex = C.int(d.DefaultFlex)
	return native
}

func (n *C.cef_box_layout_settings_t) toGo() *BoxLayoutSettings {
	if n == nil {
		return nil
	}
	var d BoxLayoutSettings
	n.intoGo(&d)
	return &d
}

func (n *C.cef_box_layout_settings_t) intoGo(d *BoxLayoutSettings) {
	d.Horizontal = int32(n.horizontal)
	d.InsideBorderHorizontalSpacing = int32(n.inside_border_horizontal_spacing)
	d.InsideBorderVerticalSpacing = int32(n.inside_border_vertical_spacing)
	n.inside_border_insets.intoGo(&d.InsideBorderInsets)
	d.BetweenChildSpacing = int32(n.between_child_spacing)
	d.MainAxisAlignment = MainAxisAlignment(n.main_axis_alignment)
	d.CrossAxisAlignment = CrossAxisAlignment(n.cross_axis_alignment)
	d.MinimumCrossAxisSize = int32(n.minimum_cross_axis_size)
	d.DefaultFlex = int32(n.default_flex)
}
