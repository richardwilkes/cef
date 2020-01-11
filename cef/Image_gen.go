// Code created from "class.go.tmpl" - don't edit by hand

package cef

import "unsafe"

import (
	// #include "capi_gen.h"
	// int gocef_image_is_empty(cef_image_t * self, int (CEF_CALLBACK *callback__)(cef_image_t *)) { return callback__(self); }
	// int gocef_image_is_same(cef_image_t * self, cef_image_t * that, int (CEF_CALLBACK *callback__)(cef_image_t *, cef_image_t *)) { return callback__(self, that); }
	// int gocef_image_add_bitmap(cef_image_t * self, float scaleFactor, int pixelWidth, int pixelHeight, cef_color_type_t colorType, cef_alpha_type_t alphaType, void * pixelData, size_t pixelDataSize, int (CEF_CALLBACK *callback__)(cef_image_t *, float, int, int, cef_color_type_t, cef_alpha_type_t, void *, size_t)) { return callback__(self, scaleFactor, pixelWidth, pixelHeight, colorType, alphaType, pixelData, pixelDataSize); }
	// int gocef_image_add_png(cef_image_t * self, float scaleFactor, void * pngData, size_t pngDataSize, int (CEF_CALLBACK *callback__)(cef_image_t *, float, void *, size_t)) { return callback__(self, scaleFactor, pngData, pngDataSize); }
	// int gocef_image_add_jpeg(cef_image_t * self, float scaleFactor, void * jpegData, size_t jpegDataSize, int (CEF_CALLBACK *callback__)(cef_image_t *, float, void *, size_t)) { return callback__(self, scaleFactor, jpegData, jpegDataSize); }
	// size_t gocef_image_get_width(cef_image_t * self, size_t (CEF_CALLBACK *callback__)(cef_image_t *)) { return callback__(self); }
	// size_t gocef_image_get_height(cef_image_t * self, size_t (CEF_CALLBACK *callback__)(cef_image_t *)) { return callback__(self); }
	// int gocef_image_has_representation(cef_image_t * self, float scaleFactor, int (CEF_CALLBACK *callback__)(cef_image_t *, float)) { return callback__(self, scaleFactor); }
	// int gocef_image_remove_representation(cef_image_t * self, float scaleFactor, int (CEF_CALLBACK *callback__)(cef_image_t *, float)) { return callback__(self, scaleFactor); }
	// int gocef_image_get_representation_info(cef_image_t * self, float scaleFactor, float * actualScaleFactor, int * pixelWidth, int * pixelHeight, int (CEF_CALLBACK *callback__)(cef_image_t *, float, float *, int *, int *)) { return callback__(self, scaleFactor, actualScaleFactor, pixelWidth, pixelHeight); }
	// cef_binary_value_t * gocef_image_get_as_bitmap(cef_image_t * self, float scaleFactor, cef_color_type_t colorType, cef_alpha_type_t alphaType, int * pixelWidth, int * pixelHeight, cef_binary_value_t * (CEF_CALLBACK *callback__)(cef_image_t *, float, cef_color_type_t, cef_alpha_type_t, int *, int *)) { return callback__(self, scaleFactor, colorType, alphaType, pixelWidth, pixelHeight); }
	// cef_binary_value_t * gocef_image_get_as_png(cef_image_t * self, float scaleFactor, int withTransparency, int * pixelWidth, int * pixelHeight, cef_binary_value_t * (CEF_CALLBACK *callback__)(cef_image_t *, float, int, int *, int *)) { return callback__(self, scaleFactor, withTransparency, pixelWidth, pixelHeight); }
	// cef_binary_value_t * gocef_image_get_as_jpeg(cef_image_t * self, float scaleFactor, int quality, int * pixelWidth, int * pixelHeight, cef_binary_value_t * (CEF_CALLBACK *callback__)(cef_image_t *, float, int, int *, int *)) { return callback__(self, scaleFactor, quality, pixelWidth, pixelHeight); }
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
func (d *Image) AddBitmap(scaleFactor float32, pixelWidth, pixelHeight int32, colorType ColorType, alphaType AlphaType, pixelData unsafe.Pointer, pixelDataSize uint64) int32 {
	return int32(C.gocef_image_add_bitmap(d.toNative(), C.float(scaleFactor), C.int(pixelWidth), C.int(pixelHeight), C.cef_color_type_t(colorType), C.cef_alpha_type_t(alphaType), pixelData, C.size_t(pixelDataSize), d.add_bitmap))
}

// AddPng (add_png)
// Add a PNG image representation for |scale_factor|. |png_data| is the image
// data of size |png_data_size|. Any alpha transparency in the PNG data will
// be maintained.
func (d *Image) AddPng(scaleFactor float32, pngData unsafe.Pointer, pngDataSize uint64) int32 {
	return int32(C.gocef_image_add_png(d.toNative(), C.float(scaleFactor), pngData, C.size_t(pngDataSize), d.add_png))
}

// AddJpeg (add_jpeg)
// Create a JPEG image representation for |scale_factor|. |jpeg_data| is the
// image data of size |jpeg_data_size|. The JPEG format does not support
// transparency so the alpha byte will be set to 0xFF for all pixels.
func (d *Image) AddJpeg(scaleFactor float32, jpegData unsafe.Pointer, jpegDataSize uint64) int32 {
	return int32(C.gocef_image_add_jpeg(d.toNative(), C.float(scaleFactor), jpegData, C.size_t(jpegDataSize), d.add_jpeg))
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
func (d *Image) HasRepresentation(scaleFactor float32) int32 {
	return int32(C.gocef_image_has_representation(d.toNative(), C.float(scaleFactor), d.has_representation))
}

// RemoveRepresentation (remove_representation)
// Removes the representation for |scale_factor|. Returns true (1) on success.
func (d *Image) RemoveRepresentation(scaleFactor float32) int32 {
	return int32(C.gocef_image_remove_representation(d.toNative(), C.float(scaleFactor), d.remove_representation))
}

// GetRepresentationInfo (get_representation_info)
// Returns information for the representation that most closely matches
// |scale_factor|. |actual_scale_factor| is the actual scale factor for the
// representation. |pixel_width| and |pixel_height| are the representation
// size in pixel coordinates. Returns true (1) on success.
func (d *Image) GetRepresentationInfo(scaleFactor float32, actualScaleFactor *float32, pixelWidth, pixelHeight *int32) int32 {
	return int32(C.gocef_image_get_representation_info(d.toNative(), C.float(scaleFactor), (*C.float)(actualScaleFactor), (*C.int)(pixelWidth), (*C.int)(pixelHeight), d.get_representation_info))
}

// GetAsBitmap (get_as_bitmap)
// Returns the bitmap representation that most closely matches |scale_factor|.
// Only 32-bit RGBA/BGRA formats are supported. |color_type| and |alpha_type|
// values specify the desired output pixel format. |pixel_width| and
// |pixel_height| are the output representation size in pixel coordinates.
// Returns a cef_binary_value_t containing the pixel data on success or NULL
// on failure.
func (d *Image) GetAsBitmap(scaleFactor float32, colorType ColorType, alphaType AlphaType, pixelWidth, pixelHeight *int32) *BinaryValue {
	return (*BinaryValue)(C.gocef_image_get_as_bitmap(d.toNative(), C.float(scaleFactor), C.cef_color_type_t(colorType), C.cef_alpha_type_t(alphaType), (*C.int)(pixelWidth), (*C.int)(pixelHeight), d.get_as_bitmap))
}

// GetAsPng (get_as_png)
// Returns the PNG representation that most closely matches |scale_factor|. If
// |with_transparency| is true (1) any alpha transparency in the image will be
// represented in the resulting PNG data. |pixel_width| and |pixel_height| are
// the output representation size in pixel coordinates. Returns a
// cef_binary_value_t containing the PNG image data on success or NULL on
// failure.
func (d *Image) GetAsPng(scaleFactor float32, withTransparency int32, pixelWidth, pixelHeight *int32) *BinaryValue {
	return (*BinaryValue)(C.gocef_image_get_as_png(d.toNative(), C.float(scaleFactor), C.int(withTransparency), (*C.int)(pixelWidth), (*C.int)(pixelHeight), d.get_as_png))
}

// GetAsJpeg (get_as_jpeg)
// Returns the JPEG representation that most closely matches |scale_factor|.
// |quality| determines the compression level with 0 == lowest and 100 ==
// highest. The JPEG format does not support alpha transparency and the alpha
// channel, if any, will be discarded. |pixel_width| and |pixel_height| are
// the output representation size in pixel coordinates. Returns a
// cef_binary_value_t containing the JPEG image data on success or NULL on
// failure.
func (d *Image) GetAsJpeg(scaleFactor float32, quality int32, pixelWidth, pixelHeight *int32) *BinaryValue {
	return (*BinaryValue)(C.gocef_image_get_as_jpeg(d.toNative(), C.float(scaleFactor), C.int(quality), (*C.int)(pixelWidth), (*C.int)(pixelHeight), d.get_as_jpeg))
}
