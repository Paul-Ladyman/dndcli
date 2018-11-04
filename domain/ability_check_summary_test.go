package domain

import (
	"testing"
)

var summary = AbilityCheckSummary{
	Strength,
	Strength,
	Skill{},
	Neutral,
	D20,
	Roll{1, Higher},
	DC,
	true,
}

func TestSummariseTarget(t *testing.T) {
	expected := "Compare against the Difficulty Class (DC): Easy - 10, Medium - 15, Hard - 20"
	result := summary.SummariseTarget()
	if result != expected {
		t.Errorf(" %q was not equal to %q", result, expected)
	}
}

func TestSummariseRoll(t *testing.T) {
	expected := "Roll a d20"
	result := summary.SummariseRoll()
	if result != expected {
		t.Errorf(" %q was not equal to %q", result, expected)
	}
}

func TestSummariseAbility(t *testing.T) {
	expected := "Make a strength check"
	result := summary.SummariseAbility()
	if result != expected {
		t.Errorf(" %q was not equal to %q", result, expected)
	}
}

func TestSummariseAbilityAdv(t *testing.T) {
	summary := AbilityCheckSummary{
		Dexterity,
		Dexterity,
		Skill{},
		Advantage,
		D20,
		Roll{1, Higher},
		DC,
		false,
	}
	expected := "Make a dexterity check with advantage"
	result := summary.SummariseAbility()
	if result != expected {
		t.Errorf(" %q was not equal to %q", result, expected)
	}
}

func TestSummariseAbilityDis(t *testing.T) {
	summary := AbilityCheckSummary{
		Intelligence,
		Intelligence,
		Skill{},
		Disadvantage,
		D20,
		Roll{1, Higher},
		DC,
		false,
	}
	expected := "Make a intelligence check with disadvantage"
	result := summary.SummariseAbility()
	if result != expected {
		t.Errorf(" %q was not equal to %q", result, expected)
	}
}

func TestSummariseAbilityWithSkill(t *testing.T) {
	summary := AbilityCheckSummary{
		Strength,
		Strength,
		Athletics,
		Neutral,
		D20,
		Roll{1, Higher},
		DC,
		false,
	}
	expected := "Make a strength (athletics) check"
	result := summary.SummariseAbility()
	if result != expected {
		t.Errorf(" %q was not equal to %q", result, expected)
	}
}

func TestSummariseModifier(t *testing.T) {
	expected := "Add the player's strength modifier (abilities section on character sheet)"
	result := summary.SummariseModifier()
	if result != expected {
		t.Errorf(" %q was not equal to %q", result, expected)
	}
}

func TestSummariseProficiencyTrue(t *testing.T) {
	expected := "Add the player's proficiency bonus (proficiency bonus section on character sheet)"
	result := summary.SummariseProficiency()
	if result != expected {
		t.Errorf(" %q was not equal to %q", result, expected)
	}
}

func TestSummariseProficiencyFalse(t *testing.T) {
	summary := AbilityCheckSummary{
		Intelligence,
		Intelligence,
		Skill{},
		Disadvantage,
		D20,
		Roll{1, Higher},
		DC,
		false,
	}
	expected := ""
	result := summary.SummariseProficiency()
	if result != expected {
		t.Errorf(" %q was not equal to %q", result, expected)
	}
}
