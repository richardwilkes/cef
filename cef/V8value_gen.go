// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// int gocef_v8value_is_valid(cef_v8value_t * self, int (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// int gocef_v8value_is_undefined(cef_v8value_t * self, int (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// int gocef_v8value_is_null(cef_v8value_t * self, int (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// int gocef_v8value_is_bool(cef_v8value_t * self, int (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// int gocef_v8value_is_int(cef_v8value_t * self, int (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// int gocef_v8value_is_uint(cef_v8value_t * self, int (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// int gocef_v8value_is_double(cef_v8value_t * self, int (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// int gocef_v8value_is_date(cef_v8value_t * self, int (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// int gocef_v8value_is_string(cef_v8value_t * self, int (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// int gocef_v8value_is_object(cef_v8value_t * self, int (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// int gocef_v8value_is_array(cef_v8value_t * self, int (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// int gocef_v8value_is_array_buffer(cef_v8value_t * self, int (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// int gocef_v8value_is_function(cef_v8value_t * self, int (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// int gocef_v8value_is_same(cef_v8value_t * self, cef_v8value_t * that, int (CEF_CALLBACK *callback__)(cef_v8value_t *, cef_v8value_t *)) { return callback__(self, that); }
	// int gocef_v8value_get_bool_value(cef_v8value_t * self, int (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// int32 gocef_v8value_get_int_value(cef_v8value_t * self, int32 (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// uint32 gocef_v8value_get_uint_value(cef_v8value_t * self, uint32 (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// double gocef_v8value_get_double_value(cef_v8value_t * self, double (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// cef_time_t gocef_v8value_get_date_value(cef_v8value_t * self, cef_time_t (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_v8value_get_string_value(cef_v8value_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// int gocef_v8value_is_user_created(cef_v8value_t * self, int (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// int gocef_v8value_has_exception(cef_v8value_t * self, int (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// cef_v8exception_t * gocef_v8value_get_exception(cef_v8value_t * self, cef_v8exception_t * (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// int gocef_v8value_clear_exception(cef_v8value_t * self, int (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// int gocef_v8value_will_rethrow_exceptions(cef_v8value_t * self, int (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// int gocef_v8value_set_rethrow_exceptions(cef_v8value_t * self, int rethrow, int (CEF_CALLBACK *callback__)(cef_v8value_t *, int)) { return callback__(self, rethrow); }
	// int gocef_v8value_has_value_bykey(cef_v8value_t * self, cef_string_t * key, int (CEF_CALLBACK *callback__)(cef_v8value_t *, cef_string_t *)) { return callback__(self, key); }
	// int gocef_v8value_has_value_byindex(cef_v8value_t * self, int index, int (CEF_CALLBACK *callback__)(cef_v8value_t *, int)) { return callback__(self, index); }
	// int gocef_v8value_delete_value_bykey(cef_v8value_t * self, cef_string_t * key, int (CEF_CALLBACK *callback__)(cef_v8value_t *, cef_string_t *)) { return callback__(self, key); }
	// int gocef_v8value_delete_value_byindex(cef_v8value_t * self, int index, int (CEF_CALLBACK *callback__)(cef_v8value_t *, int)) { return callback__(self, index); }
	// cef_v8value_t * gocef_v8value_get_value_bykey(cef_v8value_t * self, cef_string_t * key, cef_v8value_t * (CEF_CALLBACK *callback__)(cef_v8value_t *, cef_string_t *)) { return callback__(self, key); }
	// cef_v8value_t * gocef_v8value_get_value_byindex(cef_v8value_t * self, int index, cef_v8value_t * (CEF_CALLBACK *callback__)(cef_v8value_t *, int)) { return callback__(self, index); }
	// int gocef_v8value_set_value_bykey(cef_v8value_t * self, cef_string_t * key, cef_v8value_t * value, cef_v8_propertyattribute_t attribute, int (CEF_CALLBACK *callback__)(cef_v8value_t *, cef_string_t *, cef_v8value_t *, cef_v8_propertyattribute_t)) { return callback__(self, key, value, attribute); }
	// int gocef_v8value_set_value_byindex(cef_v8value_t * self, int index, cef_v8value_t * value, int (CEF_CALLBACK *callback__)(cef_v8value_t *, int, cef_v8value_t *)) { return callback__(self, index, value); }
	// int gocef_v8value_set_value_byaccessor(cef_v8value_t * self, cef_string_t * key, cef_v8_accesscontrol_t settings, cef_v8_propertyattribute_t attribute, int (CEF_CALLBACK *callback__)(cef_v8value_t *, cef_string_t *, cef_v8_accesscontrol_t, cef_v8_propertyattribute_t)) { return callback__(self, key, settings, attribute); }
	// int gocef_v8value_get_keys(cef_v8value_t * self, cef_string_list_t keys, int (CEF_CALLBACK *callback__)(cef_v8value_t *, cef_string_list_t)) { return callback__(self, keys); }
	// int gocef_v8value_set_user_data(cef_v8value_t * self, cef_base_ref_counted_t * userData, int (CEF_CALLBACK *callback__)(cef_v8value_t *, cef_base_ref_counted_t *)) { return callback__(self, userData); }
	// cef_base_ref_counted_t * gocef_v8value_get_user_data(cef_v8value_t * self, cef_base_ref_counted_t * (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// int gocef_v8value_get_externally_allocated_memory(cef_v8value_t * self, int (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// int gocef_v8value_adjust_externally_allocated_memory(cef_v8value_t * self, int changeInBytes, int (CEF_CALLBACK *callback__)(cef_v8value_t *, int)) { return callback__(self, changeInBytes); }
	// int gocef_v8value_get_array_length(cef_v8value_t * self, int (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// cef_v8array_buffer_release_callback_t * gocef_v8value_get_array_buffer_release_callback(cef_v8value_t * self, cef_v8array_buffer_release_callback_t * (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// int gocef_v8value_neuter_array_buffer(cef_v8value_t * self, int (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_v8value_get_function_name(cef_v8value_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// cef_v8handler_t * gocef_v8value_get_function_handler(cef_v8value_t * self, cef_v8handler_t * (CEF_CALLBACK *callback__)(cef_v8value_t *)) { return callback__(self); }
	// cef_v8value_t * gocef_v8value_execute_function(cef_v8value_t * self, cef_v8value_t * object, size_t argumentsCount, cef_v8value_t ** arguments, cef_v8value_t * (CEF_CALLBACK *callback__)(cef_v8value_t *, cef_v8value_t *, size_t, cef_v8value_t **)) { return callback__(self, object, argumentsCount, arguments); }
	// cef_v8value_t * gocef_v8value_execute_function_with_context(cef_v8value_t * self, cef_v8context_t * context, cef_v8value_t * object, size_t argumentsCount, cef_v8value_t ** arguments, cef_v8value_t * (CEF_CALLBACK *callback__)(cef_v8value_t *, cef_v8context_t *, cef_v8value_t *, size_t, cef_v8value_t **)) { return callback__(self, context, object, argumentsCount, arguments); }
	"C"
)

