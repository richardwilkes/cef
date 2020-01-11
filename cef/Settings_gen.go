// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Code created from "struct.go.tmpl" - don't edit by hand

package cef

import (
	"path/filepath"

	"github.com/richardwilkes/toolbox/xio/fs/paths"
)

import (
	// #include "capi_gen.h"
	"C"
)

// Settings (cef_settings_t from include/internal/cef_types.h)
// Initialization settings. Specify NULL or 0 to get the recommended default
// values. Many of these and other settings can also configured using command-
// line switches.
type Settings struct {
	// Size (size)
	// Size of this structure.
	Size uint64
	// NoSandbox (no_sandbox)
	// Set to true (1) to disable the sandbox for sub-processes. See
	// cef_sandbox_win.h for requirements to enable the sandbox on Windows. Also
	// configurable using the "no-sandbox" command-line switch.
	NoSandbox int32
	// BrowserSubprocessPath (browser_subprocess_path)
	// The path to a separate executable that will be launched for sub-processes.
	// If this value is empty on Windows or Linux then the main process executable
	// will be used. If this value is empty on macOS then a helper executable must
	// exist at "Contents/Frameworks/<app> Helper.app/Contents/MacOS/<app> Helper"
	// in the top-level app bundle. See the comments on CefExecuteProcess() for
	// details. Also configurable using the "browser-subprocess-path" command-line
	// switch.
	BrowserSubprocessPath string
	// FrameworkDirPath (framework_dir_path)
	// The path to the CEF framework directory on macOS. If this value is empty
	// then the framework must exist at "Contents/Frameworks/Chromium Embedded
	// Framework.framework" in the top-level app bundle. Also configurable using
	// the "framework-dir-path" command-line switch.
	FrameworkDirPath string
	// MultiThreadedMessageLoop (multi_threaded_message_loop)
	// Set to true (1) to have the browser process message loop run in a separate
	// thread. If false (0) than the CefDoMessageLoopWork() function must be
	// called from your application message loop. This option is only supported on
	// Windows and Linux.
	MultiThreadedMessageLoop int32
	// ExternalMessagePump (external_message_pump)
	// Set to true (1) to control browser process main (UI) thread message pump
	// scheduling via the CefBrowserProcessHandler::OnScheduleMessagePumpWork()
	// callback. This option is recommended for use in combination with the
	// CefDoMessageLoopWork() function in cases where the CEF message loop must be
	// integrated into an existing application message loop (see additional
	// comments and warnings on CefDoMessageLoopWork). Enabling this option is not
	// recommended for most users; leave this option disabled and use either the
	// CefRunMessageLoop() function or multi_threaded_message_loop if possible.
	ExternalMessagePump int32
	// WindowlessRenderingEnabled (windowless_rendering_enabled)
	// Set to true (1) to enable windowless (off-screen) rendering support. Do not
	// enable this value if the application does not use windowless rendering as
	// it may reduce rendering performance on some systems.
	WindowlessRenderingEnabled int32
	// CommandLineArgsDisabled (command_line_args_disabled)
	// Set to true (1) to disable configuration of browser process features using
	// standard CEF and Chromium command-line arguments. Configuration can still
	// be specified using CEF data structures or via the
	// CefApp::OnBeforeCommandLineProcessing() method.
	CommandLineArgsDisabled int32
	// CachePath (cache_path)
	// The location where data for the global browser cache will be stored on
	// disk. In non-empty this must be either equal to or a child directory of
	// CefSettings.root_cache_path. If empty then browsers will be created in
	// "incognito mode" where in-memory caches are used for storage and no data is
	// persisted to disk. HTML5 databases such as localStorage will only persist
	// across sessions if a cache path is specified. Can be overridden for
	// individual CefRequestContext instances via the
	// CefRequestContextSettings.cache_path value.
	CachePath string
	// RootCachePath (root_cache_path)
	// The root directory that all CefSettings.cache_path and
	// CefRequestContextSettings.cache_path values must have in common. If this
	// value is empty and CefSettings.cache_path is non-empty then this value will
	// default to the CefSettings.cache_path value. Failure to set this value
	// correctly may result in the sandbox blocking read/write access to the
	// cache_path directory.
	RootCachePath string
	// UserDataPath (user_data_path)
	// The location where user data such as spell checking dictionary files will
	// be stored on disk. If empty then the default platform-specific user data
	// directory will be used ("~/.cef_user_data" directory on Linux,
	// "~/Library/Application Support/CEF/User Data" directory on Mac OS X,
	// "Local Settings\Application Data\CEF\User Data" directory under the user
	// profile directory on Windows).
	UserDataPath string
	// PersistSessionCookies (persist_session_cookies)
	// To persist session cookies (cookies without an expiry date or validity
	// interval) by default when using the global cookie manager set this value to
	// true (1). Session cookies are generally intended to be transient and most
	// Web browsers do not persist them. A |cache_path| value must also be
	// specified to enable this feature. Also configurable using the
	// "persist-session-cookies" command-line switch. Can be overridden for
	// individual CefRequestContext instances via the
	// CefRequestContextSettings.persist_session_cookies value.
	PersistSessionCookies int32
	// PersistUserPreferences (persist_user_preferences)
	// To persist user preferences as a JSON file in the cache path directory set
	// this value to true (1). A |cache_path| value must also be specified
	// to enable this feature. Also configurable using the
	// "persist-user-preferences" command-line switch. Can be overridden for
	// individual CefRequestContext instances via the
	// CefRequestContextSettings.persist_user_preferences value.
	PersistUserPreferences int32
	// UserAgent (user_agent)
	// Value that will be returned as the User-Agent HTTP header. If empty the
	// default User-Agent string will be used. Also configurable using the
	// "user-agent" command-line switch.
	UserAgent string
	// ProductVersion (product_version)
	// Value that will be inserted as the product portion of the default
	// User-Agent string. If empty the Chromium product version will be used. If
	// |userAgent| is specified this value will be ignored. Also configurable
	// using the "product-version" command-line switch.
	ProductVersion string
	// Locale (locale)
	// The locale string that will be passed to WebKit. If empty the default
	// locale of "en-US" will be used. This value is ignored on Linux where locale
	// is determined using environment variable parsing with the precedence order:
	// LANGUAGE, LC_ALL, LC_MESSAGES and LANG. Also configurable using the "lang"
	// command-line switch.
	Locale string
	// LogFile (log_file)
	// The directory and file name to use for the debug log. If empty a default
	// log file name and location will be used. On Windows and Linux a "debug.log"
	// file will be written in the main executable directory. On Mac OS X a
	// "~/Library/Logs/<app name>_debug.log" file will be written where <app name>
	// is the name of the main app executable. Also configurable using the
	// "log-file" command-line switch.
	LogFile string
	// LogSeverity (log_severity)
	// The log severity. Only messages of this severity level or higher will be
	// logged. When set to DISABLE no messages will be written to the log file,
	// but FATAL messages will still be output to stderr. Also configurable using
	// the "log-severity" command-line switch with a value of "verbose", "info",
	// "warning", "error", "fatal" or "disable".
	LogSeverity LogSeverity
	// JavascriptFlags (javascript_flags)
	// Custom flags that will be used when initializing the V8 JavaScript engine.
	// The consequences of using custom flags may not be well tested. Also
	// configurable using the "js-flags" command-line switch.
	JavascriptFlags string
	// ResourcesDirPath (resources_dir_path)
	// The fully qualified path for the resources directory. If this value is
	// empty the cef.pak and/or devtools_resources.pak files must be located in
	// the module directory on Windows/Linux or the app bundle Resources directory
	// on Mac OS X. Also configurable using the "resources-dir-path" command-line
	// switch.
	ResourcesDirPath string
	// LocalesDirPath (locales_dir_path)
	// The fully qualified path for the locales directory. If this value is empty
	// the locales directory must be located in the module directory. This value
	// is ignored on Mac OS X where pack files are always loaded from the app
	// bundle Resources directory. Also configurable using the "locales-dir-path"
	// command-line switch.
	LocalesDirPath string
	// PackLoadingDisabled (pack_loading_disabled)
	// Set to true (1) to disable loading of pack files for resources and locales.
	// A resource bundle handler must be provided for the browser and render
	// processes via CefApp::GetResourceBundleHandler() if loading of pack files
	// is disabled. Also configurable using the "disable-pack-loading" command-
	// line switch.
	PackLoadingDisabled int32
	// RemoteDebuggingPort (remote_debugging_port)
	// Set to a value between 1024 and 65535 to enable remote debugging on the
	// specified port. For example, if 8080 is specified the remote debugging URL
	// will be http://localhost:8080. CEF can be remotely debugged from any CEF or
	// Chrome browser window. Also configurable using the "remote-debugging-port"
	// command-line switch.
	RemoteDebuggingPort int32
	// UncaughtExceptionStackSize (uncaught_exception_stack_size)
	// The number of stack trace frames to capture for uncaught exceptions.
	// Specify a positive value to enable the CefRenderProcessHandler::
	// OnUncaughtException() callback. Specify 0 (default value) and
	// OnUncaughtException() will not be called. Also configurable using the
	// "uncaught-exception-stack-size" command-line switch.
	UncaughtExceptionStackSize int32
	// IgnoreCertificateErrors (ignore_certificate_errors)
	// Set to true (1) to ignore errors related to invalid SSL certificates.
	// Enabling this setting can lead to potential security vulnerabilities like
	// "man in the middle" attacks. Applications that load content from the
	// internet should not enable this setting. Also configurable using the
	// "ignore-certificate-errors" command-line switch. Can be overridden for
	// individual CefRequestContext instances via the
	// CefRequestContextSettings.ignore_certificate_errors value.
	IgnoreCertificateErrors int32
	// EnableNetSecurityExpiration (enable_net_security_expiration)
	// Set to true (1) to enable date-based expiration of built in network
	// security information (i.e. certificate transparency logs, HSTS preloading
	// and pinning information). Enabling this option improves network security
	// but may cause HTTPS load failures when using CEF binaries built more than
	// 10 weeks in the past. See https://www.certificate-transparency.org/ and
	// https://www.chromium.org/hsts for details. Also configurable using the
	// "enable-net-security-expiration" command-line switch. Can be overridden for
	// individual CefRequestContext instances via the
	// CefRequestContextSettings.enable_net_security_expiration value.
	EnableNetSecurityExpiration int32
	// BackgroundColor (background_color)
	// Background color used for the browser before a document is loaded and when
	// no document color is specified. The alpha component must be either fully
	// opaque (0xFF) or fully transparent (0x00). If the alpha component is fully
	// opaque then the RGB components will be used as the background color. If the
	// alpha component is fully transparent for a windowed browser then the
	// default value of opaque white be used. If the alpha component is fully
	// transparent for a windowless (off-screen) browser then transparent painting
	// will be enabled.
	BackgroundColor Color
	// AcceptLanguageList (accept_language_list)
	// Comma delimited ordered list of language codes without any whitespace that
	// will be used in the "Accept-Language" HTTP header. May be overridden on a
	// per-browser basis using the CefBrowserSettings.accept_language_list value.
	// If both values are empty then "en-US,en" will be used. Can be overridden
	// for individual CefRequestContext instances via the
	// CefRequestContextSettings.accept_language_list value.
	AcceptLanguageList string
	// ApplicationClientIDForFileScanning (application_client_id_for_file_scanning)
	// GUID string used for identifying the application. This is passed to the
	// system AV function for scanning downloaded files. By default, the GUID
	// will be an empty string and the file will be treated as an untrusted
	// file when the GUID is empty.
	ApplicationClientIDForFileScanning string
}

