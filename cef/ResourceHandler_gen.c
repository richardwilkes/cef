// Code created from "callback.c.tmpl" - don't edit by hand

#include "ResourceHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_resource_handler_proxy(cef_resource_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->open = (void *)&gocef_resource_handler_open;
	self->process_request = (void *)&gocef_resource_handler_process_request;
	self->get_response_headers = (void *)&gocef_resource_handler_get_response_headers;
	self->skip = (void *)&gocef_resource_handler_skip;
	self->read = (void *)&gocef_resource_handler_read;
	self->read_response = (void *)&gocef_resource_handler_read_response;
	self->cancel = (void *)&gocef_resource_handler_cancel;
}
