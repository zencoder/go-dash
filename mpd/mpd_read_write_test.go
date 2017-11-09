package mpd

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zencoder/go-dash/helpers/testfixtures"
)

func TestReadingManifests(t *testing.T) {
	var testCases = []struct {
		err, filepath string
	}{
		{filepath: "fixtures/live_profile.mpd", err: ""},
		{filepath: "fixtures/ondemand_profile.mpd", err: ""},
		{filepath: "fixtures/invalid.mpd", err: "XML syntax error on line 3: unexpected EOF"},
		{filepath: "doesntexist.mpd", err: "open doesntexist.mpd: no such file or directory"},
	}

	for _, tc := range testCases {
		// Test reading from manifest files
		if m, err := ReadFromFile(tc.filepath); tc.err == "" {
			require.NoError(t, err, "Error while reading "+tc.filepath)
			require.NotNil(t, m, "Empty result from reading "+tc.filepath)
		} else {
			require.EqualError(t, err, tc.err)
		}

		// Test reading valid files from strings
		if tc.err == "" {
			xmlStr := testfixtures.LoadFixture(tc.filepath)
			_, err := ReadFromString(xmlStr)
			require.NotNil(t, xmlStr)
			require.NoError(t, err)
		}
	}
}

// TODO: Uncomment this test once https://github.com/golang/go/issues/9519 is fixed
//func TestItProducesTheSameFileThatItRead(t *testing.T) {
//	expectedManifest := testfixtures.LoadFixture("fixtures/live_profile.mpd")
//	m, err := ReadFromFile("fixtures/live_profile.mpd")
//	require.NoError(t, err)
//
//	writtenManifest, err := m.WriteToString()
//	require.NoError(t, err)
//
//	require.Equal(t, expectedManifest, writtenManifest)
//}

func TestNewMPDLiveWriteToString(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	xmlStr, err := m.WriteToString()
	require.Nil(t, err)
	expectedXML := `<?xml version="1.0" encoding="UTF-8"?>
<MPD xmlns="urn:mpeg:dash:schema:mpd:2011" profiles="urn:mpeg:dash:profile:isoff-live:2011" type="static" mediaPresentationDuration="PT6M16S" minBufferTime="PT1.97S">
  <Period></Period>
</MPD>
`
	require.Equal(t, expectedXML, xmlStr)
}

func TestNewMPDOnDemandWriteToString(t *testing.T) {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	xmlStr, err := m.WriteToString()
	require.Nil(t, err)
	expectedXML := `<?xml version="1.0" encoding="UTF-8"?>
<MPD xmlns="urn:mpeg:dash:schema:mpd:2011" profiles="urn:mpeg:dash:profile:isoff-on-demand:2011" type="static" mediaPresentationDuration="PT6M16S" minBufferTime="PT1.97S">
  <Period></Period>
</MPD>
`
	require.Equal(t, expectedXML, xmlStr)
}

func TestAddNewAdaptationSetAudioWriteToString(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	m.AddNewAdaptationSetAudio(DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)

	xmlStr, err := m.WriteToString()
	require.Nil(t, err)
	expectedXML := `<?xml version="1.0" encoding="UTF-8"?>
<MPD xmlns="urn:mpeg:dash:schema:mpd:2011" profiles="urn:mpeg:dash:profile:isoff-live:2011" type="static" mediaPresentationDuration="PT6M16S" minBufferTime="PT1.97S">
  <Period>
    <AdaptationSet mimeType="audio/mp4" segmentAlignment="true" startWithSAP="1" lang="en"></AdaptationSet>
  </Period>
</MPD>
`
	require.Equal(t, expectedXML, xmlStr)
}

func TestAddNewAdaptationSetVideoWriteToString(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	xmlStr, err := m.WriteToString()
	require.Nil(t, err)
	expectedXML := `<?xml version="1.0" encoding="UTF-8"?>
<MPD xmlns="urn:mpeg:dash:schema:mpd:2011" profiles="urn:mpeg:dash:profile:isoff-live:2011" type="static" mediaPresentationDuration="PT6M16S" minBufferTime="PT1.97S">
  <Period>
    <AdaptationSet mimeType="video/mp4" scanType="progressive" segmentAlignment="true" startWithSAP="1"></AdaptationSet>
  </Period>
</MPD>
`
	require.Equal(t, expectedXML, xmlStr)
}

func TestAddNewAdaptationSetSubtitleWriteToString(t *testing.T) {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	m.AddNewAdaptationSetSubtitle(DASH_MIME_TYPE_SUBTITLE_VTT, VALID_LANG)

	xmlStr, err := m.WriteToString()
	require.Nil(t, err)
	expectedXML := `<?xml version="1.0" encoding="UTF-8"?>
<MPD xmlns="urn:mpeg:dash:schema:mpd:2011" profiles="urn:mpeg:dash:profile:isoff-live:2011" type="static" mediaPresentationDuration="PT6M16S" minBufferTime="PT1.97S">
  <Period>
    <AdaptationSet mimeType="text/vtt" lang="en"></AdaptationSet>
  </Period>
</MPD>
`
	require.Equal(t, expectedXML, xmlStr)
}

