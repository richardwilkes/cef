// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// int gocef_menu_model_is_sub_menu(cef_menu_model_t * self, int (CEF_CALLBACK *callback__)(cef_menu_model_t *)) { return callback__(self); }
	// int gocef_menu_model_clear(cef_menu_model_t * self, int (CEF_CALLBACK *callback__)(cef_menu_model_t *)) { return callback__(self); }
	// int gocef_menu_model_get_count(cef_menu_model_t * self, int (CEF_CALLBACK *callback__)(cef_menu_model_t *)) { return callback__(self); }
	// int gocef_menu_model_add_separator(cef_menu_model_t * self, int (CEF_CALLBACK *callback__)(cef_menu_model_t *)) { return callback__(self); }
	// int gocef_menu_model_add_item(cef_menu_model_t * self, int commandID, cef_string_t * label, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, cef_string_t *)) { return callback__(self, commandID, label); }
	// int gocef_menu_model_add_check_item(cef_menu_model_t * self, int commandID, cef_string_t * label, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, cef_string_t *)) { return callback__(self, commandID, label); }
	// int gocef_menu_model_add_radio_item(cef_menu_model_t * self, int commandID, cef_string_t * label, int groupID, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, cef_string_t *, int)) { return callback__(self, commandID, label, groupID); }
	// cef_menu_model_t * gocef_menu_model_add_sub_menu(cef_menu_model_t * self, int commandID, cef_string_t * label, cef_menu_model_t * (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, cef_string_t *)) { return callback__(self, commandID, label); }
	// int gocef_menu_model_insert_separator_at(cef_menu_model_t * self, int index, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int)) { return callback__(self, index); }
	// int gocef_menu_model_insert_item_at(cef_menu_model_t * self, int index, int commandID, cef_string_t * label, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, int, cef_string_t *)) { return callback__(self, index, commandID, label); }
	// int gocef_menu_model_insert_check_item_at(cef_menu_model_t * self, int index, int commandID, cef_string_t * label, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, int, cef_string_t *)) { return callback__(self, index, commandID, label); }
	// int gocef_menu_model_insert_radio_item_at(cef_menu_model_t * self, int index, int commandID, cef_string_t * label, int groupID, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, int, cef_string_t *, int)) { return callback__(self, index, commandID, label, groupID); }
	// cef_menu_model_t * gocef_menu_model_insert_sub_menu_at(cef_menu_model_t * self, int index, int commandID, cef_string_t * label, cef_menu_model_t * (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, int, cef_string_t *)) { return callback__(self, index, commandID, label); }
	// int gocef_menu_model_remove(cef_menu_model_t * self, int commandID, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int)) { return callback__(self, commandID); }
	// int gocef_menu_model_remove_at(cef_menu_model_t * self, int index, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int)) { return callback__(self, index); }
	// int gocef_menu_model_get_index_of(cef_menu_model_t * self, int commandID, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int)) { return callback__(self, commandID); }
	// int gocef_menu_model_get_command_id_at(cef_menu_model_t * self, int index, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int)) { return callback__(self, index); }
	// int gocef_menu_model_set_command_id_at(cef_menu_model_t * self, int index, int commandID, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, int)) { return callback__(self, index, commandID); }
	// cef_string_userfree_t gocef_menu_model_get_label(cef_menu_model_t * self, int commandID, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_menu_model_t *, int)) { return callback__(self, commandID); }
	// cef_string_userfree_t gocef_menu_model_get_label_at(cef_menu_model_t * self, int index, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_menu_model_t *, int)) { return callback__(self, index); }
	// int gocef_menu_model_set_label(cef_menu_model_t * self, int commandID, cef_string_t * label, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, cef_string_t *)) { return callback__(self, commandID, label); }
	// int gocef_menu_model_set_label_at(cef_menu_model_t * self, int index, cef_string_t * label, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, cef_string_t *)) { return callback__(self, index, label); }
	// cef_menu_item_type_t gocef_menu_model_get_type(cef_menu_model_t * self, int commandID, cef_menu_item_type_t (CEF_CALLBACK *callback__)(cef_menu_model_t *, int)) { return callback__(self, commandID); }
	// cef_menu_item_type_t gocef_menu_model_get_type_at(cef_menu_model_t * self, int index, cef_menu_item_type_t (CEF_CALLBACK *callback__)(cef_menu_model_t *, int)) { return callback__(self, index); }
	// int gocef_menu_model_get_group_id(cef_menu_model_t * self, int commandID, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int)) { return callback__(self, commandID); }
	// int gocef_menu_model_get_group_id_at(cef_menu_model_t * self, int index, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int)) { return callback__(self, index); }
	// int gocef_menu_model_set_group_id(cef_menu_model_t * self, int commandID, int groupID, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, int)) { return callback__(self, commandID, groupID); }
	// int gocef_menu_model_set_group_id_at(cef_menu_model_t * self, int index, int groupID, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, int)) { return callback__(self, index, groupID); }
	// cef_menu_model_t * gocef_menu_model_get_sub_menu(cef_menu_model_t * self, int commandID, cef_menu_model_t * (CEF_CALLBACK *callback__)(cef_menu_model_t *, int)) { return callback__(self, commandID); }
	// cef_menu_model_t * gocef_menu_model_get_sub_menu_at(cef_menu_model_t * self, int index, cef_menu_model_t * (CEF_CALLBACK *callback__)(cef_menu_model_t *, int)) { return callback__(self, index); }
	// int gocef_menu_model_is_visible(cef_menu_model_t * self, int commandID, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int)) { return callback__(self, commandID); }
	// int gocef_menu_model_is_visible_at(cef_menu_model_t * self, int index, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int)) { return callback__(self, index); }
	// int gocef_menu_model_set_visible(cef_menu_model_t * self, int commandID, int visible, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, int)) { return callback__(self, commandID, visible); }
	// int gocef_menu_model_set_visible_at(cef_menu_model_t * self, int index, int visible, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, int)) { return callback__(self, index, visible); }
	// int gocef_menu_model_is_enabled(cef_menu_model_t * self, int commandID, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int)) { return callback__(self, commandID); }
	// int gocef_menu_model_is_enabled_at(cef_menu_model_t * self, int index, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int)) { return callback__(self, index); }
	// int gocef_menu_model_set_enabled(cef_menu_model_t * self, int commandID, int enabled, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, int)) { return callback__(self, commandID, enabled); }
	// int gocef_menu_model_set_enabled_at(cef_menu_model_t * self, int index, int enabled, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, int)) { return callback__(self, index, enabled); }
	// int gocef_menu_model_is_checked(cef_menu_model_t * self, int commandID, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int)) { return callback__(self, commandID); }
	// int gocef_menu_model_is_checked_at(cef_menu_model_t * self, int index, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int)) { return callback__(self, index); }
	// int gocef_menu_model_set_checked(cef_menu_model_t * self, int commandID, int checked, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, int)) { return callback__(self, commandID, checked); }
	// int gocef_menu_model_set_checked_at(cef_menu_model_t * self, int index, int checked, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, int)) { return callback__(self, index, checked); }
	// int gocef_menu_model_has_accelerator(cef_menu_model_t * self, int commandID, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int)) { return callback__(self, commandID); }
	// int gocef_menu_model_has_accelerator_at(cef_menu_model_t * self, int index, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int)) { return callback__(self, index); }
	// int gocef_menu_model_set_accelerator(cef_menu_model_t * self, int commandID, int keyCode, int shiftPressed, int ctrlPressed, int altPressed, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, int, int, int, int)) { return callback__(self, commandID, keyCode, shiftPressed, ctrlPressed, altPressed); }
	// int gocef_menu_model_set_accelerator_at(cef_menu_model_t * self, int index, int keyCode, int shiftPressed, int ctrlPressed, int altPressed, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, int, int, int, int)) { return callback__(self, index, keyCode, shiftPressed, ctrlPressed, altPressed); }
	// int gocef_menu_model_remove_accelerator(cef_menu_model_t * self, int commandID, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int)) { return callback__(self, commandID); }
	// int gocef_menu_model_remove_accelerator_at(cef_menu_model_t * self, int index, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int)) { return callback__(self, index); }
	// int gocef_menu_model_get_accelerator(cef_menu_model_t * self, int commandID, int * keyCode, int * shiftPressed, int * ctrlPressed, int * altPressed, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, int *, int *, int *, int *)) { return callback__(self, commandID, keyCode, shiftPressed, ctrlPressed, altPressed); }
	// int gocef_menu_model_get_accelerator_at(cef_menu_model_t * self, int index, int * keyCode, int * shiftPressed, int * ctrlPressed, int * altPressed, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, int *, int *, int *, int *)) { return callback__(self, index, keyCode, shiftPressed, ctrlPressed, altPressed); }
	// int gocef_menu_model_set_color(cef_menu_model_t * self, int commandID, cef_menu_color_type_t colorType, cef_color_t color, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, cef_menu_color_type_t, cef_color_t)) { return callback__(self, commandID, colorType, color); }
	// int gocef_menu_model_set_color_at(cef_menu_model_t * self, int index, cef_menu_color_type_t colorType, cef_color_t color, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, cef_menu_color_type_t, cef_color_t)) { return callback__(self, index, colorType, color); }
	// int gocef_menu_model_get_color(cef_menu_model_t * self, int commandID, cef_menu_color_type_t colorType, cef_color_t * color, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, cef_menu_color_type_t, cef_color_t *)) { return callback__(self, commandID, colorType, color); }
	// int gocef_menu_model_get_color_at(cef_menu_model_t * self, int index, cef_menu_color_type_t colorType, cef_color_t * color, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, cef_menu_color_type_t, cef_color_t *)) { return callback__(self, index, colorType, color); }
	// int gocef_menu_model_set_font_list(cef_menu_model_t * self, int commandID, cef_string_t * fontList, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, cef_string_t *)) { return callback__(self, commandID, fontList); }
	// int gocef_menu_model_set_font_list_at(cef_menu_model_t * self, int index, cef_string_t * fontList, int (CEF_CALLBACK *callback__)(cef_menu_model_t *, int, cef_string_t *)) { return callback__(self, index, fontList); }
	"C"
)

