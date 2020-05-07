package mpd

import (
	"encoding/xml"
	"strings"
)

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

type ContentProtection struct {
	XMLName   xml.Name `xml:"ContentProtection"`
	CENC      *ContentProtectionCENC
	Playready *ContentProtectionPlayready
	Widevine  *ContentProtectionWidevine
	Unknown   *Descriptor
}

func (c *ContentProtection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var (
		schemeUri string
	)

	for _, attr := range start.Attr {
		if attr.Name.Local == "schemeIdUri" {
			schemeUri = attr.Value
		}
	}
	switch strings.ToLower(schemeUri) {
	case CONTENT_PROTECTION_ROOT_SCHEME_ID_URI:
		cp := new(ContentProtectionCENC)
		err := d.DecodeElement(cp, &start)
		c.CENC = cp
		return err
	case CONTENT_PROTECTION_PLAYREADY_SCHEME_ID:
		cp := new(ContentProtectionPlayready)
		err := d.DecodeElement(cp, &start)
		c.Playready = cp
		return err
	case CONTENT_PROTECTION_WIDEVINE_SCHEME_ID:
		cp := new(ContentProtectionWidevine)
		err := d.DecodeElement(cp, &start)
		c.Widevine = cp
		return err
	default:
		cp := new(Descriptor)
		err := d.DecodeElement(cp, &start)
		c.Unknown = cp
		return err
	}
}

func (c *ContentProtection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	if c.CENC != nil {
		return e.EncodeElement(c.CENC.asMarshal(), start)
	} else if c.Playready != nil {
		return e.EncodeElement(c.Playready.asMarshal(), start)
	} else if c.Widevine != nil {
		return e.EncodeElement(c.Widevine.asMarshal(), start)
	} else {
		return nil
	}
}
