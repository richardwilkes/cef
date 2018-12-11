// Code generated - DO NOT EDIT.

#include "TextfieldDelegate_gen.h"
#include "_cgo_export.h"

void gocef_set_textfield_delegate_proxy(cef_textfield_delegate_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_key_event = (void *)&gocef_textfield_delegate_on_key_event;
	self->on_after_user_action = (void *)&gocef_textfield_delegate_on_after_user_action;
}
