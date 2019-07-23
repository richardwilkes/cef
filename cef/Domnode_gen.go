// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// cef_dom_node_type_t gocef_domnode_get_type(cef_domnode_t * self, cef_dom_node_type_t (CEF_CALLBACK *callback__)(cef_domnode_t *)) { return callback__(self); }
	// int gocef_domnode_is_text(cef_domnode_t * self, int (CEF_CALLBACK *callback__)(cef_domnode_t *)) { return callback__(self); }
	// int gocef_domnode_is_element(cef_domnode_t * self, int (CEF_CALLBACK *callback__)(cef_domnode_t *)) { return callback__(self); }
	// int gocef_domnode_is_editable(cef_domnode_t * self, int (CEF_CALLBACK *callback__)(cef_domnode_t *)) { return callback__(self); }
	// int gocef_domnode_is_form_control_element(cef_domnode_t * self, int (CEF_CALLBACK *callback__)(cef_domnode_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_domnode_get_form_control_element_type(cef_domnode_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_domnode_t *)) { return callback__(self); }
	// int gocef_domnode_is_same(cef_domnode_t * self, cef_domnode_t * that, int (CEF_CALLBACK *callback__)(cef_domnode_t *, cef_domnode_t *)) { return callback__(self, that); }
	// cef_string_userfree_t gocef_domnode_get_name(cef_domnode_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_domnode_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_domnode_get_value(cef_domnode_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_domnode_t *)) { return callback__(self); }
	// int gocef_domnode_set_value(cef_domnode_t * self, cef_string_t * value, int (CEF_CALLBACK *callback__)(cef_domnode_t *, cef_string_t *)) { return callback__(self, value); }
	// cef_string_userfree_t gocef_domnode_get_as_markup(cef_domnode_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_domnode_t *)) { return callback__(self); }
	// cef_domdocument_t * gocef_domnode_get_document(cef_domnode_t * self, cef_domdocument_t * (CEF_CALLBACK *callback__)(cef_domnode_t *)) { return callback__(self); }
	// cef_domnode_t * gocef_domnode_get_parent(cef_domnode_t * self, cef_domnode_t * (CEF_CALLBACK *callback__)(cef_domnode_t *)) { return callback__(self); }
	// cef_domnode_t * gocef_domnode_get_previous_sibling(cef_domnode_t * self, cef_domnode_t * (CEF_CALLBACK *callback__)(cef_domnode_t *)) { return callback__(self); }
	// cef_domnode_t * gocef_domnode_get_next_sibling(cef_domnode_t * self, cef_domnode_t * (CEF_CALLBACK *callback__)(cef_domnode_t *)) { return callback__(self); }
	// int gocef_domnode_has_children(cef_domnode_t * self, int (CEF_CALLBACK *callback__)(cef_domnode_t *)) { return callback__(self); }
	// cef_domnode_t * gocef_domnode_get_first_child(cef_domnode_t * self, cef_domnode_t * (CEF_CALLBACK *callback__)(cef_domnode_t *)) { return callback__(self); }
	// cef_domnode_t * gocef_domnode_get_last_child(cef_domnode_t * self, cef_domnode_t * (CEF_CALLBACK *callback__)(cef_domnode_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_domnode_get_element_tag_name(cef_domnode_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_domnode_t *)) { return callback__(self); }
	// int gocef_domnode_has_element_attributes(cef_domnode_t * self, int (CEF_CALLBACK *callback__)(cef_domnode_t *)) { return callback__(self); }
	// int gocef_domnode_has_element_attribute(cef_domnode_t * self, cef_string_t * attrName, int (CEF_CALLBACK *callback__)(cef_domnode_t *, cef_string_t *)) { return callback__(self, attrName); }
	// cef_string_userfree_t gocef_domnode_get_element_attribute(cef_domnode_t * self, cef_string_t * attrName, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_domnode_t *, cef_string_t *)) { return callback__(self, attrName); }
	// void gocef_domnode_get_element_attributes(cef_domnode_t * self, cef_string_map_t attrMap, void (CEF_CALLBACK *callback__)(cef_domnode_t *, cef_string_map_t)) { return callback__(self, attrMap); }
	// int gocef_domnode_set_element_attribute(cef_domnode_t * self, cef_string_t * attrName, cef_string_t * value, int (CEF_CALLBACK *callback__)(cef_domnode_t *, cef_string_t *, cef_string_t *)) { return callback__(self, attrName, value); }
	// cef_string_userfree_t gocef_domnode_get_element_inner_text(cef_domnode_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_domnode_t *)) { return callback__(self); }
	// cef_rect_t gocef_domnode_get_element_bounds(cef_domnode_t * self, cef_rect_t (CEF_CALLBACK *callback__)(cef_domnode_t *)) { return callback__(self); }
	"C"
)

