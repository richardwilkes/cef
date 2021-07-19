# NOTE:
*While this repo should still be functional for the specific revision of CEF it
was created for at the time, the project I originally created it for was
abandoned when the company decided it didn't want to embed a web browser after
all. This code has not been updated since and will likely not receive any more
attention from me. In particular, I know that the CEF project was in the process
of forcing the use of the sandbox when I stopped work on this. I did not spend
any time determining what changes would be necessary to enable that functionality
with those newer versions.*

# cef
Go bindings for the
[Chromium Embedded Framework (CEF)](https://bitbucket.org/chromiumembedded/cef).

Currently works for macOS, Windows and Linux.

## Initial setup required for Windows
- Download and run the installer from http://www.msys2.org/
- In the mingw64 msys2 console, run the following:
  - `pacman -Syu`
  - `pacman -Su`
  - `pacman -S mingw64/mingw-w64-x86_64-gcc mingw64/mingw-w64-x86_64-go mingw64/mingw-w64-x86_64-pkg-config msys/git`

## Initial setup
In the root of the repo, run:
```
go install
```
Then add `$GOPATH/bin` to your `$PATH` and run:
```
cef install
```
This will install the necessary CEF headers and libraries into /usr/local/cef.

## Example application
https://github.com/richardwilkes/webapp and
https://github.com/richardwilkes/webapp-example use these bindings to create
an example desktop application (currently macOS and Windows only).

## Updating the CEF version to be used
The CEF version can be updated in the `main.go` file by changing the
`desiredCEFVersion` variable. If a different CEF version is pulled, the
source files should be generated again by running `go generate ./...` on a
macOS machine. Code generation might be possible on other platforms, but has
not been tested there.
