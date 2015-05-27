GO ?= go
COVERAGEDIR = coverage
ifdef CIRCLE_ARTIFACTS
  COVERAGEDIR = $(CIRCLE_ARTIFACTS)
endif

all: build test cover
build:
	if [ ! -d bin ]; then mkdir bin; fi
	$(GO) build -v -o bin/go-dash
fmt:
	$(GO) fmt ./...
test:
	if [ ! -d coverage ]; then mkdir coverage; fi
	$(GO) test -v ./mpd -race -cover -coverprofile=$(COVERAGEDIR)/mpd.coverprofile
cover:
	$(GO) tool cover -html=$(COVERAGEDIR)/mpd.coverprofile -o $(COVERAGEDIR)/mpd.html
tc: test cover
coveralls:
	gover $(COVERAGEDIR) $(COVERAGEDIR)/coveralls.coverprofile
	goveralls -coverprofile=$(COVERAGEDIR)/coveralls.coverprofile -service=circle-ci -repotoken=$(COVERALLS_TOKEN)
clean:
	$(GO) clean
	rm -f bin/go-dash
	rm -rf coverage/
examples-live:
	$(GO) run examples/live.go
examples-ondemand:
	$(GO) run examples/ondemand.go
