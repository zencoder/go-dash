package mpd

import (
	"testing"

	"github.com/stretchr/testify/suite"
	ptrs "github.com/zencoder/go-dash/helpers/ptrs"
	"github.com/zencoder/go-dash/helpers/testfixtures"
)

type SegmentTimelineSuite struct {
	suite.Suite
}

func TestSegmentTimelineSuite(t *testing.T) {
	suite.Run(t, new(SegmentTimelineSuite))
}

func (s *SegmentTimelineSuite) SetupTest() {

}

func (s *SegmentTimelineSuite) SetupSuite() {

}

func (s *SegmentTimelineSuite) TestSegmentTimelineSerialization() {
	expectedXML := testfixtures.LoadFixture("fixtures/segment_timeline.mpd")
	m := getSegmentTimelineMPD()
	xml, _ := m.WriteToString()
	s.Equal(expectedXML, xml)
}

func getSegmentTimelineMPD() *MPD {
	m := NewMPD(DASH_PROFILE_LIVE, "PT65.063S", "PT2.000S")
	m.Period.BaseURL = "http://localhost:8002/public/"

	aas, _ := m.AddNewAdaptationSetAudio("audio/mp4", true, 1, "English")
	ra, _ := aas.AddNewRepresentationAudio(48000, 255000, "mp4a.40.2", "audio_1")

	ast := &SegmentTemplate{
		Timescale:       ptrs.Int64ptr(48000),
		Initialization:  ptrs.Strptr("audio/init.m4f"),
		Media:           ptrs.Strptr("audio/segment$Number$.m4f"),
		SegmentTimeline: new(SegmentTimeline),
	}
	ra.SegmentTemplate = ast

	asegs := []*SegmentTimelineSegment{}
	asegs = append(asegs, &SegmentTimelineSegment{StartTime: ptrs.Uint64ptr(0), Duration: 231424})
	asegs = append(asegs, &SegmentTimelineSegment{RepeatCount: ptrs.Intptr(2), Duration: 479232})
	asegs = append(asegs, &SegmentTimelineSegment{Duration: 10240})
	asegs = append(asegs, &SegmentTimelineSegment{Duration: 247808})
	asegs = append(asegs, &SegmentTimelineSegment{RepeatCount: ptrs.Intptr(1), Duration: 479232})
	asegs = append(asegs, &SegmentTimelineSegment{Duration: 3072})
	ast.SegmentTimeline.Segments = asegs

	vas, _ := m.AddNewAdaptationSetVideo("video/mp4", "progressive", true, 1)
	va, _ := vas.AddNewRepresentationVideo(int64(4172274), "avc1.640028", "video_1", "30000/1001", int64(1280), int64(720))

	vst := &SegmentTemplate{
		Timescale:       ptrs.Int64ptr(30000),
		Initialization:  ptrs.Strptr("video/init.m4f"),
		Media:           ptrs.Strptr("video/segment$Number$.m4f"),
		SegmentTimeline: new(SegmentTimeline),
	}
	va.SegmentTemplate = vst

	vsegs := []*SegmentTimelineSegment{}
	vsegs = append(vsegs, &SegmentTimelineSegment{StartTime: ptrs.Uint64ptr(0), Duration: 145145})
	vsegs = append(vsegs, &SegmentTimelineSegment{RepeatCount: ptrs.Intptr(2), Duration: 270270})
	vsegs = append(vsegs, &SegmentTimelineSegment{Duration: 91091})
	vsegs = append(vsegs, &SegmentTimelineSegment{Duration: 125125})
	vsegs = append(vsegs, &SegmentTimelineSegment{RepeatCount: ptrs.Intptr(1), Duration: 270270})
	vsegs = append(vsegs, &SegmentTimelineSegment{Duration: 88088})
	vst.SegmentTimeline.Segments = vsegs

	return m
}