// Domnode (cef_domnode_t from include/capi/cef_dom_capi.h)
// Structure used to represent a DOM node. The functions of this structure
// should only be called on the render process main thread.
type Domnode C.cef_domnode_t

func (d *Domnode) toNative() *C.cef_domnode_t {
	return (*C.cef_domnode_t)(d)
}

// Base (base)
// Base structure.
func (d *Domnode) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// GetType (get_type)
// Returns the type for this node.
func (d *Domnode) GetType() DOMNodeType {
	return DOMNodeType(C.gocef_domnode_get_type(d.toNative(), d.get_type))
}

// IsText (is_text)
// Returns true (1) if this is a text node.
func (d *Domnode) IsText() int32 {
	return int32(C.gocef_domnode_is_text(d.toNative(), d.is_text))
}

// IsElement (is_element)
// Returns true (1) if this is an element node.
func (d *Domnode) IsElement() int32 {
	return int32(C.gocef_domnode_is_element(d.toNative(), d.is_element))
}

// IsEditable (is_editable)
// Returns true (1) if this is an editable node.
func (d *Domnode) IsEditable() int32 {
	return int32(C.gocef_domnode_is_editable(d.toNative(), d.is_editable))
}

// IsFormControlElement (is_form_control_element)
// Returns true (1) if this is a form control element node.
func (d *Domnode) IsFormControlElement() int32 {
	return int32(C.gocef_domnode_is_form_control_element(d.toNative(), d.is_form_control_element))
}

// GetFormControlElementType (get_form_control_element_type)
// Returns the type of this form control element node.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Domnode) GetFormControlElementType() string {
	return cefuserfreestrToString(C.gocef_domnode_get_form_control_element_type(d.toNative(), d.get_form_control_element_type))
}

// IsSame (is_same)
// Returns true (1) if this object is pointing to the same handle as |that|
// object.
func (d *Domnode) IsSame(that *Domnode) int32 {
	return int32(C.gocef_domnode_is_same(d.toNative(), that.toNative(), d.is_same))
}

// GetName (get_name)
// Returns the name of this node.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Domnode) GetName() string {
	return cefuserfreestrToString(C.gocef_domnode_get_name(d.toNative(), d.get_name))
}

// GetValue (get_value)
// Returns the value of this node.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Domnode) GetValue() string {
	return cefuserfreestrToString(C.gocef_domnode_get_value(d.toNative(), d.get_value))
}

// SetValue (set_value)
// Set the value of this node. Returns true (1) on success.
func (d *Domnode) SetValue(value string) int32 {
	value_ := C.cef_string_userfree_alloc()
	setCEFStr(value, value_)
	defer func() {
		C.cef_string_userfree_free(value_)
	}()
	return int32(C.gocef_domnode_set_value(d.toNative(), (*C.cef_string_t)(value_), d.set_value))
}

// GetAsMarkup (get_as_markup)
// Returns the contents of this node as markup.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Domnode) GetAsMarkup() string {
	return cefuserfreestrToString(C.gocef_domnode_get_as_markup(d.toNative(), d.get_as_markup))
}

// GetDocument (get_document)
// Returns the document associated with this node.
func (d *Domnode) GetDocument() *Domdocument {
	return (*Domdocument)(C.gocef_domnode_get_document(d.toNative(), d.get_document))
}

// GetParent (get_parent)
// Returns the parent node.
func (d *Domnode) GetParent() *Domnode {
	return (*Domnode)(C.gocef_domnode_get_parent(d.toNative(), d.get_parent))
}

