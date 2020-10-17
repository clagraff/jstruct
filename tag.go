package main

import (
	"errors"
	"strings"
)

var ErrNoStructKey = errors.New("No struct key provided")

// Tag represents a Golang struct tag of a required key
// and optional value.
type Tag interface {
	Key() string
	Value() string

	String() string
}

func NewTag(key string, value string) (Tag, error) {
	if key == "" {
		return nil, ErrNoStructKey
	}

	return tag{
		key:   key,
		value: value,
	}, nil
}

type tag struct {
	key   string
	value string
}

func (t tag) Key() string   { return t.key }
func (t tag) Value() string { return t.value }

func (t tag) String() string {
	var builder strings.Builder

	builder.WriteString(t.Key())

	if t.Value() != "" {
		escaped := strings.Replace(t.Value(), `"`, `\"`, -1)

		builder.WriteString(`:"`) // opening quote
		builder.WriteString(escaped)
		builder.WriteString(`"`) // closing quote
	}

	return builder.String()
}
