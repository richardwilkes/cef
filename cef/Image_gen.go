// Code created from "class.go.tmpl" - don't edit by hand

package cef

import "unsafe"

import (
	// #include "capi_gen.h"
	// int gocef_image_is_empty(cef_image_t * self, int (CEF_CALLBACK *callback__)(cef_image_t *)) { return callback__(self); }
	// int gocef_image_is_same(cef_image_t * self, cef_image_t * that, int (CEF_CALLBACK *callback__)(cef_image_t *, cef_image_t *)) { return callback__(self, that); }
	// int gocef_image_add_bitmap(cef_image_t * self, float scale_factor, int pixel_width, int pixel_height, cef_color_type_t color_type, cef_alpha_type_t alpha_type, void * pixel_data, size_t pixel_data_size, int (CEF_CALLBACK *callback__)(cef_image_t *, float, int, int, cef_color_type_t, cef_alpha_type_t, void *, size_t)) { return callback__(self, scale_factor, pixel_width, pixel_height, color_type, alpha_type, pixel_data, pixel_data_size); }
	// int gocef_image_add_png(cef_image_t * self, float scale_factor, void * png_data, size_t png_data_size, int (CEF_CALLBACK *callback__)(cef_image_t *, float, void *, size_t)) { return callback__(self, scale_factor, png_data, png_data_size); }
	// int gocef_image_add_jpeg(cef_image_t * self, float scale_factor, void * jpeg_data, size_t jpeg_data_size, int (CEF_CALLBACK *callback__)(cef_image_t *, float, void *, size_t)) { return callback__(self, scale_factor, jpeg_data, jpeg_data_size); }
	// size_t gocef_image_get_width(cef_image_t * self, size_t (CEF_CALLBACK *callback__)(cef_image_t *)) { return callback__(self); }
	// size_t gocef_image_get_height(cef_image_t * self, size_t (CEF_CALLBACK *callback__)(cef_image_t *)) { return callback__(self); }
	// int gocef_image_has_representation(cef_image_t * self, float scale_factor, int (CEF_CALLBACK *callback__)(cef_image_t *, float)) { return callback__(self, scale_factor); }
	// int gocef_image_remove_representation(cef_image_t * self, float scale_factor, int (CEF_CALLBACK *callback__)(cef_image_t *, float)) { return callback__(self, scale_factor); }
	// int gocef_image_get_representation_info(cef_image_t * self, float scale_factor, float * actual_scale_factor, int * pixel_width, int * pixel_height, int (CEF_CALLBACK *callback__)(cef_image_t *, float, float *, int *, int *)) { return callback__(self, scale_factor, actual_scale_factor, pixel_width, pixel_height); }
	// cef_binary_value_t * gocef_image_get_as_bitmap(cef_image_t * self, float scale_factor, cef_color_type_t color_type, cef_alpha_type_t alpha_type, int * pixel_width, int * pixel_height, cef_binary_value_t * (CEF_CALLBACK *callback__)(cef_image_t *, float, cef_color_type_t, cef_alpha_type_t, int *, int *)) { return callback__(self, scale_factor, color_type, alpha_type, pixel_width, pixel_height); }
	// cef_binary_value_t * gocef_image_get_as_png(cef_image_t * self, float scale_factor, int with_transparency, int * pixel_width, int * pixel_height, cef_binary_value_t * (CEF_CALLBACK *callback__)(cef_image_t *, float, int, int *, int *)) { return callback__(self, scale_factor, with_transparency, pixel_width, pixel_height); }
	// cef_binary_value_t * gocef_image_get_as_jpeg(cef_image_t * self, float scale_factor, int quality, int * pixel_width, int * pixel_height, cef_binary_value_t * (CEF_CALLBACK *callback__)(cef_image_t *, float, int, int *, int *)) { return callback__(self, scale_factor, quality, pixel_width, pixel_height); }
	"C"
)

// Image (cef_image_t from include/capi/cef_image_capi.h)
// Container for a single image represented at different scale factors. All
// image representations should be the same size in density independent pixel
// (DIP) units. For example, if the image at scale factor 1.0 is 100x100 pixels
// then the image at scale factor 2.0 should be 200x200 pixels -- both images
// will display with a DIP size of 100x100 units. The functions of this
// structure must be called on the browser process UI thread.
type Image C.cef_image_t

func (d *Image) toNative() *C.cef_image_t {
	return (*C.cef_image_t)(d)
}

// Base (base)
// Base structure.
func (d *Image) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// IsEmpty (is_empty)
// Returns true (1) if this Image is NULL.
func (d *Image) IsEmpty() int32 {
	return int32(C.gocef_image_is_empty(d.toNative(), d.is_empty))
}

// IsSame (is_same)
// Returns true (1) if this Image and |that| Image share the same underlying
// storage. Will also return true (1) if both images are NULL.
func (d *Image) IsSame(that *Image) int32 {
	return int32(C.gocef_image_is_same(d.toNative(), that.toNative(), d.is_same))
}

// AddBitmap (add_bitmap)
// Add a bitmap image representation for |scale_factor|. Only 32-bit RGBA/BGRA
// formats are supported. |pixel_width| and |pixel_height| are the bitmap
// representation size in pixel coordinates. |pixel_data| is the array of
// pixel data and should be |pixel_width| x |pixel_height| x 4 bytes in size.
// |color_type| and |alpha_type| values specify the pixel format.
func (d *Image) AddBitmap(scale_factor float32, pixel_width, pixel_height int32, color_type ColorType, alpha_type AlphaType, pixel_data unsafe.Pointer, pixel_data_size uint64) int32 {
	return int32(C.gocef_image_add_bitmap(d.toNative(), C.float(scale_factor), C.int(pixel_width), C.int(pixel_height), C.cef_color_type_t(color_type), C.cef_alpha_type_t(alpha_type), pixel_data, C.size_t(pixel_data_size), d.add_bitmap))
}

