package mpd

type Representation struct {
	RepresentationBase
	AdaptationSet   *AdaptationSet   `xml:"-"`
	BaseURL         *string          `xml:"BaseURL,omitempty"`     // On-Demand Profile
	SegmentBase     *SegmentBase     `xml:"SegmentBase,omitempty"` // On-Demand Profile
	SegmentList     *SegmentList     `xml:"SegmentList,omitempty"`
	SegmentTemplate *SegmentTemplate `xml:"SegmentTemplate,omitempty"`
	Bandwidth       *int64           `xml:"bandwidth,attr"` // Audio + Video
	ID              *string          `xml:"id,attr"`        // Audio + Video
}
