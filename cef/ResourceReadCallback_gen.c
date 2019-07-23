// Code created from "callback.c.tmpl" - don't edit by hand

#include "ResourceReadCallback_gen.h"
#include "_cgo_export.h"

void gocef_set_resource_read_callback_proxy(cef_resource_read_callback_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->cont = (void *)&gocef_resource_read_callback_cont;
}
