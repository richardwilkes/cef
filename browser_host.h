#ifndef BROWSER_HOST_H_
#define BROWSER_HOST_H_
#pragma once

#include "include/capi/cef_browser_capi.h"

cef_window_handle_t gocef_windowhandle_browserhost(cef_browser_host_t *host, cef_window_handle_t (CEF_CALLBACK* get_window_handle)(cef_browser_host_t *));

#endif // BROWSER_HOST_H_
