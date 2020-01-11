// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// cef_dom_document_type_t gocef_domdocument_get_type(cef_domdocument_t * self, cef_dom_document_type_t (CEF_CALLBACK *callback__)(cef_domdocument_t *)) { return callback__(self); }
	// cef_domnode_t * gocef_domdocument_get_document(cef_domdocument_t * self, cef_domnode_t * (CEF_CALLBACK *callback__)(cef_domdocument_t *)) { return callback__(self); }
	// cef_domnode_t * gocef_domdocument_get_body(cef_domdocument_t * self, cef_domnode_t * (CEF_CALLBACK *callback__)(cef_domdocument_t *)) { return callback__(self); }
	// cef_domnode_t * gocef_domdocument_get_head(cef_domdocument_t * self, cef_domnode_t * (CEF_CALLBACK *callback__)(cef_domdocument_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_domdocument_get_title(cef_domdocument_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_domdocument_t *)) { return callback__(self); }
	// cef_domnode_t * gocef_domdocument_get_element_by_id(cef_domdocument_t * self, cef_string_t * iD, cef_domnode_t * (CEF_CALLBACK *callback__)(cef_domdocument_t *, cef_string_t *)) { return callback__(self, iD); }
	// cef_domnode_t * gocef_domdocument_get_focused_node(cef_domdocument_t * self, cef_domnode_t * (CEF_CALLBACK *callback__)(cef_domdocument_t *)) { return callback__(self); }
	// int gocef_domdocument_has_selection(cef_domdocument_t * self, int (CEF_CALLBACK *callback__)(cef_domdocument_t *)) { return callback__(self); }
	// int gocef_domdocument_get_selection_start_offset(cef_domdocument_t * self, int (CEF_CALLBACK *callback__)(cef_domdocument_t *)) { return callback__(self); }
	// int gocef_domdocument_get_selection_end_offset(cef_domdocument_t * self, int (CEF_CALLBACK *callback__)(cef_domdocument_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_domdocument_get_selection_as_markup(cef_domdocument_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_domdocument_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_domdocument_get_selection_as_text(cef_domdocument_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_domdocument_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_domdocument_get_base_url(cef_domdocument_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_domdocument_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_domdocument_get_complete_url(cef_domdocument_t * self, cef_string_t * partialURL, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_domdocument_t *, cef_string_t *)) { return callback__(self, partialURL); }
	"C"
)

// Domdocument (cef_domdocument_t from include/capi/cef_dom_capi.h)
// Structure used to represent a DOM document. The functions of this structure
// should only be called on the render process main thread thread.
type Domdocument C.cef_domdocument_t

func (d *Domdocument) toNative() *C.cef_domdocument_t {
	return (*C.cef_domdocument_t)(d)
}

// Base (base)
// Base structure.
func (d *Domdocument) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// GetType (get_type)
// Returns the document type.
func (d *Domdocument) GetType() DOMDocumentType {
	return DOMDocumentType(C.gocef_domdocument_get_type(d.toNative(), d.get_type))
}

// GetDocument (get_document)
// Returns the root document node.
func (d *Domdocument) GetDocument() *Domnode {
	return (*Domnode)(C.gocef_domdocument_get_document(d.toNative(), d.get_document))
}

// GetBody (get_body)
// Returns the BODY node of an HTML document.
func (d *Domdocument) GetBody() *Domnode {
	return (*Domnode)(C.gocef_domdocument_get_body(d.toNative(), d.get_body))
}

// GetHead (get_head)
// Returns the HEAD node of an HTML document.
func (d *Domdocument) GetHead() *Domnode {
	return (*Domnode)(C.gocef_domdocument_get_head(d.toNative(), d.get_head))
}

// GetTitle (get_title)
// Returns the title of an HTML document.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Domdocument) GetTitle() string {
	return cefuserfreestrToString(C.gocef_domdocument_get_title(d.toNative(), d.get_title))
}

// GetElementByID (get_element_by_id)
// Returns the document element with the specified ID value.
func (d *Domdocument) GetElementByID(iD string) *Domnode {
	iD_ := C.cef_string_userfree_alloc()
	setCEFStr(iD, iD_)
	defer func() {
		C.cef_string_userfree_free(iD_)
	}()
	return (*Domnode)(C.gocef_domdocument_get_element_by_id(d.toNative(), (*C.cef_string_t)(iD_), d.get_element_by_id))
}

// GetFocusedNode (get_focused_node)
// Returns the node that currently has keyboard focus.
func (d *Domdocument) GetFocusedNode() *Domnode {
	return (*Domnode)(C.gocef_domdocument_get_focused_node(d.toNative(), d.get_focused_node))
}

// HasSelection (has_selection)
// Returns true (1) if a portion of the document is selected.
func (d *Domdocument) HasSelection() int32 {
	return int32(C.gocef_domdocument_has_selection(d.toNative(), d.has_selection))
}

// GetSelectionStartOffset (get_selection_start_offset)
// Returns the selection offset within the start node.
func (d *Domdocument) GetSelectionStartOffset() int32 {
	return int32(C.gocef_domdocument_get_selection_start_offset(d.toNative(), d.get_selection_start_offset))
}

// GetSelectionEndOffset (get_selection_end_offset)
// Returns the selection offset within the end node.
func (d *Domdocument) GetSelectionEndOffset() int32 {
	return int32(C.gocef_domdocument_get_selection_end_offset(d.toNative(), d.get_selection_end_offset))
}

// GetSelectionAsMarkup (get_selection_as_markup)
// Returns the contents of this selection as markup.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Domdocument) GetSelectionAsMarkup() string {
	return cefuserfreestrToString(C.gocef_domdocument_get_selection_as_markup(d.toNative(), d.get_selection_as_markup))
}

// GetSelectionAsText (get_selection_as_text)
// Returns the contents of this selection as text.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Domdocument) GetSelectionAsText() string {
	return cefuserfreestrToString(C.gocef_domdocument_get_selection_as_text(d.toNative(), d.get_selection_as_text))
}

// GetBaseURL (get_base_url)
// Returns the base URL for the document.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Domdocument) GetBaseURL() string {
	return cefuserfreestrToString(C.gocef_domdocument_get_base_url(d.toNative(), d.get_base_url))
}

// GetCompleteURL (get_complete_url)
// Returns a complete URL based on the document base URL and the specified
// partial URL.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Domdocument) GetCompleteURL(partialURL string) string {
	partialURL_ := C.cef_string_userfree_alloc()
	setCEFStr(partialURL, partialURL_)
	defer func() {
		C.cef_string_userfree_free(partialURL_)
	}()
	return cefuserfreestrToString(C.gocef_domdocument_get_complete_url(d.toNative(), (*C.cef_string_t)(partialURL_), d.get_complete_url))
}
