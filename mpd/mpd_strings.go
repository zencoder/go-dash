package mpd

import "encoding/json"

func (m MPD) String() string {
	jb, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(jb)
}

func (p Period) String() string {
	jb, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(jb)
}

func (as AdaptationSet) String() string {
	jb, err := json.Marshal(as)
	if err != nil {
		return ""
	}
	return string(jb)
}

func (st SegmentTemplate) String() string {
	jb, err := json.Marshal(st)
	if err != nil {
		return ""
	}
	return string(jb)
}

func (r Representation) String() string {
	jb, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(jb)
}

func (sb SegmentBase) String() string {
	jb, err := json.Marshal(sb)
	if err != nil {
		return ""
	}
	return string(jb)
}

func (i Initialization) String() string {
	jb, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(jb)
}
