package mpd

import (
	"encoding/base64"
	"testing"

	. "github.com/zencoder/go-dash/helpers/ptrs"
	"github.com/zencoder/go-dash/helpers/require"
	"github.com/zencoder/go-dash/helpers/testfixtures"
)

const (
	VALID_MEDIA_PRESENTATION_DURATION string = "PT6M16S"
	VALID_MIN_BUFFER_TIME             string = "PT1.97S"
	VALID_AVAILABILITY_START_TIME     string = "1970-01-01T00:00:00Z"
	VALID_MINIMUM_UPDATE_PERIOD       string = "PT5S"
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
	VALID_PLAYREADY_PRO               string = "BgIAAAEAAQD8ATwAVwBSAE0ASABFAEEARABFAFIAIAB4AG0AbABuAHMAPQAiAGgAdAB0AHAAOgAvAC8AcwBjAGgAZQBtAGEAcwAuAG0AaQBjAHIAbwBzAG8AZgB0AC4AYwBvAG0ALwBEAFIATQAvADIAMAAwADcALwAwADMALwBQAGwAYQB5AFIAZQBhAGQAeQBIAGUAYQBkAGUAcgAiACAAdgBlAHIAcwBpAG8AbgA9ACIANAAuADAALgAwAC4AMAAiAD4APABEAEEAVABBAD4APABQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsARQBZAEwARQBOAD4AMQA2ADwALwBLAEUAWQBMAEUATgA+ADwAQQBMAEcASQBEAD4AQQBFAFMAQwBUAFIAPAAvAEEATABHAEkARAA+ADwALwBQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsASQBEAD4ATAA5AFcAOQBXAGsAcABWAEsAawArADQAMABHAEgAMwBZAFUASgBSAFYAUQA9AD0APAAvAEsASQBEAD4APABDAEgARQBDAEsAUwBVAE0APgBJAEsAegBZADIASABaAEwAQQBsAEkAPQA8AC8AQwBIAEUAQwBLAFMAVQBNAD4APAAvAEQAQQBUAEEAPgA8AC8AVwBSAE0ASABFAEEARABFAFIAPgA="
	VALID_WV_HEADER                   string = "CAESEFq91S9VSk8quNBh92FCUVUaCGNhc3RsYWJzIhhXcjNWTDFWS1R5cTQwR0gzWVVKUlZRPT0yB2RlZmF1bHQ="
	VALID_SUBTITLE_BANDWIDTH          int64  = 256
	VALID_SUBTITLE_ID                 string = "subtitle_en"
	VALID_SUBTITLE_URL                string = "http://example.com/content/sintel/subtitles/subtitles_en.vtt"
	VALID_ROLE                        string = "main"
	VALID_LOCATION                    string = "https://example.com/location.mpd"
)

func TestNewMPDLive(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME,
		AttrAvailabilityStartTime(VALID_AVAILABILITY_START_TIME))
	require.NotNil(t, m)
	expectedMPD := &MPD{
		XMLNs:                     Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles:                  Strptr((string)(DASH_PROFILE_LIVE)),
		Type:                      Strptr("static"),
		MediaPresentationDuration: Strptr(VALID_MEDIA_PRESENTATION_DURATION),
		MinBufferTime:             Strptr(VALID_MIN_BUFFER_TIME),
		AvailabilityStartTime:     Strptr(VALID_AVAILABILITY_START_TIME),
		period:                    &Period{},
		Periods:                   []*Period{{}},
	}

	expectedString, err := expectedMPD.WriteToString()
	require.NoError(t, err)
	actualString, err := m.WriteToString()
	require.NoError(t, err)

	require.EqualString(t, expectedString, actualString)
}

