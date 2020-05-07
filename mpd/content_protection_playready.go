package mpd

import (
	"encoding/xml"
	"github.com/jgert/go-dash/helpers/ptrs"
	"strconv"
)

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
	NSMSPR *string `xml:"xmlns:mspr,attr"`
	NSCENC *string `xml:"xmlns:cenc,attr"`
	PSSH   *string `xml:"cenc:pssh,omitempty"`
	//PRO         *string `xml:"mspr:pro,omitempty"`
	PRO *proMarshal `xml:"pro"`
	//KID         *string `xml:"mspr:kid,omitempty"`
	KID *kidMarshal `xml:"kid"`
	//IsEncrypted *uint   `xml:"mspr:isEncrypted,omitempty"`
	IsEncrypted *isEncryptedMarshal `xml:"IsEncrypted"`
	//IVSize      *uint   `xml:"mspr:IV_Size,omitempty"`
	IVSize *ivSizeMarshal `xml:"IV_Size"`
}

type proMarshal struct {
	NS  *string `xml:"xmlns,attr"`
	PRO *string `xml:",innerxml"`
}

type kidMarshal struct {
	NS  *string `xml:"xmlns,attr"`
	KID *string `xml:",innerxml"`
}

type isEncryptedMarshal struct {
	NS          *string `xml:"xmlns,attr"`
	IsEncrypted *string `xml:",innerxml"`
}

type ivSizeMarshal struct {
	NS     *string `xml:"xmlns,attr"`
	IVSize *string `xml:",innerxml"`
}

func RemoveIndex(s []*xml.Attr, index int) []*xml.Attr {
	return append(s[:index], s[index+1:]...)
}

func (cp *ContentProtectionPlayready) asMarshal() ContentProtectionPlayreadyMarshal {

	obj := ContentProtectionPlayreadyMarshal{
		Descriptor: cp.Descriptor,
		NSCENC:     ptrs.Strptr(CENC_XMLNS),
		NSMSPR:     ptrs.Strptr(CONTENT_PROTECTION_PLAYREADY_XMLNS),
		PSSH:       cp.PSSH,
		//PRO:         cp.PRO,
		//KID:         cp.KID,
		//IsEncrypted: cp.IsEncrypted,
		//IVSize:      cp.IVSize,
		PRO: &proMarshal{
			NS:  ptrs.Strptr(CONTENT_PROTECTION_PLAYREADY_XMLNS),
			PRO: cp.PRO,
		},
		KID: &kidMarshal{
			NS:  ptrs.Strptr(CONTENT_PROTECTION_PLAYREADY_XMLNS),
			KID: cp.KID,
		},
		IsEncrypted: &isEncryptedMarshal{
			NS:          ptrs.Strptr(CONTENT_PROTECTION_PLAYREADY_XMLNS),
			IsEncrypted: ptrs.Strptr(strconv.Itoa(int(*cp.IsEncrypted))),
		},
		IVSize: &ivSizeMarshal{
			NS:     ptrs.Strptr(CONTENT_PROTECTION_PLAYREADY_XMLNS),
			IVSize: ptrs.Strptr(strconv.Itoa(int(*cp.IVSize))),
		},
	}

	// remove cenc namespace from DescriptorAttributes
	var attributes []*xml.Attr
	for _, v := range obj.Attrs {
		if v.Name.Local != "" && v.Name.Space != "xmlns" && v.Value == CENC_XMLNS {
			attributes = append(attributes, v)
		}
	}
	obj.Attrs = attributes

	return obj
}
