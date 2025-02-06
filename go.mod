module github.com/zencoder/go-dash/v3

go 1.23

toolchain go1.23.2

require (
	github.com/Comcast/scte35-go v1.4.6
	github.com/go-chrono/chrono v0.0.0-20250124203826-0422557264a6
)

replace github.com/go-chrono/chrono v0.0.0-20250124203826-0422557264a6 => github.com/zencoder/chrono v0.0.0-20250206215435-caf544c317b8

require (
	github.com/bamiaux/iobit v0.0.0-20170418073505-498159a04883 // indirect
	golang.org/x/text v0.16.0 // indirect
)
