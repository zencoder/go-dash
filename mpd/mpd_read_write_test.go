package mpd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/zencoder/go-dash/helpers"
	"github.com/zencoder/go-dash/helpers/testfixtures"
)

type MPDReadWriteSuite struct {
	suite.Suite
}

func TestMPDReadWriteSuite(t *testing.T) {
	suite.Run(t, new(MPDReadWriteSuite))
}

func (s *MPDReadWriteSuite) SetupTest() {

}

func (s *MPDReadWriteSuite) SetupSuite() {

}

func (s *MPDReadWriteSuite) TestReadFromFileLiveProfile() {
	m, err := ReadFromFile("fixtures/live_profile.mpd")
	assert.NotNil(s.T(), m)
	assert.Nil(s.T(), err)
}

func (s *MPDReadWriteSuite) TestReadFromFileOnDemandProfile() {
	m, err := ReadFromFile("fixtures/ondemand_profile.mpd")
	assert.NotNil(s.T(), m)
	assert.Nil(s.T(), err)
}

func (s *MPDReadWriteSuite) TestReadFromFileErrorInvalidMPD() {
	m, err := ReadFromFile("fixtures/invalid.mpd")
	assert.Nil(s.T(), m)
	assert.NotNil(s.T(), err)
}

func (s *MPDReadWriteSuite) TestReadFromFileErrorInvalidFilePath() {
	m, err := ReadFromFile("this is an invalid file path")
	assert.Nil(s.T(), m)
	assert.NotNil(s.T(), err)
}

func (s *MPDReadWriteSuite) TestReadFromStringLiveProfile() {
	xmlStr := testfixtures.LoadFixture("fixtures/live_profile.mpd")
	m, err := ReadFromString(xmlStr)
	assert.NotNil(s.T(), m)
	assert.Nil(s.T(), err)
}

func (s *MPDReadWriteSuite) TestReadFromStringOnDemandProfile() {
	xmlStr := testfixtures.LoadFixture("fixtures/ondemand_profile.mpd")
	m, err := ReadFromString(xmlStr)
	assert.NotNil(s.T(), m)
	assert.Nil(s.T(), err)
}

func (s *MPDSuite) TestNewMPDLiveWriteToString() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	xmlStr, err := m.WriteToString()
	assert.Nil(s.T(), err)
	expectedXML := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<MPD xmlns=\"urn:mpeg:dash:schema:mpd:2011\" profiles=\"urn:mpeg:dash:profile:isoff-live:2011\" type=\"static\" mediaPresentationDuration=\"PT6M16S\" minBufferTime=\"PT1.97S\">\n  <Period></Period>\n</MPD>\n"
	assert.Equal(s.T(), expectedXML, xmlStr)
}

func (s *MPDSuite) TestNewMPDOnDemandWriteToString() {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	xmlStr, err := m.WriteToString()
	assert.Nil(s.T(), err)
	expectedXML := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<MPD xmlns=\"urn:mpeg:dash:schema:mpd:2011\" profiles=\"urn:mpeg:dash:profile:isoff-on-demand:2011\" type=\"static\" mediaPresentationDuration=\"PT6M16S\" minBufferTime=\"PT1.97S\">\n  <Period></Period>\n</MPD>\n"
	assert.Equal(s.T(), expectedXML, xmlStr)
}

func (s *MPDSuite) TestAddNewAdaptationSetAudioWriteToString() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	m.AddNewAdaptationSetAudio(DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)

	xmlStr, err := m.WriteToString()
	assert.Nil(s.T(), err)
	expectedXML := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<MPD xmlns=\"urn:mpeg:dash:schema:mpd:2011\" profiles=\"urn:mpeg:dash:profile:isoff-live:2011\" type=\"static\" mediaPresentationDuration=\"PT6M16S\" minBufferTime=\"PT1.97S\">\n  <Period>\n    <AdaptationSet mimeType=\"audio/mp4\" segmentAlignment=\"true\" startWithSAP=\"1\" lang=\"en\"></AdaptationSet>\n  </Period>\n</MPD>\n"
	assert.Equal(s.T(), expectedXML, xmlStr)
}

