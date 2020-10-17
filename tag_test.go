package main

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

const EXPECTED_KEY = "KEY"
const EXPECTED_VALUE = "VALUE"

func TestNewTag_noError(t *testing.T) {
	tag, err := NewTag(EXPECTED_KEY, EXPECTED_VALUE)
	if tag == nil {
		t.Errorf("expected tag, but received nil")
	}

	if err != nil {
		t.Errorf("expected nil error, but received: %v", err)
	}
}

func TestNewTag_errorOnNoKey(t *testing.T) {
	tag, err := NewTag("", EXPECTED_VALUE)
	if tag != nil {
		t.Errorf("expected nil tag, but received %v", tag)
	}

	if err == nil {
		t.Errorf("expected error, but received nil")
	}

	if !errors.Is(err, ErrNoStructKey) {
		t.Errorf("did not receive expected wrapped error: ErrNoStructKey")
	}
}

func TestTagKey(t *testing.T) {
	tag, _ := NewTag(EXPECTED_KEY, EXPECTED_VALUE)

	actualKey := tag.Key()

	if actualKey != EXPECTED_KEY {
		t.Errorf("actual %v != expected %v", actualKey, EXPECTED_KEY)
	}
}

func TestTagValue(t *testing.T) {
	tag, _ := NewTag(EXPECTED_KEY, EXPECTED_VALUE)

	actualValue := tag.Value()

	if actualValue != EXPECTED_VALUE {
		t.Errorf("actual %v != expected %v", actualValue, EXPECTED_VALUE)
	}
}

func TestTagString_keyOnly(t *testing.T) {
	tag, _ := NewTag(EXPECTED_KEY, "")

	representation := tag.String()
	if representation != tag.Key() {
		t.Errorf("actual %v != expected %v", representation, tag.Key())
	}
}

func TestTagString_withValue(t *testing.T) {
	tag, _ := NewTag(EXPECTED_KEY, EXPECTED_VALUE)

	representation := tag.String()
	expected := fmt.Sprintf(`%s:"%s"`, EXPECTED_KEY, EXPECTED_VALUE)

	if representation != expected {
		t.Errorf("actual %v != expected %v", representation, expected)
	}
}

func TestTagString_withEscapedValue(t *testing.T) {
	value := `VA"LUE`
	escapedValue := strings.Replace(value, `"`, `\"`, -1)

	tag, _ := NewTag(EXPECTED_KEY, value)

	representation := tag.String()
	expected := fmt.Sprintf(`%s:"%s"`, EXPECTED_KEY, escapedValue)

	if representation != expected {
		t.Errorf("actual %v != expected %v", representation, expected)
	}
}
