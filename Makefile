ifdef VERBOSE
V = -v
X = -x
else
.SILENT:
endif

.DEFAULT_GOAL := all

.PHONY: all
all: test

.PHONY: test
test:
	go test $(V) ./... -race

.PHONY: generate
generate:
	GENERATE_FIXTURES=true $(MAKE) test

.PHONY: fmt
fmt:
	go fmt $(X) ./...

.PHONY: clean
clean:
	go clean -i $(X) -cache -testcache