// MenuModel (cef_menu_model_t from include/capi/cef_menu_model_capi.h)
// Supports creation and modification of menus. See cef_menu_id_t for the
// command ids that have default implementations. All user-defined command ids
// should be between MENU_ID_USER_FIRST and MENU_ID_USER_LAST. The functions of
// this structure can only be accessed on the browser process the UI thread.
type MenuModel C.cef_menu_model_t

func (d *MenuModel) toNative() *C.cef_menu_model_t {
	return (*C.cef_menu_model_t)(d)
}

// Base (base)
// Base structure.
func (d *MenuModel) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// IsSubMenu (is_sub_menu)
// Returns true (1) if this menu is a submenu.
func (d *MenuModel) IsSubMenu() int32 {
	return int32(C.gocef_menu_model_is_sub_menu(d.toNative(), d.is_sub_menu))
}

// Clear (clear)
// Clears the menu. Returns true (1) on success.
func (d *MenuModel) Clear() int32 {
	return int32(C.gocef_menu_model_clear(d.toNative(), d.clear))
}

// GetCount (get_count)
// Returns the number of items in this menu.
func (d *MenuModel) GetCount() int32 {
	return int32(C.gocef_menu_model_get_count(d.toNative(), d.get_count))
}

// AddSeparator (add_separator)
// Add a separator to the menu. Returns true (1) on success.
func (d *MenuModel) AddSeparator() int32 {
	return int32(C.gocef_menu_model_add_separator(d.toNative(), d.add_separator))
}

