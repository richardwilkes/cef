// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package cmd

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/richardwilkes/toolbox"
	"github.com/richardwilkes/toolbox/atexit"
)

var (
	installPrefix = "/usr/local/cef"
	cefPlatform   string
)

func checkPlatform() {
	switch runtime.GOOS {
	case toolbox.MacOS:
		cefPlatform = "macosx64"
	case toolbox.LinuxOS:
		cefPlatform = "linux64"
	case toolbox.WindowsOS:
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
