package domain

// AbilityCheckSummary contains the details of an ability check
type AbilityCheckSummary struct {
	Ability      Ability
	Modifier     Ability
	Circumstance Circumstance
	Dice         Dice
	Throw        Throw
}
