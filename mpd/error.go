package mpd

import "errors"

// Known error variables
var (
	ErrNoDASHProfileSet               error = errors.New("No DASH profile set")
	ErrAdaptationSetNil                     = errors.New("Adaptation Set nil")
	ErrSegmentTemplateLiveProfileOnly       = errors.New("Segment template can only be used with Live Profile")
	ErrSegmentTemplateNil                   = errors.New("Segment Template nil ")
	ErrRepresentationNil                    = errors.New("Representation nil")
	ErrBaseURLEmpty                         = errors.New("Base URL empty")
	ErrSegmentBaseOnDemandProfileOnly       = errors.New("Segment Base can only be used with On-Demand Profile")
	ErrSegmentBaseNil                       = errors.New("Segment Base nil")
	ErrAudioChannelConfigurationNil         = errors.New("Audio Channel Configuration nil")
	ErrInvalidDefaultKID                    = errors.New("Invalid Default KID string, should be 32 characters")
	ErrPROEmpty                             = errors.New("PlayReady PRO empty")
	ErrContentProtectionNil                 = errors.New("Content Protection nil")
)
