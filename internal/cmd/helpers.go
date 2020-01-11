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

	"github.com/richardwilkes/toolbox/atexit"
)

func createDir(dir string, mode os.FileMode) {
	if err := os.MkdirAll(dir, mode); err != nil {
		fmt.Println(err)
		fmt.Println("You may need to run the 'cef' tool as root.")
		atexit.Exit(1)
	}
}

func checkFileError(err error, op, name string) {
	if err != nil {
		fmt.Printf("Unable to %s file %s\n", op, name)
		fmt.Println(err)
		atexit.Exit(1)
	}
}
