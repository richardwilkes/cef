// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// void gocef_textfield_set_password_input(cef_textfield_t * self, int passwordInput, void (CEF_CALLBACK *callback__)(cef_textfield_t *, int)) { return callback__(self, passwordInput); }
	// int gocef_textfield_is_password_input(cef_textfield_t * self, int (CEF_CALLBACK *callback__)(cef_textfield_t *)) { return callback__(self); }
	// void gocef_textfield_set_read_only(cef_textfield_t * self, int readOnly, void (CEF_CALLBACK *callback__)(cef_textfield_t *, int)) { return callback__(self, readOnly); }
	// int gocef_textfield_is_read_only(cef_textfield_t * self, int (CEF_CALLBACK *callback__)(cef_textfield_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_textfield_get_text(cef_textfield_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_textfield_t *)) { return callback__(self); }
	// void gocef_textfield_set_text(cef_textfield_t * self, cef_string_t * text, void (CEF_CALLBACK *callback__)(cef_textfield_t *, cef_string_t *)) { return callback__(self, text); }
	// void gocef_textfield_append_text(cef_textfield_t * self, cef_string_t * text, void (CEF_CALLBACK *callback__)(cef_textfield_t *, cef_string_t *)) { return callback__(self, text); }
	// void gocef_textfield_insert_or_replace_text(cef_textfield_t * self, cef_string_t * text, void (CEF_CALLBACK *callback__)(cef_textfield_t *, cef_string_t *)) { return callback__(self, text); }
	// int gocef_textfield_has_selection(cef_textfield_t * self, int (CEF_CALLBACK *callback__)(cef_textfield_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_textfield_get_selected_text(cef_textfield_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_textfield_t *)) { return callback__(self); }
	// void gocef_textfield_select_all(cef_textfield_t * self, int reversed, void (CEF_CALLBACK *callback__)(cef_textfield_t *, int)) { return callback__(self, reversed); }
	// void gocef_textfield_clear_selection(cef_textfield_t * self, void (CEF_CALLBACK *callback__)(cef_textfield_t *)) { return callback__(self); }
	// cef_range_t gocef_textfield_get_selected_range(cef_textfield_t * self, cef_range_t (CEF_CALLBACK *callback__)(cef_textfield_t *)) { return callback__(self); }
	// void gocef_textfield_select_range(cef_textfield_t * self, cef_range_t * range_r, void (CEF_CALLBACK *callback__)(cef_textfield_t *, cef_range_t *)) { return callback__(self, range_r); }
	// size_t gocef_textfield_get_cursor_position(cef_textfield_t * self, size_t (CEF_CALLBACK *callback__)(cef_textfield_t *)) { return callback__(self); }
	// void gocef_textfield_set_text_color(cef_textfield_t * self, cef_color_t color, void (CEF_CALLBACK *callback__)(cef_textfield_t *, cef_color_t)) { return callback__(self, color); }
	// cef_color_t gocef_textfield_get_text_color(cef_textfield_t * self, cef_color_t (CEF_CALLBACK *callback__)(cef_textfield_t *)) { return callback__(self); }
	// void gocef_textfield_set_selection_text_color(cef_textfield_t * self, cef_color_t color, void (CEF_CALLBACK *callback__)(cef_textfield_t *, cef_color_t)) { return callback__(self, color); }
	// cef_color_t gocef_textfield_get_selection_text_color(cef_textfield_t * self, cef_color_t (CEF_CALLBACK *callback__)(cef_textfield_t *)) { return callback__(self); }
	// void gocef_textfield_set_selection_background_color(cef_textfield_t * self, cef_color_t color, void (CEF_CALLBACK *callback__)(cef_textfield_t *, cef_color_t)) { return callback__(self, color); }
	// cef_color_t gocef_textfield_get_selection_background_color(cef_textfield_t * self, cef_color_t (CEF_CALLBACK *callback__)(cef_textfield_t *)) { return callback__(self); }
	// void gocef_textfield_set_font_list(cef_textfield_t * self, cef_string_t * fontList, void (CEF_CALLBACK *callback__)(cef_textfield_t *, cef_string_t *)) { return callback__(self, fontList); }
	// void gocef_textfield_apply_text_color(cef_textfield_t * self, cef_color_t color, cef_range_t * range_r, void (CEF_CALLBACK *callback__)(cef_textfield_t *, cef_color_t, cef_range_t *)) { return callback__(self, color, range_r); }
	// void gocef_textfield_apply_text_style(cef_textfield_t * self, cef_text_style_t style, int add, cef_range_t * range_r, void (CEF_CALLBACK *callback__)(cef_textfield_t *, cef_text_style_t, int, cef_range_t *)) { return callback__(self, style, add, range_r); }
	// int gocef_textfield_is_command_enabled(cef_textfield_t * self, int commandID, int (CEF_CALLBACK *callback__)(cef_textfield_t *, int)) { return callback__(self, commandID); }
	// void gocef_textfield_execute_command(cef_textfield_t * self, int commandID, void (CEF_CALLBACK *callback__)(cef_textfield_t *, int)) { return callback__(self, commandID); }
	// void gocef_textfield_clear_edit_history(cef_textfield_t * self, void (CEF_CALLBACK *callback__)(cef_textfield_t *)) { return callback__(self); }
	// void gocef_textfield_set_placeholder_text(cef_textfield_t * self, cef_string_t * text, void (CEF_CALLBACK *callback__)(cef_textfield_t *, cef_string_t *)) { return callback__(self, text); }
	// cef_string_userfree_t gocef_textfield_get_placeholder_text(cef_textfield_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_textfield_t *)) { return callback__(self); }
	// void gocef_textfield_set_placeholder_text_color(cef_textfield_t * self, cef_color_t color, void (CEF_CALLBACK *callback__)(cef_textfield_t *, cef_color_t)) { return callback__(self, color); }
	// void gocef_textfield_set_accessible_name(cef_textfield_t * self, cef_string_t * name, void (CEF_CALLBACK *callback__)(cef_textfield_t *, cef_string_t *)) { return callback__(self, name); }
	"C"
)

