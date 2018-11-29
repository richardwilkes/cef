package cef

import (
	// #include "application_darwin.h"
	"C"
)

func instantiateApplication() {
	C.gocef_instantiate_application()
}
