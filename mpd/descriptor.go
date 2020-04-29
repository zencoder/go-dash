package mpd

import "encoding/xml"

type Descriptor struct {
	SchemeIDURI *string     `xml:"schemeIdUri,attr"`
	Value       *string     `xml:"value,attr,omitempty"`
	ID          *string     `xml:"id,attr"`
	Attrs       []*xml.Attr `xml:",any,attr"`
}
