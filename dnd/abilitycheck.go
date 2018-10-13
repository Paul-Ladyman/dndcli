package dnd

import "github.com/Paul-Ladyman/dndcli/domain"

// AbilityCheck returns a summary of an ability check given an ability and an advantage circumstance
func AbilityCheck(abilityString string, circumstanceString string) (domain.AbilityCheckSummary, error) {
	circumstance, _ := domain.CircumstanceFactory(circumstanceString)
	ability, _ := domain.AbilityFactory(abilityString)
	return domain.AbilityCheckSummary{
		Ability:      ability,
		Modifier:     ability,
		Circumstance: circumstance,
		Dice:         domain.D20,
		Throw:        domain.ThrowFactory(circumstance),
	}, nil
}
