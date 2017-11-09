package mpd

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"strings"
	"time"

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
	// HbbTV Profile
	DASH_PROFILE_HBBTV_1_5_LIVE DashProfile = "urn:hbbtv:dash:profile:isoff-live:2012,urn:mpeg:dash:profile:isoff-live:2011"
)

type AudioChannelConfigurationScheme string

const (
	// Scheme for non-Dolby Audio
	AUDIO_CHANNEL_CONFIGURATION_MPEG_DASH AudioChannelConfigurationScheme = "urn:mpeg:dash:23003:3:audio_channel_configuration:2011"
	// Scheme for Dolby Audio
	AUDIO_CHANNEL_CONFIGURATION_MPEG_DOLBY AudioChannelConfigurationScheme = "tag:dolby.com,2014:dash:audio_channel_configuration:2011"
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
	ErrNoDASHProfileSet             error = errors.New("No DASH profile set")
	ErrAdaptationSetNil                   = errors.New("Adaptation Set nil")
	ErrSegmentTemplateNil                 = errors.New("Segment Template nil ")
	ErrRepresentationNil                  = errors.New("Representation nil")
	ErrBaseURLEmpty                       = errors.New("Base URL empty")
	ErrSegmentBaseNil                     = errors.New("Segment Base nil")
	ErrAudioChannelConfigurationNil       = errors.New("Audio Channel Configuration nil")
	ErrInvalidDefaultKID                  = errors.New("Invalid Default KID string, should be 32 characters")
	ErrPROEmpty                           = errors.New("PlayReady PRO empty")
)

type MPD struct {
	XMLNs                     *string `xml:"xmlns,attr"`
	Profiles                  *string `xml:"profiles,attr"`
	Type                      *string `xml:"type,attr"`
	MediaPresentationDuration *string `xml:"mediaPresentationDuration,attr"`
	MinBufferTime             *string `xml:"minBufferTime,attr"`
	BaseURL                   string  `xml:"BaseURL,omitempty"`
	period                    *Period
	Periods                   []*Period `xml:"Period,omitempty"`
}

type Period struct {
	ID              string           `xml:"id,attr,omitempty"`
	Duration        Duration         `xml:"duration,attr,omitempty"`
	Start           Duration         `xml:"start,attr,omitempty"`
	BaseURL         string           `xml:"BaseURL,omitempty"`
	SegmentBase     *SegmentBase     `xml:"SegmentBase,omitempty"`
	SegmentList     *SegmentList     `xml:"SegmentList,omitempty"`
	SegmentTemplate *SegmentTemplate `xml:"SegmentTemplate,omitempty"`
	AdaptationSets  []*AdaptationSet `xml:"AdaptationSet,omitempty"`
}

type AdaptationSet struct {
	MimeType          *string             `xml:"mimeType,attr"`
	ScanType          *string             `xml:"scanType,attr"`
	SegmentAlignment  *bool               `xml:"segmentAlignment,attr"`
	StartWithSAP      *int64              `xml:"startWithSAP,attr"`
	Lang              *string             `xml:"lang,attr"`
	ContentProtection []ContentProtection `xml:"ContentProtection,omitempty"`
	Roles             []*Role             `xml:"Role,omitempty"`
	SegmentBase       *SegmentBase        `xml:"SegmentBase,omitempty"`
	SegmentList       *SegmentList        `xml:"SegmentList,omitempty"`
	SegmentTemplate   *SegmentTemplate    `xml:"SegmentTemplate,omitempty"` // Live Profile Only
	Representations   []*Representation   `xml:"Representation,omitempty"`
}