func ExampleAddNewPeriod() {
	// a new MPD is created with a single Period
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	// you can add content to the Period
	p := m.GetCurrentPeriod()
	as, _ := p.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)
	as.SetNewSegmentTemplate(1968, "$RepresentationID$/video-1.mp4", "$RepresentationID$/video-1/seg-$Number$.m4f", 0, 1000)

	// or directly to the MPD, which will use the current Period.
	as, _ = m.AddNewAdaptationSetAudio(DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)
	as.SetNewSegmentTemplate(1968, "$RepresentationID$/audio-1.mp4", "$RepresentationID$/audio-1/seg-$Number$.m4f", 0, 1000)

	// add a second period
	p = m.AddNewPeriod()
	p.SetDuration(3 * time.Minute)
	as, _ = p.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)
	as.SetNewSegmentTemplate(1968, "$RepresentationID$/video-2.mp4", "$RepresentationID$/video-2/seg-$Number$.m4f", 0, 1000)

	as, _ = m.AddNewAdaptationSetAudio(DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)
	as.SetNewSegmentTemplate(1968, "$RepresentationID$/audio-2.mp4", "$RepresentationID$/audio-2/seg-$Number$.m4f", 0, 1000)

	xmlStr, _ := m.WriteToString()
	fmt.Print(xmlStr)
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <MPD xmlns="urn:mpeg:dash:schema:mpd:2011" profiles="urn:mpeg:dash:profile:isoff-live:2011" type="static" mediaPresentationDuration="PT6M16S" minBufferTime="PT1.97S">
	//   <Period>
	//     <AdaptationSet mimeType="video/mp4" scanType="progressive" segmentAlignment="true" startWithSAP="1">
	//       <SegmentTemplate duration="1968" initialization="$RepresentationID$/video-1.mp4" media="$RepresentationID$/video-1/seg-$Number$.m4f" startNumber="0" timescale="1000"></SegmentTemplate>
	//     </AdaptationSet>
	//     <AdaptationSet mimeType="audio/mp4" segmentAlignment="true" startWithSAP="1" lang="en">
	//       <SegmentTemplate duration="1968" initialization="$RepresentationID$/audio-1.mp4" media="$RepresentationID$/audio-1/seg-$Number$.m4f" startNumber="0" timescale="1000"></SegmentTemplate>
	//     </AdaptationSet>
	//   </Period>
	//   <Period duration="PT3M0S">
	//     <AdaptationSet mimeType="video/mp4" scanType="progressive" segmentAlignment="true" startWithSAP="1">
	//       <SegmentTemplate duration="1968" initialization="$RepresentationID$/video-2.mp4" media="$RepresentationID$/video-2/seg-$Number$.m4f" startNumber="0" timescale="1000"></SegmentTemplate>
	//     </AdaptationSet>
	//     <AdaptationSet mimeType="audio/mp4" segmentAlignment="true" startWithSAP="1" lang="en">
	//       <SegmentTemplate duration="1968" initialization="$RepresentationID$/audio-2.mp4" media="$RepresentationID$/audio-2/seg-$Number$.m4f" startNumber="0" timescale="1000"></SegmentTemplate>
	//     </AdaptationSet>
	//   </Period>
	// </MPD>
}

