package mpd

import (
	"errors"
	"strings"

	. "github.com/zencoder/go-dash/helpers/ptrs"
)

// Type definition for DASH profiles
type DashProfile string

// Constants for supported DASH profiles
const (
	// Live Profile
	DASH_PROFILE_LIVE DashProfile = "urn:mpeg:dash:profile:isoff-live:2011"
	// On Demand Profile
	DASH_PROFILE_ONDEMAND DashProfile = "urn:mpeg:dash:profile:isoff-on-demand:2011"
)

// Constants for some known MIME types, this is a limited list and others can be used.
const (
	DASH_MIME_TYPE_VIDEO_MP4     string = "video/mp4"
	DASH_MIME_TYPE_AUDIO_MP4     string = "audio/mp4"
	DASH_MIME_TYPE_SUBTITLE_VTT  string = "text/vtt"
	DASH_MIME_TYPE_SUBTITLE_TTML string = "application/ttaf+xml"
	DASH_MIME_TYPE_SUBTITLE_SRT  string = "application/x-subrip"
	DASH_MIME_TYPE_SUBTITLE_DFXP string = "application/ttaf+xml"
)

// Known error variables
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
	BaseURL        string           `xml:"BaseURL,omitempty"`
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

// Constants for DRM / ContentProtection
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

// Segment Template is for Live Profile Only
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

// SegmentBase is for On-Demand Profile Only
type SegmentBase struct {
	IndexRange     *string         `xml:"indexRange,attr"`
	Initialization *Initialization `xml:"Initialization,omitempty"`
}

// Initialization is for On-Demand Profile Only
type Initialization struct {
	Range *string `xml:"range,attr"`
}

// Creates a new MPD object.
// profile - DASH Profile (Live or OnDemand).
// mediaPresentationDuration - Media Presentation Duration (i.e. PT6M16S).
// minBufferTime - Min Buffer Time (i.e. PT1.97S).
// baseURL - base url to use to fetch fragments. Nil will be ignored.
func NewMPD(profile DashProfile, mediaPresentationDuration string, minBufferTime string, baseURL string) *MPD {
	return &MPD{
		XMLNs:    Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles: Strptr((string)(profile)),
		Type:     Strptr("static"),
		MediaPresentationDuration: Strptr(mediaPresentationDuration),
		MinBufferTime:             Strptr(minBufferTime),
		Period:                    &Period{BaseURL: baseURL},
	}
}

// Create a new Adaptation Set for Audio Assets.
// mimeType - MIME Type (i.e. audio/mp4).
// segmentAlignment - Segment Alignment(i.e. true).
// startWithSAP - Starts With SAP (i.e. 1).
// lang - Language (i.e. en).
func (m *MPD) AddNewAdaptationSetAudio(mimeType string, segmentAlignment bool, startWithSAP int64, lang string) (*AdaptationSet, error) {
	as := &AdaptationSet{
		MimeType:         Strptr(mimeType),
		SegmentAlignment: Boolptr(segmentAlignment),
		StartWithSAP:     Intptr(startWithSAP),
		Lang:             Strptr(lang),
	}
	err := m.addAdaptationSet(as)
	if err != nil {
		return nil, err
	}
	return as, nil
}

// Create a new Adaptation Set for Video Assets.
// mimeType - MIME Type (i.e. video/mp4).
// scanType - Scan Type (i.e.progressive).
// segmentAlignment - Segment Alignment(i.e. true).
// startWithSAP - Starts With SAP (i.e. 1).
func (m *MPD) AddNewAdaptationSetVideo(mimeType string, scanType string, segmentAlignment bool, startWithSAP int64) (*AdaptationSet, error) {
	as := &AdaptationSet{
		MimeType:         Strptr(mimeType),
		ScanType:         Strptr(scanType),
		SegmentAlignment: Boolptr(segmentAlignment),
		StartWithSAP:     Intptr(startWithSAP),
	}
	err := m.addAdaptationSet(as)
	if err != nil {
		return nil, err
	}
	return as, nil
}

// Create a new Adaptation Set for Subtitle Assets.
// mimeType - MIME Type (i.e. text/vtt).
// lang - Language (i.e. en).
func (m *MPD) AddNewAdaptationSetSubtitle(mimeType string, lang string) (*AdaptationSet, error) {
	as := &AdaptationSet{
		MimeType: Strptr(mimeType),
		Lang:     Strptr(lang),
	}

	err := m.addAdaptationSet(as)
	if err != nil {
		return nil, err
	}
	return as, nil
}

// Internal helper method for adding a AdapatationSet to an MPD.
func (m *MPD) addAdaptationSet(as *AdaptationSet) error {
	if as == nil {
		return ErrAdaptationSetNil
	}
	as.MPD = m
	m.Period.AdaptationSets = append(m.Period.AdaptationSets, as)
	return nil
}

// Adds a ContentProtection tag at the root level of an AdaptationSet.
// This ContentProtection tag does not include signaling for any particular DRM scheme.
// defaultKIDHex - Default Key ID as a Hex String.
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

	err := as.addContentProtection(cp)
	if err != nil {
		return nil, err
	}

	return cp, nil
}

// Adds a new content protection scheme for Widevine DRM.
// This does not include the PSSH headers in the Manifest.
func (as *AdaptationSet) AddNewContentProtectionSchemeWidevine() (*ContentProtection, error) {
	cp := &ContentProtection{
		SchemeIDURI: Strptr(CONTENT_PROTECTION_WIDEVINE_SCHEME_ID),
	}

	err := as.addContentProtection(cp)
	if err != nil {
		return nil, err
	}

	return cp, nil
}

