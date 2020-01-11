// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// int gocef_context_menu_params_get_xcoord(cef_context_menu_params_t * self, int (CEF_CALLBACK *callback__)(cef_context_menu_params_t *)) { return callback__(self); }
	// int gocef_context_menu_params_get_ycoord(cef_context_menu_params_t * self, int (CEF_CALLBACK *callback__)(cef_context_menu_params_t *)) { return callback__(self); }
	// cef_context_menu_type_flags_t gocef_context_menu_params_get_type_flags(cef_context_menu_params_t * self, cef_context_menu_type_flags_t (CEF_CALLBACK *callback__)(cef_context_menu_params_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_context_menu_params_get_link_url(cef_context_menu_params_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_context_menu_params_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_context_menu_params_get_unfiltered_link_url(cef_context_menu_params_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_context_menu_params_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_context_menu_params_get_source_url(cef_context_menu_params_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_context_menu_params_t *)) { return callback__(self); }
	// int gocef_context_menu_params_has_image_contents(cef_context_menu_params_t * self, int (CEF_CALLBACK *callback__)(cef_context_menu_params_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_context_menu_params_get_title_text(cef_context_menu_params_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_context_menu_params_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_context_menu_params_get_page_url(cef_context_menu_params_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_context_menu_params_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_context_menu_params_get_frame_url(cef_context_menu_params_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_context_menu_params_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_context_menu_params_get_frame_charset(cef_context_menu_params_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_context_menu_params_t *)) { return callback__(self); }
	// cef_context_menu_media_type_t gocef_context_menu_params_get_media_type(cef_context_menu_params_t * self, cef_context_menu_media_type_t (CEF_CALLBACK *callback__)(cef_context_menu_params_t *)) { return callback__(self); }
	// cef_context_menu_media_state_flags_t gocef_context_menu_params_get_media_state_flags(cef_context_menu_params_t * self, cef_context_menu_media_state_flags_t (CEF_CALLBACK *callback__)(cef_context_menu_params_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_context_menu_params_get_selection_text(cef_context_menu_params_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_context_menu_params_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_context_menu_params_get_misspelled_word(cef_context_menu_params_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_context_menu_params_t *)) { return callback__(self); }
	// int gocef_context_menu_params_get_dictionary_suggestions(cef_context_menu_params_t * self, cef_string_list_t suggestions, int (CEF_CALLBACK *callback__)(cef_context_menu_params_t *, cef_string_list_t)) { return callback__(self, suggestions); }
	// int gocef_context_menu_params_is_editable(cef_context_menu_params_t * self, int (CEF_CALLBACK *callback__)(cef_context_menu_params_t *)) { return callback__(self); }
	// int gocef_context_menu_params_is_spell_check_enabled(cef_context_menu_params_t * self, int (CEF_CALLBACK *callback__)(cef_context_menu_params_t *)) { return callback__(self); }
	// cef_context_menu_edit_state_flags_t gocef_context_menu_params_get_edit_state_flags(cef_context_menu_params_t * self, cef_context_menu_edit_state_flags_t (CEF_CALLBACK *callback__)(cef_context_menu_params_t *)) { return callback__(self); }
	// int gocef_context_menu_params_is_custom_menu(cef_context_menu_params_t * self, int (CEF_CALLBACK *callback__)(cef_context_menu_params_t *)) { return callback__(self); }
	// int gocef_context_menu_params_is_pepper_menu(cef_context_menu_params_t * self, int (CEF_CALLBACK *callback__)(cef_context_menu_params_t *)) { return callback__(self); }
	"C"
)

// ContextMenuParams (cef_context_menu_params_t from include/capi/cef_context_menu_handler_capi.h)
// Provides information about the context menu state. The ethods of this
// structure can only be accessed on browser process the UI thread.
type ContextMenuParams C.cef_context_menu_params_t

func (d *ContextMenuParams) toNative() *C.cef_context_menu_params_t {
	return (*C.cef_context_menu_params_t)(d)
}

// Base (base)
// Base structure.
func (d *ContextMenuParams) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// GetXcoord (get_xcoord)
// Returns the X coordinate of the mouse where the context menu was invoked.
// Coords are relative to the associated RenderView's origin.
func (d *ContextMenuParams) GetXcoord() int32 {
	return int32(C.gocef_context_menu_params_get_xcoord(d.toNative(), d.get_xcoord))
}

// GetYcoord (get_ycoord)
// Returns the Y coordinate of the mouse where the context menu was invoked.
// Coords are relative to the associated RenderView's origin.
func (d *ContextMenuParams) GetYcoord() int32 {
	return int32(C.gocef_context_menu_params_get_ycoord(d.toNative(), d.get_ycoord))
}

// GetTypeFlags (get_type_flags)
// Returns flags representing the type of node that the context menu was
// invoked on.
func (d *ContextMenuParams) GetTypeFlags() ContextMenuTypeFlags {
	return ContextMenuTypeFlags(C.gocef_context_menu_params_get_type_flags(d.toNative(), d.get_type_flags))
}

// GetLinkURL (get_link_url)
// Returns the URL of the link, if any, that encloses the node that the
// context menu was invoked on.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *ContextMenuParams) GetLinkURL() string {
	return cefuserfreestrToString(C.gocef_context_menu_params_get_link_url(d.toNative(), d.get_link_url))
}

