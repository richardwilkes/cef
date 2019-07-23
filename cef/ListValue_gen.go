// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// int gocef_list_value_is_valid(cef_list_value_t * self, int (CEF_CALLBACK *callback__)(cef_list_value_t *)) { return callback__(self); }
	// int gocef_list_value_is_owned(cef_list_value_t * self, int (CEF_CALLBACK *callback__)(cef_list_value_t *)) { return callback__(self); }
	// int gocef_list_value_is_read_only(cef_list_value_t * self, int (CEF_CALLBACK *callback__)(cef_list_value_t *)) { return callback__(self); }
	// int gocef_list_value_is_same(cef_list_value_t * self, cef_list_value_t * that, int (CEF_CALLBACK *callback__)(cef_list_value_t *, cef_list_value_t *)) { return callback__(self, that); }
	// int gocef_list_value_is_equal(cef_list_value_t * self, cef_list_value_t * that, int (CEF_CALLBACK *callback__)(cef_list_value_t *, cef_list_value_t *)) { return callback__(self, that); }
	// cef_list_value_t * gocef_list_value_copy(cef_list_value_t * self, cef_list_value_t * (CEF_CALLBACK *callback__)(cef_list_value_t *)) { return callback__(self); }
	// int gocef_list_value_set_size(cef_list_value_t * self, size_t size, int (CEF_CALLBACK *callback__)(cef_list_value_t *, size_t)) { return callback__(self, size); }
	// size_t gocef_list_value_get_size(cef_list_value_t * self, size_t (CEF_CALLBACK *callback__)(cef_list_value_t *)) { return callback__(self); }
	// int gocef_list_value_clear(cef_list_value_t * self, int (CEF_CALLBACK *callback__)(cef_list_value_t *)) { return callback__(self); }
	// int gocef_list_value_remove(cef_list_value_t * self, size_t index, int (CEF_CALLBACK *callback__)(cef_list_value_t *, size_t)) { return callback__(self, index); }
	// cef_value_type_t gocef_list_value_get_type(cef_list_value_t * self, size_t index, cef_value_type_t (CEF_CALLBACK *callback__)(cef_list_value_t *, size_t)) { return callback__(self, index); }
	// cef_value_t * gocef_list_value_get_value(cef_list_value_t * self, size_t index, cef_value_t * (CEF_CALLBACK *callback__)(cef_list_value_t *, size_t)) { return callback__(self, index); }
	// int gocef_list_value_get_bool(cef_list_value_t * self, size_t index, int (CEF_CALLBACK *callback__)(cef_list_value_t *, size_t)) { return callback__(self, index); }
	// int gocef_list_value_get_int(cef_list_value_t * self, size_t index, int (CEF_CALLBACK *callback__)(cef_list_value_t *, size_t)) { return callback__(self, index); }
	// double gocef_list_value_get_double(cef_list_value_t * self, size_t index, double (CEF_CALLBACK *callback__)(cef_list_value_t *, size_t)) { return callback__(self, index); }
	// cef_string_userfree_t gocef_list_value_get_string(cef_list_value_t * self, size_t index, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_list_value_t *, size_t)) { return callback__(self, index); }
	// cef_binary_value_t * gocef_list_value_get_binary(cef_list_value_t * self, size_t index, cef_binary_value_t * (CEF_CALLBACK *callback__)(cef_list_value_t *, size_t)) { return callback__(self, index); }
	// cef_dictionary_value_t * gocef_list_value_get_dictionary(cef_list_value_t * self, size_t index, cef_dictionary_value_t * (CEF_CALLBACK *callback__)(cef_list_value_t *, size_t)) { return callback__(self, index); }
	// cef_list_value_t * gocef_list_value_get_list(cef_list_value_t * self, size_t index, cef_list_value_t * (CEF_CALLBACK *callback__)(cef_list_value_t *, size_t)) { return callback__(self, index); }
	// int gocef_list_value_set_value(cef_list_value_t * self, size_t index, cef_value_t * value, int (CEF_CALLBACK *callback__)(cef_list_value_t *, size_t, cef_value_t *)) { return callback__(self, index, value); }
	// int gocef_list_value_set_null(cef_list_value_t * self, size_t index, int (CEF_CALLBACK *callback__)(cef_list_value_t *, size_t)) { return callback__(self, index); }
	// int gocef_list_value_set_bool(cef_list_value_t * self, size_t index, int value, int (CEF_CALLBACK *callback__)(cef_list_value_t *, size_t, int)) { return callback__(self, index, value); }
	// int gocef_list_value_set_int(cef_list_value_t * self, size_t index, int value, int (CEF_CALLBACK *callback__)(cef_list_value_t *, size_t, int)) { return callback__(self, index, value); }
	// int gocef_list_value_set_double(cef_list_value_t * self, size_t index, double value, int (CEF_CALLBACK *callback__)(cef_list_value_t *, size_t, double)) { return callback__(self, index, value); }
	// int gocef_list_value_set_string(cef_list_value_t * self, size_t index, cef_string_t * value, int (CEF_CALLBACK *callback__)(cef_list_value_t *, size_t, cef_string_t *)) { return callback__(self, index, value); }
	// int gocef_list_value_set_binary(cef_list_value_t * self, size_t index, cef_binary_value_t * value, int (CEF_CALLBACK *callback__)(cef_list_value_t *, size_t, cef_binary_value_t *)) { return callback__(self, index, value); }
	// int gocef_list_value_set_dictionary(cef_list_value_t * self, size_t index, cef_dictionary_value_t * value, int (CEF_CALLBACK *callback__)(cef_list_value_t *, size_t, cef_dictionary_value_t *)) { return callback__(self, index, value); }
	// int gocef_list_value_set_list(cef_list_value_t * self, size_t index, cef_list_value_t * value, int (CEF_CALLBACK *callback__)(cef_list_value_t *, size_t, cef_list_value_t *)) { return callback__(self, index, value); }
	"C"
)

