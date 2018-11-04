package domain

import "errors"

// Ability is a generic type for the core abilities
type Ability string

// Core abilities
const (
	Strength     Ability = "str"
	Dexterity    Ability = "dex"
	Constitution Ability = "con"
	Intelligence Ability = "int"
	Wisdom       Ability = "wis"
	Charisma     Ability = "cha"
)

func (ability Ability) String() string {
	switch ability {
	case Strength:
		return "strength"
	case Dexterity:
		return "dexterity"
	case Constitution:
		return "constitution"
	case Intelligence:
		return "intelligence"
	case Wisdom:
		return "wisdom"
	case Charisma:
		return "charisma"
	default:
		return "unknown"
	}
}

// AbilityFactory takes a string and returns the appropriate Ability
func AbilityFactory(ability string) (Ability, error) {
	switch ability {
	case string(Strength):
		return Strength, nil
	case string(Dexterity):
		return Dexterity, nil
	case string(Constitution):
		return Constitution, nil
	case string(Intelligence):
		return Intelligence, nil
	case string(Wisdom):
		return Wisdom, nil
	case string(Charisma):
		return Charisma, nil
	default:
		return "", errors.New("Invalid ability")
	}
}
