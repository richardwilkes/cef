// Code generated - DO NOT EDIT.

#include "BrowserViewDelegate_gen.h"
#include "_cgo_export.h"

void gocef_set_browser_view_delegate_proxy(cef_browser_view_delegate_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_browser_created = (void *)&gocef_browser_view_delegate_on_browser_created;
	self->on_browser_destroyed = (void *)&gocef_browser_view_delegate_on_browser_destroyed;
	self->get_delegate_for_popup_browser_view = (void *)&gocef_browser_view_delegate_get_delegate_for_popup_browser_view;
	self->on_popup_browser_view_created = (void *)&gocef_browser_view_delegate_on_popup_browser_view_created;
}
