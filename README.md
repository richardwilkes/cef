# cef
Go bindings for the
[Chromium Embedded Framework (CEF)](https://bitbucket.org/chromiumembedded/cef).

Run `setup.sh` in this directory to pull the version of CEF these bindings
were created for. The exact version can be updated in that script file by
changing the `CEF_VERSION` variable. If a different version is pulled, the
source files should be generated again by running `go generate ./...`. Note
that source file generation has only been tested on macOS -- it may or may not
work correctly when run on other platforms.

https://github.com/richardwilkes/webapp and
https://github.com/richardwilkes/webapp-example use these bindings to create
desktop applications.