// AddItem (add_item)
// Add an item to the menu. Returns true (1) on success.
func (d *MenuModel) AddItem(commandID int32, label string) int32 {
	label_ := C.cef_string_userfree_alloc()
	setCEFStr(label, label_)
	defer func() {
		C.cef_string_userfree_free(label_)
	}()
	return int32(C.gocef_menu_model_add_item(d.toNative(), C.int(commandID), (*C.cef_string_t)(label_), d.add_item))
}

// AddCheckItem (add_check_item)
// Add a check item to the menu. Returns true (1) on success.
func (d *MenuModel) AddCheckItem(commandID int32, label string) int32 {
	label_ := C.cef_string_userfree_alloc()
	setCEFStr(label, label_)
	defer func() {
		C.cef_string_userfree_free(label_)
	}()
	return int32(C.gocef_menu_model_add_check_item(d.toNative(), C.int(commandID), (*C.cef_string_t)(label_), d.add_check_item))
}

// AddRadioItem (add_radio_item)
// Add a radio item to the menu. Only a single item with the specified
// |group_id| can be checked at a time. Returns true (1) on success.
func (d *MenuModel) AddRadioItem(commandID int32, label string, groupID int32) int32 {
	label_ := C.cef_string_userfree_alloc()
	setCEFStr(label, label_)
	defer func() {
		C.cef_string_userfree_free(label_)
	}()
	return int32(C.gocef_menu_model_add_radio_item(d.toNative(), C.int(commandID), (*C.cef_string_t)(label_), C.int(groupID), d.add_radio_item))
}

