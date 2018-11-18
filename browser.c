#include "browser.h"

cef_browser_host_t *gocef_get_browser_host(cef_browser_t *browser) {
	return browser->get_host(browser);
}

cef_frame_t *gocef_get_focused_frame(cef_browser_t *browser) {
	return browser->get_focused_frame(browser);
}
