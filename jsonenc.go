package jsonenc

import (
	"io"
	"strconv"
)

type JsonEncoder interface {
	WriteStartObject()

	WriteStartObjectWithName(name string)

	WriteEndObject()

	WriteStartArray()

	WriteStartArrayWithName(name string)

	WriteEndArray()

	WriteNameValueString(name string, value string)

	WriteNameValueInt(name string, value int)

	WriteString(value string)

	WriteInt(value int)

}

type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

func (enc Encoder) WriteStartObject() {
	enc.w.Write([]byte("{"))
}

func (enc Encoder) WriteStartObjectWithName(name string) {
	enc.writeName(name)
	enc.WriteStartObject()
}


func (enc Encoder) WriteEndObject() {
	enc.w.Write("}")
}

func (enc Encoder) WriteStartArray() {
	enc.w.Write("[")
}

func (enc Encoder) WriteStartArrayWithName(name string) {
	enc.writeName(name)
	enc.WriteStartArray()
}

func (enc Encoder) WriteEndArray() {
	enc.w.Write("]")
}

func (enc Encoder) writeName(name string) {
	enc.w.Write("\"")
	enc.w.Write([]byte(name))
	enc.w.Write("\": ")
}

func (enc Encoder) WriteString(value string) {
	enc.w.Write("\"")
	enc.w.Write([]byte(value))
	enc.w.Write("\"")
}

func (enc Encoder) WriteInt(value int) {
	enc.w.Write([]byte(strconv.Itoa(value)))
}

func (enc Encoder) WriteNameValueString(name string, value string) {
	enc.writeName(name)
	enc.WriteString(value)
}

func (enc Encoder) WriteNameValueInt(name string, value int) {
	enc.writeName(name)
	enc.WriteInt(value)
}
