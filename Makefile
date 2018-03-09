COVERAGEDIR = coverage
ifdef CIRCLE_ARTIFACTS
  COVERAGEDIR = $(CIRCLE_ARTIFACTS)
endif

ifdef VERBOSE
V = -v
else
.SILENT:
endif

all: test cover

fmt:
	find . -not -path "./vendor/*" -name '*.go' -type f | sed 's#\(.*\)/.*#\1#' | sort -u | xargs -n1 -I {} bash -c "cd {} && goimports -w *.go && gofmt -w -s -l *.go"

test:
	mkdir -p coverage
	go test $(V) ./mpd -race -cover -coverprofile=$(COVERAGEDIR)/mpd.coverprofile

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

generate:
	GENERATE_FIXTURES=true $(MAKE) test
