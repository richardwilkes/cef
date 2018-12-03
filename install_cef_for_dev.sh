#!/usr/bin/env bash
set -eo pipefail

CEF_VERSION=3.3538.1852.gcb937fc

case $(uname -s) in
    Darwin*)  PLATFORM=macosx64 ;;
#    Linux*)   PLATFORM=linux64 ;;
    MINGW64*) PLATFORM=windows64 ;;
    *)        echo "Unsupported OS: $(uname -s)"; false ;;
esac

CEF_PREFIX=/usr/local/cef
if [ -e "$CEF_PREFIX/include/cef_version.h" ]; then
    EXISTING=`grep "#define CEF_VERSION " "$CEF_PREFIX/include/cef_version.h" | cut -f 2 -d '"'`
fi
if [ $CEF_VERSION != "$EXISTING" ] || [ "$1" == "force" ]; then
    /bin/rm -rf "$CEF_PREFIX"
    mkdir -p "$CEF_PREFIX"
    curl -L http://opensource.spotify.com/cefbuilds/cef_binary_${CEF_VERSION}_${PLATFORM}_minimal.tar.bz2 | bunzip2 | tar xpof - -C "$CEF_PREFIX" --strip-components 1
    chmod -R go+rX "$CEF_PREFIX"
fi
