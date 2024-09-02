package duration

import (
	"strconv"
	"time"
)

func From(duration time.Duration) *Duration {
	d := Duration(duration)
	return &d
}

type Duration time.Duration

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

func (d *Duration) MarshalText() (text []byte, err error) {
	if d == nil {
		return []byte{}, err
	}
	s := time.Duration(*d).String()
	return []byte(s), nil
}

func (d *Duration) MarshalJSON() ([]byte, error) {
	if d == nil {
		return []byte{}, nil
	}
	s := time.Duration(*d).String()
	return []byte(strconv.Quote(s)), nil
}

func (d *Duration) Duration() time.Duration {
	if d != nil {
		return time.Duration(*d)
	}
	return 0
}
