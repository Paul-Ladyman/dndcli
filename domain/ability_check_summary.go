package domain

import "fmt"

// AbilityCheckSummary contains the details of an ability check
type AbilityCheckSummary struct {
	Ability      Ability
	Modifier     Ability
	Circumstance Circumstance
	Dice         Dice
	Roll         Roll
	Target       Target
	Proficient   bool
}

func (summary AbilityCheckSummary) String() string {
	return fmt.Sprintf(`AbilityCheckSummary {
		Ability: %s,
		Modifier: %s,
		Circumstance: %s,
		Dice: %s,
		Roll: %s,
		Target: %s,
		Proficient: %t
	}`, summary.Ability, summary.Modifier, summary.Circumstance, summary.Dice, summary.Roll, summary.Target, summary.Proficient)
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

// SummariseRoll summarises what dice to roll and how many times
func (summary AbilityCheckSummary) SummariseRoll() string {
	return fmt.Sprintf(summary.Roll.Summarise(summary.Dice))
}

// SummariseTarget summarises whether to use DC or AC as the check target
func (summary AbilityCheckSummary) SummariseTarget() string {
	return fmt.Sprintf("Compare against %s", summary.Target.Summarise())
}

// SummariseModifier summarises which modifier to apply
func (summary AbilityCheckSummary) SummariseModifier() string {
	return fmt.Sprintf("Add the player's %s modifier (abilities section on character sheet)", summary.Modifier)
}

// SummariseProficiency summarises whether a proficiency bonus should be added
func (summary AbilityCheckSummary) SummariseProficiency() string {
	if summary.Proficient {
		return "Add the player's proficiency bonus (proficiency bonus section on character sheet)"
	}
	return ""
}
