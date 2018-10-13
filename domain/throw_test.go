package domain

import (
	"testing"
)

func TestThrowFactory(t *testing.T) {
	expected := Throw{Number: 1, Preference: Only}
	result := ThrowFactory(Neutral)
	if result != expected {
		t.Errorf("Wrong throw returned. %q was not equal to %q", result, expected)
	}

	expectedString := "Throw a d20"
	actualString := result.Summarise(D20)
	if actualString != expectedString {
		t.Errorf("Wrong throw string returned. %q was not equal to %q", actualString, expectedString)
	}
}

func TestThrowFactoryWithAdvantage(t *testing.T) {
	expected := Throw{Number: 2, Preference: Higher}
	result := ThrowFactory(Advantage)
	if result != expected {
		t.Errorf("Wrong throw returned. %q was not equal to %q", result, expected)
	}

	expectedString := "Throw a 2d12 and use the higher value"
	actualString := result.Summarise(D12)
	if actualString != expectedString {
		t.Errorf("Wrong throw string returned. %q was not equal to %q", actualString, expectedString)
	}
}

func TestThrowFactoryWithDisadvantage(t *testing.T) {
	expected := Throw{Number: 2, Preference: Lower}
	result := ThrowFactory(Disadvantage)
	if result != expected {
		t.Errorf("Wrong throw returned. %q was not equal to %q", result, expected)
	}

	expectedString := "Throw a 2d10 and use the lower value"
	actualString := result.Summarise(D10)
	if actualString != expectedString {
		t.Errorf("Wrong throw string returned. %q was not equal to %q", actualString, expectedString)
	}
}