// V8value (cef_v8value_t from include/capi/cef_v8_capi.h)
// Structure representing a V8 value handle. V8 handles can only be accessed
// from the thread on which they are created. Valid threads for creating a V8
// handle include the render process main thread (TID_RENDERER) and WebWorker
// threads. A task runner for posting tasks on the associated thread can be
// retrieved via the cef_v8context_t::get_task_runner() function.
type V8value C.cef_v8value_t

func (d *V8value) toNative() *C.cef_v8value_t {
	return (*C.cef_v8value_t)(d)
}

// Base (base)
// Base structure.
func (d *V8value) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// IsValid (is_valid)
// Returns true (1) if the underlying handle is valid and it can be accessed
// on the current thread. Do not call any other functions if this function
// returns false (0).
func (d *V8value) IsValid() int32 {
	return int32(C.gocef_v8value_is_valid(d.toNative(), d.is_valid))
}

// IsUndefined (is_undefined)
// True if the value type is undefined.
func (d *V8value) IsUndefined() int32 {
	return int32(C.gocef_v8value_is_undefined(d.toNative(), d.is_undefined))
}

// IsNull (is_null)
// True if the value type is null.
func (d *V8value) IsNull() int32 {
	return int32(C.gocef_v8value_is_null(d.toNative(), d.is_null))
}