// GetUnfilteredLinkURL (get_unfiltered_link_url)
// Returns the link URL, if any, to be used ONLY for "copy link address". We
// don't validate this field in the frontend process.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *ContextMenuParams) GetUnfilteredLinkURL() string {
	return cefuserfreestrToString(C.gocef_context_menu_params_get_unfiltered_link_url(d.toNative(), d.get_unfiltered_link_url))
}

// GetSourceURL (get_source_url)
// Returns the source URL, if any, for the element that the context menu was
// invoked on. Example of elements with source URLs are img, audio, and video.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *ContextMenuParams) GetSourceURL() string {
	return cefuserfreestrToString(C.gocef_context_menu_params_get_source_url(d.toNative(), d.get_source_url))
}

// HasImageContents (has_image_contents)
// Returns true (1) if the context menu was invoked on an image which has non-
// NULL contents.
func (d *ContextMenuParams) HasImageContents() int32 {
	return int32(C.gocef_context_menu_params_has_image_contents(d.toNative(), d.has_image_contents))
}

// GetTitleText (get_title_text)
// Returns the title text or the alt text if the context menu was invoked on
// an image.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *ContextMenuParams) GetTitleText() string {
	return cefuserfreestrToString(C.gocef_context_menu_params_get_title_text(d.toNative(), d.get_title_text))
}

// GetPageURL (get_page_url)
// Returns the URL of the top level page that the context menu was invoked on.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *ContextMenuParams) GetPageURL() string {
	return cefuserfreestrToString(C.gocef_context_menu_params_get_page_url(d.toNative(), d.get_page_url))
}

// GetFrameURL (get_frame_url)
// Returns the URL of the subframe that the context menu was invoked on.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *ContextMenuParams) GetFrameURL() string {
	return cefuserfreestrToString(C.gocef_context_menu_params_get_frame_url(d.toNative(), d.get_frame_url))
}

// GetFrameCharset (get_frame_charset)
// Returns the character encoding of the subframe that the context menu was
// invoked on.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *ContextMenuParams) GetFrameCharset() string {
	return cefuserfreestrToString(C.gocef_context_menu_params_get_frame_charset(d.toNative(), d.get_frame_charset))
}

// GetMediaType (get_media_type)
// Returns the type of context node that the context menu was invoked on.
func (d *ContextMenuParams) GetMediaType() ContextMenuMediaType {
	return ContextMenuMediaType(C.gocef_context_menu_params_get_media_type(d.toNative(), d.get_media_type))
}

// GetMediaStateFlags (get_media_state_flags)
// Returns flags representing the actions supported by the media element, if
// any, that the context menu was invoked on.
func (d *ContextMenuParams) GetMediaStateFlags() ContextMenuMediaStateFlags {
	return ContextMenuMediaStateFlags(C.gocef_context_menu_params_get_media_state_flags(d.toNative(), d.get_media_state_flags))
}

// GetSelectionText (get_selection_text)
// Returns the text of the selection, if any, that the context menu was
// invoked on.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *ContextMenuParams) GetSelectionText() string {
	return cefuserfreestrToString(C.gocef_context_menu_params_get_selection_text(d.toNative(), d.get_selection_text))
}

// GetMisspelledWord (get_misspelled_word)
// Returns the text of the misspelled word, if any, that the context menu was
// invoked on.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *ContextMenuParams) GetMisspelledWord() string {
	return cefuserfreestrToString(C.gocef_context_menu_params_get_misspelled_word(d.toNative(), d.get_misspelled_word))
}

// GetDictionarySuggestions (get_dictionary_suggestions)
// Returns true (1) if suggestions exist, false (0) otherwise. Fills in
// |suggestions| from the spell check service for the misspelled word if there
// is one.
func (d *ContextMenuParams) GetDictionarySuggestions(suggestions StringList) int32 {
	return int32(C.gocef_context_menu_params_get_dictionary_suggestions(d.toNative(), C.cef_string_list_t(suggestions), d.get_dictionary_suggestions))
}

// IsEditable (is_editable)
// Returns true (1) if the context menu was invoked on an editable node.
func (d *ContextMenuParams) IsEditable() int32 {
	return int32(C.gocef_context_menu_params_is_editable(d.toNative(), d.is_editable))
}

// IsSpellCheckEnabled (is_spell_check_enabled)
// Returns true (1) if the context menu was invoked on an editable node where
// spell-check is enabled.
func (d *ContextMenuParams) IsSpellCheckEnabled() int32 {
	return int32(C.gocef_context_menu_params_is_spell_check_enabled(d.toNative(), d.is_spell_check_enabled))
}

// GetEditStateFlags (get_edit_state_flags)
// Returns flags representing the actions supported by the editable node, if
// any, that the context menu was invoked on.
func (d *ContextMenuParams) GetEditStateFlags() ContextMenuEditStateFlags {
	return ContextMenuEditStateFlags(C.gocef_context_menu_params_get_edit_state_flags(d.toNative(), d.get_edit_state_flags))
}

// IsCustomMenu (is_custom_menu)
// Returns true (1) if the context menu contains items specified by the
// renderer process (for example, plugin placeholder or pepper plugin menu
// items).
func (d *ContextMenuParams) IsCustomMenu() int32 {
	return int32(C.gocef_context_menu_params_is_custom_menu(d.toNative(), d.is_custom_menu))
}

// IsPepperMenu (is_pepper_menu)
// Returns true (1) if the context menu was invoked from a pepper plugin.
func (d *ContextMenuParams) IsPepperMenu() int32 {
	return int32(C.gocef_context_menu_params_is_pepper_menu(d.toNative(), d.is_pepper_menu))
}
