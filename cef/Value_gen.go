// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// int gocef_value_is_valid(cef_value_t * self, int (CEF_CALLBACK *callback__)(cef_value_t *)) { return callback__(self); }
	// int gocef_value_is_owned(cef_value_t * self, int (CEF_CALLBACK *callback__)(cef_value_t *)) { return callback__(self); }
	// int gocef_value_is_read_only(cef_value_t * self, int (CEF_CALLBACK *callback__)(cef_value_t *)) { return callback__(self); }
	// int gocef_value_is_same(cef_value_t * self, cef_value_t * that, int (CEF_CALLBACK *callback__)(cef_value_t *, cef_value_t *)) { return callback__(self, that); }
	// int gocef_value_is_equal(cef_value_t * self, cef_value_t * that, int (CEF_CALLBACK *callback__)(cef_value_t *, cef_value_t *)) { return callback__(self, that); }
	// cef_value_t * gocef_value_copy(cef_value_t * self, cef_value_t * (CEF_CALLBACK *callback__)(cef_value_t *)) { return callback__(self); }
	// cef_value_type_t gocef_value_get_type(cef_value_t * self, cef_value_type_t (CEF_CALLBACK *callback__)(cef_value_t *)) { return callback__(self); }
	// int gocef_value_get_bool(cef_value_t * self, int (CEF_CALLBACK *callback__)(cef_value_t *)) { return callback__(self); }
	// int gocef_value_get_int(cef_value_t * self, int (CEF_CALLBACK *callback__)(cef_value_t *)) { return callback__(self); }
	// double gocef_value_get_double(cef_value_t * self, double (CEF_CALLBACK *callback__)(cef_value_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_value_get_string(cef_value_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_value_t *)) { return callback__(self); }
	// cef_binary_value_t * gocef_value_get_binary(cef_value_t * self, cef_binary_value_t * (CEF_CALLBACK *callback__)(cef_value_t *)) { return callback__(self); }
	// cef_dictionary_value_t * gocef_value_get_dictionary(cef_value_t * self, cef_dictionary_value_t * (CEF_CALLBACK *callback__)(cef_value_t *)) { return callback__(self); }
	// cef_list_value_t * gocef_value_get_list(cef_value_t * self, cef_list_value_t * (CEF_CALLBACK *callback__)(cef_value_t *)) { return callback__(self); }
	// int gocef_value_set_null(cef_value_t * self, int (CEF_CALLBACK *callback__)(cef_value_t *)) { return callback__(self); }
	// int gocef_value_set_bool(cef_value_t * self, int value, int (CEF_CALLBACK *callback__)(cef_value_t *, int)) { return callback__(self, value); }
	// int gocef_value_set_int(cef_value_t * self, int value, int (CEF_CALLBACK *callback__)(cef_value_t *, int)) { return callback__(self, value); }
	// int gocef_value_set_double(cef_value_t * self, double value, int (CEF_CALLBACK *callback__)(cef_value_t *, double)) { return callback__(self, value); }
	// int gocef_value_set_string(cef_value_t * self, cef_string_t * value, int (CEF_CALLBACK *callback__)(cef_value_t *, cef_string_t *)) { return callback__(self, value); }
	// int gocef_value_set_binary(cef_value_t * self, cef_binary_value_t * value, int (CEF_CALLBACK *callback__)(cef_value_t *, cef_binary_value_t *)) { return callback__(self, value); }
	// int gocef_value_set_dictionary(cef_value_t * self, cef_dictionary_value_t * value, int (CEF_CALLBACK *callback__)(cef_value_t *, cef_dictionary_value_t *)) { return callback__(self, value); }
	// int gocef_value_set_list(cef_value_t * self, cef_list_value_t * value, int (CEF_CALLBACK *callback__)(cef_value_t *, cef_list_value_t *)) { return callback__(self, value); }
	"C"
)

// Value (cef_value_t from include/capi/cef_values_capi.h)
// Structure that wraps other data value types. Complex types (binary,
// dictionary and list) will be referenced but not owned by this object. Can be
// used on any process and thread.
type Value C.cef_value_t

func (d *Value) toNative() *C.cef_value_t {
	return (*C.cef_value_t)(d)
}

// Base (base)
// Base structure.
func (d *Value) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// IsValid (is_valid)
// Returns true (1) if the underlying data is valid. This will always be true
// (1) for simple types. For complex types (binary, dictionary and list) the
// underlying data may become invalid if owned by another object (e.g. list or
// dictionary) and that other object is then modified or destroyed. This value
// object can be re-used by calling Set*() even if the underlying data is
// invalid.
func (d *Value) IsValid() int32 {
	return int32(C.gocef_value_is_valid(d.toNative(), d.is_valid))
}

// IsOwned (is_owned)
// Returns true (1) if the underlying data is owned by another object.
func (d *Value) IsOwned() int32 {
	return int32(C.gocef_value_is_owned(d.toNative(), d.is_owned))
}

// IsReadOnly (is_read_only)
// Returns true (1) if the underlying data is read-only. Some APIs may expose
// read-only objects.
func (d *Value) IsReadOnly() int32 {
	return int32(C.gocef_value_is_read_only(d.toNative(), d.is_read_only))
}

// IsSame (is_same)
// Returns true (1) if this object and |that| object have the same underlying
// data. If true (1) modifications to this object will also affect |that|
// object and vice-versa.
func (d *Value) IsSame(that *Value) int32 {
	return int32(C.gocef_value_is_same(d.toNative(), that.toNative(), d.is_same))
}