// IsBool (is_bool)
// True if the value type is bool.
func (d *V8value) IsBool() int32 {
	return int32(C.gocef_v8value_is_bool(d.toNative(), d.is_bool))
}

// IsInt (is_int)
// True if the value type is int.
func (d *V8value) IsInt() int32 {
	return int32(C.gocef_v8value_is_int(d.toNative(), d.is_int))
}

// IsUint (is_uint)
// True if the value type is unsigned int.
func (d *V8value) IsUint() int32 {
	return int32(C.gocef_v8value_is_uint(d.toNative(), d.is_uint))
}

// IsDouble (is_double)
// True if the value type is double.
func (d *V8value) IsDouble() int32 {
	return int32(C.gocef_v8value_is_double(d.toNative(), d.is_double))
}

// IsDate (is_date)
// True if the value type is Date.
func (d *V8value) IsDate() int32 {
	return int32(C.gocef_v8value_is_date(d.toNative(), d.is_date))
}

// IsString (is_string)
// True if the value type is string.
func (d *V8value) IsString() int32 {
	return int32(C.gocef_v8value_is_string(d.toNative(), d.is_string))
}

// IsObject (is_object)
// True if the value type is object.
func (d *V8value) IsObject() int32 {
	return int32(C.gocef_v8value_is_object(d.toNative(), d.is_object))
}

// IsArray (is_array)
// True if the value type is array.
func (d *V8value) IsArray() int32 {
	return int32(C.gocef_v8value_is_array(d.toNative(), d.is_array))
}

// IsArrayBuffer (is_array_buffer)
// True if the value type is an ArrayBuffer.
func (d *V8value) IsArrayBuffer() int32 {
	return int32(C.gocef_v8value_is_array_buffer(d.toNative(), d.is_array_buffer))
}

// IsFunction (is_function)
// True if the value type is function.
func (d *V8value) IsFunction() int32 {
	return int32(C.gocef_v8value_is_function(d.toNative(), d.is_function))
}

// IsSame (is_same)
// Returns true (1) if this object is pointing to the same handle as |that|
// object.
func (d *V8value) IsSame(that *V8value) int32 {
	return int32(C.gocef_v8value_is_same(d.toNative(), that.toNative(), d.is_same))
}

// GetBoolValue (get_bool_value)
// Return a bool value.
func (d *V8value) GetBoolValue() int32 {
	return int32(C.gocef_v8value_get_bool_value(d.toNative(), d.get_bool_value))
}

// GetIntValue (get_int_value)
// Return an int value.
func (d *V8value) GetIntValue() int32 {
	return int32(C.gocef_v8value_get_int_value(d.toNative(), d.get_int_value))
}

// GetUintValue (get_uint_value)
// Return an unsigned int value.
func (d *V8value) GetUintValue() uint32 {
	return uint32(C.gocef_v8value_get_uint_value(d.toNative(), d.get_uint_value))
}

// GetDoubleValue (get_double_value)
// Return a double value.
func (d *V8value) GetDoubleValue() float64 {
	return float64(C.gocef_v8value_get_double_value(d.toNative(), d.get_double_value))
}

// GetDateValue (get_date_value)
// Return a Date value.
func (d *V8value) GetDateValue() Time {
	cresult__ := C.gocef_v8value_get_date_value(d.toNative(), d.get_date_value)
	var result__ Time
	(&cresult__).intoGo(&result__)
	return result__
}

