package domain

import (
	"testing"
)

func TestCircumstanceFactoryAdvantage(t *testing.T) {
	expected := Advantage
	result, _ := CircumstanceFactory("adv")
	if result != expected {
		t.Errorf("Wrong circumstance returned. %q was not equal to %q", result, expected)
	}

	expectedString := "advantage"
	actualString := result.String()
	if actualString != expectedString {
		t.Errorf("Wrong circumstance string. %q was not equal to %q", actualString, expectedString)
	}
}

func TestCircumstanceFactoryDisadvantage(t *testing.T) {
	expected := Disadvantage
	result, _ := CircumstanceFactory("dis")
	if result != expected {
		t.Errorf("Wrong circumstance returned. %q was not equal to %q", result, expected)
	}

	expectedString := "disadvantage"
	actualString := result.String()
	if actualString != expectedString {
		t.Errorf("Wrong circumstance string. %q was not equal to %q", actualString, expectedString)
	}
}

func TestCircumstanceFactoryNeutral(t *testing.T) {
	expected := Neutral
	result, _ := CircumstanceFactory("")
	if result != expected {
		t.Errorf("Wrong circumstance returned. %q was not equal to %q", result, expected)
	}

	expectedString := ""
	actualString := result.String()
	if actualString != expectedString {
		t.Errorf("Wrong circumstance string. %q was not equal to %q", actualString, expectedString)
	}
}

func TestCircumstanceFactoryUnknown(t *testing.T) {
	result, err := CircumstanceFactory("unknown")
	if err == nil {
		t.Errorf("Expected error, actually got %q", result)
	}
}

func TestResolveNeutralIndividual(t *testing.T) {
	expected := Neutral
	circumstance, _ := CircumstanceFactory("")
	result := circumstance.Resolve(Individual)
	if result != expected {
		t.Errorf("Wrong circumstance returned. %q was not equal to %q", result, expected)
	}
}

func TestResolveNeutralHelp(t *testing.T) {
	expected := Advantage
	circumstance, _ := CircumstanceFactory("")
	result := circumstance.Resolve(Help)
	if result != expected {
		t.Errorf("Wrong circumstance returned. %q was not equal to %q", result, expected)
	}
}

func TestResolveAdvantageIndividual(t *testing.T) {
	expected := Advantage
	circumstance, _ := CircumstanceFactory("adv")
	result := circumstance.Resolve(Individual)
	if result != expected {
		t.Errorf("Wrong circumstance returned. %q was not equal to %q", result, expected)
	}
}

func TestResolveAdvantageHelp(t *testing.T) {
	expected := Advantage
	circumstance, _ := CircumstanceFactory("adv")
	result := circumstance.Resolve(Help)
	if result != expected {
		t.Errorf("Wrong circumstance returned. %q was not equal to %q", result, expected)
	}
}

func TestResolveDisadvantageIndividual(t *testing.T) {
	expected := Disadvantage
	circumstance, _ := CircumstanceFactory("dis")
	result := circumstance.Resolve(Individual)
	if result != expected {
		t.Errorf("Wrong circumstance returned. %q was not equal to %q", result, expected)
	}
}

func TestResolveDisadvantageHelp(t *testing.T) {
	expected := Neutral
	circumstance, _ := CircumstanceFactory("dis")
	result := circumstance.Resolve(Help)
	if result != expected {
		t.Errorf("Wrong circumstance returned. %q was not equal to %q", result, expected)
	}
}
