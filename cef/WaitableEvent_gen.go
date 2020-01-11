// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// void gocef_waitable_event_reset(cef_waitable_event_t * self, void (CEF_CALLBACK *callback__)(cef_waitable_event_t *)) { return callback__(self); }
	// void gocef_waitable_event_signal(cef_waitable_event_t * self, void (CEF_CALLBACK *callback__)(cef_waitable_event_t *)) { return callback__(self); }
	// int gocef_waitable_event_is_signaled(cef_waitable_event_t * self, int (CEF_CALLBACK *callback__)(cef_waitable_event_t *)) { return callback__(self); }
	// void gocef_waitable_event_wait(cef_waitable_event_t * self, void (CEF_CALLBACK *callback__)(cef_waitable_event_t *)) { return callback__(self); }
	// int gocef_waitable_event_timed_wait(cef_waitable_event_t * self, int64 maxMs, int (CEF_CALLBACK *callback__)(cef_waitable_event_t *, int64)) { return callback__(self, maxMs); }
	"C"
)

// WaitableEvent (cef_waitable_event_t from include/capi/cef_waitable_event_capi.h)
// WaitableEvent is a thread synchronization tool that allows one thread to wait
// for another thread to finish some work. This is equivalent to using a
// Lock+ConditionVariable to protect a simple boolean value. However, using
// WaitableEvent in conjunction with a Lock to wait for a more complex state
// change (e.g., for an item to be added to a queue) is not recommended. In that
// case consider using a ConditionVariable instead of a WaitableEvent. It is
// safe to create and/or signal a WaitableEvent from any thread. Blocking on a
// WaitableEvent by calling the *wait() functions is not allowed on the browser
// process UI or IO threads.
type WaitableEvent C.cef_waitable_event_t

func (d *WaitableEvent) toNative() *C.cef_waitable_event_t {
	return (*C.cef_waitable_event_t)(d)
}

// Base (base)
// Base structure.
func (d *WaitableEvent) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Reset (reset)
// Put the event in the un-signaled state.
func (d *WaitableEvent) Reset() {
	C.gocef_waitable_event_reset(d.toNative(), d.reset)
}

// Signal (signal)
// Put the event in the signaled state. This causes any thread blocked on Wait
// to be woken up.
func (d *WaitableEvent) Signal() {
	C.gocef_waitable_event_signal(d.toNative(), d.signal)
}

// IsSignaled (is_signaled)
// Returns true (1) if the event is in the signaled state, else false (0). If
// the event was created with |automatic_reset| set to true (1) then calling
// this function will also cause a reset.
func (d *WaitableEvent) IsSignaled() int32 {
	return int32(C.gocef_waitable_event_is_signaled(d.toNative(), d.is_signaled))
}

// Wait (wait)
// Wait indefinitely for the event to be signaled. This function will not
// return until after the call to signal() has completed. This function cannot
// be called on the browser process UI or IO threads.
func (d *WaitableEvent) Wait() {
	C.gocef_waitable_event_wait(d.toNative(), d.wait)
}

// TimedWait (timed_wait)
// Wait up to |max_ms| milliseconds for the event to be signaled. Returns true
// (1) if the event was signaled. A return value of false (0) does not
// necessarily mean that |max_ms| was exceeded. This function will not return
// until after the call to signal() has completed. This function cannot be
// called on the browser process UI or IO threads.
func (d *WaitableEvent) TimedWait(maxMs int64) int32 {
	return int32(C.gocef_waitable_event_timed_wait(d.toNative(), C.int64(maxMs), d.timed_wait))
}
