package jsonull

import (
	"encoding/json"
	"github.com/lib/pq"
	"time"
)

const DateFormat = "2006-01-02 15:04:05"

type JsonNullTimestamp struct {
	pq.NullTime
}

func (v JsonNullTimestamp) MarshalJSON() ([]byte, error) {
	if v.Valid {
		stamp := v.Time.Format(DateFormat)
		return json.Marshal(stamp)
	} else {
		return json.Marshal(nil)
	}
}

func (v *JsonNullTimestamp) UnmarshalJSON(data []byte) error {
	var x *string
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		if result, e := time.Parse(DateFormat, *x); e != nil {
			return e
		} else {
			v.Time = result
			v.Valid = true
		}
	} else {
		v.Valid = false
	}
	return nil
}
