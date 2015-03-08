package mpd

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
)

type DashProfile string

const (
	DASH_PROFILE_LIVE     DashProfile = "urn:mpeg:dash:profile:isoff-live:2011"
	DASH_PROFILE_ONDEMAND DashProfile = "urn:mpeg:dash:profile:isoff-on-demand:2011"
)

type MPD struct {
	XMLNs                     *string `xml:"xmlns,attr"`
	Profiles                  *string `xml:"profiles,attr"`
	Type                      *string `xml:"type,attr"`
	MediaPresentationDuration *string `xml:"mediaPresentationDuration,attr"`
	MinBufferTime             *string `xml:"minBufferTime,attr"`
	Period                    *Period `xml:"Period"`
}

func (m MPD) String() string {
	jb, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(jb)
}

type Period struct {
	AdaptationSets []*AdaptationSet `xml:"AdaptationSet"`
}

func (p Period) String() string {
	jb, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(jb)
}

type AdaptationSet struct {
	MimeType         *string           `xml:"mimeType,attr"`
	ScanType         *string           `xml:"scanType,attr"`
	SegmentAlignment *bool             `xml:"segmentAlignment,attr"`
	StartWithSAP     *int64            `xml:"startWithSAP,attr"`
	Lang             *string           `xml:"lang,attr"`
	SegmentTemplate  *SegmentTemplate  `xml:"SegmentTemplate"` // Live Profile Only
	Representations  []*Representation `xml:"Representation"`
}

func (as AdaptationSet) String() string {
	jb, err := json.Marshal(as)
	if err != nil {
		return ""
	}
	return string(jb)
}

// Live Profile Only
type SegmentTemplate struct {
	Duration       *int64  `xml:"duration,attr"`
	Initialization *string `xml:"initialization,attr"`
	Media          *string `xml:"media,attr"`
	StartNumber    *int64  `xml:"startNumber,attr"`
	Timescale      *int64  `xml:"timescale,attr"`
}

func (st SegmentTemplate) String() string {
	jb, err := json.Marshal(st)
	if err != nil {
		return ""
	}
	return string(jb)
}

type Representation struct {
	AudioSamplingRate *int64       `xml:"audioSamplingRate,attr"` // Audio
	Bandwidth         *int64       `xml:"bandwidth,attr"`         // Audio + Video
	Codecs            *string      `xml:"codecs,attr"`            // Audio + Video
	ID                *string      `xml:"id,attr"`                // Audio + Video
	FrameRate         *string      `xml:"frameRate,attr"`         // Video
	Width             *int64       `xml:"width,attr"`             // Video
	Height            *int64       `xml:"height,attr"`            // Video
	BaseURL           *string      `xml:"xml:BaseURL"`            // On-Demand Profile
	SegmentBase       *SegmentBase `xml:"SegmentBase"`            // On-Demand Profile
}

func (r Representation) String() string {
	jb, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(jb)
}

// On-Demand Profile
type SegmentBase struct {
	IndexRange     *string         `xml:"indexRange,attr"`
	Initialization *Initialization `xml:"Initialization"`
}

func (sb SegmentBase) String() string {
	jb, err := json.Marshal(sb)
	if err != nil {
		return ""
	}
	return string(jb)
}

// On-Demand Profile
type Initialization struct {
	Range *string `xml:"range,attr"`
}

func (i Initialization) String() string {
	jb, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(jb)
}

func ReadMPD(path string) (*MPD, error) {
	var mpd MPD
	f, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	d := xml.NewDecoder(f)
	d.Decode(&mpd)

	return &mpd, nil
}

