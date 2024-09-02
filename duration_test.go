package duration

import (
	"encoding/json"
	"testing"
	"time"
)

type example struct {
	D *Duration `json:"d"`
}

func TestJsonUnmarshal(t *testing.T) {
	samples := []struct {
		str  string
		want time.Duration
	}{
		{str: `{"d":"1h2m"}`, want: time.Hour + time.Minute*2},
		{str: `{"d":"25s"}`, want: time.Second * 25},
		{str: `{"d":"1000ms"}`, want: time.Millisecond * 1000},
		{str: `{"d":"0s"}`, want: 0},
		{str: `{"d":"1ns"}`, want: time.Nanosecond},
	}

	for _, sample := range samples {
		e := example{}
		err := json.Unmarshal([]byte(sample.str), &e)
		if err != nil {
			t.Fatal(err)
		}

		if e.D.Duration() != sample.want {
			t.Fatalf("expected %v, got %v", sample.want, e.D.Duration())
		}
	}
}

func TestJsonMarshal(t *testing.T) {
	samples := []struct {
		data time.Duration
		want string
	}{
		{time.Second, `{"d":"1s"}`},
		{time.Nanosecond, `{"d":"1ns"}`},
		{time.Hour + time.Minute, `{"d":"1h1m0s"}`},
	}

	for _, sample := range samples {
		e := example{D: From(sample.data)}
		marshal, err := json.Marshal(e)
		if err != nil {
			t.Fatal(err)
		}
		if string(marshal) != sample.want {
			t.Fatalf("expected %v, got %v", sample.want, string(marshal))
		}
	}
}