// GetPreviousSibling (get_previous_sibling)
// Returns the previous sibling node.
func (d *Domnode) GetPreviousSibling() *Domnode {
	return (*Domnode)(C.gocef_domnode_get_previous_sibling(d.toNative(), d.get_previous_sibling))
}

// GetNextSibling (get_next_sibling)
// Returns the next sibling node.
func (d *Domnode) GetNextSibling() *Domnode {
	return (*Domnode)(C.gocef_domnode_get_next_sibling(d.toNative(), d.get_next_sibling))
}

// HasChildren (has_children)
// Returns true (1) if this node has child nodes.
func (d *Domnode) HasChildren() int32 {
	return int32(C.gocef_domnode_has_children(d.toNative(), d.has_children))
}

// GetFirstChild (get_first_child)
// Return the first child node.
func (d *Domnode) GetFirstChild() *Domnode {
	return (*Domnode)(C.gocef_domnode_get_first_child(d.toNative(), d.get_first_child))
}

// GetLastChild (get_last_child)
// Returns the last child node.
func (d *Domnode) GetLastChild() *Domnode {
	return (*Domnode)(C.gocef_domnode_get_last_child(d.toNative(), d.get_last_child))
}

// GetElementTagName (get_element_tag_name)
// Returns the tag name of this element.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Domnode) GetElementTagName() string {
	return cefuserfreestrToString(C.gocef_domnode_get_element_tag_name(d.toNative(), d.get_element_tag_name))
}

// HasElementAttributes (has_element_attributes)
// Returns true (1) if this element has attributes.
func (d *Domnode) HasElementAttributes() int32 {
	return int32(C.gocef_domnode_has_element_attributes(d.toNative(), d.has_element_attributes))
}

// HasElementAttribute (has_element_attribute)
// Returns true (1) if this element has an attribute named |attrName|.
func (d *Domnode) HasElementAttribute(attrName string) int32 {
	attrName_ := C.cef_string_userfree_alloc()
	setCEFStr(attrName, attrName_)
	defer func() {
		C.cef_string_userfree_free(attrName_)
	}()
	return int32(C.gocef_domnode_has_element_attribute(d.toNative(), (*C.cef_string_t)(attrName_), d.has_element_attribute))
}

// GetElementAttribute (get_element_attribute)
// Returns the element attribute named |attrName|.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Domnode) GetElementAttribute(attrName string) string {
	attrName_ := C.cef_string_userfree_alloc()
	setCEFStr(attrName, attrName_)
	defer func() {
		C.cef_string_userfree_free(attrName_)
	}()
	return cefuserfreestrToString(C.gocef_domnode_get_element_attribute(d.toNative(), (*C.cef_string_t)(attrName_), d.get_element_attribute))
}

// GetElementAttributes (get_element_attributes)
// Returns a map of all element attributes.
func (d *Domnode) GetElementAttributes(attrMap StringMap) {
	C.gocef_domnode_get_element_attributes(d.toNative(), C.cef_string_map_t(attrMap), d.get_element_attributes)
}

// SetElementAttribute (set_element_attribute)
// Set the value for the element attribute named |attrName|. Returns true (1)
// on success.
func (d *Domnode) SetElementAttribute(attrName, value string) int32 {
	attrName_ := C.cef_string_userfree_alloc()
	setCEFStr(attrName, attrName_)
	defer func() {
		C.cef_string_userfree_free(attrName_)
	}()
	value_ := C.cef_string_userfree_alloc()
	setCEFStr(value, value_)
	defer func() {
		C.cef_string_userfree_free(value_)
	}()
	return int32(C.gocef_domnode_set_element_attribute(d.toNative(), (*C.cef_string_t)(attrName_), (*C.cef_string_t)(value_), d.set_element_attribute))
}

// GetElementInnerText (get_element_inner_text)
// Returns the inner text of the element.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Domnode) GetElementInnerText() string {
	return cefuserfreestrToString(C.gocef_domnode_get_element_inner_text(d.toNative(), d.get_element_inner_text))
}

// GetElementBounds (get_element_bounds)
// Returns the bounds of the element.
func (d *Domnode) GetElementBounds() Rect {
	cresult__ := C.gocef_domnode_get_element_bounds(d.toNative(), d.get_element_bounds)
	var result__ Rect
	(&cresult__).intoGo(&result__)
	return result__
}
