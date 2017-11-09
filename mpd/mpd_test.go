package mpd

import (
	"encoding/base64"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/require"
	. "github.com/zencoder/go-dash/helpers/ptrs"
)

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

func TestNewMPDLive(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	require.NotNil(t, m)
	expectedMPD := &MPD{
		XMLNs:    Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles: Strptr((string)(DASH_PROFILE_LIVE)),
		Type:     Strptr("static"),
		MediaPresentationDuration: Strptr(VALID_MEDIA_PRESENTATION_DURATION),
		MinBufferTime:             Strptr(VALID_MIN_BUFFER_TIME),
		period:                    &Period{},
		Periods:                   []*Period{{}},
	}
	require.Equal(t, expectedMPD, m)
}

func TestNewMPDLiveWithBaseURLInMPD(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	m.BaseURL = VALID_BASE_URL_VIDEO
	require.NotNil(t, m)
	expectedMPD := &MPD{
		XMLNs:    Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles: Strptr((string)(DASH_PROFILE_LIVE)),
		Type:     Strptr("static"),
		MediaPresentationDuration: Strptr(VALID_MEDIA_PRESENTATION_DURATION),
		MinBufferTime:             Strptr(VALID_MIN_BUFFER_TIME),
		period:                    &Period{},
		Periods:                   []*Period{{}},
		BaseURL:                   VALID_BASE_URL_VIDEO,
	}
	require.Equal(t, expectedMPD, m)
}

func TestNewMPDLiveWithBaseURLInPeriod(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	m.period.BaseURL = VALID_BASE_URL_VIDEO
	require.NotNil(t, m)
	period := &Period{
		BaseURL: VALID_BASE_URL_VIDEO,
	}
	expectedMPD := &MPD{
		XMLNs:    Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles: Strptr((string)(DASH_PROFILE_LIVE)),
		Type:     Strptr("static"),
		MediaPresentationDuration: Strptr(VALID_MEDIA_PRESENTATION_DURATION),
		MinBufferTime:             Strptr(VALID_MIN_BUFFER_TIME),
		period:                    period,
		Periods:                   []*Period{period},
	}
	require.Equal(t, expectedMPD, m)
}

func TestNewMPDHbbTV(t *testing.T) {
	m := NewMPD(DASH_PROFILE_HBBTV_1_5_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	require.NotNil(t, m)
	expectedMPD := &MPD{
		XMLNs:    Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles: Strptr((string)(DASH_PROFILE_HBBTV_1_5_LIVE)),
		Type:     Strptr("static"),
		MediaPresentationDuration: Strptr(VALID_MEDIA_PRESENTATION_DURATION),
		MinBufferTime:             Strptr(VALID_MIN_BUFFER_TIME),
		period:                    &Period{},
		Periods:                   []*Period{{}},
	}
	require.Equal(t, expectedMPD, m)
}

func TestNewMPDOnDemand(t *testing.T) {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	require.NotNil(t, m)
	expectedMPD := &MPD{
		XMLNs:    Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles: Strptr((string)(DASH_PROFILE_ONDEMAND)),
		Type:     Strptr("static"),
		MediaPresentationDuration: Strptr(VALID_MEDIA_PRESENTATION_DURATION),
		MinBufferTime:             Strptr(VALID_MIN_BUFFER_TIME),
		period:                    &Period{},
		Periods:                   []*Period{{}},
	}
	require.Equal(t, expectedMPD, m)
}

func TestAddNewAdaptationSetAudio(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	as, err := m.AddNewAdaptationSetAudio(DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)
	require.NotNil(t, as)
	require.Nil(t, err)
	expectedAS := &AdaptationSet{
		MimeType:         Strptr(VALID_MIME_TYPE_AUDIO),
		SegmentAlignment: Boolptr(VALID_SEGMENT_ALIGNMENT),
		StartWithSAP:     Int64ptr(VALID_START_WITH_SAP),
		Lang:             Strptr(VALID_LANG),
	}
	require.Equal(t, expectedAS, as)
}

func TestAddNewAdaptationSetVideo(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	as, err := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)
	require.NotNil(t, as)
	require.Nil(t, err)
	expectedAS := &AdaptationSet{
		MimeType:         Strptr(VALID_MIME_TYPE_VIDEO),
		ScanType:         Strptr(VALID_SCAN_TYPE),
		SegmentAlignment: Boolptr(VALID_SEGMENT_ALIGNMENT),
		StartWithSAP:     Int64ptr(VALID_START_WITH_SAP),
	}
	require.Equal(t, expectedAS, as)
}

