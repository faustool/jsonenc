package jsonenc

import (
	"io"
	"github.com/fausto/stack"
)

const (
	FIRST = iota
	NOT_FIRST = iota
)

type JsonStream interface {
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

	prependComma()
}

type Stream struct {
	w         io.Writer
	state     stack.Stack
}

func NewEncoder(w io.Writer) *Stream {
	return &Stream{w: w, state: stack.NewStack()}
}

func (stream *Stream) WriteStartObject() {
	stream.prependComma()
	stream.writeLiteral("{")
	stream.state.Push(FIRST)
}

func (stream *Stream) WriteStartObjectWithName(name string) {
	stream.prependComma()
	stream.writeName(name)
	stream.writeLiteral("{")
	stream.state.Push(FIRST)
}

func (stream *Stream) WriteEndObject() {
	stream.writeLiteral("}")
	stream.state.Pop()
	stream.state.Push(NOT_FIRST)
}

func (stream *Stream) WriteStartArray() {
	stream.prependComma()
	stream.writeLiteral("[")
	stream.state.Push(FIRST)
}

func (stream *Stream) WriteStartArrayWithName(name string) {
	stream.prependComma()
	stream.writeName(name)
	stream.writeLiteral("[")
	stream.state.Push(FIRST)
}

func (stream *Stream) WriteEndArray() {
	stream.writeLiteral("]")
	stream.state.Pop()
	stream.state.Push(NOT_FIRST)
}

func (stream *Stream) WriteStringValue(value string) {
	stream.prependComma()
	stream.writeString(value)
	stream.state.Push(NOT_FIRST)
}

func (stream *Stream) WriteNameValueString(name string, value string) {
	stream.prependComma()
	stream.writeName(name)
	stream.writeString(value)
	stream.state.Push(NOT_FIRST)
}

func (stream *Stream) WriteLiteralValue(value string) {
	stream.prependComma()
	stream.writeLiteral(value)
	stream.state.Push(NOT_FIRST)
}

func (stream *Stream) WriteNameValueLiteral(name string, value string) {
	stream.prependComma()
	stream.writeName(name)
	stream.writeLiteral(value)
	stream.state.Push(NOT_FIRST)
}

func (stream *Stream) writeName(name string) {
	stream.writeString(name)
	stream.writeLiteral(":")
}

func (stream *Stream) writeLiteral(value string) {
	stream.w.Write([]byte(value))
}

func (stream *Stream) writeString(value string) {
	stream.w.Write([]byte("\""))
	stream.w.Write([]byte(value))
	stream.w.Write([]byte("\""))
}

func (stream *Stream) prependComma() {
	value, err := stream.state.Peek()
	if err == nil && value == NOT_FIRST {
		stream.w.Write([]byte(","))
	}
}

