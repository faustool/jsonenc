package jsonenc

import (
	"io"
	"github.com/fausto/stack"
)

const (
	FIRST = iota
	NOT_FIRST = iota
)

type JsonEncoder interface {
	WriteStartObject()

	WriteStartObjectWithName(name string)

	WriteEndObject()

	WriteStartArray()

	WriteStartArrayWithName(name string)

	WriteEndArray()

	WriteStringValue(value string)

	WriteNameValueString(name string, value string)

	WriteLiteralValue(value string)

	WriteNameValueLiteral(name string, value string)

	writeName(name string)

	writeLiteral(value string)

	writeString(value string)

	prependComma();
}

type Encoder struct {
	w         io.Writer
	state     stack.Stack
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w, state: stack.NewStack()}
}

func (enc *Encoder) WriteStartObject() {
	enc.prependComma()
	enc.writeLiteral("{")
	enc.state.Push(FIRST)
}

func (enc *Encoder) WriteStartObjectWithName(name string) {
	enc.prependComma()
	enc.writeName(name)
	enc.writeLiteral("{")
	enc.state.Push(FIRST)
}

func (enc *Encoder) WriteEndObject() {
	enc.writeLiteral("}")
	enc.state.Pop()
	enc.state.Push(NOT_FIRST)
}

func (enc *Encoder) WriteStartArray() {
	enc.prependComma()
	enc.writeLiteral("[")
	enc.state.Push(FIRST)
}

func (enc *Encoder) WriteStartArrayWithName(name string) {
	enc.prependComma()
	enc.writeName(name)
	enc.writeLiteral("[")
	enc.state.Push(FIRST)
}

func (enc *Encoder) WriteEndArray() {
	enc.writeLiteral("]")
	enc.state.Pop()
	enc.state.Push(NOT_FIRST)
}

func (enc *Encoder) WriteStringValue(value string) {
	enc.prependComma()
	enc.writeString(value)
	enc.state.Push(NOT_FIRST)
}

func (enc *Encoder) WriteNameValueString(name string, value string) {
	enc.prependComma()
	enc.writeName(name)
	enc.writeString(value)
	enc.state.Push(NOT_FIRST)
}

func (enc *Encoder) WriteLiteralValue(value string) {
	enc.prependComma()
	enc.writeLiteral(value)
	enc.state.Push(NOT_FIRST)
}

func (enc *Encoder) WriteNameValueLiteral(name string, value string) {
	enc.prependComma()
	enc.writeName(name)
	enc.writeLiteral(value);
	enc.state.Push(NOT_FIRST)
}

func (enc *Encoder) writeName(name string) {
	enc.writeString(name)
	enc.writeLiteral(":")
}

func (enc *Encoder) writeLiteral(value string) {
	enc.w.Write([]byte(value))
}

func (enc *Encoder) writeString(value string) {
	enc.w.Write([]byte("\""))
	enc.w.Write([]byte(value))
	enc.w.Write([]byte("\""))
}

func (enc *Encoder) prependComma() {
	value, error := enc.state.Peek();
	if (error == nil && value == NOT_FIRST) {
		enc.w.Write([]byte(","))
	}
}