func TestAddNewAdaptationSetSubtitle(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	as, err := m.AddNewAdaptationSetSubtitle(DASH_MIME_TYPE_SUBTITLE_VTT, VALID_LANG)
	require.NotNil(t, as)
	require.Nil(t, err)
	expectedAS := &AdaptationSet{
		MimeType: Strptr(VALID_MIME_TYPE_SUBTITLE_VTT),
		Lang:     Strptr(VALID_LANG),
	}
	require.Equal(t, expectedAS, as)
}

func TestAddAdaptationSetErrorNil(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	err := m.period.addAdaptationSet(nil)
	require.NotNil(t, err)
	require.Equal(t, ErrAdaptationSetNil, err)
}

func TestAddNewContentProtectionRoot(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	s, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	cp, err := s.AddNewContentProtectionRoot(VALID_DEFAULT_KID_HEX)
	require.Nil(t, err)
	require.NotNil(t, cp)
	expectedCP := ContentProtection{
		DefaultKID: Strptr(VALID_DEFAULT_KID),
		Value:      Strptr(CONTENT_PROTECTION_ROOT_VALUE),
	}
	expectedCP.SchemeIDURI = Strptr(CONTENT_PROTECTION_ROOT_SCHEME_ID_URI)
	expectedCP.XMLNS = Strptr(CENC_XMLNS)

	require.Equal(t, expectedCP, cp)
}

func TestAddNewContentProtection_Proprietary(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	as, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	cp := ContentProtection{
		SchemeIDURI: Strptr(CONTENT_PROTECTION_ROOT_SCHEME_ID_URI),
		CustomAttrs: []xml.Attr{
			{Name: xml.Name{Local: "a"}, Value: "foo"},
			{Name: xml.Name{Local: "b"}, Value: "bar"},
		},
	}

	x, err := xml.Marshal(cp)
	require.NoError(t, err)

	require.Equal(t, `<ContentProtection schemeIdUri="urn:mpeg:dash:mp4protection:2011" a="foo" b="bar"></ContentProtection>`, string(x))
	as.AddContentProtection(cp)
	require.Equal(t, as.ContentProtection, []ContentProtection{cp})
}

func TestAddNewContentProtectionRootErrorInvalidLengthDefaultKID(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	s, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	_, err := s.AddNewContentProtectionRoot("invalidkid")
	require.NotNil(t, err)
	require.Equal(t, ErrInvalidDefaultKID, err)
}

func TestAddNewContentProtectionRootErrorEmptyDefaultKID(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	s, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	_, err := s.AddNewContentProtectionRoot("")
	require.NotNil(t, err)
	require.Equal(t, ErrInvalidDefaultKID, err)
}

func TestAddNewContentProtectionSchemeWidevineWithPSSH(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	s, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	cp, err := s.AddNewContentProtectionSchemeWidevineWithPSSH(getValidWVHeaderBytes())
	require.Nil(t, err)
	require.NotNil(t, cp)
	expectedCP := ContentProtection{
		PSSH: Strptr("AAAAYXBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAAEEIARIQWr3VL1VKTyq40GH3YUJRVRoIY2FzdGxhYnMiGFdyM1ZMMVZLVHlxNDBHSDNZVUpSVlE9PTIHZGVmYXVsdA=="),
	}
	expectedCP.SchemeIDURI = Strptr(CONTENT_PROTECTION_WIDEVINE_SCHEME_ID)
	expectedCP.XMLNS = Strptr(CENC_XMLNS)
	require.Equal(t, expectedCP, cp)
}

func TestAddNewContentProtectionSchemeWidevine(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	s, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	cp, err := s.AddNewContentProtectionSchemeWidevine()
	require.Nil(t, err)
	require.NotNil(t, cp)
	expectedCP := ContentProtection{}
	expectedCP.SchemeIDURI = Strptr(CONTENT_PROTECTION_WIDEVINE_SCHEME_ID)
	require.Equal(t, expectedCP, cp)
}

func TestAddNewContentProtectionSchemePlayreadyErrorEmptyPRO(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	s, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	_, err := s.AddNewContentProtectionSchemePlayready("")
	require.NotNil(t, err)
	require.Equal(t, ErrPROEmpty, err)
}