// Textfield (cef_textfield_t from include/capi/views/cef_textfield_capi.h)
// A Textfield supports editing of text. This control is custom rendered with no
// platform-specific code. Methods must be called on the browser process UI
// thread unless otherwise indicated.
type Textfield C.cef_textfield_t

func (d *Textfield) toNative() *C.cef_textfield_t {
	return (*C.cef_textfield_t)(d)
}

// Base (base)
// Base structure.
func (d *Textfield) Base() *View {
	return (*View)(&d.base)
}

// SetPasswordInput (set_password_input)
// Sets whether the text will be displayed as asterisks.
func (d *Textfield) SetPasswordInput(passwordInput int32) {
	C.gocef_textfield_set_password_input(d.toNative(), C.int(passwordInput), d.set_password_input)
}

// IsPasswordInput (is_password_input)
// Returns true (1) if the text will be displayed as asterisks.
func (d *Textfield) IsPasswordInput() int32 {
	return int32(C.gocef_textfield_is_password_input(d.toNative(), d.is_password_input))
}

// SetReadOnly (set_read_only)
// Sets whether the text will read-only.
func (d *Textfield) SetReadOnly(readOnly int32) {
	C.gocef_textfield_set_read_only(d.toNative(), C.int(readOnly), d.set_read_only)
}

// IsReadOnly (is_read_only)
// Returns true (1) if the text is read-only.
func (d *Textfield) IsReadOnly() int32 {
	return int32(C.gocef_textfield_is_read_only(d.toNative(), d.is_read_only))
}

// GetText (get_text)
// Returns the currently displayed text.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Textfield) GetText() string {
	return cefuserfreestrToString(C.gocef_textfield_get_text(d.toNative(), d.get_text))
}

// SetText (set_text)
// Sets the contents to |text|. The cursor will be moved to end of the text if
// the current position is outside of the text range.
func (d *Textfield) SetText(text string) {
	text_ := C.cef_string_userfree_alloc()
	setCEFStr(text, text_)
	defer func() {
		C.cef_string_userfree_free(text_)
	}()
	C.gocef_textfield_set_text(d.toNative(), (*C.cef_string_t)(text_), d.set_text)
}

