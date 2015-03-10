package mpd

import (
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
	VALID_TIMESCALE                   int64  = 1000
	VALID_AUDIO_SAMPLE_RATE           int64  = 44100
	VALID_AUDIO_BITRATE               int64  = 67095
	VALID_AUDIO_CODEC                 string = "mp4a.40.2"
	VALID_AUDIO_ID                    string = "800"
	VALID_VIDEO_BITRATE               int64  = 1518664
	VALID_VIDEO_CODEC                 string = "avc1.4d401f"
	VALID_VIDEO_ID                    string = "800"
	VALID_VIDEO_FRAMERATE             string = "30000/1001"
	VALID_VIDEO_WIDTH                 int64  = 960
	VALID_VIDEO_HEIGHT                int64  = 540
	VALID_BASE_URL_VIDEO              string = "800k/output-video-1.mp4"
	VALID_INDEX_RANGE                 string = "629-756"
	VALID_INIT_RANGE                  string = "0-628"
)

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

func (s *MPDSuite) TestAddAdaptationSetErrorNil() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	err := m.AddAdaptationSet(nil)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrAdaptationSetNil, err)
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

func (s *MPDSuite) TestSetNewSegmentTemplateErrorNoDASHProfile() {
	m := &MPD{
		XMLNs:    Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles: nil,
		Type:     Strptr("static"),
		MediaPresentationDuration: Strptr(VALID_MEDIA_PRESENTATION_DURATION),
		MinBufferTime:             Strptr(VALID_MIN_BUFFER_TIME),
		Period:                    &Period{},
	}
	audioAS, _ := m.AddNewAdaptationSetAudio(VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)
	st, err := audioAS.SetNewSegmentTemplate(VALID_DURATION, VALID_INIT_PATH_AUDIO, VALID_MEDIA_PATH_AUDIO, VALID_START_NUMBER, VALID_TIMESCALE)
	assert.Nil(s.T(), st)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrNoDASHProfileSet, err)
}

func (s *MPDSuite) TestAddRepresentationAudio() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	audioAS, _ := m.AddNewAdaptationSetAudio(VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)

	r, err := audioAS.AddNewRepresentationAudio(VALID_AUDIO_SAMPLE_RATE, VALID_AUDIO_BITRATE, VALID_AUDIO_CODEC, VALID_AUDIO_ID)

	assert.NotNil(s.T(), r)
	assert.Nil(s.T(), err)
}

func (s *MPDSuite) TestAddRepresentationVideo() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	videoAS, _ := m.AddNewAdaptationSetVideo(VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, err := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	assert.NotNil(s.T(), r)
	assert.Nil(s.T(), err)
}

func (s *MPDSuite) TestAddRepresentationErrorNil() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	videoAS, _ := m.AddNewAdaptationSetVideo(VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	err := videoAS.AddRepresentation(nil)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrRepresentationNil, err)
}

func (s *MPDSuite) TestSetSegmentTemplateErrorNil() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	audioAS, _ := m.AddNewAdaptationSetAudio(VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)
	err := audioAS.SetSegmentTemplate(nil)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrSegmentTemplateNil, err)
}

func (s *MPDSuite) TestSetNewBaseURL() {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	videoAS, _ := m.AddNewAdaptationSetVideo(VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, _ := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	err := r.SetNewBaseURL(VALID_BASE_URL_VIDEO)

	assert.Nil(s.T(), err)
}

func (s *MPDSuite) TestSetNewBaseURLErrorNoDASHProfile() {
	m := &MPD{
		XMLNs:    Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles: nil,
		Type:     Strptr("static"),
		MediaPresentationDuration: Strptr(VALID_MEDIA_PRESENTATION_DURATION),
		MinBufferTime:             Strptr(VALID_MIN_BUFFER_TIME),
		Period:                    &Period{},
	}
	videoAS, _ := m.AddNewAdaptationSetVideo(VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, _ := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	err := r.SetNewBaseURL(VALID_BASE_URL_VIDEO)

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrNoDASHProfileSet, err)
}

func (s *MPDSuite) TestSetNewBaseURLErrorInvalidDASHProfile() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	videoAS, _ := m.AddNewAdaptationSetVideo(VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, _ := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	err := r.SetNewBaseURL(VALID_BASE_URL_VIDEO)

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrBaseURLOnDemandProfileOnly, err)
}

func (s *MPDSuite) TestSetNewBaseURLErrorEmpty() {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	videoAS, _ := m.AddNewAdaptationSetVideo(VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, _ := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	err := r.SetNewBaseURL("")

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrBaseURLEmpty, err)
}

func (s *MPDSuite) TestSetNewSegmentBase() {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	videoAS, _ := m.AddNewAdaptationSetVideo(VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, _ := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	sb, err := r.AddNewSegmentBase(VALID_INDEX_RANGE, VALID_INIT_RANGE)
	assert.NotNil(s.T(), sb)
	assert.Nil(s.T(), err)
}

func (s *MPDSuite) TestSetNewSegmentBaseErrorInvalidDASHProfile() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	videoAS, _ := m.AddNewAdaptationSetVideo(VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, _ := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	sb, err := r.AddNewSegmentBase(VALID_INDEX_RANGE, VALID_INIT_RANGE)
	assert.Nil(s.T(), sb)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrSegmentBaseOnDemandProfileOnly, err)
}

func (s *MPDSuite) TestSetNewSegmentBaseErrorNoDASHProfile() {
	m := &MPD{
		XMLNs:    Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles: nil,
		Type:     Strptr("static"),
		MediaPresentationDuration: Strptr(VALID_MEDIA_PRESENTATION_DURATION),
		MinBufferTime:             Strptr(VALID_MIN_BUFFER_TIME),
		Period:                    &Period{},
	}
	videoAS, _ := m.AddNewAdaptationSetVideo(VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, _ := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	sb, err := r.AddNewSegmentBase(VALID_INDEX_RANGE, VALID_INIT_RANGE)
	assert.Nil(s.T(), sb)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrNoDASHProfileSet, err)
}

func (s *MPDSuite) TestSetSegmentBaseErrorNil() {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	videoAS, _ := m.AddNewAdaptationSetVideo(VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, _ := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	err := r.SetSegmentBase(nil)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrSegmentBaseNil, err)
}