func TestAddNewContentProtectionSchemePlayready(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	s, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	cp, err := s.AddNewContentProtectionSchemePlayready(VALID_PLAYREADY_PRO)
	require.Nil(t, err)
	require.NotNil(t, cp)
	expectedCP := ContentProtection{
		PlayreadyXMLNS: Strptr(VALID_PLAYREADY_XMLNS),
		PRO:            Strptr(VALID_PLAYREADY_PRO),
	}
	expectedCP.SchemeIDURI = Strptr(CONTENT_PROTECTION_PLAYREADY_SCHEME_ID)

	require.Equal(t, expectedCP, cp)
}

func TestAddNewContentProtectionSchemePlayreadyV10ErrorEmptyPRO(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	s, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	_, err := s.AddNewContentProtectionSchemePlayreadyV10("")
	require.NotNil(t, err)
	require.Equal(t, ErrPROEmpty, err)
}

func TestAddNewContentProtectionSchemePlayreadyV10(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	s, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	cp, err := s.AddNewContentProtectionSchemePlayreadyV10(VALID_PLAYREADY_PRO)
	require.Nil(t, err)
	require.NotNil(t, cp)
	expectedCP := ContentProtection{
		PlayreadyXMLNS: Strptr(VALID_PLAYREADY_XMLNS),
		PRO:            Strptr(VALID_PLAYREADY_PRO),
	}
	expectedCP.SchemeIDURI = Strptr(CONTENT_PROTECTION_PLAYREADY_SCHEME_V10_ID)

	require.Equal(t, expectedCP, cp)
}

func TestAddNewContentProtectionSchemePlayreadyWithPSSH(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	s, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	cp, err := s.AddNewContentProtectionSchemePlayreadyWithPSSH(VALID_PLAYREADY_PRO)
	require.Nil(t, err)
	require.NotNil(t, cp)
	expectedCP := ContentProtection{
		PlayreadyXMLNS: Strptr(VALID_PLAYREADY_XMLNS),
		PRO:            Strptr(VALID_PLAYREADY_PRO),
	}
	expectedCP.SchemeIDURI = Strptr(CONTENT_PROTECTION_PLAYREADY_SCHEME_ID)
	expectedCP.XMLNS = Strptr(CENC_XMLNS)
	expectedCP.PSSH = Strptr("AAACJnBzc2gAAAAAmgTweZhAQoarkuZb4IhflQAAAgYGAgAAAQABAPwBPABXAFIATQBIAEUAQQBEAEUAUgAgAHgAbQBsAG4AcwA9ACIAaAB0AHQAcAA6AC8ALwBzAGMAaABlAG0AYQBzAC4AbQBpAGMAcgBvAHMAbwBmAHQALgBjAG8AbQAvAEQAUgBNAC8AMgAwADAANwAvADAAMwAvAFAAbABhAHkAUgBlAGEAZAB5AEgAZQBhAGQAZQByACIAIAB2AGUAcgBzAGkAbwBuAD0AIgA0AC4AMAAuADAALgAwACIAPgA8AEQAQQBUAEEAPgA8AFAAUgBPAFQARQBDAFQASQBOAEYATwA+ADwASwBFAFkATABFAE4APgAxADYAPAAvAEsARQBZAEwARQBOAD4APABBAEwARwBJAEQAPgBBAEUAUwBDAFQAUgA8AC8AQQBMAEcASQBEAD4APAAvAFAAUgBPAFQARQBDAFQASQBOAEYATwA+ADwASwBJAEQAPgBMADkAVwA5AFcAawBwAFYASwBrACsANAAwAEcASAAzAFkAVQBKAFIAVgBRAD0APQA8AC8ASwBJAEQAPgA8AEMASABFAEMASwBTAFUATQA+AEkASwB6AFkAMgBIAFoATABBAGwASQA9ADwALwBDAEgARQBDAEsAUwBVAE0APgA8AC8ARABBAFQAQQA+ADwALwBXAFIATQBIAEUAQQBEAEUAUgA+AA==")

	require.Equal(t, expectedCP, cp)
}

