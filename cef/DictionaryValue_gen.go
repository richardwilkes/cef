// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// int gocef_dictionary_value_is_valid(cef_dictionary_value_t * self, int (CEF_CALLBACK *callback__)(cef_dictionary_value_t *)) { return callback__(self); }
	// int gocef_dictionary_value_is_owned(cef_dictionary_value_t * self, int (CEF_CALLBACK *callback__)(cef_dictionary_value_t *)) { return callback__(self); }
	// int gocef_dictionary_value_is_read_only(cef_dictionary_value_t * self, int (CEF_CALLBACK *callback__)(cef_dictionary_value_t *)) { return callback__(self); }
	// int gocef_dictionary_value_is_same(cef_dictionary_value_t * self, cef_dictionary_value_t * that, int (CEF_CALLBACK *callback__)(cef_dictionary_value_t *, cef_dictionary_value_t *)) { return callback__(self, that); }
	// int gocef_dictionary_value_is_equal(cef_dictionary_value_t * self, cef_dictionary_value_t * that, int (CEF_CALLBACK *callback__)(cef_dictionary_value_t *, cef_dictionary_value_t *)) { return callback__(self, that); }
	// cef_dictionary_value_t * gocef_dictionary_value_copy(cef_dictionary_value_t * self, int excludeEmptyChildren, cef_dictionary_value_t * (CEF_CALLBACK *callback__)(cef_dictionary_value_t *, int)) { return callback__(self, excludeEmptyChildren); }
	// size_t gocef_dictionary_value_get_size(cef_dictionary_value_t * self, size_t (CEF_CALLBACK *callback__)(cef_dictionary_value_t *)) { return callback__(self); }
	// int gocef_dictionary_value_clear(cef_dictionary_value_t * self, int (CEF_CALLBACK *callback__)(cef_dictionary_value_t *)) { return callback__(self); }
	// int gocef_dictionary_value_has_key(cef_dictionary_value_t * self, cef_string_t * key, int (CEF_CALLBACK *callback__)(cef_dictionary_value_t *, cef_string_t *)) { return callback__(self, key); }
	// int gocef_dictionary_value_get_keys(cef_dictionary_value_t * self, cef_string_list_t keys, int (CEF_CALLBACK *callback__)(cef_dictionary_value_t *, cef_string_list_t)) { return callback__(self, keys); }
	// int gocef_dictionary_value_remove(cef_dictionary_value_t * self, cef_string_t * key, int (CEF_CALLBACK *callback__)(cef_dictionary_value_t *, cef_string_t *)) { return callback__(self, key); }
	// cef_value_type_t gocef_dictionary_value_get_type(cef_dictionary_value_t * self, cef_string_t * key, cef_value_type_t (CEF_CALLBACK *callback__)(cef_dictionary_value_t *, cef_string_t *)) { return callback__(self, key); }
	// cef_value_t * gocef_dictionary_value_get_value(cef_dictionary_value_t * self, cef_string_t * key, cef_value_t * (CEF_CALLBACK *callback__)(cef_dictionary_value_t *, cef_string_t *)) { return callback__(self, key); }
	// int gocef_dictionary_value_get_bool(cef_dictionary_value_t * self, cef_string_t * key, int (CEF_CALLBACK *callback__)(cef_dictionary_value_t *, cef_string_t *)) { return callback__(self, key); }
	// int gocef_dictionary_value_get_int(cef_dictionary_value_t * self, cef_string_t * key, int (CEF_CALLBACK *callback__)(cef_dictionary_value_t *, cef_string_t *)) { return callback__(self, key); }
	// double gocef_dictionary_value_get_double(cef_dictionary_value_t * self, cef_string_t * key, double (CEF_CALLBACK *callback__)(cef_dictionary_value_t *, cef_string_t *)) { return callback__(self, key); }
	// cef_string_userfree_t gocef_dictionary_value_get_string(cef_dictionary_value_t * self, cef_string_t * key, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_dictionary_value_t *, cef_string_t *)) { return callback__(self, key); }
	// cef_binary_value_t * gocef_dictionary_value_get_binary(cef_dictionary_value_t * self, cef_string_t * key, cef_binary_value_t * (CEF_CALLBACK *callback__)(cef_dictionary_value_t *, cef_string_t *)) { return callback__(self, key); }
	// cef_dictionary_value_t * gocef_dictionary_value_get_dictionary(cef_dictionary_value_t * self, cef_string_t * key, cef_dictionary_value_t * (CEF_CALLBACK *callback__)(cef_dictionary_value_t *, cef_string_t *)) { return callback__(self, key); }
	// cef_list_value_t * gocef_dictionary_value_get_list(cef_dictionary_value_t * self, cef_string_t * key, cef_list_value_t * (CEF_CALLBACK *callback__)(cef_dictionary_value_t *, cef_string_t *)) { return callback__(self, key); }
	// int gocef_dictionary_value_set_value(cef_dictionary_value_t * self, cef_string_t * key, cef_value_t * value, int (CEF_CALLBACK *callback__)(cef_dictionary_value_t *, cef_string_t *, cef_value_t *)) { return callback__(self, key, value); }
	// int gocef_dictionary_value_set_null(cef_dictionary_value_t * self, cef_string_t * key, int (CEF_CALLBACK *callback__)(cef_dictionary_value_t *, cef_string_t *)) { return callback__(self, key); }
	// int gocef_dictionary_value_set_bool(cef_dictionary_value_t * self, cef_string_t * key, int value, int (CEF_CALLBACK *callback__)(cef_dictionary_value_t *, cef_string_t *, int)) { return callback__(self, key, value); }
	// int gocef_dictionary_value_set_int(cef_dictionary_value_t * self, cef_string_t * key, int value, int (CEF_CALLBACK *callback__)(cef_dictionary_value_t *, cef_string_t *, int)) { return callback__(self, key, value); }
	// int gocef_dictionary_value_set_double(cef_dictionary_value_t * self, cef_string_t * key, double value, int (CEF_CALLBACK *callback__)(cef_dictionary_value_t *, cef_string_t *, double)) { return callback__(self, key, value); }
	// int gocef_dictionary_value_set_string(cef_dictionary_value_t * self, cef_string_t * key, cef_string_t * value, int (CEF_CALLBACK *callback__)(cef_dictionary_value_t *, cef_string_t *, cef_string_t *)) { return callback__(self, key, value); }
	// int gocef_dictionary_value_set_binary(cef_dictionary_value_t * self, cef_string_t * key, cef_binary_value_t * value, int (CEF_CALLBACK *callback__)(cef_dictionary_value_t *, cef_string_t *, cef_binary_value_t *)) { return callback__(self, key, value); }
	// int gocef_dictionary_value_set_dictionary(cef_dictionary_value_t * self, cef_string_t * key, cef_dictionary_value_t * value, int (CEF_CALLBACK *callback__)(cef_dictionary_value_t *, cef_string_t *, cef_dictionary_value_t *)) { return callback__(self, key, value); }
	// int gocef_dictionary_value_set_list(cef_dictionary_value_t * self, cef_string_t * key, cef_list_value_t * value, int (CEF_CALLBACK *callback__)(cef_dictionary_value_t *, cef_string_t *, cef_list_value_t *)) { return callback__(self, key, value); }
	"C"
)