// AddSubMenu (add_sub_menu)
// Add a sub-menu to the menu. The new sub-menu is returned.
func (d *MenuModel) AddSubMenu(commandID int32, label string) *MenuModel {
	label_ := C.cef_string_userfree_alloc()
	setCEFStr(label, label_)
	defer func() {
		C.cef_string_userfree_free(label_)
	}()
	return (*MenuModel)(C.gocef_menu_model_add_sub_menu(d.toNative(), C.int(commandID), (*C.cef_string_t)(label_), d.add_sub_menu))
}

// InsertSeparatorAt (insert_separator_at)
// Insert a separator in the menu at the specified |index|. Returns true (1)
// on success.
func (d *MenuModel) InsertSeparatorAt(index int32) int32 {
	return int32(C.gocef_menu_model_insert_separator_at(d.toNative(), C.int(index), d.insert_separator_at))
}

// InsertItemAt (insert_item_at)
// Insert an item in the menu at the specified |index|. Returns true (1) on
// success.
func (d *MenuModel) InsertItemAt(index, commandID int32, label string) int32 {
	label_ := C.cef_string_userfree_alloc()
	setCEFStr(label, label_)
	defer func() {
		C.cef_string_userfree_free(label_)
	}()
	return int32(C.gocef_menu_model_insert_item_at(d.toNative(), C.int(index), C.int(commandID), (*C.cef_string_t)(label_), d.insert_item_at))
}

// InsertCheckItemAt (insert_check_item_at)
// Insert a check item in the menu at the specified |index|. Returns true (1)
// on success.
func (d *MenuModel) InsertCheckItemAt(index, commandID int32, label string) int32 {
	label_ := C.cef_string_userfree_alloc()
	setCEFStr(label, label_)
	defer func() {
		C.cef_string_userfree_free(label_)
	}()
	return int32(C.gocef_menu_model_insert_check_item_at(d.toNative(), C.int(index), C.int(commandID), (*C.cef_string_t)(label_), d.insert_check_item_at))
}

// InsertRadioItemAt (insert_radio_item_at)
// Insert a radio item in the menu at the specified |index|. Only a single
// item with the specified |group_id| can be checked at a time. Returns true
// (1) on success.
func (d *MenuModel) InsertRadioItemAt(index, commandID int32, label string, groupID int32) int32 {
	label_ := C.cef_string_userfree_alloc()
	setCEFStr(label, label_)
	defer func() {
		C.cef_string_userfree_free(label_)
	}()
	return int32(C.gocef_menu_model_insert_radio_item_at(d.toNative(), C.int(index), C.int(commandID), (*C.cef_string_t)(label_), C.int(groupID), d.insert_radio_item_at))
}

// InsertSubMenuAt (insert_sub_menu_at)
// Insert a sub-menu in the menu at the specified |index|. The new sub-menu is
// returned.
func (d *MenuModel) InsertSubMenuAt(index, commandID int32, label string) *MenuModel {
	label_ := C.cef_string_userfree_alloc()
	setCEFStr(label, label_)
	defer func() {
		C.cef_string_userfree_free(label_)
	}()
	return (*MenuModel)(C.gocef_menu_model_insert_sub_menu_at(d.toNative(), C.int(index), C.int(commandID), (*C.cef_string_t)(label_), d.insert_sub_menu_at))
}

// Remove (remove)
// Removes the item with the specified |command_id|. Returns true (1) on
// success.
func (d *MenuModel) Remove(commandID int32) int32 {
	return int32(C.gocef_menu_model_remove(d.toNative(), C.int(commandID), d.remove))
}

