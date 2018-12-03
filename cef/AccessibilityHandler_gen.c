// Code generated - DO NOT EDIT.

#include "AccessibilityHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_accessibility_handler_proxy(cef_accessibility_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_accessibility_tree_change = (void *)&gocef_accessibility_handler_on_accessibility_tree_change;
	self->on_accessibility_location_change = (void *)&gocef_accessibility_handler_on_accessibility_location_change;
}
