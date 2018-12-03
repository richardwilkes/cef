// Code generated - DO NOT EDIT.

#include "RequestHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_request_handler_proxy(cef_request_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_before_browse = (void *)&gocef_request_handler_on_before_browse;
	self->on_open_urlfrom_tab = (void *)&gocef_request_handler_on_open_urlfrom_tab;
	self->on_before_resource_load = (void *)&gocef_request_handler_on_before_resource_load;
	self->get_resource_handler = (void *)&gocef_request_handler_get_resource_handler;
	self->on_resource_redirect = (void *)&gocef_request_handler_on_resource_redirect;
	self->on_resource_response = (void *)&gocef_request_handler_on_resource_response;
	self->get_resource_response_filter = (void *)&gocef_request_handler_get_resource_response_filter;
	self->on_resource_load_complete = (void *)&gocef_request_handler_on_resource_load_complete;
	self->get_auth_credentials = (void *)&gocef_request_handler_get_auth_credentials;
	self->can_get_cookies = (void *)&gocef_request_handler_can_get_cookies;
	self->can_set_cookie = (void *)&gocef_request_handler_can_set_cookie;
	self->on_quota_request = (void *)&gocef_request_handler_on_quota_request;
	self->on_protocol_execution = (void *)&gocef_request_handler_on_protocol_execution;
	self->on_certificate_error = (void *)&gocef_request_handler_on_certificate_error;
	self->on_select_client_certificate = (void *)&gocef_request_handler_on_select_client_certificate;
	self->on_plugin_crashed = (void *)&gocef_request_handler_on_plugin_crashed;
	self->on_render_view_ready = (void *)&gocef_request_handler_on_render_view_ready;
	self->on_render_process_terminated = (void *)&gocef_request_handler_on_render_process_terminated;
}
