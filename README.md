# go-dash

[![godoc](https://godoc.org/github.com/zencoder/go-dash/mpd?status.svg)](http://godoc.org/github.com/zencoder/go-dash/mpd)

A Go library for generating MPEG-DASH manifests.

## Install

This library uses [Glide](https://github.com/Masterminds/glide) to manage it's dependencies. Please refer to the Glide documentation to see how to install glide.

```bash
mkdir -p $GOPATH/src/github.com/zencoder
cd $GOPATH/src/github.com/zencoder
git clone https://github.com/zencoder/go-dash
cd go-dash
export GO15VENDOREXPERIMENT=1
glide install
go install ./...
```

## Supported Features

* Profiles
  * Live
  * On Demand
* Adaption Sets / Representations / Roles
  * Audio
  * Video
  * Subtitles
  * Multiple periods (multi-part playlist)
* DRM (ContentProtection)
  * PlayReady
  * Widevine

## Known Limitations (for now) (PRs welcome)

* No PSSH/PRO generation
* Limited Profile Support

## Example Usage

See [examples/](https://github.com/zencoder/go-dash/tree/master/examples)

To run (Live Profile example):
```
make examples-live
```

To run (OnDemand Profile example):
```
make examples-ondemand
```

## Development

### Dependencies

Tested on go 1.8.3.

### Build and run unit tests

    make test

### CI

[This library builds on Circle CI, here.](https://circleci.com/gh/zencoder/go-dash/)

## License

[Apache License Version 2.0](LICENSE)
