// Code created from "callback.c.tmpl" - don't edit by hand

#include "AudioHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_audio_handler_proxy(cef_audio_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_audio_stream_started = (void *)&gocef_audio_handler_on_audio_stream_started;
	self->on_audio_stream_packet = (void *)&gocef_audio_handler_on_audio_stream_packet;
	self->on_audio_stream_stopped = (void *)&gocef_audio_handler_on_audio_stream_stopped;
}
