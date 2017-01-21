package jsonenc

import (
	"testing"
	"bytes"
	"encoding/json"
	"github.com/fausto/jsonenc/test"
	"github.com/stretchr/testify/assert"
)

func TestEnc(t *testing.T) {

	buffer := bytes.NewBufferString("")

	enc := NewEncoder(buffer)
	enc.WriteStartObject()
	enc.WriteNameValueString("stringField", "my string field")
	enc.WriteNameValueLiteral("intField", "10")
	enc.WriteStartArrayWithName("stringArray")
	enc.WriteStringValue("value 1")
	enc.WriteStringValue("value 2")
	enc.WriteEndArray()
	enc.WriteStartArrayWithName("intArray")
	enc.WriteLiteralValue("1")
	enc.WriteLiteralValue("2")
	enc.WriteLiteralValue("3")
	enc.WriteEndArray()
	enc.WriteStartArrayWithName("objectArray")
	enc.WriteStartObject()
	enc.WriteNameValueString("field", "object 1")
	enc.WriteEndObject()
	enc.WriteStartObject()
	enc.WriteNameValueString("field", "object 2")
	enc.WriteEndObject()
	enc.WriteEndArray()
	enc.WriteEndObject()

	actual := string(buffer.Bytes())

	expectedJson := test.ExpectedJson{
		StringField: "my string field",
		IntField: 10,
		StringArray: []string{"value 1", "value 2"},
		IntArray: []int{1, 2, 3},
		ObjectArray:[]test.Object{
			{Field:"object 1"},
			{Field:"object 2"}}}

	b, err := json.Marshal(expectedJson)
	if (err != nil) {
		t.Fatal(err)
	}
	expected := string(b)

	assert.Equal(t, expected, actual)

}