// DictionaryValue (cef_dictionary_value_t from include/capi/cef_values_capi.h)
// Structure representing a dictionary value. Can be used on any process and
// thread.
type DictionaryValue C.cef_dictionary_value_t

func (d *DictionaryValue) toNative() *C.cef_dictionary_value_t {
	return (*C.cef_dictionary_value_t)(d)
}

// Base (base)
// Base structure.
func (d *DictionaryValue) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// IsValid (is_valid)
// Returns true (1) if this object is valid. This object may become invalid if
// the underlying data is owned by another object (e.g. list or dictionary)
// and that other object is then modified or destroyed. Do not call any other
// functions if this function returns false (0).
func (d *DictionaryValue) IsValid() int32 {
	return int32(C.gocef_dictionary_value_is_valid(d.toNative(), d.is_valid))
}

// IsOwned (is_owned)
// Returns true (1) if this object is currently owned by another object.
func (d *DictionaryValue) IsOwned() int32 {
	return int32(C.gocef_dictionary_value_is_owned(d.toNative(), d.is_owned))
}

// IsReadOnly (is_read_only)
// Returns true (1) if the values of this object are read-only. Some APIs may
// expose read-only objects.
func (d *DictionaryValue) IsReadOnly() int32 {
	return int32(C.gocef_dictionary_value_is_read_only(d.toNative(), d.is_read_only))
}

// IsSame (is_same)
// Returns true (1) if this object and |that| object have the same underlying
// data. If true (1) modifications to this object will also affect |that|
// object and vice-versa.
func (d *DictionaryValue) IsSame(that *DictionaryValue) int32 {
	return int32(C.gocef_dictionary_value_is_same(d.toNative(), that.toNative(), d.is_same))
}

