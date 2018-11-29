#include "V8arrayBufferReleaseCallback_gen.h"
#include "_cgo_export.h"

void gocef_set_v8array_buffer_release_callback_proxy(cef_v8array_buffer_release_callback_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->release_buffer = (void *)&gocef_v8array_buffer_release_callback_release_buffer;
}
