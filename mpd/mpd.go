package mpd

import (
	"errors"
	"strings"

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
	ErrInvalidDefaultKID                    = errors.New("Invalid Default KID string, should be 32 characters")
	ErrPROEmpty                             = errors.New("PlayReady PRO empty")
	ErrContentProtectionNil                 = errors.New("Content Protection nil")
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
	MPD               *MPD                 `xml:"-"`
	MimeType          *string              `xml:"mimeType,attr"`
	ScanType          *string              `xml:"scanType,attr"`
	SegmentAlignment  *bool                `xml:"segmentAlignment,attr"`
	StartWithSAP      *int64               `xml:"startWithSAP,attr"`
	Lang              *string              `xml:"lang,attr"`
	ContentProtection []*ContentProtection `xml:"ContentProtection,omitempty"`
	SegmentTemplate   *SegmentTemplate     `xml:"SegmentTemplate,omitempty"` // Live Profile Only
	Representations   []*Representation    `xml:"Representation,omitempty"`
}

/*
<ContentProtection cenc:default_KID="09e36702-8f33-436c-a5dd60ffe6671e70" schemeIdUri="urn:mpeg:dash:mp4protection:2011" value="cenc" xmlns:cenc="urn:mpeg:cenc:2013"/>
<ContentProtection schemeIdUri="urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed"/>
<ContentProtection schemeIdUri="urn:uuid:9a04f079-9840-4286-ab92-e65be0885f95">
	<mspr:pro>mgIAAAEAAQCQAjwAVwBSAE0ASABFAEEARABFAFIAIAB4AG0AbABuAHMAPQAiAGgAdAB0AHAAOgAvAC8AcwBjAGgAZQBtAGEAcwAuAG0AaQBjAHIAbwBzAG8AZgB0AC4AYwBvAG0ALwBEAFIATQAvADIAMAAwADcALwAwADMALwBQAGwAYQB5AFIAZQBhAGQAeQBIAGUAYQBkAGUAcgAiACAAdgBlAHIAcwBpAG8AbgA9ACIANAAuADAALgAwAC4AMAAiAD4APABEAEEAVABBAD4APABQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsARQBZAEwARQBOAD4AMQA2ADwALwBLAEUAWQBMAEUATgA+ADwAQQBMAEcASQBEAD4AQQBFAFMAQwBUAFIAPAAvAEEATABHAEkARAA+ADwALwBQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsASQBEAD4AQQBtAGYAagBDAFQATwBQAGIARQBPAGwAMwBXAEQALwA1AG0AYwBlAGMAQQA9AD0APAAvAEsASQBEAD4APABDAEgARQBDAEsAUwBVAE0APgBCAEcAdwAxAGEAWQBaADEAWQBYAE0APQA8AC8AQwBIAEUAQwBLAFMAVQBNAD4APABMAEEAXwBVAFIATAA+AGgAdAB0AHAAOgAvAC8AcABsAGEAeQByAGUAYQBkAHkALgBkAGkAcgBlAGMAdAB0AGEAcABzAC4AbgBlAHQALwBwAHIALwBzAHYAYwAvAHIAaQBnAGgAdABzAG0AYQBuAGEAZwBlAHIALgBhAHMAbQB4ADwALwBMAEEAXwBVAFIATAA+ADwALwBEAEEAVABBAD4APAAvAFcAUgBNAEgARQBBAEQARQBSAD4A</mspr:pro>
</ContentProtection>
*/

const (
	CONTENT_PROTECTION_ROOT_SCHEME_ID_URI  = "urn:mpeg:dash:mp4protection:2011"
	CONTENT_PROTECTION_ROOT_VALUE          = "cenc"
	CONTENT_PROTECTION_ROOT_XMLNS          = "urn:mpeg:cenc:2013"
	CONTENT_PROTECTION_WIDEVINE_SCHEME_ID  = "urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed"
	CONTENT_PROTECTION_PLAYREADY_SCHEME_ID = "urn:uuid:9a04f079-9840-4286-ab92-e65be0885f95"
)

