package mpd

import (
	"strconv"
	"testing"
	"time"

	"github.com/liuyanhit/go-dash/helpers/ptrs"
	"github.com/liuyanhit/go-dash/helpers/require"
	"github.com/liuyanhit/go-dash/helpers/testfixtures"
)

func TestSegmentTimelineSerialization(t *testing.T) {
	testcases := []struct {
		In  *MPD
		Out string
	}{
		{In: getSegmentTimelineMPD(), Out: "segment_timeline.mpd"},
		{In: getMultiPeriodSegmentTimelineMPD(), Out: "segment_timeline_multi_period.mpd"},
	}
	for _, tc := range testcases {
		t.Run(tc.Out, func(t *testing.T) {
			found, err := tc.In.WriteToString()
			require.NoError(t, err)
			testfixtures.CompareFixture(t, "fixtures/"+tc.Out, found)
		})
	}
}

func TestSegmentTimelineDeserialization(t *testing.T) {
	xml := testfixtures.LoadFixture("fixtures/segment_timeline.mpd")
	m, err := ReadFromString(xml)
	require.NoError(t, err)
	expected := getSegmentTimelineMPD()
	require.EqualString(t, expected.Periods[0].BaseURL, m.Periods[0].BaseURL)

	expectedAudioSegTimeline := expected.Periods[0].AdaptationSets[0].Representations[0].SegmentTemplate.SegmentTimeline
	audioSegTimeline := m.Periods[0].AdaptationSets[0].Representations[0].SegmentTemplate.SegmentTimeline

	for i := range expectedAudioSegTimeline.Segments {
		require.EqualUInt64(t, expectedAudioSegTimeline.Segments[i].Duration, audioSegTimeline.Segments[i].Duration)
		require.EqualUInt64Ptr(t, expectedAudioSegTimeline.Segments[i].StartTime, audioSegTimeline.Segments[i].StartTime)
		require.EqualIntPtr(t, expectedAudioSegTimeline.Segments[i].RepeatCount, audioSegTimeline.Segments[i].RepeatCount)
	}

	expectedVideoSegTimeline := expected.Periods[0].AdaptationSets[1].Representations[0].SegmentTemplate.SegmentTimeline
	videoSegTimeline := m.Periods[0].AdaptationSets[1].Representations[0].SegmentTemplate.SegmentTimeline

	for i := range expectedVideoSegTimeline.Segments {
		require.EqualUInt64(t, expectedVideoSegTimeline.Segments[i].Duration, videoSegTimeline.Segments[i].Duration)
		require.EqualUInt64Ptr(t, expectedVideoSegTimeline.Segments[i].StartTime, videoSegTimeline.Segments[i].StartTime)
		require.EqualIntPtr(t, expectedVideoSegTimeline.Segments[i].RepeatCount, videoSegTimeline.Segments[i].RepeatCount)
	}
}

func getMultiPeriodSegmentTimelineMPD() *MPD {
	m := NewMPD(DASH_PROFILE_LIVE, "PT65.063S", "PT2.000S")
	for i := 0; i < 4; i++ {
		if i > 0 {
			m.AddNewPeriod()
		}
		p := m.GetCurrentPeriod()
		p.ID = strconv.Itoa(i)
		p.Duration = Duration(30 * time.Second)
		aas, _ := p.AddNewAdaptationSetAudioWithID("1", "audio/mp4", true, 1, "en")
		_, _ = aas.AddNewRepresentationAudio(48000, 92000, "mp4a.40.2", "audio_1")
		aas.SegmentTemplate = &SegmentTemplate{
			Timescale:      ptrs.Int64ptr(48000),
			Initialization: ptrs.Strptr("audio/init.m4f"),
			Media:          ptrs.Strptr("audio/segment$Number$.m4f"),
			SegmentTimeline: &SegmentTimeline{
				Segments: []*SegmentTimelineSegment{
					{Duration: 95232, RepeatCount: ptrs.Intptr(14)},
					{Duration: 15360},
				},
			},
		}
		vas, _ := p.AddNewAdaptationSetVideoWithID("2", "video/mp4", "progressive", true, 1)
		_, _ = vas.AddNewRepresentationVideo(3532000, "avc1.640028", "video_1", "2997/100", 2048, 854)
		_, _ = vas.AddNewRepresentationVideo(453000, "avc1.420016", "video_2", "2997/100", 648, 270)
		vas.SegmentTemplate = &SegmentTemplate{
			Timescale:      ptrs.Int64ptr(30000),
			Initialization: ptrs.Strptr("video/$RepresentationID$/init.m4f"),
			Media:          ptrs.Strptr("video/$RepresentationID$/segment$Number$.m4f"),
			SegmentTimeline: &SegmentTimeline{
				Segments: []*SegmentTimelineSegment{
					{Duration: 58058, RepeatCount: ptrs.Intptr(14)},
					{Duration: 31031},
				},
			},
		}
		// Add Special flags on 3rd Period, to simulate cutting a piece of content midway
		if i == 2 {
			aas.SegmentTemplate.StartNumber = ptrs.Int64ptr(17)
			aas.SegmentTemplate.PresentationTimeOffset = ptrs.Uint64ptr(743424)
			vas.SegmentTemplate.StartNumber = ptrs.Int64ptr(17)
			vas.SegmentTemplate.PresentationTimeOffset = ptrs.Uint64ptr(464464)
		}
	}
	return m
}

func getSegmentTimelineMPD() *MPD {
	m := NewMPD(DASH_PROFILE_LIVE, "PT65.063S", "PT2.000S")
	m.period.BaseURL = "http://localhost:8002/public/"

	aas, _ := m.AddNewAdaptationSetAudioWithID("1", "audio/mp4", true, 1, "English")
	ra, _ := aas.AddNewRepresentationAudio(48000, 255000, "mp4a.40.2", "audio_1")

	ra.SegmentTemplate = &SegmentTemplate{
		Timescale:      ptrs.Int64ptr(48000),
		Initialization: ptrs.Strptr("audio/init.m4f"),
		Media:          ptrs.Strptr("audio/segment$Number$.m4f"),
		SegmentTimeline: &SegmentTimeline{
			Segments: []*SegmentTimelineSegment{
				{StartTime: ptrs.Uint64ptr(0), Duration: 231424},
				{RepeatCount: ptrs.Intptr(2), Duration: 479232},
				{Duration: 10240},
				{Duration: 247808},
				{RepeatCount: ptrs.Intptr(1), Duration: 479232},
				{Duration: 3072},
			},
		},
	}

	vas, _ := m.AddNewAdaptationSetVideoWithID("2", "video/mp4", "progressive", true, 1)
	va, _ := vas.AddNewRepresentationVideo(int64(4172274), "avc1.640028", "video_1", "30000/1001", int64(1280), int64(720))

	va.SegmentTemplate = &SegmentTemplate{
		Timescale:      ptrs.Int64ptr(30000),
		Initialization: ptrs.Strptr("video/init.m4f"),
		Media:          ptrs.Strptr("video/segment$Number$.m4f"),
		SegmentTimeline: &SegmentTimeline{
			Segments: []*SegmentTimelineSegment{
				{StartTime: ptrs.Uint64ptr(0), Duration: 145145},
				{RepeatCount: ptrs.Intptr(2), Duration: 270270},
				{Duration: 91091},
				{Duration: 125125},
				{RepeatCount: ptrs.Intptr(1), Duration: 270270},
				{Duration: 88088},
			},
		},
	}
	return m
}
