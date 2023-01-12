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

.PHONY: test
test:
	make e2e

.PHONY: e2e
e2e:
	@if [ $$(docker ps | grep -c smocker) -eq 0 ]; then \
		docker run -d --restart=always -p 8080:8080 -p 8081:8081 --name smocker thiht/smocker; \
	fi
	docker build -t setzer -f e2e/Dockerfile .
	docker -v run -i --rm --link smocker setzer

.PHONY: unit-test
unit-test:
	./test/units.sh
