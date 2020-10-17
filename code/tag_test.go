package code

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

const expectedKey = "KEY"
const expectedValue = "VALUE"

func TestNewTag_noError(t *testing.T) {
	tag, err := NewTag(expectedKey, expectedValue)
	if tag == nil {
		t.Errorf("expected tag, but received nil")
	}

	if err != nil {
		t.Errorf("expected nil error, but received: %v", err)
	}
}

func TestNewTag_errorOnNoKey(t *testing.T) {
	tag, err := NewTag("", expectedValue)
	if tag != nil {
		t.Errorf("expected nil tag, but received %v", tag)
	}

	if err == nil {
		t.Errorf("expected error, but received nil")
	}

	if !errors.Is(err, ErrNoTagKey) {
		t.Errorf("did not receive expected wrapped error: ErrNoTagKey")
	}
}

func TestTagKey(t *testing.T) {
	tag, _ := NewTag(expectedKey, expectedValue)

	actualKey := tag.Key()

	if actualKey != expectedKey {
		t.Errorf("actual %v != expected %v", actualKey, expectedKey)
	}
}

func TestTagValue(t *testing.T) {
	tag, _ := NewTag(expectedKey, expectedValue)

	actualValue := tag.Value()

	if actualValue != expectedValue {
		t.Errorf("actual %v != expected %v", actualValue, expectedValue)
	}
}

func TestTagString_keyOnly(t *testing.T) {
	tag, _ := NewTag(expectedKey, "")

	representation := tag.String()
	if representation != tag.Key() {
		t.Errorf("actual %v != expected %v", representation, tag.Key())
	}
}

func TestTagString_withValue(t *testing.T) {
	tag, _ := NewTag(expectedKey, expectedValue)

	representation := tag.String()
	expected := fmt.Sprintf(`%s:"%s"`, expectedKey, expectedValue)

	if representation != expected {
		t.Errorf("actual %v != expected %v", representation, expected)
	}
}

func TestTagString_withEscapedValue(t *testing.T) {
	value := `VA"LUE`
	escapedValue := strings.Replace(value, `"`, `\"`, -1)

	tag, _ := NewTag(expectedKey, value)

	representation := tag.String()
	expected := fmt.Sprintf(`%s:"%s"`, expectedKey, escapedValue)

	if representation != expected {
		t.Errorf("actual %v != expected %v", representation, expected)
	}
}
