package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strconv"
	"time"

	"github.com/richardwilkes/toolbox/atexit"
	"github.com/richardwilkes/toolbox/cmdline"
	"github.com/richardwilkes/toolbox/xio/fs"
)

type dist struct {
	root           string
	displayName    string
	bundleName     string
	bundleID       string
	exeName        string
	icon           string
	version        string
	shortVersion   string
	copyrightYears string
	copyrightOwner string
}

// NewDist returns the dist command.
func NewDist() cmdline.Cmd {
	d := &dist{
		root:           "dist",
		displayName:    "Example",
		bundleName:     "Example",
		bundleID:       "com.example",
		exeName:        "example",
		version:        "1.0.0",
		shortVersion:   "1.0",
		copyrightYears: strconv.Itoa(time.Now().Year()),
		copyrightOwner: "Unknown",
	}
	switch runtime.GOOS {
	case "darwin":
		d.root = path.Join(d.root, "macos")
		d.icon = "AppIcon.icns"
	default:
		d.root = path.Join(d.root, runtime.GOOS)
	}
	return d
}

func (d *dist) Name() string {
	return "dist"
}

func (d *dist) Usage() string {
	return "Creates a distribution tree, adding the necessary CEF libraries."
}

func (d *dist) Run(cl *cmdline.CmdLine, args []string) error {
	cl.NewStringOption(&d.root).SetSingle('d').SetName("dir").SetUsage("Set the root distribution directory")
	cl.NewStringOption(&d.displayName).SetSingle('n').SetName("name").SetUsage("Set the display name (macOS-only)")
	cl.NewStringOption(&d.bundleName).SetSingle('b').SetName("bundle").SetUsage("Set the bundle name (macOS-only)")
	cl.NewStringOption(&d.bundleID).SetSingle('B').SetName("id").SetUsage("Set the bundle ID (macOS-only)")
	cl.NewStringOption(&d.exeName).SetSingle('e').SetName("executable").SetUsage("Set the executable name (macOS-only)")
	cl.NewStringOption(&d.icon).SetSingle('i').SetName("icon").SetUsage("Set the icon (macOS-only)")
	cl.NewStringOption(&d.version).SetSingle('r').SetName("release").SetUsage("Set the release version (macOS-only)")
	cl.NewStringOption(&d.shortVersion).SetSingle('s').SetName("short-release").SetUsage("Set the short release version (macOS-only)")
	cl.NewStringOption(&d.copyrightYears).SetSingle('y').SetName("year").SetUsage("Set the copyright year(s) (macOS-only)")
	cl.NewStringOption(&d.copyrightOwner).SetSingle('o').SetName("owner").SetUsage("Set the copyright owner (macOS-only)")
	cl.Parse(args)
	checkPlatform()
	if err := os.RemoveAll(d.root); err != nil {
		fmt.Printf("Unable to remove old dist at %s\n", d.root)
		fmt.Println(err)
		atexit.Exit(1)
	}
	createDir(d.root, 0755)
	switch runtime.GOOS {
	case "darwin":
		d.distMacOS()
	case "windows":
		d.distWindows()
	default:
		return fmt.Errorf("Unhandled OS: %s", runtime.GOOS)
	}
	return nil
}

