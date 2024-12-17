package mpd

import "encoding/xml"

type EventStream struct {
	XMLName     xml.Name `xml:"EventStream"`
	SchemeIDURI *string  `xml:"schemeIdUri,attr"`
	Value       *string  `xml:"value,attr,omitempty"`
	Timescale   *uint    `xml:"timescale,attr"`
	Events      []Event  `xml:"Event,omitempty"`
}

type Event struct {
	XMLName          xml.Name `xml:"Event"`
	ID               *string  `xml:"id,attr,omitempty"`
	PresentationTime *uint64  `xml:"presentationTime,attr,omitempty"`
	Duration         *uint64  `xml:"duration,attr,omitempty"`
	Signals          []Signal `xml:"Signal,omitempty"`
}

type ByPresentationTime []Event

func (p ByPresentationTime) Len() int {
	return len(p)
}

func (p ByPresentationTime) Less(i, j int) bool {
	return *p[i].PresentationTime < *p[j].PresentationTime
}

func (p ByPresentationTime) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