// RemoveAt (remove_at)
// Removes the item at the specified |index|. Returns true (1) on success.
func (d *MenuModel) RemoveAt(index int32) int32 {
	return int32(C.gocef_menu_model_remove_at(d.toNative(), C.int(index), d.remove_at))
}

// GetIndexOf (get_index_of)
// Returns the index associated with the specified |command_id| or -1 if not
// found due to the command id not existing in the menu.
func (d *MenuModel) GetIndexOf(commandID int32) int32 {
	return int32(C.gocef_menu_model_get_index_of(d.toNative(), C.int(commandID), d.get_index_of))
}

// GetCommandIDAt (get_command_id_at)
// Returns the command id at the specified |index| or -1 if not found due to
// invalid range or the index being a separator.
func (d *MenuModel) GetCommandIDAt(index int32) int32 {
	return int32(C.gocef_menu_model_get_command_id_at(d.toNative(), C.int(index), d.get_command_id_at))
}

// SetCommandIDAt (set_command_id_at)
// Sets the command id at the specified |index|. Returns true (1) on success.
func (d *MenuModel) SetCommandIDAt(index, commandID int32) int32 {
	return int32(C.gocef_menu_model_set_command_id_at(d.toNative(), C.int(index), C.int(commandID), d.set_command_id_at))
}

// GetLabel (get_label)
// Returns the label for the specified |command_id| or NULL if not found.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *MenuModel) GetLabel(commandID int32) string {
	return cefuserfreestrToString(C.gocef_menu_model_get_label(d.toNative(), C.int(commandID), d.get_label))
}

// GetLabelAt (get_label_at)
// Returns the label at the specified |index| or NULL if not found due to
// invalid range or the index being a separator.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *MenuModel) GetLabelAt(index int32) string {
	return cefuserfreestrToString(C.gocef_menu_model_get_label_at(d.toNative(), C.int(index), d.get_label_at))
}

// SetLabel (set_label)
// Sets the label for the specified |command_id|. Returns true (1) on success.
func (d *MenuModel) SetLabel(commandID int32, label string) int32 {
	label_ := C.cef_string_userfree_alloc()
	setCEFStr(label, label_)
	defer func() {
		C.cef_string_userfree_free(label_)
	}()
	return int32(C.gocef_menu_model_set_label(d.toNative(), C.int(commandID), (*C.cef_string_t)(label_), d.set_label))
}

// SetLabelAt (set_label_at)
// Set the label at the specified |index|. Returns true (1) on success.
func (d *MenuModel) SetLabelAt(index int32, label string) int32 {
	label_ := C.cef_string_userfree_alloc()
	setCEFStr(label, label_)
	defer func() {
		C.cef_string_userfree_free(label_)
	}()
	return int32(C.gocef_menu_model_set_label_at(d.toNative(), C.int(index), (*C.cef_string_t)(label_), d.set_label_at))
}

// GetType (get_type)
// Returns the item type for the specified |command_id|.
func (d *MenuModel) GetType(commandID int32) MenuItemType {
	return MenuItemType(C.gocef_menu_model_get_type(d.toNative(), C.int(commandID), d.get_type))
}

// GetTypeAt (get_type_at)
// Returns the item type at the specified |index|.
func (d *MenuModel) GetTypeAt(index int32) MenuItemType {
	return MenuItemType(C.gocef_menu_model_get_type_at(d.toNative(), C.int(index), d.get_type_at))
}

// GetGroupID (get_group_id)
// Returns the group id for the specified |command_id| or -1 if invalid.
func (d *MenuModel) GetGroupID(commandID int32) int32 {
	return int32(C.gocef_menu_model_get_group_id(d.toNative(), C.int(commandID), d.get_group_id))
}

// GetGroupIDAt (get_group_id_at)
// Returns the group id at the specified |index| or -1 if invalid.
func (d *MenuModel) GetGroupIDAt(index int32) int32 {
	return int32(C.gocef_menu_model_get_group_id_at(d.toNative(), C.int(index), d.get_group_id_at))
}

// SetGroupID (set_group_id)
// Sets the group id for the specified |command_id|. Returns true (1) on
// success.
func (d *MenuModel) SetGroupID(commandID, groupID int32) int32 {
	return int32(C.gocef_menu_model_set_group_id(d.toNative(), C.int(commandID), C.int(groupID), d.set_group_id))
}

