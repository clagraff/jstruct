package code

import (
	"errors"
	"strings"
)

// ErrNoTagKey represents an error where an empty key was provided.
var ErrNoTagKey = errors.New("No struct key provided")

// Tag represents a Golang struct tag of a required key
// and optional value.
type Tag interface {
	// Key returns the tag key; it will never be an empty string.
	Key() string

	// Value returns the tag value; it may be an empty string.
	Value() string

	// String returns a Golang code representation of the tag.
	//
	// It follows the standard established by the reflect package,
	// resulting in `key` or `key:"value"` output, depending on
	// the presence of a non-empty `value`.
	//
	// If present, any double-quote characters in `value` will
	// be escaped with a prefix slash, a la `VA\"UE`.
	String() string
}

// NewTag returns a Tag with the provided key and optional value.
//
// The key MUST be a non-empty string, otherwise an error is returned.
// Value may be empty.
func NewTag(key string, value string) (Tag, error) {
	if key == "" {
		return nil, ErrNoTagKey
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
