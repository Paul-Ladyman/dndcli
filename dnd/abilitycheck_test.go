package dnd

import (
	"testing"

	"github.com/Paul-Ladyman/dndcli/domain"
)

func basicCheck(ability domain.Ability) domain.AbilityCheckSummary {
	return domain.AbilityCheckSummary{
		Ability:      ability,
		Modifier:     ability,
		Circumstance: domain.Neutral,
		Dice:         domain.D20,
		Roll:         domain.Roll{Number: 1, Preference: domain.Only},
		Target:       domain.DC,
		Proficient:   false,
		Cooperation:  domain.Individual,
	}
}

func skillCheck(proficient bool) domain.AbilityCheckSummary {
	return domain.AbilityCheckSummary{
		Ability:      domain.Strength,
		Modifier:     domain.Strength,
		Skill:        domain.Athletics,
		Circumstance: domain.Disadvantage,
		Dice:         domain.D20,
		Roll:         domain.Roll{Number: 2, Preference: domain.Lower},
		Target:       domain.DC,
		Proficient:   proficient,
		Cooperation:  domain.Individual,
	}
}

func TestStrengthAbilityCheck(t *testing.T) {
	expected := basicCheck(domain.Strength)
	result, _ := AbilityCheck("str", "", false, false, "")
	if result != expected {
		t.Errorf("Wrong ability check summary returned. %q was not equal to %q", result, expected)
	}
}

func TestDexterityAbilityCheck(t *testing.T) {
	expected := basicCheck(domain.Dexterity)
	result, _ := AbilityCheck("dex", "", false, false, "")
	if result != expected {
		t.Errorf("Wrong ability check summary returned. %q was not equal to %q", result, expected)
	}
}

func TestConstitutionAbilityCheck(t *testing.T) {
	expected := basicCheck(domain.Constitution)
	result, _ := AbilityCheck("con", "", false, false, "")
	if result != expected {
		t.Errorf("Wrong ability check summary returned. %q was not equal to %q", result, expected)
	}
}

func TestIntelligenceAbilityCheck(t *testing.T) {
	expected := basicCheck(domain.Intelligence)
	result, _ := AbilityCheck("int", "", false, false, "")
	if result != expected {
		t.Errorf("Wrong ability check summary returned. %q was not equal to %q", result, expected)
	}
}

func TestWisdomAbilityCheck(t *testing.T) {
	expected := basicCheck(domain.Wisdom)
	result, _ := AbilityCheck("wis", "", false, false, "")
	if result != expected {
		t.Errorf("Wrong ability check summary returned. %q was not equal to %q", result, expected)
	}
}

func TestCharismaAbilityCheck(t *testing.T) {
	expected := basicCheck(domain.Charisma)
	result, _ := AbilityCheck("cha", "", false, false, "")
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
		Cooperation:  domain.Individual,
	}
	result, _ := AbilityCheck("cha", "adv", false, false, "")
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
		Cooperation:  domain.Individual,
	}
	result, _ := AbilityCheck("cha", "dis", false, false, "")
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
		Cooperation:  domain.Individual,
	}
	result, _ := AbilityCheck("cha", "dis", true, false, "")
	if result != expected {
		t.Errorf("Wrong ability check summary returned. %q was not equal to %q", result, expected)
	}
}

func TestAbilityCheckWithSkillProficiency(t *testing.T) {
	expected := skillCheck(true)
	result, _ := AbilityCheck("ath", "dis", false, true, "")
	if result != expected {
		t.Errorf("Wrong ability check summary returned. %q was not equal to %q", result, expected)
	}
}

func TestAbilityCheckWithSkill(t *testing.T) {
	expected := skillCheck(false)
	result, _ := AbilityCheck("ath", "dis", false, false, "")
	if result != expected {
		t.Errorf("Wrong ability check summary returned. %q was not equal to %q", result, expected)
	}
}

func TestAbilityCheckWithSkillAndToolProficiency(t *testing.T) {
	expected := skillCheck(true)
	result, _ := AbilityCheck("ath", "dis", true, true, "")
	if result != expected {
		t.Errorf("Wrong ability check summary returned. %q was not equal to %q", result, expected)
	}
}