func (m *MPD) WriteToFile(path string) error {
	// Open the file to write the XML to
	f, err := os.OpenFile(path, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	// Write out the XML Header
	f.Write([]byte(xml.Header))
	// Write out the DASH XML manifest
	e := xml.NewEncoder(f)
	err = e.Encode(m)
	if err != nil {
		return err
	}
	return nil
}

func (m *MPD) WriteToString() (string, error) {
	b, err := xml.Marshal(m)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", xml.Header, b), nil
}

func NewMPD(profile DashProfile, mediaPresentationDuration string, minBufferTime string) *MPD {
	return &MPD{
		XMLNs:    Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles: Strptr((string)(profile)),
		Type:     Strptr("static"),
		MediaPresentationDuration: Strptr(mediaPresentationDuration),
		MinBufferTime:             Strptr(minBufferTime),
		Period:                    &Period{},
	}
}

func NewAdaptationSetAudio(segmentAlignment bool, startWithSAP int64, lang string) *AdaptationSet {
	return &AdaptationSet{
		MimeType:         Strptr("audio/mp4"),
		SegmentAlignment: Boolptr(segmentAlignment),
		StartWithSAP:     Intptr(startWithSAP),
		Lang:             Strptr(lang),
	}
}

func NewAdaptationSetVideo(scanType string, segmentAlignment bool, startWithSAP int64) *AdaptationSet {
	return &AdaptationSet{
		MimeType:         Strptr("video/mp4"),
		ScanType:         Strptr(scanType),
		SegmentAlignment: Boolptr(segmentAlignment),
		StartWithSAP:     Intptr(startWithSAP),
	}
}

func (m *MPD) AddAdaptationSet(as *AdaptationSet) {
	m.Period.AdaptationSets = append(m.Period.AdaptationSets, as)
}

func NewSegmentTemplate(duration int64, init string, media string, startNumber int64, timescale int64) *SegmentTemplate {
	return &SegmentTemplate{
		Duration:       Intptr(duration),
		Initialization: Strptr(init),
		Media:          Strptr(media),
		StartNumber:    Intptr(startNumber),
		Timescale:      Intptr(timescale),
	}
}

func (as *AdaptationSet) SetSegmentTemplate(st *SegmentTemplate, profile DashProfile) error {
	if profile != DASH_PROFILE_LIVE {
		return errors.New("Segment template can only be used with Live Profile")
	}
	if st == nil {
		return errors.New("nil Segment template")
	}
	as.SegmentTemplate = st
	return nil
}

func NewRepresentationAudio(samplingRate int64, bandwidth int64, codecs string, id string) *Representation {
	return &Representation{
		AudioSamplingRate: Intptr(samplingRate),
		Bandwidth:         Intptr(bandwidth),
		Codecs:            Strptr(codecs),
		ID:                Strptr(id),
	}
}

func NewRepresentationVideo(bandwidth int64, codecs string, id string, frameRate string, width int64, height int64) *Representation {
	return &Representation{
		Bandwidth: Intptr(bandwidth),
		Codecs:    Strptr(codecs),
		ID:        Strptr(id),
		FrameRate: Strptr(frameRate),
		Width:     Intptr(width),
		Height:    Intptr(height),
	}
}

func (r *Representation) SetBaseURL(baseURL string, profile DashProfile) error {
	if profile != DASH_PROFILE_ONDEMAND {
		return errors.New("Base URL can only be used with On-Demand Profile")
	}
	if baseURL == "" {
		return errors.New("Base URL empty")
	}
	r.BaseURL = Strptr(baseURL)
	return nil
}

func NewSegmentBase(indexRange string, init *Initialization) *SegmentBase {
	return &SegmentBase{
		IndexRange:     Strptr(indexRange),
		Initialization: init,
	}
}

func NewInitialization(initRange string) *Initialization {
	return &Initialization{
		Range: Strptr(initRange),
	}
}

func (r *Representation) SetSegmentBase(sb *SegmentBase, profile DashProfile) error {
	if profile != DASH_PROFILE_ONDEMAND {
		return errors.New("Segment Base can only be used with On-Demand Profile")
	}
	if sb == nil {
		return errors.New("Segment Base not set")
	}
	r.SegmentBase = sb
	return nil
}