type ContentProtection struct {
	AdaptationSet *AdaptationSet `xml:"-"`
	DefaultKID    *string        `xml:"cenc:default_KID,attr"`
	SchemeIDURI   *string        `xml:"schemeIdUri,attr"` // Default: urn:mpeg:dash:mp4protection:2011
	Value         *string        `xml:"value,attr"`       // Default: cenc
	XMLNS         *string        `xml:"xmlns:cenc,attr"`  // Default: urn:mpeg:cenc:2013
	PlayreadyPRO  *string        `xml:"mspr:pro,omitempty"`
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

/*
<ContentProtection cenc:default_KID="09e36702-8f33-436c-a5dd60ffe6671e70" schemeIdUri="urn:mpeg:dash:mp4protection:2011" value="cenc" xmlns:cenc="urn:mpeg:cenc:2013"/>
<ContentProtection schemeIdUri="urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed"/>
<ContentProtection schemeIdUri="urn:uuid:9a04f079-9840-4286-ab92-e65be0885f95">
	<mspr:pro>mgIAAAEAAQCQAjwAVwBSAE0ASABFAEEARABFAFIAIAB4AG0AbABuAHMAPQAiAGgAdAB0AHAAOgAvAC8AcwBjAGgAZQBtAGEAcwAuAG0AaQBjAHIAbwBzAG8AZgB0AC4AYwBvAG0ALwBEAFIATQAvADIAMAAwADcALwAwADMALwBQAGwAYQB5AFIAZQBhAGQAeQBIAGUAYQBkAGUAcgAiACAAdgBlAHIAcwBpAG8AbgA9ACIANAAuADAALgAwAC4AMAAiAD4APABEAEEAVABBAD4APABQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsARQBZAEwARQBOAD4AMQA2ADwALwBLAEUAWQBMAEUATgA+ADwAQQBMAEcASQBEAD4AQQBFAFMAQwBUAFIAPAAvAEEATABHAEkARAA+ADwALwBQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsASQBEAD4AQQBtAGYAagBDAFQATwBQAGIARQBPAGwAMwBXAEQALwA1AG0AYwBlAGMAQQA9AD0APAAvAEsASQBEAD4APABDAEgARQBDAEsAUwBVAE0APgBCAEcAdwAxAGEAWQBaADEAWQBYAE0APQA8AC8AQwBIAEUAQwBLAFMAVQBNAD4APABMAEEAXwBVAFIATAA+AGgAdAB0AHAAOgAvAC8AcABsAGEAeQByAGUAYQBkAHkALgBkAGkAcgBlAGMAdAB0AGEAcABzAC4AbgBlAHQALwBwAHIALwBzAHYAYwAvAHIAaQBnAGgAdABzAG0AYQBuAGEAZwBlAHIALgBhAHMAbQB4ADwALwBMAEEAXwBVAFIATAA+ADwALwBEAEEAVABBAD4APAAvAFcAUgBNAEgARQBBAEQARQBSAD4A</mspr:pro>
</ContentProtection>


type ContentProtection struct {
	AdaptationSet *AdaptationSet `xml:"-"`
	DefaultKID    *string        `xml:"cenc:default_KID,attr"`
	SchemeIDURI   *string        `xml:"schemeIdUri,attr"` // Default: urn:mpeg:dash:mp4protection:2011
	Value         *string        `xml:"value,attr"`       // Default: cenc
	XMLNS         *string        `xml:"xmlns:cenc,attr"`  // Default: urn:mpeg:cenc:2013
	PlayReadyPRO  *string        `xml:"mspr:pro,omitempty"`
}

const (
	CONTENT_PROTECTION_ROOT_SCHEME_ID_URI = "urn:mpeg:dash:mp4protection:2011"
	CONTENT_PROTECTION_ROOT_VALUE = "cenc"
	CONTENT_PROTECTION_ROOT_XMLNS = "urn:mpeg:cenc:2013"
	)
*/

func (as *AdaptationSet) AddNewContentProtectionRoot(defaultKIDHex string) (*ContentProtection, error) {
	if len(defaultKIDHex) != 32 || defaultKIDHex == "" {
		return nil, ErrInvalidDefaultKID
	}

	// Convert the KID into the correct format
	defaultKID := strings.ToLower(defaultKIDHex[0:8] + "-" + defaultKIDHex[8:12] + "-" + defaultKIDHex[12:16] + "-" + defaultKIDHex[16:32])

	cp := &ContentProtection{
		DefaultKID:  Strptr(defaultKID),
		SchemeIDURI: Strptr(CONTENT_PROTECTION_ROOT_SCHEME_ID_URI),
		Value:       Strptr(CONTENT_PROTECTION_ROOT_VALUE),
		XMLNS:       Strptr(CONTENT_PROTECTION_ROOT_XMLNS),
	}

	err := as.AddContentProtection(cp)
	if err != nil {
		return nil, err
	}

	return cp, nil
}

func (as *AdaptationSet) AddNewContentProtectionSchemeWidevine() (*ContentProtection, error) {
	cp := &ContentProtection{
		SchemeIDURI: Strptr(CONTENT_PROTECTION_WIDEVINE_SCHEME_ID),
	}

	err := as.AddContentProtection(cp)
	if err != nil {
		return nil, err
	}

	return cp, nil
}

func (as *AdaptationSet) AddNewContentProtectionSchemePlayready(pro string) (*ContentProtection, error) {
	if pro == "" {
		return nil, ErrPROEmpty
	}

	cp := &ContentProtection{
		SchemeIDURI:  Strptr(CONTENT_PROTECTION_PLAYREADY_SCHEME_ID),
		PlayreadyPRO: Strptr(pro),
	}

	err := as.AddContentProtection(cp)
	if err != nil {
		return nil, err
	}

	return cp, nil
}

func (as *AdaptationSet) AddContentProtection(cp *ContentProtection) error {
	if cp == nil {
		return ErrContentProtectionNil
	}

	as.ContentProtection = append(as.ContentProtection, cp)
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