// SetGroupIDAt (set_group_id_at)
// Sets the group id at the specified |index|. Returns true (1) on success.
func (d *MenuModel) SetGroupIDAt(index, groupID int32) int32 {
	return int32(C.gocef_menu_model_set_group_id_at(d.toNative(), C.int(index), C.int(groupID), d.set_group_id_at))
}

// GetSubMenu (get_sub_menu)
// Returns the submenu for the specified |command_id| or NULL if invalid.
func (d *MenuModel) GetSubMenu(commandID int32) *MenuModel {
	return (*MenuModel)(C.gocef_menu_model_get_sub_menu(d.toNative(), C.int(commandID), d.get_sub_menu))
}

// GetSubMenuAt (get_sub_menu_at)
// Returns the submenu at the specified |index| or NULL if invalid.
func (d *MenuModel) GetSubMenuAt(index int32) *MenuModel {
	return (*MenuModel)(C.gocef_menu_model_get_sub_menu_at(d.toNative(), C.int(index), d.get_sub_menu_at))
}

// IsVisible (is_visible)
// Returns true (1) if the specified |command_id| is visible.
func (d *MenuModel) IsVisible(commandID int32) int32 {
	return int32(C.gocef_menu_model_is_visible(d.toNative(), C.int(commandID), d.is_visible))
}

// IsVisibleAt (is_visible_at)
// Returns true (1) if the specified |index| is visible.
func (d *MenuModel) IsVisibleAt(index int32) int32 {
	return int32(C.gocef_menu_model_is_visible_at(d.toNative(), C.int(index), d.is_visible_at))
}

// SetVisible (set_visible)
// Change the visibility of the specified |command_id|. Returns true (1) on
// success.
func (d *MenuModel) SetVisible(commandID, visible int32) int32 {
	return int32(C.gocef_menu_model_set_visible(d.toNative(), C.int(commandID), C.int(visible), d.set_visible))
}

// SetVisibleAt (set_visible_at)
// Change the visibility at the specified |index|. Returns true (1) on
// success.
func (d *MenuModel) SetVisibleAt(index, visible int32) int32 {
	return int32(C.gocef_menu_model_set_visible_at(d.toNative(), C.int(index), C.int(visible), d.set_visible_at))
}

// IsEnabled (is_enabled)
// Returns true (1) if the specified |command_id| is enabled.
func (d *MenuModel) IsEnabled(commandID int32) int32 {
	return int32(C.gocef_menu_model_is_enabled(d.toNative(), C.int(commandID), d.is_enabled))
}

// IsEnabledAt (is_enabled_at)
// Returns true (1) if the specified |index| is enabled.
func (d *MenuModel) IsEnabledAt(index int32) int32 {
	return int32(C.gocef_menu_model_is_enabled_at(d.toNative(), C.int(index), d.is_enabled_at))
}

// SetEnabled (set_enabled)
// Change the enabled status of the specified |command_id|. Returns true (1)
// on success.
func (d *MenuModel) SetEnabled(commandID, enabled int32) int32 {
	return int32(C.gocef_menu_model_set_enabled(d.toNative(), C.int(commandID), C.int(enabled), d.set_enabled))
}

// SetEnabledAt (set_enabled_at)
// Change the enabled status at the specified |index|. Returns true (1) on
// success.
func (d *MenuModel) SetEnabledAt(index, enabled int32) int32 {
	return int32(C.gocef_menu_model_set_enabled_at(d.toNative(), C.int(index), C.int(enabled), d.set_enabled_at))
}

// IsChecked (is_checked)
// Returns true (1) if the specified |command_id| is checked. Only applies to
// check and radio items.
func (d *MenuModel) IsChecked(commandID int32) int32 {
	return int32(C.gocef_menu_model_is_checked(d.toNative(), C.int(commandID), d.is_checked))
}

// IsCheckedAt (is_checked_at)
// Returns true (1) if the specified |index| is checked. Only applies to check
// and radio items.
func (d *MenuModel) IsCheckedAt(index int32) int32 {
	return int32(C.gocef_menu_model_is_checked_at(d.toNative(), C.int(index), d.is_checked_at))
}

