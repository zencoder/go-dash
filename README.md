# go-dash

[![godoc](https://godoc.org/github.com/zencoder/go-dash/mpd?status.svg)](http://godoc.org/github.com/zencoder/go-dash/mpd)

A Go library for generating MPEG-DASH manifests.

## Install

	go get github.com/zencoder/go-dash/mpd

## Supported Features

* Profiles
  * Live
  * On Demand
* Adaption Sets / Representations
  * Audio
  * Video
  * Subtitles
* DRM (ContentProtection)
  * PlayReady
  * Widevine

## Known Limitations (for now) (PRs welcome)

* Single Period
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

Tested on go 1.5.1.

### Build and run unit tests

    make test
    
### CI

[This library builds on Circle CI, here.](https://circleci.com/gh/zencoder/go-dash/)

## License

[Apache License Version 2.0](LICENSE)
