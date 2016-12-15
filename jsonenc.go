package jsonenc

import (
	"io"
	"strconv"
)

type JsonEncoder interface {
	WriteStartObject()

	WriteEndObject()

	WriteStartArray()

	WriteEndArray()

	WriteString(value string)

	WriteName(name string)

	WriteStartObjectWithName(name string)

	WriteStartArrayWithName(name string)

	WriteNameValueString(name string, value string)

	WriteNameValueInt(name string, value int)

	WriteInt(value int)
}

type Encoder struct {
	w                io.Writer
	prependWithComma bool
	writingArray     bool
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

func (enc *Encoder) writeCommaInsideArrayIfApplicable() {
	if (enc.writingArray) {
		enc.writeCommaIfApplicable()
	}
}

func (enc *Encoder) writeCommaIfApplicable() {
	if (enc.prependWithComma) {
		enc.w.Write([]byte(","))
	}
	enc.prependWithComma = true
}

func (enc *Encoder) WriteStartObject() {
	enc.writeCommaInsideArrayIfApplicable()
	enc.w.Write([]byte("{"))
	enc.prependWithComma = false
}

func (enc *Encoder) WriteEndObject() {
	enc.w.Write([]byte("}"))
	enc.prependWithComma = true
}

func (enc *Encoder) WriteStartArray() {
	enc.w.Write([]byte("["))
	enc.prependWithComma = false
	enc.writingArray = true
}

func (enc *Encoder) WriteEndArray() {
	enc.w.Write([]byte("]"))
	enc.prependWithComma = true
	enc.writingArray = false
}

func (enc *Encoder) WriteString(value string) {
	enc.writeCommaInsideArrayIfApplicable()
	enc.w.Write([]byte("\""))
	enc.w.Write([]byte(value))
	enc.w.Write([]byte("\""))
}

func (enc *Encoder) WriteInt(value int) {
	enc.writeCommaInsideArrayIfApplicable()
	enc.w.Write([]byte(strconv.Itoa(value)))
}

func (enc *Encoder) WriteName(name string) {
	enc.writeCommaIfApplicable()
	enc.WriteString(name)
	enc.w.Write([]byte(":"))
}

func (enc *Encoder) WriteStartObjectWithName(name string) {
	enc.WriteName(name)
	enc.WriteStartObject()
}

func (enc *Encoder) WriteStartArrayWithName(name string) {
	enc.WriteName(name)
	enc.WriteStartArray()
}

func (enc *Encoder) WriteNameValueString(name string, value string) {
	enc.WriteName(name)
	enc.WriteString(value)
}

func (enc *Encoder) WriteNameValueInt(name string, value int) {
	enc.WriteName(name)
	enc.WriteInt(value)
}
