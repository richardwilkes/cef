// Code created from "class.go.tmpl" - don't edit by hand

package cef

import "unsafe"

import (
	// #include "capi_gen.h"
	// cef_task_runner_t * gocef_server_get_task_runner(cef_server_t * self, cef_task_runner_t * (CEF_CALLBACK *callback__)(cef_server_t *)) { return callback__(self); }
	// void gocef_server_shutdown(cef_server_t * self, void (CEF_CALLBACK *callback__)(cef_server_t *)) { return callback__(self); }
	// int gocef_server_is_running(cef_server_t * self, int (CEF_CALLBACK *callback__)(cef_server_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_server_get_address(cef_server_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_server_t *)) { return callback__(self); }
	// int gocef_server_has_connection(cef_server_t * self, int (CEF_CALLBACK *callback__)(cef_server_t *)) { return callback__(self); }
	// int gocef_server_is_valid_connection(cef_server_t * self, int connection_id, int (CEF_CALLBACK *callback__)(cef_server_t *, int)) { return callback__(self, connection_id); }
	// void gocef_server_send_http200response(cef_server_t * self, int connection_id, cef_string_t * content_type, void * data, size_t data_size, void (CEF_CALLBACK *callback__)(cef_server_t *, int, cef_string_t *, void *, size_t)) { return callback__(self, connection_id, content_type, data, data_size); }
	// void gocef_server_send_http404response(cef_server_t * self, int connection_id, void (CEF_CALLBACK *callback__)(cef_server_t *, int)) { return callback__(self, connection_id); }
	// void gocef_server_send_http500response(cef_server_t * self, int connection_id, cef_string_t * error_message, void (CEF_CALLBACK *callback__)(cef_server_t *, int, cef_string_t *)) { return callback__(self, connection_id, error_message); }
	// void gocef_server_send_http_response(cef_server_t * self, int connection_id, int response_code, cef_string_t * content_type, int64 content_length, cef_string_multimap_t extra_headers, void (CEF_CALLBACK *callback__)(cef_server_t *, int, int, cef_string_t *, int64, cef_string_multimap_t)) { return callback__(self, connection_id, response_code, content_type, content_length, extra_headers); }
	// void gocef_server_send_raw_data(cef_server_t * self, int connection_id, void * data, size_t data_size, void (CEF_CALLBACK *callback__)(cef_server_t *, int, void *, size_t)) { return callback__(self, connection_id, data, data_size); }
	// void gocef_server_close_connection(cef_server_t * self, int connection_id, void (CEF_CALLBACK *callback__)(cef_server_t *, int)) { return callback__(self, connection_id); }
	// void gocef_server_send_web_socket_message(cef_server_t * self, int connection_id, void * data, size_t data_size, void (CEF_CALLBACK *callback__)(cef_server_t *, int, void *, size_t)) { return callback__(self, connection_id, data, data_size); }
	"C"
)

// Server (cef_server_t from include/capi/cef_server_capi.h)
// Structure representing a server that supports HTTP and WebSocket requests.
// Server capacity is limited and is intended to handle only a small number of
// simultaneous connections (e.g. for communicating between applications on
// localhost). The functions of this structure are safe to call from any thread
// in the brower process unless otherwise indicated.
type Server C.cef_server_t

func (d *Server) toNative() *C.cef_server_t {
	return (*C.cef_server_t)(d)
}

// Base (base)
// Base structure.
func (d *Server) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// GetTaskRunner (get_task_runner)
// Returns the task runner for the dedicated server thread.
func (d *Server) GetTaskRunner() *TaskRunner {
	return (*TaskRunner)(C.gocef_server_get_task_runner(d.toNative(), d.get_task_runner))
}

// Shutdown (shutdown)
// Stop the server and shut down the dedicated server thread. See
// cef_server_handler_t::OnServerCreated documentation for a description of
// server lifespan.
func (d *Server) Shutdown() {
	C.gocef_server_shutdown(d.toNative(), d.shutdown)
}

// IsRunning (is_running)
// Returns true (1) if the server is currently running and accepting incoming
// connections. See cef_server_handler_t::OnServerCreated documentation for a
// description of server lifespan. This function must be called on the
// dedicated server thread.
func (d *Server) IsRunning() int32 {
	return int32(C.gocef_server_is_running(d.toNative(), d.is_running))
}

// GetAddress (get_address)
// Returns the server address including the port number.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Server) GetAddress() string {
	return cefuserfreestrToString(C.gocef_server_get_address(d.toNative(), d.get_address))
}

// HasConnection (has_connection)
// Returns true (1) if the server currently has a connection. This function
// must be called on the dedicated server thread.
func (d *Server) HasConnection() int32 {
	return int32(C.gocef_server_has_connection(d.toNative(), d.has_connection))
}

