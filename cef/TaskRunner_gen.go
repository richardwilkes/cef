// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// int gocef_task_runner_is_same(cef_task_runner_t * self, cef_task_runner_t * that, int (CEF_CALLBACK *callback__)(cef_task_runner_t *, cef_task_runner_t *)) { return callback__(self, that); }
	// int gocef_task_runner_belongs_to_current_thread(cef_task_runner_t * self, int (CEF_CALLBACK *callback__)(cef_task_runner_t *)) { return callback__(self); }
	// int gocef_task_runner_belongs_to_thread(cef_task_runner_t * self, cef_thread_id_t threadID, int (CEF_CALLBACK *callback__)(cef_task_runner_t *, cef_thread_id_t)) { return callback__(self, threadID); }
	// int gocef_task_runner_post_task(cef_task_runner_t * self, cef_task_t * task, int (CEF_CALLBACK *callback__)(cef_task_runner_t *, cef_task_t *)) { return callback__(self, task); }
	// int gocef_task_runner_post_delayed_task(cef_task_runner_t * self, cef_task_t * task, int64 delayMs, int (CEF_CALLBACK *callback__)(cef_task_runner_t *, cef_task_t *, int64)) { return callback__(self, task, delayMs); }
	"C"
)

// TaskRunner (cef_task_runner_t from include/capi/cef_task_capi.h)
// Structure that asynchronously executes tasks on the associated thread. It is
// safe to call the functions of this structure on any thread.
//
// CEF maintains multiple internal threads that are used for handling different
// types of tasks in different processes. The cef_thread_id_t definitions in
// cef_types.h list the common CEF threads. Task runners are also available for
// other CEF threads as appropriate (for example, V8 WebWorker threads).
type TaskRunner C.cef_task_runner_t

func (d *TaskRunner) toNative() *C.cef_task_runner_t {
	return (*C.cef_task_runner_t)(d)
}

// Base (base)
// Base structure.
func (d *TaskRunner) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// IsSame (is_same)
// Returns true (1) if this object is pointing to the same task runner as
// |that| object.
func (d *TaskRunner) IsSame(that *TaskRunner) int32 {
	return int32(C.gocef_task_runner_is_same(d.toNative(), that.toNative(), d.is_same))
}

// BelongsToCurrentThread (belongs_to_current_thread)
// Returns true (1) if this task runner belongs to the current thread.
func (d *TaskRunner) BelongsToCurrentThread() int32 {
	return int32(C.gocef_task_runner_belongs_to_current_thread(d.toNative(), d.belongs_to_current_thread))
}

// BelongsToThread (belongs_to_thread)
// Returns true (1) if this task runner is for the specified CEF thread.
func (d *TaskRunner) BelongsToThread(threadID ThreadID) int32 {
	return int32(C.gocef_task_runner_belongs_to_thread(d.toNative(), C.cef_thread_id_t(threadID), d.belongs_to_thread))
}

// PostTask (post_task)
// Post a task for execution on the thread associated with this task runner.
// Execution will occur asynchronously.
func (d *TaskRunner) PostTask(task *Task) int32 {
	return int32(C.gocef_task_runner_post_task(d.toNative(), task.toNative(), d.post_task))
}

// PostDelayedTask (post_delayed_task)
// Post a task for delayed execution on the thread associated with this task
// runner. Execution will occur asynchronously. Delayed tasks are not
// supported on V8 WebWorker threads and will be executed without the
// specified delay.
func (d *TaskRunner) PostDelayedTask(task *Task, delayMs int64) int32 {
	return int32(C.gocef_task_runner_post_delayed_task(d.toNative(), task.toNative(), C.int64(delayMs), d.post_delayed_task))
}
