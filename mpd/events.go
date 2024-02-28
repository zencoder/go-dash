package mpd

import "encoding/xml"

type EventStream struct {
	XMLName     xml.Name `xml:"EventStream"`
	SchemeIDURI *string  `xml:"schemeIdUri,attr"`
	Value       *string  `xml:"value,attr,omitempty"`
	Timescale   *uint64  `xml:"timescale,attr"`
	Events      []Event  `xml:"Event,omitempty"`
}

type Event struct {
	XMLName           xml.Name                 `xml:"Event"`
	ID                *string                  `xml:"id,attr,omitempty"`
	SpliceInfoSection *Scte35SpliceInfoSection `xml:"SpliceInfoSection,omitempty"`
	PresentationTime  *int64                   `xml:"presentationTime,attr,omitempty"`
	Duration          *int64                   `xml:"duration,attr,omitempty"`
	MessageData       *string                  `xml:"messageData,attr,omitempty"`
}
