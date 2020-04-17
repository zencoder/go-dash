package mpd

type Period struct {
	ID              string           `xml:"id,attr,omitempty"`
	Duration        *Duration        `xml:"duration,attr,omitempty"`
	Start           *Duration        `xml:"start,attr,omitempty"`
	BaseURL         string           `xml:"BaseURL,omitempty"`
	EventStream     *EventStream     `xml:"EventStream,omitempty"`
	SegmentBase     *SegmentBase     `xml:"SegmentBase,omitempty"`
	SegmentList     *SegmentList     `xml:"SegmentList,omitempty"`
	SegmentTemplate *SegmentTemplate `xml:"SegmentTemplate,omitempty"`
	AdaptationSets  []*AdaptationSet `xml:"AdaptationSet,omitempty"`
}