// AppendText (append_text)
// Appends |text| to the previously-existing text.
func (d *Textfield) AppendText(text string) {
	text_ := C.cef_string_userfree_alloc()
	setCEFStr(text, text_)
	defer func() {
		C.cef_string_userfree_free(text_)
	}()
	C.gocef_textfield_append_text(d.toNative(), (*C.cef_string_t)(text_), d.append_text)
}

// InsertOrReplaceText (insert_or_replace_text)
// Inserts |text| at the current cursor position replacing any selected text.
func (d *Textfield) InsertOrReplaceText(text string) {
	text_ := C.cef_string_userfree_alloc()
	setCEFStr(text, text_)
	defer func() {
		C.cef_string_userfree_free(text_)
	}()
	C.gocef_textfield_insert_or_replace_text(d.toNative(), (*C.cef_string_t)(text_), d.insert_or_replace_text)
}

// HasSelection (has_selection)
// Returns true (1) if there is any selected text.
func (d *Textfield) HasSelection() int32 {
	return int32(C.gocef_textfield_has_selection(d.toNative(), d.has_selection))
}

// GetSelectedText (get_selected_text)
// Returns the currently selected text.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Textfield) GetSelectedText() string {
	return cefuserfreestrToString(C.gocef_textfield_get_selected_text(d.toNative(), d.get_selected_text))
}

// SelectAll (select_all)
// Selects all text. If |reversed| is true (1) the range will end at the
// logical beginning of the text; this generally shows the leading portion of
// text that overflows its display area.
func (d *Textfield) SelectAll(reversed int32) {
	C.gocef_textfield_select_all(d.toNative(), C.int(reversed), d.select_all)
}

// ClearSelection (clear_selection)
// Clears the text selection and sets the caret to the end.
func (d *Textfield) ClearSelection() {
	C.gocef_textfield_clear_selection(d.toNative(), d.clear_selection)
}

// GetSelectedRange (get_selected_range)
// Returns the selected logical text range.
func (d *Textfield) GetSelectedRange() Range {
	cresult__ := C.gocef_textfield_get_selected_range(d.toNative(), d.get_selected_range)
	var result__ Range
	(&cresult__).intoGo(&result__)
	return result__
}

// SelectRange (select_range)
// Selects the specified logical text range.
func (d *Textfield) SelectRange(range_r *Range) {
	C.gocef_textfield_select_range(d.toNative(), range_r.toNative(&C.cef_range_t{}), d.select_range)
}

// GetCursorPosition (get_cursor_position)
// Returns the current cursor position.
func (d *Textfield) GetCursorPosition() uint64 {
	return uint64(C.gocef_textfield_get_cursor_position(d.toNative(), d.get_cursor_position))
}

// SetTextColor (set_text_color)
// Sets the text color.
func (d *Textfield) SetTextColor(color Color) {
	C.gocef_textfield_set_text_color(d.toNative(), C.cef_color_t(color), d.set_text_color)
}

// GetTextColor (get_text_color)
// Returns the text color.
func (d *Textfield) GetTextColor() Color {
	return Color(C.gocef_textfield_get_text_color(d.toNative(), d.get_text_color))
}

// SetSelectionTextColor (set_selection_text_color)
// Sets the selection text color.
func (d *Textfield) SetSelectionTextColor(color Color) {
	C.gocef_textfield_set_selection_text_color(d.toNative(), C.cef_color_t(color), d.set_selection_text_color)
}

// GetSelectionTextColor (get_selection_text_color)
// Returns the selection text color.
func (d *Textfield) GetSelectionTextColor() Color {
	return Color(C.gocef_textfield_get_selection_text_color(d.toNative(), d.get_selection_text_color))
}

// SetSelectionBackgroundColor (set_selection_background_color)
// Sets the selection background color.
func (d *Textfield) SetSelectionBackgroundColor(color Color) {
	C.gocef_textfield_set_selection_background_color(d.toNative(), C.cef_color_t(color), d.set_selection_background_color)
}

