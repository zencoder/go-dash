package mpd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"os"

	. "github.com/zencoder/go-dash/helpers/ptrs"
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
	Period                    *Period `xml:"Period,omitempty"`
}

func (m MPD) String() string {
	jb, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(jb)
}

type Period struct {
	AdaptationSets []*AdaptationSet `xml:"AdaptationSet,omitempty"`
}

func (p Period) String() string {
	jb, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(jb)
}

type AdaptationSet struct {
	MPD              *MPD              `xml:"-"`
	MimeType         *string           `xml:"mimeType,attr"`
	ScanType         *string           `xml:"scanType,attr"`
	SegmentAlignment *bool             `xml:"segmentAlignment,attr"`
	StartWithSAP     *int64            `xml:"startWithSAP,attr"`
	Lang             *string           `xml:"lang,attr"`
	SegmentTemplate  *SegmentTemplate  `xml:"SegmentTemplate,omitempty"` // Live Profile Only
	Representations  []*Representation `xml:"Representation,omitempty"`
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
	AdaptationSet  *AdaptationSet `xml:"-"`
	Duration       *int64         `xml:"duration,attr"`
	Initialization *string        `xml:"initialization,attr"`
	Media          *string        `xml:"media,attr"`
	StartNumber    *int64         `xml:"startNumber,attr"`
	Timescale      *int64         `xml:"timescale,attr"`
}

func (st SegmentTemplate) String() string {
	jb, err := json.Marshal(st)
	if err != nil {
		return ""
	}
	return string(jb)
}

