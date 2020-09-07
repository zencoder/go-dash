package mpd

import (
	"encoding/base64"
	"encoding/hex"
	"github.com/liuyanhit/go-dash/helpers/require"
	"testing"
)

func TestMakePSSHBox_Widevine(t *testing.T) {
	expectedPSSH, err := base64.StdEncoding.DecodeString("AAAAYXBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAAEEIARIQWr3VL1VKTyq40GH3YUJRVRoIY2FzdGxhYnMiGFdyM1ZMMVZLVHlxNDBHSDNZVUpSVlE9PTIHZGVmYXVsdA==")
	if err != nil {
		panic(err.Error())
	}

	payload, err := base64.StdEncoding.DecodeString(VALID_WV_HEADER)
	if err != nil {
		panic(err.Error())
	}

	wvSystemID, err := hex.DecodeString(CONTENT_PROTECTION_WIDEVINE_SCHEME_HEX)
	if err != nil {
		panic(err.Error())
	}

	psshBox, err := makePSSHBox(wvSystemID, payload)
	require.NoError(t, err)

	require.EqualString(t, string(expectedPSSH), string(psshBox))
}

func TestMakePSSHBox_Playready(t *testing.T) {
	expectedPSSH, err := base64.StdEncoding.DecodeString("AAACJnBzc2gAAAAAmgTweZhAQoarkuZb4IhflQAAAgYGAgAAAQABAPwBPABXAFIATQBIAEUAQQBEAEUAUgAgAHgAbQBsAG4AcwA9ACIAaAB0AHQAcAA6AC8ALwBzAGMAaABlAG0AYQBzAC4AbQBpAGMAcgBvAHMAbwBmAHQALgBjAG8AbQAvAEQAUgBNAC8AMgAwADAANwAvADAAMwAvAFAAbABhAHkAUgBlAGEAZAB5AEgAZQBhAGQAZQByACIAIAB2AGUAcgBzAGkAbwBuAD0AIgA0AC4AMAAuADAALgAwACIAPgA8AEQAQQBUAEEAPgA8AFAAUgBPAFQARQBDAFQASQBOAEYATwA+ADwASwBFAFkATABFAE4APgAxADYAPAAvAEsARQBZAEwARQBOAD4APABBAEwARwBJAEQAPgBBAEUAUwBDAFQAUgA8AC8AQQBMAEcASQBEAD4APAAvAFAAUgBPAFQARQBDAFQASQBOAEYATwA+ADwASwBJAEQAPgBMADkAVwA5AFcAawBwAFYASwBrACsANAAwAEcASAAzAFkAVQBKAFIAVgBRAD0APQA8AC8ASwBJAEQAPgA8AEMASABFAEMASwBTAFUATQA+AEkASwB6AFkAMgBIAFoATABBAGwASQA9ADwALwBDAEgARQBDAEsAUwBVAE0APgA8AC8ARABBAFQAQQA+ADwALwBXAFIATQBIAEUAQQBEAEUAUgA+AA==")
	if err != nil {
		panic(err.Error())
	}

	// Base64 PRO
	payload, err := base64.StdEncoding.DecodeString(VALID_PLAYREADY_PRO)
	if err != nil {
		panic(err.Error())
	}

	wvSystemID, err := hex.DecodeString(CONTENT_PROTECTION_PLAYREADY_SCHEME_HEX)
	if err != nil {
		panic(err.Error())
	}

	psshBox, err := makePSSHBox(wvSystemID, payload)
	require.NoError(t, err)

	require.EqualString(t, string(expectedPSSH), string(psshBox))
}

func TestMakePSSHBox_BadSystemID(t *testing.T) {
	_, err := makePSSHBox([]byte("meaningless byte array"), nil)
	require.EqualError(t, err, "SystemID must be 16 bytes, was: 22")
}

func TestMakePSSHBox_NilSystemID(t *testing.T) {
	_, err := makePSSHBox(nil, nil)
	require.EqualError(t, err, "SystemID must be 16 bytes, was: 0")
}
