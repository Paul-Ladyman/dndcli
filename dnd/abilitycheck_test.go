package dnd

import (
	"testing"

	"github.com/Paul-Ladyman/dndcli/domain"
)

func TestStrengthAbilityCheck(t *testing.T) {
	expected := domain.AbilityCheckSummary{
		Ability:      domain.Strength,
		Modifier:     domain.Strength,
		Circumstance: domain.Neutral,
		Dice:         domain.D20,
		Roll:         domain.Roll{Number: 1, Preference: domain.Only},
		Target:       domain.DC,
		Proficient:   false,
	}
	result, _ := AbilityCheck("str", "", false)
	if result != expected {
		t.Errorf("Wrong ability check summary returned. %q was not equal to %q", result, expected)
	}
}

func TestDexterityAbilityCheck(t *testing.T) {
	expected := domain.AbilityCheckSummary{
		Ability:      domain.Dexterity,
		Modifier:     domain.Dexterity,
		Circumstance: domain.Neutral,
		Dice:         domain.D20,
		Roll:         domain.Roll{Number: 1, Preference: domain.Only},
		Target:       domain.DC,
		Proficient:   false,
	}
	result, _ := AbilityCheck("dex", "", false)
	if result != expected {
		t.Errorf("Wrong ability check summary returned. %q was not equal to %q", result, expected)
	}
}

func TestConstitutionAbilityCheck(t *testing.T) {
	expected := domain.AbilityCheckSummary{
		Ability:      domain.Constitution,
		Modifier:     domain.Constitution,
		Circumstance: domain.Neutral,
		Dice:         domain.D20,
		Roll:         domain.Roll{Number: 1, Preference: domain.Only},
		Target:       domain.DC,
		Proficient:   false,
	}
	result, _ := AbilityCheck("con", "", false)
	if result != expected {
		t.Errorf("Wrong ability check summary returned. %q was not equal to %q", result, expected)
	}
}

func TestIntelligenceAbilityCheck(t *testing.T) {
	expected := domain.AbilityCheckSummary{
		Ability:      domain.Intelligence,
		Modifier:     domain.Intelligence,
		Circumstance: domain.Neutral,
		Dice:         domain.D20,
		Roll:         domain.Roll{Number: 1, Preference: domain.Only},
		Target:       domain.DC,
		Proficient:   false,
	}
	result, _ := AbilityCheck("int", "", false)
	if result != expected {
		t.Errorf("Wrong ability check summary returned. %q was not equal to %q", result, expected)
	}
}

func TestWisdomAbilityCheck(t *testing.T) {
	expected := domain.AbilityCheckSummary{
		Ability:      domain.Wisdom,
		Modifier:     domain.Wisdom,
		Circumstance: domain.Neutral,
		Dice:         domain.D20,
		Roll:         domain.Roll{Number: 1, Preference: domain.Only},
		Target:       domain.DC,
		Proficient:   false,
	}
	result, _ := AbilityCheck("wis", "", false)
	if result != expected {
		t.Errorf("Wrong ability check summary returned. %q was not equal to %q", result, expected)
	}
}

func TestCharismaAbilityCheck(t *testing.T) {
	expected := domain.AbilityCheckSummary{
		Ability:      domain.Charisma,
		Modifier:     domain.Charisma,
		Circumstance: domain.Neutral,
		Dice:         domain.D20,
		Roll:         domain.Roll{Number: 1, Preference: domain.Only},
		Target:       domain.DC,
		Proficient:   false,
	}
	result, _ := AbilityCheck("cha", "", false)
	if result != expected {
		t.Errorf("Wrong ability check summary returned. %q was not equal to %q", result, expected)
	}
}

func TestAbilityCheckWithAdvantage(t *testing.T) {
	expected := domain.AbilityCheckSummary{
		Ability:      domain.Charisma,
		Modifier:     domain.Charisma,
		Circumstance: domain.Advantage,
		Dice:         domain.D20,
		Roll:         domain.Roll{Number: 2, Preference: domain.Higher},
		Target:       domain.DC,
		Proficient:   false,
	}
	result, _ := AbilityCheck("cha", "adv", false)
	if result != expected {
		t.Errorf("Wrong ability check summary returned. %q was not equal to %q", result, expected)
	}
}

func TestAbilityCheckWithDisadvantage(t *testing.T) {
	expected := domain.AbilityCheckSummary{
		Ability:      domain.Charisma,
		Modifier:     domain.Charisma,
		Circumstance: domain.Disadvantage,
		Dice:         domain.D20,
		Roll:         domain.Roll{Number: 2, Preference: domain.Lower},
		Target:       domain.DC,
		Proficient:   false,
	}
	result, _ := AbilityCheck("cha", "dis", false)
	if result != expected {
		t.Errorf("Wrong ability check summary returned. %q was not equal to %q", result, expected)
	}
}

func TestAbilityCheckWithToolProficiency(t *testing.T) {
	expected := domain.AbilityCheckSummary{
		Ability:      domain.Charisma,
		Modifier:     domain.Charisma,
		Circumstance: domain.Disadvantage,
		Dice:         domain.D20,
		Roll:         domain.Roll{Number: 2, Preference: domain.Lower},
		Target:       domain.DC,
		Proficient:   true,
	}
	result, _ := AbilityCheck("cha", "dis", true)
	if result != expected {
		t.Errorf("Wrong ability check summary returned. %q was not equal to %q", result, expected)
	}
}
