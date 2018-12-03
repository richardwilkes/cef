// Code generated - DO NOT EDIT.

#include "FocusHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_focus_handler_proxy(cef_focus_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_take_focus = (void *)&gocef_focus_handler_on_take_focus;
	self->on_set_focus = (void *)&gocef_focus_handler_on_set_focus;
	self->on_got_focus = (void *)&gocef_focus_handler_on_got_focus;
}
