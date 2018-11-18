package cef

import (
	// #include "app_darwin.h"
	"C"
)

func instantiateCEFApplication() {
	C.gocef_instantiate_cef_application()
}
