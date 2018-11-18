#!/usr/bin/env bash
set -eo pipefail

CEF_VERSION=3.3538.1852.gcb937fc

# Setup PLATFORM
case $(uname -s) in
    Darwin*)  PLATFORM=macosx64 ;;
    Linux*)   PLATFORM=linux64 ;;
    MINGW64*) PLATFORM=windows64 ;;
    *)        echo "Unsupported OS: $(uname -s)"; false ;;
esac

if [ -e cef/include/cef_version.h ]; then
    EXISTING=`grep "#define CEF_VERSION " cef/include/cef_version.h | cut -f 2 -d '"'`
fi

if [ $CEF_VERSION != "$EXISTING" ]; then
    /bin/rm -rf cef cef_binary_${CEF_VERSION}_${PLATFORM}_minimal
    curl -L http://opensource.spotify.com/cefbuilds/cef_binary_${CEF_VERSION}_${PLATFORM}_minimal.tar.bz2 | bunzip2 | tar xf -
    mv cef_binary_${CEF_VERSION}_${PLATFORM}_minimal cef
fi
