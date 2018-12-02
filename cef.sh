#!/usr/bin/env bash
set -eo pipefail

CEF_VERSION=3.3538.1852.gcb937fc

SELF=$0
function help() {
    echo "$SELF [options]"
    echo ""
    echo "    --clear           Clears the cache"
    echo "    --headers <path>  Install the CEF headers into <path>, replacing existing contents"
    echo "    --libs <path>     Install the CEF libraries into <path>, replacing existing contents"
    echo "    -c                Same as --clear"
    echo "    -H <path>         Same as --headers"
    echo "    -l <path>         Same as --libs"
}

# Process the command line
while (( "$#" )); do
    case "$1" in
        -c|--clear)
            CLEAR_CACHE=1
            shift
            ;;
        -H|--headers)
            if [ "x$2" == "x" ]; then
                echo "$1 requires a path"
                exit 1
            fi
            INSTALL_HEADERS=$2
            shift 2
            ;;
        -l|--libs)
            if [ "x$2" == "x" ]; then
                echo "$1 requires a path"
                exit 1
            fi
            INSTALL_LIBS=$2
            shift 2
            ;;
        *)
            help
            exit 1
            ;;
    esac
done
if [ "x$CLEAR_CACHE" == "x" ] && [ "x$INSTALL_HEADERS" == "x" ] && [ "x$INSTALL_LIBS" == "x" ]; then
    help
    exit 1
fi

# Setup PLATFORM and CACHEDIR
case $(uname -s) in
    Darwin*)
        PLATFORM=macosx64
        CACHEDIR="$HOME/Library/Caches"
        ;;
#    Linux*)
#        PLATFORM=linux64
#        if [ "$XDG_CACHE_HOME" != "" ]; then
#            CACHEDIR="$XDG_CACHE_HOME"
#        else
#            CACHEDIR="$HOME/.cache"
#        fi
#        ;;
    MINGW64*)
        PLATFORM=windows64
        CACHEDIR="$LocalAppData"
        ;;
    *)
        echo "Unsupported OS: $(uname -s)"
        exit 1
        ;;
esac

CACHEDIR="$CACHEDIR/gocef"
if [ "x$CLEAR_CACHE" == "x1" ]; then
    /bin/rm -rf "$CACHEDIR"
fi

CEF_BASE=cef_binary_${CEF_VERSION}_${PLATFORM}_minimal
CEF_ARCHIVE=$CEF_BASE.tar.bz2

# Function to download the CEF archive if it doesn't already exist
function download_cef_archive() {
    if [ ! -f "$CACHEDIR/$CEF_ARCHIVE" ]; then
        mkdir -p "$CACHEDIR"
        curl -L -o "$CACHEDIR/$CEF_ARCHIVE" http://opensource.spotify.com/cefbuilds/$CEF_ARCHIVE
    fi
}

# Install the CEF headers
if [ "x$INSTALL_HEADERS" != "x" ]; then
    EXISTING=
    if [ -e "$INSTALL_HEADERS/include/cef_version.h" ]; then
        EXISTING=`grep "#define CEF_VERSION " "$INSTALL_HEADERS/include/cef_version.h" | cut -f 2 -d '"'`
    fi
    if [ $CEF_VERSION != "$EXISTING" ]; then
        download_cef_archive
        /bin/rm -rf "$INSTALL_HEADERS"
        mkdir -p "$INSTALL_HEADERS"
        EXTRA_INCLUDE=
        if [ $PLATFORM == "macosx64" ]; then
            EXTRA_INCLUDE="--include $CEF_BASE/include/cef_application_mac.h"
        fi
        bunzip2 --stdout "$CACHEDIR/$CEF_ARCHIVE" | tar xf - -C "$INSTALL_HEADERS" \
            --strip-components 1 $EXTRA_INCLUDE \
            --include $CEF_BASE/include/cef_version.h \
            --include $CEF_BASE/include/base \
            --include $CEF_BASE/include/internal \
            --include $CEF_BASE/include/capi \
            --exclude $CEF_BASE/include/capi/cef_\*_util_capi.h \
            --exclude $CEF_BASE/include/capi/cef_parser_capi.h \
            --exclude $CEF_BASE/include/capi/cef_thread_capi.h \
            --exclude $CEF_BASE/include/capi/cef_trace_capi.h \
            --exclude $CEF_BASE/include/capi/cef_xml_reader_capi.h \
            --exclude $CEF_BASE/include/capi/cef_zip_reader_capi.h \
            --exclude $CEF_BASE/include/capi/test \
            --exclude $CEF_BASE/include/capi/views
    fi
fi

# Install the CEF libraries
if [ "x$INSTALL_LIBS" != "x" ]; then
    EXISTING=
    if [ -e "$INSTALL_LIBS/version.txt" ]; then
        EXISTING=`cat "$INSTALL_LIBS/version.txt"`
    fi
    if [ $CEF_VERSION != "$EXISTING" ]; then
        download_cef_archive
        /bin/rm -rf "$INSTALL_LIBS"
        mkdir -p "$INSTALL_LIBS"
        EXTRA_INCLUDE=
        if [ $PLATFORM == "windows64" ]; then
            EXTRA_INCLUDE="--include $CEF_BASE/Resources"
        fi
        bunzip2 --stdout "$CACHEDIR/$CEF_ARCHIVE" | tar xf - -C "$INSTALL_LIBS" \
            --strip-components 2 $EXTRA_INCLUDE \
            --include $CEF_BASE/Release \
            --exclude $CEF_BASE/Release/cef_sandbox.a
        echo "$CEF_VERSION" > "$INSTALL_LIBS/version.txt"
    fi
fi
