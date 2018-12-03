// Code generated - DO NOT EDIT.

#include "NavigationEntryVisitor_gen.h"
#include "_cgo_export.h"

void gocef_set_navigation_entry_visitor_proxy(cef_navigation_entry_visitor_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->visit = (void *)&gocef_navigation_entry_visitor_visit;
}
