// Code created from "callback.c.tmpl" - don't edit by hand

#include "ContextMenuHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_context_menu_handler_proxy(cef_context_menu_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_before_context_menu = (void *)&gocef_context_menu_handler_on_before_context_menu;
	self->run_context_menu = (void *)&gocef_context_menu_handler_run_context_menu;
	self->on_context_menu_command = (void *)&gocef_context_menu_handler_on_context_menu_command;
	self->on_context_menu_dismissed = (void *)&gocef_context_menu_handler_on_context_menu_dismissed;
}