func (s *MPDSuite) TestAddNewAdaptationSetVideoWriteToString() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	xmlStr, err := m.WriteToString()
	assert.Nil(s.T(), err)
	expectedXML := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<MPD xmlns=\"urn:mpeg:dash:schema:mpd:2011\" profiles=\"urn:mpeg:dash:profile:isoff-live:2011\" type=\"static\" mediaPresentationDuration=\"PT6M16S\" minBufferTime=\"PT1.97S\">\n  <Period>\n    <AdaptationSet mimeType=\"video/mp4\" scanType=\"progressive\" segmentAlignment=\"true\" startWithSAP=\"1\"></AdaptationSet>\n  </Period>\n</MPD>\n"
	assert.Equal(s.T(), expectedXML, xmlStr)
}

func (s *MPDSuite) TestAddNewAdaptationSetSubtitleWriteToString() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	m.AddNewAdaptationSetSubtitle(DASH_MIME_TYPE_SUBTITLE_VTT, VALID_LANG)

	xmlStr, err := m.WriteToString()
	assert.Nil(s.T(), err)
	expectedXML := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<MPD xmlns=\"urn:mpeg:dash:schema:mpd:2011\" profiles=\"urn:mpeg:dash:profile:isoff-live:2011\" type=\"static\" mediaPresentationDuration=\"PT6M16S\" minBufferTime=\"PT1.97S\">\n  <Period>\n    <AdaptationSet mimeType=\"text/vtt\" lang=\"en\"></AdaptationSet>\n  </Period>\n</MPD>\n"
	assert.Equal(s.T(), expectedXML, xmlStr)
}

