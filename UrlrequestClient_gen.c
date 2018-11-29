#include "UrlrequestClient_gen.h"
#include "_cgo_export.h"

void gocef_set_urlrequest_client_proxy(cef_urlrequest_client_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_request_complete = (void *)&gocef_urlrequest_client_on_request_complete;
	self->on_upload_progress = (void *)&gocef_urlrequest_client_on_upload_progress;
	self->on_download_progress = (void *)&gocef_urlrequest_client_on_download_progress;
	self->on_download_data = (void *)&gocef_urlrequest_client_on_download_data;
	self->get_auth_credentials = (void *)&gocef_urlrequest_client_get_auth_credentials;
}
