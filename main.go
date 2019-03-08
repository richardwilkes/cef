package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/richardwilkes/cef/internal/cmd"
	"github.com/richardwilkes/toolbox"
	"github.com/richardwilkes/toolbox/atexit"
	"github.com/richardwilkes/toolbox/cmdline"
)

func main() {
	// This is normally the same for all platforms, however, the March 8, 2019
	// build has a slightly older build for Windows. Given the security issue
	// this build addresses, I felt it was better to have the platforms
	// deviate on the specific build than to wait for them to return to being
	// in sync with each other.
	var desiredCEFVersion string
	if runtime.GOOS == toolbox.WindowsOS {
		desiredCEFVersion = "3.3578.1870.gc974488"
	} else {
		desiredCEFVersion = "3.3626.1895.g7001d56"
	}

	cmdline.CopyrightYears = "2018-2019"
	cmdline.CopyrightHolder = "Richard A. Wilkes"
	cmdline.AppIdentifier = "com.trollworks.cef"
	cl := cmdline.New(true)
	cl.Description = "Utilities for managing setup of the Chromium Embedded Framework."
	cl.AddCommand(cmd.NewInstall(desiredCEFVersion))
	cl.AddCommand(cmd.NewDist())
	if err := cl.RunCommand(cl.Parse(os.Args[1:])); err != nil {
		fmt.Fprintln(os.Stderr, err)
		atexit.Exit(1)
	}
}