// GetStringValue (get_string_value)
// Return a string value.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *V8value) GetStringValue() string {
	return cefuserfreestrToString(C.gocef_v8value_get_string_value(d.toNative(), d.get_string_value))
}

// IsUserCreated (is_user_created)
// Returns true (1) if this is a user created object.
func (d *V8value) IsUserCreated() int32 {
	return int32(C.gocef_v8value_is_user_created(d.toNative(), d.is_user_created))
}

// HasException (has_exception)
// Returns true (1) if the last function call resulted in an exception. This
// attribute exists only in the scope of the current CEF value object.
func (d *V8value) HasException() int32 {
	return int32(C.gocef_v8value_has_exception(d.toNative(), d.has_exception))
}

// GetException (get_exception)
// Returns the exception resulting from the last function call. This attribute
// exists only in the scope of the current CEF value object.
func (d *V8value) GetException() *V8exception {
	return (*V8exception)(C.gocef_v8value_get_exception(d.toNative(), d.get_exception))
}

// ClearException (clear_exception)
// Clears the last exception and returns true (1) on success.
func (d *V8value) ClearException() int32 {
	return int32(C.gocef_v8value_clear_exception(d.toNative(), d.clear_exception))
}

// WillRethrowExceptions (will_rethrow_exceptions)
// Returns true (1) if this object will re-throw future exceptions. This
// attribute exists only in the scope of the current CEF value object.
func (d *V8value) WillRethrowExceptions() int32 {
	return int32(C.gocef_v8value_will_rethrow_exceptions(d.toNative(), d.will_rethrow_exceptions))
}

// SetRethrowExceptions (set_rethrow_exceptions)
// Set whether this object will re-throw future exceptions. By default
// exceptions are not re-thrown. If a exception is re-thrown the current
// context should not be accessed again until after the exception has been
// caught and not re-thrown. Returns true (1) on success. This attribute
// exists only in the scope of the current CEF value object.
func (d *V8value) SetRethrowExceptions(rethrow int32) int32 {
	return int32(C.gocef_v8value_set_rethrow_exceptions(d.toNative(), C.int(rethrow), d.set_rethrow_exceptions))
}

// HasValueBykey (has_value_bykey)
// Returns true (1) if the object has a value with the specified identifier.
func (d *V8value) HasValueBykey(key string) int32 {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	return int32(C.gocef_v8value_has_value_bykey(d.toNative(), (*C.cef_string_t)(key_), d.has_value_bykey))
}

// HasValueByindex (has_value_byindex)
// Returns true (1) if the object has a value with the specified identifier.
func (d *V8value) HasValueByindex(index int32) int32 {
	return int32(C.gocef_v8value_has_value_byindex(d.toNative(), C.int(index), d.has_value_byindex))
}

// DeleteValueBykey (delete_value_bykey)
// Deletes the value with the specified identifier and returns true (1) on
// success. Returns false (0) if this function is called incorrectly or an
// exception is thrown. For read-only and don't-delete values this function
// will return true (1) even though deletion failed.
func (d *V8value) DeleteValueBykey(key string) int32 {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	return int32(C.gocef_v8value_delete_value_bykey(d.toNative(), (*C.cef_string_t)(key_), d.delete_value_bykey))
}

// DeleteValueByindex (delete_value_byindex)
// Deletes the value with the specified identifier and returns true (1) on
// success. Returns false (0) if this function is called incorrectly, deletion
// fails or an exception is thrown. For read-only and don't-delete values this
// function will return true (1) even though deletion failed.
func (d *V8value) DeleteValueByindex(index int32) int32 {
	return int32(C.gocef_v8value_delete_value_byindex(d.toNative(), C.int(index), d.delete_value_byindex))
}

// GetValueBykey (get_value_bykey)
// Returns the value with the specified identifier on success. Returns NULL if
// this function is called incorrectly or an exception is thrown.
func (d *V8value) GetValueBykey(key string) *V8value {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	return (*V8value)(C.gocef_v8value_get_value_bykey(d.toNative(), (*C.cef_string_t)(key_), d.get_value_bykey))
}

