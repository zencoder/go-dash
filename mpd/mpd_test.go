package mpd

import (
	"encoding/base64"
	"encoding/xml"
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
	VALID_MIME_TYPE_SUBTITLE_VTT      string = "text/vtt"
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
	VALID_DEFAULT_KID_HEX             string = "08e367028f33436ca5dd60ffe5571e60"
	VALID_DEFAULT_KID                 string = "08e36702-8f33-436c-a5dd60ffe5571e60"
	VALID_PLAYREADY_XMLNS             string = "urn:microsoft:playready"
	VALID_PLAYREADY_PRO               string = "BgIAAAEAAQD8ATwAVwBSAE0ASABFAEEARABFAFIAIAB4AG0AbABuAHMAPQAiAGgAdAB0AHAAOgAvAC8AcwBjAGgAZQBtAGEAcwAuAG0AaQBjAHIAbwBzAG8AZgB0AC4AYwBvAG0ALwBEAFIATQAvADIAMAAwADcALwAwADMALwBQAGwAYQB5AFIAZQBhAGQAeQBIAGUAYQBkAGUAcgAiACAAdgBlAHIAcwBpAG8AbgA9ACIANAAuADAALgAwAC4AMAAiAD4APABEAEEAVABBAD4APABQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsARQBZAEwARQBOAD4AMQA2ADwALwBLAEUAWQBMAEUATgA+ADwAQQBMAEcASQBEAD4AQQBFAFMAQwBUAFIAPAAvAEEATABHAEkARAA+ADwALwBQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsASQBEAD4ATAA5AFcAOQBXAGsAcABWAEsAawArADQAMABHAEgAMwBZAFUASgBSAFYAUQA9AD0APAAvAEsASQBEAD4APABDAEgARQBDAEsAUwBVAE0APgBJAEsAegBZADIASABaAEwAQQBsAEkAPQA8AC8AQwBIAEUAQwBLAFMAVQBNAD4APAAvAEQAQQBUAEEAPgA8AC8AVwBSAE0ASABFAEEARABFAFIAPgA="
	VALID_WV_HEADER                   string = "CAESEFq91S9VSk8quNBh92FCUVUaCGNhc3RsYWJzIhhXcjNWTDFWS1R5cTQwR0gzWVVKUlZRPT0yB2RlZmF1bHQ="
	VALID_SUBTITLE_BANDWIDTH          int64  = 256
	VALID_SUBTITLE_ID                 string = "subtitle_en"
	VALID_SUBTITLE_URL                string = "http://example.com/content/sintel/subtitles/subtitles_en.vtt"
	VALID_ROLE                        string = "main"
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

func (s *MPDSuite) TestContentProtection_ImplementsInterface() {
	cp := (*ContentProtectioner)(nil)
	s.Implements(cp, &ContentProtection{})
	s.Implements(cp, ContentProtection{})
}

func (s *MPDSuite) TestCENCContentProtection_ImplementsInterface() {
	cp := (*ContentProtectioner)(nil)
	s.Implements(cp, &CENCContentProtection{})
	s.Implements(cp, CENCContentProtection{})
}

func (s *MPDSuite) TestPlayreadyContentProtection_ImplementsInterface() {
	cp := (*ContentProtectioner)(nil)
	s.Implements(cp, &PlayreadyContentProtection{})
	s.Implements(cp, PlayreadyContentProtection{})
}

func (s *MPDSuite) TestWidevineContentProtection_ImplementsInterface() {
	cp := (*ContentProtectioner)(nil)
	s.Implements(cp, &WidevineContentProtection{})
	s.Implements(cp, WidevineContentProtection{})
}

func (s *MPDSuite) TestNewMPDLiveWithBaseURLInMPD() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	m.BaseURL = VALID_BASE_URL_VIDEO
	assert.NotNil(s.T(), m)
	expectedMPD := &MPD{
		XMLNs:    Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles: Strptr((string)(DASH_PROFILE_LIVE)),
		Type:     Strptr("static"),
		MediaPresentationDuration: Strptr(VALID_MEDIA_PRESENTATION_DURATION),
		MinBufferTime:             Strptr(VALID_MIN_BUFFER_TIME),
		Period:                    &Period{},
		BaseURL:                   VALID_BASE_URL_VIDEO,
	}
	assert.Equal(s.T(), expectedMPD, m)
}

