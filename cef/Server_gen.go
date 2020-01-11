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
	// int gocef_server_is_valid_connection(cef_server_t * self, int connectionID, int (CEF_CALLBACK *callback__)(cef_server_t *, int)) { return callback__(self, connectionID); }
	// void gocef_server_send_http200response(cef_server_t * self, int connectionID, cef_string_t * contentType, void * data, size_t dataSize, void (CEF_CALLBACK *callback__)(cef_server_t *, int, cef_string_t *, void *, size_t)) { return callback__(self, connectionID, contentType, data, dataSize); }
	// void gocef_server_send_http404response(cef_server_t * self, int connectionID, void (CEF_CALLBACK *callback__)(cef_server_t *, int)) { return callback__(self, connectionID); }
	// void gocef_server_send_http500response(cef_server_t * self, int connectionID, cef_string_t * errorMessage, void (CEF_CALLBACK *callback__)(cef_server_t *, int, cef_string_t *)) { return callback__(self, connectionID, errorMessage); }
	// void gocef_server_send_http_response(cef_server_t * self, int connectionID, int responseCode, cef_string_t * contentType, int64 contentLength, cef_string_multimap_t extraHeaders, void (CEF_CALLBACK *callback__)(cef_server_t *, int, int, cef_string_t *, int64, cef_string_multimap_t)) { return callback__(self, connectionID, responseCode, contentType, contentLength, extraHeaders); }
	// void gocef_server_send_raw_data(cef_server_t * self, int connectionID, void * data, size_t dataSize, void (CEF_CALLBACK *callback__)(cef_server_t *, int, void *, size_t)) { return callback__(self, connectionID, data, dataSize); }
	// void gocef_server_close_connection(cef_server_t * self, int connectionID, void (CEF_CALLBACK *callback__)(cef_server_t *, int)) { return callback__(self, connectionID); }
	// void gocef_server_send_web_socket_message(cef_server_t * self, int connectionID, void * data, size_t dataSize, void (CEF_CALLBACK *callback__)(cef_server_t *, int, void *, size_t)) { return callback__(self, connectionID, data, dataSize); }
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
func (d *Server) IsValidConnection(connectionID int32) int32 {
	return int32(C.gocef_server_is_valid_connection(d.toNative(), C.int(connectionID), d.is_valid_connection))
}

// SendHTTP200Response (send_http200response)
// Send an HTTP 200 "OK" response to the connection identified by
// |connection_id|. |content_type| is the response content type (e.g.
// "text/html"), |data| is the response content, and |data_size| is the size
// of |data| in bytes. The contents of |data| will be copied. The connection
// will be closed automatically after the response is sent.
func (d *Server) SendHTTP200Response(connectionID int32, contentType string, data unsafe.Pointer, dataSize uint64) {
	contentType_ := C.cef_string_userfree_alloc()
	setCEFStr(contentType, contentType_)
	defer func() {
		C.cef_string_userfree_free(contentType_)
	}()
	C.gocef_server_send_http200response(d.toNative(), C.int(connectionID), (*C.cef_string_t)(contentType_), data, C.size_t(dataSize), d.send_http200response)
}

// SendHTTP404Response (send_http404response)
// Send an HTTP 404 "Not Found" response to the connection identified by
// |connection_id|. The connection will be closed automatically after the
// response is sent.
func (d *Server) SendHTTP404Response(connectionID int32) {
	C.gocef_server_send_http404response(d.toNative(), C.int(connectionID), d.send_http404response)
}

// SendHTTP500Response (send_http500response)
// Send an HTTP 500 "Internal Server Error" response to the connection
// identified by |connection_id|. |error_message| is the associated error
// message. The connection will be closed automatically after the response is
// sent.
func (d *Server) SendHTTP500Response(connectionID int32, errorMessage string) {
	errorMessage_ := C.cef_string_userfree_alloc()
	setCEFStr(errorMessage, errorMessage_)
	defer func() {
		C.cef_string_userfree_free(errorMessage_)
	}()
	C.gocef_server_send_http500response(d.toNative(), C.int(connectionID), (*C.cef_string_t)(errorMessage_), d.send_http500response)
}

// SendHTTPResponse (send_http_response)
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
func (d *Server) SendHTTPResponse(connectionID, responseCode int32, contentType string, contentLength int64, extraHeaders StringMultimap) {
	contentType_ := C.cef_string_userfree_alloc()
	setCEFStr(contentType, contentType_)
	defer func() {
		C.cef_string_userfree_free(contentType_)
	}()
	C.gocef_server_send_http_response(d.toNative(), C.int(connectionID), C.int(responseCode), (*C.cef_string_t)(contentType_), C.int64(contentLength), C.cef_string_multimap_t(extraHeaders), d.send_http_response)
}

// SendRawData (send_raw_data)
// Send raw data directly to the connection identified by |connection_id|.
// |data| is the raw data and |data_size| is the size of |data| in bytes. The
// contents of |data| will be copied. No validation of |data| is performed
// internally so the client should be careful to send the amount indicated by
// the "Content-Length" header, if specified. See SendHttpResponse
// documentation for intended usage.
func (d *Server) SendRawData(connectionID int32, data unsafe.Pointer, dataSize uint64) {
	C.gocef_server_send_raw_data(d.toNative(), C.int(connectionID), data, C.size_t(dataSize), d.send_raw_data)
}

// CloseConnection (close_connection)
// Close the connection identified by |connection_id|. See SendHttpResponse
// documentation for intended usage.
func (d *Server) CloseConnection(connectionID int32) {
	C.gocef_server_close_connection(d.toNative(), C.int(connectionID), d.close_connection)
}

// SendWebSocketMessage (send_web_socket_message)
// Send a WebSocket message to the connection identified by |connection_id|.
// |data| is the response content and |data_size| is the size of |data| in
// bytes. The contents of |data| will be copied. See
// cef_server_handler_t::OnWebSocketRequest documentation for intended usage.
func (d *Server) SendWebSocketMessage(connectionID int32, data unsafe.Pointer, dataSize uint64) {
	C.gocef_server_send_web_socket_message(d.toNative(), C.int(connectionID), data, C.size_t(dataSize), d.send_web_socket_message)
}
