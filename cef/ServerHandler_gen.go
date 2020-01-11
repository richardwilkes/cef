// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

import (
	// #include "ServerHandler_gen.h"
	"C"
)

// ServerHandlerProxy defines methods required for using ServerHandler.
type ServerHandlerProxy interface {
	OnServerCreated(self *ServerHandler, server *Server)
	OnServerDestroyed(self *ServerHandler, server *Server)
	OnClientConnected(self *ServerHandler, server *Server, connectionID int32)
	OnClientDisconnected(self *ServerHandler, server *Server, connectionID int32)
	OnHTTPRequest(self *ServerHandler, server *Server, connectionID int32, clientAddress string, request *Request)
	OnWebSocketRequest(self *ServerHandler, server *Server, connectionID int32, clientAddress string, request *Request, callback *Callback)
	OnWebSocketConnected(self *ServerHandler, server *Server, connectionID int32)
	OnWebSocketMessage(self *ServerHandler, server *Server, connectionID int32, data unsafe.Pointer, dataSize uint64)
}

// ServerHandler (cef_server_handler_t from include/capi/cef_server_capi.h)
// Implement this structure to handle HTTP server requests. A new thread will be
// created for each cef_server_t::CreateServer call (the "dedicated server
// thread"), and the functions of this structure will be called on that thread.
// It is therefore recommended to use a different cef_server_handler_t instance
// for each cef_server_t::CreateServer call to avoid thread safety issues in the
// cef_server_handler_t implementation.
type ServerHandler C.cef_server_handler_t

// NewServerHandler creates a new ServerHandler with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewServerHandler(proxy ServerHandlerProxy) *ServerHandler {
	result := (*ServerHandler)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_server_handler_t, proxy)))
	if proxy != nil {
		C.gocef_set_server_handler_proxy(result.toNative())
	}
	return result
}

func (d *ServerHandler) toNative() *C.cef_server_handler_t {
	return (*C.cef_server_handler_t)(d)
}

func lookupServerHandlerProxy(obj *BaseRefCounted) ServerHandlerProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(ServerHandlerProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type ServerHandlerProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *ServerHandler) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnServerCreated (on_server_created)
// Called when |server| is created. If the server was started successfully
// then cef_server_t::IsRunning will return true (1). The server will continue
// running until cef_server_t::Shutdown is called, after which time
// OnServerDestroyed will be called. If the server failed to start then
// OnServerDestroyed will be called immediately after this function returns.
func (d *ServerHandler) OnServerCreated(server *Server) {
	lookupServerHandlerProxy(d.Base()).OnServerCreated(d, server)
}

//nolint:gocritic
//export gocef_server_handler_on_server_created
func gocef_server_handler_on_server_created(self *C.cef_server_handler_t, server *C.cef_server_t) {
	me__ := (*ServerHandler)(self)
	proxy__ := lookupServerHandlerProxy(me__.Base())
	proxy__.OnServerCreated(me__, (*Server)(server))
}

// OnServerDestroyed (on_server_destroyed)
// Called when |server| is destroyed. The server thread will be stopped after
// this function returns. The client should release any references to |server|
// when this function is called. See OnServerCreated documentation for a
// description of server lifespan.
func (d *ServerHandler) OnServerDestroyed(server *Server) {
	lookupServerHandlerProxy(d.Base()).OnServerDestroyed(d, server)
}

//nolint:gocritic
//export gocef_server_handler_on_server_destroyed
func gocef_server_handler_on_server_destroyed(self *C.cef_server_handler_t, server *C.cef_server_t) {
	me__ := (*ServerHandler)(self)
	proxy__ := lookupServerHandlerProxy(me__.Base())
	proxy__.OnServerDestroyed(me__, (*Server)(server))
}

// OnClientConnected (on_client_connected)
// Called when a client connects to |server|. |connection_id| uniquely
// identifies the connection. Each call to this function will have a matching
// call to OnClientDisconnected.
func (d *ServerHandler) OnClientConnected(server *Server, connectionID int32) {
	lookupServerHandlerProxy(d.Base()).OnClientConnected(d, server, connectionID)
}

//nolint:gocritic
//export gocef_server_handler_on_client_connected
func gocef_server_handler_on_client_connected(self *C.cef_server_handler_t, server *C.cef_server_t, connectionID C.int) {
	me__ := (*ServerHandler)(self)
	proxy__ := lookupServerHandlerProxy(me__.Base())
	proxy__.OnClientConnected(me__, (*Server)(server), int32(connectionID))
}

// OnClientDisconnected (on_client_disconnected)
// Called when a client disconnects from |server|. |connection_id| uniquely
// identifies the connection. The client should release any data associated
// with |connection_id| when this function is called and |connection_id|
// should no longer be passed to cef_server_t functions. Disconnects can
// originate from either the client or the server. For example, the server
// will disconnect automatically after a cef_server_t::SendHttpXXXResponse
// function is called.
func (d *ServerHandler) OnClientDisconnected(server *Server, connectionID int32) {
	lookupServerHandlerProxy(d.Base()).OnClientDisconnected(d, server, connectionID)
}

