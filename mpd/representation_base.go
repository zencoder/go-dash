package mpd

type RepresentationBase struct {
	AudioChannelConfiguration []*Descriptor        `xml:"AudioChannelConfiguration"`
	ContentProtection         []*ContentProtection `xml:"ContentProtection,omitempty"`
	EssentialProperty         []*Descriptor        `xml:"EssentialProperty"`
	FramePacking              []*Descriptor        `xml:"FramePacking"`
	InbandEventStream         []*Descriptor        `xml:"InbandEventStream"`
	SupplementalProperty      []*Descriptor        `xml:"SupplementalProperty"`

	AudioSamplingRate *int64  `xml:"audioSamplingRate,attr"`
	Codecs            *string `xml:"codecs,attr"`
	FrameRate         *string `xml:"frameRate,attr"`
	Height            *uint32 `xml:"height,attr"`
	MaximumSAPPeriod  *string `xml:"maximumSAPPeriod,attr"`
	MaxPlayoutRate    *string `xml:"maxPlayoutRate,attr"`
	MimeType          *string `xml:"mimeType,attr"`
	Profiles          *string `xml:"profiles,attr"`
	Sar               *string `xml:"sar,attr"`
	ScanType          *string `xml:"scanType,attr"`
	SegmentProfiles   *string `xml:"segmentProfiles,attr"`
	StartWithSAP      *int64  `xml:"startWithSAP,attr"`
	Width             *uint32 `xml:"width,attr"`
}
