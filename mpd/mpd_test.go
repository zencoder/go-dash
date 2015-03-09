package mpd

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	. "github.com/zencoder/go-dash/helpers/ptrs"
)

type MPDSuite struct {
	suite.Suite
}

func TestMPDSuite(t *testing.T) {
	suite.Run(t, new(MPDSuite))
}

func (s *MPDSuite) SetupTest() {

}

func (s *MPDSuite) SetupSuite() {

}

const (
	VALID_MEDIA_PRESENTATION_DURATION string = "PT6M16S"
	VALID_MIN_BUFFER_TIME             string = "PT1.97S"
	VALID_MIME_TYPE_VIDEO             string = "video/mp4"
	VALID_MIME_TYPE_AUDIO             string = "audio/mp4"
	VALID_SCAN_TYPE                   string = "progressive"
	VALID_SEGMENT_ALIGNMENT           bool   = true
	VALID_START_WITH_SAP              int64  = 1
	VALID_LANG                        string = "en"
	VALID_DURATION                    int64  = 1968
	VALID_INIT_PATH_AUDIO             string = "$RepresentationID$/audio/en/init.mp4"
	VALID_MEDIA_PATH_AUDIO            string = "$RepresentationID$/audio/en/seg-$Number$.m4f"
	VALID_START_NUMBER                int64  = 0
	VALID_TIMESCALE                   int64  = 0
)

func (s *MPDSuite) TestReadMPDLiveProfile() {
	m, err := ReadFromFile("fixtures/live_profile.mpd")
	assert.NotNil(s.T(), m)
	assert.Nil(s.T(), err)
	//assert.Equal(s.T(), &MPD{}, m)
}

func (s *MPDSuite) TestReadMPDOnDemandProfile() {
	m, err := ReadFromFile("fixtures/ondemand_profile.mpd")
	assert.NotNil(s.T(), m)
	assert.Nil(s.T(), err)
	//assert.Equal(s.T(), &MPD{}, m)
}

func (s *MPDSuite) TestNewMPDLive() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	assert.NotNil(s.T(), m)
	expectedMPD := &MPD{
		XMLNs:    Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles: Strptr((string)(DASH_PROFILE_LIVE)),
		Type:     Strptr("static"),
		MediaPresentationDuration: Strptr(VALID_MEDIA_PRESENTATION_DURATION),
		MinBufferTime:             Strptr(VALID_MIN_BUFFER_TIME),
		Period:                    &Period{},
	}
	assert.Equal(s.T(), expectedMPD, m)
}

func (s *MPDSuite) TestNewMPDLiveWriteToString() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	xmlStr, err := m.WriteToString()
	assert.Nil(s.T(), err)
	expectedXML := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<MPD xmlns=\"urn:mpeg:dash:schema:mpd:2011\" profiles=\"urn:mpeg:dash:profile:isoff-live:2011\" type=\"static\" mediaPresentationDuration=\"PT6M16S\" minBufferTime=\"PT1.97S\">\n  <Period></Period>\n</MPD>\n"
	assert.Equal(s.T(), expectedXML, xmlStr)
}

func (s *MPDSuite) TestNewMPDOnDemand() {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	assert.NotNil(s.T(), m)
	expectedMPD := &MPD{
		XMLNs:    Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles: Strptr((string)(DASH_PROFILE_ONDEMAND)),
		Type:     Strptr("static"),
		MediaPresentationDuration: Strptr(VALID_MEDIA_PRESENTATION_DURATION),
		MinBufferTime:             Strptr(VALID_MIN_BUFFER_TIME),
		Period:                    &Period{},
	}
	assert.Equal(s.T(), expectedMPD, m)
}

func (s *MPDSuite) TestNewMPDOnDemandWriteToString() {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	xmlStr, err := m.WriteToString()
	assert.Nil(s.T(), err)
	expectedXML := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<MPD xmlns=\"urn:mpeg:dash:schema:mpd:2011\" profiles=\"urn:mpeg:dash:profile:isoff-on-demand:2011\" type=\"static\" mediaPresentationDuration=\"PT6M16S\" minBufferTime=\"PT1.97S\">\n  <Period></Period>\n</MPD>\n"
	assert.Equal(s.T(), expectedXML, xmlStr)
}

func (s *MPDSuite) TestAddNewAdaptationSetAudio() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	as, err := m.AddNewAdaptationSetAudio(VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)
	assert.NotNil(s.T(), as)
	assert.Nil(s.T(), err)
	expectedAS := &AdaptationSet{
		MPD:              m,
		MimeType:         Strptr(VALID_MIME_TYPE_AUDIO),
		SegmentAlignment: Boolptr(VALID_SEGMENT_ALIGNMENT),
		StartWithSAP:     Intptr(VALID_START_WITH_SAP),
		Lang:             Strptr(VALID_LANG),
	}
	assert.Equal(s.T(), expectedAS, as)
}

func (s *MPDSuite) TestAddNewAdaptationSetAudioWriteToString() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	m.AddNewAdaptationSetAudio(VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)

	xmlStr, err := m.WriteToString()
	assert.Nil(s.T(), err)
	expectedXML := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<MPD xmlns=\"urn:mpeg:dash:schema:mpd:2011\" profiles=\"urn:mpeg:dash:profile:isoff-live:2011\" type=\"static\" mediaPresentationDuration=\"PT6M16S\" minBufferTime=\"PT1.97S\">\n  <Period>\n    <AdaptationSet mimeType=\"audio/mp4\" segmentAlignment=\"true\" startWithSAP=\"1\" lang=\"en\"></AdaptationSet>\n  </Period>\n</MPD>\n"
	assert.Equal(s.T(), expectedXML, xmlStr)
}

