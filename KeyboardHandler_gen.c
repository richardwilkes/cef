// Code generated - DO NOT EDIT.

#include "KeyboardHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_keyboard_handler_proxy(cef_keyboard_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_pre_key_event = (void *)&gocef_keyboard_handler_on_pre_key_event;
	self->on_key_event = (void *)&gocef_keyboard_handler_on_key_event;
}
