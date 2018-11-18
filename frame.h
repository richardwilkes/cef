#ifndef GOCEF_FRAME_H_
#define GOCEF_FRAME_H_
#pragma once

#include "include/capi/cef_browser_capi.h"
#include "include/capi/cef_frame_capi.h"

int gocef_call_int_frame(cef_frame_t *frame, int (CEF_CALLBACK *callback)(cef_frame_t *));
int64 gocef_call_int64_frame(cef_frame_t *frame, int64 (CEF_CALLBACK *callback)(cef_frame_t *));
cef_string_userfree_t gocef_call_string_frame(cef_frame_t *frame, cef_string_userfree_t (CEF_CALLBACK *callback)(cef_frame_t *));
void gocef_call_void_frame(cef_frame_t *frame, void (CEF_CALLBACK *callback)(cef_frame_t *));
void gocef_call_void_frame_string_visitor(cef_frame_t *frame, cef_string_visitor_t *visitor, void (CEF_CALLBACK *callback)(cef_frame_t *, cef_string_visitor_t *));
void gocef_call_void_frame_string(cef_frame_t *frame, cef_string_t *str, void (CEF_CALLBACK *callback)(cef_frame_t *, const cef_string_t *));
void gocef_call_void_frame_string_string(cef_frame_t *frame, cef_string_t *str1, cef_string_t *str2, void (CEF_CALLBACK *callback)(cef_frame_t *, const cef_string_t *, const cef_string_t *));
void gocef_call_void_frame_string_string_int(cef_frame_t *frame, cef_string_t *str1, cef_string_t *str2, int num, void (CEF_CALLBACK *callback)(cef_frame_t *, const cef_string_t *, const cef_string_t *, int));
cef_frame_t *gocef_call_frame_frame(cef_frame_t *frame, cef_frame_t *(CEF_CALLBACK *callback)(cef_frame_t *));
cef_browser_t *gocef_call_browser_frame(cef_frame_t *frame, cef_browser_t *(CEF_CALLBACK *callback)(cef_frame_t *));

#endif // GOCEF_FRAME_H_