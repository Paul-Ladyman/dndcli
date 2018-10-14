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

func (summary AbilityCheckSummary) SummariseTarget() string {
	return fmt.Sprintf("Compare against %s", summary.Target.Summarise())
}