// SetChecked (set_checked)
// Check the specified |command_id|. Only applies to check and radio items.
// Returns true (1) on success.
func (d *MenuModel) SetChecked(commandID, checked int32) int32 {
	return int32(C.gocef_menu_model_set_checked(d.toNative(), C.int(commandID), C.int(checked), d.set_checked))
}

// SetCheckedAt (set_checked_at)
// Check the specified |index|. Only applies to check and radio items. Returns
// true (1) on success.
func (d *MenuModel) SetCheckedAt(index, checked int32) int32 {
	return int32(C.gocef_menu_model_set_checked_at(d.toNative(), C.int(index), C.int(checked), d.set_checked_at))
}

// HasAccelerator (has_accelerator)
// Returns true (1) if the specified |command_id| has a keyboard accelerator
// assigned.
func (d *MenuModel) HasAccelerator(commandID int32) int32 {
	return int32(C.gocef_menu_model_has_accelerator(d.toNative(), C.int(commandID), d.has_accelerator))
}

// HasAcceleratorAt (has_accelerator_at)
// Returns true (1) if the specified |index| has a keyboard accelerator
// assigned.
func (d *MenuModel) HasAcceleratorAt(index int32) int32 {
	return int32(C.gocef_menu_model_has_accelerator_at(d.toNative(), C.int(index), d.has_accelerator_at))
}

// SetAccelerator (set_accelerator)
// Set the keyboard accelerator for the specified |command_id|. |key_code| can
// be any virtual key or character value. Returns true (1) on success.
func (d *MenuModel) SetAccelerator(commandID, keyCode, shiftPressed, ctrlPressed, altPressed int32) int32 {
	return int32(C.gocef_menu_model_set_accelerator(d.toNative(), C.int(commandID), C.int(keyCode), C.int(shiftPressed), C.int(ctrlPressed), C.int(altPressed), d.set_accelerator))
}

// SetAcceleratorAt (set_accelerator_at)
// Set the keyboard accelerator at the specified |index|. |key_code| can be
// any virtual key or character value. Returns true (1) on success.
func (d *MenuModel) SetAcceleratorAt(index, keyCode, shiftPressed, ctrlPressed, altPressed int32) int32 {
	return int32(C.gocef_menu_model_set_accelerator_at(d.toNative(), C.int(index), C.int(keyCode), C.int(shiftPressed), C.int(ctrlPressed), C.int(altPressed), d.set_accelerator_at))
}

// RemoveAccelerator (remove_accelerator)
// Remove the keyboard accelerator for the specified |command_id|. Returns
// true (1) on success.
func (d *MenuModel) RemoveAccelerator(commandID int32) int32 {
	return int32(C.gocef_menu_model_remove_accelerator(d.toNative(), C.int(commandID), d.remove_accelerator))
}

// RemoveAcceleratorAt (remove_accelerator_at)
// Remove the keyboard accelerator at the specified |index|. Returns true (1)
// on success.
func (d *MenuModel) RemoveAcceleratorAt(index int32) int32 {
	return int32(C.gocef_menu_model_remove_accelerator_at(d.toNative(), C.int(index), d.remove_accelerator_at))
}

// GetAccelerator (get_accelerator)
// Retrieves the keyboard accelerator for the specified |command_id|. Returns
// true (1) on success.
func (d *MenuModel) GetAccelerator(commandID int32, keyCode, shiftPressed, ctrlPressed, altPressed *int32) int32 {
	return int32(C.gocef_menu_model_get_accelerator(d.toNative(), C.int(commandID), (*C.int)(keyCode), (*C.int)(shiftPressed), (*C.int)(ctrlPressed), (*C.int)(altPressed), d.get_accelerator))
}

// GetAcceleratorAt (get_accelerator_at)
// Retrieves the keyboard accelerator for the specified |index|. Returns true
// (1) on success.
func (d *MenuModel) GetAcceleratorAt(index int32, keyCode, shiftPressed, ctrlPressed, altPressed *int32) int32 {
	return int32(C.gocef_menu_model_get_accelerator_at(d.toNative(), C.int(index), (*C.int)(keyCode), (*C.int)(shiftPressed), (*C.int)(ctrlPressed), (*C.int)(altPressed), d.get_accelerator_at))
}

