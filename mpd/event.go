package mpd

import (
	"encoding/base64"
	"encoding/xml"
	"github.com/Comcast/gots/scte35"
)

type Event struct {
	ID               string  `xml:"id,attr"`
	Duration         int64   `xml:"duration,attr"`
	PresentationTime int64   `xml:"presentationTime,attr"`
	Signal           *Signal `xml:"Signal"`
}

type Signal struct {
	XMLNS  string `xml:"xmlns,attr"`
	Binary Binary `xml:"Binary"`
}

type Binary struct {
	scte35.SCTE35
}

func (b *Binary) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	value := base64.StdEncoding.EncodeToString(b.SCTE35.UpdateData())
	return e.EncodeElement(value, start)
}

func (b *Binary) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var o string
	if err := d.DecodeElement(&o, &start); err != nil {
		return err
	}

	bytes, err := base64.StdEncoding.DecodeString(o)
	if err != nil {
		return err
	}

	bytes = append([]byte{0x00}, bytes...)

	marker, err := scte35.NewSCTE35(bytes)

	b.SCTE35 = marker

	return err
}
