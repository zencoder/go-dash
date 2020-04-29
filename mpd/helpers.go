package mpd

/*

// Creates a new static MPD object.
// profile - DASH Profile (Live or OnDemand).
// mediaPresentationDuration - Media Presentation Duration (i.e. PT6M16S).
// minBufferTime - Min Buffer Time (i.e. PT1.97S).
// attributes - Other attributes (optional).
func NewMPD(profile DashProfile, mediaPresentationDuration, minBufferTime string, attributes ...AttrMPD) *MPD {
	period := &Period{}
	mpd := &MPD{
		XMLNs:                     Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles:                  Strptr((string)(profile)),
		Type:                      Strptr("static"),
		MediaPresentationDuration: Strptr(mediaPresentationDuration),
		MinBufferTime:             Strptr(minBufferTime),
		period:                    period,
		Periods:                   []*Period{period},
	}

	for i := range attributes {
		switch attr := attributes[i].(type) {
		case *attrAvailabilityStartTime:
			mpd.AvailabilityStartTime = attr.GetStrptr()
		}
	}

	return mpd
}

// Creates a new dynamic MPD object.
// profile - DASH Profile (Live or OnDemand).
// availabilityStartTime - anchor for the computation of the earliest availability time (in UTC).
// minBufferTime - Min Buffer Time (i.e. PT1.97S).
// attributes - Other attributes (optional).
func NewDynamicMPD(profile DashProfile, availabilityStartTime, minBufferTime string, attributes ...AttrMPD) *MPD {
	period := &Period{}
	mpd := &MPD{
		XMLNs:                 Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles:              Strptr((string)(profile)),
		Type:                  Strptr("dynamic"),
		AvailabilityStartTime: Strptr(availabilityStartTime),
		MinBufferTime:         Strptr(minBufferTime),
		period:                period,
		Periods:               []*Period{period},
		UTCTiming:             &Descriptor{},
	}

	for i := range attributes {
		switch attr := attributes[i].(type) {
		case *attrMinimumUpdatePeriod:
			mpd.MinimumUpdatePeriod = attr.GetStrptr()
		case *attrMediaPresentationDuration:
			mpd.MediaPresentationDuration = attr.GetStrptr()
		case *attrSuggestedPresentationDelay:
			mpd.SuggestedPresentationDelay = attr.GetStrptr()
		}
	}

	return mpd
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
	period.Duration = DurationPtr(Duration(d))
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
func (m *MPD) AddNewAdaptationSetAudioWithID(id string, mimeType string, segmentAlignment bool, startWithSAP int64, lang string) (*AdaptationSet, error) {
	return m.period.AddNewAdaptationSetAudioWithID(id, mimeType, segmentAlignment, startWithSAP, lang)
}

// Create a new Adaptation Set for Audio Assets.
// mimeType - MIME Type (i.e. audio/mp4).
// segmentAlignment - Segment Alignment(i.e. true).
// startWithSAP - Starts With SAP (i.e. 1).
// lang - Language (i.e. en).
func (period *Period) AddNewAdaptationSetAudio(mimeType string, segmentAlignment bool, startWithSAP int64, lang string) (*AdaptationSet, error) {
	as := &AdaptationSet{
		SegmentAlignment: Boolptr(segmentAlignment),
		Lang:             Strptr(lang),
		RepresentationBase: RepresentationBase{
			MimeType:     Strptr(mimeType),
			StartWithSAP: Int64ptr(startWithSAP),
		},
	}
	err := period.addAdaptationSet(as)
	if err != nil {
		return nil, err
	}
	return as, nil
}

// Create a new Adaptation Set for Audio Assets.
// mimeType - MIME Type (i.e. audio/mp4).
// segmentAlignment - Segment Alignment(i.e. true).
// startWithSAP - Starts With SAP (i.e. 1).
// lang - Language (i.e. en).
func (period *Period) AddNewAdaptationSetAudioWithID(id string, mimeType string, segmentAlignment bool, startWithSAP int64, lang string) (*AdaptationSet, error) {
	as := &AdaptationSet{
		ID:               Strptr(id),
		SegmentAlignment: Boolptr(segmentAlignment),
		Lang:             Strptr(lang),
		RepresentationBase: RepresentationBase{
			MimeType:     Strptr(mimeType),
			StartWithSAP: Int64ptr(startWithSAP),
		},
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
func (m *MPD) AddNewAdaptationSetVideoWithID(id string, mimeType string, scanType string, segmentAlignment bool, startWithSAP int64) (*AdaptationSet, error) {
	return m.period.AddNewAdaptationSetVideoWithID(id, mimeType, scanType, segmentAlignment, startWithSAP)
}

// Create a new Adaptation Set for Video Assets.
// mimeType - MIME Type (i.e. video/mp4).
// scanType - Scan Type (i.e.progressive).
// segmentAlignment - Segment Alignment(i.e. true).
// startWithSAP - Starts With SAP (i.e. 1).
func (period *Period) AddNewAdaptationSetVideo(mimeType string, scanType string, segmentAlignment bool, startWithSAP int64) (*AdaptationSet, error) {
	as := &AdaptationSet{
		SegmentAlignment: Boolptr(segmentAlignment),
		RepresentationBase: RepresentationBase{
			MimeType:     Strptr(mimeType),
			StartWithSAP: Int64ptr(startWithSAP),
			ScanType:     Strptr(scanType),
		},
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
func (period *Period) AddNewAdaptationSetVideoWithID(id string, mimeType string, scanType string, segmentAlignment bool, startWithSAP int64) (*AdaptationSet, error) {
	as := &AdaptationSet{
		SegmentAlignment: Boolptr(segmentAlignment),
		ID:               Strptr(id),
		RepresentationBase: RepresentationBase{
			MimeType:     Strptr(mimeType),
			StartWithSAP: Int64ptr(startWithSAP),
			ScanType:     Strptr(scanType),
		},
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
func (m *MPD) AddNewAdaptationSetSubtitleWithID(id string, mimeType string, lang string) (*AdaptationSet, error) {
	return m.period.AddNewAdaptationSetSubtitleWithID(id, mimeType, lang)
}

// Create a new Adaptation Set for Subtitle Assets.
// mimeType - MIME Type (i.e. text/vtt).
// lang - Language (i.e. en).
func (period *Period) AddNewAdaptationSetSubtitle(mimeType string, lang string) (*AdaptationSet, error) {
	as := &AdaptationSet{
		Lang: Strptr(lang),
		RepresentationBase: RepresentationBase{
			MimeType: Strptr(mimeType),
		},
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
func (period *Period) AddNewAdaptationSetSubtitleWithID(id string, mimeType string, lang string) (*AdaptationSet, error) {
	as := &AdaptationSet{
		ID:   Strptr(id),
		Lang: Strptr(lang),
		RepresentationBase: RepresentationBase{
			MimeType: Strptr(mimeType),
		},
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
//
// NOTE: this is only here for Legacy purposes. This will create an invalid UUID.
func (as *AdaptationSet) AddNewContentProtectionRootLegacyUUID(defaultKIDHex string) (*ContentProtection, error) {
	if len(defaultKIDHex) != 32 || defaultKIDHex == "" {
		return nil, ErrInvalidDefaultKID
	}

	// Convert the KID into the correct format
	defaultKID := strings.ToLower(defaultKIDHex[0:8] + "-" + defaultKIDHex[8:12] + "-" + defaultKIDHex[12:16] + "-" + defaultKIDHex[16:32])

	cp := &ContentProtection{
		CENC: ContentProtectionCENC{
			Descriptor: Descriptor{
				Value:       Strptr(CONTENT_PROTECTION_ROOT_VALUE),
				SchemeIDURI: Strptr(CONTENT_PROTECTION_ROOT_SCHEME_ID_URI),
			},
			DefaultKID: Strptr(defaultKID),
		},
	}

	err := as.AddContentProtection(cp)
	if err != nil {
		return nil, err
	}
	return cp, nil
}

// Adds a ContentProtection tag at the root level of an AdaptationSet.
// This ContentProtection tag does not include signaling for any particular DRM scheme.
// defaultKIDHex - Default Key ID as a Hex String.
func (as *AdaptationSet) AddNewContentProtectionRoot(defaultKIDHex string) (*ContentProtectionCenc, error) {
	if len(defaultKIDHex) != 32 || defaultKIDHex == "" {
		return nil, ErrInvalidDefaultKID
	}

	// Convert the KID into the correct format
	defaultKID := strings.ToLower(defaultKIDHex[0:8] + "-" + defaultKIDHex[8:12] + "-" + defaultKIDHex[12:16] + "-" + defaultKIDHex[16:20] + "-" + defaultKIDHex[20:32])

	cp := &CENCContentProtection{}
	cp.DefaultKID = Strptr(defaultKID)
	cp.Value = Strptr(CONTENT_PROTECTION_ROOT_VALUE)
	cp.SchemeIDURI = Strptr(CONTENT_PROTECTION_ROOT_SCHEME_ID_URI)

	err := as.AddContentProtection(cp)
	if err != nil {
		return nil, err
	}
	return cp, nil
}

// AddNewContentProtectionSchemeWidevine adds a new content protection scheme for Widevine DRM to the adaptation set. With
// a <cenc:pssh> element that contains a Base64 encoded PSSH box
// wvHeader - binary representation of Widevine Header
// !!! Note: this function will accept any byte slice as a wvHeader value !!!
func (as *AdaptationSet) AddNewContentProtectionSchemeWidevineWithPSSH(wvHeader []byte) (*WidevineContentProtection, error) {
	cp, err := NewWidevineContentProtection(wvHeader)
	if err != nil {
		return nil, err
	}

	err = as.AddContentProtection(cp)
	if err != nil {
		return nil, err
	}
	return cp, nil
}

// AddNewContentProtectionSchemeWidevine adds a new content protection scheme for Widevine DRM to the adaptation set.
func (as *AdaptationSet) AddNewContentProtectionSchemeWidevine() (*WidevineContentProtection, error) {
	cp, err := NewWidevineContentProtection(nil)
	if err != nil {
		return nil, err
	}

	err = as.AddContentProtection(cp)
	if err != nil {
		return nil, err
	}
	return cp, nil
}

func NewWidevineContentProtection(wvHeader []byte) (*WidevineContentProtection, error) {
	cp := &WidevineContentProtection{}
	cp.SchemeIDURI = Strptr(CONTENT_PROTECTION_WIDEVINE_SCHEME_ID)

	if len(wvHeader) > 0 {
		//cp.XMLNS = Strptr(CENC_XMLNS)
		wvSystemID, err := hex.DecodeString(CONTENT_PROTECTION_WIDEVINE_SCHEME_HEX)
		if err != nil {
			panic(err.Error())
		}
		psshBox, err := makePSSHBox(wvSystemID, wvHeader)
		if err != nil {
			return nil, err
		}

		psshB64 := base64.StdEncoding.EncodeToString(psshBox)
		cp.PSSH = &psshB64
	}
	return cp, nil
}

// AddNewContentProtectionSchemePlayready adds a new content protection scheme for PlayReady DRM.
// pro - PlayReady Object Header, as a Base64 encoded string.
func (as *AdaptationSet) AddNewContentProtectionSchemePlayready(pro string) (*PlayreadyContentProtection, error) {
	cp, err := newPlayreadyContentProtection(pro, CONTENT_PROTECTION_PLAYREADY_SCHEME_ID)
	if err != nil {
		return nil, err
	}

	err = as.AddContentProtection(cp)
	if err != nil {
		return nil, err
	}
	return cp, nil
}

// AddNewContentProtectionSchemePlayreadyV10 adds a new content protection scheme for PlayReady v1.0 DRM.
// pro - PlayReady Object Header, as a Base64 encoded string.
func (as *AdaptationSet) AddNewContentProtectionSchemePlayreadyV10(pro string) (*PlayreadyContentProtection, error) {
	cp, err := newPlayreadyContentProtection(pro, CONTENT_PROTECTION_PLAYREADY_SCHEME_V10_ID)
	if err != nil {
		return nil, err
	}

	err = as.AddContentProtection(cp)
	if err != nil {
		return nil, err
	}
	return cp, nil
}

func newPlayreadyContentProtection(pro string, schemeIDURI string) (*PlayreadyContentProtection, error) {
	if pro == "" {
		return nil, ErrPROEmpty
	}

	cp := &PlayreadyContentProtection{
		PlayreadyXMLNS: Strptr(CONTENT_PROTECTION_PLAYREADY_XMLNS),
		PRO:            Strptr(pro),
	}
	cp.SchemeIDURI = Strptr(schemeIDURI)

	return cp, nil
}

// AddNewContentProtectionSchemePlayreadyWithPSSH adds a new content protection scheme for PlayReady DRM. The scheme
// will include both ms:pro and cenc:pssh subelements
// pro - PlayReady Object Header, as a Base64 encoded string.
func (as *AdaptationSet) AddNewContentProtectionSchemePlayreadyWithPSSH(pro string) (*PlayreadyContentProtection, error) {
	cp, err := newPlayreadyContentProtection(pro, CONTENT_PROTECTION_PLAYREADY_SCHEME_ID)
	if err != nil {
		return nil, err
	}
	//cp.XMLNS = Strptr(CENC_XMLNS)
	prSystemID, err := hex.DecodeString(CONTENT_PROTECTION_PLAYREADY_SCHEME_HEX)
	if err != nil {
		panic(err.Error())
	}

	proBin, err := base64.StdEncoding.DecodeString(pro)
	if err != nil {
		return nil, err
	}

	psshBox, err := makePSSHBox(prSystemID, proBin)
	if err != nil {
		return nil, err
	}
	cp.PSSH = Strptr(base64.StdEncoding.EncodeToString(psshBox))

	err = as.AddContentProtection(cp)
	if err != nil {
		return nil, err
	}
	return cp, nil
}

// AddNewContentProtectionSchemePlayreadyV10WithPSSH adds a new content protection scheme for PlayReady v1.0 DRM. The scheme
// will include both ms:pro and cenc:pssh subelements
// pro - PlayReady Object Header, as a Base64 encoded string.
func (as *AdaptationSet) AddNewContentProtectionSchemePlayreadyV10WithPSSH(pro string) (*PlayreadyContentProtection, error) {
	cp, err := newPlayreadyContentProtection(pro, CONTENT_PROTECTION_PLAYREADY_SCHEME_V10_ID)
	if err != nil {
		return nil, err
	}
	//cp.XMLNS = Strptr(CENC_XMLNS)
	prSystemID, err := hex.DecodeString(CONTENT_PROTECTION_PLAYREADY_SCHEME_V10_HEX)
	if err != nil {
		panic(err.Error())
	}

	proBin, err := base64.StdEncoding.DecodeString(pro)
	if err != nil {
		return nil, err
	}

	psshBox, err := makePSSHBox(prSystemID, proBin)
	if err != nil {
		return nil, err
	}
	cp.PSSH = Strptr(base64.StdEncoding.EncodeToString(psshBox))

	err = as.AddContentProtection(cp)
	if err != nil {
		return nil, err
	}
	return cp, nil
}

// Internal helper method for adding a ContentProtection to an AdaptationSet.
func (as *AdaptationSet) AddContentProtection(cp Protection) error {
	if cp == nil {
		return ErrContentProtectionNil
	}

	cpc := ContentProtectionContainer{
		Protection: cp,
	}
	as.ContentProtection = append(as.ContentProtection, cpc)
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
		RepresentationBase: RepresentationBase{
			AudioSamplingRate: Int64ptr(samplingRate),
			Codecs:            Strptr(codecs),
		},
		ID:        Strptr(id),
		Bandwidth: Int64ptr(bandwidth),
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
func (as *AdaptationSet) AddNewRepresentationVideo(bandwidth int64, codecs string, id string, frameRate string, width uint32, height uint32) (*Representation, error) {
	r := &Representation{
		RepresentationBase: RepresentationBase{
			Width:     Uint32ptr(width),
			Height:    Uint32ptr(height),
			FrameRate: Strptr(frameRate),
			Codecs:    Strptr(codecs),
		},
		ID:        Strptr(id),
		Bandwidth: Int64ptr(bandwidth),
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
func (r *Representation) AddNewAudioChannelConfiguration(scheme AudioChannelConfigurationScheme, channelConfiguration string) (*Descriptor, error) {
	acc := &Descriptor{
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
func (r *Representation) setAudioChannelConfiguration(acc *Descriptor) error {
	if acc == nil {
		return ErrAudioChannelConfigurationNil
	}
	r.AudioChannelConfiguration = append(r.AudioChannelConfiguration, acc)
	return nil
}
*/
