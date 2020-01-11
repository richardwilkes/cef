// Code created from "class.go.tmpl" - don't edit by hand

package cef

import "unsafe"

import (
	// #include "capi_gen.h"
	// int gocef_command_line_is_valid(cef_command_line_t * self, int (CEF_CALLBACK *callback__)(cef_command_line_t *)) { return callback__(self); }
	// int gocef_command_line_is_read_only(cef_command_line_t * self, int (CEF_CALLBACK *callback__)(cef_command_line_t *)) { return callback__(self); }
	// cef_command_line_t * gocef_command_line_copy(cef_command_line_t * self, cef_command_line_t * (CEF_CALLBACK *callback__)(cef_command_line_t *)) { return callback__(self); }
	// void gocef_command_line_init_from_argv(cef_command_line_t * self, int argc, char ** argv, void (CEF_CALLBACK *callback__)(cef_command_line_t *, int, char **)) { return callback__(self, argc, argv); }
	// void gocef_command_line_init_from_string(cef_command_line_t * self, cef_string_t * commandLine, void (CEF_CALLBACK *callback__)(cef_command_line_t *, cef_string_t *)) { return callback__(self, commandLine); }
	// void gocef_command_line_reset(cef_command_line_t * self, void (CEF_CALLBACK *callback__)(cef_command_line_t *)) { return callback__(self); }
	// void gocef_command_line_get_argv(cef_command_line_t * self, cef_string_list_t argv, void (CEF_CALLBACK *callback__)(cef_command_line_t *, cef_string_list_t)) { return callback__(self, argv); }
	// cef_string_userfree_t gocef_command_line_get_command_line_string(cef_command_line_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_command_line_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_command_line_get_program(cef_command_line_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_command_line_t *)) { return callback__(self); }
	// void gocef_command_line_set_program(cef_command_line_t * self, cef_string_t * program, void (CEF_CALLBACK *callback__)(cef_command_line_t *, cef_string_t *)) { return callback__(self, program); }
	// int gocef_command_line_has_switches(cef_command_line_t * self, int (CEF_CALLBACK *callback__)(cef_command_line_t *)) { return callback__(self); }
	// int gocef_command_line_has_switch(cef_command_line_t * self, cef_string_t * name, int (CEF_CALLBACK *callback__)(cef_command_line_t *, cef_string_t *)) { return callback__(self, name); }
	// cef_string_userfree_t gocef_command_line_get_switch_value(cef_command_line_t * self, cef_string_t * name, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_command_line_t *, cef_string_t *)) { return callback__(self, name); }
	// void gocef_command_line_get_switches(cef_command_line_t * self, cef_string_map_t switches, void (CEF_CALLBACK *callback__)(cef_command_line_t *, cef_string_map_t)) { return callback__(self, switches); }
	// void gocef_command_line_append_switch(cef_command_line_t * self, cef_string_t * name, void (CEF_CALLBACK *callback__)(cef_command_line_t *, cef_string_t *)) { return callback__(self, name); }
	// void gocef_command_line_append_switch_with_value(cef_command_line_t * self, cef_string_t * name, cef_string_t * value, void (CEF_CALLBACK *callback__)(cef_command_line_t *, cef_string_t *, cef_string_t *)) { return callback__(self, name, value); }
	// int gocef_command_line_has_arguments(cef_command_line_t * self, int (CEF_CALLBACK *callback__)(cef_command_line_t *)) { return callback__(self); }
	// void gocef_command_line_get_arguments(cef_command_line_t * self, cef_string_list_t arguments, void (CEF_CALLBACK *callback__)(cef_command_line_t *, cef_string_list_t)) { return callback__(self, arguments); }
	// void gocef_command_line_append_argument(cef_command_line_t * self, cef_string_t * argument, void (CEF_CALLBACK *callback__)(cef_command_line_t *, cef_string_t *)) { return callback__(self, argument); }
	// void gocef_command_line_prepend_wrapper(cef_command_line_t * self, cef_string_t * wrapper, void (CEF_CALLBACK *callback__)(cef_command_line_t *, cef_string_t *)) { return callback__(self, wrapper); }
	"C"
)

