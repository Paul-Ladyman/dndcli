package view

import (
	"fmt"
	"strings"

	"github.com/Paul-Ladyman/dndcli/domain"
)

func circumstanceMessage(circumstance string) string {
	if circumstance == "" {
		return ""
	}
	return fmt.Sprintf("with %s", circumstance)
}

func abilityMessage(ability string) string {
	return fmt.Sprintf("Make a %s check", ability)
}

// ExplainAbilityCheck will guide the DM through making the specified ability check
func ExplainAbilityCheck(summary domain.AbilityCheckSummary) {
	messages := [2]string{
		abilityMessage(summary.Ability.String()),
		circumstanceMessage(summary.Circumstance.String()),
	}
	message := strings.Join(messages[0:2], " ")

	fmt.Printf("\n%s\n", message)
	fmt.Println(summary.Throw.Summarise(summary.Dice))
	fmt.Printf("Add a %s modifier (abilities section on character sheet)\n", summary.Modifier.String())
	fmt.Println(summary.SummariseTarget())
}