// IsEqual (is_equal)
// Returns true (1) if this object and |that| object have an equivalent
// underlying value but are not necessarily the same object.
func (d *Value) IsEqual(that *Value) int32 {
	return int32(C.gocef_value_is_equal(d.toNative(), that.toNative(), d.is_equal))
}

// Copy (copy)
// Returns a copy of this object. The underlying data will also be copied.
func (d *Value) Copy() *Value {
	return (*Value)(C.gocef_value_copy(d.toNative(), d.copy))
}

// GetType (get_type)
// Returns the underlying value type.
func (d *Value) GetType() ValueType {
	return ValueType(C.gocef_value_get_type(d.toNative(), d.get_type))
}

// GetBool (get_bool)
// Returns the underlying value as type bool.
func (d *Value) GetBool() int32 {
	return int32(C.gocef_value_get_bool(d.toNative(), d.get_bool))
}

// GetInt (get_int)
// Returns the underlying value as type int.
func (d *Value) GetInt() int32 {
	return int32(C.gocef_value_get_int(d.toNative(), d.get_int))
}

// GetDouble (get_double)
// Returns the underlying value as type double.
func (d *Value) GetDouble() float64 {
	return float64(C.gocef_value_get_double(d.toNative(), d.get_double))
}

// GetString (get_string)
// Returns the underlying value as type string.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Value) GetString() string {
	return cefuserfreestrToString(C.gocef_value_get_string(d.toNative(), d.get_string))
}

// GetBinary (get_binary)
// Returns the underlying value as type binary. The returned reference may
// become invalid if the value is owned by another object or if ownership is
// transferred to another object in the future. To maintain a reference to the
// value after assigning ownership to a dictionary or list pass this object to
// the set_value() function instead of passing the returned reference to
// set_binary().
func (d *Value) GetBinary() *BinaryValue {
	return (*BinaryValue)(C.gocef_value_get_binary(d.toNative(), d.get_binary))
}

// GetDictionary (get_dictionary)
// Returns the underlying value as type dictionary. The returned reference may
// become invalid if the value is owned by another object or if ownership is
// transferred to another object in the future. To maintain a reference to the
// value after assigning ownership to a dictionary or list pass this object to
// the set_value() function instead of passing the returned reference to
// set_dictionary().
func (d *Value) GetDictionary() *DictionaryValue {
	return (*DictionaryValue)(C.gocef_value_get_dictionary(d.toNative(), d.get_dictionary))
}

// GetList (get_list)
// Returns the underlying value as type list. The returned reference may
// become invalid if the value is owned by another object or if ownership is
// transferred to another object in the future. To maintain a reference to the
// value after assigning ownership to a dictionary or list pass this object to
// the set_value() function instead of passing the returned reference to
// set_list().
func (d *Value) GetList() *ListValue {
	return (*ListValue)(C.gocef_value_get_list(d.toNative(), d.get_list))
}

// SetNull (set_null)
// Sets the underlying value as type null. Returns true (1) if the value was
// set successfully.
func (d *Value) SetNull() int32 {
	return int32(C.gocef_value_set_null(d.toNative(), d.set_null))
}

// SetBool (set_bool)
// Sets the underlying value as type bool. Returns true (1) if the value was
// set successfully.
func (d *Value) SetBool(value int32) int32 {
	return int32(C.gocef_value_set_bool(d.toNative(), C.int(value), d.set_bool))
}

// SetInt (set_int)
// Sets the underlying value as type int. Returns true (1) if the value was
// set successfully.
func (d *Value) SetInt(value int32) int32 {
	return int32(C.gocef_value_set_int(d.toNative(), C.int(value), d.set_int))
}

// SetDouble (set_double)
// Sets the underlying value as type double. Returns true (1) if the value was
// set successfully.
func (d *Value) SetDouble(value float64) int32 {
	return int32(C.gocef_value_set_double(d.toNative(), C.double(value), d.set_double))
}

// SetString (set_string)
// Sets the underlying value as type string. Returns true (1) if the value was
// set successfully.
func (d *Value) SetString(value string) int32 {
	value_ := C.cef_string_userfree_alloc()
	setCEFStr(value, value_)
	defer func() {
		C.cef_string_userfree_free(value_)
	}()
	return int32(C.gocef_value_set_string(d.toNative(), (*C.cef_string_t)(value_), d.set_string))
}

// SetBinary (set_binary)
// Sets the underlying value as type binary. Returns true (1) if the value was
// set successfully. This object keeps a reference to |value| and ownership of
// the underlying data remains unchanged.
func (d *Value) SetBinary(value *BinaryValue) int32 {
	return int32(C.gocef_value_set_binary(d.toNative(), value.toNative(), d.set_binary))
}

// SetDictionary (set_dictionary)
// Sets the underlying value as type dict. Returns true (1) if the value was
// set successfully. This object keeps a reference to |value| and ownership of
// the underlying data remains unchanged.
func (d *Value) SetDictionary(value *DictionaryValue) int32 {
	return int32(C.gocef_value_set_dictionary(d.toNative(), value.toNative(), d.set_dictionary))
}

// SetList (set_list)
// Sets the underlying value as type list. Returns true (1) if the value was
// set successfully. This object keeps a reference to |value| and ownership of
// the underlying data remains unchanged.
func (d *Value) SetList(value *ListValue) int32 {
	return int32(C.gocef_value_set_list(d.toNative(), value.toNative(), d.set_list))
}
