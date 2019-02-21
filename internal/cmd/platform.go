package cmd

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/richardwilkes/toolbox/atexit"
)

// Constants for comparison to runtime.GOOS
const (
	MacOS     = "darwin"
	WindowsOS = "windows"
	LinuxOS   = "linux"
)

var (
	installPrefix = "/usr/local/cef"
	cefPlatform   string
)

func checkPlatform() {
	switch runtime.GOOS {
	case MacOS:
		cefPlatform = "macosx64"
	case LinuxOS:
		cefPlatform = "linux64"
	case WindowsOS:
		if os.Getenv("MSYSTEM") != "MINGW64" {
			fmt.Println("Windows is only supported through the use of MINGW64")
			atexit.Exit(1)
		}
		cefPlatform = "windows64"
		installPrefix = path.Join(path.Dir(os.Getenv("MINGW_PREFIX")), installPrefix)
	default:
		fmt.Println("Unsupported OS: ", runtime.GOOS)
		atexit.Exit(1)
	}
}
