// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package main

import (
	"strings"

	"github.com/richardwilkes/toolbox/txt"
)

var (
	nameExceptions = []string{
		"acl", "api", "ascii", "com", "cpu", "css", "ct", "dns", "dom", "eof",
		"guid", "html", "http", "http11", "http200", "http404", "http500",
		"https", "id", "io", "ip", "js", "json", "lhs", "pdf", "qps", "ram",
		"rhs", "rpc", "sla", "smtp", "sql", "ssh", "ssl", "ssl2", "ssl3",
		"st", "tcp", "tid", "tls", "tls1", "tls13", "tp", "ts", "tt", "ttl",
		"udp", "ui", "uid", "uuid", "uri", "url", "utf16", "utf16le",
		"utf16be", "utf8", "v8", "vm", "x509", "xml", "xmpp", "xsrf", "xss",
	}
	paramRenames = []string{
		"chan", "defer", "error", "fallthrough", "func", "go", "import",
		"interface", "map", "package", "range", "select", "string", "type",
		"var",
	}
)

func translateName(in string) string {
	in = strings.TrimPrefix(strings.TrimSpace(in), "CEF_")
	if strings.HasSuffix(in, "response") {
		in = strings.TrimSuffix(in, "response") + "_response"
	}
	in = strings.ToLower(txt.ToSnakeCase(in))
	for _, s := range nameExceptions {
		if s == in {
			in = strings.ToUpper(s)
			break
		}
		upper := strings.ToUpper(s)
		if strings.HasPrefix(in, s+"_") {
			in = upper + "_" + in[len(s)+1:]
		}
		if strings.HasSuffix(in, "_"+s) {
			in = in[:len(in)-(1+len(s))] + "_" + upper
		}
		for {
			if i := strings.Index(in, "_"+s+"_"); i == -1 {
				break
			} else {
				in = in[:i+1] + upper + in[i+1+len(s):]
			}
		}
	}
	in = txt.ToCamelCase(in)
	return in
}

func adjustedParamName(name string) string {
	for _, one := range paramRenames {
		if name == one {
			return name + "_r"
		}
	}
	return txt.FirstToLower(translateName(name))
}
