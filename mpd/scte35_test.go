package mpd

import (
	"strconv"
	"testing"

	. "github.com/zencoder/go-dash/v3/helpers/ptrs"
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

	expectedMPD := &MPD{
		XMLNs:                     Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles:                  Strptr((string)(DASH_PROFILE_LIVE)),
		Type:                      Strptr("static"),
		MediaPresentationDuration: Strptr(VALID_MEDIA_PRESENTATION_DURATION),
		MinBufferTime:             Strptr(VALID_MIN_BUFFER_TIME),
		AvailabilityStartTime:     Strptr(VALID_AVAILABILITY_START_TIME),
		period:                    nil,
		Periods: []*Period{
			{ID: "0"},
			{ID: "1", EventStreams: []EventStream{
				{
					Timescale:   Uintptr(VALID_EVENT_STREAM_TIMESCALE),
					SchemeIDURI: Strptr(SCTE352014SchemaUri),
					Events: []Event{
						{
							ID:               Strptr("1"),
							PresentationTime: Uint64ptr(10000),
							Signals: []Signal{
								{
									XMLNs: Strptr(SCTE352014SchemaUri),
									Binaries: []Binary{
										{
											XMLNs:      Strptr(SCTE352014SchemaUri),
											BinaryData: Strptr(""),
										},
									},
								},
							},
						},
						{
							ID:               Strptr("2"),
							PresentationTime: Uint64ptr(15000),
							Signals: []Signal{
								{
									XMLNs: Strptr(SCTE352014SchemaUri),
									Binaries: []Binary{
										{
											XMLNs:      Strptr(SCTE352014SchemaUri),
											BinaryData: Strptr(""),
										},
									},
								},
							},
						},
						{
							ID:               Strptr("3"),
							PresentationTime: Uint64ptr(20000),
							Signals: []Signal{
								{
									XMLNs: Strptr(SCTE352014SchemaUri),
									Binaries: []Binary{
										{
											XMLNs:      Strptr(SCTE352014SchemaUri),
											BinaryData: Strptr("/DAvAAAAAAAAAP/wFAUAAAsgf/+AAAAAAAcT/w=="),
										},
									},
								},
							},
						},
					},
				},
			}},
			{ID: "2"},
		},
	}

	expectedString, err := expectedMPD.WriteToString()
	require.NoError(t, err)
	actualString, err := m.WriteToString()
	require.NoError(t, err)

	require.EqualString(t, expectedString, actualString)
}