// GetValueByindex (get_value_byindex)
// Returns the value with the specified identifier on success. Returns NULL if
// this function is called incorrectly or an exception is thrown.
func (d *V8value) GetValueByindex(index int32) *V8value {
	return (*V8value)(C.gocef_v8value_get_value_byindex(d.toNative(), C.int(index), d.get_value_byindex))
}

// SetValueBykey (set_value_bykey)
// Associates a value with the specified identifier and returns true (1) on
// success. Returns false (0) if this function is called incorrectly or an
// exception is thrown. For read-only values this function will return true
// (1) even though assignment failed.
func (d *V8value) SetValueBykey(key string, value *V8value, attribute V8Propertyattribute) int32 {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	return int32(C.gocef_v8value_set_value_bykey(d.toNative(), (*C.cef_string_t)(key_), value.toNative(), C.cef_v8_propertyattribute_t(attribute), d.set_value_bykey))
}

// SetValueByindex (set_value_byindex)
// Associates a value with the specified identifier and returns true (1) on
// success. Returns false (0) if this function is called incorrectly or an
// exception is thrown. For read-only values this function will return true
// (1) even though assignment failed.
func (d *V8value) SetValueByindex(index int32, value *V8value) int32 {
	return int32(C.gocef_v8value_set_value_byindex(d.toNative(), C.int(index), value.toNative(), d.set_value_byindex))
}

// SetValueByaccessor (set_value_byaccessor)
// Registers an identifier and returns true (1) on success. Access to the
// identifier will be forwarded to the cef_v8accessor_t instance passed to
// cef_v8value_t::cef_v8value_create_object(). Returns false (0) if this
// function is called incorrectly or an exception is thrown. For read-only
// values this function will return true (1) even though assignment failed.
func (d *V8value) SetValueByaccessor(key string, settings V8Accesscontrol, attribute V8Propertyattribute) int32 {
	key_ := C.cef_string_userfree_alloc()
	setCEFStr(key, key_)
	defer func() {
		C.cef_string_userfree_free(key_)
	}()
	return int32(C.gocef_v8value_set_value_byaccessor(d.toNative(), (*C.cef_string_t)(key_), C.cef_v8_accesscontrol_t(settings), C.cef_v8_propertyattribute_t(attribute), d.set_value_byaccessor))
}

// GetKeys (get_keys)
// Read the keys for the object's values into the specified vector. Integer-
// based keys will also be returned as strings.
func (d *V8value) GetKeys(keys StringList) int32 {
	return int32(C.gocef_v8value_get_keys(d.toNative(), C.cef_string_list_t(keys), d.get_keys))
}

// SetUserData (set_user_data)
// Sets the user data for this object and returns true (1) on success. Returns
// false (0) if this function is called incorrectly. This function can only be
// called on user created objects.
func (d *V8value) SetUserData(userData *BaseRefCounted) int32 {
	return int32(C.gocef_v8value_set_user_data(d.toNative(), userData.toNative(), d.set_user_data))
}

// GetUserData (get_user_data)
// Returns the user data, if any, assigned to this object.
func (d *V8value) GetUserData() *BaseRefCounted {
	return (*BaseRefCounted)(C.gocef_v8value_get_user_data(d.toNative(), d.get_user_data))
}

// GetExternallyAllocatedMemory (get_externally_allocated_memory)
// Returns the amount of externally allocated memory registered for the
// object.
func (d *V8value) GetExternallyAllocatedMemory() int32 {
	return int32(C.gocef_v8value_get_externally_allocated_memory(d.toNative(), d.get_externally_allocated_memory))
}

