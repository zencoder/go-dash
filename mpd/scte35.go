package mpd

import (
	"encoding/xml"
	"strings"
)

type Scte35NS struct {
	XmlName xml.Name
	Value   string
}

func (s *Scte35NS) UnmarshalXMLAttr(attr xml.Attr) error {
	s.XmlName = attr.Name
	s.Value = attr.Value
	return nil
}

func (s *Scte35NS) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if strings.Contains(s.XmlName.Local, "scte35") {
		return xml.Attr{Name: xml.Name{Local: "xmlns:scte35"}, Value: s.Value}, nil
	}
	return xml.Attr{}, nil
}

type Scte35SpliceInfoSection struct {
	ProtocolVersion    *string             `xml:"protocolVersion,attr,omitempty"`
	PtsAdjustment      *int64              `xml:"ptsAdjustment,attr,omitempty"`
	Tier               *int64              `xml:"tier,attr,omitempty"`
	Scte35SpliceInsert *Scte35SpliceInsert `xml:"SpliceInsert,omitempty"`
}

type Scte35SpliceInsert struct {
	SpliceEventId              *string              `xml:"spliceEventId,attr,omitempty"`
	SpliceEventCancelIndicator bool                 `xml:"spliceEventCancelIndicator,attr"`
	OutOfNetworkIndicator      bool                 `xml:"outOfNetworkIndicator,attr,omitempty"`
	SpliceImmediateFlag        bool                 `xml:"spliceImmediateFlag,attr"`
	UniqueProgramId            *string              `xml:"uniqueProgramId,attr,omitempty"`
	AvailNum                   *int64               `xml:"availNum,attr,omitempty"`
	AvailsExpected             *int64               `xml:"availsExpected,attr,omitempty"`
	Program                    *Scte35Program       `xml:"Program,omitempty"`
	BreakDuration              *Scte35BreakDuration `xml:"BreakDuration,omitempty"`
}

type Scte35Program struct {
	Scte35SpliceTime *Scte35SpliceTime `xml:"SpliceTime,omitempty"`
}

type Scte35SpliceTime struct {
	PtsTime *int64 `xml:"ptsTime,attr,omitempty"`
}

type Scte35BreakDuration struct {
	AutoReturn bool   `xml:"autoReturn,attr,omitempty"`
	Duration   *int64 `xml:"duration,attr,omitempty"`
}

// Wrappers for handlings unmarshalling, since golang's xml namespace handling is horrifying.
type wrappedScte35Program Scte35Program
type wrappedScte35SpliceTime Scte35SpliceTime
type wrappedScte35BreakDuration Scte35BreakDuration
type wrappedScte35SpliceInsert Scte35SpliceInsert
type wrappedScte35SpliceInfoSection Scte35SpliceInfoSection

type dtoScte35Program struct {
	wrappedScte35Program
}

type dtoScte35SpliceTime struct {
	wrappedScte35SpliceTime
}

type dtoScte35BreakDuration struct {
	wrappedScte35BreakDuration
}

type dtoScte35SpliceInsert struct {
	wrappedScte35SpliceInsert
}

type dtoScte35SpliceInfoSection struct {
	wrappedScte35SpliceInfoSection
}

// Wrappers for handling marshalling, since golangs xml namespace handling is horrifying.
type Scte35ProgramMarshal struct {
	XMLName          xml.Name
	Scte35SpliceTime *Scte35SpliceTime `xml:",omitempty"`
}

type Scte35SpliceTimeMarshal struct {
	XMLName xml.Name
	PtsTime *int64 `xml:"ptsTime,attr,omitempty"`
}

type Scte35BreakDurationMarshal struct {
	XMLName    xml.Name
	AutoReturn bool   `xml:"autoReturn,attr,omitempty"`
	Duration   *int64 `xml:"duration,attr,omitempty"`
}

type Scte35SpliceInfoSectionMarshal struct {
	XMLName            xml.Name
	ProtocolVersion    *string             `xml:"protocolVersion,attr,omitempty"`
	PtsAdjustment      *int64              `xml:"ptsAdjustment,attr,omitempty"`
	Tier               *int64              `xml:"tier,attr,omitempty"`
	Scte35SpliceInsert *Scte35SpliceInsert `xml:"SpliceInsert,omitempty"`
}