func (s *MPDSuite) TestNewMPDLiveWithBaseURLInPeriod() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	m.Period.BaseURL = VALID_BASE_URL_VIDEO
	assert.NotNil(s.T(), m)
	expectedMPD := &MPD{
		XMLNs:    Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles: Strptr((string)(DASH_PROFILE_LIVE)),
		Type:     Strptr("static"),
		MediaPresentationDuration: Strptr(VALID_MEDIA_PRESENTATION_DURATION),
		MinBufferTime:             Strptr(VALID_MIN_BUFFER_TIME),
		Period: &Period{
			BaseURL: VALID_BASE_URL_VIDEO,
		},
	}
	assert.Equal(s.T(), expectedMPD, m)
}

func (s *MPDSuite) TestNewMPDHbbTV() {
	m := NewMPD(DASH_PROFILE_HBBTV_1_5_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	assert.NotNil(s.T(), m)
	expectedMPD := &MPD{
		XMLNs:    Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles: Strptr((string)(DASH_PROFILE_HBBTV_1_5_LIVE)),
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
	as, err := m.AddNewAdaptationSetAudio(DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)
	assert.NotNil(s.T(), as)
	assert.Nil(s.T(), err)
	expectedAS := &AdaptationSet{
		MPD:              m,
		MimeType:         Strptr(VALID_MIME_TYPE_AUDIO),
		SegmentAlignment: Boolptr(VALID_SEGMENT_ALIGNMENT),
		StartWithSAP:     Int64ptr(VALID_START_WITH_SAP),
		Lang:             Strptr(VALID_LANG),
	}
	assert.Equal(s.T(), expectedAS, as)
}

func (s *MPDSuite) TestAddNewAdaptationSetVideo() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	as, err := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)
	assert.NotNil(s.T(), as)
	assert.Nil(s.T(), err)
	expectedAS := &AdaptationSet{
		MPD:              m,
		MimeType:         Strptr(VALID_MIME_TYPE_VIDEO),
		ScanType:         Strptr(VALID_SCAN_TYPE),
		SegmentAlignment: Boolptr(VALID_SEGMENT_ALIGNMENT),
		StartWithSAP:     Int64ptr(VALID_START_WITH_SAP),
	}
	assert.Equal(s.T(), expectedAS, as)
}

func (s *MPDSuite) TestAddNewAdaptationSetSubtitle() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	as, err := m.AddNewAdaptationSetSubtitle(DASH_MIME_TYPE_SUBTITLE_VTT, VALID_LANG)
	assert.NotNil(s.T(), as)
	assert.Nil(s.T(), err)
	expectedAS := &AdaptationSet{
		MPD:      m,
		MimeType: Strptr(VALID_MIME_TYPE_SUBTITLE_VTT),
		Lang:     Strptr(VALID_LANG),
	}
	assert.Equal(s.T(), expectedAS, as)
}

func (s *MPDSuite) TestAddAdaptationSetErrorNil() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	err := m.addAdaptationSet(nil)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrAdaptationSetNil, err)
}

func (s *MPDSuite) TestAddNewContentProtectionRoot() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	as, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	cp, err := as.AddNewContentProtectionRoot(VALID_DEFAULT_KID_HEX)
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), cp)
	expectedCP := &CENCContentProtection{
		DefaultKID: Strptr(VALID_DEFAULT_KID),
		Value:      Strptr(CONTENT_PROTECTION_ROOT_VALUE),
	}
	expectedCP.SchemeIDURI = Strptr(CONTENT_PROTECTION_ROOT_SCHEME_ID_URI)
	expectedCP.XMLNS = Strptr(CENC_XMLNS)

	assert.Equal(s.T(), expectedCP, cp)
}

