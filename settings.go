package cef

import (
	// #include <stdlib.h>
	// #include "include/internal/cef_types.h"
	"C"
)

// LogSeverity holds a log severity level.
type LogSeverity int

// Possible LogSeverity levels
const (
	LogSeverityDefault = iota
	LogSeverityVerbose
	LogSeverityInfo
	LogSeverityWarning
	LogSeverityError
	LogSeverityDebug   = LogSeverityVerbose
	LogSeverityDisable = 99
)

// Settings holds initialization settings.
//
// Defined in include/internal/cef_types.h
type Settings struct {
	native *C.cef_settings_t
}

// NewSettings creates a new default Settings instance.
func NewSettings() *Settings {
	settings := &Settings{native: (*C.cef_settings_t)(C.calloc(1, C.sizeof_struct__cef_settings_t))}
	settings.native.size = C.sizeof_struct__cef_settings_t
	settings.native.no_sandbox = 1
	settings.native.command_line_args_disabled = 1
	settings.SetLogSeverity(LogSeverityError)
	return settings
}

// SetBrowserSubProcessPath sets the path to a separate executable that will
// be launched for sub-processes. If this value is empty on Windows or Linux
// then the main process executable will be used. If this value is empty on
// macOS then a helper executable must exist at
// "Contents/Frameworks/<app> Helper.app/Contents/MacOS/<app> Helper" in the
// top-level app bundle. See the comments on ExecuteProcess for details.
func (s *Settings) SetBrowserSubProcessPath(path string) {
	setCEFStr(path, &s.native.browser_subprocess_path)
}

// SetFrameworkDirPath sets the path to the CEF framework directory on macOS.
// Defaults to "Contents/Frameworks/Chromium Embedded Framework.framework" in
// the top-level app bundle.
func (s *Settings) SetFrameworkDirPath(path string) {
	setCEFStr(path, &s.native.framework_dir_path)
}

// EnableMultiThreadedMessageLoop causes the browser process message loop to
// run in a separate thread. This option is only supported on Windows and
// Linux.
func (s *Settings) EnableMultiThreadedMessageLoop() {
	s.native.multi_threaded_message_loop = 1
}

// EnableExternalMessagePump to call DoMessageLoopWork within your own message
// loop. Enabling this option is not recommended for most users.
func (s *Settings) EnableExternalMessagePump() {
	s.native.external_message_pump = 1
}

// EnableWindowlessRendering enables windowless (off-screen) rendering
// support. Do not enable this value if the application does not use
// windowless rendering as it may reduce rendering performance on some
// systems.
func (s *Settings) EnableWindowlessRendering() {
	s.native.windowless_rendering_enabled = 1
}

// SetCachePath sets the location where cache data will be stored on disk. If
// empty then browsers will be created in "incognito mode" where in-memory
// caches are used for storage and no data is persisted to disk. HTML5
// databases such as localStorage will only persist across sessions if a
// cache path is specified. Can be overridden for individual RequestContext
// instances.
func (s *Settings) SetCachePath(path string) {
	setCEFStr(path, &s.native.cache_path)
}

// SetUserDataPath sets the location where user data such as spell checking
// dictionary files will be stored on disk. If empty then the default
// platform-specific user data directory will be used ("~/.cef_user_data"
// directory on Linux, "~/Library/Application Support/CEF/User Data" directory
// on macOS, "Local Settings\Application Data\CEF\User Data" directory under
// the user profile directory on Windows).
func (s *Settings) SetUserDataPath(path string) {
	setCEFStr(path, &s.native.user_data_path)
}

// EnablePersistentSessionCookies to persist session cookies (cookies without
// an expiry date or validity interval) by default when using the global
// cookie manager. Session cookies are generally intended to be transient and
// most Web browsers do not persist them. The cache path (via SetCachePath)
// must also be specified to enable this feature. Can be overridden for
// individual RequestContext instances.
func (s *Settings) EnablePersistentSessionCookies() {
	s.native.persist_session_cookies = 1
}

// EnablePersistentUserPreferences to persist user preferences as a JSON file
// in the cache path directory. The cache path (via SetCachePath) must also be
// specified to enable this feature. Can be overridden for individual
// RequestContext instances.
func (s *Settings) EnablePersistentUserPreferences() {
	s.native.persist_user_preferences = 1
}

// SetUserAgent sets the value that will be returned as the User-Agent HTTP
// header.
func (s *Settings) SetUserAgent(userAgent string) {
	setCEFStr(userAgent, &s.native.user_agent)
}

// SetProductVersion sets the value that will be inserted as the product
// portion of the default User-Agent string. Defaults to the Chromium product
// version. If a user agent has been set via a call to SetUserAgent, this
// value will be ignored.
func (s *Settings) SetProductVersion(version string) {
	setCEFStr(version, &s.native.product_version)
}

