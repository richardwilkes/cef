// Code generated - DO NOT EDIT.

#include "ResourceHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_resource_handler_proxy(cef_resource_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->process_request = (void *)&gocef_resource_handler_process_request;
	self->get_response_headers = (void *)&gocef_resource_handler_get_response_headers;
	self->read_response = (void *)&gocef_resource_handler_read_response;
	self->can_get_cookie = (void *)&gocef_resource_handler_can_get_cookie;
	self->can_set_cookie = (void *)&gocef_resource_handler_can_set_cookie;
	self->cancel = (void *)&gocef_resource_handler_cancel;
}
