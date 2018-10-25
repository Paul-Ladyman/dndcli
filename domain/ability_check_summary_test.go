package domain

import (
	"testing"
)

var summary = AbilityCheckSummary{
	Strength,
	Strength,
	Neutral,
	D20,
	Throw{1, Higher},
	DC,
}

func TestSummariseTarget(t *testing.T) {
	expected := "Compare against the Difficulty Class (DC): Easy - 10, Medium - 15, Hard - 20"
	result := summary.SummariseTarget()
	if result != expected {
		t.Errorf(" %q was not equal to %q", result, expected)
	}
}

func TestSummariseThrow(t *testing.T) {
	expected := "Throw a d20"
	result := summary.SummariseThrow()
	if result != expected {
		t.Errorf(" %q was not equal to %q", result, expected)
	}
}

func TestSummariseAbilityAndCircumstance(t *testing.T) {
	expected := "Make a strength check"
	result := summary.SummariseAbilityAndCircumstance()
	if result != expected {
		t.Errorf(" %q was not equal to %q", result, expected)
	}
}

func TestSummariseModifier(t *testing.T) {
	expected := "Add a strength modifier (abilities section on character sheet)"
	result := summary.SummariseModifier()
	if result != expected {
		t.Errorf(" %q was not equal to %q", result, expected)
	}
}

func TestSummariseAbilityAndCircumstanceAdv(t *testing.T) {
	summary := AbilityCheckSummary{
		Dexterity,
		Dexterity,
		Advantage,
		D20,
		Throw{1, Higher},
		DC,
	}
	expected := "Make a dexterity check with advantage"
	result := summary.SummariseAbilityAndCircumstance()
	if result != expected {
		t.Errorf(" %q was not equal to %q", result, expected)
	}
}

func TestSummariseAbilityAndCircumstanceDis(t *testing.T) {
	summary := AbilityCheckSummary{
		Intelligence,
		Intelligence,
		Disadvantage,
		D20,
		Throw{1, Higher},
		DC,
	}
	expected := "Make a intelligence check with disadvantage"
	result := summary.SummariseAbilityAndCircumstance()
	if result != expected {
		t.Errorf(" %q was not equal to %q", result, expected)
	}
}
