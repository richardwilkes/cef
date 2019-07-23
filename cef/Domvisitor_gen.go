// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// void gocef_domvisitor_visit(cef_domvisitor_t * self, cef_domdocument_t * document, void (CEF_CALLBACK *callback__)(cef_domvisitor_t *, cef_domdocument_t *)) { return callback__(self, document); }
	"C"
)

// Domvisitor (cef_domvisitor_t from include/capi/cef_dom_capi.h)
// Structure to implement for visiting the DOM. The functions of this structure
// will be called on the render process main thread.
type Domvisitor C.cef_domvisitor_t

func (d *Domvisitor) toNative() *C.cef_domvisitor_t {
	return (*C.cef_domvisitor_t)(d)
}

// Base (base)
// Base structure.
func (d *Domvisitor) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Visit (visit)
// Method executed for visiting the DOM. The document object passed to this
// function represents a snapshot of the DOM at the time this function is
// executed. DOM objects are only valid for the scope of this function. Do not
// keep references to or attempt to access any DOM objects outside the scope
// of this function.
func (d *Domvisitor) Visit(document *Domdocument) {
	C.gocef_domvisitor_visit(d.toNative(), document.toNative(), d.visit)
}