// SetLocale sets the locale string that will be passed to WebKit. Defaults to
// "en-US". This value is ignored on Linux where locale is determined using
// environment variable parsing with the precedence order: LANGUAGE, LC_ALL,
// LC_MESSAGES and LANG.
func (s *Settings) SetLocale(locale string) {
	setCEFStr(locale, &s.native.locale)
}

// SetLogFile sets the directory and file name to use for the debug log. By
// default, on Windows and Linux a "debug.log" file will be written in the
// main executable directory, while on macOS a
// "~/Library/Logs/<app name>_debug.log" file will be written where <app name>
// is the name of the main app executable.
func (s *Settings) SetLogFile(locale string) {
	setCEFStr(locale, &s.native.log_file)
}

// SetLogSeverity sets the log severity. Only messages of this severity level
// or higher will be logged. Defaults to LogSeverityError.
func (s *Settings) SetLogSeverity(level LogSeverity) {
	s.native.log_severity = C.cef_log_severity_t(level)
}

// SetJavaScriptFLags sets custom flags that will be used when initializing
// the V8 JavaScript engine. The consequences of using custom flags may not be
// well tested.
func (s *Settings) SetJavaScriptFLags(flags string) {
	setCEFStr(flags, &s.native.javascript_flags)
}

// SetResourcesDirPath sets the fully qualified path for the resources
// directory. If this value is not set, the cef.pak and/or
// devtools_resources.pak files must be located in the module directory on
// Windows/Linux or the app bundle Resources directory on macOS.
func (s *Settings) SetResourcesDirPath(path string) {
	setCEFStr(path, &s.native.resources_dir_path)
}

// SetLocalesDirPath sets the fully qualified path for the locales directory.
// If this value is not set, the locales directory must be located in the
// module directory. This value is ignored on macOS where pack files are
// always loaded from the app bundle Resources directory.
func (s *Settings) SetLocalesDirPath(path string) {
	setCEFStr(path, &s.native.locales_dir_path)
}

// DisablePackLoading to disable loading of pack files for resources and
// locales. A resource bundle handler must be provided for the browser and
// render processes if loading of pack files is disabled.
func (s *Settings) DisablePackLoading() {
	s.native.pack_loading_disabled = 1
}

// SetRemoteDebuggingPort enables remote debugging on the specified port,
// which must be in the range 1024 to 65535. For example, if 8080 is
// specified, the remote debugging URL will be http://localhost:8080. CEF can
// be remotely debugged from any CEF or Chrome browser window.
func (s *Settings) SetRemoteDebuggingPort(port int) {
	if port >= 1024 && port <= 65535 {
		s.native.remote_debugging_port = C.int(port)
	}
}

// SetUncaughtExceptionStackSize sets the number of stack trace frames to
// capture for uncaught exceptions. A zero or negative value turns off stack
// traces. Defaults to 0.
func (s *Settings) SetUncaughtExceptionStackSize(size int) {
	s.native.uncaught_exception_stack_size = C.int(size)
}

// IgnoreCertificateErrors causes errors related to invalid SSL certificates
// to be ignored. Enabling this setting can lead to potential security
// vulnerabilities like "man in the middle" attacks. Applications that load
// content from the internet should not enable this setting. Can be overridden
// for individual RequestContext instances.
func (s *Settings) IgnoreCertificateErrors() {
	s.native.ignore_certificate_errors = 1
}

// EnableNetSecurityExpiration enables date-based expiration of built in
// network security information (i.e. certificate transparency logs, HSTS
// preloading and pinning information). Enabling this option improves network
// security but may cause HTTPS load failures when using CEF binaries built
// more than 10 weeks in the past. See
// https://www.certificate-transparency.org/ and https://www.chromium.org/hsts
// for details. Can be overridden for individual RequestContext instances.
func (s *Settings) EnableNetSecurityExpiration() {
	s.native.enable_net_security_expiration = 1
}

// SetBackgroundColor sets the background color used for the browser before a
// document is loaded and when no document color is specified. The alpha
// component must be either fully opaque (0xFF) or fully transparent (0x00).
// If the alpha component is fully opaque then the RGB components will be used
// as the background color. If the alpha component is fully transparent for a
// windowed browser then the background color value will be used. If the alpha
// component is fully transparent for a windowless (off-screen) browser then
// transparent painting will be enabled. May be overridden by BrowserSettings.
func (s *Settings) SetBackgroundColor(color Color) {
	s.native.background_color = C.cef_color_t(color)
}

// SetAcceptLanguageList sets a comma delimited ordered list of language codes
// without any whitespace that will be used in the "Accept-Language" HTTP
// header. May be overridden by BrowserSettings. Defaults to "en-US,en".
func (s *Settings) SetAcceptLanguageList(list string) {
	setCEFStr(list, &s.native.accept_language_list)
}
