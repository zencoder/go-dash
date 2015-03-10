package mpd

import (
	"errors"

	. "github.com/zencoder/go-dash/helpers/ptrs"
)

type DashProfile string

const (
	DASH_PROFILE_LIVE     DashProfile = "urn:mpeg:dash:profile:isoff-live:2011"
	DASH_PROFILE_ONDEMAND DashProfile = "urn:mpeg:dash:profile:isoff-on-demand:2011"
)

var (
	ErrNoDASHProfileSet               error = errors.New("No DASH profile set")
	ErrAdaptationSetNil                     = errors.New("Adaptation Set nil")
	ErrSegmentTemplateLiveProfileOnly       = errors.New("Segment template can only be used with Live Profile")
	ErrSegmentTemplateNil                   = errors.New("Segment Template nil ")
	ErrRepresentationNil                    = errors.New("Representation nil")
	ErrBaseURLOnDemandProfileOnly           = errors.New("Base URL can only be used with On-Demand Profile")
	ErrBaseURLEmpty                         = errors.New("Base URL empty")
	ErrSegmentBaseOnDemandProfileOnly       = errors.New("Segment Base can only be used with On-Demand Profile")
	ErrSegmentBaseNil                       = errors.New("Segment Base nil")
)

type MPD struct {
	XMLNs                     *string `xml:"xmlns,attr"`
	Profiles                  *string `xml:"profiles,attr"`
	Type                      *string `xml:"type,attr"`
	MediaPresentationDuration *string `xml:"mediaPresentationDuration,attr"`
	MinBufferTime             *string `xml:"minBufferTime,attr"`
	Period                    *Period `xml:"Period,omitempty"`
}

type Period struct {
	AdaptationSets []*AdaptationSet `xml:"AdaptationSet,omitempty"`
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

// Live Profile Only
type SegmentTemplate struct {
	AdaptationSet  *AdaptationSet `xml:"-"`
	Duration       *int64         `xml:"duration,attr"`
	Initialization *string        `xml:"initialization,attr"`
	Media          *string        `xml:"media,attr"`
	StartNumber    *int64         `xml:"startNumber,attr"`
	Timescale      *int64         `xml:"timescale,attr"`
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

// On-Demand Profile
type SegmentBase struct {
	IndexRange     *string         `xml:"indexRange,attr"`
	Initialization *Initialization `xml:"Initialization,omitempty"`
}

// On-Demand Profile
type Initialization struct {
	Range *string `xml:"range,attr"`
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
		return ErrAdaptationSetNil
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
		return ErrNoDASHProfileSet
	}
	if *as.MPD.Profiles != (string)(DASH_PROFILE_LIVE) {
		return ErrSegmentTemplateLiveProfileOnly
	}
	if st == nil {
		return ErrSegmentTemplateNil
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
		return ErrRepresentationNil
	}
	r.AdaptationSet = as
	as.Representations = append(as.Representations, r)
	return nil
}

func (r *Representation) SetNewBaseURL(baseURL string) error {
	if r.AdaptationSet == nil || r.AdaptationSet.MPD == nil || r.AdaptationSet.MPD.Profiles == nil {
		return ErrNoDASHProfileSet
	}
	if *r.AdaptationSet.MPD.Profiles != (string)(DASH_PROFILE_ONDEMAND) {
		return ErrBaseURLOnDemandProfileOnly
	}
	if baseURL == "" {
		return ErrBaseURLEmpty
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
		return ErrNoDASHProfileSet
	}
	if *r.AdaptationSet.MPD.Profiles != (string)(DASH_PROFILE_ONDEMAND) {
		return ErrSegmentBaseOnDemandProfileOnly
	}
	if sb == nil {
		return ErrSegmentBaseNil
	}
	r.SegmentBase = sb
	return nil
}
