#include "string_visitor.h"
#include "ref.h"
#include "_cgo_export.h"

void gocef_async_string_visitor_visit(cef_string_visitor_t *self, const cef_string_t *str) {
	cef_string_userfree_utf8_t conv = cef_string_userfree_utf8_alloc();
	cef_string_to_utf8(str->str, str->length, (cef_string_utf8_t *)conv);
	asyncStringVisitorCallback(((gocef_async_string_visitor_t *)self)->id, conv->str);
	cef_string_userfree_utf8_free(conv);
}

int gocef_async_string_visitor_release(cef_base_ref_counted_t *self) {
	gocef_async_string_visitor_t *visitor = (gocef_async_string_visitor_t *)self;
	int32_t id = visitor->id;
	int result = visitor->release(self);
	if (result) {
		asyncStringVisitorDisposedCallback(id);
	}
	return result;
}

cef_string_visitor_t *gocef_new_async_string_visitor(int32_t id) {
	gocef_async_string_visitor_t *visitor = (gocef_async_string_visitor_t *)gocef_refcnt_alloc(sizeof(gocef_async_string_visitor_t));
	visitor->id = id;
	visitor->visitor.visit = gocef_async_string_visitor_visit;
	visitor->release = visitor->visitor.base.release;
	visitor->visitor.base.release = gocef_async_string_visitor_release;
	return &visitor->visitor;
}
