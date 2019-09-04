package main

import (
	"fmt"
	"os"

	"github.com/richardwilkes/cef/internal/cmd"
	"github.com/richardwilkes/toolbox/atexit"
	"github.com/richardwilkes/toolbox/cmdline"
)

const desiredCEFVersion = "76.1.13+gf19c584+chromium-76.0.3809.132"

func main() {
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
