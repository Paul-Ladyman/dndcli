package dnd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Paul-Ladyman/dndcli/domain"
)

func getToolProficiency(response string) bool {
	if strings.TrimSpace(response) == "y" {
		return true
	}
	return false
}

// AbilityCheck returns a summary of an ability check given an ability and an advantage circumstance
func AbilityCheck(abilityString string, circumstanceString string, toolProficiency bool) (domain.AbilityCheckSummary, error) {
	circumstance, _ := domain.CircumstanceFactory(circumstanceString)
	ability, _ := domain.AbilityFactory(abilityString)
	return domain.AbilityCheckSummary{
		Ability:      ability,
		Modifier:     ability,
		Circumstance: circumstance,
		Dice:         domain.D20,
		Roll:         domain.RollFactory(circumstance),
		Target:       domain.DC,
		Proficient:   toolProficiency,
	}, nil
}

// AbilityCheckCli interactively retrieves information from the user relevant to ability checks
func AbilityCheckCli(ability string) (domain.AbilityCheckSummary, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Does the player have Advantage (adv), Disadvantage (dis) or neither (press enter)?")
	circumstance, _ := reader.ReadString('\n')

	fmt.Println("Is the player proficient in any relevant tools (y/n)?")
	toolProficiencyResponse, _ := reader.ReadString('\n')

	toolProficiency := getToolProficiency(toolProficiencyResponse)

	return AbilityCheck(ability, strings.TrimSpace(circumstance), toolProficiency)
}
