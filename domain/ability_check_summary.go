package domain

import "fmt"

// AbilityCheckSummary contains the details of an ability check
type AbilityCheckSummary struct {
	Ability      Ability
	Modifier     Ability
	Skill        Skill
	Circumstance Circumstance
	Dice         Dice
	Roll         Roll
	Target       Target
	Proficient   bool
	Cooperation  Cooperation
}

func (summary AbilityCheckSummary) String() string {
	return fmt.Sprintf(`AbilityCheckSummary {
		Ability: %s,
		Modifier: %s,
		Skill: %s,
		Circumstance: %s,
		Dice: %s,
		Roll: %s,
		Target: %s,
		Proficient: %t
		Cooperation: %s
	}`, summary.Ability, summary.Modifier, summary.Skill, summary.Circumstance, summary.Dice, summary.Roll, summary.Target, summary.Proficient, summary.Cooperation)
}

func getCheck(ability Ability, skill Skill) string {
	if (skill != Skill{}) {
		return fmt.Sprintf("%s (%s)", ability, skill)
	}
	return ability.String()
}

// SummariseAbility summarises what check to make and whether it has
// advantage, disadvantage or neither
func (summary AbilityCheckSummary) SummariseAbility() string {
	abilitySummary := fmt.Sprintf("Make a %s check", getCheck(summary.Ability, summary.Skill))

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
	subject := "the player's"
	if summary.Cooperation == Help {
		subject = "the highest available"
	}
	return fmt.Sprintf("Add %s %s modifier (abilities section on character sheet)", subject, summary.Modifier)
}

// SummariseProficiency summarises whether a proficiency bonus should be added
func (summary AbilityCheckSummary) SummariseProficiency() string {
	if summary.Proficient {
		subject := "the player's"
		if summary.Cooperation == Help {
			subject = "the same player's"
		}
		return fmt.Sprintf("Add %s proficiency bonus (proficiency bonus section on character sheet)", subject)
	}
	return ""
}