// CommandLine (cef_command_line_t from include/capi/cef_command_line_capi.h)
// Structure used to create and/or parse command line arguments. Arguments with
// '--', '-' and, on Windows, '/' prefixes are considered switches. Switches
// will always precede any arguments without switch prefixes. Switches can
// optionally have a value specified using the '=' delimiter (e.g.
// "-switch=value"). An argument of "--" will terminate switch parsing with all
// subsequent tokens, regardless of prefix, being interpreted as non-switch
// arguments. Switch names are considered case-insensitive. This structure can
// be used before cef_initialize() is called.
type CommandLine C.cef_command_line_t

func (d *CommandLine) toNative() *C.cef_command_line_t {
	return (*C.cef_command_line_t)(d)
}

// Base (base)
// Base structure.
func (d *CommandLine) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// IsValid (is_valid)
// Returns true (1) if this object is valid. Do not call any other functions
// if this function returns false (0).
func (d *CommandLine) IsValid() int32 {
	return int32(C.gocef_command_line_is_valid(d.toNative(), d.is_valid))
}

// IsReadOnly (is_read_only)
// Returns true (1) if the values of this object are read-only. Some APIs may
// expose read-only objects.
func (d *CommandLine) IsReadOnly() int32 {
	return int32(C.gocef_command_line_is_read_only(d.toNative(), d.is_read_only))
}

// Copy (copy)
// Returns a writable copy of this object.
func (d *CommandLine) Copy() *CommandLine {
	return (*CommandLine)(C.gocef_command_line_copy(d.toNative(), d.copy))
}

// InitFromArgv (init_from_argv)
// Initialize the command line with the specified |argc| and |argv| values.
// The first argument must be the name of the program. This function is only
// supported on non-Windows platforms.
func (d *CommandLine) InitFromArgv(argc int32, argv []string) {
	argv_ := C.calloc(C.size_t(len(argv)), C.size_t(unsafe.Sizeof(uintptr(0))))
	argv_p := (*[1<<30 - 1]*C.char)(argv_)
	for i, one := range argv {
		argv_p[i] = C.CString(one)
	}
	C.gocef_command_line_init_from_argv(d.toNative(), C.int(argc), (**C.char)(argv_), d.init_from_argv)
}

// InitFromString (init_from_string)
// Initialize the command line with the string returned by calling
// GetCommandLineW(). This function is only supported on Windows.
func (d *CommandLine) InitFromString(commandLine string) {
	commandLine_ := C.cef_string_userfree_alloc()
	setCEFStr(commandLine, commandLine_)
	defer func() {
		C.cef_string_userfree_free(commandLine_)
	}()
	C.gocef_command_line_init_from_string(d.toNative(), (*C.cef_string_t)(commandLine_), d.init_from_string)
}

// Reset (reset)
// Reset the command-line switches and arguments but leave the program
// component unchanged.
func (d *CommandLine) Reset() {
	C.gocef_command_line_reset(d.toNative(), d.reset)
}

// GetArgv (get_argv)
// Retrieve the original command line string as a vector of strings. The argv
// array: { program, [(--|-|/)switch[=value]]*, [--], [argument]* }
func (d *CommandLine) GetArgv(argv StringList) {
	C.gocef_command_line_get_argv(d.toNative(), C.cef_string_list_t(argv), d.get_argv)
}

// GetCommandLineString (get_command_line_string)
// Constructs and returns the represented command line string. Use this
// function cautiously because quoting behavior is unclear.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *CommandLine) GetCommandLineString() string {
	return cefuserfreestrToString(C.gocef_command_line_get_command_line_string(d.toNative(), d.get_command_line_string))
}

// GetProgram (get_program)
// Get the program part of the command line string (the first item).
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *CommandLine) GetProgram() string {
	return cefuserfreestrToString(C.gocef_command_line_get_program(d.toNative(), d.get_program))
}

// SetProgram (set_program)
// Set the program part of the command line string (the first item).
func (d *CommandLine) SetProgram(program string) {
	program_ := C.cef_string_userfree_alloc()
	setCEFStr(program, program_)
	defer func() {
		C.cef_string_userfree_free(program_)
	}()
	C.gocef_command_line_set_program(d.toNative(), (*C.cef_string_t)(program_), d.set_program)
}

