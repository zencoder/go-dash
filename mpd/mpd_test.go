package mpd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MPDSuite struct {
	suite.Suite
}

func TestMPDSuite(t *testing.T) {
	suite.Run(t, new(MPDSuite))
}

func (s *MPDSuite) SetupTest() {

}

func (s *MPDSuite) SetupSuite() {

}

const (
	VALID_MEDIA_PRESENTATION_DURATION string = "PT6M16S"
	VALID_MIN_BUFFER_TIME             string = "PT1.97S"
)

func (s *MPDSuite) TestReadMPDLiveProfile() {
	m, err := ReadMPD("fixtures/live_profile.mpd")
	assert.NotNil(s.T(), m)
	assert.Nil(s.T(), err)
	//assert.Equal(s.T(), &MPD{}, m)
}

func (s *MPDSuite) TestReadMPDOnDemandProfile() {
	m, err := ReadMPD("fixtures/ondemand_profile.mpd")
	assert.NotNil(s.T(), m)
	assert.Nil(s.T(), err)
	//assert.Equal(s.T(), &MPD{}, m)
}

func (s *MPDSuite) TestNewMPDLive() {
	m := NewMPD(DASH_PROFILE_LIVE, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	assert.NotNil(s.T(), m)
	expectedMPD := &MPD{
		XMLNs:    Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles: Strptr((string)(DASH_PROFILE_LIVE)),
		Type:     Strptr("static"),
		MediaPresentationDuration: Strptr(VALID_MEDIA_PRESENTATION_DURATION),
		MinBufferTime:             Strptr(VALID_MIN_BUFFER_TIME),
		Period:                    &Period{},
	}
	assert.Equal(s.T(), expectedMPD, m)

	xmlStr, err := m.WriteToString()
	assert.Nil(s.T(), err)
	expectedMPDStr := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<MPD xmlns=\"urn:mpeg:dash:schema:mpd:2011\" profiles=\"urn:mpeg:dash:profile:isoff-live:2011\" type=\"static\" mediaPresentationDuration=\"PT6M16S\" minBufferTime=\"PT1.97S\"><Period></Period></MPD>"
	assert.Equal(s.T(), expectedMPDStr, xmlStr)
}

func (s *MPDSuite) TestNewMPDOnDemand() {
	m := NewMPD(DASH_PROFILE_ONDEMAND, VALID_MEDIA_PRESENTATION_DURATION, VALID_MIN_BUFFER_TIME)
	assert.NotNil(s.T(), m)
	expectedMPD := &MPD{
		XMLNs:    Strptr("urn:mpeg:dash:schema:mpd:2011"),
		Profiles: Strptr((string)(DASH_PROFILE_ONDEMAND)),
		Type:     Strptr("static"),
		MediaPresentationDuration: Strptr(VALID_MEDIA_PRESENTATION_DURATION),
		MinBufferTime:             Strptr(VALID_MIN_BUFFER_TIME),
		Period:                    &Period{},
	}
	assert.Equal(s.T(), expectedMPD, m)

	xmlStr, err := m.WriteToString()
	assert.Nil(s.T(), err)
	expectedMPDStr := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<MPD xmlns=\"urn:mpeg:dash:schema:mpd:2011\" profiles=\"urn:mpeg:dash:profile:isoff-on-demand:2011\" type=\"static\" mediaPresentationDuration=\"PT6M16S\" minBufferTime=\"PT1.97S\"><Period></Period></MPD>"
	assert.Equal(s.T(), expectedMPDStr, xmlStr)
}

/*
type MPD struct {
	XMLNs                     *string `xml:"xmlns,attr"`
	Profiles                  *string `xml:"profiles,attr"`
	Type                      *string `xml:"type,attr"`
	MediaPresentationDuration *string `xml:"mediaPresentationDuration,attr"`
	MinBufferTime             *string `xml:"minBufferTime,attr"`
	Period                    *Period `xml:"Period"`
}
*/
