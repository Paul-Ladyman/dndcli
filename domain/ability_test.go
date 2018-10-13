package domain

import (
	"testing"
)

func TestAbilityFactoryStrength(t *testing.T) {
	expected := Strength
	result, _ := AbilityFactory("str")
	if result != expected {
		t.Errorf("Wrong ability returned. %q was not equal to %q", result, expected)
	}

	expectedString := "strength"
	actualString := result.String()
	if actualString != expectedString {
		t.Errorf("Wrong ability string. %q was not equal to %q", actualString, expectedString)
	}
}

func TestAbilityFactoryDexterity(t *testing.T) {
	expected := Dexterity
	result, _ := AbilityFactory("dex")
	if result != expected {
		t.Errorf("Wrong ability returned. %q was not equal to %q", result, expected)
	}

	expectedString := "dexterity"
	actualString := result.String()
	if actualString != expectedString {
		t.Errorf("Wrong ability string. %q was not equal to %q", actualString, expectedString)
	}
}

func TestAbilityFactoryConstitution(t *testing.T) {
	expected := Constitution
	result, _ := AbilityFactory("con")
	if result != expected {
		t.Errorf("Wrong ability returned. %q was not equal to %q", result, expected)
	}

	expectedString := "constitution"
	actualString := result.String()
	if actualString != expectedString {
		t.Errorf("Wrong ability string. %q was not equal to %q", actualString, expectedString)
	}
}

func TestAbilityFactoryIntelligence(t *testing.T) {
	expected := Intelligence
	result, _ := AbilityFactory("int")
	if result != expected {
		t.Errorf("Wrong ability returned. %q was not equal to %q", result, expected)
	}

	expectedString := "intelligence"
	actualString := result.String()
	if actualString != expectedString {
		t.Errorf("Wrong ability string. %q was not equal to %q", actualString, expectedString)
	}
}

func TestAbilityFactoryWisdom(t *testing.T) {
	expected := Wisdom
	result, _ := AbilityFactory("wis")
	if result != expected {
		t.Errorf("Wrong ability returned. %q was not equal to %q", result, expected)
	}

	expectedString := "wisdom"
	actualString := result.String()
	if actualString != expectedString {
		t.Errorf("Wrong ability string. %q was not equal to %q", actualString, expectedString)
	}
}

func TestAbilityFactoryCharisma(t *testing.T) {
	expected := Charisma
	result, _ := AbilityFactory("cha")
	if result != expected {
		t.Errorf("Wrong ability returned. %q was not equal to %q", result, expected)
	}

	expectedString := "charisma"
	actualString := result.String()
	if actualString != expectedString {
		t.Errorf("Wrong ability string. %q was not equal to %q", actualString, expectedString)
	}
}

func TestAbilityFactoryUnknown(t *testing.T) {
	result, err := AbilityFactory("unknown")
	if err == nil {
		t.Errorf("Expected error, actually got %q", result)
	}
}
