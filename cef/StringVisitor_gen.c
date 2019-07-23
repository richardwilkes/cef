// Code created from "callback.c.tmpl" - don't edit by hand

#include "StringVisitor_gen.h"
#include "_cgo_export.h"

void gocef_set_string_visitor_proxy(cef_string_visitor_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->visit = (void *)&gocef_string_visitor_visit;
}
