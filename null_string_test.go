package jsonull

import (
	"database/sql"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

type objectNullString struct {
	Name JsonNullString `json:"name"`
}

func TestJsonNullString_MarshallNonNullString(t *testing.T) {
	object := &objectNullString{
		Name: JsonNullString{sql.NullString{Valid: true, String: "John Doe"}},
	}
	actual, _ := json.Marshal(object)
	assert.NotEmpty(t, string(actual))
	assert.Equal(t, `{"name":"John Doe"}`, string(actual))
}

func TestJsonNullString_MarshallNullString(t *testing.T) {
	object := &objectNullString{
		Name: JsonNullString{sql.NullString{Valid: false,}},
	}
	actual, _ := json.Marshal(object)
	assert.NotEmpty(t, string(actual))
	assert.Equal(t, `{"name":null}`, string(actual))
}

func TestJsonNullString_UnMarshallNonNullString(t *testing.T) {
	body := `{"name": "John Doe"}`
	var object objectNullString
	e := json.Unmarshal([]byte(body), &object)
	assert.Nil(t, e)
	assert.NotNil(t, object)
	assert.True(t, object.Name.Valid)
	assert.Equal(t, "John Doe", object.Name.String)
}

func TestJsonNullString_UnMarshallNullString(t *testing.T) {
	body := `{"name":null}`
	var object objectNullString
	e := json.Unmarshal([]byte(body), &object)
	assert.Nil(t, e)
	assert.NotNil(t, object)
	assert.False(t, object.Name.Valid)
}

func TestJsonNullString_UnMarshallInvalidString(t *testing.T) {
	body := `{"name": 123}`
	var object objectNullString
	e := json.Unmarshal([]byte(body), &object)
	assert.NotNil(t, e)
}
