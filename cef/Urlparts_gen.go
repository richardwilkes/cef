// Code created from "struct.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	"C"
)

// Urlparts (cef_urlparts_t from include/internal/cef_types.h)
// URL component parts.
type Urlparts struct {
	// Spec (spec)
	// The complete URL specification.
	Spec string
	// Scheme (scheme)
	// Scheme component not including the colon (e.g., "http").
	Scheme string
	// Username (username)
	// User name component.
	Username string
	// Password (password)
	// Password component.
	Password string
	// Host (host)
	// Host component. This may be a hostname, an IPv4 address or an IPv6 literal
	// surrounded by square brackets (e.g., "[2001:db8::1]").
	Host string
	// Port (port)
	// Port number component.
	Port string
	// Origin (origin)
	// Origin contains just the scheme, host, and port from a URL. Equivalent to
	// clearing any username and password, replacing the path with a slash, and
	// clearing everything after that. This value will be empty for non-standard
	// URLs.
	Origin string
	// Path (path)
	// Path component including the first slash following the host.
	Path string
	// Query (query)
	// Query string component (i.e., everything following the '?').
	Query string
}

// NewUrlparts creates a new Urlparts.
func NewUrlparts() *Urlparts {
	return &Urlparts{}
}

func (d *Urlparts) toNative(native *C.cef_urlparts_t) *C.cef_urlparts_t {
	if d == nil {
		return nil
	}
	setCEFStr(d.Spec, &native.spec)
	setCEFStr(d.Scheme, &native.scheme)
	setCEFStr(d.Username, &native.username)
	setCEFStr(d.Password, &native.password)
	setCEFStr(d.Host, &native.host)
	setCEFStr(d.Port, &native.port)
	setCEFStr(d.Origin, &native.origin)
	setCEFStr(d.Path, &native.path)
	setCEFStr(d.Query, &native.query)
	return native
}

func (n *C.cef_urlparts_t) toGo() *Urlparts {
	if n == nil {
		return nil
	}
	var d Urlparts
	n.intoGo(&d)
	return &d
}

func (n *C.cef_urlparts_t) intoGo(d *Urlparts) {
	d.Spec = cefstrToString(&n.spec)
	d.Scheme = cefstrToString(&n.scheme)
	d.Username = cefstrToString(&n.username)
	d.Password = cefstrToString(&n.password)
	d.Host = cefstrToString(&n.host)
	d.Port = cefstrToString(&n.port)
	d.Origin = cefstrToString(&n.origin)
	d.Path = cefstrToString(&n.path)
	d.Query = cefstrToString(&n.query)
}
