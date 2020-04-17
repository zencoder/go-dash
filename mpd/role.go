package mpd

type Role struct {
	AdaptationSet *AdaptationSet `xml:"-"`
	SchemeIDURI   *string        `xml:"schemeIdUri,attr"`
	Value         *string        `xml:"value,attr"`
}
