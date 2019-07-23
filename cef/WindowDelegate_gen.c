// Code created from "callback.c.tmpl" - don't edit by hand

#include "WindowDelegate_gen.h"
#include "_cgo_export.h"

void gocef_set_window_delegate_proxy(cef_window_delegate_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_window_created = (void *)&gocef_window_delegate_on_window_created;
	self->on_window_destroyed = (void *)&gocef_window_delegate_on_window_destroyed;
	self->get_parent_window = (void *)&gocef_window_delegate_get_parent_window;
	self->is_frameless = (void *)&gocef_window_delegate_is_frameless;
	self->can_resize = (void *)&gocef_window_delegate_can_resize;
	self->can_maximize = (void *)&gocef_window_delegate_can_maximize;
	self->can_minimize = (void *)&gocef_window_delegate_can_minimize;
	self->can_close = (void *)&gocef_window_delegate_can_close;
	self->on_accelerator = (void *)&gocef_window_delegate_on_accelerator;
	self->on_key_event = (void *)&gocef_window_delegate_on_key_event;
}
