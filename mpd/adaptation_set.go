package mpd

import "encoding/xml"

type AdaptationSet struct {
	RepresentationBase
	XMLName           xml.Name            `xml:"AdaptationSet"`
	ID                *string             `xml:"id,attr"`
	SegmentAlignment  *bool               `xml:"segmentAlignment,attr"`
	Lang              *string             `xml:"lang,attr"`
	Group             *string             `xml:"group,attr"`
	PAR               *string             `xml:"par,attr"`
	MinBandwidth      *string             `xml:"minBandwidth,attr"`
	MaxBandwidth      *string             `xml:"maxBandwidth,attr"`
	MinWidth          *string             `xml:"minWidth,attr"`
	MaxHeight         *string             `xml:"maxHeight,attr,omitempty"`
	MinFrameRate      *string             `xml:"minFrameRate,attr,omitempty"`
	MaxFrameRate      *string             `xml:"maxFrameRate,attr,omitempty"`
	MaxWidth          *string             `xml:"maxWidth,attr"`
	ContentType       *string             `xml:"contentType,attr"`
	ContentProtection []ContentProtection `xml:"ContentProtection,omitempty"` // Common attribute, can be deprecated here
	Roles             []*Role             `xml:"Role,omitempty"`
	SegmentBase       *SegmentBase        `xml:"SegmentBase,omitempty"`
	SegmentList       *SegmentList        `xml:"SegmentList,omitempty"`
	SegmentTemplate   *SegmentTemplate    `xml:"SegmentTemplate,omitempty"` // Live Profile Only
	Representations   []*Representation   `xml:"Representation,omitempty"`
}
