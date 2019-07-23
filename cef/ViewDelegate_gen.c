// Code created from "callback.c.tmpl" - don't edit by hand

#include "ViewDelegate_gen.h"
#include "_cgo_export.h"

void gocef_set_view_delegate_proxy(cef_view_delegate_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->get_preferred_size = (void *)&gocef_view_delegate_get_preferred_size;
	self->get_minimum_size = (void *)&gocef_view_delegate_get_minimum_size;
	self->get_maximum_size = (void *)&gocef_view_delegate_get_maximum_size;
	self->get_height_for_width = (void *)&gocef_view_delegate_get_height_for_width;
	self->on_parent_view_changed = (void *)&gocef_view_delegate_on_parent_view_changed;
	self->on_child_view_changed = (void *)&gocef_view_delegate_on_child_view_changed;
	self->on_focus = (void *)&gocef_view_delegate_on_focus;
	self->on_blur = (void *)&gocef_view_delegate_on_blur;
}
