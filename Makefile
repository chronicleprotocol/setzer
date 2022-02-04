default: test

SHELL = bash
dirs = {bin,libexec}
prefix ?= /usr/local

dirs:; mkdir -p $(prefix)/$(dirs)
files = $(shell ls -d $(dirs)/*)
install:; cp -r -n $(dirs) $(prefix)
link: uninstall dirs; for x in $(files); do \
ln -s `pwd`/$$x $(prefix)/$$x; done
uninstall:; rm -rf $(addprefix $(prefix)/,$(files))

test:; ! grep '^#!/bin/sh' libexec/*/* && \
grep '^#!/usr/bin/env bash' libexec/*/* | \
cut -d: -f1 | xargs shellcheck

.PHONY: e2e
e2e:
	@if [ $$(docker ps | grep -c smocker) -eq 0 ]; then \
		docker run -d --restart=always -p 8080:8080 -p 8081:8081 --name smocker thiht/smocker; \
	fi
	docker build -t setzer -f e2e/Dockerfile .
	docker -v run -i --rm --link smocker setzer

# Suss out sources/pairs from script and call directly. Useful for debug. Note
# that this blindly pulls all pairs and tries to obtain a price from all
# sources, regardless of whether the source knows anything about the pair.
DEBUG=false
sf=libexec/setzer/setzer-x-price
.PHONY: price-dump
price-dump:
	@export DEBUG=$(DEBUG) && \
	for source in $$(grep -E '^  \w+)' $(sf) | awk '{print $$1}' | tr -d ')' ); do \
		for pair in $$(grep -E '^      \w+)' $(sf) | awk '{print $$1}' | tr -d ')' | sort -u); do \
			echo "=====> $$source / $$pair" && ./bin/setzer x-price $$source $$pair; echo "Exit status: $$?"; \
		done \
	done
