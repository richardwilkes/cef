// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// int gocef_print_settings_is_valid(cef_print_settings_t * self, int (CEF_CALLBACK *callback__)(cef_print_settings_t *)) { return callback__(self); }
	// int gocef_print_settings_is_read_only(cef_print_settings_t * self, int (CEF_CALLBACK *callback__)(cef_print_settings_t *)) { return callback__(self); }
	// cef_print_settings_t * gocef_print_settings_copy(cef_print_settings_t * self, cef_print_settings_t * (CEF_CALLBACK *callback__)(cef_print_settings_t *)) { return callback__(self); }
	// void gocef_print_settings_set_orientation(cef_print_settings_t * self, int landscape, void (CEF_CALLBACK *callback__)(cef_print_settings_t *, int)) { return callback__(self, landscape); }
	// int gocef_print_settings_is_landscape(cef_print_settings_t * self, int (CEF_CALLBACK *callback__)(cef_print_settings_t *)) { return callback__(self); }
	// void gocef_print_settings_set_printer_printable_area(cef_print_settings_t * self, cef_size_t * physicalSizeDeviceUnits, cef_rect_t * printableAreaDeviceUnits, int landscapeNeedsFlip, void (CEF_CALLBACK *callback__)(cef_print_settings_t *, cef_size_t *, cef_rect_t *, int)) { return callback__(self, physicalSizeDeviceUnits, printableAreaDeviceUnits, landscapeNeedsFlip); }
	// void gocef_print_settings_set_device_name(cef_print_settings_t * self, cef_string_t * name, void (CEF_CALLBACK *callback__)(cef_print_settings_t *, cef_string_t *)) { return callback__(self, name); }
	// cef_string_userfree_t gocef_print_settings_get_device_name(cef_print_settings_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_print_settings_t *)) { return callback__(self); }
	// void gocef_print_settings_set_dpi(cef_print_settings_t * self, int dpi, void (CEF_CALLBACK *callback__)(cef_print_settings_t *, int)) { return callback__(self, dpi); }
	// int gocef_print_settings_get_dpi(cef_print_settings_t * self, int (CEF_CALLBACK *callback__)(cef_print_settings_t *)) { return callback__(self); }
	// void gocef_print_settings_set_page_ranges(cef_print_settings_t * self, size_t rangesCount, cef_range_t * ranges, void (CEF_CALLBACK *callback__)(cef_print_settings_t *, size_t, cef_range_t *)) { return callback__(self, rangesCount, ranges); }
	// size_t gocef_print_settings_get_page_ranges_count(cef_print_settings_t * self, size_t (CEF_CALLBACK *callback__)(cef_print_settings_t *)) { return callback__(self); }
	// void gocef_print_settings_get_page_ranges(cef_print_settings_t * self, size_t * rangesCount, cef_range_t * ranges, void (CEF_CALLBACK *callback__)(cef_print_settings_t *, size_t *, cef_range_t *)) { return callback__(self, rangesCount, ranges); }
	// void gocef_print_settings_set_selection_only(cef_print_settings_t * self, int selectionOnly, void (CEF_CALLBACK *callback__)(cef_print_settings_t *, int)) { return callback__(self, selectionOnly); }
	// int gocef_print_settings_is_selection_only(cef_print_settings_t * self, int (CEF_CALLBACK *callback__)(cef_print_settings_t *)) { return callback__(self); }
	// void gocef_print_settings_set_collate(cef_print_settings_t * self, int collate, void (CEF_CALLBACK *callback__)(cef_print_settings_t *, int)) { return callback__(self, collate); }
	// int gocef_print_settings_will_collate(cef_print_settings_t * self, int (CEF_CALLBACK *callback__)(cef_print_settings_t *)) { return callback__(self); }
	// void gocef_print_settings_set_color_model(cef_print_settings_t * self, cef_color_model_t model, void (CEF_CALLBACK *callback__)(cef_print_settings_t *, cef_color_model_t)) { return callback__(self, model); }
	// cef_color_model_t gocef_print_settings_get_color_model(cef_print_settings_t * self, cef_color_model_t (CEF_CALLBACK *callback__)(cef_print_settings_t *)) { return callback__(self); }
	// void gocef_print_settings_set_copies(cef_print_settings_t * self, int copies, void (CEF_CALLBACK *callback__)(cef_print_settings_t *, int)) { return callback__(self, copies); }
	// int gocef_print_settings_get_copies(cef_print_settings_t * self, int (CEF_CALLBACK *callback__)(cef_print_settings_t *)) { return callback__(self); }
	// void gocef_print_settings_set_duplex_mode(cef_print_settings_t * self, cef_duplex_mode_t mode, void (CEF_CALLBACK *callback__)(cef_print_settings_t *, cef_duplex_mode_t)) { return callback__(self, mode); }
	// cef_duplex_mode_t gocef_print_settings_get_duplex_mode(cef_print_settings_t * self, cef_duplex_mode_t (CEF_CALLBACK *callback__)(cef_print_settings_t *)) { return callback__(self); }
	"C"
)