func TestNewDynamicMPDLive(t *testing.T) {
	m := NewDynamicMPD(DASH_PROFILE_LIVE, VALID_AVAILABILITY_START_TIME, VALID_MIN_BUFFER_TIME,
		AttrMediaPresentationDuration(VALID_MEDIA_PRESENTATION_DURATION),
		AttrMinimumUpdatePeriod(VALID_MINIMUM_UPDATE_PERIOD))
	require.NotNil(t, m)
	expectedMPD := &MPD{
		XMLNs:                     Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles:                  Strptr((string)(DASH_PROFILE_LIVE)),
		Type:                      Strptr("dynamic"),
		MediaPresentationDuration: Strptr(VALID_MEDIA_PRESENTATION_DURATION),
		MinBufferTime:             Strptr(VALID_MIN_BUFFER_TIME),
		AvailabilityStartTime:     Strptr(VALID_AVAILABILITY_START_TIME),
		MinimumUpdatePeriod:       Strptr(VALID_MINIMUM_UPDATE_PERIOD),
		period:                    &Period{},
		Periods:                   []*Period{{}},
		UTCTiming:                 &DescriptorType{},
	}

	expectedString, err := expectedMPD.WriteToString()
	require.NoError(t, err)
	actualString, err := m.WriteToString()
	require.NoError(t, err)

	require.EqualString(t, expectedString, actualString)
}

func TestContentProtection_ImplementsInterface(t *testing.T) {
	cp := (*ContentProtectioner)(nil)
	require.Implements(t, cp, &ContentProtection{})
	require.Implements(t, cp, ContentProtection{})
}

func TestCENCContentProtection_ImplementsInterface(t *testing.T) {
	cp := (*ContentProtectioner)(nil)
	require.Implements(t, cp, &CENCContentProtection{})
	require.Implements(t, cp, CENCContentProtection{})
}

func TestPlayreadyContentProtection_ImplementsInterface(t *testing.T) {
	cp := (*ContentProtectioner)(nil)
	require.Implements(t, cp, &PlayreadyContentProtection{})
	require.Implements(t, cp, PlayreadyContentProtection{})
}

func TestWidevineContentProtection_ImplementsInterface(t *testing.T) {
	cp := (*ContentProtectioner)(nil)
	require.Implements(t, cp, &WidevineContentProtection{})
	require.Implements(t, cp, WidevineContentProtection{})
}

func TestNewMPDLiveWithBaseURLInMPD(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	m.BaseURL = VALID_BASE_URL_VIDEO
	require.NotNil(t, m)
	expectedMPD := &MPD{
		XMLNs:                     Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles:                  Strptr((string)(DASH_PROFILE_LIVE)),
		Type:                      Strptr("static"),
		MediaPresentationDuration: Strptr(VALID_MEDIA_PRESENTATION_DURATION),
		MinBufferTime:             Strptr(VALID_MIN_BUFFER_TIME),
		period:                    &Period{},
		Periods:                   []*Period{{}},
		BaseURL:                   VALID_BASE_URL_VIDEO,
	}

	expectedString, err := expectedMPD.WriteToString()
	require.NoError(t, err)
	actualString, err := m.WriteToString()
	require.NoError(t, err)

	require.EqualString(t, expectedString, actualString)
}

func TestNewMPDLiveWithBaseURLInPeriod(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	m.period.BaseURL = VALID_BASE_URL_VIDEO
	require.NotNil(t, m)
	period := &Period{
		BaseURL: VALID_BASE_URL_VIDEO,
	}
	expectedMPD := &MPD{
		XMLNs:                     Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles:                  Strptr((string)(DASH_PROFILE_LIVE)),
		Type:                      Strptr("static"),
		MediaPresentationDuration: Strptr(VALID_MEDIA_PRESENTATION_DURATION),
		MinBufferTime:             Strptr(VALID_MIN_BUFFER_TIME),
		period:                    period,
		Periods:                   []*Period{period},
	}

	expectedString, err := expectedMPD.WriteToString()
	require.NoError(t, err)
	actualString, err := m.WriteToString()
	require.NoError(t, err)

	require.EqualString(t, expectedString, actualString)
}

