package mpd

import "encoding/xml"

// Constants for DRM / ContentProtection
const (
	CONTENT_PROTECTION_ROOT_SCHEME_ID_URI       = "urn:mpeg:dash:mp4protection:2011"
	CONTENT_PROTECTION_ROOT_VALUE               = "cenc"
	CENC_XMLNS                                  = "urn:mpeg:cenc:2013"
	CONTENT_PROTECTION_WIDEVINE_SCHEME_ID       = "urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed"
	CONTENT_PROTECTION_WIDEVINE_SCHEME_HEX      = "edef8ba979d64acea3c827dcd51d21ed"
	CONTENT_PROTECTION_PLAYREADY_SCHEME_ID      = "urn:uuid:9a04f079-9840-4286-ab92-e65be0885f95"
	CONTENT_PROTECTION_PLAYREADY_SCHEME_HEX     = "9a04f07998404286ab92e65be0885f95"
	CONTENT_PROTECTION_PLAYREADY_SCHEME_V10_ID  = "urn:uuid:79f0049a-4098-8642-ab92-e65be0885f95"
	CONTENT_PROTECTION_PLAYREADY_SCHEME_V10_HEX = "79f0049a40988642ab92e65be0885f95"
	CONTENT_PROTECTION_PLAYREADY_XMLNS          = "urn:microsoft:playready"
)

type ContentProtectionContainer struct {
	XMLName             xml.Name `xml:"ContentProtection"`
	ContentProtectioner ContentProtectioner
}

func (c *ContentProtectionContainer) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var (
		schemeUri string
		cp        ContentProtectioner
	)

	for _, attr := range start.Attr {
		if attr.Name.Local == "schemeIdUri" {
			schemeUri = attr.Value
		}
	}
	switch schemeUri {
	case CONTENT_PROTECTION_ROOT_SCHEME_ID_URI:
		cp = new(CENCContentProtection)
	case CONTENT_PROTECTION_PLAYREADY_SCHEME_ID:
		cp = new(PlayreadyContentProtection)
	case CONTENT_PROTECTION_WIDEVINE_SCHEME_ID:
		cp = new(WidevineContentProtection)
	default:
		cp = new(ContentProtection)
	}

	err := d.DecodeElement(cp, &start)
	if err != nil {
		return err
	}
	c.ContentProtectioner = cp
	//*c = ContentProtectionContainer{ContentProtectioner:cp}
	return nil
}

func (c *ContentProtectionContainer) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(c.ContentProtectioner)
}

type ContentProtectioner interface {
	ContentProtected()
}

//
//type ContentProtection struct {
//	Descriptor
//	XMLName       xml.Name       `xml:"ContentProtection"`
//}

type ContentProtection struct {
	Descriptor
	XMLName xml.Name `xml:"ContentProtection"`
	//AdaptationSet *AdaptationSet `xml:"-"`
	//
	//SchemeIDURI   *string        `xml:"schemeIdUri,attr"` // Default: urn:mpeg:dash:mp4protection:2011
	//XMLNS         *string        `xml:"cenc,attr"`        // Default: urn:mpeg:cenc:2013
	//Attrs         []*xml.Attr    `xml:",any,attr"`
}

type CENCContentProtection struct {
	ContentProtection
	DefaultKID *string `xml:"default_KID,attr"`
}

type PlayreadyContentProtection struct {
	ContentProtection
	PlayreadyXMLNS *string `xml:"mspr,attr,omitempty"`
	PRO            *string `xml:"pro,omitempty"`
	PSSH           *string `xml:"pssh,omitempty"`
}

type WidevineContentProtection struct {
	ContentProtection
	PSSH *string `xml:"pssh,omitempty"`
}

type ContentProtectionMarshal struct {
	//XMLName     xml.Name    `xml:"ContentProtection"`
	//SchemeIDURI *string     `xml:"schemeIdUri,attr"` // Default: urn:mpeg:dash:mp4protection:2011
	//XMLNS       *string     `xml:"xmlns:cenc,attr"`  // Default: urn:mpeg:cenc:2013
	//Attrs       []*xml.Attr `xml:",any,attr"`
}

type CENCContentProtectionMarshal struct {
	ContentProtectionMarshal
	DefaultKID *string `xml:"cenc:default_KID,attr"`
	Value      *string `xml:"value,attr"` // Default: cenc
}

type PlayreadyContentProtectionMarshal struct {
	ContentProtectionMarshal
	PlayreadyXMLNS *string `xml:"xmlns:mspr,attr,omitempty"`
	PRO            *string `xml:"mspr:pro,omitempty"`
	PSSH           *string `xml:"cenc:pssh,omitempty"`
}

type WidevineContentProtectionMarshal struct {
	ContentProtectionMarshal
	PSSH *string `xml:"cenc:pssh,omitempty"`
}

func (s ContentProtection) ContentProtected() {}

func (s ContentProtection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	//err := e.Encode(&ContentProtectionMarshal{
	//	s.XMLName,
	//	s.SchemeIDURI,
	//	//s.XMLNS,
	//	s.Attrs,
	//})
	//if err != nil {
	//	return err
	//}
	return nil
}

func (s CENCContentProtection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	//err := e.Encode(&CENCContentProtectionMarshal{
	//	ContentProtectionMarshal{
	//		s.XMLName,
	//		s.SchemeIDURI,
	//		//s.XMLNS,
	//		s.Attrs,
	//	},
	//	s.DefaultKID,
	//	s.Value,
	//})
	//if err != nil {
	//	return err
	//}
	return nil
}

func (s PlayreadyContentProtection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	//err := e.Encode(&PlayreadyContentProtectionMarshal{
	//	ContentProtectionMarshal{
	//		s.AdaptationSet,
	//		s.XMLName,
	//		s.SchemeIDURI,
	//		s.XMLNS,
	//		s.Attrs,
	//	},
	//	s.PlayreadyXMLNS,
	//	s.PRO,
	//	s.PSSH,
	//})
	//if err != nil {
	//	return err
	//}
	return nil
}

func (s WidevineContentProtection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	//err := e.Encode(&WidevineContentProtectionMarshal{
	//	ContentProtectionMarshal{
	//		s.AdaptationSet,
	//		s.XMLName,
	//		s.SchemeIDURI,
	//		s.XMLNS,
	//		s.Attrs,
	//	},
	//	s.PSSH,
	//})
	//if err != nil {
	//	return err
	//}
	return nil
}