// IsEqual (is_equal)
// Returns true (1) if this object and |that| object have an equivalent
// underlying value but are not necessarily the same object.
func (d *DictionaryValue) IsEqual(that *DictionaryValue) int32 {
	return int32(C.gocef_dictionary_value_is_equal(d.toNative(), that.toNative(), d.is_equal))
}

// Copy (copy)
// Returns a writable copy of this object. If |exclude_NULL_children| is true
// (1) any NULL dictionaries or lists will be excluded from the copy.
func (d *DictionaryValue) Copy(excludeEmptyChildren int32) *DictionaryValue {
	return (*DictionaryValue)(C.gocef_dictionary_value_copy(d.toNative(), C.int(excludeEmptyChildren), d.copy))
}

// GetSize (get_size)
// Returns the number of values.
func (d *DictionaryValue) GetSize() uint64 {
	return uint64(C.gocef_dictionary_value_get_size(d.toNative(), d.get_size))
}

// Clear (clear)
// Removes all values. Returns true (1) on success.
func (d *DictionaryValue) Clear() int32 {
	return int32(C.gocef_dictionary_value_clear(d.toNative(), d.clear))
}

// HasKey (has_key)
// Returns true (1) if the current dictionary has a value for the given key.
func (d *DictionaryValue) HasKey(key string) int32 {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	return int32(C.gocef_dictionary_value_has_key(d.toNative(), (*C.cef_string_t)(key_), d.has_key))
}

// GetKeys (get_keys)
// Reads all keys for this dictionary into the specified vector.
func (d *DictionaryValue) GetKeys(keys StringList) int32 {
	return int32(C.gocef_dictionary_value_get_keys(d.toNative(), C.cef_string_list_t(keys), d.get_keys))
}

// Remove (remove)
// Removes the value at the specified key. Returns true (1) is the value was
// removed successfully.
func (d *DictionaryValue) Remove(key string) int32 {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	return int32(C.gocef_dictionary_value_remove(d.toNative(), (*C.cef_string_t)(key_), d.remove))
}

// GetType (get_type)
// Returns the value type for the specified key.
func (d *DictionaryValue) GetType(key string) ValueType {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	return ValueType(C.gocef_dictionary_value_get_type(d.toNative(), (*C.cef_string_t)(key_), d.get_type))
}

// GetValue (get_value)
// Returns the value at the specified key. For simple types the returned value
// will copy existing data and modifications to the value will not modify this
// object. For complex types (binary, dictionary and list) the returned value
// will reference existing data and modifications to the value will modify
// this object.
func (d *DictionaryValue) GetValue(key string) *Value {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	return (*Value)(C.gocef_dictionary_value_get_value(d.toNative(), (*C.cef_string_t)(key_), d.get_value))
}

// GetBool (get_bool)
// Returns the value at the specified key as type bool.
func (d *DictionaryValue) GetBool(key string) int32 {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	return int32(C.gocef_dictionary_value_get_bool(d.toNative(), (*C.cef_string_t)(key_), d.get_bool))
}

// GetInt (get_int)
// Returns the value at the specified key as type int.
func (d *DictionaryValue) GetInt(key string) int32 {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	return int32(C.gocef_dictionary_value_get_int(d.toNative(), (*C.cef_string_t)(key_), d.get_int))
}

// GetDouble (get_double)
// Returns the value at the specified key as type double.
func (d *DictionaryValue) GetDouble(key string) float64 {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	return float64(C.gocef_dictionary_value_get_double(d.toNative(), (*C.cef_string_t)(key_), d.get_double))
}

// GetString (get_string)
// Returns the value at the specified key as type string.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *DictionaryValue) GetString(key string) string {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	return cefuserfreestrToString(C.gocef_dictionary_value_get_string(d.toNative(), (*C.cef_string_t)(key_), d.get_string))
}

// GetBinary (get_binary)
// Returns the value at the specified key as type binary. The returned value
// will reference existing data.
func (d *DictionaryValue) GetBinary(key string) *BinaryValue {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	return (*BinaryValue)(C.gocef_dictionary_value_get_binary(d.toNative(), (*C.cef_string_t)(key_), d.get_binary))
}

// GetDictionary (get_dictionary)
// Returns the value at the specified key as type dictionary. The returned
// value will reference existing data and modifications to the value will
// modify this object.
func (d *DictionaryValue) GetDictionary(key string) *DictionaryValue {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	return (*DictionaryValue)(C.gocef_dictionary_value_get_dictionary(d.toNative(), (*C.cef_string_t)(key_), d.get_dictionary))
}

