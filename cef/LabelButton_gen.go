// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// cef_menu_button_t * gocef_label_button_as_menu_button(cef_label_button_t * self, cef_menu_button_t * (CEF_CALLBACK *callback__)(cef_label_button_t *)) { return callback__(self); }
	// void gocef_label_button_set_text(cef_label_button_t * self, cef_string_t * text, void (CEF_CALLBACK *callback__)(cef_label_button_t *, cef_string_t *)) { return callback__(self, text); }
	// cef_string_userfree_t gocef_label_button_get_text(cef_label_button_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_label_button_t *)) { return callback__(self); }
	// void gocef_label_button_set_image(cef_label_button_t * self, cef_button_state_t buttonState, cef_image_t * image, void (CEF_CALLBACK *callback__)(cef_label_button_t *, cef_button_state_t, cef_image_t *)) { return callback__(self, buttonState, image); }
	// cef_image_t * gocef_label_button_get_image(cef_label_button_t * self, cef_button_state_t buttonState, cef_image_t * (CEF_CALLBACK *callback__)(cef_label_button_t *, cef_button_state_t)) { return callback__(self, buttonState); }
	// void gocef_label_button_set_text_color(cef_label_button_t * self, cef_button_state_t forState, cef_color_t color, void (CEF_CALLBACK *callback__)(cef_label_button_t *, cef_button_state_t, cef_color_t)) { return callback__(self, forState, color); }
	// void gocef_label_button_set_enabled_text_colors(cef_label_button_t * self, cef_color_t color, void (CEF_CALLBACK *callback__)(cef_label_button_t *, cef_color_t)) { return callback__(self, color); }
	// void gocef_label_button_set_font_list(cef_label_button_t * self, cef_string_t * fontList, void (CEF_CALLBACK *callback__)(cef_label_button_t *, cef_string_t *)) { return callback__(self, fontList); }
	// void gocef_label_button_set_horizontal_alignment(cef_label_button_t * self, cef_horizontal_alignment_t alignment, void (CEF_CALLBACK *callback__)(cef_label_button_t *, cef_horizontal_alignment_t)) { return callback__(self, alignment); }
	// void gocef_label_button_set_minimum_size(cef_label_button_t * self, cef_size_t * size, void (CEF_CALLBACK *callback__)(cef_label_button_t *, cef_size_t *)) { return callback__(self, size); }
	// void gocef_label_button_set_maximum_size(cef_label_button_t * self, cef_size_t * size, void (CEF_CALLBACK *callback__)(cef_label_button_t *, cef_size_t *)) { return callback__(self, size); }
	"C"
)

// LabelButton (cef_label_button_t from include/capi/views/cef_label_button_capi.h)
// LabelButton is a button with optional text and/or icon. Methods must be
// called on the browser process UI thread unless otherwise indicated.
type LabelButton C.cef_label_button_t

func (d *LabelButton) toNative() *C.cef_label_button_t {
	return (*C.cef_label_button_t)(d)
}

// Base (base)
// Base structure.
func (d *LabelButton) Base() *Button {
	return (*Button)(&d.base)
}

// AsMenuButton (as_menu_button)
// Returns this LabelButton as a MenuButton or NULL if this is not a
// MenuButton.
func (d *LabelButton) AsMenuButton() *MenuButton {
	return (*MenuButton)(C.gocef_label_button_as_menu_button(d.toNative(), d.as_menu_button))
}

// SetText (set_text)
// Sets the text shown on the LabelButton. By default |text| will also be used
// as the accessible name.
func (d *LabelButton) SetText(text string) {
	text_ := C.cef_string_userfree_alloc()
	setCEFStr(text, text_)
	defer func() {
		C.cef_string_userfree_free(text_)
	}()
	C.gocef_label_button_set_text(d.toNative(), (*C.cef_string_t)(text_), d.set_text)
}

// GetText (get_text)
// Returns the text shown on the LabelButton.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *LabelButton) GetText() string {
	return cefuserfreestrToString(C.gocef_label_button_get_text(d.toNative(), d.get_text))
}

// SetImage (set_image)
// Sets the image shown for |button_state|. When this Button is drawn if no
// image exists for the current state then the image for
// CEF_BUTTON_STATE_NORMAL, if any, will be shown.
func (d *LabelButton) SetImage(buttonState ButtonState, image *Image) {
	C.gocef_label_button_set_image(d.toNative(), C.cef_button_state_t(buttonState), image.toNative(), d.set_image)
}

// GetImage (get_image)
// Returns the image shown for |button_state|. If no image exists for that
// state then the image for CEF_BUTTON_STATE_NORMAL will be returned.
func (d *LabelButton) GetImage(buttonState ButtonState) *Image {
	return (*Image)(C.gocef_label_button_get_image(d.toNative(), C.cef_button_state_t(buttonState), d.get_image))
}

// SetTextColor (set_text_color)
// Sets the text color shown for the specified button |for_state| to |color|.
func (d *LabelButton) SetTextColor(forState ButtonState, color Color) {
	C.gocef_label_button_set_text_color(d.toNative(), C.cef_button_state_t(forState), C.cef_color_t(color), d.set_text_color)
}

// SetEnabledTextColors (set_enabled_text_colors)
// Sets the text colors shown for the non-disabled states to |color|.
func (d *LabelButton) SetEnabledTextColors(color Color) {
	C.gocef_label_button_set_enabled_text_colors(d.toNative(), C.cef_color_t(color), d.set_enabled_text_colors)
}

// SetFontList (set_font_list)
// Sets the font list. The format is "<FONT_FAMILY_LIST>,[STYLES] <SIZE>",
// where: - FONT_FAMILY_LIST is a comma-separated list of font family names, -
// STYLES is an optional space-separated list of style names (case-sensitive
//   "Bold" and "Italic" are supported), and
// - SIZE is an integer font size in pixels with the suffix "px".
//
// Here are examples of valid font description strings: - "Arial, Helvetica,
// Bold Italic 14px" - "Arial, 14px"
func (d *LabelButton) SetFontList(fontList string) {
	fontList_ := C.cef_string_userfree_alloc()
	setCEFStr(fontList, fontList_)
	defer func() {
		C.cef_string_userfree_free(fontList_)
	}()
	C.gocef_label_button_set_font_list(d.toNative(), (*C.cef_string_t)(fontList_), d.set_font_list)
}

// SetHorizontalAlignment (set_horizontal_alignment)
// Sets the horizontal alignment; reversed in RTL. Default is
// CEF_HORIZONTAL_ALIGNMENT_CENTER.
func (d *LabelButton) SetHorizontalAlignment(alignment HorizontalAlignment) {
	C.gocef_label_button_set_horizontal_alignment(d.toNative(), C.cef_horizontal_alignment_t(alignment), d.set_horizontal_alignment)
}

// SetMinimumSize (set_minimum_size)
// Reset the minimum size of this LabelButton to |size|.
func (d *LabelButton) SetMinimumSize(size *Size) {
	C.gocef_label_button_set_minimum_size(d.toNative(), size.toNative(&C.cef_size_t{}), d.set_minimum_size)
}

// SetMaximumSize (set_maximum_size)
// Reset the maximum size of this LabelButton to |size|.
func (d *LabelButton) SetMaximumSize(size *Size) {
	C.gocef_label_button_set_maximum_size(d.toNative(), size.toNative(&C.cef_size_t{}), d.set_maximum_size)
}