// Constants for DRM / ContentProtection
const (
	CONTENT_PROTECTION_ROOT_SCHEME_ID_URI       = "urn:mpeg:dash:mp4protection:2011"
	CONTENT_PROTECTION_ROOT_VALUE               = "cenc"
	CENC_XMLNS                                  = "urn:mpeg:cenc:2013"
	CONTENT_PROTECTION_WIDEVINE_SCHEME_ID       = "urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed"
	CONTENT_PROTECTION_WIDEVINE_SCHEME_HEX      = "edef8ba979d64acea3c827dcd51d21ed"
	CONTENT_PROTECTION_PLAYREADY_SCHEME_ID      = "urn:uuid:9a04f079-9840-4286-ab92-e65be0885f95"
	CONTENT_PROTECTION_PLAYREADY_SCHEME_HEX     = "9a04f07998404286ab92e65be0885f95"
	CONTENT_PROTECTION_PLAYREADY_SCHEME_V10_ID  = "urn:uuid:79f0049a-4098-8642-ab92-e65be0885f95"
	CONTENT_PROTECTION_PLAYREADY_SCHEME_V10_HEX = "79f0049a40988642ab92e65be0885f95"
	CONTENT_PROTECTION_PLAYREADY_XMLNS          = "urn:microsoft:playready"
)

type ContentProtection struct {
	AdaptationSet *AdaptationSet `xml:"-"`
	XMLName       xml.Name       `xml:"ContentProtection"`
	SchemeIDURI   *string        `xml:"schemeIdUri,attr"` // Default: urn:mpeg:dash:mp4protection:2011
	XMLNS         *string        `xml:"xmlns:cenc,attr"`  // Default: urn:mpeg:cenc:2013

	// CENC
	DefaultKID *string `xml:"cenc:default_KID,attr"`
	Value      *string `xml:"value,attr"` // Default: cenc

	// Playready
	PlayreadyXMLNS *string `xml:"xmlns:mspr,attr,omitempty"`
	PRO            *string `xml:"mspr:pro,omitempty"`

	// Playready / Widevine
	PSSH *string `xml:"cenc:pssh,omitempty"`

	// For proprietary DRM
	CustomAttrs []xml.Attr `xml:",attr,all"`
}

type Role struct {
	AdaptationSet *AdaptationSet `xml:"-"`
	SchemeIDURI   *string        `xml:"schemeIdUri,attr"`
	Value         *string        `xml:"value,attr"`
}

// Segment Template is for Live Profile Only
type SegmentTemplate struct {
	AdaptationSet          *AdaptationSet   `xml:"-"`
	SegmentTimeline        *SegmentTimeline `xml:"SegmentTimeline,omitempty"`
	PresentationTimeOffset *uint64          `xml:"presentationTimeOffset,attr,omitempty"`
	Duration               *int64           `xml:"duration,attr"`
	Initialization         *string          `xml:"initialization,attr"`
	Media                  *string          `xml:"media,attr"`
	StartNumber            *int64           `xml:"startNumber,attr"`
	Timescale              *int64           `xml:"timescale,attr"`
}

type Representation struct {
	AdaptationSet             *AdaptationSet             `xml:"-"`
	AudioChannelConfiguration *AudioChannelConfiguration `xml:"AudioChannelConfiguration,omitempty"`
	AudioSamplingRate         *int64                     `xml:"audioSamplingRate,attr"`   // Audio
	Bandwidth                 *int64                     `xml:"bandwidth,attr"`           // Audio + Video
	Codecs                    *string                    `xml:"codecs,attr"`              // Audio + Video
	FrameRate                 *string                    `xml:"frameRate,attr,omitempty"` // Video
	Height                    *int64                     `xml:"height,attr"`              // Video
	ID                        *string                    `xml:"id,attr"`                  // Audio + Video
	Width                     *int64                     `xml:"width,attr"`               // Video
	BaseURL                   *string                    `xml:"BaseURL,omitempty"`        // On-Demand Profile
	SegmentBase               *SegmentBase               `xml:"SegmentBase,omitempty"`    // On-Demand Profile
	SegmentList               *SegmentList               `xml:"SegmentList,omitempty"`
	SegmentTemplate           *SegmentTemplate           `xml:"SegmentTemplate,omitempty"`
}

type AudioChannelConfiguration struct {
	SchemeIDURI *string `xml:"schemeIdUri,attr"`
	// Value will be an int for non-Dolby Schemes, and a hexstring for Dolby Schemes, hence we make it a string
	Value *string `xml:"value,attr"`
}

