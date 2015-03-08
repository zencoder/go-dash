package mpd

func Strptr(v string) *string {
	p := new(string)
	*p = v
	return p
}

func Intptr(v int64) *int64 {
	p := new(int64)
	*p = v
	return p
}

func Uintptr(v uint) *uint {
	p := new(uint)
	*p = v
	return p
}

func Boolptr(v bool) *bool {
	p := new(bool)
	*p = v
	return p
}

func Float64ptr(v float64) *float64 {
	p := new(float64)
	*p = v
	return p
}
