package connect

import (
	"encoding/json"
	"time"
)

// Time is a type masking a time.Time capable of parsing the JSON from
// Garmin Connect.
type Time struct{ time.Time }

// UnmarshalJSON implements json.Unmarshaler. It can parse timestamps
// returned from connect.garmin.com.
func (t *Time) UnmarshalJSON(value []byte) error {
	// FIXME: Somehow we should deal with timezones :-/
	layouts := []string{
		"2006-01-02T15:04:05.0",
		"2006-01-02 15:04:05",
	}

	var blip string
	err := json.Unmarshal(value, &blip)
	if err != nil {
		return err
	}

	var proxy time.Time
	for _, l := range layouts {
		proxy, err = time.Parse(l, blip)
		if err == nil {
			break
		}
	}

	t.Time = proxy

	return nil
}