// NewSettings creates a new Settings.
func NewSettings() *Settings {
	return &Settings{
		Size:                    C.sizeof_struct__cef_settings_t,
		NoSandbox:               1,
		CommandLineArgsDisabled: 1,
		LogSeverity:             LogseverityWarning,
		CachePath:               filepath.Join(paths.AppDataDir(), "cache"),
		UserDataPath:            filepath.Join(paths.AppDataDir(), "data"),
		LogFile:                 filepath.Join(paths.AppLogDir(), "cef.log"),
	}
}

func (d *Settings) toNative(native *C.cef_settings_t) *C.cef_settings_t {
	if d == nil {
		return nil
	}
	native.size = C.size_t(d.Size)
	native.no_sandbox = C.int(d.NoSandbox)
	setCEFStr(d.BrowserSubprocessPath, &native.browser_subprocess_path)
	setCEFStr(d.FrameworkDirPath, &native.framework_dir_path)
	native.multi_threaded_message_loop = C.int(d.MultiThreadedMessageLoop)
	native.external_message_pump = C.int(d.ExternalMessagePump)
	native.windowless_rendering_enabled = C.int(d.WindowlessRenderingEnabled)
	native.command_line_args_disabled = C.int(d.CommandLineArgsDisabled)
	setCEFStr(d.CachePath, &native.cache_path)
	setCEFStr(d.RootCachePath, &native.root_cache_path)
	setCEFStr(d.UserDataPath, &native.user_data_path)
	native.persist_session_cookies = C.int(d.PersistSessionCookies)
	native.persist_user_preferences = C.int(d.PersistUserPreferences)
	setCEFStr(d.UserAgent, &native.user_agent)
	setCEFStr(d.ProductVersion, &native.product_version)
	setCEFStr(d.Locale, &native.locale)
	setCEFStr(d.LogFile, &native.log_file)
	native.log_severity = C.cef_log_severity_t(d.LogSeverity)
	setCEFStr(d.JavascriptFlags, &native.javascript_flags)
	setCEFStr(d.ResourcesDirPath, &native.resources_dir_path)
	setCEFStr(d.LocalesDirPath, &native.locales_dir_path)
	native.pack_loading_disabled = C.int(d.PackLoadingDisabled)
	native.remote_debugging_port = C.int(d.RemoteDebuggingPort)
	native.uncaught_exception_stack_size = C.int(d.UncaughtExceptionStackSize)
	native.ignore_certificate_errors = C.int(d.IgnoreCertificateErrors)
	native.enable_net_security_expiration = C.int(d.EnableNetSecurityExpiration)
	native.background_color = C.cef_color_t(d.BackgroundColor)
	setCEFStr(d.AcceptLanguageList, &native.accept_language_list)
	setCEFStr(d.ApplicationClientIDForFileScanning, &native.application_client_id_for_file_scanning)
	return native
}

