GO15VENDOREXPERIMENT := 1
COVERAGEDIR = coverage
ifdef CIRCLE_ARTIFACTS
  COVERAGEDIR = $(CIRCLE_ARTIFACTS)
endif

all: build test cover
fmt:
	go fmt ./...
test:
	if [ ! -d coverage ]; then mkdir coverage; fi
	go test -v ./mpd -race -cover -coverprofile=$(COVERAGEDIR)/mpd.coverprofile
cover:
	go tool cover -html=$(COVERAGEDIR)/mpd.coverprofile -o $(COVERAGEDIR)/mpd.html
tc: test cover
coveralls:
	gover $(COVERAGEDIR) $(COVERAGEDIR)/coveralls.coverprofile
	goveralls -coverprofile=$(COVERAGEDIR)/coveralls.coverprofile -service=circle-ci -repotoken=$(COVERALLS_TOKEN)
clean:
	go clean
	rm -rf coverage/
examples-live:
	go run examples/live.go
examples-ondemand:
	go run examples/ondemand.go