// AdjustExternallyAllocatedMemory (adjust_externally_allocated_memory)
// Adjusts the amount of registered external memory for the object. Used to
// give V8 an indication of the amount of externally allocated memory that is
// kept alive by JavaScript objects. V8 uses this information to decide when
// to perform global garbage collection. Each cef_v8value_t tracks the amount
// of external memory associated with it and automatically decreases the
// global total by the appropriate amount on its destruction.
// |change_in_bytes| specifies the number of bytes to adjust by. This function
// returns the number of bytes associated with the object after the
// adjustment. This function can only be called on user created objects.
func (d *V8value) AdjustExternallyAllocatedMemory(changeInBytes int32) int32 {
	return int32(C.gocef_v8value_adjust_externally_allocated_memory(d.toNative(), C.int(changeInBytes), d.adjust_externally_allocated_memory))
}

// GetArrayLength (get_array_length)
// Returns the number of elements in the array.
func (d *V8value) GetArrayLength() int32 {
	return int32(C.gocef_v8value_get_array_length(d.toNative(), d.get_array_length))
}

// GetArrayBufferReleaseCallback (get_array_buffer_release_callback)
// Returns the ReleaseCallback object associated with the ArrayBuffer or NULL
// if the ArrayBuffer was not created with CreateArrayBuffer.
func (d *V8value) GetArrayBufferReleaseCallback() *V8arrayBufferReleaseCallback {
	return (*V8arrayBufferReleaseCallback)(C.gocef_v8value_get_array_buffer_release_callback(d.toNative(), d.get_array_buffer_release_callback))
}

// NeuterArrayBuffer (neuter_array_buffer)
// Prevent the ArrayBuffer from using it's memory block by setting the length
// to zero. This operation cannot be undone. If the ArrayBuffer was created
// with CreateArrayBuffer then
// cef_v8array_buffer_release_callback_t::ReleaseBuffer will be called to
// release the underlying buffer.
func (d *V8value) NeuterArrayBuffer() int32 {
	return int32(C.gocef_v8value_neuter_array_buffer(d.toNative(), d.neuter_array_buffer))
}

// GetFunctionName (get_function_name)
// Returns the function name.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *V8value) GetFunctionName() string {
	return cefuserfreestrToString(C.gocef_v8value_get_function_name(d.toNative(), d.get_function_name))
}

// GetFunctionHandler (get_function_handler)
// Returns the function handler or NULL if not a CEF-created function.
func (d *V8value) GetFunctionHandler() *V8handler {
	return (*V8handler)(C.gocef_v8value_get_function_handler(d.toNative(), d.get_function_handler))
}

// ExecuteFunction (execute_function)
// Execute the function using the current V8 context. This function should
// only be called from within the scope of a cef_v8handler_t or
// cef_v8accessor_t callback, or in combination with calling enter() and
// exit() on a stored cef_v8context_t reference. |object| is the receiver
// ('this' object) of the function. If |object| is NULL the current context's
// global object will be used. |arguments| is the list of arguments that will
// be passed to the function. Returns the function return value on success.
// Returns NULL if this function is called incorrectly or an exception is
// thrown.
func (d *V8value) ExecuteFunction(object *V8value, argumentsCount uint64, arguments **V8value) *V8value {
	arguments_ := (*arguments).toNative()
	return (*V8value)(C.gocef_v8value_execute_function(d.toNative(), object.toNative(), C.size_t(argumentsCount), &arguments_, d.execute_function))
}

// ExecuteFunctionWithContext (execute_function_with_context)
// Execute the function using the specified V8 context. |object| is the
// receiver ('this' object) of the function. If |object| is NULL the specified
// context's global object will be used. |arguments| is the list of arguments
// that will be passed to the function. Returns the function return value on
// success. Returns NULL if this function is called incorrectly or an
// exception is thrown.
func (d *V8value) ExecuteFunctionWithContext(context *V8context, object *V8value, argumentsCount uint64, arguments **V8value) *V8value {
	arguments_ := (*arguments).toNative()
	return (*V8value)(C.gocef_v8value_execute_function_with_context(d.toNative(), context.toNative(), object.toNative(), C.size_t(argumentsCount), &arguments_, d.execute_function_with_context))
}
