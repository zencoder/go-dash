// based on code from golang src/time/time.go

package mpd

import (
	"encoding/xml"
	"errors"
	"time"

	"github.com/go-chrono/chrono"
)

type Duration time.Duration

var unsupportedFormatErr = errors.New("duration must be in the format: P[nD][T[nH][nM][nS]]")

func (d *Duration) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: name, Value: d.String()}, nil
}

func (d *Duration) UnmarshalXMLAttr(attr xml.Attr) error {
	duration, err := ParseDuration(attr.Value)
	if err != nil {
		return err
	}
	*d = Duration(duration)
	return nil
}

// String parses the duration into a string with the use of the chrono library.
func (d *Duration) String() string {
	if d == nil {
		return "PT0S"
	}

	return chrono.DurationOf(chrono.Extent(*d)).String()
}

// ParseDuration converts the given string into a time.Duration with the use of
// the chrono library. The function doesn't allow the use of negative durations,
// decimal valued periods, or the use of the year, month, or week units as they
// don't make sense.
func ParseDuration(str string) (time.Duration, error) {
	period, duration, err := chrono.ParseDuration(str)
	if err != nil {
		return 0, unsupportedFormatErr
	}

	hasDecimalDays := period.Days != float32(int64(period.Days))
	hasUnsupportedUnits := period.Years+period.Months+period.Years > 0
	if hasDecimalDays || hasUnsupportedUnits {
		return 0, unsupportedFormatErr
	}

	durationDays := chrono.Extent(period.Days) * 24 * chrono.Hour
	totalDur := duration.Add(chrono.DurationOf(durationDays))

	if totalDur.Compare(chrono.Duration{}) == -1 {
		return 0, errors.New("duration cannot be negative")
	}

	return time.Duration(totalDur.Nanoseconds()), nil
}
