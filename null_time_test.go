package jsonull

import (
	"encoding/json"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type objectNullTime struct {
	UpdatedAt JsonNullTime `json:"updated_at"`
}

func TestJsonNullTime_MarshallNullTime(t *testing.T) {

	object := &objectNullTime{
		UpdatedAt: JsonNullTime{pq.NullTime{Valid: false}},
	}
	actual, _ := json.Marshal(object)
	assert.NotEmpty(t, string(actual))
	assert.JSONEq(t, `{"updated_at":null}`, string(actual))
}

func TestJsonNullTime_MarshallNonNullTime(t *testing.T) {
	object := &objectNullTime{
		UpdatedAt: JsonNullTime{pq.NullTime{Valid: true, Time: time.Now()}},
	}
	actual, _ := json.Marshal(object)
	assert.NotEmpty(t, string(actual))
	assert.NotEqual(t, `{"updated_at":null}`, string(actual))
}

func TestJsonNullTime_UnMarshallNonNullTime(t *testing.T) {
	body := `{"updated_at":1554640929535}`
	var object objectNullTime
	e := json.Unmarshal([]byte(body), &object)
	assert.Nil(t, e)
	assert.NotNil(t, object.UpdatedAt)
	assert.True(t, object.UpdatedAt.Valid)
	assert.Contains(t, object.UpdatedAt.Time.String(), "2019-04-07")
}

func TestJsonNullTime_UnMarshallInvalidTime(t *testing.T) {
	body := `{"updated_at": "abc"}`
	var object objectNullTime
	e := json.Unmarshal([]byte(body), &object)
	assert.NotNil(t, e)
}
