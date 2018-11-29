// Code generated - DO NOT EDIT.

#include "DownloadHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_download_handler_proxy(cef_download_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_before_download = (void *)&gocef_download_handler_on_before_download;
	self->on_download_updated = (void *)&gocef_download_handler_on_download_updated;
}
