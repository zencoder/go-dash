package mpd

import "github.com/jgert/go-dash/helpers/ptrs"

type ContentProtectionPlayready struct {
	Descriptor
	PSSH        *string `xml:"pssh,omitempty"`
	PRO         *string `xml:"pro,omitempty"`
	KID         *string `xml:"kid,omitempty"`
	IsEncrypted *uint   `xml:"isEncrypted,omitempty"`
	IVSize      *uint   `xml:"IV_Size,omitempty"`
}

type ContentProtectionPlayreadyMarshal struct {
	Descriptor
	NSMSPR      *string `xml:"xmlns:mspr,attr"`
	PSSH        *string `xml:"cenc:pssh,omitempty"`
	PRO         *string `xml:"mspr:pro,omitempty"`
	KID         *string `xml:"mspr:kid,omitempty"`
	IsEncrypted *uint   `xml:"mspr:isEncrypted,omitempty"`
	IVSize      *uint   `xml:"mspr:IV_Size,omitempty"`
}

func (cp *ContentProtectionPlayready) asMarshal() ContentProtectionPlayreadyMarshal {

	return ContentProtectionPlayreadyMarshal{
		NSMSPR:      ptrs.Strptr(CONTENT_PROTECTION_PLAYREADY_XMLNS),
		Descriptor:  cp.Descriptor,
		PSSH:        cp.PSSH,
		PRO:         cp.PRO,
		KID:         cp.KID,
		IsEncrypted: cp.IsEncrypted,
		IVSize:      cp.IVSize,
	}
}
