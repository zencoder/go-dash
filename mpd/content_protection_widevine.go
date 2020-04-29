package mpd

type ContentProtectionWidevine struct {
	Descriptor
	XMLNS *string `xml:"cenc,attr"`
	PSSH  *string `xml:"pssh,omitempty"`
}

type ContentProtectionWidevineMarshal struct {
	Descriptor
	XMLNS *string `xml:"xmlns:cenc,attr"`
	PSSH  *string `xml:"cenc:pssh,omitempty"`
}

func (cp *ContentProtectionWidevine) asMarshal() ContentProtectionWidevineMarshal {
	return ContentProtectionWidevineMarshal{
		cp.Descriptor,
		cp.XMLNS,
		cp.PSSH,
	}
}
