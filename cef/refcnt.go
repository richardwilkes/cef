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
