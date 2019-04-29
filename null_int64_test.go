package jsonull

import (
	"database/sql"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

type objectNullInt64 struct {
	CreatedBy JsonNullInt64 `json:"id"`
}

func TestJsonNullInt64_MarshallNonNull64(t *testing.T) {

	object := &objectNullInt64{
		CreatedBy: JsonNullInt64{sql.NullInt64{Valid: true, Int64: 12}},
	}
	actual, _ := json.Marshal(object)
	assert.NotEmpty(t, string(actual))
	assert.Equal(t, `{"id":12}`, string(actual))
}

func TestJsonNullInt64_MarshallNullInt64(t *testing.T) {

	object := &objectNullInt64{
		CreatedBy: JsonNullInt64{sql.NullInt64{Valid: false}},
	}
	actual, _ := json.Marshal(object)
	assert.NotEmpty(t, string(actual))
	assert.Equal(t, `{"id":null}`, string(actual))
}

func TestJsonNullInt64_UnMarshallNonNullInt64(t *testing.T) {
	body := `{"id":15}`
	var object objectNullInt64
	e := json.Unmarshal([]byte(body), &object)
	assert.Nil(t, e)
	assert.NotNil(t, object)
	assert.True(t, object.CreatedBy.Valid)
	assert.Equal(t, int64(15), object.CreatedBy.Int64)
}

func TestJsonNullInt64_UnMarshallNullInt64(t *testing.T) {
	body := `{"id":null}`
	var object objectNullInt64
	e := json.Unmarshal([]byte(body), &object)
	assert.Nil(t, e)
	assert.NotNil(t, object)
	assert.False(t, object.CreatedBy.Valid)
}

func TestJsonNullInt64_UnMarshallInvalidInt64(t *testing.T) {
	body := `{"id":"abc"}`
	var object objectNullInt64
	e := json.Unmarshal([]byte(body), &object)
	assert.NotNil(t, e)
}