func TestNewMPDHbbTV(t *testing.T) {
	m := NewMPD(DASH_PROFILE_HBBTV_1_5_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	require.NotNil(t, m)
	expectedMPD := &MPD{
		XMLNs:                     Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles:                  Strptr((string)(DASH_PROFILE_HBBTV_1_5_LIVE)),
		Type:                      Strptr("static"),
		MediaPresentationDuration: Strptr(VALID_MEDIA_PRESENTATION_DURATION),
		MinBufferTime:             Strptr(VALID_MIN_BUFFER_TIME),
		period:                    &Period{},
		Periods:                   []*Period{{}},
	}

	expectedString, err := expectedMPD.WriteToString()
	require.NoError(t, err)
	actualString, err := m.WriteToString()
	require.NoError(t, err)

	require.EqualString(t, expectedString, actualString)
}

func TestNewMPDOnDemand(t *testing.T) {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	require.NotNil(t, m)
	expectedMPD := &MPD{
		XMLNs:                     Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles:                  Strptr((string)(DASH_PROFILE_ONDEMAND)),
		Type:                      Strptr("static"),
		MediaPresentationDuration: Strptr(VALID_MEDIA_PRESENTATION_DURATION),
		MinBufferTime:             Strptr(VALID_MIN_BUFFER_TIME),
		period:                    &Period{},
		Periods:                   []*Period{{}},
	}

	expectedString, err := expectedMPD.WriteToString()
	require.NoError(t, err)
	actualString, err := m.WriteToString()
	require.NoError(t, err)

	require.EqualString(t, expectedString, actualString)
}

func TestAddAdaptationSetErrorNil(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	err := m.period.addAdaptationSet(nil)
	require.NotNil(t, err)
	require.EqualErr(t, ErrAdaptationSetNil, err)
}

type TestProprietaryContentProtection struct {
	ContentProtectionMarshal
	TestAttrA string `xml:"a,attr,omitempty"`
	TestAttrB string `xml:"b,attr,omitempty"`
}

func (s *TestProprietaryContentProtection) ContentProtected() {}

func TestAddNewContentProtectionRootErrorInvalidLengthDefaultKID(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	s, _ := m.AddNewAdaptationSetVideoWithID("7357", DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	cp, err := s.AddNewContentProtectionRoot("invalidkid")
	require.NotNil(t, err)
	require.EqualErr(t, ErrInvalidDefaultKID, err)
	require.Nil(t, cp)
}

func TestAddNewContentProtectionRootErrorEmptyDefaultKID(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	s, _ := m.AddNewAdaptationSetVideoWithID("7357", DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	cp, err := s.AddNewContentProtectionRoot("")
	require.NotNil(t, err)
	require.EqualErr(t, ErrInvalidDefaultKID, err)
	require.Nil(t, cp)
}

func TestAddNewContentProtectionSchemePlayreadyErrorEmptyPRO(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	s, _ := m.AddNewAdaptationSetVideoWithID("7357", DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	cp, err := s.AddNewContentProtectionSchemePlayready("")
	require.NotNil(t, err)
	require.EqualErr(t, ErrPROEmpty, err)
	require.Nil(t, cp)
}

func TestAddNewContentProtectionSchemePlayreadyV10ErrorEmptyPRO(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	s, _ := m.AddNewAdaptationSetVideoWithID("7357", DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	cp, err := s.AddNewContentProtectionSchemePlayreadyV10("")
	require.NotNil(t, err)
	require.EqualErr(t, ErrPROEmpty, err)
	require.Nil(t, cp)
}

func TestSetNewSegmentTemplate(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	audioAS, _ := m.AddNewAdaptationSetAudioWithID("7357", DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)
	st, err := audioAS.SetNewSegmentTemplate(VALID_DURATION, VALID_INIT_PATH_AUDIO, VALID_MEDIA_PATH_AUDIO, VALID_START_NUMBER, VALID_TIMESCALE)
	require.NotNil(t, st)
	require.NoError(t, err)
}

func TestSetNewSegmentTemplateErrorNoDASHProfile(t *testing.T) {
	m := &MPD{
		XMLNs:                     Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles:                  nil,
		Type:                      Strptr("static"),
		MediaPresentationDuration: Strptr(VALID_MEDIA_PRESENTATION_DURATION),
		MinBufferTime:             Strptr(VALID_MIN_BUFFER_TIME),
		period:                    &Period{},
	}
	audioAS, _ := m.AddNewAdaptationSetAudioWithID("7357", DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)
	_, _ = audioAS.SetNewSegmentTemplate(VALID_DURATION, VALID_INIT_PATH_AUDIO, VALID_MEDIA_PATH_AUDIO, VALID_START_NUMBER, VALID_TIMESCALE)
	err := m.Validate()
	require.EqualErr(t, ErrNoDASHProfileSet, err)
}

func TestAddRepresentationAudio(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	audioAS, _ := m.AddNewAdaptationSetAudioWithID("7357", DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)

	r, err := audioAS.AddNewRepresentationAudio(VALID_AUDIO_SAMPLE_RATE, VALID_AUDIO_BITRATE, VALID_AUDIO_CODEC, VALID_AUDIO_ID)

	require.NotNil(t, r)
	require.NoError(t, err)
}

func TestAddAudioChannelConfiguration(t *testing.T) {
	m := NewMPD(DASH_PROFILE_HBBTV_1_5_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	audioAS, _ := m.AddNewAdaptationSetAudioWithID("7357", DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)

	r, _ := audioAS.AddNewRepresentationAudio(VALID_AUDIO_SAMPLE_RATE, VALID_AUDIO_BITRATE, VALID_AUDIO_CODEC, VALID_AUDIO_ID)

	acc, err := r.AddNewAudioChannelConfiguration(AUDIO_CHANNEL_CONFIGURATION_MPEG_DASH, "2")

	require.NotNil(t, acc)
	require.NoError(t, err)
}

func TestAddRepresentationVideo(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	videoAS, _ := m.AddNewAdaptationSetVideoWithID("7357", DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, err := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	require.NotNil(t, r)
	require.NoError(t, err)
}

func TestAddRepresentationSubtitle(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	subtitleAS, _ := m.AddNewAdaptationSetSubtitleWithID("7357", DASH_MIME_TYPE_SUBTITLE_VTT, VALID_LANG)

	r, err := subtitleAS.AddNewRepresentationSubtitle(VALID_SUBTITLE_BANDWIDTH, VALID_SUBTITLE_ID)

	require.NotNil(t, r)
	require.NoError(t, err)
}

func TestAddRepresentationErrorNil(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	videoAS, _ := m.AddNewAdaptationSetVideoWithID("7357", DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	err := videoAS.addRepresentation(nil)
	require.NotNil(t, err)
	require.EqualErr(t, ErrRepresentationNil, err)
}

func TestAddRole(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	audioAS, _ := m.AddNewAdaptationSetAudioWithID("7357", DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)

	r, err := audioAS.AddNewRole("urn:mpeg:dash:role:2011", VALID_ROLE)

	require.NotNil(t, r)
	require.NoError(t, err)
}

func TestSetSegmentTemplateErrorNil(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	audioAS, _ := m.AddNewAdaptationSetAudioWithID("7357", DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)
	err := audioAS.setSegmentTemplate(nil)
	require.NotNil(t, err)
	require.EqualErr(t, ErrSegmentTemplateNil, err)
}

func TestSetNewBaseURLVideo(t *testing.T) {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	videoAS, _ := m.AddNewAdaptationSetVideoWithID("7357", DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, _ := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	err := r.SetNewBaseURL(VALID_BASE_URL_VIDEO)

	require.NoError(t, err)
}

func TestSetNewBaseURLSubtitle(t *testing.T) {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	subtitleAS, _ := m.AddNewAdaptationSetSubtitleWithID("7357", DASH_MIME_TYPE_SUBTITLE_VTT, VALID_LANG)

	r, _ := subtitleAS.AddNewRepresentationSubtitle(VALID_SUBTITLE_BANDWIDTH, VALID_SUBTITLE_ID)

	err := r.SetNewBaseURL(VALID_SUBTITLE_URL)

	require.NoError(t, err)
}

func TestSetNewBaseURLErrorNoDASHProfile(t *testing.T) {
	m := &MPD{
		XMLNs:                     Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles:                  nil,
		Type:                      Strptr("static"),
		MediaPresentationDuration: Strptr(VALID_MEDIA_PRESENTATION_DURATION),
		MinBufferTime:             Strptr(VALID_MIN_BUFFER_TIME),
		period:                    &Period{},
	}
	videoAS, _ := m.AddNewAdaptationSetVideoWithID("7357", DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, _ := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	_ = r.SetNewBaseURL(VALID_BASE_URL_VIDEO)
	err := m.Validate()

	require.NotNil(t, err)
	require.EqualErr(t, ErrNoDASHProfileSet, err)
}

func TestSetNewBaseURLErrorEmpty(t *testing.T) {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	videoAS, _ := m.AddNewAdaptationSetVideoWithID("7357", DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, _ := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	err := r.SetNewBaseURL("")

	require.NotNil(t, err)
	require.EqualErr(t, ErrBaseURLEmpty, err)
}

func TestSetNewSegmentBase(t *testing.T) {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	videoAS, _ := m.AddNewAdaptationSetVideoWithID("7357", DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, _ := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	sb, err := r.AddNewSegmentBase(VALID_INDEX_RANGE, VALID_INIT_RANGE)
	require.NotNil(t, sb)
	require.NoError(t, err)
}

func TestSetNewSegmentBaseErrorNoDASHProfile(t *testing.T) {
	m := &MPD{
		XMLNs:                     Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles:                  nil,
		Type:                      Strptr("static"),
		MediaPresentationDuration: Strptr(VALID_MEDIA_PRESENTATION_DURATION),
		MinBufferTime:             Strptr(VALID_MIN_BUFFER_TIME),
		period:                    &Period{},
	}
	videoAS, _ := m.AddNewAdaptationSetVideoWithID("7357", DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, _ := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	_, _ = r.AddNewSegmentBase(VALID_INDEX_RANGE, VALID_INIT_RANGE)

	err := m.Validate()
	require.EqualErr(t, ErrNoDASHProfileSet, err)
}

func TestSetSegmentBaseErrorNil(t *testing.T) {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	videoAS, _ := m.AddNewAdaptationSetVideoWithID("7357", DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, _ := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	err := r.setSegmentBase(nil)
	require.NotNil(t, err)
	require.EqualErr(t, ErrSegmentBaseNil, err)
}

func getValidWVHeaderBytes() []byte {
	wvHeader, err := base64.StdEncoding.DecodeString(VALID_WV_HEADER)
	if err != nil {
		panic(err.Error())
	}
	return wvHeader
}

func TestAddNewAccessibilityElement(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	audioAS, err := m.AddNewAdaptationSetAudioWithID("7357", DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT,
		VALID_START_WITH_SAP, VALID_LANG)
	if err != nil {
		t.Errorf("AddNewAccessibilityElement() error adding audio adaptation set: %v", err)
		return
	}

	_, err = audioAS.AddNewAccessibilityElement(ACCESSIBILITY_ELEMENT_SCHEME_DESCRIPTIVE_AUDIO, "1")
	if err != nil {
		t.Errorf("AddNewAccessibilityElement() error adding accessibility element: %v", err)
		return
	}

	if g, e := len(audioAS.AccessibilityElems), 1; g != e {
		t.Errorf("AddNewAccessibilityElement() wrong number of accessibility elements, got: %d, expected: %d",
			g, e)
		return
	}

	elem := audioAS.AccessibilityElems[0]

	require.EqualStringPtr(t, Strptr((string)(ACCESSIBILITY_ELEMENT_SCHEME_DESCRIPTIVE_AUDIO)), elem.SchemeIdUri)
	require.EqualStringPtr(t, Strptr("1"), elem.Value)
}

func TestLocationWriteToString(t *testing.T) {
	m := &MPD{
		XMLNs:                 Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles:              Strptr((string)(DASH_PROFILE_LIVE)),
		Type:                  Strptr("dynamic"),
		AvailabilityStartTime: Strptr(VALID_AVAILABILITY_START_TIME),
		MinimumUpdatePeriod:   Strptr(VALID_MINIMUM_UPDATE_PERIOD),
		PublishTime:           Strptr(VALID_AVAILABILITY_START_TIME),
		Location:              VALID_LOCATION,
	}

	got, err := m.WriteToString()
	require.NoError(t, err)

	testfixtures.CompareFixture(t, "fixtures/location.mpd", got)
}

func TestReadLocation(t *testing.T) {
	m, err := ReadFromFile("fixtures/location.mpd")
	require.NoError(t, err)

	got, err := m.WriteToString()
	require.NoError(t, err)

	testfixtures.CompareFixture(t, "fixtures/location.mpd", got)
}
