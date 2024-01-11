#!/usr/bin/make --no-print-directory --jobs=1 --environment-overrides -f

VERSION_TAGS += REPLACE
REPLACE_MK_SUMMARY := go-corelibs/replace
REPLACE_MK_VERSION := v1.0.0

include CoreLibs.mk