// GetList (get_list)
// Returns the value at the specified key as type list. The returned value
// will reference existing data and modifications to the value will modify
// this object.
func (d *DictionaryValue) GetList(key string) *ListValue {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	return (*ListValue)(C.gocef_dictionary_value_get_list(d.toNative(), (*C.cef_string_t)(key_), d.get_list))
}

// SetValue (set_value)
// Sets the value at the specified key. Returns true (1) if the value was set
// successfully. If |value| represents simple data then the underlying data
// will be copied and modifications to |value| will not modify this object. If
// |value| represents complex data (binary, dictionary or list) then the
// underlying data will be referenced and modifications to |value| will modify
// this object.
func (d *DictionaryValue) SetValue(key string, value *Value) int32 {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	return int32(C.gocef_dictionary_value_set_value(d.toNative(), (*C.cef_string_t)(key_), value.toNative(), d.set_value))
}

// SetNull (set_null)
// Sets the value at the specified key as type null. Returns true (1) if the
// value was set successfully.
func (d *DictionaryValue) SetNull(key string) int32 {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	return int32(C.gocef_dictionary_value_set_null(d.toNative(), (*C.cef_string_t)(key_), d.set_null))
}

// SetBool (set_bool)
// Sets the value at the specified key as type bool. Returns true (1) if the
// value was set successfully.
func (d *DictionaryValue) SetBool(key string, value int32) int32 {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	return int32(C.gocef_dictionary_value_set_bool(d.toNative(), (*C.cef_string_t)(key_), C.int(value), d.set_bool))
}

// SetInt (set_int)
// Sets the value at the specified key as type int. Returns true (1) if the
// value was set successfully.
func (d *DictionaryValue) SetInt(key string, value int32) int32 {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	return int32(C.gocef_dictionary_value_set_int(d.toNative(), (*C.cef_string_t)(key_), C.int(value), d.set_int))
}

// SetDouble (set_double)
// Sets the value at the specified key as type double. Returns true (1) if the
// value was set successfully.
func (d *DictionaryValue) SetDouble(key string, value float64) int32 {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	return int32(C.gocef_dictionary_value_set_double(d.toNative(), (*C.cef_string_t)(key_), C.double(value), d.set_double))
}

// SetString (set_string)
// Sets the value at the specified key as type string. Returns true (1) if the
// value was set successfully.
func (d *DictionaryValue) SetString(key, value string) int32 {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	value_ := C.cef_string_userfree_alloc()
	setCEFStr(value, value_)
	defer func() {
		C.cef_string_userfree_free(value_)
	}()
	return int32(C.gocef_dictionary_value_set_string(d.toNative(), (*C.cef_string_t)(key_), (*C.cef_string_t)(value_), d.set_string))
}

// SetBinary (set_binary)
// Sets the value at the specified key as type binary. Returns true (1) if the
// value was set successfully. If |value| is currently owned by another object
// then the value will be copied and the |value| reference will not change.
// Otherwise, ownership will be transferred to this object and the |value|
// reference will be invalidated.
func (d *DictionaryValue) SetBinary(key string, value *BinaryValue) int32 {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	return int32(C.gocef_dictionary_value_set_binary(d.toNative(), (*C.cef_string_t)(key_), value.toNative(), d.set_binary))
}

// SetDictionary (set_dictionary)
// Sets the value at the specified key as type dict. Returns true (1) if the
// value was set successfully. If |value| is currently owned by another object
// then the value will be copied and the |value| reference will not change.
// Otherwise, ownership will be transferred to this object and the |value|
// reference will be invalidated.
func (d *DictionaryValue) SetDictionary(key string, value *DictionaryValue) int32 {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	return int32(C.gocef_dictionary_value_set_dictionary(d.toNative(), (*C.cef_string_t)(key_), value.toNative(), d.set_dictionary))
}

// SetList (set_list)
// Sets the value at the specified key as type list. Returns true (1) if the
// value was set successfully. If |value| is currently owned by another object
// then the value will be copied and the |value| reference will not change.
// Otherwise, ownership will be transferred to this object and the |value|
// reference will be invalidated.
func (d *DictionaryValue) SetList(key string, value *ListValue) int32 {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	return int32(C.gocef_dictionary_value_set_list(d.toNative(), (*C.cef_string_t)(key_), value.toNative(), d.set_list))
}
