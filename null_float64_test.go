package jsonull

import (
	"database/sql"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

type objectNullFloat64 struct {
	Score JsonNullFloat64 `json:"score"`
}

func TestJsonNullInt64_MarshallNonNullFloat64(t *testing.T) {

	object := &objectNullFloat64{
		Score: JsonNullFloat64{sql.NullFloat64{Valid: true, Float64: 7.5}},
	}
	actual, _ := json.Marshal(object)
	assert.NotEmpty(t, string(actual))
	assert.Equal(t, `{"score":7.5}`, string(actual))
}

func TestJsonNullInt64_MarshallNullFloat64(t *testing.T) {

	object := &objectNullFloat64{
		Score: JsonNullFloat64{sql.NullFloat64{Valid: false}},
	}
	actual, _ := json.Marshal(object)
	assert.NotEmpty(t, string(actual))
	assert.Equal(t, `{"score":null}`, string(actual))
}

func TestJsonNullInt64_UnMarshallNonNullFloat64(t *testing.T) {
	body := `{"score": 7.5}`
	var object objectNullFloat64
	e := json.Unmarshal([]byte(body), &object)
	assert.Nil(t, e)
	assert.NotNil(t, object)
	assert.True(t, object.Score.Valid)
	assert.Equal(t, float64(7.5), object.Score.Float64)
}

func TestJsonNullInt64_UnMarshallNullFloat64(t *testing.T) {
	body := `{"score":null}`
	var object objectNullFloat64
	e := json.Unmarshal([]byte(body), &object)
	assert.Nil(t, e)
	assert.NotNil(t, object)
	assert.False(t, object.Score.Valid)
}

func TestJsonNullInt64_UnMarshallInvalidFloat64(t *testing.T) {
	body := `{"score":"abc"}`
	var object objectNullFloat64
	e := json.Unmarshal([]byte(body), &object)
	assert.NotNil(t, e)
}
