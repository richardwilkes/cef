// Code created from "callback.c.tmpl" - don't edit by hand

#include "DownloadImageCallback_gen.h"
#include "_cgo_export.h"

void gocef_set_download_image_callback_proxy(cef_download_image_callback_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_download_image_finished = (void *)&gocef_download_image_callback_on_download_image_finished;
}
