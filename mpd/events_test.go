package mpd

import (
	"testing"

	"github.com/garkettleung/go-dash/v3/helpers/ptrs"
	"github.com/garkettleung/go-dash/v3/helpers/require"
	"github.com/garkettleung/go-dash/v3/helpers/testfixtures"
)

const (
	VALID_EVENT_STREAM_SCHEME_ID_URI      = "urn:example:eventstream"
	VALID_EVENT_STREAM_VALUE              = "eventstream"
	VALID_EVENT_STREAM_TIMESCALE     uint = 10
)

func newEventStreamMPD() *MPD {
	m := NewDynamicMPD(
		DASH_PROFILE_LIVE,
		VALID_AVAILABILITY_START_TIME,
		VALID_MIN_BUFFER_TIME,
	)
	p := m.GetCurrentPeriod()

	es := EventStream{
		SchemeIDURI: ptrs.Strptr(VALID_EVENT_STREAM_SCHEME_ID_URI),
		Value:       ptrs.Strptr(VALID_EVENT_STREAM_VALUE),
		Timescale:   ptrs.Uintptr(VALID_EVENT_STREAM_TIMESCALE),
	}

	e0 := Event{
		ID:               ptrs.Strptr("event-0"),
		PresentationTime: ptrs.Uint64ptr(100),
		Duration:         ptrs.Uint64ptr(50),
	}

	e1 := Event{
		ID:               ptrs.Strptr("event-1"),
		PresentationTime: ptrs.Uint64ptr(200),
		Duration:         ptrs.Uint64ptr(50),
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