// Creates a new MPD object.
// profile - DASH Profile (Live or OnDemand).
// mediaPresentationDuration - Media Presentation Duration (i.e. PT6M16S).
// minBufferTime - Min Buffer Time (i.e. PT1.97S).
func NewMPD(profile DashProfile, mediaPresentationDuration string, minBufferTime string) *MPD {
	period := &Period{}
	return &MPD{
		XMLNs:    Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles: Strptr((string)(profile)),
		Type:     Strptr("static"),
		MediaPresentationDuration: Strptr(mediaPresentationDuration),
		MinBufferTime:             Strptr(minBufferTime),
		period:                    period,
		Periods:                   []*Period{period},
	}
}

// AddNewPeriod creates a new Period and make it the currently active one.
func (m *MPD) AddNewPeriod() *Period {
	period := &Period{}
	m.Periods = append(m.Periods, period)
	m.period = period
	return period
}

// GetCurrentPeriod returns the current Period.
func (m *MPD) GetCurrentPeriod() *Period {
	return m.period
}

func (period *Period) SetDuration(d time.Duration) {
	period.Duration = Duration(d)
}

// Create a new Adaptation Set for Audio Assets.
// mimeType - MIME Type (i.e. audio/mp4).
// segmentAlignment - Segment Alignment(i.e. true).
// startWithSAP - Starts With SAP (i.e. 1).
// lang - Language (i.e. en).
func (m *MPD) AddNewAdaptationSetAudio(mimeType string, segmentAlignment bool, startWithSAP int64, lang string) (*AdaptationSet, error) {
	return m.period.AddNewAdaptationSetAudio(mimeType, segmentAlignment, startWithSAP, lang)
}

// Create a new Adaptation Set for Audio Assets.
// mimeType - MIME Type (i.e. audio/mp4).
// segmentAlignment - Segment Alignment(i.e. true).
// startWithSAP - Starts With SAP (i.e. 1).
// lang - Language (i.e. en).
func (period *Period) AddNewAdaptationSetAudio(mimeType string, segmentAlignment bool, startWithSAP int64, lang string) (*AdaptationSet, error) {
	as := &AdaptationSet{
		MimeType:         Strptr(mimeType),
		SegmentAlignment: Boolptr(segmentAlignment),
		StartWithSAP:     Int64ptr(startWithSAP),
		Lang:             Strptr(lang),
	}
	err := period.addAdaptationSet(as)
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
	return m.period.AddNewAdaptationSetVideo(mimeType, scanType, segmentAlignment, startWithSAP)
}

// Create a new Adaptation Set for Video Assets.
// mimeType - MIME Type (i.e. video/mp4).
// scanType - Scan Type (i.e.progressive).
// segmentAlignment - Segment Alignment(i.e. true).
// startWithSAP - Starts With SAP (i.e. 1).
func (period *Period) AddNewAdaptationSetVideo(mimeType string, scanType string, segmentAlignment bool, startWithSAP int64) (*AdaptationSet, error) {
	as := &AdaptationSet{
		MimeType:         Strptr(mimeType),
		ScanType:         Strptr(scanType),
		SegmentAlignment: Boolptr(segmentAlignment),
		StartWithSAP:     Int64ptr(startWithSAP),
	}
	err := period.addAdaptationSet(as)
	if err != nil {
		return nil, err
	}
	return as, nil
}

// Create a new Adaptation Set for Subtitle Assets.
// mimeType - MIME Type (i.e. text/vtt).
// lang - Language (i.e. en).
func (m *MPD) AddNewAdaptationSetSubtitle(mimeType string, lang string) (*AdaptationSet, error) {
	return m.period.AddNewAdaptationSetSubtitle(mimeType, lang)
}

// Create a new Adaptation Set for Subtitle Assets.
// mimeType - MIME Type (i.e. text/vtt).
// lang - Language (i.e. en).
func (period *Period) AddNewAdaptationSetSubtitle(mimeType string, lang string) (*AdaptationSet, error) {
	as := &AdaptationSet{
		MimeType: Strptr(mimeType),
		Lang:     Strptr(lang),
	}
	err := period.addAdaptationSet(as)
	if err != nil {
		return nil, err
	}
	return as, nil
}

