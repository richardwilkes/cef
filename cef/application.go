package cef

// InstantiateApplication is only needed on macOS. It will instantiate a
// special version of NSApplication for CEF.
func InstantiateApplication() {
	instantiateApplication()
}
