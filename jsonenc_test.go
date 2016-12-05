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

	enc.WriteNameValueString("stringField", "value")
	enc.WriteNameValueInt("intField", 10)

	expected := string(buffer.Bytes())

	actualJson := test.ActualJson{}
	actualJson.StringField = "my string field"
	actualJson.IntField = 10
	actualJson.StringArray = []string{"value 1", "value 2"}
	actualJson.IntArray = []int{1, 2, 3}
	actualJson.ObjectArray = []test.Object{test.Object{Field:"object 1"}, test.Object{Field:"object 2"}}

	b, err := json.Marshal(actualJson)
	if (err != null) {
		t.Fatal(err)
	}
	actual := string(b)

	assert.Equal(t, expected, actual)

}