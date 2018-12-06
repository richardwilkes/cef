package cmd

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/richardwilkes/toolbox/atexit"
)

var (
	installPrefix = "/usr/local/cef"
	cefPlatform   string
)

func checkPlatform() {
	switch runtime.GOOS {
	case "darwin":
		cefPlatform = "macosx64"
	case "linux":
		cefPlatform = "linux64"
	case "windows":
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