func TestAddNewContentProtectionSchemePlayreadyV10WithPSSH(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	s, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	cp, err := s.AddNewContentProtectionSchemePlayreadyV10WithPSSH(VALID_PLAYREADY_PRO)
	require.Nil(t, err)
	require.NotNil(t, cp)
	expectedCP := ContentProtection{
		PlayreadyXMLNS: Strptr(VALID_PLAYREADY_XMLNS),
		PRO:            Strptr(VALID_PLAYREADY_PRO),
	}
	expectedCP.SchemeIDURI = Strptr(CONTENT_PROTECTION_PLAYREADY_SCHEME_V10_ID)
	expectedCP.XMLNS = Strptr(CENC_XMLNS)
	expectedCP.PSSH = Strptr("AAACJnBzc2gAAAAAefAEmkCYhkKrkuZb4IhflQAAAgYGAgAAAQABAPwBPABXAFIATQBIAEUAQQBEAEUAUgAgAHgAbQBsAG4AcwA9ACIAaAB0AHQAcAA6AC8ALwBzAGMAaABlAG0AYQBzAC4AbQBpAGMAcgBvAHMAbwBmAHQALgBjAG8AbQAvAEQAUgBNAC8AMgAwADAANwAvADAAMwAvAFAAbABhAHkAUgBlAGEAZAB5AEgAZQBhAGQAZQByACIAIAB2AGUAcgBzAGkAbwBuAD0AIgA0AC4AMAAuADAALgAwACIAPgA8AEQAQQBUAEEAPgA8AFAAUgBPAFQARQBDAFQASQBOAEYATwA+ADwASwBFAFkATABFAE4APgAxADYAPAAvAEsARQBZAEwARQBOAD4APABBAEwARwBJAEQAPgBBAEUAUwBDAFQAUgA8AC8AQQBMAEcASQBEAD4APAAvAFAAUgBPAFQARQBDAFQASQBOAEYATwA+ADwASwBJAEQAPgBMADkAVwA5AFcAawBwAFYASwBrACsANAAwAEcASAAzAFkAVQBKAFIAVgBRAD0APQA8AC8ASwBJAEQAPgA8AEMASABFAEMASwBTAFUATQA+AEkASwB6AFkAMgBIAFoATABBAGwASQA9ADwALwBDAEgARQBDAEsAUwBVAE0APgA8AC8ARABBAFQAQQA+ADwALwBXAFIATQBIAEUAQQBEAEUAUgA+AA==")

	require.Equal(t, expectedCP, cp)
}

func TestSetNewSegmentTemplate(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	audioAS, _ := m.AddNewAdaptationSetAudio(DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)
	st, err := audioAS.SetNewSegmentTemplate(VALID_DURATION, VALID_INIT_PATH_AUDIO, VALID_MEDIA_PATH_AUDIO, VALID_START_NUMBER, VALID_TIMESCALE)
	require.NotNil(t, st)
	require.Nil(t, err)
}

func TestSetNewSegmentTemplateErrorNoDASHProfile(t *testing.T) {
	m := &MPD{
		XMLNs:    Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles: nil,
		Type:     Strptr("static"),
		MediaPresentationDuration: Strptr(VALID_MEDIA_PRESENTATION_DURATION),
		MinBufferTime:             Strptr(VALID_MIN_BUFFER_TIME),
		period:                    &Period{},
	}
	audioAS, _ := m.AddNewAdaptationSetAudio(DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)
	audioAS.SetNewSegmentTemplate(VALID_DURATION, VALID_INIT_PATH_AUDIO, VALID_MEDIA_PATH_AUDIO, VALID_START_NUMBER, VALID_TIMESCALE)
	err := m.Validate()
	require.Equal(t, ErrNoDASHProfileSet, err)
}

func TestAddRepresentationAudio(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	audioAS, _ := m.AddNewAdaptationSetAudio(DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)

	r, err := audioAS.AddNewRepresentationAudio(VALID_AUDIO_SAMPLE_RATE, VALID_AUDIO_BITRATE, VALID_AUDIO_CODEC, VALID_AUDIO_ID)

	require.NotNil(t, r)
	require.Nil(t, err)
}

