package mpd

import (
	"encoding/xml"
	"sort"

	. "github.com/zencoder/go-dash/v3/helpers/ptrs"
)

const SCTE352014SchemaUri = "urn:scte:scte35:2014:xml+bin"

type Signal struct {
	XMLName  xml.Name `xml:"Signal"`
	XMLNs    *string  `xml:"xmlns,attr,omitempty"`
	Binaries []Binary `xml:"Binary,omitempty"`
}

type Binary struct {
	XMLName    xml.Name `xml:"Binary"`
	XMLNs      *string  `xml:"xmlns,attr,omitempty"`
	BinaryData *string  `xml:",chardata"`
}

// AddNewSCTE35Break will create a new period with an empty SCTE-35 period
// compliant with urn:scte:scte35:2014:xml+bin. The optional body can be passed,
// if not, an empty scte35.SpliceInfoSection will be generated.
func (period *Period) AddNewSCTE35Break(
	timeScale uint,
	presentationTime uint64,
	id, bodyBinary string,
) {
	var eventStreams []EventStream
	if period.EventStreams != nil {
		eventStreams = period.EventStreams
	}

	var (
		scte35EventStream EventStream
		eventStreamIndex  = -1
	)
	for i, eventStream := range eventStreams {
		if eventStream.SchemeIDURI == nil || *eventStream.SchemeIDURI != SCTE352014SchemaUri {
			continue
		}

		scte35EventStream = eventStream
		eventStreamIndex = i
		break
	}

	scte35Break := Event{
		ID:               Strptr(id),
		PresentationTime: Uint64ptr(presentationTime),
		Signals: []Signal{
			{
				XMLNs: Strptr(SCTE352014SchemaUri),
				Binaries: []Binary{
					{
						XMLNs:      Strptr(SCTE352014SchemaUri),
						BinaryData: Strptr(bodyBinary),
					},
				},
			},
		},
	}

	scte35EventStream.Events = append(scte35EventStream.Events, scte35Break)

	sort.Sort(ByPresentationTime(scte35EventStream.Events))

	if eventStreamIndex != -1 {
		period.EventStreams[eventStreamIndex] = scte35EventStream
		return
	}

	scte35EventStream.SchemeIDURI = Strptr(SCTE352014SchemaUri)
	scte35EventStream.Timescale = Uintptr(timeScale)
	period.EventStreams = append(period.EventStreams, scte35EventStream)
}