// AddPng (add_png)
// Add a PNG image representation for |scale_factor|. |png_data| is the image
// data of size |png_data_size|. Any alpha transparency in the PNG data will
// be maintained.
func (d *Image) AddPng(scale_factor float32, png_data unsafe.Pointer, png_data_size uint64) int32 {
	return int32(C.gocef_image_add_png(d.toNative(), C.float(scale_factor), png_data, C.size_t(png_data_size), d.add_png))
}

// AddJpeg (add_jpeg)
// Create a JPEG image representation for |scale_factor|. |jpeg_data| is the
// image data of size |jpeg_data_size|. The JPEG format does not support
// transparency so the alpha byte will be set to 0xFF for all pixels.
func (d *Image) AddJpeg(scale_factor float32, jpeg_data unsafe.Pointer, jpeg_data_size uint64) int32 {
	return int32(C.gocef_image_add_jpeg(d.toNative(), C.float(scale_factor), jpeg_data, C.size_t(jpeg_data_size), d.add_jpeg))
}

// GetWidth (get_width)
// Returns the image width in density independent pixel (DIP) units.
func (d *Image) GetWidth() uint64 {
	return uint64(C.gocef_image_get_width(d.toNative(), d.get_width))
}

// GetHeight (get_height)
// Returns the image height in density independent pixel (DIP) units.
func (d *Image) GetHeight() uint64 {
	return uint64(C.gocef_image_get_height(d.toNative(), d.get_height))
}

// HasRepresentation (has_representation)
// Returns true (1) if this image contains a representation for
// |scale_factor|.
func (d *Image) HasRepresentation(scale_factor float32) int32 {
	return int32(C.gocef_image_has_representation(d.toNative(), C.float(scale_factor), d.has_representation))
}

// RemoveRepresentation (remove_representation)
// Removes the representation for |scale_factor|. Returns true (1) on success.
func (d *Image) RemoveRepresentation(scale_factor float32) int32 {
	return int32(C.gocef_image_remove_representation(d.toNative(), C.float(scale_factor), d.remove_representation))
}

// GetRepresentationInfo (get_representation_info)
// Returns information for the representation that most closely matches
// |scale_factor|. |actual_scale_factor| is the actual scale factor for the
// representation. |pixel_width| and |pixel_height| are the representation
// size in pixel coordinates. Returns true (1) on success.
func (d *Image) GetRepresentationInfo(scale_factor float32, actual_scale_factor *float32, pixel_width, pixel_height *int32) int32 {
	return int32(C.gocef_image_get_representation_info(d.toNative(), C.float(scale_factor), (*C.float)(actual_scale_factor), (*C.int)(pixel_width), (*C.int)(pixel_height), d.get_representation_info))
}

// GetAsBitmap (get_as_bitmap)
// Returns the bitmap representation that most closely matches |scale_factor|.
// Only 32-bit RGBA/BGRA formats are supported. |color_type| and |alpha_type|
// values specify the desired output pixel format. |pixel_width| and
// |pixel_height| are the output representation size in pixel coordinates.
// Returns a cef_binary_value_t containing the pixel data on success or NULL
// on failure.
func (d *Image) GetAsBitmap(scale_factor float32, color_type ColorType, alpha_type AlphaType, pixel_width, pixel_height *int32) *BinaryValue {
	return (*BinaryValue)(C.gocef_image_get_as_bitmap(d.toNative(), C.float(scale_factor), C.cef_color_type_t(color_type), C.cef_alpha_type_t(alpha_type), (*C.int)(pixel_width), (*C.int)(pixel_height), d.get_as_bitmap))
}

// GetAsPng (get_as_png)
// Returns the PNG representation that most closely matches |scale_factor|. If
// |with_transparency| is true (1) any alpha transparency in the image will be
// represented in the resulting PNG data. |pixel_width| and |pixel_height| are
// the output representation size in pixel coordinates. Returns a
// cef_binary_value_t containing the PNG image data on success or NULL on
// failure.
func (d *Image) GetAsPng(scale_factor float32, with_transparency int32, pixel_width, pixel_height *int32) *BinaryValue {
	return (*BinaryValue)(C.gocef_image_get_as_png(d.toNative(), C.float(scale_factor), C.int(with_transparency), (*C.int)(pixel_width), (*C.int)(pixel_height), d.get_as_png))
}

// GetAsJpeg (get_as_jpeg)
// Returns the JPEG representation that most closely matches |scale_factor|.
// |quality| determines the compression level with 0 == lowest and 100 ==
// highest. The JPEG format does not support alpha transparency and the alpha
// channel, if any, will be discarded. |pixel_width| and |pixel_height| are
// the output representation size in pixel coordinates. Returns a
// cef_binary_value_t containing the JPEG image data on success or NULL on
// failure.
func (d *Image) GetAsJpeg(scale_factor float32, quality int32, pixel_width, pixel_height *int32) *BinaryValue {
	return (*BinaryValue)(C.gocef_image_get_as_jpeg(d.toNative(), C.float(scale_factor), C.int(quality), (*C.int)(pixel_width), (*C.int)(pixel_height), d.get_as_jpeg))
}
