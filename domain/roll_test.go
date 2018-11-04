package domain

import (
	"testing"
)

func TestRollFactory(t *testing.T) {
	expected := Roll{Number: 1, Preference: Only}
	result := RollFactory(Neutral)
	if result != expected {
		t.Errorf("Wrong roll returned. %q was not equal to %q", result, expected)
	}

	expectedString := "Roll a d20"
	actualString := result.Summarise(D20)
	if actualString != expectedString {
		t.Errorf("Wrong roll string returned. %q was not equal to %q", actualString, expectedString)
	}
}

func TestRollFactoryWithAdvantage(t *testing.T) {
	expected := Roll{Number: 2, Preference: Higher}
	result := RollFactory(Advantage)
	if result != expected {
		t.Errorf("Wrong roll returned. %q was not equal to %q", result, expected)
	}

	expectedString := "Roll a 2d12 and use the higher value"
	actualString := result.Summarise(D12)
	if actualString != expectedString {
		t.Errorf("Wrong roll string returned. %q was not equal to %q", actualString, expectedString)
	}
}

func TestRollFactoryWithDisadvantage(t *testing.T) {
	expected := Roll{Number: 2, Preference: Lower}
	result := RollFactory(Disadvantage)
	if result != expected {
		t.Errorf("Wrong roll returned. %q was not equal to %q", result, expected)
	}

	expectedString := "Roll a 2d10 and use the lower value"
	actualString := result.Summarise(D10)
	if actualString != expectedString {
		t.Errorf("Wrong roll string returned. %q was not equal to %q", actualString, expectedString)
	}
}
