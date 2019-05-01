package jsonull

import (
	"encoding/json"
	"github.com/lib/pq"
	"time"
)

type JsonNullTime struct {
	pq.NullTime
}

func (v JsonNullTime) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Time.UnixNano() / int64(time.Millisecond))
	} else {
		return json.Marshal(nil)
	}
}

func (v *JsonNullTime) UnmarshalJSON(data []byte) error {
	var x *int64
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Time = time.Unix(0, *x * int64(time.Millisecond))
	} else {
		v.Valid = false
	}
	return nil
}
