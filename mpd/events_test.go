package mpd

import (
	"testing"

	"github.com/zencoder/go-dash/helpers/ptrs"
	"github.com/zencoder/go-dash/helpers/require"
	"github.com/zencoder/go-dash/helpers/testfixtures"
)

const (
	VALID_EVENT_STREAM_SCHEME_ID_URI       = "urn:example:eventstream"
	VALID_EVENT_STREAM_VALUE               = "eventstream"
	VALID_EVENT_STREAM_TIMESCALE     int64 = 10
)

func newEventStreamMPD() *MPD {
	m := NewDynamicMPD(
		DASH_PROFILE_LIVE,
		VALID_AVAILABILITY_START_TIME,
		VALID_MIN_BUFFER_TIME,
	)
	p := m.GetCurrentPeriod()

	es := &EventStream{
		SchemeIDURI: ptrs.Strptr(VALID_EVENT_STREAM_SCHEME_ID_URI),
		Value:       ptrs.Strptr(VALID_EVENT_STREAM_VALUE),
		Timescale:   ptrs.Int64ptr(VALID_EVENT_STREAM_TIMESCALE),
	}

	e0 := &Event{
		ID:               ptrs.Strptr("event-0"),
		PresentationTime: ptrs.Int64ptr(100),
		Duration:         ptrs.Int64ptr(50),
	}

	e1 := &Event{
		ID:               ptrs.Strptr("event-1"),
		PresentationTime: ptrs.Int64ptr(200),
		Duration:         ptrs.Int64ptr(50),
	}

	es.Events = append(es.Events, e0, e1)
	p.EventStreams = append(p.EventStreams, es)

	return m
}

func TestEventStreamsWriteToString(t *testing.T) {
	m := newEventStreamMPD()

	got, err := m.WriteToString()
	require.NoError(t, err)

	testfixtures.CompareFixture(t, "fixtures/events.mpd", got)
}

func TestReadEventStreams(t *testing.T) {
	m, err := ReadFromFile("fixtures/events.mpd")
	require.NoError(t, err)

	got, err := m.WriteToString()
	require.NoError(t, err)

	testfixtures.CompareFixture(t, "fixtures/events.mpd", got)
}