func LiveProfile() *MPD {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	audioAS, _ := m.AddNewAdaptationSetAudio(DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)

	audioAS.AddNewContentProtectionRoot("08e367028f33436ca5dd60ffe5571e60")
	audioAS.AddNewContentProtectionSchemeWidevineWithPSSH(getValidWVHeaderBytes())
	audioAS.AddNewContentProtectionSchemePlayreadyWithPSSH(VALID_PLAYREADY_PRO)

	audioAS.AddNewRole("urn:mpeg:dash:role:2011", VALID_ROLE)

	audioAS.SetNewSegmentTemplate(1968, "$RepresentationID$/audio/en/init.mp4", "$RepresentationID$/audio/en/seg-$Number$.m4f", 0, 1000)
	audioAS.AddNewRepresentationAudio(44100, 67095, "mp4a.40.2", "800")

	videoAS, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	videoAS.AddNewContentProtectionRoot("08e367028f33436ca5dd60ffe5571e60")
	videoAS.AddNewContentProtectionSchemeWidevineWithPSSH(getValidWVHeaderBytes())
	videoAS.AddNewContentProtectionSchemePlayreadyWithPSSH(VALID_PLAYREADY_PRO)

	videoAS.AddNewRole("urn:mpeg:dash:role:2011", VALID_ROLE)

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

func TestFullLiveProfileWriteToString(t *testing.T) {
	m := LiveProfile()
	require.NotNil(t, m)
	xmlStr, err := m.WriteToString()
	require.Nil(t, err)
	expectedXML := testfixtures.LoadFixture("fixtures/live_profile.mpd")
	require.Equal(t, expectedXML, xmlStr)
}

func TestFullLiveProfileWriteToFile(t *testing.T) {
	m := LiveProfile()
	require.NotNil(t, m)
	err := m.WriteToFile("test_live.mpd")
	defer os.Remove("test_live.mpd")
	require.Nil(t, err)
}

func HbbTVProfile() *MPD {
	m := NewMPD(DASH_PROFILE_HBBTV_1_5_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)

	audioAS, _ := m.AddNewAdaptationSetAudio(DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, VALID_LANG)

	audioAS.AddNewContentProtectionRoot("08e367028f33436ca5dd60ffe5571e60")
	audioAS.AddNewContentProtectionSchemeWidevineWithPSSH(getValidWVHeaderBytes())
	audioAS.AddNewContentProtectionSchemePlayreadyWithPSSH(VALID_PLAYREADY_PRO)

	audioAS.AddNewRole("urn:mpeg:dash:role:2011", VALID_ROLE)

	audioAS.SetNewSegmentTemplate(1968, "$RepresentationID$/audio/en/init.mp4", "$RepresentationID$/audio/en/seg-$Number$.m4f", 0, 1000)
	r, _ := audioAS.AddNewRepresentationAudio(44100, 67095, "mp4a.40.2", "800")
	r.AddNewAudioChannelConfiguration(AUDIO_CHANNEL_CONFIGURATION_MPEG_DASH, "2")

	videoAS, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	videoAS.AddNewContentProtectionRoot("08e367028f33436ca5dd60ffe5571e60")
	videoAS.AddNewContentProtectionSchemeWidevineWithPSSH(getValidWVHeaderBytes())
	videoAS.AddNewContentProtectionSchemePlayreadyWithPSSH(VALID_PLAYREADY_PRO)

	videoAS.AddNewRole("urn:mpeg:dash:role:2011", VALID_ROLE)

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

func TestFullHbbTVProfileWriteToString(t *testing.T) {
	m := HbbTVProfile()
	require.NotNil(t, m)
	xmlStr, err := m.WriteToString()
	require.Nil(t, err)
	expectedXML := testfixtures.LoadFixture("fixtures/hbbtv_profile.mpd")
	require.Equal(t, expectedXML, xmlStr)
}

func TestFullHbbTVProfileWriteToFile(t *testing.T) {
	m := HbbTVProfile()
	require.NotNil(t, m)
	err := m.WriteToFile("test_hbbtv.mpd")
	defer os.Remove("test_hbbtv.mpd")
	require.Nil(t, err)
}

func OnDemandProfile() *MPD {
	m := NewMPD(DASH_PROFILE_ONDEMAND, "PT30S", VALID_MIN_BUFFER_TIME)

	audioAS, _ := m.AddNewAdaptationSetAudio(DASH_MIME_TYPE_AUDIO_MP4, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP, "und")

	audioAS.AddNewContentProtectionRoot("08e367028f33436ca5dd60ffe5571e60")
	audioAS.AddNewContentProtectionSchemeWidevineWithPSSH(getValidWVHeaderBytes())
	audioAS.AddNewContentProtectionSchemePlayreadyWithPSSH(VALID_PLAYREADY_PRO)

	audioRep, _ := audioAS.AddNewRepresentationAudio(44100, 128558, "mp4a.40.5", "800k/audio-und")
	audioRep.SetNewBaseURL("800k/output-audio-und.mp4")
	audioRep.AddNewSegmentBase("629-756", "0-628")

	videoAS, _ := m.AddNewAdaptationSetVideo(DASH_MIME_TYPE_VIDEO_MP4, VALID_SCAN_TYPE, VALID_SEGMENT_ALIGNMENT, VALID_START_WITH_SAP)

	videoAS.AddNewContentProtectionRoot("08e367028f33436ca5dd60ffe5571e60")
	videoAS.AddNewContentProtectionSchemeWidevineWithPSSH(getValidWVHeaderBytes())
	videoAS.AddNewContentProtectionSchemePlayreadyWithPSSH(VALID_PLAYREADY_PRO)

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

func TestFullOnDemandProfileWriteToString(t *testing.T) {
	m := OnDemandProfile()
	require.NotNil(t, m)
	xmlStr, err := m.WriteToString()
	require.Nil(t, err)
	expectedXML := testfixtures.LoadFixture("fixtures/ondemand_profile.mpd")
	require.Equal(t, expectedXML, xmlStr)
}

func TestFullOnDemandProfileWriteToFile(t *testing.T) {
	m := OnDemandProfile()
	require.NotNil(t, m)
	err := m.WriteToFile("test-ondemand.mpd")
	defer os.Remove("test-ondemand.mpd")
	require.Nil(t, err)
}

func TestWriteToFileInvalidFilePath(t *testing.T) {
	m := LiveProfile()
	require.NotNil(t, m)
	err := m.WriteToFile("")
	require.NotNil(t, err)
}
