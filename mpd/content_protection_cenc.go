package mpd

import "encoding/xml"

type ContentProtectionCENC struct {
	Descriptor
	XMLNS      *string `xml:"cenc,attr"`
	DefaultKID *string `xml:"default_KID,attr"`
}

type ContentProtectionCENCMarshal struct {
	Descriptor
	XMLNS      *string `xml:"xmlns:cenc,attr"`
	DefaultKID *string `xml:"cenc:default_KID,attr"`
}

func (cp *ContentProtectionCENC) EncodeElement(e *xml.Encoder, start xml.StartElement) {

}

func (cp *ContentProtectionCENC) asMarshal() ContentProtectionCENCMarshal {
	return ContentProtectionCENCMarshal{
		cp.Descriptor,
		cp.XMLNS,
		cp.DefaultKID,
	}
}
