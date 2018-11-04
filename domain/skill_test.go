package domain

import (
	"testing"
)

func assertAbility(expected Ability, actual Ability, t *testing.T) {
	if actual != expected {
		t.Errorf("Wrong ability for skill. %q was not equal to %q", actual, expected)
	}
}

func assertSkill(expected Skill, actual Skill, t *testing.T) {
	if actual != expected {
		t.Errorf("Wrong skill returned. %q was not equal to %q", actual, expected)
	}
}

func assertSkillString(expected string, actual string, t *testing.T) {
	if actual != expected {
		t.Errorf("Wrong skill string. %q was not equal to %q", actual, expected)
	}
}

func TestSkillFactoryAthletics(t *testing.T) {
	result, _ := SkillFactory("ath")
	assertSkill(Athletics, result, t)
	assertSkillString("athletics", result.String(), t)
}

func TestAthletics(t *testing.T) {
	assertAbility(Strength, Athletics.Ability, t)
}

func TestSkillFactoryUnknown(t *testing.T) {
	result, err := SkillFactory("unknown")
	if err == nil {
		t.Errorf("Expected error, actually got %q", result)
	}
}
