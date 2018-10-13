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
		Throw:        domain.Throw{Number: 1, Preference: domain.Only},
	}
	result, _ := AbilityCheck("str", "")
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
		Throw:        domain.Throw{Number: 1, Preference: domain.Only},
	}
	result, _ := AbilityCheck("dex", "")
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
		Throw:        domain.Throw{Number: 1, Preference: domain.Only},
	}
	result, _ := AbilityCheck("con", "")
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
		Throw:        domain.Throw{Number: 1, Preference: domain.Only},
	}
	result, _ := AbilityCheck("int", "")
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
		Throw:        domain.Throw{Number: 1, Preference: domain.Only},
	}
	result, _ := AbilityCheck("wis", "")
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
		Throw:        domain.Throw{Number: 1, Preference: domain.Only},
	}
	result, _ := AbilityCheck("cha", "")
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
		Throw:        domain.Throw{Number: 2, Preference: domain.Higher},
	}
	result, _ := AbilityCheck("cha", "adv")
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
		Throw:        domain.Throw{Number: 2, Preference: domain.Lower},
	}
	result, _ := AbilityCheck("cha", "dis")
	if result != expected {
		t.Errorf("Wrong ability check summary returned. %q was not equal to %q", result, expected)
	}
}
