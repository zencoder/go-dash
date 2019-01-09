# go-dash [![godoc](https://godoc.org/github.com/zencoder/go-dash/mpd?status.svg)](http://godoc.org/github.com/zencoder/go-dash/mpd)

A [Go](https://golang.org) library for generating [MPEG-DASH](https://en.wikipedia.org/wiki/Dynamic_Adaptive_Streaming_over_HTTP) manifests.

## Install

```
go get -u github.com/zencoder/go-dash
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

See the [examples/](https://github.com/zencoder/go-dash/tree/master/examples) directory.

## Development

```
make test
```

### CI

[This project builds in Circle CI](https://circleci.com/gh/zencoder/go-dash/)

## License

[Apache License Version 2.0](LICENSE)
