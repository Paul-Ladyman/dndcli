package domain

import "fmt"

// AbilityCheckSummary contains the details of an ability check
type AbilityCheckSummary struct {
	Ability      Ability
	Modifier     Ability
	Circumstance Circumstance
	Dice         Dice
	Throw        Throw
	Target       Target
}

// SummariseAbilityAndCircumstance summarises what check to make and whether it has
// advantage, disadvantage or neither
func (summary AbilityCheckSummary) SummariseAbilityAndCircumstance() string {
	abilitySummary := fmt.Sprintf("Make a %s check", summary.Ability)

	if summary.Circumstance == Advantage || summary.Circumstance == Disadvantage {
		return fmt.Sprintf("%s with %s", abilitySummary, summary.Circumstance)
	}

	return abilitySummary
}

// SummariseThrow summarises what dice to throw and how many times
func (summary AbilityCheckSummary) SummariseThrow() string {
	return fmt.Sprintf(summary.Throw.Summarise(summary.Dice))
}

// SummariseTarget summarises whether to use DC or AC as the check target
func (summary AbilityCheckSummary) SummariseTarget() string {
	return fmt.Sprintf("Compare against %s", summary.Target.Summarise())
}

// SummariseModifier summarises which modifier to apply
func (summary AbilityCheckSummary) SummariseModifier() string {
	return fmt.Sprintf("Add a %s modifier (abilities section on character sheet)", summary.Modifier)
}