// ListValue (cef_list_value_t from include/capi/cef_values_capi.h)
// Structure representing a list value. Can be used on any process and thread.
type ListValue C.cef_list_value_t

func (d *ListValue) toNative() *C.cef_list_value_t {
	return (*C.cef_list_value_t)(d)
}

// Base (base)
// Base structure.
func (d *ListValue) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// IsValid (is_valid)
// Returns true (1) if this object is valid. This object may become invalid if
// the underlying data is owned by another object (e.g. list or dictionary)
// and that other object is then modified or destroyed. Do not call any other
// functions if this function returns false (0).
func (d *ListValue) IsValid() int32 {
	return int32(C.gocef_list_value_is_valid(d.toNative(), d.is_valid))
}

// IsOwned (is_owned)
// Returns true (1) if this object is currently owned by another object.
func (d *ListValue) IsOwned() int32 {
	return int32(C.gocef_list_value_is_owned(d.toNative(), d.is_owned))
}

// IsReadOnly (is_read_only)
// Returns true (1) if the values of this object are read-only. Some APIs may
// expose read-only objects.
func (d *ListValue) IsReadOnly() int32 {
	return int32(C.gocef_list_value_is_read_only(d.toNative(), d.is_read_only))
}

// IsSame (is_same)
// Returns true (1) if this object and |that| object have the same underlying
// data. If true (1) modifications to this object will also affect |that|
// object and vice-versa.
func (d *ListValue) IsSame(that *ListValue) int32 {
	return int32(C.gocef_list_value_is_same(d.toNative(), that.toNative(), d.is_same))
}

// IsEqual (is_equal)
// Returns true (1) if this object and |that| object have an equivalent
// underlying value but are not necessarily the same object.
func (d *ListValue) IsEqual(that *ListValue) int32 {
	return int32(C.gocef_list_value_is_equal(d.toNative(), that.toNative(), d.is_equal))
}

// Copy (copy)
// Returns a writable copy of this object.
func (d *ListValue) Copy() *ListValue {
	return (*ListValue)(C.gocef_list_value_copy(d.toNative(), d.copy))
}

// SetSize (set_size)
// Sets the number of values. If the number of values is expanded all new
// value slots will default to type null. Returns true (1) on success.
func (d *ListValue) SetSize(size uint64) int32 {
	return int32(C.gocef_list_value_set_size(d.toNative(), C.size_t(size), d.set_size))
}

// GetSize (get_size)
// Returns the number of values.
func (d *ListValue) GetSize() uint64 {
	return uint64(C.gocef_list_value_get_size(d.toNative(), d.get_size))
}

// Clear (clear)
// Removes all values. Returns true (1) on success.
func (d *ListValue) Clear() int32 {
	return int32(C.gocef_list_value_clear(d.toNative(), d.clear))
}

// Remove (remove)
// Removes the value at the specified index.
func (d *ListValue) Remove(index uint64) int32 {
	return int32(C.gocef_list_value_remove(d.toNative(), C.size_t(index), d.remove))
}

// GetType (get_type)
// Returns the value type at the specified index.
func (d *ListValue) GetType(index uint64) ValueType {
	return ValueType(C.gocef_list_value_get_type(d.toNative(), C.size_t(index), d.get_type))
}

// GetValue (get_value)
// Returns the value at the specified index. For simple types the returned
// value will copy existing data and modifications to the value will not
// modify this object. For complex types (binary, dictionary and list) the
// returned value will reference existing data and modifications to the value
// will modify this object.
func (d *ListValue) GetValue(index uint64) *Value {
	return (*Value)(C.gocef_list_value_get_value(d.toNative(), C.size_t(index), d.get_value))
}

// GetBool (get_bool)
// Returns the value at the specified index as type bool.
func (d *ListValue) GetBool(index uint64) int32 {
	return int32(C.gocef_list_value_get_bool(d.toNative(), C.size_t(index), d.get_bool))
}

// GetInt (get_int)
// Returns the value at the specified index as type int.
func (d *ListValue) GetInt(index uint64) int32 {
	return int32(C.gocef_list_value_get_int(d.toNative(), C.size_t(index), d.get_int))
}