// PrintSettings (cef_print_settings_t from include/capi/cef_print_settings_capi.h)
// Structure representing print settings.
type PrintSettings C.cef_print_settings_t

func (d *PrintSettings) toNative() *C.cef_print_settings_t {
	return (*C.cef_print_settings_t)(d)
}

// Base (base)
// Base structure.
func (d *PrintSettings) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// IsValid (is_valid)
// Returns true (1) if this object is valid. Do not call any other functions
// if this function returns false (0).
func (d *PrintSettings) IsValid() int32 {
	return int32(C.gocef_print_settings_is_valid(d.toNative(), d.is_valid))
}

// IsReadOnly (is_read_only)
// Returns true (1) if the values of this object are read-only. Some APIs may
// expose read-only objects.
func (d *PrintSettings) IsReadOnly() int32 {
	return int32(C.gocef_print_settings_is_read_only(d.toNative(), d.is_read_only))
}

// Copy (copy)
// Returns a writable copy of this object.
func (d *PrintSettings) Copy() *PrintSettings {
	return (*PrintSettings)(C.gocef_print_settings_copy(d.toNative(), d.copy))
}

// SetOrientation (set_orientation)
// Set the page orientation.
func (d *PrintSettings) SetOrientation(landscape int32) {
	C.gocef_print_settings_set_orientation(d.toNative(), C.int(landscape), d.set_orientation)
}

// IsLandscape (is_landscape)
// Returns true (1) if the orientation is landscape.
func (d *PrintSettings) IsLandscape() int32 {
	return int32(C.gocef_print_settings_is_landscape(d.toNative(), d.is_landscape))
}

// SetPrinterPrintableArea (set_printer_printable_area)
// Set the printer printable area in device units. Some platforms already
// provide flipped area. Set |landscape_needs_flip| to false (0) on those
// platforms to avoid double flipping.
func (d *PrintSettings) SetPrinterPrintableArea(physicalSizeDeviceUnits *Size, printableAreaDeviceUnits *Rect, landscapeNeedsFlip int32) {
	C.gocef_print_settings_set_printer_printable_area(d.toNative(), physicalSizeDeviceUnits.toNative(&C.cef_size_t{}), printableAreaDeviceUnits.toNative(&C.cef_rect_t{}), C.int(landscapeNeedsFlip), d.set_printer_printable_area)
}

// SetDeviceName (set_device_name)
// Set the device name.
func (d *PrintSettings) SetDeviceName(name string) {
	name_ := C.cef_string_userfree_alloc()
	setCEFStr(name, name_)
	defer func() {
		C.cef_string_userfree_free(name_)
	}()
	C.gocef_print_settings_set_device_name(d.toNative(), (*C.cef_string_t)(name_), d.set_device_name)
}

// GetDeviceName (get_device_name)
// Get the device name.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *PrintSettings) GetDeviceName() string {
	return cefuserfreestrToString(C.gocef_print_settings_get_device_name(d.toNative(), d.get_device_name))
}

