package utils

import (
	"testing"
)

func TestPrettifyProperDay(t *testing.T) {
	input := "c-ordinary-22"
	expectedOutput := "Twenty Second Sunday in Ordinary Time"

	output, err := PrettifyProperDay(input)
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	if output != expectedOutput {
		t.Errorf("Expected %s, got %s", expectedOutput, output)
	}
}

func TestPrettifyProperDayNamedFeast(t *testing.T) {
	input := "c-the-transfiguration-of-the-lord"
	expectedOutput := "The Transfiguration of the Lord"

	output, err := PrettifyProperDay(input)
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	if output != expectedOutput {
		t.Errorf("Expected %s, got %s", expectedOutput, output)
	}
}