func TestAddAudioChannelConfiguration(t *testing.T) {
	m := NewMPD(DASH_PROFILE_HBBTV_1_5_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	audioAS, _ := m.AddNewAdaptationSetAudio(DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)

	r, _ := audioAS.AddNewRepresentationAudio(VALID_AUDIO_SAMPLE_RATE, VALID_AUDIO_BITRATE, VALID_AUDIO_CODEC, VALID_AUDIO_ID)

	acc, err := r.AddNewAudioChannelConfiguration(AUDIO_CHANNEL_CONFIGURATION_MPEG_DASH, "2")

	require.NotNil(t, acc)
	require.Nil(t, err)
}

func TestAddRepresentationVideo(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	videoAS, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, err := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	require.NotNil(t, r)
	require.Nil(t, err)
}

func TestAddRepresentationSubtitle(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	subtitleAS, _ := m.AddNewAdaptationSetSubtitle(DASH_MIME_TYPE_SUBTITLE_VTT, VALID_LANG)

	r, err := subtitleAS.AddNewRepresentationSubtitle(VALID_SUBTITLE_BANDWIDTH, VALID_SUBTITLE_ID)

	require.NotNil(t, r)
	require.Nil(t, err)
}

func TestAddRepresentationErrorNil(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	videoAS, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	err := videoAS.addRepresentation(nil)
	require.NotNil(t, err)
	require.Equal(t, ErrRepresentationNil, err)
}

func TestAddRole(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	audioAS, _ := m.AddNewAdaptationSetAudio(DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)

	r, err := audioAS.AddNewRole("urn:mpeg:dash:role:2011", VALID_ROLE)

	require.NotNil(t, r)
	require.Nil(t, err)
}

func TestSetSegmentTemplateErrorNil(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	audioAS, _ := m.AddNewAdaptationSetAudio(DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)
	err := audioAS.setSegmentTemplate(nil)
	require.NotNil(t, err)
	require.Equal(t, ErrSegmentTemplateNil, err)
}

func TestSetNewBaseURLVideo(t *testing.T) {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	videoAS, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, _ := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	err := r.SetNewBaseURL(VALID_BASE_URL_VIDEO)

	require.Nil(t, err)
}

func TestSetNewBaseURLSubtitle(t *testing.T) {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	subtitleAS, _ := m.AddNewAdaptationSetSubtitle(DASH_MIME_TYPE_SUBTITLE_VTT, VALID_LANG)

	r, _ := subtitleAS.AddNewRepresentationSubtitle(VALID_SUBTITLE_BANDWIDTH, VALID_SUBTITLE_ID)

	err := r.SetNewBaseURL(VALID_SUBTITLE_URL)

	require.Nil(t, err)
}

func TestSetNewBaseURLErrorNoDASHProfile(t *testing.T) {
	m := &MPD{
		XMLNs:    Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles: nil,
		Type:     Strptr("static"),
		MediaPresentationDuration: Strptr(VALID_MEDIA_PRESENTATION_DURATION),
		MinBufferTime:             Strptr(VALID_MIN_BUFFER_TIME),
		period:                    &Period{},
	}
	videoAS, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, _ := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	r.SetNewBaseURL(VALID_BASE_URL_VIDEO)
	err := m.Validate()

	require.NotNil(t, err)
	require.Equal(t, ErrNoDASHProfileSet, err)
}

func TestSetNewBaseURLErrorEmpty(t *testing.T) {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	videoAS, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, _ := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	err := r.SetNewBaseURL("")

	require.NotNil(t, err)
	require.Equal(t, ErrBaseURLEmpty, err)
}

func TestSetNewSegmentBase(t *testing.T) {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	videoAS, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, _ := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	sb, err := r.AddNewSegmentBase(VALID_INDEX_RANGE, VALID_INIT_RANGE)
	require.NotNil(t, sb)
	require.Nil(t, err)
}

func TestSetNewSegmentBaseErrorNoDASHProfile(t *testing.T) {
	m := &MPD{
		XMLNs:    Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles: nil,
		Type:     Strptr("static"),
		MediaPresentationDuration: Strptr(VALID_MEDIA_PRESENTATION_DURATION),
		MinBufferTime:             Strptr(VALID_MIN_BUFFER_TIME),
		period:                    &Period{},
	}
	videoAS, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, _ := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	r.AddNewSegmentBase(VALID_INDEX_RANGE, VALID_INIT_RANGE)

	err := m.Validate()
	require.Equal(t, ErrNoDASHProfileSet, err)
}

func TestSetSegmentBaseErrorNil(t *testing.T) {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	videoAS, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	r, _ := videoAS.AddNewRepresentationVideo(VALID_VIDEO_BITRATE, VALID_VIDEO_CODEC, VALID_VIDEO_ID, VALID_VIDEO_FRAMERATE, VALID_VIDEO_WIDTH, VALID_VIDEO_HEIGHT)

	err := r.setSegmentBase(nil)
	require.NotNil(t, err)
	require.Equal(t, ErrSegmentBaseNil, err)
}

func getValidWVHeaderBytes() []byte {
	wvHeader, err := base64.StdEncoding.DecodeString(VALID_WV_HEADER)
	if err != nil {
		panic(err.Error())
	}
	return wvHeader
}