// Internal helper method for adding a AdapatationSet.
func (period *Period) addAdaptationSet(as *AdaptationSet) error {
	if as == nil {
		return ErrAdaptationSetNil
	}
	period.AdaptationSets = append(period.AdaptationSets, as)
	return nil
}

// Adds a ContentProtection tag at the root level of an AdaptationSet.
// This ContentProtection tag does not include signaling for any particular DRM scheme.
// defaultKIDHex - Default Key ID as a Hex String.
func (as *AdaptationSet) AddNewContentProtectionRoot(defaultKIDHex string) (ContentProtection, error) {
	if len(defaultKIDHex) != 32 || defaultKIDHex == "" {
		return ContentProtection{}, ErrInvalidDefaultKID
	}

	// Convert the KID into the correct format
	defaultKID := strings.ToLower(defaultKIDHex[0:8] + "-" + defaultKIDHex[8:12] + "-" + defaultKIDHex[12:16] + "-" + defaultKIDHex[16:32])

	cp := ContentProtection{
		DefaultKID: Strptr(defaultKID),
		Value:      Strptr(CONTENT_PROTECTION_ROOT_VALUE),
	}
	cp.SchemeIDURI = Strptr(CONTENT_PROTECTION_ROOT_SCHEME_ID_URI)
	cp.XMLNS = Strptr(CENC_XMLNS)

	as.AddContentProtection(cp)
	return cp, nil
}

// AddNewContentProtectionSchemeWidevine adds a new content protection scheme for Widevine DRM to the adaptation set. With
// a <cenc:pssh> element that contains a Base64 encoded PSSH box
// wvHeader - binary representation of Widevine Header
// !!! Note: this function will accept any byte slice as a wvHeader value !!!
func (as *AdaptationSet) AddNewContentProtectionSchemeWidevineWithPSSH(wvHeader []byte) (ContentProtection, error) {
	cp, err := NewWidevineContentProtection(wvHeader)
	if err != nil {
		return ContentProtection{}, err
	}

	as.AddContentProtection(cp)
	return cp, nil
}

// AddNewContentProtectionSchemeWidevine adds a new content protection scheme for Widevine DRM to the adaptation set.
func (as *AdaptationSet) AddNewContentProtectionSchemeWidevine() (ContentProtection, error) {
	cp, err := NewWidevineContentProtection(nil)
	if err != nil {
		return ContentProtection{}, err
	}

	as.AddContentProtection(cp)
	return cp, nil
}

func NewWidevineContentProtection(wvHeader []byte) (ContentProtection, error) {
	cp := ContentProtection{}
	cp.SchemeIDURI = Strptr(CONTENT_PROTECTION_WIDEVINE_SCHEME_ID)

	if len(wvHeader) > 0 {
		cp.XMLNS = Strptr(CENC_XMLNS)
		wvSystemID, err := hex.DecodeString(CONTENT_PROTECTION_WIDEVINE_SCHEME_HEX)
		if err != nil {
			return ContentProtection{}, err
		}
		psshBox, err := makePSSHBox(wvSystemID, wvHeader)
		if err != nil {
			return ContentProtection{}, err
		}

		psshB64 := base64.StdEncoding.EncodeToString(psshBox)
		cp.PSSH = &psshB64
	}
	return cp, nil
}

// AddNewContentProtectionSchemePlayready adds a new content protection scheme for PlayReady DRM.
// pro - PlayReady Object Header, as a Base64 encoded string.
func (as *AdaptationSet) AddNewContentProtectionSchemePlayready(pro string) (ContentProtection, error) {
	cp, err := newPlayreadyContentProtection(pro, CONTENT_PROTECTION_PLAYREADY_SCHEME_ID)
	if err != nil {
		return ContentProtection{}, err
	}

	as.AddContentProtection(cp)
	return cp, nil
}

