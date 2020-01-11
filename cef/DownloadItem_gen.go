// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// int gocef_download_item_is_valid(cef_download_item_t * self, int (CEF_CALLBACK *callback__)(cef_download_item_t *)) { return callback__(self); }
	// int gocef_download_item_is_in_progress(cef_download_item_t * self, int (CEF_CALLBACK *callback__)(cef_download_item_t *)) { return callback__(self); }
	// int gocef_download_item_is_complete(cef_download_item_t * self, int (CEF_CALLBACK *callback__)(cef_download_item_t *)) { return callback__(self); }
	// int gocef_download_item_is_canceled(cef_download_item_t * self, int (CEF_CALLBACK *callback__)(cef_download_item_t *)) { return callback__(self); }
	// int64 gocef_download_item_get_current_speed(cef_download_item_t * self, int64 (CEF_CALLBACK *callback__)(cef_download_item_t *)) { return callback__(self); }
	// int gocef_download_item_get_percent_complete(cef_download_item_t * self, int (CEF_CALLBACK *callback__)(cef_download_item_t *)) { return callback__(self); }
	// int64 gocef_download_item_get_total_bytes(cef_download_item_t * self, int64 (CEF_CALLBACK *callback__)(cef_download_item_t *)) { return callback__(self); }
	// int64 gocef_download_item_get_received_bytes(cef_download_item_t * self, int64 (CEF_CALLBACK *callback__)(cef_download_item_t *)) { return callback__(self); }
	// cef_time_t gocef_download_item_get_start_time(cef_download_item_t * self, cef_time_t (CEF_CALLBACK *callback__)(cef_download_item_t *)) { return callback__(self); }
	// cef_time_t gocef_download_item_get_end_time(cef_download_item_t * self, cef_time_t (CEF_CALLBACK *callback__)(cef_download_item_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_download_item_get_full_path(cef_download_item_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_download_item_t *)) { return callback__(self); }
	// uint32 gocef_download_item_get_id(cef_download_item_t * self, uint32 (CEF_CALLBACK *callback__)(cef_download_item_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_download_item_get_url(cef_download_item_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_download_item_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_download_item_get_original_url(cef_download_item_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_download_item_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_download_item_get_suggested_file_name(cef_download_item_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_download_item_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_download_item_get_content_disposition(cef_download_item_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_download_item_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_download_item_get_mime_type(cef_download_item_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_download_item_t *)) { return callback__(self); }
	"C"
)

// DownloadItem (cef_download_item_t from include/capi/cef_download_item_capi.h)
// Structure used to represent a download item.
type DownloadItem C.cef_download_item_t

func (d *DownloadItem) toNative() *C.cef_download_item_t {
	return (*C.cef_download_item_t)(d)
}

// Base (base)
// Base structure.
func (d *DownloadItem) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// IsValid (is_valid)
// Returns true (1) if this object is valid. Do not call any other functions
// if this function returns false (0).
func (d *DownloadItem) IsValid() int32 {
	return int32(C.gocef_download_item_is_valid(d.toNative(), d.is_valid))
}

// IsInProgress (is_in_progress)
// Returns true (1) if the download is in progress.
func (d *DownloadItem) IsInProgress() int32 {
	return int32(C.gocef_download_item_is_in_progress(d.toNative(), d.is_in_progress))
}

// IsComplete (is_complete)
// Returns true (1) if the download is complete.
func (d *DownloadItem) IsComplete() int32 {
	return int32(C.gocef_download_item_is_complete(d.toNative(), d.is_complete))
}

// IsCanceled (is_canceled)
// Returns true (1) if the download has been canceled or interrupted.
func (d *DownloadItem) IsCanceled() int32 {
	return int32(C.gocef_download_item_is_canceled(d.toNative(), d.is_canceled))
}

// GetCurrentSpeed (get_current_speed)
// Returns a simple speed estimate in bytes/s.
func (d *DownloadItem) GetCurrentSpeed() int64 {
	return int64(C.gocef_download_item_get_current_speed(d.toNative(), d.get_current_speed))
}

// GetPercentComplete (get_percent_complete)
// Returns the rough percent complete or -1 if the receive total size is
// unknown.
func (d *DownloadItem) GetPercentComplete() int32 {
	return int32(C.gocef_download_item_get_percent_complete(d.toNative(), d.get_percent_complete))
}

// GetTotalBytes (get_total_bytes)
// Returns the total number of bytes.
func (d *DownloadItem) GetTotalBytes() int64 {
	return int64(C.gocef_download_item_get_total_bytes(d.toNative(), d.get_total_bytes))
}

// GetReceivedBytes (get_received_bytes)
// Returns the number of received bytes.
func (d *DownloadItem) GetReceivedBytes() int64 {
	return int64(C.gocef_download_item_get_received_bytes(d.toNative(), d.get_received_bytes))
}

// GetStartTime (get_start_time)
// Returns the time that the download started.
func (d *DownloadItem) GetStartTime() Time {
	cresult__ := C.gocef_download_item_get_start_time(d.toNative(), d.get_start_time)
	var result__ Time
	(&cresult__).intoGo(&result__)
	return result__
}

// GetEndTime (get_end_time)
// Returns the time that the download ended.
func (d *DownloadItem) GetEndTime() Time {
	cresult__ := C.gocef_download_item_get_end_time(d.toNative(), d.get_end_time)
	var result__ Time
	(&cresult__).intoGo(&result__)
	return result__
}

// GetFullPath (get_full_path)
// Returns the full path to the downloaded or downloading file.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *DownloadItem) GetFullPath() string {
	return cefuserfreestrToString(C.gocef_download_item_get_full_path(d.toNative(), d.get_full_path))
}

// GetID (get_id)
// Returns the unique identifier for this download.
func (d *DownloadItem) GetID() uint32 {
	return uint32(C.gocef_download_item_get_id(d.toNative(), d.get_id))
}

// GetURL (get_url)
// Returns the URL.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *DownloadItem) GetURL() string {
	return cefuserfreestrToString(C.gocef_download_item_get_url(d.toNative(), d.get_url))
}

// GetOriginalURL (get_original_url)
// Returns the original URL before any redirections.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *DownloadItem) GetOriginalURL() string {
	return cefuserfreestrToString(C.gocef_download_item_get_original_url(d.toNative(), d.get_original_url))
}

// GetSuggestedFileName (get_suggested_file_name)
// Returns the suggested file name.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *DownloadItem) GetSuggestedFileName() string {
	return cefuserfreestrToString(C.gocef_download_item_get_suggested_file_name(d.toNative(), d.get_suggested_file_name))
}

// GetContentDisposition (get_content_disposition)
// Returns the content disposition.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *DownloadItem) GetContentDisposition() string {
	return cefuserfreestrToString(C.gocef_download_item_get_content_disposition(d.toNative(), d.get_content_disposition))
}

// GetMimeType (get_mime_type)
// Returns the mime type.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *DownloadItem) GetMimeType() string {
	return cefuserfreestrToString(C.gocef_download_item_get_mime_type(d.toNative(), d.get_mime_type))
}
