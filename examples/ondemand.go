package main

import (
	"fmt"

	"github.com/zencoder/go-dash/mpd"
)

func exampleOndemand() {
	m := mpd.NewMPD(mpd.DASH_PROFILE_ONDEMAND, "PT30S", "PT1.97S")

	audioAS, _ := m.AddNewAdaptationSetAudio(mpd.DASH_MIME_TYPE_AUDIO_MP4, true, 1, "und")
	_, _ = audioAS.AddNewContentProtectionRoot("08e367028f33436ca5dd60ffe5571e60")
	_, _ = audioAS.AddNewContentProtectionSchemeWidevineWithPSSH([]byte{})
	_, _ = audioAS.AddNewContentProtectionSchemePlayready("mgIAAAEAAQCQAjwAVwBSAE0ASABFAEEARABFAFIAIAB4AG0AbABuAHMAPQAiAGgAdAB0AHAAOgAvAC8AcwBjAGgAZQBtAGEAcwAuAG0AaQBjAHIAbwBzAG8AZgB0AC4AYwBvAG0ALwBEAFIATQAvADIAMAAwADcALwAwADMALwBQAGwAYQB5AFIAZQBhAGQAeQBIAGUAYQBkAGUAcgAiACAAdgBlAHIAcwBpAG8AbgA9ACIANAAuADAALgAwAC4AMAAiAD4APABEAEEAVABBAD4APABQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsARQBZAEwARQBOAD4AMQA2ADwALwBLAEUAWQBMAEUATgA+ADwAQQBMAEcASQBEAD4AQQBFAFMAQwBUAFIAPAAvAEEATABHAEkARAA+ADwALwBQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsASQBEAD4AQQBtAGYAagBDAFQATwBQAGIARQBPAGwAMwBXAEQALwA1AG0AYwBlAGMAQQA9AD0APAAvAEsASQBEAD4APABDAEgARQBDAEsAUwBVAE0APgBCAEcAdwAxAGEAWQBaADEAWQBYAE0APQA8AC8AQwBIAEUAQwBLAFMAVQBNAD4APABMAEEAXwBVAFIATAA+AGgAdAB0AHAAOgAvAC8AcABsAGEAeQByAGUAYQBkAHkALgBkAGkAcgBlAGMAdAB0AGEAcABzAC4AbgBlAHQALwBwAHIALwBzAHYAYwAvAHIAaQBnAGgAdABzAG0AYQBuAGEAZwBlAHIALgBhAHMAbQB4ADwALwBMAEEAXwBVAFIATAA+ADwALwBEAEEAVABBAD4APAAvAFcAUgBNAEgARQBBAEQARQBSAD4A")

	audioRep, _ := audioAS.AddNewRepresentationAudio(44100, 128558, "mp4a.40.5", "800k/audio-und")
	_ = audioRep.SetNewBaseURL("800k/output-audio-und.mp4")
	_, _ = audioRep.AddNewSegmentBase("629-756", "0-628")

	videoAS, _ := m.AddNewAdaptationSetVideo(mpd.DASH_MIME_TYPE_VIDEO_MP4, "progressive", true, 1)
	_, _ = videoAS.AddNewContentProtectionRoot("08e367028f33436ca5dd60ffe5571e60")
	_, _ = videoAS.AddNewContentProtectionSchemeWidevineWithPSSH([]byte{})
	_, _ = videoAS.AddNewContentProtectionSchemePlayready("mgIAAAEAAQCQAjwAVwBSAE0ASABFAEEARABFAFIAIAB4AG0AbABuAHMAPQAiAGgAdAB0AHAAOgAvAC8AcwBjAGgAZQBtAGEAcwAuAG0AaQBjAHIAbwBzAG8AZgB0AC4AYwBvAG0ALwBEAFIATQAvADIAMAAwADcALwAwADMALwBQAGwAYQB5AFIAZQBhAGQAeQBIAGUAYQBkAGUAcgAiACAAdgBlAHIAcwBpAG8AbgA9ACIANAAuADAALgAwAC4AMAAiAD4APABEAEEAVABBAD4APABQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsARQBZAEwARQBOAD4AMQA2ADwALwBLAEUAWQBMAEUATgA+ADwAQQBMAEcASQBEAD4AQQBFAFMAQwBUAFIAPAAvAEEATABHAEkARAA+ADwALwBQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsASQBEAD4AQQBtAGYAagBDAFQATwBQAGIARQBPAGwAMwBXAEQALwA1AG0AYwBlAGMAQQA9AD0APAAvAEsASQBEAD4APABDAEgARQBDAEsAUwBVAE0APgBCAEcAdwAxAGEAWQBaADEAWQBYAE0APQA8AC8AQwBIAEUAQwBLAFMAVQBNAD4APABMAEEAXwBVAFIATAA+AGgAdAB0AHAAOgAvAC8AcABsAGEAeQByAGUAYQBkAHkALgBkAGkAcgBlAGMAdAB0AGEAcABzAC4AbgBlAHQALwBwAHIALwBzAHYAYwAvAHIAaQBnAGgAdABzAG0AYQBuAGEAZwBlAHIALgBhAHMAbQB4ADwALwBMAEEAXwBVAFIATAA+ADwALwBEAEEAVABBAD4APAAvAFcAUgBNAEgARQBBAEQARQBSAD4A")

	videoRep1, _ := videoAS.AddNewRepresentationVideo(1100690, "avc1.4d401e", "800k/video-1", "30000/1001", 640, 360)
	_ = videoRep1.SetNewBaseURL("800k/output-video-1.mp4")
	_, _ = videoRep1.AddNewSegmentBase("686-813", "0-685")

	videoRep2, _ := videoAS.AddNewRepresentationVideo(1633516, "avc1.4d401f", "1200k/video-1", "30000/1001", 960, 540)
	_ = videoRep2.SetNewBaseURL("1200k/output-video-1.mp4")
	_, _ = videoRep2.AddNewSegmentBase("686-813", "0-685")

	subtitleAS, _ := m.AddNewAdaptationSetSubtitle(mpd.DASH_MIME_TYPE_SUBTITLE_VTT, "en")
	subtitleRep, _ := subtitleAS.AddNewRepresentationSubtitle(256, "captions_en")
	_ = subtitleRep.SetNewBaseURL("http://example.com/content/sintel/subtitles/subtitles_en.vtt")

	thumbnailsAS, _ := m.AddNewAdaptationSetThumbnails(mpd.DASH_MIME_TYPE_IMAGE_JPEG)
	_, _ = thumbnailsAS.SetNewSegmentTemplateThumbnails(1801800, "$RepresentationID$/$Number$.jpg", 0, 30000)
	_, _ = thumbnailsAS.AddNewRepresentationThumbnails(50000, "5x4", 1600, 720, "http://dashif.org/guidelines/thumbnail_tile")

	mpdStr, _ := m.WriteToString()
	fmt.Println(mpdStr)
}