// SetDpi (set_dpi)
// Set the DPI (dots per inch).
func (d *PrintSettings) SetDpi(dpi int32) {
	C.gocef_print_settings_set_dpi(d.toNative(), C.int(dpi), d.set_dpi)
}

// GetDpi (get_dpi)
// Get the DPI (dots per inch).
func (d *PrintSettings) GetDpi() int32 {
	return int32(C.gocef_print_settings_get_dpi(d.toNative(), d.get_dpi))
}

// SetPageRanges (set_page_ranges)
// Set the page ranges.
func (d *PrintSettings) SetPageRanges(rangesCount uint64, ranges *Range) {
	C.gocef_print_settings_set_page_ranges(d.toNative(), C.size_t(rangesCount), ranges.toNative(&C.cef_range_t{}), d.set_page_ranges)
}

// GetPageRangesCount (get_page_ranges_count)
// Returns the number of page ranges that currently exist.
func (d *PrintSettings) GetPageRangesCount() uint64 {
	return uint64(C.gocef_print_settings_get_page_ranges_count(d.toNative(), d.get_page_ranges_count))
}

// GetPageRanges (get_page_ranges)
// Retrieve the page ranges.
func (d *PrintSettings) GetPageRanges(rangesCount *uint64, ranges *Range) {
	C.gocef_print_settings_get_page_ranges(d.toNative(), (*C.size_t)(rangesCount), ranges.toNative(&C.cef_range_t{}), d.get_page_ranges)
}

// SetSelectionOnly (set_selection_only)
// Set whether only the selection will be printed.
func (d *PrintSettings) SetSelectionOnly(selectionOnly int32) {
	C.gocef_print_settings_set_selection_only(d.toNative(), C.int(selectionOnly), d.set_selection_only)
}

// IsSelectionOnly (is_selection_only)
// Returns true (1) if only the selection will be printed.
func (d *PrintSettings) IsSelectionOnly() int32 {
	return int32(C.gocef_print_settings_is_selection_only(d.toNative(), d.is_selection_only))
}

// SetCollate (set_collate)
// Set whether pages will be collated.
func (d *PrintSettings) SetCollate(collate int32) {
	C.gocef_print_settings_set_collate(d.toNative(), C.int(collate), d.set_collate)
}

// WillCollate (will_collate)
// Returns true (1) if pages will be collated.
func (d *PrintSettings) WillCollate() int32 {
	return int32(C.gocef_print_settings_will_collate(d.toNative(), d.will_collate))
}

// SetColorModel (set_color_model)
// Set the color model.
func (d *PrintSettings) SetColorModel(model ColorModel) {
	C.gocef_print_settings_set_color_model(d.toNative(), C.cef_color_model_t(model), d.set_color_model)
}

// GetColorModel (get_color_model)
// Get the color model.
func (d *PrintSettings) GetColorModel() ColorModel {
	return ColorModel(C.gocef_print_settings_get_color_model(d.toNative(), d.get_color_model))
}

// SetCopies (set_copies)
// Set the number of copies.
func (d *PrintSettings) SetCopies(copies int32) {
	C.gocef_print_settings_set_copies(d.toNative(), C.int(copies), d.set_copies)
}

// GetCopies (get_copies)
// Get the number of copies.
func (d *PrintSettings) GetCopies() int32 {
	return int32(C.gocef_print_settings_get_copies(d.toNative(), d.get_copies))
}

// SetDuplexMode (set_duplex_mode)
// Set the duplex mode.
func (d *PrintSettings) SetDuplexMode(mode DuplexMode) {
	C.gocef_print_settings_set_duplex_mode(d.toNative(), C.cef_duplex_mode_t(mode), d.set_duplex_mode)
}

// GetDuplexMode (get_duplex_mode)
// Get the duplex mode.
func (d *PrintSettings) GetDuplexMode() DuplexMode {
	return DuplexMode(C.gocef_print_settings_get_duplex_mode(d.toNative(), d.get_duplex_mode))
}