type Representation struct {
	AdaptationSet     *AdaptationSet `xml:"-"`
	AudioSamplingRate *int64         `xml:"audioSamplingRate,attr"` // Audio
	Bandwidth         *int64         `xml:"bandwidth,attr"`         // Audio + Video
	Codecs            *string        `xml:"codecs,attr"`            // Audio + Video
	FrameRate         *string        `xml:"frameRate,attr"`         // Video
	Height            *int64         `xml:"height,attr"`            // Video
	ID                *string        `xml:"id,attr"`                // Audio + Video
	Width             *int64         `xml:"width,attr"`             // Video
	BaseURL           *string        `xml:"BaseURL,omitempty"`      // On-Demand Profile
	SegmentBase       *SegmentBase   `xml:"SegmentBase,omitempty"`  // On-Demand Profile
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
	Initialization *Initialization `xml:"Initialization,omitempty"`
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

func ReadFromFile(path string) (*MPD, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return Read(f)
}

func ReadFromString(xmlStr string) (*MPD, error) {
	b := bytes.NewBufferString(xmlStr)
	return Read(b)
}

func Read(r io.Reader) (*MPD, error) {
	var mpd MPD
	d := xml.NewDecoder(r)
	err := d.Decode(&mpd)
	if err != nil {
		return nil, err
	}
	return &mpd, nil
}

func (m *MPD) WriteToFile(path string) error {
	// Open the file to write the XML to
	f, err := os.OpenFile(path, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	if err = m.Write(f); err != nil {
		return err
	}
	if err = f.Sync(); err != nil {
		return err
	}
	return err
}

func (m *MPD) WriteToString() (string, error) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	err := m.Write(w)
	if err != nil {
		return "", err
	}
	err = w.Flush()
	if err != nil {
		return "", err
	}
	return b.String(), err
}

func (m *MPD) Write(w io.Writer) error {
	// Write out the XML Header
	w.Write([]byte(xml.Header))
	// Write out the DASH XML manifest
	e := xml.NewEncoder(w)
	e.Indent("", "  ")
	err := e.Encode(m)
	if err != nil {
		return err
	}
	w.Write([]byte("\n"))
	return nil
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

func (m *MPD) AddNewAdaptationSetAudio(segmentAlignment bool, startWithSAP int64, lang string) (*AdaptationSet, error) {
	as := &AdaptationSet{
		MimeType:         Strptr("audio/mp4"),
		SegmentAlignment: Boolptr(segmentAlignment),
		StartWithSAP:     Intptr(startWithSAP),
		Lang:             Strptr(lang),
	}
	err := m.AddAdaptationSet(as)
	if err != nil {
		return nil, err
	}
	return as, nil
}

func (m *MPD) AddNewAdaptationSetVideo(scanType string, segmentAlignment bool, startWithSAP int64) (*AdaptationSet, error) {
	as := &AdaptationSet{
		MimeType:         Strptr("video/mp4"),
		ScanType:         Strptr(scanType),
		SegmentAlignment: Boolptr(segmentAlignment),
		StartWithSAP:     Intptr(startWithSAP),
	}
	err := m.AddAdaptationSet(as)
	if err != nil {
		return nil, err
	}
	return as, nil
}

func (m *MPD) AddAdaptationSet(as *AdaptationSet) error {
	if as == nil {
		return errors.New("Adaptation set is nil")
	}
	as.MPD = m
	m.Period.AdaptationSets = append(m.Period.AdaptationSets, as)
	return nil
}

func (as *AdaptationSet) SetNewSegmentTemplate(duration int64, init string, media string, startNumber int64, timescale int64) (*SegmentTemplate, error) {
	st := &SegmentTemplate{
		Duration:       Intptr(duration),
		Initialization: Strptr(init),
		Media:          Strptr(media),
		StartNumber:    Intptr(startNumber),
		Timescale:      Intptr(timescale),
	}

	err := as.SetSegmentTemplate(st)
	if err != nil {
		return nil, err
	}
	return st, nil
}

func (as *AdaptationSet) SetSegmentTemplate(st *SegmentTemplate) error {
	if as.MPD == nil || as.MPD.Profiles == nil {
		return errors.New("No DASH profile set")
	}
	if *as.MPD.Profiles != (string)(DASH_PROFILE_LIVE) {
		return errors.New("Segment template can only be used with Live Profile")
	}
	if st == nil {
		return errors.New("nil Segment template")
	}
	st.AdaptationSet = as
	as.SegmentTemplate = st
	return nil
}

func (as *AdaptationSet) AddNewRepresentationAudio(samplingRate int64, bandwidth int64, codecs string, id string) (*Representation, error) {
	r := &Representation{
		AudioSamplingRate: Intptr(samplingRate),
		Bandwidth:         Intptr(bandwidth),
		Codecs:            Strptr(codecs),
		ID:                Strptr(id),
	}

	err := as.AddRepresentation(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (as *AdaptationSet) AddNewRepresentationVideo(bandwidth int64, codecs string, id string, frameRate string, width int64, height int64) (*Representation, error) {
	r := &Representation{
		Bandwidth: Intptr(bandwidth),
		Codecs:    Strptr(codecs),
		ID:        Strptr(id),
		FrameRate: Strptr(frameRate),
		Width:     Intptr(width),
		Height:    Intptr(height),
	}

	err := as.AddRepresentation(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (as *AdaptationSet) AddRepresentation(r *Representation) error {
	if r == nil {
		return errors.New("Representation is nil")
	}
	r.AdaptationSet = as
	as.Representations = append(as.Representations, r)
	return nil
}

func (r *Representation) SetNewBaseURL(baseURL string) error {
	if r.AdaptationSet == nil || r.AdaptationSet.MPD == nil || r.AdaptationSet.MPD.Profiles == nil {
		return errors.New("No DASH profile set")
	}
	if *r.AdaptationSet.MPD.Profiles != (string)(DASH_PROFILE_ONDEMAND) {
		return errors.New("Base URL can only be used with On-Demand Profile")
	}
	if baseURL == "" {
		return errors.New("Base URL empty")
	}
	r.BaseURL = Strptr(baseURL)
	return nil
}

func (r *Representation) AddNewSegmentBase(indexRange string, init string) (*SegmentBase, error) {
	sb := &SegmentBase{
		IndexRange: Strptr(indexRange),
		Initialization: &Initialization{
			Range: Strptr(init),
		},
	}

	err := r.SetSegmentBase(sb)
	if err != nil {
		return nil, err
	}
	return sb, nil
}

func (r *Representation) SetSegmentBase(sb *SegmentBase) error {
	if r.AdaptationSet == nil || r.AdaptationSet.MPD == nil || r.AdaptationSet.MPD.Profiles == nil {
		return errors.New("No DASH profile set")
	}
	if *r.AdaptationSet.MPD.Profiles != (string)(DASH_PROFILE_ONDEMAND) {
		return errors.New("Segment Base can only be used with On-Demand Profile")
	}
	if sb == nil {
		return errors.New("Segment Base not set")
	}
	r.SegmentBase = sb
	return nil
}
