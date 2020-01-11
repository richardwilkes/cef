// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// cef_task_runner_t * gocef_v8context_get_task_runner(cef_v8context_t * self, cef_task_runner_t * (CEF_CALLBACK *callback__)(cef_v8context_t *)) { return callback__(self); }
	// int gocef_v8context_is_valid(cef_v8context_t * self, int (CEF_CALLBACK *callback__)(cef_v8context_t *)) { return callback__(self); }
	// cef_browser_t * gocef_v8context_get_browser(cef_v8context_t * self, cef_browser_t * (CEF_CALLBACK *callback__)(cef_v8context_t *)) { return callback__(self); }
	// cef_frame_t * gocef_v8context_get_frame(cef_v8context_t * self, cef_frame_t * (CEF_CALLBACK *callback__)(cef_v8context_t *)) { return callback__(self); }
	// cef_v8value_t * gocef_v8context_get_global(cef_v8context_t * self, cef_v8value_t * (CEF_CALLBACK *callback__)(cef_v8context_t *)) { return callback__(self); }
	// int gocef_v8context_enter(cef_v8context_t * self, int (CEF_CALLBACK *callback__)(cef_v8context_t *)) { return callback__(self); }
	// int gocef_v8context_exit(cef_v8context_t * self, int (CEF_CALLBACK *callback__)(cef_v8context_t *)) { return callback__(self); }
	// int gocef_v8context_is_same(cef_v8context_t * self, cef_v8context_t * that, int (CEF_CALLBACK *callback__)(cef_v8context_t *, cef_v8context_t *)) { return callback__(self, that); }
	// int gocef_v8context_eval(cef_v8context_t * self, cef_string_t * code, cef_string_t * scriptURL, int startLine, cef_v8value_t ** retval, cef_v8exception_t ** exception, int (CEF_CALLBACK *callback__)(cef_v8context_t *, cef_string_t *, cef_string_t *, int, cef_v8value_t **, cef_v8exception_t **)) { return callback__(self, code, scriptURL, startLine, retval, exception); }
	"C"
)

// V8context (cef_v8context_t from include/capi/cef_v8_capi.h)
// Structure representing a V8 context handle. V8 handles can only be accessed
// from the thread on which they are created. Valid threads for creating a V8
// handle include the render process main thread (TID_RENDERER) and WebWorker
// threads. A task runner for posting tasks on the associated thread can be
// retrieved via the cef_v8context_t::get_task_runner() function.
type V8context C.cef_v8context_t

func (d *V8context) toNative() *C.cef_v8context_t {
	return (*C.cef_v8context_t)(d)
}

// Base (base)
// Base structure.
func (d *V8context) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// GetTaskRunner (get_task_runner)
// Returns the task runner associated with this context. V8 handles can only
// be accessed from the thread on which they are created. This function can be
// called on any render process thread.
func (d *V8context) GetTaskRunner() *TaskRunner {
	return (*TaskRunner)(C.gocef_v8context_get_task_runner(d.toNative(), d.get_task_runner))
}

// IsValid (is_valid)
// Returns true (1) if the underlying handle is valid and it can be accessed
// on the current thread. Do not call any other functions if this function
// returns false (0).
func (d *V8context) IsValid() int32 {
	return int32(C.gocef_v8context_is_valid(d.toNative(), d.is_valid))
}

// GetBrowser (get_browser)
// Returns the browser for this context. This function will return an NULL
// reference for WebWorker contexts.
func (d *V8context) GetBrowser() *Browser {
	return (*Browser)(C.gocef_v8context_get_browser(d.toNative(), d.get_browser))
}

// GetFrame (get_frame)
// Returns the frame for this context. This function will return an NULL
// reference for WebWorker contexts.
func (d *V8context) GetFrame() *Frame {
	return (*Frame)(C.gocef_v8context_get_frame(d.toNative(), d.get_frame))
}

// GetGlobal (get_global)
// Returns the global object for this context. The context must be entered
// before calling this function.
func (d *V8context) GetGlobal() *V8value {
	return (*V8value)(C.gocef_v8context_get_global(d.toNative(), d.get_global))
}

// Enter (enter)
// Enter this context. A context must be explicitly entered before creating a
// V8 Object, Array, Function or Date asynchronously. exit() must be called
// the same number of times as enter() before releasing this context. V8
// objects belong to the context in which they are created. Returns true (1)
// if the scope was entered successfully.
func (d *V8context) Enter() int32 {
	return int32(C.gocef_v8context_enter(d.toNative(), d.enter))
}

// Exit (exit)
// Exit this context. Call this function only after calling enter(). Returns
// true (1) if the scope was exited successfully.
func (d *V8context) Exit() int32 {
	return int32(C.gocef_v8context_exit(d.toNative(), d.exit))
}

// IsSame (is_same)
// Returns true (1) if this object is pointing to the same handle as |that|
// object.
func (d *V8context) IsSame(that *V8context) int32 {
	return int32(C.gocef_v8context_is_same(d.toNative(), that.toNative(), d.is_same))
}

// Eval (eval)
// Execute a string of JavaScript code in this V8 context. The |script_url|
// parameter is the URL where the script in question can be found, if any. The
// |start_line| parameter is the base line number to use for error reporting.
// On success |retval| will be set to the return value, if any, and the
// function will return true (1). On failure |exception| will be set to the
// exception, if any, and the function will return false (0).
func (d *V8context) Eval(code, scriptURL string, startLine int32, retval **V8value, exception **V8exception) int32 {
	code_ := C.cef_string_userfree_alloc()
	setCEFStr(code, code_)
	defer func() {
		C.cef_string_userfree_free(code_)
	}()
	scriptURL_ := C.cef_string_userfree_alloc()
	setCEFStr(scriptURL, scriptURL_)
	defer func() {
		C.cef_string_userfree_free(scriptURL_)
	}()
	retval_ := (*retval).toNative()
	exception_ := (*exception).toNative()
	return int32(C.gocef_v8context_eval(d.toNative(), (*C.cef_string_t)(code_), (*C.cef_string_t)(scriptURL_), C.int(startLine), &retval_, &exception_, d.eval))
}