type TestProprietaryContentProtection struct {
	ContentProtection
	TestAttrA string `xml:"a,attr,omitempty"`
	TestAttrB string `xml:"b,attr,omitempty"`
}

func (s *TestProprietaryContentProtection) ContentProtected() {}

func (s *MPDSuite) TestAddNewContentProtection_Proprietary() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	as, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	cp := &ContentProtection{
		SchemeIDURI: Strptr(CONTENT_PROTECTION_ROOT_SCHEME_ID_URI),
	}

	pcp := &TestProprietaryContentProtection{*cp, "foo", "bar"}
	x, _ := xml.Marshal(pcp)
	s.Equal(`<ContentProtection schemeIdUri="urn:mpeg:dash:mp4protection:2011" a="foo" b="bar"></ContentProtection>`, string(x))
	as.AddContentProtection(pcp)
	s.Equal(as.ContentProtection, []ContentProtectioner{pcp})
}

func (s *MPDSuite) TestAddNewContentProtectionRootErrorInvalidLengthDefaultKID() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	as, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	cp, err := as.AddNewContentProtectionRoot("invalidkid")
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrInvalidDefaultKID, err)
	assert.Nil(s.T(), cp)
}

func (s *MPDSuite) TestAddNewContentProtectionRootErrorEmptyDefaultKID() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	as, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	cp, err := as.AddNewContentProtectionRoot("")
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrInvalidDefaultKID, err)
	assert.Nil(s.T(), cp)
}

func (s *MPDSuite) TestAddNewContentProtectionSchemeWidevineWithPSSH() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	as, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	cp, err := as.AddNewContentProtectionSchemeWidevineWithPSSH(getValidWVHeaderBytes())
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), cp)
	expectedCP := &WidevineContentProtection{
		PSSH: Strptr("AAAAYXBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAAEEIARIQWr3VL1VKTyq40GH3YUJRVRoIY2FzdGxhYnMiGFdyM1ZMMVZLVHlxNDBHSDNZVUpSVlE9PTIHZGVmYXVsdA=="),
	}
	expectedCP.SchemeIDURI = Strptr(CONTENT_PROTECTION_WIDEVINE_SCHEME_ID)
	expectedCP.XMLNS = Strptr(CENC_XMLNS)
	assert.Equal(s.T(), expectedCP, cp)
}

func (s *MPDSuite) TestAddNewContentProtectionSchemeWidevine() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	as, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	cp, err := as.AddNewContentProtectionSchemeWidevine()
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), cp)
	expectedCP := &WidevineContentProtection{}
	expectedCP.SchemeIDURI = Strptr(CONTENT_PROTECTION_WIDEVINE_SCHEME_ID)
	assert.Equal(s.T(), expectedCP, cp)
}

func (s *MPDSuite) TestAddNewContentProtectionSchemePlayreadyErrorEmptyPRO() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	as, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	cp, err := as.AddNewContentProtectionSchemePlayready("")
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrPROEmpty, err)
	assert.Nil(s.T(), cp)
}

func (s *MPDSuite) TestAddNewContentProtectionSchemePlayready() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	as, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	cp, err := as.AddNewContentProtectionSchemePlayready(VALID_PLAYREADY_PRO)
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), cp)
	expectedCP := &PlayreadyContentProtection{
		PlayreadyXMLNS: Strptr(VALID_PLAYREADY_XMLNS),
		PRO:            Strptr(VALID_PLAYREADY_PRO),
	}
	expectedCP.SchemeIDURI = Strptr(CONTENT_PROTECTION_PLAYREADY_SCHEME_ID)

	assert.Equal(s.T(), expectedCP, cp)
}