func (n *C.cef_settings_t) toGo() *Settings {
	if n == nil {
		return nil
	}
	var d Settings
	n.intoGo(&d)
	return &d
}

func (n *C.cef_settings_t) intoGo(d *Settings) {
	d.Size = uint64(n.size)
	d.NoSandbox = int32(n.no_sandbox)
	d.BrowserSubprocessPath = cefstrToString(&n.browser_subprocess_path)
	d.FrameworkDirPath = cefstrToString(&n.framework_dir_path)
	d.MultiThreadedMessageLoop = int32(n.multi_threaded_message_loop)
	d.ExternalMessagePump = int32(n.external_message_pump)
	d.WindowlessRenderingEnabled = int32(n.windowless_rendering_enabled)
	d.CommandLineArgsDisabled = int32(n.command_line_args_disabled)
	d.CachePath = cefstrToString(&n.cache_path)
	d.RootCachePath = cefstrToString(&n.root_cache_path)
	d.UserDataPath = cefstrToString(&n.user_data_path)
	d.PersistSessionCookies = int32(n.persist_session_cookies)
	d.PersistUserPreferences = int32(n.persist_user_preferences)
	d.UserAgent = cefstrToString(&n.user_agent)
	d.ProductVersion = cefstrToString(&n.product_version)
	d.Locale = cefstrToString(&n.locale)
	d.LogFile = cefstrToString(&n.log_file)
	d.LogSeverity = LogSeverity(n.log_severity)
	d.JavascriptFlags = cefstrToString(&n.javascript_flags)
	d.ResourcesDirPath = cefstrToString(&n.resources_dir_path)
	d.LocalesDirPath = cefstrToString(&n.locales_dir_path)
	d.PackLoadingDisabled = int32(n.pack_loading_disabled)
	d.RemoteDebuggingPort = int32(n.remote_debugging_port)
	d.UncaughtExceptionStackSize = int32(n.uncaught_exception_stack_size)
	d.IgnoreCertificateErrors = int32(n.ignore_certificate_errors)
	d.EnableNetSecurityExpiration = int32(n.enable_net_security_expiration)
	d.BackgroundColor = Color(n.background_color)
	d.AcceptLanguageList = cefstrToString(&n.accept_language_list)
	d.ApplicationClientIDForFileScanning = cefstrToString(&n.application_client_id_for_file_scanning)
}
