#! /usr/bin/make
#
# Makefile for goa v2 plugins

PLUGINS=$(shell find . -mindepth 1 -maxdepth 1 -not -path "*/\.*" -type d)

all: test-plugins test-aliaser

gen:
	@for p in $(PLUGINS) ; do \
		make -C $$p gen || exit 1; \
	done

aliases:
	@for p in $(PLUGINS) ; do \
		make -C $$p aliases || exit 1; \
	done

test-plugins:
	@for p in $(PLUGINS) ; do \
		make -C $$p || exit 1; \
	done

test-aliaser:
	@for p in $(PLUGINS) ; do \
		make -k -C $$p test-aliaser || exit 1; \
	done
