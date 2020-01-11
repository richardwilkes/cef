// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package main

import (
	"fmt"
	"os"

	"github.com/richardwilkes/cef/internal/cmd"
	"github.com/richardwilkes/toolbox/atexit"
	"github.com/richardwilkes/toolbox/cmdline"
)

const desiredCEFVersion = "75.1.4+g4210896+chromium-75.0.3770.100"

func main() {
	cmdline.CopyrightYears = "2018-2020"
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
