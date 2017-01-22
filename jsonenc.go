package jsonenc

import (
	"io"
	"github.com/fausto/stack"
)

const (
	// this represents a state where the element is first in it's container
	FIRST     = iota
	// this represents a state where other elements were added to the same container before
	NOT_FIRST = iota
)

// Json stream API intended to be super fast. There's no validation of method call order whatsoever so use it with
// care and test thoroughly your code to make sure you're generating a valid Json object
type JsonStream interface {
	// Write the beginning of a Json object with {
	WriteStartObject()

	// Write the beginning of a Json object with a name as in "name": {
	WriteStartObjectWithName(name string)

	// Write the end of a Json object with }
	WriteEndObject()

	// Write the beginning of a Json array with [
	WriteStartArray()

	// Write the beginning of a Json array with a name as in "name": [
	WriteStartArrayWithName(name string)

	// Write the end of a Json array with ]
	WriteEndArray()

	// Write a String value, i.e. between double quotes. This is meant to be used to write array items.
	WriteStringValue(value string)

	// Write a Json property and its String value as in "name": "value"
	WriteNameValueString(name string, value string)

	// Write a value without double quotes. This is mean to be used to write integer o boolean values
	WriteLiteralValue(value string)

	// Write a Json property and its value witout double quotes as in "name": value
	WriteNameValueLiteral(name string, value string)

	// Write a Json property name and its succeeding column char as in "name":
	// This is useful when you're writing an object from outside
	WriteName(name string)

	// Write a Json content built from outside. Use with care!
	WriteJsonContent(object string)

	// Write a Json name - for internal purposes
	writeName(name string)

	// Write a literal value i.e. without quotes - for internal purposes
	writeLiteral(value string)

	// Write a String value i.e. with quotes - for internal purposes
	writeString(value string)

	// Prepend a comma if applicable - for internal purposes
	prependCommaIfApplicable()
}

type Stream struct {
	w     io.Writer
	state stack.Stack
}

func NewJsonStream(w io.Writer) JsonStream {
	return Stream{w: w, state: stack.NewStack()}
}

func (stream Stream) WriteStartObject() {
	stream.prependCommaIfApplicable()
	stream.writeLiteral("{")
	stream.state.Push(FIRST)
}

func (stream Stream) WriteStartObjectWithName(name string) {
	stream.prependCommaIfApplicable()
	stream.writeName(name)
	stream.writeLiteral("{")
	stream.state.Push(FIRST)
}

func (stream Stream) WriteEndObject() {
	stream.writeLiteral("}")
	stream.state.Pop()
	stream.state.Push(NOT_FIRST)
}

func (stream Stream) WriteStartArray() {
	stream.prependCommaIfApplicable()
	stream.writeLiteral("[")
	stream.state.Push(FIRST)
}

func (stream Stream) WriteStartArrayWithName(name string) {
	stream.prependCommaIfApplicable()
	stream.writeName(name)
	stream.writeLiteral("[")
	stream.state.Push(FIRST)
}

func (stream Stream) WriteEndArray() {
	stream.writeLiteral("]")
	stream.state.Pop()
	stream.state.Push(NOT_FIRST)
}

func (stream Stream) WriteStringValue(value string) {
	stream.prependCommaIfApplicable()
	stream.writeString(value)
	stream.state.Push(NOT_FIRST)
}

func (stream Stream) WriteNameValueString(name string, value string) {
	stream.prependCommaIfApplicable()
	stream.writeName(name)
	stream.writeString(value)
	stream.state.Push(NOT_FIRST)
}

func (stream Stream) WriteLiteralValue(value string) {
	stream.prependCommaIfApplicable()
	stream.writeLiteral(value)
	stream.state.Push(NOT_FIRST)
}

func (stream Stream) WriteNameValueLiteral(name string, value string) {
	stream.prependCommaIfApplicable()
	stream.writeName(name)
	stream.writeLiteral(value)
	stream.state.Push(NOT_FIRST)
}

func (stream Stream) WriteName(name string) {
	stream.prependCommaIfApplicable()
	stream.writeName(name)
}

func (stream Stream) WriteJsonContent(object string) {
	stream.w.Write([]byte(object))
}

func (stream Stream) writeName(name string) {
	stream.writeString(name)
	stream.writeLiteral(":")
}

func (stream Stream) writeLiteral(value string) {
	stream.w.Write([]byte(value))
}

func (stream Stream) writeString(value string) {
	stream.w.Write([]byte("\""))
	stream.w.Write([]byte(value))
	stream.w.Write([]byte("\""))
}

func (stream Stream) prependCommaIfApplicable() {
	value, err := stream.state.Peek()
	if err == nil && value == NOT_FIRST {
		stream.w.Write([]byte(","))
	}
}
