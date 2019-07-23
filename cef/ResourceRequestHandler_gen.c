// Code created from "callback.c.tmpl" - don't edit by hand

#include "ResourceRequestHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_resource_request_handler_proxy(cef_resource_request_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->get_cookie_access_filter = (void *)&gocef_resource_request_handler_get_cookie_access_filter;
	self->on_before_resource_load = (void *)&gocef_resource_request_handler_on_before_resource_load;
	self->get_resource_handler = (void *)&gocef_resource_request_handler_get_resource_handler;
	self->on_resource_redirect = (void *)&gocef_resource_request_handler_on_resource_redirect;
	self->on_resource_response = (void *)&gocef_resource_request_handler_on_resource_response;
	self->get_resource_response_filter = (void *)&gocef_resource_request_handler_get_resource_response_filter;
	self->on_resource_load_complete = (void *)&gocef_resource_request_handler_on_resource_load_complete;
	self->on_protocol_execution = (void *)&gocef_resource_request_handler_on_protocol_execution;
}
