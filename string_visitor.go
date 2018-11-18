package cef

import (
	// #include "string_visitor.h"
	"C"
	"sync"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

var (
	asyncStringVisitorCallbackLock sync.Mutex
	asyncStringVisitorCallbackMap  = make(map[C.int32_t]func(string))
	asyncStringNextID              C.int32_t
)

func newAsyncStringVisitor(callback func(string)) *C.cef_string_visitor_t {
	asyncStringVisitorCallbackLock.Lock()
	id := asyncStringNextID
	asyncStringNextID++
	asyncStringVisitorCallbackMap[id] = callback
	asyncStringVisitorCallbackLock.Unlock()
	return C.gocef_new_async_string_visitor(id)
}

//export asyncStringVisitorCallback
func asyncStringVisitorCallback(id C.int32_t, str *C.char) {
	asyncStringVisitorCallbackLock.Lock()
	f, ok := asyncStringVisitorCallbackMap[id]
	asyncStringVisitorCallbackLock.Unlock()
	if ok {
		f(C.GoString(str))
	} else {
		jot.Error(errs.Newf("id %d not found for callback", id))
	}
}

//export asyncStringVisitorDisposedCallback
func asyncStringVisitorDisposedCallback(id C.int32_t) {
	asyncStringVisitorCallbackLock.Lock()
	delete(asyncStringVisitorCallbackMap, id)
	asyncStringVisitorCallbackLock.Unlock()
}
