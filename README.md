# cef
Go bindings for the
[Chromium Embedded Framework (CEF)](https://bitbucket.org/chromiumembedded/cef).

Currently works for macOS and Windows. Linux support will be coming at some
point as well.

## Initial setup
Run `setup.sh` in this directory to pull the version of CEF these bindings
were created for.

## Extra environment setup required for Windows
- Download and run the installer from http://www.msys2.org/
- In the mingw64 msys2 console, run the following:
  - `pacman -Syu`
  - `pacman -Su`
  - `pacman -S mingw64/mingw-w64-x86_64-gcc mingw64/mingw-w64-x86_64-go msys/git`

## Example application
https://github.com/richardwilkes/webapp and
https://github.com/richardwilkes/webapp-example use these bindings to create
an example desktop application.

## Updating the CEF version to be used
The CEF version can be updated in the `setup.sh` script file by changing the
`CEF_VERSION` variable. If a different version is pulled, the source files
should be generated again by running `go generate ./...` on a macOS machine.
Code generation might be possible on other platforms, but has not been tested
there.