func LiveProfile() *MPD {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	audioAS, _ := m.AddNewAdaptationSetAudio(DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)

	audioAS.AddNewContentProtectionRoot("08e367028f33436ca5dd60ffe5571e60")
	audioAS.AddNewContentProtectionSchemeWidevine(helpers.Strptr(VALID_WV_HEADER))
	audioAS.AddNewContentProtectionSchemePlayready("mgIAAAEAAQCQAjwAVwBSAE0ASABFAEEARABFAFIAIAB4AG0AbABuAHMAPQAiAGgAdAB0AHAAOgAvAC8AcwBjAGgAZQBtAGEAcwAuAG0AaQBjAHIAbwBzAG8AZgB0AC4AYwBvAG0ALwBEAFIATQAvADIAMAAwADcALwAwADMALwBQAGwAYQB5AFIAZQBhAGQAeQBIAGUAYQBkAGUAcgAiACAAdgBlAHIAcwBpAG8AbgA9ACIANAAuADAALgAwAC4AMAAiAD4APABEAEEAVABBAD4APABQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsARQBZAEwARQBOAD4AMQA2ADwALwBLAEUAWQBMAEUATgA+ADwAQQBMAEcASQBEAD4AQQBFAFMAQwBUAFIAPAAvAEEATABHAEkARAA+ADwALwBQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsASQBEAD4AQQBtAGYAagBDAFQATwBQAGIARQBPAGwAMwBXAEQALwA1AG0AYwBlAGMAQQA9AD0APAAvAEsASQBEAD4APABDAEgARQBDAEsAUwBVAE0APgBCAEcAdwAxAGEAWQBaADEAWQBYAE0APQA8AC8AQwBIAEUAQwBLAFMAVQBNAD4APABMAEEAXwBVAFIATAA+AGgAdAB0AHAAOgAvAC8AcABsAGEAeQByAGUAYQBkAHkALgBkAGkAcgBlAGMAdAB0AGEAcABzAC4AbgBlAHQALwBwAHIALwBzAHYAYwAvAHIAaQBnAGgAdABzAG0AYQBuAGEAZwBlAHIALgBhAHMAbQB4ADwALwBMAEEAXwBVAFIATAA+ADwALwBEAEEAVABBAD4APAAvAFcAUgBNAEgARQBBAEQARQBSAD4A")

	audioAS.SetNewSegmentTemplate(1968, "$RepresentationID$/audio/en/init.mp4", "$RepresentationID$/audio/en/seg-$Number$.m4f", 0, 1000)
	audioAS.AddNewRepresentationAudio(44100, 67095, "mp4a.40.2", "800")

	videoAS, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	videoAS.AddNewContentProtectionRoot("08e367028f33436ca5dd60ffe5571e60")
	videoAS.AddNewContentProtectionSchemeWidevine(helpers.Strptr(VALID_WV_HEADER))
	videoAS.AddNewContentProtectionSchemePlayready("mgIAAAEAAQCQAjwAVwBSAE0ASABFAEEARABFAFIAIAB4AG0AbABuAHMAPQAiAGgAdAB0AHAAOgAvAC8AcwBjAGgAZQBtAGEAcwAuAG0AaQBjAHIAbwBzAG8AZgB0AC4AYwBvAG0ALwBEAFIATQAvADIAMAAwADcALwAwADMALwBQAGwAYQB5AFIAZQBhAGQAeQBIAGUAYQBkAGUAcgAiACAAdgBlAHIAcwBpAG8AbgA9ACIANAAuADAALgAwAC4AMAAiAD4APABEAEEAVABBAD4APABQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsARQBZAEwARQBOAD4AMQA2ADwALwBLAEUAWQBMAEUATgA+ADwAQQBMAEcASQBEAD4AQQBFAFMAQwBUAFIAPAAvAEEATABHAEkARAA+ADwALwBQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsASQBEAD4AQQBtAGYAagBDAFQATwBQAGIARQBPAGwAMwBXAEQALwA1AG0AYwBlAGMAQQA9AD0APAAvAEsASQBEAD4APABDAEgARQBDAEsAUwBVAE0APgBCAEcAdwAxAGEAWQBaADEAWQBYAE0APQA8AC8AQwBIAEUAQwBLAFMAVQBNAD4APABMAEEAXwBVAFIATAA+AGgAdAB0AHAAOgAvAC8AcABsAGEAeQByAGUAYQBkAHkALgBkAGkAcgBlAGMAdAB0AGEAcABzAC4AbgBlAHQALwBwAHIALwBzAHYAYwAvAHIAaQBnAGgAdABzAG0AYQBuAGEAZwBlAHIALgBhAHMAbQB4ADwALwBMAEEAXwBVAFIATAA+ADwALwBEAEEAVABBAD4APAAvAFcAUgBNAEgARQBBAEQARQBSAD4A")

	videoAS.SetNewSegmentTemplate(1968, "$RepresentationID$/video/1/init.mp4", "$RepresentationID$/video/1/seg-$Number$.m4f", 0, 1000)
	videoAS.AddNewRepresentationVideo(1518664, "avc1.4d401f", "800", "30000/1001", 960, 540)
	videoAS.AddNewRepresentationVideo(1911775, "avc1.4d401f", "1000", "30000/1001", 1024, 576)
	videoAS.AddNewRepresentationVideo(2295158, "avc1.4d401f", "1200", "30000/1001", 1024, 576)
	videoAS.AddNewRepresentationVideo(2780732, "avc1.4d401f", "1500", "30000/1001", 1280, 720)

	subtitleAS, _ := m.AddNewAdaptationSetSubtitle(DASH_MIME_TYPE_SUBTITLE_VTT, VALID_LANG)
	subtitleRep, _ := subtitleAS.AddNewRepresentationSubtitle(VALID_SUBTITLE_BANDWIDTH, VALID_SUBTITLE_ID)
	subtitleRep.SetNewBaseURL(VALID_SUBTITLE_URL)

	return m
}