//nolint:gocritic
//export gocef_server_handler_on_client_disconnected
func gocef_server_handler_on_client_disconnected(self *C.cef_server_handler_t, server *C.cef_server_t, connectionID C.int) {
	me__ := (*ServerHandler)(self)
	proxy__ := lookupServerHandlerProxy(me__.Base())
	proxy__.OnClientDisconnected(me__, (*Server)(server), int32(connectionID))
}

// OnHTTPRequest (on_http_request)
// Called when |server| receives an HTTP request. |connection_id| uniquely
// identifies the connection, |client_address| is the requesting IPv4 or IPv6
// client address including port number, and |request| contains the request
// contents (URL, function, headers and optional POST data). Call cef_server_t
// functions either synchronously or asynchronusly to send a response.
func (d *ServerHandler) OnHTTPRequest(server *Server, connectionID int32, clientAddress string, request *Request) {
	lookupServerHandlerProxy(d.Base()).OnHTTPRequest(d, server, connectionID, clientAddress, request)
}

//nolint:gocritic
//export gocef_server_handler_on_http_request
func gocef_server_handler_on_http_request(self *C.cef_server_handler_t, server *C.cef_server_t, connectionID C.int, clientAddress *C.cef_string_t, request *C.cef_request_t) {
	me__ := (*ServerHandler)(self)
	proxy__ := lookupServerHandlerProxy(me__.Base())
	clientAddress_ := cefstrToString(clientAddress)
	proxy__.OnHTTPRequest(me__, (*Server)(server), int32(connectionID), clientAddress_, (*Request)(request))
}

// OnWebSocketRequest (on_web_socket_request)
// Called when |server| receives a WebSocket request. |connection_id| uniquely
// identifies the connection, |client_address| is the requesting IPv4 or IPv6
// client address including port number, and |request| contains the request
// contents (URL, function, headers and optional POST data). Execute
// |callback| either synchronously or asynchronously to accept or decline the
// WebSocket connection. If the request is accepted then OnWebSocketConnected
// will be called after the WebSocket has connected and incoming messages will
// be delivered to the OnWebSocketMessage callback. If the request is declined
// then the client will be disconnected and OnClientDisconnected will be
// called. Call the cef_server_t::SendWebSocketMessage function after
// receiving the OnWebSocketConnected callback to respond with WebSocket
// messages.
func (d *ServerHandler) OnWebSocketRequest(server *Server, connectionID int32, clientAddress string, request *Request, callback *Callback) {
	lookupServerHandlerProxy(d.Base()).OnWebSocketRequest(d, server, connectionID, clientAddress, request, callback)
}

//nolint:gocritic
//export gocef_server_handler_on_web_socket_request
func gocef_server_handler_on_web_socket_request(self *C.cef_server_handler_t, server *C.cef_server_t, connectionID C.int, clientAddress *C.cef_string_t, request *C.cef_request_t, callback *C.cef_callback_t) {
	me__ := (*ServerHandler)(self)
	proxy__ := lookupServerHandlerProxy(me__.Base())
	clientAddress_ := cefstrToString(clientAddress)
	proxy__.OnWebSocketRequest(me__, (*Server)(server), int32(connectionID), clientAddress_, (*Request)(request), (*Callback)(callback))
}

// OnWebSocketConnected (on_web_socket_connected)
// Called after the client has accepted the WebSocket connection for |server|
// and |connection_id| via the OnWebSocketRequest callback. See
// OnWebSocketRequest documentation for intended usage.
func (d *ServerHandler) OnWebSocketConnected(server *Server, connectionID int32) {
	lookupServerHandlerProxy(d.Base()).OnWebSocketConnected(d, server, connectionID)
}

//nolint:gocritic
//export gocef_server_handler_on_web_socket_connected
func gocef_server_handler_on_web_socket_connected(self *C.cef_server_handler_t, server *C.cef_server_t, connectionID C.int) {
	me__ := (*ServerHandler)(self)
	proxy__ := lookupServerHandlerProxy(me__.Base())
	proxy__.OnWebSocketConnected(me__, (*Server)(server), int32(connectionID))
}

// OnWebSocketMessage (on_web_socket_message)
// Called when |server| receives an WebSocket message. |connection_id|
// uniquely identifies the connection, |data| is the message content and
// |data_size| is the size of |data| in bytes. Do not keep a reference to
// |data| outside of this function. See OnWebSocketRequest documentation for
// intended usage.
func (d *ServerHandler) OnWebSocketMessage(server *Server, connectionID int32, data unsafe.Pointer, dataSize uint64) {
	lookupServerHandlerProxy(d.Base()).OnWebSocketMessage(d, server, connectionID, data, dataSize)
}

//nolint:gocritic
//export gocef_server_handler_on_web_socket_message
func gocef_server_handler_on_web_socket_message(self *C.cef_server_handler_t, server *C.cef_server_t, connectionID C.int, data unsafe.Pointer, dataSize C.size_t) {
	me__ := (*ServerHandler)(self)
	proxy__ := lookupServerHandlerProxy(me__.Base())
	proxy__.OnWebSocketMessage(me__, (*Server)(server), int32(connectionID), data, uint64(dataSize))
}
