package domain

import "fmt"

// ThrowPreference is a generic type for describing which of multiple dice throws to use
type ThrowPreference string

// ThrowPreferences for selecting the higher or lower of multiple dice or using the only dice thrown for one dice
const (
	Higher ThrowPreference = "higher"
	Lower  ThrowPreference = "lower"
	Only   ThrowPreference = "only"
)

// Throw specifies how many dice to throw and the ThrowPreference
type Throw struct {
	Number     int
	Preference ThrowPreference
}

func (throw Throw) String() string {
	return fmt.Sprintf("Throw {Number: %d, Preference: %s}", throw.Number, throw.Preference)
}

func (throw Throw) Summarise(dice Dice) string {
	if throw.Number > 1 {
		return fmt.Sprintf("Throw a %d%s and use the %s value", throw.Number, dice, throw.Preference)
	}
	return fmt.Sprintf("Throw a %s", dice)
}

func getNumber(circumstance Circumstance) int {
	switch circumstance {
	case Neutral:
		return 1
	default:
		return 2
	}
}

func getPreference(circumstance Circumstance) ThrowPreference {
	switch circumstance {
	case Advantage:
		return Higher
	case Disadvantage:
		return Lower
	default:
		return Only
	}
}

// ThrowFactory returns a Throw given a circumstance
func ThrowFactory(circumstance Circumstance) Throw {
	return Throw{Number: getNumber(circumstance), Preference: getPreference(circumstance)}
}
