// Code generated - DO NOT EDIT.

#include "ServerHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_server_handler_proxy(cef_server_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_server_created = (void *)&gocef_server_handler_on_server_created;
	self->on_server_destroyed = (void *)&gocef_server_handler_on_server_destroyed;
	self->on_client_connected = (void *)&gocef_server_handler_on_client_connected;
	self->on_client_disconnected = (void *)&gocef_server_handler_on_client_disconnected;
	self->on_http_request = (void *)&gocef_server_handler_on_http_request;
	self->on_web_socket_request = (void *)&gocef_server_handler_on_web_socket_request;
	self->on_web_socket_connected = (void *)&gocef_server_handler_on_web_socket_connected;
	self->on_web_socket_message = (void *)&gocef_server_handler_on_web_socket_message;
}
