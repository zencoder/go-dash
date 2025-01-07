package mpd

import (
	"os"
	"strconv"
	"testing"

	"github.com/zencoder/go-dash/v3/helpers/require"
)

func TestPeriod_AddNewSCTE35Break(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME,
		AttrAvailabilityStartTime(VALID_AVAILABILITY_START_TIME))
	require.NotNil(t, m)

	for i := 0; i < 3; i++ {
		period := m.AddNewPeriod()
		period.ID = strconv.Itoa(i)
	}

	m.Periods[1].AddNewSCTE35Break(VALID_EVENT_STREAM_TIMESCALE, 10000, "1", "")
	m.Periods[1].AddNewSCTE35Break(VALID_EVENT_STREAM_TIMESCALE, 20000, "3", "/DAvAAAAAAAAAP/wFAUAAAsgf/+AAAAAAAcT/w==")
	m.Periods[1].AddNewSCTE35Break(VALID_EVENT_STREAM_TIMESCALE, 15000, "2", "")

	expected, err := os.ReadFile("fixtures/scte35.mpd")
	require.NoError(t, err)

	actual, err := m.WriteToString()
	require.NoError(t, err)

	require.EqualString(t, string(expected), actual)
}