// GetSelectionBackgroundColor (get_selection_background_color)
// Returns the selection background color.
func (d *Textfield) GetSelectionBackgroundColor() Color {
	return Color(C.gocef_textfield_get_selection_background_color(d.toNative(), d.get_selection_background_color))
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
func (d *Textfield) SetFontList(fontList string) {
	fontList_ := C.cef_string_userfree_alloc()
	setCEFStr(fontList, fontList_)
	defer func() {
		C.cef_string_userfree_free(fontList_)
	}()
	C.gocef_textfield_set_font_list(d.toNative(), (*C.cef_string_t)(fontList_), d.set_font_list)
}

// ApplyTextColor (apply_text_color)
// Applies |color| to the specified |range| without changing the default
// color. If |range| is NULL the color will be set on the complete text
// contents.
func (d *Textfield) ApplyTextColor(color Color, range_r *Range) {
	C.gocef_textfield_apply_text_color(d.toNative(), C.cef_color_t(color), range_r.toNative(&C.cef_range_t{}), d.apply_text_color)
}

// ApplyTextStyle (apply_text_style)
// Applies |style| to the specified |range| without changing the default
// style. If |add| is true (1) the style will be added, otherwise the style
// will be removed. If |range| is NULL the style will be set on the complete
// text contents.
func (d *Textfield) ApplyTextStyle(style TextStyle, add int32, range_r *Range) {
	C.gocef_textfield_apply_text_style(d.toNative(), C.cef_text_style_t(style), C.int(add), range_r.toNative(&C.cef_range_t{}), d.apply_text_style)
}

// IsCommandEnabled (is_command_enabled)
// Returns true (1) if the action associated with the specified command id is
// enabled. See additional comments on execute_command().
func (d *Textfield) IsCommandEnabled(commandID int32) int32 {
	return int32(C.gocef_textfield_is_command_enabled(d.toNative(), C.int(commandID), d.is_command_enabled))
}

// ExecuteCommand (execute_command)
// Performs the action associated with the specified command id. Valid values
// include IDS_APP_UNDO, IDS_APP_REDO, IDS_APP_CUT, IDS_APP_COPY,
// IDS_APP_PASTE, IDS_APP_DELETE, IDS_APP_SELECT_ALL, IDS_DELETE_* and
// IDS_MOVE_*. See include/cef_pack_strings.h for definitions.
func (d *Textfield) ExecuteCommand(commandID int32) {
	C.gocef_textfield_execute_command(d.toNative(), C.int(commandID), d.execute_command)
}

// ClearEditHistory (clear_edit_history)
// Clears Edit history.
func (d *Textfield) ClearEditHistory() {
	C.gocef_textfield_clear_edit_history(d.toNative(), d.clear_edit_history)
}

// SetPlaceholderText (set_placeholder_text)
// Sets the placeholder text that will be displayed when the Textfield is
// NULL.
func (d *Textfield) SetPlaceholderText(text string) {
	text_ := C.cef_string_userfree_alloc()
	setCEFStr(text, text_)
	defer func() {
		C.cef_string_userfree_free(text_)
	}()
	C.gocef_textfield_set_placeholder_text(d.toNative(), (*C.cef_string_t)(text_), d.set_placeholder_text)
}

// GetPlaceholderText (get_placeholder_text)
// Returns the placeholder text that will be displayed when the Textfield is
// NULL.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Textfield) GetPlaceholderText() string {
	return cefuserfreestrToString(C.gocef_textfield_get_placeholder_text(d.toNative(), d.get_placeholder_text))
}

// SetPlaceholderTextColor (set_placeholder_text_color)
// Sets the placeholder text color.
func (d *Textfield) SetPlaceholderTextColor(color Color) {
	C.gocef_textfield_set_placeholder_text_color(d.toNative(), C.cef_color_t(color), d.set_placeholder_text_color)
}

// SetAccessibleName (set_accessible_name)
// Set the accessible name that will be exposed to assistive technology (AT).
func (d *Textfield) SetAccessibleName(name string) {
	name_ := C.cef_string_userfree_alloc()
	setCEFStr(name, name_)
	defer func() {
		C.cef_string_userfree_free(name_)
	}()
	C.gocef_textfield_set_accessible_name(d.toNative(), (*C.cef_string_t)(name_), d.set_accessible_name)
}