// AddNewContentProtectionSchemePlayreadyV10 adds a new content protection scheme for PlayReady v1.0 DRM.
// pro - PlayReady Object Header, as a Base64 encoded string.
func (as *AdaptationSet) AddNewContentProtectionSchemePlayreadyV10(pro string) (ContentProtection, error) {
	cp, err := newPlayreadyContentProtection(pro, CONTENT_PROTECTION_PLAYREADY_SCHEME_V10_ID)
	if err != nil {
		return ContentProtection{}, err
	}

	as.AddContentProtection(cp)
	return cp, nil
}

func newPlayreadyContentProtection(pro string, schemeIDURI string) (ContentProtection, error) {
	if pro == "" {
		return ContentProtection{}, ErrPROEmpty
	}

	cp := ContentProtection{
		PlayreadyXMLNS: Strptr(CONTENT_PROTECTION_PLAYREADY_XMLNS),
		PRO:            Strptr(pro),
	}
	cp.SchemeIDURI = Strptr(schemeIDURI)

	return cp, nil
}

// AddNewContentProtectionSchemePlayreadyWithPSSH adds a new content protection scheme for PlayReady DRM. The scheme
// will include both ms:pro and cenc:pssh subelements
// pro - PlayReady Object Header, as a Base64 encoded string.
func (as *AdaptationSet) AddNewContentProtectionSchemePlayreadyWithPSSH(pro string) (ContentProtection, error) {
	cp, err := newPlayreadyContentProtection(pro, CONTENT_PROTECTION_PLAYREADY_SCHEME_ID)
	if err != nil {
		return ContentProtection{}, err
	}
	cp.XMLNS = Strptr(CENC_XMLNS)
	prSystemID, err := hex.DecodeString(CONTENT_PROTECTION_PLAYREADY_SCHEME_HEX)
	if err != nil {
		return ContentProtection{}, err
	}

	proBin, err := base64.StdEncoding.DecodeString(pro)
	if err != nil {
		return ContentProtection{}, err
	}

	psshBox, err := makePSSHBox(prSystemID, proBin)
	if err != nil {
		return ContentProtection{}, err
	}
	cp.PSSH = Strptr(base64.StdEncoding.EncodeToString(psshBox))

	as.AddContentProtection(cp)
	return cp, nil
}

// AddNewContentProtectionSchemePlayreadyV10WithPSSH adds a new content protection scheme for PlayReady v1.0 DRM. The scheme
// will include both ms:pro and cenc:pssh subelements
// pro - PlayReady Object Header, as a Base64 encoded string.
func (as *AdaptationSet) AddNewContentProtectionSchemePlayreadyV10WithPSSH(pro string) (ContentProtection, error) {
	cp, err := newPlayreadyContentProtection(pro, CONTENT_PROTECTION_PLAYREADY_SCHEME_V10_ID)
	if err != nil {
		return ContentProtection{}, err
	}
	cp.XMLNS = Strptr(CENC_XMLNS)
	prSystemID, err := hex.DecodeString(CONTENT_PROTECTION_PLAYREADY_SCHEME_V10_HEX)
	if err != nil {
		return ContentProtection{}, err
	}

	proBin, err := base64.StdEncoding.DecodeString(pro)
	if err != nil {
		return ContentProtection{}, err
	}

	psshBox, err := makePSSHBox(prSystemID, proBin)
	if err != nil {
		return ContentProtection{}, err
	}
	cp.PSSH = Strptr(base64.StdEncoding.EncodeToString(psshBox))

	as.AddContentProtection(cp)
	return cp, nil
}

// Internal helper method for adding a ContentProtection to an AdaptationSet.
func (as *AdaptationSet) AddContentProtection(cp ContentProtection) {
	as.ContentProtection = append(as.ContentProtection, cp)
}

