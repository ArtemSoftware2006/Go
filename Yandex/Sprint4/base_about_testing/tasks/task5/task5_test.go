package main

import (
	"errors"
	"testing"
)

func TestGetUTFLengthValid(t *testing.T) {
	input := []byte("Hello, world!")
	expectedLength := 13
	actualLength, err := GetUTFLength(input)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if actualLength != expectedLength {
		t.Errorf("Expected length: %d, got: %d", expectedLength, actualLength)
	}
}

func TestGetUTFLengthInvalid(t *testing.T) {
	input := []byte("\x80") // Invalid UTF-8 byte
	expectedErr := ErrInvalidUTF8

	_, err := GetUTFLength(input)

	if err == nil {
		t.Errorf("Expected error: %v, got nil", expectedErr)
	}

	if !errors.Is(err, expectedErr) {
		t.Errorf("Expected error: %v, got: %v", expectedErr, err)
	}
}
