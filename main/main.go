package main

import (
	"fmt"
	"bytes"
	"github.com/fausto/jsonenc"
)

func main() {
	buffer := bytes.NewBufferString("")

	enc := jsonenc.NewEncoder(buffer)
	enc.WriteStartObject()
	enc.WriteNameValueString("stringField", "my string field")
	enc.WriteNameValueInt("intField", 10)
	enc.WriteStartArrayWithName("stringArray")
	enc.WriteString("value 1")
	enc.WriteString("value 2")
	enc.WriteEndArray()
	enc.WriteStartArrayWithName("intArray")
	enc.WriteInt(1)
	enc.WriteInt(2)
	enc.WriteInt(3)
	enc.WriteEndArray()
	enc.WriteStartArrayWithName("objectArray")
	enc.WriteStartObject()
	enc.WriteNameValueString("field", "object 1")
	enc.WriteEndObject()
	enc.WriteStartObject()
	enc.WriteNameValueString("field", "object 2")
	enc.WriteEndObject()

	actual := string(buffer.Bytes())

	fmt.Println(actual)
}
