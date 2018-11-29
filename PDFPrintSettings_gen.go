// Code generated - DO NOT EDIT.

package cef

import (
	// #include "capi_gen.h"
	"C"
)

// PDFPrintSettings (cef_pdf_print_settings_t from include/internal/cef_types.h)
// Structure representing PDF print settings.
type PDFPrintSettings struct {
	// HeaderFooterTitle (header_footer_title)
	// Page title to display in the header. Only used if |header_footer_enabled|
	// is set to true (1).
	HeaderFooterTitle string
	// HeaderFooterUrl (header_footer_url)
	// URL to display in the footer. Only used if |header_footer_enabled| is set
	// to true (1).
	HeaderFooterUrl string
	// PageWidth (page_width)
	// Output page size in microns. If either of these values is less than or
	// equal to zero then the default paper size (A4) will be used.
	PageWidth int32
	// PageHeight (page_height)
	PageHeight int32
	// ScaleFactor (scale_factor)
	// The percentage to scale the PDF by before printing (e.g. 50 is 50%).
	// If this value is less than or equal to zero the default value of 100
	// will be used.
	ScaleFactor int32
	// MarginTop (margin_top)
	// Margins in millimeters. Only used if |margin_type| is set to
	// PDF_PRINT_MARGIN_CUSTOM.
	MarginTop float64
	// MarginRight (margin_right)
	MarginRight float64
	// MarginBottom (margin_bottom)
	MarginBottom float64
	// MarginLeft (margin_left)
	MarginLeft float64
	// MarginType (margin_type)
	// Margin type.
	MarginType PDFPrintMarginType
	// HeaderFooterEnabled (header_footer_enabled)
	// Set to true (1) to print headers and footers or false (0) to not print
	// headers and footers.
	HeaderFooterEnabled int32
	// SelectionOnly (selection_only)
	// Set to true (1) to print the selection only or false (0) to print all.
	SelectionOnly int32
	// Landscape (landscape)
	// Set to true (1) for landscape mode or false (0) for portrait mode.
	Landscape int32
	// BackgroundsEnabled (backgrounds_enabled)
	// Set to true (1) to print background graphics or false (0) to not print
	// background graphics.
	BackgroundsEnabled int32
}

// NewPDFPrintSettings creates a new PDFPrintSettings.
func NewPDFPrintSettings() *PDFPrintSettings {
	return &PDFPrintSettings{}
}

func (d *PDFPrintSettings) toNative(native *C.cef_pdf_print_settings_t) *C.cef_pdf_print_settings_t {
	setCEFStr(d.HeaderFooterTitle, &native.header_footer_title)
	setCEFStr(d.HeaderFooterUrl, &native.header_footer_url)
	native.page_width = C.int(d.PageWidth)
	native.page_height = C.int(d.PageHeight)
	native.scale_factor = C.int(d.ScaleFactor)
	native.margin_top = C.double(d.MarginTop)
	native.margin_right = C.double(d.MarginRight)
	native.margin_bottom = C.double(d.MarginBottom)
	native.margin_left = C.double(d.MarginLeft)
	native.margin_type = C.cef_pdf_print_margin_type_t(d.MarginType)
	native.header_footer_enabled = C.int(d.HeaderFooterEnabled)
	native.selection_only = C.int(d.SelectionOnly)
	native.landscape = C.int(d.Landscape)
	native.backgrounds_enabled = C.int(d.BackgroundsEnabled)
	return native
}

func (d *PDFPrintSettings) fromNative(native *C.cef_pdf_print_settings_t) *PDFPrintSettings {
	d.HeaderFooterTitle = cefstrToString(&native.header_footer_title)
	d.HeaderFooterUrl = cefstrToString(&native.header_footer_url)
	d.PageWidth = int32(native.page_width)
	d.PageHeight = int32(native.page_height)
	d.ScaleFactor = int32(native.scale_factor)
	d.MarginTop = float64(native.margin_top)
	d.MarginRight = float64(native.margin_right)
	d.MarginBottom = float64(native.margin_bottom)
	d.MarginLeft = float64(native.margin_left)
	d.MarginType = PDFPrintMarginType(native.margin_type)
	d.HeaderFooterEnabled = int32(native.header_footer_enabled)
	d.SelectionOnly = int32(native.selection_only)
	d.Landscape = int32(native.landscape)
	d.BackgroundsEnabled = int32(native.backgrounds_enabled)
	return d
}
