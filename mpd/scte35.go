package mpd

import (
	"encoding/xml"
	"sort"

	"github.com/Comcast/scte35-go/pkg/scte35"
	. "github.com/zencoder/go-dash/v3/helpers/ptrs"
)

const (
	SCTE352014SchemeIdUri = "urn:scte:scte35:2014:xml+bin"
	SCTE35352016Namespace = "http://www.scte.org/schemas/35/2016"
)

type Signal struct {
	XMLName   xml.Name `xml:"Signal"`
	XMLNs     *string  `xml:"xmlns,attr,omitempty"`
	Namespace *string  `xml:"namespace,attr,omitempty"`
	Binaries  []Binary `xml:"Binary,omitempty"`
}

type Binary struct {
	XMLName    xml.Name `xml:"Binary"`
	XMLNs      *string  `xml:"xmlns,attr,omitempty"`
	BinaryData *string  `xml:",chardata"`
}

// SCTE35EventOption is used to create options to modify the scte35Break Event.
type SCTE35EventOption func(scte35Break *Event)

// AddNewSCTE35Break will create a new period with an empty SCTE-35 period
// compliant with urn:scte:scte35:2014:xml+bin. This function accepts optional
// SCTE35EventOptions to further modify the break.
func (period *Period) AddNewSCTE35Break(
	timeScale uint,
	presentationTime uint64,
	id string,
	eventOptions ...SCTE35EventOption,
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
		if eventStream.SchemeIDURI == nil || *eventStream.SchemeIDURI != SCTE352014SchemeIdUri {
			continue
		}

		scte35EventStream = eventStream
		eventStreamIndex = i
		break
	}

	scte35Break := Event{
		ID:               Strptr(id),
		PresentationTime: Uint64ptr(presentationTime),
	}

	for _, eventOption := range eventOptions {
		eventOption(&scte35Break)
	}

	scte35EventStream.Events = append(scte35EventStream.Events, scte35Break)

	sort.Sort(ByPresentationTime(scte35EventStream.Events))

	if eventStreamIndex != -1 {
		period.EventStreams[eventStreamIndex] = scte35EventStream
		return
	}

	scte35EventStream.SchemeIDURI = Strptr(SCTE352014SchemeIdUri)
	scte35EventStream.Timescale = Uintptr(timeScale)
	period.EventStreams = append(period.EventStreams, scte35EventStream)
}

func getBreakSignal(scte35Break *Event) Signal {
	var signal Signal
	if len(scte35Break.Signals) != 0 {
		signal = scte35Break.Signals[0]
	}

	return signal
}

// WithBodyBinary sets the provided body binary in the break. This will also
// set the XMLNs schema to the SCTE-35 2014 Schema.
func WithBodyBinary(bodyBinary string) SCTE35EventOption {
	return func(scte35Break *Event) {
		signal := getBreakSignal(scte35Break)

		signal.XMLNs = Strptr(SCTE352014SchemeIdUri)
		signal.Binaries = []Binary{
			{
				XMLNs:      Strptr(SCTE352014SchemeIdUri),
				BinaryData: Strptr(bodyBinary),
			},
		}

		scte35Break.Signals = []Signal{signal}
	}
}

// WithNameSpace sets the provided namespace as the break's signal namespace.
func WithNameSpace(nameSpace string) SCTE35EventOption {
	return func(scte35Break *Event) {
		signal := getBreakSignal(scte35Break)

		signal.Namespace = Strptr(nameSpace)

		scte35Break.Signals = []Signal{signal}
	}
}

// WithSpliceInfoSection takes the provided scte35.SpliceInfoSection, encodes it
// and then sets it as the binary data for the break. This will also set the
// XMLNs schema to the SCTE-35 2014 Schema.
func WithSpliceInfoSection(spliceInfoSection *scte35.SpliceInfoSection) SCTE35EventOption {
	return func(scte35Break *Event) {
		signal := getBreakSignal(scte35Break)

		signal.XMLNs = Strptr(SCTE352014SchemeIdUri)
		signal.Binaries = []Binary{
			{
				XMLNs:      Strptr(SCTE352014SchemeIdUri),
				BinaryData: Strptr(spliceInfoSection.Base64()),
			},
		}

		scte35Break.Signals = []Signal{signal}
	}
}

// WithSpliceInsertCommand creates a new scte35.SpliceInfoSection with the
// splice_insert command type, and calls WithSpliceInfoSection.
func WithSpliceInsertCommand() SCTE35EventOption {
	return func(scte35Break *Event) {
		spliceInfoSection := &scte35.SpliceInfoSection{
			SpliceCommand: scte35.NewSpliceCommand(scte35.SpliceInsertType),
		}

		WithSpliceInfoSection(spliceInfoSection)(scte35Break)
	}
}
