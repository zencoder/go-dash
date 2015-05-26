# go-dash

[![godoc](https://godoc.org/github.com/zencoder/go-dash?status.svg)](http://godoc.org/github.com/zencoder/go-dash)

A Go library for generating MPEG-DASH manifests.

Install
-------

	go get github.com/zencoder/go-dash

Supported Features
-------

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

Known Limitations (for now) (PRs welcome)
--------
* Single Period
* No PSSH/PRO generation
* Limited Profile Support

Examples
--------

See [examples/](https://github.com/zencoder/go-dash/tree/master/examples)

To run (Live Profile example):
```
make examples-live
```

To run (OnDemand Profile example):
```
make examples-ondemand
```