// IsValidConnection (is_valid_connection)
// Returns true (1) if |connection_id| represents a valid connection. This
// function must be called on the dedicated server thread.
func (d *Server) IsValidConnection(connection_id int32) int32 {
	return int32(C.gocef_server_is_valid_connection(d.toNative(), C.int(connection_id), d.is_valid_connection))
}

// SendHttp200response (send_http200response)
// Send an HTTP 200 "OK" response to the connection identified by
// |connection_id|. |content_type| is the response content type (e.g.
// "text/html"), |data| is the response content, and |data_size| is the size
// of |data| in bytes. The contents of |data| will be copied. The connection
// will be closed automatically after the response is sent.
func (d *Server) SendHttp200response(connection_id int32, content_type string, data unsafe.Pointer, data_size uint64) {
	content_type_ := C.cef_string_userfree_alloc()
	setCEFStr(content_type, content_type_)
	defer func() {
		C.cef_string_userfree_free(content_type_)
	}()
	C.gocef_server_send_http200response(d.toNative(), C.int(connection_id), (*C.cef_string_t)(content_type_), data, C.size_t(data_size), d.send_http200response)
}

// SendHttp404response (send_http404response)
// Send an HTTP 404 "Not Found" response to the connection identified by
// |connection_id|. The connection will be closed automatically after the
// response is sent.
func (d *Server) SendHttp404response(connection_id int32) {
	C.gocef_server_send_http404response(d.toNative(), C.int(connection_id), d.send_http404response)
}

// SendHttp500response (send_http500response)
// Send an HTTP 500 "Internal Server Error" response to the connection
// identified by |connection_id|. |error_message| is the associated error
// message. The connection will be closed automatically after the response is
// sent.
func (d *Server) SendHttp500response(connection_id int32, error_message string) {
	error_message_ := C.cef_string_userfree_alloc()
	setCEFStr(error_message, error_message_)
	defer func() {
		C.cef_string_userfree_free(error_message_)
	}()
	C.gocef_server_send_http500response(d.toNative(), C.int(connection_id), (*C.cef_string_t)(error_message_), d.send_http500response)
}

// SendHttpResponse (send_http_response)
// Send a custom HTTP response to the connection identified by
// |connection_id|. |response_code| is the HTTP response code sent in the
// status line (e.g. 200), |content_type| is the response content type sent as
// the "Content-Type" header (e.g. "text/html"), |content_length| is the
// expected content length, and |extra_headers| is the map of extra response
// headers. If |content_length| is >= 0 then the "Content-Length" header will
// be sent. If |content_length| is 0 then no content is expected and the
// connection will be closed automatically after the response is sent. If
// |content_length| is < 0 then no "Content-Length" header will be sent and
// the client will continue reading until the connection is closed. Use the
// SendRawData function to send the content, if applicable, and call
// CloseConnection after all content has been sent.
func (d *Server) SendHttpResponse(connection_id, response_code int32, content_type string, content_length int64, extra_headers StringMultimap) {
	content_type_ := C.cef_string_userfree_alloc()
	setCEFStr(content_type, content_type_)
	defer func() {
		C.cef_string_userfree_free(content_type_)
	}()
	C.gocef_server_send_http_response(d.toNative(), C.int(connection_id), C.int(response_code), (*C.cef_string_t)(content_type_), C.int64(content_length), C.cef_string_multimap_t(extra_headers), d.send_http_response)
}

// SendRawData (send_raw_data)
// Send raw data directly to the connection identified by |connection_id|.
// |data| is the raw data and |data_size| is the size of |data| in bytes. The
// contents of |data| will be copied. No validation of |data| is performed
// internally so the client should be careful to send the amount indicated by
// the "Content-Length" header, if specified. See SendHttpResponse
// documentation for intended usage.
func (d *Server) SendRawData(connection_id int32, data unsafe.Pointer, data_size uint64) {
	C.gocef_server_send_raw_data(d.toNative(), C.int(connection_id), data, C.size_t(data_size), d.send_raw_data)
}

// CloseConnection (close_connection)
// Close the connection identified by |connection_id|. See SendHttpResponse
// documentation for intended usage.
func (d *Server) CloseConnection(connection_id int32) {
	C.gocef_server_close_connection(d.toNative(), C.int(connection_id), d.close_connection)
}

// SendWebSocketMessage (send_web_socket_message)
// Send a WebSocket message to the connection identified by |connection_id|.
// |data| is the response content and |data_size| is the size of |data| in
// bytes. The contents of |data| will be copied. See
// cef_server_handler_t::OnWebSocketRequest documentation for intended usage.
func (d *Server) SendWebSocketMessage(connection_id int32, data unsafe.Pointer, data_size uint64) {
	C.gocef_server_send_web_socket_message(d.toNative(), C.int(connection_id), data, C.size_t(data_size), d.send_web_socket_message)
}
