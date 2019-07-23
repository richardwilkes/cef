// Code created from "callback.c.tmpl" - don't edit by hand

#include "DragHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_drag_handler_proxy(cef_drag_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_drag_enter = (void *)&gocef_drag_handler_on_drag_enter;
	self->on_draggable_regions_changed = (void *)&gocef_drag_handler_on_draggable_regions_changed;
}
