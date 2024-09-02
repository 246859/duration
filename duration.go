package duration

import (
	"strconv"
	"time"
)

const (
	Nanosecond  = Duration(time.Nanosecond)
	Microsecond = Duration(time.Microsecond)
	Millisecond = Duration(time.Millisecond)
	Second      = Duration(time.Second)
	Minute      = Duration(time.Minute)
	Hour        = Duration(time.Hour)
)

func From(duration time.Duration) Duration {
	return Duration(duration)
}

// Duration is a wrapper of time.Duration, which implements TextUnmarshaler, TextMarshaler, MarshalJSON()
type Duration time.Duration

func (d Duration) String() string {
	return time.Duration(d).String()
}

func (d *Duration) UnmarshalText(text []byte) error {
	if d == nil {
		return nil
	}
	parsedDura, err := time.ParseDuration(string(text))
	if err != nil {
		return err
	}
	*d = Duration(parsedDura)
	return nil
}

func (d Duration) MarshalText() (text []byte, err error) {
	s := time.Duration(d).String()
	return []byte(s), nil
}

func (d Duration) MarshalJSON() ([]byte, error) {
	s := time.Duration(d).String()
	return []byte(strconv.Quote(s)), nil
}

func (d *Duration) Duration() time.Duration {
	if d != nil {
		return time.Duration(*d)
	}
	return 0
}
