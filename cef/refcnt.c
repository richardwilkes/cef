// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#include <stdatomic.h>
#include <stdlib.h>
#include "refcnt.h"
#include "_cgo_export.h"

typedef struct _gocef_counted_id_t {
	volatile atomic_int_fast32_t count;
	uint32_t                     id;
} gocef_counted_id_t;

static volatile atomic_uint_fast32_t gocef_next_id = 1;

static gocef_counted_id_t *gocef_refcnt_counted_id_base(cef_base_ref_counted_t *p) {
	return ((gocef_counted_id_t *)(((void *)p) + p->size - sizeof(gocef_counted_id_t)));
}

static volatile atomic_int_fast32_t *gocef_refcnt_count(cef_base_ref_counted_t *p) {
	return &gocef_refcnt_counted_id_base(p)->count;
}

uint32_t gocef_refcnt_id(cef_base_ref_counted_t *p) {
	return gocef_refcnt_counted_id_base(p)->id;
}

static void gocef_refcnt_add(cef_base_ref_counted_t *p) {
	atomic_fetch_add(gocef_refcnt_count(p), 1);
}

static int gocef_refcnt_has_one(cef_base_ref_counted_t *p) {
	return atomic_load(gocef_refcnt_count(p)) == 1;
}

static int gocef_refcnt_has_at_least_one(cef_base_ref_counted_t *p) {
	return atomic_load(gocef_refcnt_count(p)) > 0;
}

static int gocef_refcnt_release(cef_base_ref_counted_t *p) {
	if (atomic_fetch_sub(gocef_refcnt_count(p), 1) == 0) {
		freeObjByID(gocef_refcnt_id(p));
		free(p);
		return 1;
	}
	return 0;
}

cef_base_ref_counted_t *gocef_refcnt_alloc(size_t size) {
	size += 7 + sizeof(gocef_counted_id_t);
	size -= (size % 8);
	cef_base_ref_counted_t *p = (cef_base_ref_counted_t *)calloc(1, size);
	p->size = size;
	p->add_ref = gocef_refcnt_add;
	p->release = gocef_refcnt_release;
	p->has_one_ref = gocef_refcnt_has_one;
	p->has_at_least_one_ref = gocef_refcnt_has_at_least_one;
	gocef_refcnt_counted_id_base(p)->id = (uint32_t)atomic_fetch_add(&gocef_next_id, 1);
	return p;
}