// Adds a new content protection scheme for PlayReady DRM.
// pro - PlayReady Object Header, as a Base64 encoded string.
func (as *AdaptationSet) AddNewContentProtectionSchemePlayready(pro string) (*ContentProtection, error) {
	if pro == "" {
		return nil, ErrPROEmpty
	}

	cp := &ContentProtection{
		SchemeIDURI:  Strptr(CONTENT_PROTECTION_PLAYREADY_SCHEME_ID),
		PlayreadyPRO: Strptr(pro),
	}

	err := as.addContentProtection(cp)
	if err != nil {
		return nil, err
	}

	return cp, nil
}

// Internal helper method for adding a ContentProtection to an AdaptationSet.
func (as *AdaptationSet) addContentProtection(cp *ContentProtection) error {
	if cp == nil {
		return ErrContentProtectionNil
	}

	as.ContentProtection = append(as.ContentProtection, cp)
	return nil
}

// Sets up a new SegmentTemplate for an AdaptationSet.
// duration - relative to timescale (i.e. 2000).
// init - template string for init segment (i.e. $RepresentationID$/audio/en/init.mp4).
// media - template string for media segments.
// startNumber - the number to start segments from ($Number$) (i.e. 0).
// timescale - sets the timescale for duration (i.e. 1000, represents milliseconds).
func (as *AdaptationSet) SetNewSegmentTemplate(duration int64, init string, media string, startNumber int64, timescale int64) (*SegmentTemplate, error) {
	st := &SegmentTemplate{
		Duration:       Intptr(duration),
		Initialization: Strptr(init),
		Media:          Strptr(media),
		StartNumber:    Intptr(startNumber),
		Timescale:      Intptr(timescale),
	}

	err := as.setSegmentTemplate(st)
	if err != nil {
		return nil, err
	}
	return st, nil
}

// Internal helper method for setting the Segment Template on an AdaptationSet.
func (as *AdaptationSet) setSegmentTemplate(st *SegmentTemplate) error {
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

// Adds a new Audio representation to an AdaptationSet.
// samplingRate - in Hz (i.e. 44100).
// bandwidth - in Bits/s (i.e. 67095).
// codecs - codec string for Audio Only (in RFC6381, https://tools.ietf.org/html/rfc6381) (i.e. mp4a.40.2).
// id - ID for this representation, will get used as $RepresentationID$ in template strings.
func (as *AdaptationSet) AddNewRepresentationAudio(samplingRate int64, bandwidth int64, codecs string, id string) (*Representation, error) {
	r := &Representation{
		AudioSamplingRate: Intptr(samplingRate),
		Bandwidth:         Intptr(bandwidth),
		Codecs:            Strptr(codecs),
		ID:                Strptr(id),
	}

	err := as.addRepresentation(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Adds a new Video representation to an AdaptationSet.
// bandwidth - in Bits/s (i.e. 1518664).
// codecs - codec string for Audio Only (in RFC6381, https://tools.ietf.org/html/rfc6381) (i.e. avc1.4d401f).
// id - ID for this representation, will get used as $RepresentationID$ in template strings.
// frameRate - video frame rate (as a fraction) (i.e. 30000/1001).
// width - width of the video (i.e. 1280).
// height - height of the video (i.e 720).
func (as *AdaptationSet) AddNewRepresentationVideo(bandwidth int64, codecs string, id string, frameRate string, width int64, height int64) (*Representation, error) {
	r := &Representation{
		Bandwidth: Intptr(bandwidth),
		Codecs:    Strptr(codecs),
		ID:        Strptr(id),
		FrameRate: Strptr(frameRate),
		Width:     Intptr(width),
		Height:    Intptr(height),
	}

	err := as.addRepresentation(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Adds a new Subtitle representation to an AdaptationSet.
// bandwidth - in Bits/s (i.e. 256).
// id - ID for this representation, will get used as $RepresentationID$ in template strings.
func (as *AdaptationSet) AddNewRepresentationSubtitle(bandwidth int64, id string) (*Representation, error) {
	r := &Representation{
		Bandwidth: Intptr(bandwidth),
		ID:        Strptr(id),
	}

	err := as.addRepresentation(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Internal helper method for adding a Representation to an AdaptationSet.
func (as *AdaptationSet) addRepresentation(r *Representation) error {
	if r == nil {
		return ErrRepresentationNil
	}
	r.AdaptationSet = as
	as.Representations = append(as.Representations, r)
	return nil
}

// Sets the BaseURL for a Representation.
// baseURL - Base URL as a string (i.e. 800k/output-audio-und.mp4)
func (r *Representation) SetNewBaseURL(baseURL string) error {
	if r.AdaptationSet == nil || r.AdaptationSet.MPD == nil || r.AdaptationSet.MPD.Profiles == nil {
		return ErrNoDASHProfileSet
	}
	if baseURL == "" {
		return ErrBaseURLEmpty
	}
	r.BaseURL = Strptr(baseURL)
	return nil
}

// Sets a new SegmentBase on a Respresentation.
// This is for On Demand profile.
// indexRange - Byte range to the index (sidx)atom.
// init - Byte range to the init atoms (ftyp+moov).
func (r *Representation) AddNewSegmentBase(indexRange string, initRange string) (*SegmentBase, error) {
	sb := &SegmentBase{
		IndexRange: Strptr(indexRange),
		Initialization: &Initialization{
			Range: Strptr(initRange),
		},
	}

	err := r.setSegmentBase(sb)
	if err != nil {
		return nil, err
	}
	return sb, nil
}

// Internal helper method for setting the SegmentBase on a Respresentation.
func (r *Representation) setSegmentBase(sb *SegmentBase) error {
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
