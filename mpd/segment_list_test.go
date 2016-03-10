package mpd

import (
	"testing"

	"github.com/stretchr/testify/suite"
	ptrs "github.com/zencoder/go-dash/helpers/ptrs"
	"github.com/zencoder/go-dash/helpers/testfixtures"
)

type SegmentListSuite struct {
	suite.Suite
}

func TestSegmentListSuite(t *testing.T) {
	suite.Run(t, new(SegmentListSuite))
}

func (s *SegmentListSuite) SetupTest() {

}

func (s *SegmentListSuite) SetupSuite() {

}

func (s *SegmentListSuite) TestSegmentListSerialization() {
	expectedXML := testfixtures.LoadFixture("fixtures/segment_list.mpd")
	m := getSegmentListMPD()
	xml, _ := m.WriteToString()
	s.Equal(expectedXML, xml)
}

func getSegmentListMPD() *MPD {
	m := NewMPD(DASH_PROFILE_LIVE, "PT30.016S", "PT2.000S")
	m.Period.BaseURL = "http://localhost:8002/dash/"

	aas, _ := m.AddNewAdaptationSetAudio("audio/mp4", true, 1, "English")
	ra, _ := aas.AddNewRepresentationAudio(48000, 255000, "mp4a.40.2", "audio_1")

	asl := new(SegmentList)
	asl.Timescale = uint32ptr(48000)
	asl.Duration = uint32ptr(479232)
	asl.Initialization = &URL{SourceURL: ptrs.Strptr("b4324d65-ad06-4735-9535-5cd4af84ebb6/dcb11457-9092-4410-b204-67b3c6d9a9e2/init.m4f")}

	asegs := []*SegmentURL{}
	asegs = append(asegs, &SegmentURL{Media: ptrs.Strptr("b4324d65-ad06-4735-9535-5cd4af84ebb6/dcb11457-9092-4410-b204-67b3c6d9a9e2/segment0.m4f")})
	asegs = append(asegs, &SegmentURL{Media: ptrs.Strptr("b4324d65-ad06-4735-9535-5cd4af84ebb6/dcb11457-9092-4410-b204-67b3c6d9a9e2/segment1.m4f")})
	asegs = append(asegs, &SegmentURL{Media: ptrs.Strptr("b4324d65-ad06-4735-9535-5cd4af84ebb6/dcb11457-9092-4410-b204-67b3c6d9a9e2/segment2.m4f")})
	asegs = append(asegs, &SegmentURL{Media: ptrs.Strptr("b4324d65-ad06-4735-9535-5cd4af84ebb6/dcb11457-9092-4410-b204-67b3c6d9a9e2/segment3.m4f")})
	asl.SegmentURL = asegs

	ra.SegmentList = asl

	vas, _ := m.AddNewAdaptationSetVideo("video/mp4", "progressive", true, 1)
	va, _ := vas.AddNewRepresentationVideo(int64(4172274), "avc1.640028", "video_1", "30000/1001", int64(1280), int64(720))

	vsl := new(SegmentList)
	vsl.Timescale = uint32ptr(30000)
	vsl.Duration = uint32ptr(225120)
	vsl.Initialization = &URL{SourceURL: ptrs.Strptr("b4324d65-ad06-4735-9535-5cd4af84ebb6/f2ad47b2-5362-46e6-ad1d-dff7b10f00b8/init.m4f")}

	vsegs := []*SegmentURL{}
	vsegs = append(vsegs, &SegmentURL{Media: ptrs.Strptr("b4324d65-ad06-4735-9535-5cd4af84ebb6/f2ad47b2-5362-46e6-ad1d-dff7b10f00b8/segment0.m4f")})
	vsegs = append(vsegs, &SegmentURL{Media: ptrs.Strptr("b4324d65-ad06-4735-9535-5cd4af84ebb6/f2ad47b2-5362-46e6-ad1d-dff7b10f00b8/segment1.m4f")})
	vsegs = append(vsegs, &SegmentURL{Media: ptrs.Strptr("b4324d65-ad06-4735-9535-5cd4af84ebb6/f2ad47b2-5362-46e6-ad1d-dff7b10f00b8/segment2.m4f")})
	vsegs = append(vsegs, &SegmentURL{Media: ptrs.Strptr("b4324d65-ad06-4735-9535-5cd4af84ebb6/f2ad47b2-5362-46e6-ad1d-dff7b10f00b8/segment3.m4f")})
	vsl.SegmentURL = vsegs

	va.SegmentList = vsl

	return m
}