func (s *MPDSuite) TestAddNewAdaptationSetVideo() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	as, err := m.AddNewAdaptationSetVideo(VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)
	assert.NotNil(s.T(), as)
	assert.Nil(s.T(), err)
	expectedAS := &AdaptationSet{
		MPD:              m,
		MimeType:         Strptr(VALID_MIME_TYPE_VIDEO),
		ScanType:         Strptr(VALID_SCAN_TYPE),
		SegmentAlignment: Boolptr(VALID_SEGMENT_ALIGNMENT),
		StartWithSAP:     Intptr(VALID_START_WITH_SAP),
	}
	assert.Equal(s.T(), expectedAS, as)
}

func (s *MPDSuite) TestAddNewAdaptationSetVideoWriteToString() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	m.AddNewAdaptationSetVideo(VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	xmlStr, err := m.WriteToString()
	assert.Nil(s.T(), err)
	expectedXML := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<MPD xmlns=\"urn:mpeg:dash:schema:mpd:2011\" profiles=\"urn:mpeg:dash:profile:isoff-live:2011\" type=\"static\" mediaPresentationDuration=\"PT6M16S\" minBufferTime=\"PT1.97S\">\n  <Period>\n    <AdaptationSet mimeType=\"video/mp4\" scanType=\"progressive\" segmentAlignment=\"true\" startWithSAP=\"1\"></AdaptationSet>\n  </Period>\n</MPD>\n"
	assert.Equal(s.T(), expectedXML, xmlStr)
}

func (s *MPDSuite) TestSetNewSegmentTemplate() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	audioAS, _ := m.AddNewAdaptationSetAudio(VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)
	st, err := audioAS.SetNewSegmentTemplate(VALID_DURATION, VALID_INIT_PATH_AUDIO, VALID_MEDIA_PATH_AUDIO, VALID_START_NUMBER, VALID_TIMESCALE)
	assert.NotNil(s.T(), st)
	assert.Nil(s.T(), err)
}

func (s *MPDSuite) TestSetNewSegmentTemplateErrorInvalidProfile() {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	audioAS, _ := m.AddNewAdaptationSetAudio(VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)
	st, err := audioAS.SetNewSegmentTemplate(VALID_DURATION, VALID_INIT_PATH_AUDIO, VALID_MEDIA_PATH_AUDIO, VALID_START_NUMBER, VALID_TIMESCALE)
	assert.Nil(s.T(), st)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrSegmentTemplateLiveProfileOnly, err)
}

func (s *MPDSuite) TestFullLiveProfileWriteToString() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	audioAS, _ := m.AddNewAdaptationSetAudio(VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)
	audioAS.SetNewSegmentTemplate(1968, "$RepresentationID$/audio/en/init.mp4", "$RepresentationID$/audio/en/seg-$Number$.m4f", 0, 1000)
	audioAS.AddNewRepresentationAudio(44100, 67095, "mp4a.40.2", "800")

	videoAS, _ := m.AddNewAdaptationSetVideo(VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)
	videoAS.SetNewSegmentTemplate(1968, "$RepresentationID$/video/1/init.mp4", "$RepresentationID$/video/1/seg-$Number$.m4f", 0, 1000)
	videoAS.AddNewRepresentationVideo(1518664, "avc1.4d401f", "800", "30000/1001", 960, 540)
	videoAS.AddNewRepresentationVideo(1911775, "avc1.4d401f", "1000", "30000/1001", 1024, 576)
	videoAS.AddNewRepresentationVideo(2295158, "avc1.4d401f", "1200", "30000/1001", 1024, 576)
	videoAS.AddNewRepresentationVideo(2780732, "avc1.4d401f", "1500", "30000/1001", 1280, 720)

	xmlStr, err := m.WriteToString()
	assert.Nil(s.T(), err)
	expectedXML, err := ioutil.ReadFile("fixtures/live_profile.mpd")
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), (string)(expectedXML), xmlStr)
}

func (s *MPDSuite) TestFullOnDemandProfileWriteToString() {
	m := NewMPD(DASH_PROFILE_ONDEMAND, "PT30S", VALID_MIN_BUFFER_TIME)

	audioAS, _ := m.AddNewAdaptationSetAudio(VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, "und")
	audioRep, _ := audioAS.AddNewRepresentationAudio(44100, 128558, "mp4a.40.5", "800k/audio-und")
	audioRep.SetNewBaseURL("800k/output-audio-und.mp4")
	audioRep.AddNewSegmentBase("629-756", "0-628")

	videoAS, _ := m.AddNewAdaptationSetVideo(VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)
	videoRep1, _ := videoAS.AddNewRepresentationVideo(1100690, "avc1.4d401e", "800k/video-1", "30000/1001", 640, 360)
	videoRep1.SetNewBaseURL("800k/output-video-1.mp4")
	videoRep1.AddNewSegmentBase("686-813", "0-685")

	videoRep2, _ := videoAS.AddNewRepresentationVideo(1633516, "avc1.4d401f", "1200k/video-1", "30000/1001", 960, 540)
	videoRep2.SetNewBaseURL("1200k/output-video-1.mp4")
	videoRep2.AddNewSegmentBase("686-813", "0-685")

	xmlStr, err := m.WriteToString()
	assert.Nil(s.T(), err)
	expectedXML, err := ioutil.ReadFile("fixtures/ondemand_profile.mpd")
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), (string)(expectedXML), xmlStr)
}
