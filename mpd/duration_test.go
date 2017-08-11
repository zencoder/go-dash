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
