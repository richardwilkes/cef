// Code generated - DO NOT EDIT.

package cef

import (
	// #include "capi_gen.h"
	"C"
)

// Time (cef_time_t from include/internal/cef_time.h)
// Time information. Values should always be in UTC.
type Time struct {
	// Year (year)
	Year int32
	// Month (month)
	//   Windows, 1970 to 2038 on 32-bit POSIX)
	Month int32
	// DayOfWeek (day_of_week)
	DayOfWeek int32
	// DayOfMonth (day_of_month)
	DayOfMonth int32
	// Hour (hour)
	Hour int32
	// Minute (minute)
	Minute int32
	// Second (second)
	Second int32
	// Millisecond (millisecond)
	//   seconds which may take it up to 60).
	Millisecond int32
}

// NewTime creates a new Time.
func NewTime() *Time {
	return &Time{}
}

func (d *Time) toNative(native *C.cef_time_t) *C.cef_time_t {
	native.year = C.int(d.Year)
	native.month = C.int(d.Month)
	native.day_of_week = C.int(d.DayOfWeek)
	native.day_of_month = C.int(d.DayOfMonth)
	native.hour = C.int(d.Hour)
	native.minute = C.int(d.Minute)
	native.second = C.int(d.Second)
	native.millisecond = C.int(d.Millisecond)
	return native
}

func (d *Time) fromNative(native *C.cef_time_t) *Time {
	d.Year = int32(native.year)
	d.Month = int32(native.month)
	d.DayOfWeek = int32(native.day_of_week)
	d.DayOfMonth = int32(native.day_of_month)
	d.Hour = int32(native.hour)
	d.Minute = int32(native.minute)
	d.Second = int32(native.second)
	d.Millisecond = int32(native.millisecond)
	return d
}
