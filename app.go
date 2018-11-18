package cef

import (
	// // Note: Only one file per package needs the #cgo directives.
	// //       Includes are still needed on a per-file basis.
	//
	// #cgo CFLAGS: -I ${SRCDIR}/cef
	// #cgo darwin LDFLAGS: -framework Cocoa -F ${SRCDIR}/cef/Release -framework "Chromium Embedded Framework"
	// #cgo windows LDFLAGS: -L${SRCDIR}/cef/Release -lcef -Wl,--subsystem,windows
	// #include "include/capi/cef_app_capi.h"
	"C"

	"github.com/richardwilkes/toolbox/atexit"
	"github.com/richardwilkes/toolbox/errs"
)

// ExecuteProcess is used to start the secondary CEF processes. If this is
// the main process, this call will do nothing and return. If it is a
// secondary process, the call will not return.
func ExecuteProcess() error {
	args, err := mainArgs()
	if err != nil {
		return err
	}
	if code := C.cef_execute_process(args, nil, nil); code >= 0 {
		atexit.Exit(int(code))
	}
	instantiateCEFApplication()
	return nil
}

// Initialize CEF.
func Initialize(settings Settings) error {
	args, err := mainArgs()
	if err != nil {
		return err
	}
	if C.cef_initialize(args, settings, nil, nil) != 1 {
		return errs.New("Unable to initialize CEF")
	}
	return nil
}

// RunMessageLoop runs the application event loop.
func RunMessageLoop() {
	C.cef_run_message_loop()
}

// QuitMessageLoop quits the message loop in preparation for exiting.
func QuitMessageLoop() {
	C.cef_quit_message_loop()
}

// Shutdown CEF and the application.
func Shutdown() {
	C.cef_shutdown()
}

// EnableHighResolutionSupport enables CEF's high-resolution support. This
// should be called before any other CEF function if this support is desired.
func EnableHighResolutionSupport() {
	C.cef_enable_highdpi_support()
}
