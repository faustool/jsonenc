package jsonenc

import (
	"io"
	"strconv"
	"github.com/fausto/stack"
)

const (
	NEW_PROPERTY = iota
	NEXT_PROPERTY = iota
	NEW_ARRAY_VALUE = iota
	NEXT_ARRAY_VALUE = iota
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

	WriteComma()
}

type Encoder struct {
	w         io.Writer
	state     stack.Stack
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w, state: stack.NewStack()}
}

func (enc *Encoder) WriteStartObject() {
	enc.w.Write([]byte("{"))
	enc.state.Push(NEW_PROPERTY)
}

func (enc *Encoder) WriteEndObject() {
	enc.w.Write([]byte("}"))
	enc.state.Pop()
}

func (enc *Encoder) WriteStartArray() {
	enc.w.Write([]byte("["))
	enc.state.Push(NEW_ARRAY_VALUE)
}

func (enc *Encoder) WriteEndArray() {
	enc.w.Write([]byte("]"))
	enc.state.Pop()
}

func (enc *Encoder) WriteString(value string) {
	currentState, err := enc.state.Peek()
	if (err == nil) {
		if (currentState == NEW_ARRAY_VALUE) {
			enc.state.Pop()
			enc.state.Push(NEXT_ARRAY_VALUE)
		} else {
			enc.WriteComma()
		}
	}
	enc.w.Write([]byte("\""))
	enc.w.Write([]byte(value))
	enc.w.Write([]byte("\""))
}

func (enc *Encoder) WriteInt(value int) {
	enc.w.Write([]byte(strconv.Itoa(value)))
}

func (enc *Encoder) WriteName(name string) {
	currentState, err := enc.state.Peek()
	if (err == nil) {
		if (currentState == NEW_PROPERTY) {
			enc.state.Pop()
			enc.state.Push(NEXT_PROPERTY)
		} else {
			enc.WriteComma()
		}
	}
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

func (enc *Encoder) WriteComma() {
	enc.w.Write([]byte(","))
}
