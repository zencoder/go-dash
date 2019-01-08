ifdef VERBOSE
V = -v
X = -x
else
.SILENT:
endif

.DEFAULT_GOAL := all

.PHONY: all
all: test

vendor:
	glide install

.PHONY: test
test: vendor
	go test $(V) ./... -race

.PHONY: generate
generate: vendor
	GENERATE_FIXTURES=true $(MAKE) test

.PHONY: fmt
fmt:
	find . -not -path "./vendor/*" -name '*.go' -type f | sed 's#\(.*\)/.*#\1#' | sort -u | xargs -n1 -I {} bash -c "cd {} && goimports -w *.go && gofmt -w -s -l *.go"

.PHONY: clean
clean:
	rm -rf vendor/
	go clean -i $(X) -cache -testcache