// GetDouble (get_double)
// Returns the value at the specified index as type double.
func (d *ListValue) GetDouble(index uint64) float64 {
	return float64(C.gocef_list_value_get_double(d.toNative(), C.size_t(index), d.get_double))
}

// GetString (get_string)
// Returns the value at the specified index as type string.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *ListValue) GetString(index uint64) string {
	return cefuserfreestrToString(C.gocef_list_value_get_string(d.toNative(), C.size_t(index), d.get_string))
}

// GetBinary (get_binary)
// Returns the value at the specified index as type binary. The returned value
// will reference existing data.
func (d *ListValue) GetBinary(index uint64) *BinaryValue {
	return (*BinaryValue)(C.gocef_list_value_get_binary(d.toNative(), C.size_t(index), d.get_binary))
}

// GetDictionary (get_dictionary)
// Returns the value at the specified index as type dictionary. The returned
// value will reference existing data and modifications to the value will
// modify this object.
func (d *ListValue) GetDictionary(index uint64) *DictionaryValue {
	return (*DictionaryValue)(C.gocef_list_value_get_dictionary(d.toNative(), C.size_t(index), d.get_dictionary))
}

// GetList (get_list)
// Returns the value at the specified index as type list. The returned value
// will reference existing data and modifications to the value will modify
// this object.
func (d *ListValue) GetList(index uint64) *ListValue {
	return (*ListValue)(C.gocef_list_value_get_list(d.toNative(), C.size_t(index), d.get_list))
}

// SetValue (set_value)
// Sets the value at the specified index. Returns true (1) if the value was
// set successfully. If |value| represents simple data then the underlying
// data will be copied and modifications to |value| will not modify this
// object. If |value| represents complex data (binary, dictionary or list)
// then the underlying data will be referenced and modifications to |value|
// will modify this object.
func (d *ListValue) SetValue(index uint64, value *Value) int32 {
	return int32(C.gocef_list_value_set_value(d.toNative(), C.size_t(index), value.toNative(), d.set_value))
}

// SetNull (set_null)
// Sets the value at the specified index as type null. Returns true (1) if the
// value was set successfully.
func (d *ListValue) SetNull(index uint64) int32 {
	return int32(C.gocef_list_value_set_null(d.toNative(), C.size_t(index), d.set_null))
}

// SetBool (set_bool)
// Sets the value at the specified index as type bool. Returns true (1) if the
// value was set successfully.
func (d *ListValue) SetBool(index uint64, value int32) int32 {
	return int32(C.gocef_list_value_set_bool(d.toNative(), C.size_t(index), C.int(value), d.set_bool))
}

// SetInt (set_int)
// Sets the value at the specified index as type int. Returns true (1) if the
// value was set successfully.
func (d *ListValue) SetInt(index uint64, value int32) int32 {
	return int32(C.gocef_list_value_set_int(d.toNative(), C.size_t(index), C.int(value), d.set_int))
}

// SetDouble (set_double)
// Sets the value at the specified index as type double. Returns true (1) if
// the value was set successfully.
func (d *ListValue) SetDouble(index uint64, value float64) int32 {
	return int32(C.gocef_list_value_set_double(d.toNative(), C.size_t(index), C.double(value), d.set_double))
}

// SetString (set_string)
// Sets the value at the specified index as type string. Returns true (1) if
// the value was set successfully.
func (d *ListValue) SetString(index uint64, value string) int32 {
	value_ := C.cef_string_userfree_alloc()
	setCEFStr(value, value_)
	defer func() {
		C.cef_string_userfree_free(value_)
	}()
	return int32(C.gocef_list_value_set_string(d.toNative(), C.size_t(index), (*C.cef_string_t)(value_), d.set_string))
}

// SetBinary (set_binary)
// Sets the value at the specified index as type binary. Returns true (1) if
// the value was set successfully. If |value| is currently owned by another
// object then the value will be copied and the |value| reference will not
// change. Otherwise, ownership will be transferred to this object and the
// |value| reference will be invalidated.
func (d *ListValue) SetBinary(index uint64, value *BinaryValue) int32 {
	return int32(C.gocef_list_value_set_binary(d.toNative(), C.size_t(index), value.toNative(), d.set_binary))
}

// SetDictionary (set_dictionary)
// Sets the value at the specified index as type dict. Returns true (1) if the
// value was set successfully. If |value| is currently owned by another object
// then the value will be copied and the |value| reference will not change.
// Otherwise, ownership will be transferred to this object and the |value|
// reference will be invalidated.
func (d *ListValue) SetDictionary(index uint64, value *DictionaryValue) int32 {
	return int32(C.gocef_list_value_set_dictionary(d.toNative(), C.size_t(index), value.toNative(), d.set_dictionary))
}

// SetList (set_list)
// Sets the value at the specified index as type list. Returns true (1) if the
// value was set successfully. If |value| is currently owned by another object
// then the value will be copied and the |value| reference will not change.
// Otherwise, ownership will be transferred to this object and the |value|
// reference will be invalidated.
func (d *ListValue) SetList(index uint64, value *ListValue) int32 {
	return int32(C.gocef_list_value_set_list(d.toNative(), C.size_t(index), value.toNative(), d.set_list))
}