type Scte35SpliceInsertMarshal struct {
	XMLName                    xml.Name
	SpliceEventId              *string              `xml:"spliceEventId,attr,omitempty"`
	SpliceEventCancelIndicator bool                 `xml:"spliceEventCancelIndicator,attr"`
	OutOfNetworkIndicator      bool                 `xml:"outOfNetworkIndicator,attr,omitempty"`
	SpliceImmediateFlag        bool                 `xml:"spliceImmediateFlag,attr"`
	UniqueProgramId            *string              `xml:"uniqueProgramId,attr,omitempty"`
	AvailNum                   *int64               `xml:"availNum,attr,omitempty"`
	AvailsExpected             *int64               `xml:"availsExpected,attr,omitempty"`
	Program                    *Scte35Program       `xml:"Program,omitempty"`
	BreakDuration              *Scte35BreakDuration `xml:"BreakDuration,omitempty"`
}

// unmarshal/marshal functions for Scte35 structures
func (s *Scte35Program) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var _s dtoScte35Program
	if err := d.DecodeElement(&_s, &start); err != nil {
		return err
	}
	*s = Scte35Program(_s.wrappedScte35Program)
	return nil
}

func (s *Scte35Program) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(&Scte35ProgramMarshal{
		XMLName:          xml.Name{Local: "scte35:Program"},
		Scte35SpliceTime: s.Scte35SpliceTime,
	})
}

func (s *Scte35SpliceTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var _s dtoScte35SpliceTime
	if err := d.DecodeElement(&_s, &start); err != nil {
		return err
	}
	*s = Scte35SpliceTime(_s.wrappedScte35SpliceTime)
	return nil
}

func (s *Scte35SpliceTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(&Scte35SpliceTimeMarshal{
		XMLName: xml.Name{Local: "scte35:SpliceTime"},
		PtsTime: s.PtsTime,
	})
}

func (s *Scte35BreakDuration) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var _s dtoScte35BreakDuration
	if err := d.DecodeElement(&_s, &start); err != nil {
		return err
	}
	*s = Scte35BreakDuration(_s.wrappedScte35BreakDuration)
	return nil
}

func (s *Scte35BreakDuration) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(&Scte35BreakDurationMarshal{
		XMLName:    xml.Name{Local: "scte35:BreakDuration"},
		AutoReturn: s.AutoReturn,
		Duration:   s.Duration,
	})
}

func (s *Scte35SpliceInsert) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var _s dtoScte35SpliceInsert
	if err := d.DecodeElement(&_s, &start); err != nil {
		return err
	}
	*s = Scte35SpliceInsert(_s.wrappedScte35SpliceInsert)
	return nil
}

func (s *Scte35SpliceInsert) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(&Scte35SpliceInsertMarshal{
		XMLName:                    xml.Name{Local: "scte35:SpliceInsert"},
		SpliceEventId:              s.SpliceEventId,
		SpliceEventCancelIndicator: s.SpliceEventCancelIndicator,
		OutOfNetworkIndicator:      s.OutOfNetworkIndicator,
		SpliceImmediateFlag:        s.SpliceImmediateFlag,
		UniqueProgramId:            s.UniqueProgramId,
		AvailNum:                   s.AvailNum,
		AvailsExpected:             s.AvailsExpected,
		Program:                    s.Program,
		BreakDuration:              s.BreakDuration,
	})
}

func (s *Scte35SpliceInfoSection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var _s dtoScte35SpliceInfoSection
	if err := d.DecodeElement(&_s, &start); err != nil {
		return err
	}
	*s = Scte35SpliceInfoSection(_s.wrappedScte35SpliceInfoSection)
	return nil
}

func (s *Scte35SpliceInfoSection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(&Scte35SpliceInfoSectionMarshal{
		XMLName:            xml.Name{Local: "scte35:SpliceInfoSection"},
		ProtocolVersion:    s.ProtocolVersion,
		PtsAdjustment:      s.PtsAdjustment,
		Tier:               s.Tier,
		Scte35SpliceInsert: s.Scte35SpliceInsert,
	})
}