func (s *MPDSuite) TestAddNewContentProtectionSchemePlayreadyV10ErrorEmptyPRO() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	as, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	cp, err := as.AddNewContentProtectionSchemePlayreadyV10("")
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrPROEmpty, err)
	assert.Nil(s.T(), cp)
}

func (s *MPDSuite) TestAddNewContentProtectionSchemePlayreadyV10() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	as, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	cp, err := as.AddNewContentProtectionSchemePlayreadyV10(VALID_PLAYREADY_PRO)
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), cp)
	expectedCP := &PlayreadyContentProtection{
		PlayreadyXMLNS: Strptr(VALID_PLAYREADY_XMLNS),
		PRO:            Strptr(VALID_PLAYREADY_PRO),
	}
	expectedCP.SchemeIDURI = Strptr(CONTENT_PROTECTION_PLAYREADY_SCHEME_V10_ID)

	assert.Equal(s.T(), expectedCP, cp)
}

func (s *MPDSuite) TestAddNewContentProtectionSchemePlayreadyWithPSSH() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	as, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	cp, err := as.AddNewContentProtectionSchemePlayreadyWithPSSH(VALID_PLAYREADY_PRO)
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), cp)
	expectedCP := &PlayreadyContentProtection{
		PlayreadyXMLNS: Strptr(VALID_PLAYREADY_XMLNS),
		PRO:            Strptr(VALID_PLAYREADY_PRO),
	}
	expectedCP.SchemeIDURI = Strptr(CONTENT_PROTECTION_PLAYREADY_SCHEME_ID)
	expectedCP.XMLNS = Strptr(CENC_XMLNS)
	expectedCP.PSSH = Strptr("AAACJnBzc2gAAAAAmgTweZhAQoarkuZb4IhflQAAAgYGAgAAAQABAPwBPABXAFIATQBIAEUAQQBEAEUAUgAgAHgAbQBsAG4AcwA9ACIAaAB0AHQAcAA6AC8ALwBzAGMAaABlAG0AYQBzAC4AbQBpAGMAcgBvAHMAbwBmAHQALgBjAG8AbQAvAEQAUgBNAC8AMgAwADAANwAvADAAMwAvAFAAbABhAHkAUgBlAGEAZAB5AEgAZQBhAGQAZQByACIAIAB2AGUAcgBzAGkAbwBuAD0AIgA0AC4AMAAuADAALgAwACIAPgA8AEQAQQBUAEEAPgA8AFAAUgBPAFQARQBDAFQASQBOAEYATwA+ADwASwBFAFkATABFAE4APgAxADYAPAAvAEsARQBZAEwARQBOAD4APABBAEwARwBJAEQAPgBBAEUAUwBDAFQAUgA8AC8AQQBMAEcASQBEAD4APAAvAFAAUgBPAFQARQBDAFQASQBOAEYATwA+ADwASwBJAEQAPgBMADkAVwA5AFcAawBwAFYASwBrACsANAAwAEcASAAzAFkAVQBKAFIAVgBRAD0APQA8AC8ASwBJAEQAPgA8AEMASABFAEMASwBTAFUATQA+AEkASwB6AFkAMgBIAFoATABBAGwASQA9ADwALwBDAEgARQBDAEsAUwBVAE0APgA8AC8ARABBAFQAQQA+ADwALwBXAFIATQBIAEUAQQBEAEUAUgA+AA==")

	assert.Equal(s.T(), expectedCP, cp)
}

