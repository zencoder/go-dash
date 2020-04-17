package mpd

// Segment Template is for Live Profile Only
type SegmentTemplate struct {
	AdaptationSet          *AdaptationSet   `xml:"-"`
	SegmentTimeline        *SegmentTimeline `xml:"SegmentTimeline,omitempty"`
	PresentationTimeOffset *uint64          `xml:"presentationTimeOffset,attr,omitempty"`
	Duration               *int64           `xml:"duration,attr"`
	Initialization         *string          `xml:"initialization,attr"`
	Media                  *string          `xml:"media,attr"`
	StartNumber            *int64           `xml:"startNumber,attr"`
	Timescale              *int64           `xml:"timescale,attr"`
}
