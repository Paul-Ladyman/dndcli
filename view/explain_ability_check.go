package view

import (
	"fmt"

	"github.com/Paul-Ladyman/dndcli/domain"
)

// ExplainAbilityCheck will guide the DM through making the specified ability check
func ExplainAbilityCheck(summary domain.AbilityCheckSummary) {
	fmt.Printf("%s:\n", summary.SummariseAbilityAndCircumstance())
	fmt.Println(summary.SummariseThrow())
	fmt.Println(summary.SummariseModifier())
	fmt.Println(summary.SummariseTarget())
}
