// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package cef

import "sync"

import (
	// #include "refcnt.h"
	"C"
)

var (
	refLock sync.Mutex
	refMap  = make(map[C.uint32_t]interface{})
)

func newRefCntObj(size uint, proxy interface{}) *BaseRefCounted {
	cobj := C.gocef_refcnt_alloc(C.size_t(size))
	id := C.gocef_refcnt_id(cobj)
	obj := (*BaseRefCounted)(cobj)
	refLock.Lock()
	refMap[id] = proxy
	refLock.Unlock()
	return obj
}

func lookupProxy(base *BaseRefCounted) (proxy interface{}, exists bool) {
	id := C.gocef_refcnt_id(base.toNative())
	refLock.Lock()
	proxy, exists = refMap[id]
	refLock.Unlock()
	return proxy, exists
}

//nolint: gocritic
//export freeObjByID
func freeObjByID(cid C.uint32_t) {
	refLock.Lock()
	delete(refMap, cid)
	refLock.Unlock()
}
