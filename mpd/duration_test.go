package mpd

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDuration(t *testing.T) {
	in := map[string]string{
		"0s":    "PT0S",
		"6m16s": "PT6M16S",
		"1.97s": "PT1.97S",
	}
	for ins, ex := range in {
		timeDur, err := time.ParseDuration(ins)
		assert.Equal(t, nil, err)
		dur := Duration(timeDur)
		assert.Equal(t, ex, dur.String())
	}
}

func TestParseDuration(t *testing.T) {
	in := map[string]float64{
		"PT0S":                 0,
		"PT1M":                 60,
		"PT2H":                 7200,
		"PT6M16S":              376,
		"PT1.97S":              1.97,
		"PT1H2M3.456S":         3723.456,
		"P1YT2M3S":             (365*24*time.Hour + 2*time.Minute + 3*time.Second).Seconds(),
		"P1Y2M3DT12H34M56.78S": (365*24*time.Hour + 59*24*time.Hour + 3*24*time.Hour + 12*time.Hour + 34*time.Minute + 56*time.Second + 780*time.Millisecond).Seconds(),
		"P1DT2H":               (26 * time.Hour).Seconds(),
		"P20M":                 (608 * 24 * time.Hour).Seconds(),
		"PT20M":                (20 * time.Minute).Seconds(),
		"P0Y20M0D":             (608 * 24 * time.Hour).Seconds(),
		"P0Y":                  0,
		"-P60D":                -(60 * 24 * time.Hour).Seconds(),
		"PT1M30.5S":            (time.Minute + 30*time.Second + 500*time.Millisecond).Seconds(),
	}
	for ins, ex := range in {
		act, err := parseDuration(ins)
		assert.NoError(t, err, ins)
		assert.Equal(t, ex, act.Seconds(), ins)
	}
}
