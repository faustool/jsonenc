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

	stream := NewJsonStream(buffer)
	stream.WriteStartObject()
	stream.WriteNameValueString("stringField", "my string field")
	stream.WriteNameValueLiteral("intField", "10")
	stream.WriteStartArrayWithName("stringArray")
	stream.WriteStringValue("value 1")
	stream.WriteStringValue("value 2")
	stream.WriteEndArray()
	stream.WriteStartArrayWithName("intArray")
	stream.WriteLiteralValue("1")
	stream.WriteLiteralValue("2")
	stream.WriteLiteralValue("3")
	stream.WriteEndArray()
	stream.WriteStartArrayWithName("objectArray")
	stream.WriteStartObject()
	stream.WriteNameValueString("field", "object 1")
	stream.WriteEndObject()
	stream.WriteStartObject()
	stream.WriteNameValueString("field", "object 2")
	stream.WriteEndObject()
	stream.WriteEndArray()
	stream.WriteEndObject()

	actual := string(buffer.Bytes())

	expectedJson := test.ExpectedJson{
		StringField: "my string field",
		IntField:    10,
		StringArray: []string{"value 1", "value 2"},
		IntArray:    []int{1, 2, 3},
		ObjectArray:[]test.Object{
			{Field:"object 1"},
			{Field:"object 2"}}}

	b, err := json.Marshal(expectedJson)
	if err != nil {
		t.Fatal(err)
	}
	expected := string(b)

	assert.Equal(t, expected, actual)

}

func TestStream_WriteJsonContent(t *testing.T) {

	buffer := bytes.NewBufferString("")

	stream := NewJsonStream(buffer)
	stream.WriteStartObject()
	stream.WriteName("Object")
	stream.WriteJsonContent("{\"name\":\"value\"}");
	stream.WriteEndObject()

	expectedString := "{\"Object\": {\"name\": \"value\"}}";
	var expected map[string]interface{}
	err := json.Unmarshal([]byte(expectedString), &expected)
	if err != nil {
		t.Fatal(err)
	}

	var actual map[string]interface{}
	err = json.Unmarshal([]byte(buffer.Bytes()), &actual)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual, ":(")
}
