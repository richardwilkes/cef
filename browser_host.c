#include "browser_host.h"

cef_window_handle_t gocef_windowhandle_browserhost(cef_browser_host_t *host, cef_window_handle_t(CEF_CALLBACK *callback)(cef_browser_host_t *)) {
	return callback(host);
}
