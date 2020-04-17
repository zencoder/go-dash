package mpd

type EventStream struct {
	SchemeIdUri *string  `xml:"schemeIdUri,attr"`
	Value       *string  `xml:"value,attr,omitempty"`
	Timescale   *int64   `xml:"timescale,attr"`
	Events      []*Event `xml:"Event,omitempty"`
}