// SetColor (set_color)
// Set the explicit color for |command_id| and |color_type| to |color|.
// Specify a |color| value of 0 to remove the explicit color. If no explicit
// color or default color is set for |color_type| then the system color will
// be used. Returns true (1) on success.
func (d *MenuModel) SetColor(commandID int32, colorType MenuColorType, color Color) int32 {
	return int32(C.gocef_menu_model_set_color(d.toNative(), C.int(commandID), C.cef_menu_color_type_t(colorType), C.cef_color_t(color), d.set_color))
}

// SetColorAt (set_color_at)
// Set the explicit color for |command_id| and |index| to |color|. Specify a
// |color| value of 0 to remove the explicit color. Specify an |index| value
// of -1 to set the default color for items that do not have an explicit color
// set. If no explicit color or default color is set for |color_type| then the
// system color will be used. Returns true (1) on success.
func (d *MenuModel) SetColorAt(index int32, colorType MenuColorType, color Color) int32 {
	return int32(C.gocef_menu_model_set_color_at(d.toNative(), C.int(index), C.cef_menu_color_type_t(colorType), C.cef_color_t(color), d.set_color_at))
}

// GetColor (get_color)
// Returns in |color| the color that was explicitly set for |command_id| and
// |color_type|. If a color was not set then 0 will be returned in |color|.
// Returns true (1) on success.
func (d *MenuModel) GetColor(commandID int32, colorType MenuColorType, color *Color) int32 {
	return int32(C.gocef_menu_model_get_color(d.toNative(), C.int(commandID), C.cef_menu_color_type_t(colorType), (*C.cef_color_t)(color), d.get_color))
}

// GetColorAt (get_color_at)
// Returns in |color| the color that was explicitly set for |command_id| and
// |color_type|. Specify an |index| value of -1 to return the default color in
// |color|. If a color was not set then 0 will be returned in |color|. Returns
// true (1) on success.
func (d *MenuModel) GetColorAt(index int32, colorType MenuColorType, color *Color) int32 {
	return int32(C.gocef_menu_model_get_color_at(d.toNative(), C.int(index), C.cef_menu_color_type_t(colorType), (*C.cef_color_t)(color), d.get_color_at))
}

// SetFontList (set_font_list)
// Sets the font list for the specified |command_id|. If |font_list| is NULL
// the system font will be used. Returns true (1) on success. The format is
// "<FONT_FAMILY_LIST>,[STYLES] <SIZE>", where: - FONT_FAMILY_LIST is a comma-
// separated list of font family names, - STYLES is an optional space-
// separated list of style names (case-sensitive
//   "Bold" and "Italic" are supported), and
// - SIZE is an integer font size in pixels with the suffix "px".
//
// Here are examples of valid font description strings: - "Arial, Helvetica,
// Bold Italic 14px" - "Arial, 14px"
func (d *MenuModel) SetFontList(commandID int32, fontList string) int32 {
	fontList_ := C.cef_string_userfree_alloc()
	setCEFStr(fontList, fontList_)
	defer func() {
		C.cef_string_userfree_free(fontList_)
	}()
	return int32(C.gocef_menu_model_set_font_list(d.toNative(), C.int(commandID), (*C.cef_string_t)(fontList_), d.set_font_list))
}

// SetFontListAt (set_font_list_at)
// Sets the font list for the specified |index|. Specify an |index| value of
// -1 to set the default font. If |font_list| is NULL the system font will be
// used. Returns true (1) on success. The format is
// "<FONT_FAMILY_LIST>,[STYLES] <SIZE>", where: - FONT_FAMILY_LIST is a comma-
// separated list of font family names, - STYLES is an optional space-
// separated list of style names (case-sensitive
//   "Bold" and "Italic" are supported), and
// - SIZE is an integer font size in pixels with the suffix "px".
//
// Here are examples of valid font description strings: - "Arial, Helvetica,
// Bold Italic 14px" - "Arial, 14px"
func (d *MenuModel) SetFontListAt(index int32, fontList string) int32 {
	fontList_ := C.cef_string_userfree_alloc()
	setCEFStr(fontList, fontList_)
	defer func() {
		C.cef_string_userfree_free(fontList_)
	}()
	return int32(C.gocef_menu_model_set_font_list_at(d.toNative(), C.int(index), (*C.cef_string_t)(fontList_), d.set_font_list_at))
}
