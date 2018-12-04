package main

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/bzip2"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/richardwilkes/toolbox/atexit"
	"github.com/richardwilkes/toolbox/cmdline"
	"github.com/richardwilkes/toolbox/xio"
)

const desiredCEFVersion = "3.3538.1852.gcb937fc"

var (
	cefVersionRegex = regexp.MustCompile(`^\s*#define\s+CEF_VERSION\s+"(\d+\.\d+\.\d+\.\w+)"\s*$`)
	installPrefix   = "/usr/local/cef"
	cefPlatform     string
)

func main() {
	cmdline.CopyrightYears = "2018"
	cmdline.CopyrightHolder = "Richard A. Wilkes"
	cmdline.AppIdentifier = "com.trollworks.cef"
	cl := cmdline.New(true)
	cl.Description = "Downloads and installs the headers and libraries necessary use the github.com/richardwilkes/cef/cef package."
	var needInstall bool
	cl.NewBoolOption(&needInstall).SetSingle('f').SetName("force").SetUsage("Force an install")
	cl.Parse(os.Args[1:])

	switch runtime.GOOS {
	case "darwin":
		cefPlatform = "macosx64"
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

	if !needInstall {
		var existingCEFVersion string
		if f, err := os.Open(path.Join(installPrefix, "include/cef_version.h")); err == nil {
			s := bufio.NewScanner(f)
			for s.Scan() {
				line := s.Text()
				if result := cefVersionRegex.FindStringSubmatch(line); len(result) == 2 {
					existingCEFVersion = result[1]
					break
				}
			}
			xio.CloseIgnoringErrors(f)
		}
		needInstall = existingCEFVersion != desiredCEFVersion
	}

	if needInstall {
		fmt.Printf("Installing into %s...\n", installPrefix)
		if err := os.RemoveAll(installPrefix); err != nil {
			fmt.Printf("Unable to remove old installation at %s\n", installPrefix)
			fmt.Println(err)
			fmt.Println("You may need to run this as root.")
			atexit.Exit(1)
		}
		createDir(installPrefix, 0755)
		untar(bytes.NewBuffer(downloadAndUncompressArchive()))
		if runtime.GOOS == "windows" {
			dir := path.Join(path.Dir(os.Getenv("MINGW_PREFIX")), "lib/pkgconfig")
			createDir(dir, 0755)
			name := path.Join(dir, "cef.pc")
			f, err := os.Create(name)
			checkFileError(err, "create", name)
			_, err = fmt.Fprintf(f, `Name: cef
Description: Chromium Embedded Framework
Version: %[1]s

Requires:
Libs: -L%[2]s/Release -lcef
Cflags: -I%[2]s
`, desiredCEFVersion, installPrefix)
			checkFileError(err, "write", name)
			checkFileError(f.Close(), "write", name)
		}
	}

	atexit.Exit(0)
}

func createDir(dir string, mode os.FileMode) {
	if err := os.MkdirAll(dir, mode); err != nil {
		fmt.Println(err)
		fmt.Println("You may need to run this as root.")
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

func archiveName() string {
	return fmt.Sprintf("cef_binary_%s_%s_minimal", desiredCEFVersion, cefPlatform)
}

func downloadAndUncompressArchive() []byte {
	client := http.Client{Timeout: 10 * time.Minute}
	url := fmt.Sprintf("http://opensource.spotify.com/cefbuilds/%s.tar.bz2", archiveName())
	fmt.Println("  Downloading...")
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Unable to download CEF archive at %s\n", url)
		fmt.Println(err)
		atexit.Exit(1)
	}
	buffer, err := ioutil.ReadAll(resp.Body)
	xio.CloseIgnoringErrors(resp.Body)
	if err != nil {
		fmt.Printf("Unable to download CEF archive at %s\n", url)
		fmt.Println(err)
		atexit.Exit(1)
	}
	if resp.StatusCode > 299 {
		fmt.Printf("Unable to download CEF archive at %s\n", url)
		fmt.Printf("Status: %s\n", resp.Status)
		atexit.Exit(1)
	}
	fmt.Println("  Uncompressing...")
	buffer, err = ioutil.ReadAll(bzip2.NewReader(bytes.NewReader(buffer)))
	if err != nil {
		fmt.Printf("Unable to uncompress CEF archive from %s\n", url)
		fmt.Println(err)
		atexit.Exit(1)
	}
	return buffer
}

func untar(in io.Reader) {
	prefix := archiveName()
	fmt.Println("  Unarchiving...")
	r := tar.NewReader(in)
	for {
		h, err := r.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Unable to read tar entry from archive")
			fmt.Println(err)
			atexit.Exit(1)
		}
		name := strings.Trim(strings.TrimPrefix(h.Name, prefix), "/")
		if name != "" && !strings.Contains(name, "..") {
			name = path.Join(installPrefix, name)
			switch h.Typeflag {
			case tar.TypeDir:
				createDir(name, os.FileMode(h.Mode|0555))
			case tar.TypeReg:
				buffer, err := ioutil.ReadAll(r)
				checkFileError(err, "read archive data for", name)
				checkFileError(ioutil.WriteFile(name, buffer, os.FileMode(h.Mode|0444)), "write", name)
			default:
				fmt.Printf("Unexpected type flag: %d\n", h.Typeflag)
				atexit.Exit(1)
			}
		}
	}
}
