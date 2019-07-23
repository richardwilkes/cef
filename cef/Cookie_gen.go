// Code created from "struct.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	"C"
)

// Cookie (cef_cookie_t from include/internal/cef_types.h)
// Cookie information.
type Cookie struct {
	// Name (name)
	// The cookie name.
	Name string
	// Value (value)
	// The cookie value.
	Value string
	// Domain (domain)
	// If |domain| is empty a host cookie will be created instead of a domain
	// cookie. Domain cookies are stored with a leading "." and are visible to
	// sub-domains whereas host cookies are not.
	Domain string
	// Path (path)
	// If |path| is non-empty only URLs at or below the path will get the cookie
	// value.
	Path string
	// Secure (secure)
	// If |secure| is true the cookie will only be sent for HTTPS requests.
	Secure int32
	// Httponly (httponly)
	// If |httponly| is true the cookie will only be sent for HTTP requests.
	Httponly int32
	// Creation (creation)
	// The cookie creation date. This is automatically populated by the system on
	// cookie creation.
	Creation Time
	// LastAccess (last_access)
	// The cookie last access date. This is automatically populated by the system
	// on access.
	LastAccess Time
	// HasExpires (has_expires)
	// The cookie expiration date is only valid if |has_expires| is true.
	HasExpires int32
	// Expires (expires)
	Expires Time
}

// NewCookie creates a new Cookie.
func NewCookie() *Cookie {
	return &Cookie{}
}

func (d *Cookie) toNative(native *C.cef_cookie_t) *C.cef_cookie_t {
	if d == nil {
		return nil
	}
	setCEFStr(d.Name, &native.name)
	setCEFStr(d.Value, &native.value)
	setCEFStr(d.Domain, &native.domain)
	setCEFStr(d.Path, &native.path)
	native.secure = C.int(d.Secure)
	native.httponly = C.int(d.Httponly)
	d.Creation.toNative(&native.creation)
	d.LastAccess.toNative(&native.last_access)
	native.has_expires = C.int(d.HasExpires)
	d.Expires.toNative(&native.expires)
	return native
}

func (n *C.cef_cookie_t) toGo() *Cookie {
	if n == nil {
		return nil
	}
	var d Cookie
	n.intoGo(&d)
	return &d
}

func (n *C.cef_cookie_t) intoGo(d *Cookie) {
	d.Name = cefstrToString(&n.name)
	d.Value = cefstrToString(&n.value)
	d.Domain = cefstrToString(&n.domain)
	d.Path = cefstrToString(&n.path)
	d.Secure = int32(n.secure)
	d.Httponly = int32(n.httponly)
	n.creation.intoGo(&d.Creation)
	n.last_access.intoGo(&d.LastAccess)
	d.HasExpires = int32(n.has_expires)
	n.expires.intoGo(&d.Expires)
}