func (s *MPDSuite) TestAddNewContentProtectionSchemePlayreadyV10WithPSSH() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	as, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	cp, err := as.AddNewContentProtectionSchemePlayreadyV10WithPSSH(VALID_PLAYREADY_PRO)
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), cp)
	expectedCP := &PlayreadyContentProtection{
		PlayreadyXMLNS: Strptr(VALID_PLAYREADY_XMLNS),
		PRO:            Strptr(VALID_PLAYREADY_PRO),
	}
	expectedCP.SchemeIDURI = Strptr(CONTENT_PROTECTION_PLAYREADY_SCHEME_V10_ID)
	expectedCP.XMLNS = Strptr(CENC_XMLNS)
	expectedCP.PSSH = Strptr("AAACJnBzc2gAAAAAefAEmkCYhkKrkuZb4IhflQAAAgYGAgAAAQABAPwBPABXAFIATQBIAEUAQQBEAEUAUgAgAHgAbQBsAG4AcwA9ACIAaAB0AHQAcAA6AC8ALwBzAGMAaABlAG0AYQBzAC4AbQBpAGMAcgBvAHMAbwBmAHQALgBjAG8AbQAvAEQAUgBNAC8AMgAwADAANwAvADAAMwAvAFAAbABhAHkAUgBlAGEAZAB5AEgAZQBhAGQAZQByACIAIAB2AGUAcgBzAGkAbwBuAD0AIgA0AC4AMAAuADAALgAwACIAPgA8AEQAQQBUAEEAPgA8AFAAUgBPAFQARQBDAFQASQBOAEYATwA+ADwASwBFAFkATABFAE4APgAxADYAPAAvAEsARQBZAEwARQBOAD4APABBAEwARwBJAEQAPgBBAEUAUwBDAFQAUgA8AC8AQQBMAEcASQBEAD4APAAvAFAAUgBPAFQARQBDAFQASQBOAEYATwA+ADwASwBJAEQAPgBMADkAVwA5AFcAawBwAFYASwBrACsANAAwAEcASAAzAFkAVQBKAFIAVgBRAD0APQA8AC8ASwBJAEQAPgA8AEMASABFAEMASwBTAFUATQA+AEkASwB6AFkAMgBIAFoATABBAGwASQA9ADwALwBDAEgARQBDAEsAUwBVAE0APgA8AC8ARABBAFQAQQA+ADwALwBXAFIATQBIAEUAQQBEAEUAUgA+AA==")

	assert.Equal(s.T(), expectedCP, cp)
}

func (s *MPDSuite) TestSetNewSegmentTemplate() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	audioAS, _ := m.AddNewAdaptationSetAudio(DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)
	st, err := audioAS.SetNewSegmentTemplate(VALID_DURATION, VALID_INIT_PATH_AUDIO, VALID_MEDIA_PATH_AUDIO, VALID_START_NUMBER, VALID_TIMESCALE)
	assert.NotNil(s.T(), st)
	assert.Nil(s.T(), err)
}

func (s *MPDSuite) TestSetNewSegmentTemplateErrorInvalidProfile() {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	audioAS, _ := m.AddNewAdaptationSetAudio(DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)
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
	audioAS, _ := m.AddNewAdaptationSetAudio(DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)
	st, err := audioAS.SetNewSegmentTemplate(VALID_DURATION, VALID_INIT_PATH_AUDIO, VALID_MEDIA_PATH_AUDIO, VALID_START_NUMBER, VALID_TIMESCALE)
	assert.Nil(s.T(), st)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrNoDASHProfileSet, err)
}

func (s *MPDSuite) TestAddRepresentationAudio() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	audioAS, _ := m.AddNewAdaptationSetAudio(DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)

	r, err := audioAS.AddNewRepresentationAudio(VALID_AUDIO_SAMPLE_RATE, VALID_AUDIO_BITRATE, VALID_AUDIO_CODEC, VALID_AUDIO_ID)

	assert.NotNil(s.T(), r)
	assert.Nil(s.T(), err)
}

