// Code generated - DO NOT EDIT.

#include "DownloadItemCallback_gen.h"
#include "_cgo_export.h"

void gocef_set_download_item_callback_proxy(cef_download_item_callback_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->cancel = (void *)&gocef_download_item_callback_cancel;
	self->pause = (void *)&gocef_download_item_callback_pause;
	self->resume = (void *)&gocef_download_item_callback_resume;
}
