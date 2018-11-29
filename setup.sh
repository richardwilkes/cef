#!/usr/bin/env bash
set -eo pipefail

CEF_VERSION=3.3538.1852.gcb937fc

# Setup PLATFORM
case $(uname -s) in
    Darwin*)  PLATFORM=macosx64 ;;
    #Linux*)   PLATFORM=linux64 ;;
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
    # Prune out the things we don't need
    /bin/rm -rf cef/cmake cef/Debug \
        cef/include/capi/test \
        cef/include/capi/views \
        cef/include/test \
        cef/include/views \
        cef/include/wrapper \
        cef/libcef_dll
    /bin/rm -f cef/CMakeLists.txt \
        cef/Release/cef_sandbox.a \
        cef/include/capi/cef_*_util_capi.h \
        cef/include/capi/cef_parser_capi.h \
        cef/include/capi/cef_thread_capi.h \
        cef/include/capi/cef_trace_capi.h \
        cef/include/capi/cef_xml_reader_capi.h \
        cef/include/capi/cef_zip_reader_capi.h
    mv cef/include/cef_version.h cef/
    if [ $PLATFORM == "macosx64" ]; then
        mv cef/include/cef_application_mac.h cef/
    fi
    /bin/rm -f cef/include/*.h
    mv cef/*.h cef/include/
fi
