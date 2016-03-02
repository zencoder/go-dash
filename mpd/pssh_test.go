package mpd

import (
	"encoding/base64"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakePSSHBox(t *testing.T) {
	expectedOutput, err := base64.StdEncoding.DecodeString("AAAAYXBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAAEEIARIQWr3VL1VKTyq40GH3YUJRVRoIY2FzdGxhYnMiGFdyM1ZMMVZLVHlxNDBHSDNZVUpSVlE9PTIHZGVmYXVsdA==")
	if err != nil {
		panic(err.Error())
	}

	payload, err := base64.StdEncoding.DecodeString("CAESEFq91S9VSk8quNBh92FCUVUaCGNhc3RsYWJzIhhXcjNWTDFWS1R5cTQwR0gzWVVKUlZRPT0yB2RlZmF1bHQ=")
	if err != nil {
		panic(err.Error())
	}

	wvSystemID, err := hex.DecodeString(CONTENT_PROTECTION_WIDEVINE_SCHEME_HEX)
	if err != nil {
		panic(err.Error())
	}

	psshBox, err := makePSSHBox(wvSystemID, payload)

	assert.Equal(t, expectedOutput, psshBox)
}
