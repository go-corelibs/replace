#!/usr/bin/make --no-print-directory --jobs=1 --environment-overrides -f

CORELIB_PKG := go-corelibs/replace
VERSION_TAGS += MAIN
MAIN_MK_SUMMARY := ${CORELIB_PKG}
MAIN_MK_VERSION := v1.3.2

include CoreLibs.mk
