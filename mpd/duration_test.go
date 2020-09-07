package mpd

import (
	"fmt"
	"testing"
	"time"

	"github.com/liuyanhit/go-dash/helpers/require"
)

func TestDuration(t *testing.T) {
	in := map[string]string{
		"0s":    "PT0S",
		"6m16s": "PT6M16S",
		"1.97s": "PT1.97S",
	}
	for ins, ex := range in {
		timeDur, err := time.ParseDuration(ins)
		require.NoError(t, err)
		dur := Duration(timeDur)
		require.EqualString(t, ex, dur.String())
	}
}

func TestParseDuration(t *testing.T) {
	in := map[string]float64{
		"PT0S":          0,
		"PT1M":          60,
		"PT2H":          7200,
		"PT6M16S":       376,
		"PT1.97S":       1.97,
		"PT1H2M3.456S":  3723.456,
		"P1DT2H":        (26 * time.Hour).Seconds(),
		"PT20M":         (20 * time.Minute).Seconds(),
		"PT1M30.5S":     (time.Minute + 30*time.Second + 500*time.Millisecond).Seconds(),
		"PT1004199059S": (1004199059 * time.Second).Seconds(),
	}
	for ins, ex := range in {
		act, err := ParseDuration(ins)
		require.NoError(t, err, ins)
		require.EqualFloat64(t, ex, act.Seconds(), ins)
	}
}

func TestParseBadDurations(t *testing.T) {
	in := map[string]string{
		"P20M":   `Duration must be in the format: P[nD][T[nH][nM][nS]]`, // We don't allow Months (doesn't make sense when converting to duration)
		"P20Y":   `Duration must be in the format: P[nD][T[nH][nM][nS]]`, // We don't allow Years (doesn't make sense when converting to duration)
		"P15.5D": `Duration must be in the format: P[nD][T[nH][nM][nS]]`, // Only seconds can be expressed as a decimal
		"P2H":    `Duration must be in the format: P[nD][T[nH][nM][nS]]`, // "T" must be present to separate days and hours
		"2DT1H":  `Duration must be in the format: P[nD][T[nH][nM][nS]]`, // "P" must always be present
		"PT2M1H": `Duration must be in the format: P[nD][T[nH][nM][nS]]`, // Hours must appear before Minutes
		"P":      `At least one number and designator are required`,      // At least one number and designator are required
		"-P20H":  `Duration cannot be negative`,                          // Negative duration doesn't make sense
	}
	for ins, msg := range in {
		_, err := ParseDuration(ins)
		require.EqualError(t, err, msg, fmt.Sprintf("Expected an error for: %s", ins))
	}
}
