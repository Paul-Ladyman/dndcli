package dnd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Paul-Ladyman/dndcli/domain"
)

func getBoolInput(response string) bool {
	if strings.TrimSpace(response) == "y" {
		return true
	}
	return false
}

func getAbility(abilityString string) domain.Ability {
	ability, abilityError := domain.AbilityFactory(abilityString)

	if abilityError != nil {
		skill, _ := domain.SkillFactory(abilityString)
		return skill.Ability
	}
	return ability
}

// AbilityCheck returns a summary of an ability check given an ability and an advantage circumstance
func AbilityCheck(abilityString string, circumstanceString string, toolProficiency bool, skillProficiency bool) (domain.AbilityCheckSummary, error) {
	circumstance, _ := domain.CircumstanceFactory(circumstanceString)
	ability := getAbility(abilityString)

	return domain.AbilityCheckSummary{
		Ability:      ability,
		Modifier:     ability,
		Circumstance: circumstance,
		Dice:         domain.D20,
		Roll:         domain.RollFactory(circumstance),
		Target:       domain.DC,
		Proficient:   toolProficiency || skillProficiency,
	}, nil
}

// AbilityCheckCli interactively retrieves information from the user relevant to ability checks
func AbilityCheckCli(ability string) (domain.AbilityCheckSummary, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Does the player have Advantage (adv), Disadvantage (dis) or neither (press enter)?")
	circumstance, _ := reader.ReadString('\n')

	skill, skillErr := domain.SkillFactory(ability)
	var skillProficiency = false
	if skillErr == nil {
		fmt.Printf("Is the player proficient in the %s skill?\n", skill)
		skillProficiencyResponse, _ := reader.ReadString('\n')
		skillProficiency = getBoolInput(skillProficiencyResponse)
	}

	fmt.Println("Is the player proficient with any relevant tools (y/n)?")
	toolProficiencyResponse, _ := reader.ReadString('\n')

	toolProficiency := getBoolInput(toolProficiencyResponse)

	return AbilityCheck(ability, strings.TrimSpace(circumstance), toolProficiency, skillProficiency)
}
