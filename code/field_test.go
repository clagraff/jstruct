package code

import (
	"errors"
	"testing"
)

const expectedName = "fieldName"
const expectedType = "fieldType"

var expectedTags = []Tag{MustNewTag("key", "value")}

func TestNewField_noError(t *testing.T) {
	field, err := NewField(expectedName, expectedType, nil)
	if field == nil {
		t.Errorf("expected field, but received nil")
	}

	if err != nil {
		t.Errorf("expected nil error, but received: %v", err)
	}
}

func TestNewField_errorOnNoName(t *testing.T) {
	field, err := NewField("", expectedType, nil)
	if field != nil {
		t.Errorf("expected nil field, but received %v", field)
	}

	if err == nil {
		t.Errorf("expected error, but received nil")
	}

	if !errors.Is(err, ErrNoFieldName) {
		t.Errorf("did not receive expected wrapped error: ErrNoFieldName")
	}
}

func TestNewField_errorOnNoType(t *testing.T) {
	field, err := NewField(expectedName, "", nil)
	if field != nil {
		t.Errorf("expected nil field, but received %v", field)
	}

	if err == nil {
		t.Errorf("expected error, but received nil")
	}

	if !errors.Is(err, ErrNoFieldType) {
		t.Errorf("did not receive expected wrapped error: ErrNoFieldType")
	}
}

func TestFieldName(t *testing.T) {
	field, _ := NewField(expectedName, expectedType, nil)

	actual := field.Name()

	if actual != expectedName {
		t.Errorf("actual %v != expected %v", actual, expectedName)
	}
}

func TestFieldType(t *testing.T) {
	field, _ := NewField(expectedName, expectedType, nil)

	actual := field.Type()

	if actual != expectedType {
		t.Errorf("actual %v != expected %v", actual, expectedType)
	}
}

func TestFieldTags_noTags(t *testing.T) {
	field, _ := NewField(expectedName, expectedType, nil)

	actual := field.Tags()

	if len(actual) != 0 {
		t.Errorf("actual %v != expected %v", len(actual), 0)
	}
}

func TestFieldTags_withTags(t *testing.T) {
	field, _ := NewField(expectedName, expectedType, expectedTags)

	actual := field.Tags()

	if len(actual) != len(expectedTags) {
		t.Errorf("actual %v != expected %v", len(actual), len(expectedTags))
	}

	if actual[0].String() != expectedTags[0].String() {
		t.Errorf(
			"actual %v != expected %v",
			actual[0].String(),
			expectedTags[0].String(),
		)
	}
}

func TestFieldString_noTags(t *testing.T) {
	field, _ := NewField(expectedName, expectedType, nil)

	representation := field.String()
	expected := expectedName + " " + expectedType
	if representation != expected {
		t.Errorf("actual %v != expected %v", representation, expected)
	}
}

func TestFieldString_oneTag(t *testing.T) {
	field, _ := NewField(expectedName, expectedType, expectedTags)

	representation := field.String()
	expected := expectedName + " " + expectedType + " `" + expectedTags[0].String() + "`"
	if representation != expected {
		t.Errorf("actual %v != expected %v", representation, expected)
	}
}

func TestFieldString_multipleTags(t *testing.T) {
	twoTags := []Tag{
		MustNewTag("keyOne", "valueOne"),
		MustNewTag("keyTwo", ""),
	}

	field, _ := NewField(expectedName, expectedType, twoTags)

	representation := field.String()
	expected := expectedName +
		" " +
		expectedType +
		" `" +
		twoTags[0].String() + " " +
		twoTags[1].String() +
		"`"

	if representation != expected {
		t.Errorf("actual %v != expected %v", representation, expected)
	}
}
