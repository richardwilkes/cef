#include <stdatomic.h>
#include <stdlib.h>
#include "ref.h"

typedef struct _gocef_refcnt_t {
	volatile atomic_int_fast32_t count;
	int32_t                      dummy; // Just here to pad it out to 8 bytes
} gocef_refcnt_t;

void gocef_refcnt_add(cef_base_ref_counted_t *p) {
	gocef_refcnt_t *ref = (gocef_refcnt_t *)(((void *)p) - sizeof(gocef_refcnt_t));
	atomic_fetch_add(&ref->count, 1);
}

int gocef_refcnt_release(cef_base_ref_counted_t *p) {
	gocef_refcnt_t *ref = (gocef_refcnt_t *)(((void *)p) - sizeof(gocef_refcnt_t));
	if (atomic_fetch_sub(&ref->count, 1) == 0) {
		free(ref);
		return 1;
	}
	return 0;
}

int gocef_refcnt_has_one(cef_base_ref_counted_t *p) {
	gocef_refcnt_t *ref = (gocef_refcnt_t *)(((void *)p) - sizeof(gocef_refcnt_t));
	return atomic_load(&ref->count) == 1;
}

int gocef_refcnt_has_at_least_one(cef_base_ref_counted_t *p) {
	gocef_refcnt_t *ref = (gocef_refcnt_t *)(((void *)p) - sizeof(gocef_refcnt_t));
	return atomic_load(&ref->count) > 0;
}

cef_base_ref_counted_t *gocef_refcnt_alloc(int size) {
	void *p = calloc(1, size + sizeof(gocef_refcnt_t));
	cef_base_ref_counted_t *ret = (cef_base_ref_counted_t *)(p + sizeof(gocef_refcnt_t));
	ret->size = size;
	ret->add_ref = gocef_refcnt_add;
	ret->release = gocef_refcnt_release;
	ret->has_one_ref = gocef_refcnt_has_one;
	ret->has_at_least_one_ref = gocef_refcnt_has_at_least_one;
	return ret;
}
