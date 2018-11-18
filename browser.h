#ifndef GOCEF_BROWSER_H_
#define GOCEF_BROWSER_H_
#pragma once

#include "include/capi/cef_browser_capi.h"

cef_browser_host_t *gocef_get_browser_host(cef_browser_t *browser);
cef_frame_t *gocef_get_focused_frame(cef_browser_t *browser);

#endif // GOCEF_BROWSER_H_