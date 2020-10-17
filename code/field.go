package code

import (
	"errors"
	"strings"
)

// ErrNoFieldName represents an error where an an empty field name was provided.
var ErrNoFieldName = errors.New("No field name provided")

// ErrNoFieldType represents an error where an an empty field type was provided.
var ErrNoFieldType = errors.New("No field type provided")

// Field represents a struct field.
//
// Field name & type are required; a field tag is optional.
type Field interface {
	// Name returns the name of the field.
	Name() string

	// Fields returns a slice of tags for the field. May be empty.
	Tags() []Tag

	// Type returns a string representation of the type of the field.
	Type() string

	// String returns a Golang code representation of the field.
	//
	// If present, will include space-separated tag representations.
	String() string
}

// NewField returns a Field with the provided name and type, and optional tags.
func NewField(name, typeRepr string, tags []Tag) (Field, error) {
	if name == "" {
		return nil, ErrNoFieldName
	}

	if typeRepr == "" {
		return nil, ErrNoFieldType
	}

	if tags == nil {
		tags = make([]Tag, 0)
	}

	return field{
		name:     name,
		typeRepr: typeRepr,
		tags:     tags,
	}, nil
}

type field struct {
	name     string
	typeRepr string
	tags     []Tag
}

func (f field) Name() string { return f.name }
func (f field) Type() string { return f.typeRepr }
func (f field) Tags() []Tag  { return f.tags }

func (f field) String() string {
	var builder strings.Builder

	builder.WriteString(f.Name())
	builder.WriteString(" ")
	builder.WriteString(f.Type())

	if len(f.Tags()) != 0 {
		builder.WriteString(" `") // Opening backtick

		tags := make([]string, 0)
		for _, tag := range f.Tags() {
			tags = append(tags, tag.String())
		}
		builder.WriteString(strings.Join(tags, " "))

		builder.WriteString("`") // Closing backtick
	}

	return builder.String()
}
