// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#ifndef GOCEF_REFCNT_H_
#define GOCEF_REFCNT_H_
#pragma once

#include "include/capi/cef_base_capi.h"

cef_base_ref_counted_t *gocef_refcnt_alloc(size_t size);
uint32_t gocef_refcnt_id(cef_base_ref_counted_t *p);

#endif // GOCEF_REFCNT_H_