// HasSwitches (has_switches)
// Returns true (1) if the command line has switches.
func (d *CommandLine) HasSwitches() int32 {
	return int32(C.gocef_command_line_has_switches(d.toNative(), d.has_switches))
}

// HasSwitch (has_switch)
// Returns true (1) if the command line contains the given switch.
func (d *CommandLine) HasSwitch(name string) int32 {
	name_ := C.cef_string_userfree_alloc()
	setCEFStr(name, name_)
	defer func() {
		C.cef_string_userfree_free(name_)
	}()
	return int32(C.gocef_command_line_has_switch(d.toNative(), (*C.cef_string_t)(name_), d.has_switch))
}

// GetSwitchValue (get_switch_value)
// Returns the value associated with the given switch. If the switch has no
// value or isn't present this function returns the NULL string.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *CommandLine) GetSwitchValue(name string) string {
	name_ := C.cef_string_userfree_alloc()
	setCEFStr(name, name_)
	defer func() {
		C.cef_string_userfree_free(name_)
	}()
	return cefuserfreestrToString(C.gocef_command_line_get_switch_value(d.toNative(), (*C.cef_string_t)(name_), d.get_switch_value))
}

// GetSwitches (get_switches)
// Returns the map of switch names and values. If a switch has no value an
// NULL string is returned.
func (d *CommandLine) GetSwitches(switches StringMap) {
	C.gocef_command_line_get_switches(d.toNative(), C.cef_string_map_t(switches), d.get_switches)
}

// AppendSwitch (append_switch)
// Add a switch to the end of the command line. If the switch has no value
// pass an NULL value string.
func (d *CommandLine) AppendSwitch(name string) {
	name_ := C.cef_string_userfree_alloc()
	setCEFStr(name, name_)
	defer func() {
		C.cef_string_userfree_free(name_)
	}()
	C.gocef_command_line_append_switch(d.toNative(), (*C.cef_string_t)(name_), d.append_switch)
}

// AppendSwitchWithValue (append_switch_with_value)
// Add a switch with the specified value to the end of the command line.
func (d *CommandLine) AppendSwitchWithValue(name, value string) {
	name_ := C.cef_string_userfree_alloc()
	setCEFStr(name, name_)
	defer func() {
		C.cef_string_userfree_free(name_)
	}()
	value_ := C.cef_string_userfree_alloc()
	setCEFStr(value, value_)
	defer func() {
		C.cef_string_userfree_free(value_)
	}()
	C.gocef_command_line_append_switch_with_value(d.toNative(), (*C.cef_string_t)(name_), (*C.cef_string_t)(value_), d.append_switch_with_value)
}

// HasArguments (has_arguments)
// True if there are remaining command line arguments.
func (d *CommandLine) HasArguments() int32 {
	return int32(C.gocef_command_line_has_arguments(d.toNative(), d.has_arguments))
}

// GetArguments (get_arguments)
// Get the remaining command line arguments.
func (d *CommandLine) GetArguments(arguments StringList) {
	C.gocef_command_line_get_arguments(d.toNative(), C.cef_string_list_t(arguments), d.get_arguments)
}

// AppendArgument (append_argument)
// Add an argument to the end of the command line.
func (d *CommandLine) AppendArgument(argument string) {
	argument_ := C.cef_string_userfree_alloc()
	setCEFStr(argument, argument_)
	defer func() {
		C.cef_string_userfree_free(argument_)
	}()
	C.gocef_command_line_append_argument(d.toNative(), (*C.cef_string_t)(argument_), d.append_argument)
}

// PrependWrapper (prepend_wrapper)
// Insert a command before the current command. Common for debuggers, like
// "valgrind" or "gdb --args".
func (d *CommandLine) PrependWrapper(wrapper string) {
	wrapper_ := C.cef_string_userfree_alloc()
	setCEFStr(wrapper, wrapper_)
	defer func() {
		C.cef_string_userfree_free(wrapper_)
	}()
	C.gocef_command_line_prepend_wrapper(d.toNative(), (*C.cef_string_t)(wrapper_), d.prepend_wrapper)
}
