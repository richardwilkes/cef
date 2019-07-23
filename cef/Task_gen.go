// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

import (
	// #include "Task_gen.h"
	"C"
)

// TaskProxy defines methods required for using Task.
type TaskProxy interface {
	Execute(self *Task)
}

// Task (cef_task_t from include/capi/cef_task_capi.h)
// Implement this structure for asynchronous task execution. If the task is
// posted successfully and if the associated message loop is still running then
// the execute() function will be called on the target thread. If the task fails
// to post then the task object may be destroyed on the source thread instead of
// the target thread. For this reason be cautious when performing work in the
// task object destructor.
type Task C.cef_task_t

// NewTask creates a new Task with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewTask(proxy TaskProxy) *Task {
	result := (*Task)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_task_t, proxy)))
	if proxy != nil {
		C.gocef_set_task_proxy(result.toNative())
	}
	return result
}

func (d *Task) toNative() *C.cef_task_t {
	return (*C.cef_task_t)(d)
}

func lookupTaskProxy(obj *BaseRefCounted) TaskProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(TaskProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type TaskProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *Task) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Execute (execute)
// Method that will be executed on the target thread.
func (d *Task) Execute() {
	lookupTaskProxy(d.Base()).Execute(d)
}

//nolint:gocritic
//export gocef_task_execute
func gocef_task_execute(self *C.cef_task_t) {
	me__ := (*Task)(self)
	proxy__ := lookupTaskProxy(me__.Base())
	proxy__.Execute(me__)
}