func (s *MPDSuite) TestFullLiveProfileWriteToString() {
	m := LiveProfile()
	assert.NotNil(s.T(), m)
	xmlStr, err := m.WriteToString()
	assert.Nil(s.T(), err)
	expectedXML := testfixtures.LoadFixture("fixtures/live_profile.mpd")
	assert.Equal(s.T(), expectedXML, xmlStr)
}

func (s *MPDSuite) TestFullLiveProfileWriteToFile() {
	m := LiveProfile()
	assert.NotNil(s.T(), m)
	err := m.WriteToFile("test_live.mpd")
	defer os.Remove("test_live.mpd")
	assert.Nil(s.T(), err)
}

func OnDemandProfile() *MPD {
	m := NewMPD(DASH_PROFILE_ONDEMAND, "PT30S", VALID_MIN_BUFFER_TIME)

	audioAS, _ := m.AddNewAdaptationSetAudio(DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, "und")

	audioAS.AddNewContentProtectionRoot("08e367028f33436ca5dd60ffe5571e60")
	audioAS.AddNewContentProtectionSchemeWidevine(helpers.Strptr(VALID_WV_HEADER))
	audioAS.AddNewContentProtectionSchemePlayready("mgIAAAEAAQCQAjwAVwBSAE0ASABFAEEARABFAFIAIAB4AG0AbABuAHMAPQAiAGgAdAB0AHAAOgAvAC8AcwBjAGgAZQBtAGEAcwAuAG0AaQBjAHIAbwBzAG8AZgB0AC4AYwBvAG0ALwBEAFIATQAvADIAMAAwADcALwAwADMALwBQAGwAYQB5AFIAZQBhAGQAeQBIAGUAYQBkAGUAcgAiACAAdgBlAHIAcwBpAG8AbgA9ACIANAAuADAALgAwAC4AMAAiAD4APABEAEEAVABBAD4APABQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsARQBZAEwARQBOAD4AMQA2ADwALwBLAEUAWQBMAEUATgA+ADwAQQBMAEcASQBEAD4AQQBFAFMAQwBUAFIAPAAvAEEATABHAEkARAA+ADwALwBQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsASQBEAD4AQQBtAGYAagBDAFQATwBQAGIARQBPAGwAMwBXAEQALwA1AG0AYwBlAGMAQQA9AD0APAAvAEsASQBEAD4APABDAEgARQBDAEsAUwBVAE0APgBCAEcAdwAxAGEAWQBaADEAWQBYAE0APQA8AC8AQwBIAEUAQwBLAFMAVQBNAD4APABMAEEAXwBVAFIATAA+AGgAdAB0AHAAOgAvAC8AcABsAGEAeQByAGUAYQBkAHkALgBkAGkAcgBlAGMAdAB0AGEAcABzAC4AbgBlAHQALwBwAHIALwBzAHYAYwAvAHIAaQBnAGgAdABzAG0AYQBuAGEAZwBlAHIALgBhAHMAbQB4ADwALwBMAEEAXwBVAFIATAA+ADwALwBEAEEAVABBAD4APAAvAFcAUgBNAEgARQBBAEQARQBSAD4A")

	audioRep, _ := audioAS.AddNewRepresentationAudio(44100, 128558, "mp4a.40.5", "800k/audio-und")
	audioRep.SetNewBaseURL("800k/output-audio-und.mp4")
	audioRep.AddNewSegmentBase("629-756", "0-628")

	videoAS, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	videoAS.AddNewContentProtectionRoot("08e367028f33436ca5dd60ffe5571e60")
	videoAS.AddNewContentProtectionSchemeWidevine(helpers.Strptr(VALID_WV_HEADER))
	videoAS.AddNewContentProtectionSchemePlayready("mgIAAAEAAQCQAjwAVwBSAE0ASABFAEEARABFAFIAIAB4AG0AbABuAHMAPQAiAGgAdAB0AHAAOgAvAC8AcwBjAGgAZQBtAGEAcwAuAG0AaQBjAHIAbwBzAG8AZgB0AC4AYwBvAG0ALwBEAFIATQAvADIAMAAwADcALwAwADMALwBQAGwAYQB5AFIAZQBhAGQAeQBIAGUAYQBkAGUAcgAiACAAdgBlAHIAcwBpAG8AbgA9ACIANAAuADAALgAwAC4AMAAiAD4APABEAEEAVABBAD4APABQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsARQBZAEwARQBOAD4AMQA2ADwALwBLAEUAWQBMAEUATgA+ADwAQQBMAEcASQBEAD4AQQBFAFMAQwBUAFIAPAAvAEEATABHAEkARAA+ADwALwBQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsASQBEAD4AQQBtAGYAagBDAFQATwBQAGIARQBPAGwAMwBXAEQALwA1AG0AYwBlAGMAQQA9AD0APAAvAEsASQBEAD4APABDAEgARQBDAEsAUwBVAE0APgBCAEcAdwAxAGEAWQBaADEAWQBYAE0APQA8AC8AQwBIAEUAQwBLAFMAVQBNAD4APABMAEEAXwBVAFIATAA+AGgAdAB0AHAAOgAvAC8AcABsAGEAeQByAGUAYQBkAHkALgBkAGkAcgBlAGMAdAB0AGEAcABzAC4AbgBlAHQALwBwAHIALwBzAHYAYwAvAHIAaQBnAGgAdABzAG0AYQBuAGEAZwBlAHIALgBhAHMAbQB4ADwALwBMAEEAXwBVAFIATAA+ADwALwBEAEEAVABBAD4APAAvAFcAUgBNAEgARQBBAEQARQBSAD4A")

	videoRep1, _ := videoAS.AddNewRepresentationVideo(1100690, "avc1.4d401e", "800k/video-1", "30000/1001", 640, 360)
	videoRep1.SetNewBaseURL("800k/output-video-1.mp4")
	videoRep1.AddNewSegmentBase("686-813", "0-685")

	videoRep2, _ := videoAS.AddNewRepresentationVideo(1633516, "avc1.4d401f", "1200k/video-1", "30000/1001", 960, 540)
	videoRep2.SetNewBaseURL("1200k/output-video-1.mp4")
	videoRep2.AddNewSegmentBase("686-813", "0-685")

	subtitleAS, _ := m.AddNewAdaptationSetSubtitle(DASH_MIME_TYPE_SUBTITLE_VTT, VALID_LANG)
	subtitleRep, _ := subtitleAS.AddNewRepresentationSubtitle(VALID_SUBTITLE_BANDWIDTH, VALID_SUBTITLE_ID)
	subtitleRep.SetNewBaseURL(VALID_SUBTITLE_URL)

	return m
}

func (s *MPDSuite) TestFullOnDemandProfileWriteToString() {
	m := OnDemandProfile()
	assert.NotNil(s.T(), m)
	xmlStr, err := m.WriteToString()
	assert.Nil(s.T(), err)
	expectedXML := testfixtures.LoadFixture("fixtures/ondemand_profile.mpd")
	assert.Equal(s.T(), expectedXML, xmlStr)
}

func (s *MPDSuite) TestFullOnDemandProfileWriteToFile() {
	m := OnDemandProfile()
	assert.NotNil(s.T(), m)
	err := m.WriteToFile("test-ondemand.mpd")
	defer os.Remove("test-ondemand.mpd")
	assert.Nil(s.T(), err)
}

func (s *MPDSuite) TestWriteToFileInvalidFilePath() {
	m := LiveProfile()
	assert.NotNil(s.T(), m)
	err := m.WriteToFile("")
	assert.NotNil(s.T(), err)
}
