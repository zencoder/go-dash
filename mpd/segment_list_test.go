package mpd

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zencoder/go-dash/helpers/ptrs"
	"github.com/zencoder/go-dash/helpers/testfixtures"
)

func TestSegmentListSerialization(t *testing.T) {
	m := getSegmentListMPD()
	xml, err := m.WriteToString()
	require.NoError(t, err)
	testfixtures.CompareFixture(t, "fixtures/segment_list.mpd", xml)
}

func TestSegmentListDeserialization(t *testing.T) {
	xml := testfixtures.LoadFixture("fixtures/segment_list.mpd")
	m, err := ReadFromString(xml)

	require.Nil(t, err)
	if err == nil {
		expected := getSegmentListMPD()

		require.Equal(t, expected.Periods[0].BaseURL, m.Periods[0].BaseURL)

		expectedAudioSegList := expected.Periods[0].AdaptationSets[0].Representations[0].SegmentList
		audioSegList := m.Periods[0].AdaptationSets[0].Representations[0].SegmentList

		require.Equal(t, expectedAudioSegList.Timescale, audioSegList.Timescale)
		require.Equal(t, expectedAudioSegList.Duration, audioSegList.Duration)
		require.Equal(t, expectedAudioSegList.Initialization, audioSegList.Initialization)

		for i := range expectedAudioSegList.SegmentURLs {
			require.Equal(t, expectedAudioSegList.SegmentURLs[i], audioSegList.SegmentURLs[i])
		}

		expectedVideoSegList := expected.Periods[0].AdaptationSets[1].Representations[0].SegmentList
		videoSegList := m.Periods[0].AdaptationSets[1].Representations[0].SegmentList

		require.Equal(t, expectedVideoSegList.Timescale, videoSegList.Timescale)
		require.Equal(t, expectedVideoSegList.Duration, videoSegList.Duration)
		require.Equal(t, expectedVideoSegList.Initialization, videoSegList.Initialization)

		for i := range expectedVideoSegList.SegmentURLs {
			require.Equal(t, expectedVideoSegList.SegmentURLs[i], videoSegList.SegmentURLs[i])
		}
	}
}

func getSegmentListMPD() *MPD {
	m := NewMPD(DASH_PROFILE_LIVE, "PT30.016S", "PT2.000S")
	m.period.BaseURL = "http://localhost:8002/dash/"

	aas, _ := m.AddNewAdaptationSetAudioWithID("1", "audio/mp4", true, 1, "English")
	ra, _ := aas.AddNewRepresentationAudio(48000, 255000, "mp4a.40.2", "audio_1")

	asl := new(SegmentList)
	asl.Timescale = ptrs.Uint32ptr(48000)
	asl.Duration = ptrs.Uint32ptr(479232)
	asl.Initialization = &URL{SourceURL: ptrs.Strptr("b4324d65-ad06-4735-9535-5cd4af84ebb6/dcb11457-9092-4410-b204-67b3c6d9a9e2/init.m4f")}

	asegs := []*SegmentURL{}
	asegs = append(asegs, &SegmentURL{Media: ptrs.Strptr("b4324d65-ad06-4735-9535-5cd4af84ebb6/dcb11457-9092-4410-b204-67b3c6d9a9e2/segment0.m4f")})
	asegs = append(asegs, &SegmentURL{Media: ptrs.Strptr("b4324d65-ad06-4735-9535-5cd4af84ebb6/dcb11457-9092-4410-b204-67b3c6d9a9e2/segment1.m4f")})
	asegs = append(asegs, &SegmentURL{Media: ptrs.Strptr("b4324d65-ad06-4735-9535-5cd4af84ebb6/dcb11457-9092-4410-b204-67b3c6d9a9e2/segment2.m4f")})
	asegs = append(asegs, &SegmentURL{Media: ptrs.Strptr("b4324d65-ad06-4735-9535-5cd4af84ebb6/dcb11457-9092-4410-b204-67b3c6d9a9e2/segment3.m4f")})
	asl.SegmentURLs = asegs

	ra.SegmentList = asl

	vas, _ := m.AddNewAdaptationSetVideoWithID("2", "video/mp4", "progressive", true, 1)
	va, _ := vas.AddNewRepresentationVideo(int64(4172274), "avc1.640028", "video_1", "30000/1001", int64(1280), int64(720))

	vsl := new(SegmentList)
	vsl.Timescale = ptrs.Uint32ptr(30000)
	vsl.Duration = ptrs.Uint32ptr(225120)
	vsl.Initialization = &URL{SourceURL: ptrs.Strptr("b4324d65-ad06-4735-9535-5cd4af84ebb6/f2ad47b2-5362-46e6-ad1d-dff7b10f00b8/init.m4f")}

	vsegs := []*SegmentURL{}
	vsegs = append(vsegs, &SegmentURL{Media: ptrs.Strptr("b4324d65-ad06-4735-9535-5cd4af84ebb6/f2ad47b2-5362-46e6-ad1d-dff7b10f00b8/segment0.m4f")})
	vsegs = append(vsegs, &SegmentURL{Media: ptrs.Strptr("b4324d65-ad06-4735-9535-5cd4af84ebb6/f2ad47b2-5362-46e6-ad1d-dff7b10f00b8/segment1.m4f")})
	vsegs = append(vsegs, &SegmentURL{Media: ptrs.Strptr("b4324d65-ad06-4735-9535-5cd4af84ebb6/f2ad47b2-5362-46e6-ad1d-dff7b10f00b8/segment2.m4f")})
	vsegs = append(vsegs, &SegmentURL{Media: ptrs.Strptr("b4324d65-ad06-4735-9535-5cd4af84ebb6/f2ad47b2-5362-46e6-ad1d-dff7b10f00b8/segment3.m4f")})
	vsl.SegmentURLs = vsegs

	va.SegmentList = vsl

	return m
}
