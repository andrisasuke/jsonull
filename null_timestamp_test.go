package jsonull

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

type objectNullTimestamp struct {
	UpdatedAt JsonNullTimestamp `json:"updated_at"`
}

func TestJsonNullTime_MarshallNullTimestamp(t *testing.T) {

	object := &objectNullTimestamp{
		UpdatedAt: JsonNullTimestamp{pq.NullTime{Valid: false}},
	}
	actual, _ := json.Marshal(object)
	assert.NotEmpty(t, string(actual))
	assert.JSONEq(t, `{"updated_at":null}`, string(actual))
}

func TestJsonNullTime_MarshallNonNullTimestamp(t *testing.T) {
	object := &objectNullTimestamp{
		UpdatedAt: JsonNullTimestamp{pq.NullTime{Valid: true, Time: time.Now()}},
	}
	actual, _ := json.Marshal(object)
	assert.NotEmpty(t, string(actual))
	assert.NotEqual(t, `{"updated_at":null}`, string(actual))
}

func TestJsonNullTime_UnMarshallNonNullTimestamp(t *testing.T) {
	body := `{"updated_at":"2019-07-04 07:53:26"}`
	var object objectNullTimestamp
	e := json.Unmarshal([]byte(body), &object)
	assert.Nil(t, e)
	assert.NotNil(t, object.UpdatedAt)
	assert.True(t, object.UpdatedAt.Valid)
	assert.Contains(t, object.UpdatedAt.Time.String(), "2019-07-04")
}

func TestJsonNullTime_UnMarshallNullTimestamp(t *testing.T) {
	body := `{"updated_at": null}`
	var object objectNullTimestamp
	e := json.Unmarshal([]byte(body), &object)
	assert.Nil(t, e)
	assert.NotNil(t, object.UpdatedAt)
	assert.False(t, object.UpdatedAt.Valid)
	assert.Contains(t, object.UpdatedAt.Time.String(), "0001-01-01")
}
func TestJsonNullTime_UnMarshallInvalidTimestamp(t *testing.T) {
	body := `{"updated_at": "abc"}`
	var object objectNullTimestamp
	e := json.Unmarshal([]byte(body), &object)
	assert.NotNil(t, e)

	body = `{"updated_at": 123}`
	e = json.Unmarshal([]byte(body), &object)
	assert.NotNil(t, e)
}
