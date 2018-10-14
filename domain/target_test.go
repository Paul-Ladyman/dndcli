package domain

import (
	"testing"
)

func TestDCSummarise(t *testing.T) {
	expected := "the Difficulty Class (DC): Easy - 10, Medium - 15, Hard - 20"
	result := DC.Summarise()
	if result != expected {
		t.Errorf("%q was not equal to %q", result, expected)
	}
}

func TestACSummarise(t *testing.T) {
	expected := "the character's Armor Class (AC on character sheet)"
	result := AC.Summarise()
	if result != expected {
		t.Errorf("%q was not equal to %q", result, expected)
	}
}
