// based on code from golang src/time/time.go

package mpd

import (
	"encoding/xml"
	"errors"
	"strconv"
	"strings"
	"time"
)

type Duration time.Duration

func (d Duration) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{name, d.String()}, nil
}

func (d *Duration) UnmarshalXMLAttr(attr xml.Attr) error {
	dur, err := parseDuration(attr.Value)
	if err != nil {
		return err
	}
	*d = Duration(dur)
	return nil
}

// String renders a Duration in XML Duration Data Type format
func (d *Duration) String() string {
	// Largest time is 2540400h10m10.000000000s
	var buf [32]byte
	w := len(buf)

	u := uint64(*d)
	neg := *d < 0
	if neg {
		u = -u
	}

	if u < uint64(time.Second) {
		// Special case: if duration is smaller than a second,
		// use smaller units, like 1.2ms
		var prec int
		w--
		buf[w] = 'S'
		w--
		if u == 0 {
			return "PT0S"
		}
		/*
			switch {
			case u < uint64(Millisecond):
				// print microseconds
				prec = 3
				// U+00B5 'µ' micro sign == 0xC2 0xB5
				w-- // Need room for two bytes.
				copy(buf[w:], "µ")
			default:
				// print milliseconds
				prec = 6
				buf[w] = 'm'
			}
		*/
		w, u = fmtFrac(buf[:w], u, prec)
		w = fmtInt(buf[:w], u)
	} else {
		w--
		buf[w] = 'S'

		w, u = fmtFrac(buf[:w], u, 9)

		// u is now integer seconds
		w = fmtInt(buf[:w], u%60)
		u /= 60

		// u is now integer minutes
		if u > 0 {
			w--
			buf[w] = 'M'
			w = fmtInt(buf[:w], u%60)
			u /= 60

			// u is now integer hours
			// Stop at hours because days can be different lengths.
			if u > 0 {
				w--
				buf[w] = 'H'
				w = fmtInt(buf[:w], u)
			}
		}
	}

	if neg {
		w--
		buf[w] = '-'
	}

	return "PT" + string(buf[w:])
}

// fmtFrac formats the fraction of v/10**prec (e.g., ".12345") into the
// tail of buf, omitting trailing zeros.  it omits the decimal
// point too when the fraction is 0.  It returns the index where the
// output bytes begin and the value v/10**prec.
func fmtFrac(buf []byte, v uint64, prec int) (nw int, nv uint64) {
	// Omit trailing zeros up to and including decimal point.
	w := len(buf)
	print := false
	for i := 0; i < prec; i++ {
		digit := v % 10
		print = print || digit != 0
		if print {
			w--
			buf[w] = byte(digit) + '0'
		}
		v /= 10
	}
	if print {
		w--
		buf[w] = '.'
	}
	return w, v
}

// fmtInt formats v into the tail of buf.
// It returns the index where the output begins.
func fmtInt(buf []byte, v uint64) int {
	w := len(buf)
	if v == 0 {
		w--
		buf[w] = '0'
	} else {
		for v > 0 {
			w--
			buf[w] = byte(v%10) + '0'
			v /= 10
		}
	}
	return w
}

func parseDuration(str string) (time.Duration, error) {
	if len(str) < 3 {
		return 0, errors.New("input duration too short")
	}

	var minus bool
	offset := 0
	if str[offset] == '-' {
		minus = true
		offset++
	}

	if str[offset] != 'P' {
		return 0, errors.New("input duration does not have a valid prefix")
	}
	offset++

	base := time.Unix(0, 0)
	t := base

	var dateStr, timeStr string
	if i := strings.IndexByte(str[offset:], 'T'); i != -1 {
		dateStr = str[offset : offset+i]
		timeStr = str[offset+i+1:]
	} else {
		dateStr = str[offset:]
	}

	if len(dateStr) > 0 {
		var pos int
		var err error
		var years, months, days int
		if i := strings.IndexByte(dateStr[pos:], 'Y'); i != -1 {
			years, err = strconv.Atoi(dateStr[pos : pos+i])
			if err != nil {
				return 0, err
			}
			pos += i + 1
		}

		if i := strings.IndexByte(dateStr[pos:], 'M'); i != -1 {
			months, err = strconv.Atoi(dateStr[pos : pos+i])
			if err != nil {
				return 0, err
			}
			pos += i + 1
		}

		if i := strings.IndexByte(dateStr[pos:], 'D'); i != -1 {
			days, err = strconv.Atoi(dateStr[pos : pos+i])
			if err != nil {
				return 0, err
			}
		}
		t = t.AddDate(years, months, days)
	}

	if len(timeStr) > 0 {
		var pos int
		var sum float64
		if i := strings.IndexByte(timeStr[pos:], 'H'); i != -1 {
			hours, err := strconv.ParseInt(timeStr[pos:pos+i], 10, 64)
			if err != nil {
				return 0, err
			}
			sum += float64(hours) * 3600
			pos += i + 1
		}
		if i := strings.IndexByte(timeStr[pos:], 'M'); i != -1 {
			minutes, err := strconv.ParseInt(timeStr[pos:pos+i], 10, 64)
			if err != nil {
				return 0, err
			}
			sum += float64(minutes) * 60
			pos += i + 1
		}
		if i := strings.IndexByte(timeStr[pos:], 'S'); i != -1 {
			seconds, err := strconv.ParseFloat(timeStr[pos:pos+i], 64)
			if err != nil {
				return 0, err
			}
			sum += seconds
		}
		t = t.Add(time.Duration(sum * float64(time.Second)))
	}

	if minus {
		return -t.Sub(base), nil
	}
	return t.Sub(base), nil
}