func (s *MPDSuite) TestAddAudioChannelConfiguration() {
	m := NewMPD(DASH_PROFILE_HBBTV_1_5_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	audioAS, _ := m.AddNewAdaptationSetAudio(DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)

	r, _ := audioAS.AddNewRepresentationAudio(VALID_AUDIO_SAMPLE_RATE, VALID_AUDIO_BITRATE, VALID_AUDIO_CODEC, VALID_AUDIO_ID)

	acc, err := r.AddNewAudioChannelConfiguration(AUDIO_CHANNEL_CONFIGURATION_MPEG_DASH, "2")

	assert.NotNil(s.T(), acc)
	assert.Nil(s.T(), err)
}

func (s *MPDSuite) TestAddRepresentationVideo() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	videoAS, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, err := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	assert.NotNil(s.T(), r)
	assert.Nil(s.T(), err)
}

func (s *MPDSuite) TestAddRepresentationSubtitle() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	subtitleAS, _ := m.AddNewAdaptationSetSubtitle(DASH_MIME_TYPE_SUBTITLE_VTT, VALID_LANG)

	r, err := subtitleAS.AddNewRepresentationSubtitle(VALID_SUBTITLE_BANDWIDTH, VALID_SUBTITLE_ID)

	assert.NotNil(s.T(), r)
	assert.Nil(s.T(), err)
}

func (s *MPDSuite) TestAddRepresentationErrorNil() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	videoAS, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	err := videoAS.addRepresentation(nil)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrRepresentationNil, err)
}

func (s *MPDSuite) TestAddRole() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	audioAS, _ := m.AddNewAdaptationSetAudio(DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)

	r, err := audioAS.AddNewRole("urn:mpeg:dash:role:2011", VALID_ROLE)

	assert.NotNil(s.T(), r)
	assert.Nil(s.T(), err)
}

func (s *MPDSuite) TestSetSegmentTemplateErrorNil() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	audioAS, _ := m.AddNewAdaptationSetAudio(DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)
	err := audioAS.setSegmentTemplate(nil)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrSegmentTemplateNil, err)
}

func (s *MPDSuite) TestSetNewBaseURLVideo() {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	videoAS, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, _ := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	err := r.SetNewBaseURL(VALID_BASE_URL_VIDEO)

	assert.Nil(s.T(), err)
}

func (s *MPDSuite) TestSetNewBaseURLSubtitle() {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	subtitleAS, _ := m.AddNewAdaptationSetSubtitle(DASH_MIME_TYPE_SUBTITLE_VTT, VALID_LANG)

	r, _ := subtitleAS.AddNewRepresentationSubtitle(VALID_SUBTITLE_BANDWIDTH, VALID_SUBTITLE_ID)

	err := r.SetNewBaseURL(VALID_SUBTITLE_URL)

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
	videoAS, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, _ := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	err := r.SetNewBaseURL(VALID_BASE_URL_VIDEO)

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrNoDASHProfileSet, err)
}

func (s *MPDSuite) TestSetNewBaseURLErrorEmpty() {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	videoAS, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, _ := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	err := r.SetNewBaseURL("")

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrBaseURLEmpty, err)
}

func (s *MPDSuite) TestSetNewSegmentBase() {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	videoAS, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, _ := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	sb, err := r.AddNewSegmentBase(VALID_INDEX_RANGE, VALID_INIT_RANGE)
	assert.NotNil(s.T(), sb)
	assert.Nil(s.T(), err)
}

func (s *MPDSuite) TestSetNewSegmentBaseErrorInvalidDASHProfile() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	videoAS, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

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
	videoAS, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, _ := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	sb, err := r.AddNewSegmentBase(VALID_INDEX_RANGE, VALID_INIT_RANGE)
	assert.Nil(s.T(), sb)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrNoDASHProfileSet, err)
}

func (s *MPDSuite) TestSetSegmentBaseErrorNil() {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	videoAS, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, _ := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	err := r.setSegmentBase(nil)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrSegmentBaseNil, err)
}

func getValidWVHeaderBytes() []byte {
	wvHeader, err := base64.StdEncoding.DecodeString(VALID_WV_HEADER)
	if err != nil {
		panic(err.Error())
	}
	return wvHeader
}
