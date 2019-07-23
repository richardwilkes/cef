// Code created from "callback.c.tmpl" - don't edit by hand

#include "MenuButtonDelegate_gen.h"
#include "_cgo_export.h"

void gocef_set_menu_button_delegate_proxy(cef_menu_button_delegate_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_menu_button_pressed = (void *)&gocef_menu_button_delegate_on_menu_button_pressed;
}
