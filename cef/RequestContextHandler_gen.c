// Code created from "callback.c.tmpl" - don't edit by hand

#include "RequestContextHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_request_context_handler_proxy(cef_request_context_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_request_context_initialized = (void *)&gocef_request_context_handler_on_request_context_initialized;
	self->on_before_plugin_load = (void *)&gocef_request_context_handler_on_before_plugin_load;
	self->get_resource_request_handler = (void *)&gocef_request_context_handler_get_resource_request_handler;
}