func (d *dist) distMacOS() {
	appBundleContentsDir := path.Join(d.root, d.displayName+".app", "Contents")
	createDir(path.Join(appBundleContentsDir, "MacOS"), 0755)
	appBundleResourcesDir := path.Join(appBundleContentsDir, "Resources")
	createDir(appBundleResourcesDir, 0755)
	appFrameworksDir := path.Join(appBundleContentsDir, "Frameworks")
	helperAppBundleContentsDir := path.Join(appFrameworksDir, d.exeName+" Helper.app", "Contents")
	helperAppBundleMacOSDir := path.Join(helperAppBundleContentsDir, "MacOS")
	createDir(helperAppBundleMacOSDir, 0755)
	createDir(path.Join(helperAppBundleContentsDir, "Frameworks"), 0755)
	releaseDir := path.Join(installPrefix, "Release")
	cc := exec.Command("cc", "-I", installPrefix, path.Join(installPrefix, "helper", "helper.c"), "-F", releaseDir, "-framework", "Chromium Embedded Framework", "-o", path.Join(helperAppBundleMacOSDir, d.exeName+" Helper"))
	if result, err := cc.CombinedOutput(); err != nil {
		fmt.Println("Failed to compile the helper.")
		fmt.Println(err)
		fmt.Println(string(result))
		atexit.Exit(1)
	}
	if err := fs.Copy(path.Join(releaseDir, "Chromium Embedded Framework.framework"), path.Join(appFrameworksDir, "Chromium Embedded Framework.framework")); err != nil {
		fmt.Println(err)
		atexit.Exit(1)
	}
	if err := os.Symlink("../../../Chromium Embedded Framework.framework", path.Join(helperAppBundleContentsDir, "Frameworks", "Chromium Embedded Framework.framework")); err != nil {
		fmt.Println(err)
		atexit.Exit(1)
	}
	if err := fs.Copy(d.icon, path.Join(appBundleResourcesDir, "AppIcon.icns")); err != nil {
		fmt.Println(err)
		atexit.Exit(1)
	}

	plist := path.Join(appBundleContentsDir, "Info.plist")
	f, err := os.Create(plist)
	checkFileError(err, "create", plist)
	_, err = fmt.Fprintf(f, `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleInfoDictionaryVersion</key>
	<string>6.0</string>
	<key>CFBundleDisplayName</key>
	<string>%s</string>
	<key>CFBundleName</key>
	<string>%s</string>
	<key>CFBundleExecutable</key>
	<string>%s</string>
	<key>CFBundleIconFile</key>
	<string>AppIcon.icns</string>
	<key>CFBundleIdentifier</key>
	<string>%s</string>
	<key>CFBundlePackageType</key>
	<string>APPL</string>
	<key>CFBundleVersion</key>
	<string>%s</string>
	<key>CFBundleShortVersionString</key>
	<string>%s</string>
	<key>NSHumanReadableCopyright</key>
	<string>© %s by %s. All rights reserved.</string>
	<key>NSHighResolutionCapable</key>
	<true/>
	<key>NSSupportsAutomaticGraphicsSwitching</key>
	<true/>
</dict>
</plist>
`, d.displayName, d.bundleName, d.exeName, d.bundleID, d.version, d.shortVersion, d.copyrightYears, d.copyrightOwner)
	checkFileError(err, "write", plist)
	checkFileError(f.Close(), "write", plist)

	plist = path.Join(helperAppBundleContentsDir, "Info.plist")
	f, err = os.Create(plist)
	checkFileError(err, "create", plist)
	_, err = fmt.Fprintf(f, `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleInfoDictionaryVersion</key>
	<string>6.0</string>
	<key>CFBundleDisplayName</key>
	<string>%s Helper</string>
	<key>CFBundleName</key>
	<string>%s Helper</string>
	<key>CFBundleExecutable</key>
	<string>%s Helper</string>
	<key>CFBundleIdentifier</key>
	<string>%s.helper</string>
	<key>CFBundlePackageType</key>
	<string>APPL</string>
	<key>CFBundleVersion</key>
	<string>%s</string>
	<key>CFBundleShortVersionString</key>
	<string>%s</string>
	<key>NSHumanReadableCopyright</key>
	<string>© %s by %s. All rights reserved.</string>
	<key>NSHighResolutionCapable</key>
	<true/>
	<key>NSSupportsAutomaticGraphicsSwitching</key>
	<true/>
</dict>
</plist>
`, d.displayName, d.bundleName, d.exeName, d.bundleID, d.version, d.shortVersion, d.copyrightYears, d.copyrightOwner)
	checkFileError(err, "write", plist)
	checkFileError(f.Close(), "write", plist)
}

func (d *dist) distWindows() {
	copyDirContents(path.Join(installPrefix, "Release"), d.root)
	copyDirContents(path.Join(installPrefix, "Resources"), d.root)
}

func copyDirContents(srcdir, dstdir string) {
	list, err := ioutil.ReadDir(srcdir)
	if err != nil {
		fmt.Println(err)
		atexit.Exit(1)
	}
	for _, one := range list {
		name := one.Name()
		if err := fs.Copy(path.Join(srcdir, name), path.Join(dstdir, name)); err != nil {
			fmt.Println(err)
			atexit.Exit(1)
		}
	}
}