// Sets up a new SegmentTemplate for an AdaptationSet.
// duration - relative to timescale (i.e. 2000).
// init - template string for init segment (i.e. $RepresentationID$/audio/en/init.mp4).
// media - template string for media segments.
// startNumber - the number to start segments from ($Number$) (i.e. 0).
// timescale - sets the timescale for duration (i.e. 1000, represents milliseconds).
func (as *AdaptationSet) SetNewSegmentTemplate(duration int64, init string, media string, startNumber int64, timescale int64) (*SegmentTemplate, error) {
	st := &SegmentTemplate{
		Duration:       Int64ptr(duration),
		Initialization: Strptr(init),
		Media:          Strptr(media),
		StartNumber:    Int64ptr(startNumber),
		Timescale:      Int64ptr(timescale),
	}

	err := as.setSegmentTemplate(st)
	if err != nil {
		return nil, err
	}
	return st, nil
}

// Internal helper method for setting the Segment Template on an AdaptationSet.
func (as *AdaptationSet) setSegmentTemplate(st *SegmentTemplate) error {
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
		AudioSamplingRate: Int64ptr(samplingRate),
		Bandwidth:         Int64ptr(bandwidth),
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
		Bandwidth: Int64ptr(bandwidth),
		Codecs:    Strptr(codecs),
		ID:        Strptr(id),
		FrameRate: Strptr(frameRate),
		Width:     Int64ptr(width),
		Height:    Int64ptr(height),
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
		Bandwidth: Int64ptr(bandwidth),
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

// Adds a new Role to an AdaptationSet
// schemeIdUri - Scheme ID URI string (i.e. urn:mpeg:dash:role:2011)
// value - Value for this role, (i.e. caption, subtitle, main, alternate, supplementary, commentary, dub)
func (as *AdaptationSet) AddNewRole(schemeIDURI string, value string) (*Role, error) {
	r := &Role{
		SchemeIDURI: Strptr(schemeIDURI),
		Value:       Strptr(value),
	}
	r.AdaptationSet = as
	as.Roles = append(as.Roles, r)
	return r, nil
}

// Sets the BaseURL for a Representation.
// baseURL - Base URL as a string (i.e. 800k/output-audio-und.mp4)
func (r *Representation) SetNewBaseURL(baseURL string) error {
	if baseURL == "" {
		return ErrBaseURLEmpty
	}
	r.BaseURL = Strptr(baseURL)
	return nil
}

// Sets a new SegmentBase on a Representation.
// This is for On Demand profile.
// indexRange - Byte range to the index (sidx)atom.
// init - Byte range to the init atoms (ftyp+moov).
func (r *Representation) AddNewSegmentBase(indexRange string, initRange string) (*SegmentBase, error) {
	sb := &SegmentBase{
		IndexRange:     Strptr(indexRange),
		Initialization: &URL{Range: Strptr(initRange)},
	}

	err := r.setSegmentBase(sb)
	if err != nil {
		return nil, err
	}
	return sb, nil
}

// Internal helper method for setting the SegmentBase on a Representation.
func (r *Representation) setSegmentBase(sb *SegmentBase) error {
	if r.AdaptationSet == nil {
		return ErrNoDASHProfileSet
	}
	if sb == nil {
		return ErrSegmentBaseNil
	}
	r.SegmentBase = sb
	return nil
}

// Sets a new AudioChannelConfiguration on a Representation.
// This is required for the HbbTV profile.
// scheme - One of the two AudioConfigurationSchemes.
// channelConfiguration - string that represents the channel configuration.
func (r *Representation) AddNewAudioChannelConfiguration(scheme AudioChannelConfigurationScheme, channelConfiguration string) (*AudioChannelConfiguration, error) {
	acc := &AudioChannelConfiguration{
		SchemeIDURI: Strptr((string)(scheme)),
		Value:       Strptr(channelConfiguration),
	}

	err := r.setAudioChannelConfiguration(acc)
	if err != nil {
		return nil, err
	}

	return acc, nil
}

// Internal helper method for setting the SegmentBase on a Representation.
func (r *Representation) setAudioChannelConfiguration(acc *AudioChannelConfiguration) error {
	if acc == nil {
		return ErrAudioChannelConfigurationNil
	}
	r.AudioChannelConfiguration = acc
	return nil
}
