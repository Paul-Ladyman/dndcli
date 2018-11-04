package domain

import "fmt"

// RollPreference is a generic type for describing which of multiple dice rolls to use
type RollPreference string

// RollPreferences for selecting the higher or lower of multiple dice or using the only dice rolled for one dice
const (
	Higher RollPreference = "higher"
	Lower  RollPreference = "lower"
	Only   RollPreference = "only"
)

// Roll specifies how many dice to roll and the RollPreference
type Roll struct {
	Number     int
	Preference RollPreference
}

func (roll Roll) String() string {
	return fmt.Sprintf("Roll {Number: %d, Preference: %s}", roll.Number, roll.Preference)
}

func (roll Roll) Summarise(dice Dice) string {
	if roll.Number > 1 {
		return fmt.Sprintf("Roll a %d%s and use the %s value", roll.Number, dice, roll.Preference)
	}
	return fmt.Sprintf("Roll a %s", dice)
}

func getNumber(circumstance Circumstance) int {
	switch circumstance {
	case Neutral:
		return 1
	default:
		return 2
	}
}

func getPreference(circumstance Circumstance) RollPreference {
	switch circumstance {
	case Advantage:
		return Higher
	case Disadvantage:
		return Lower
	default:
		return Only
	}
}

// RollFactory returns a Roll given a circumstance
func RollFactory(circumstance Circumstance) Roll {
	return Roll{Number: getNumber(circumstance), Preference: getPreference(circumstance)}
}
