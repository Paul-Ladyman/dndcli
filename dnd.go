package main

import (
	"flag"

	"github.com/Paul-Ladyman/dndcli/dnd"
	"github.com/Paul-Ladyman/dndcli/view"
)

func main() {
	var ability = flag.String("abcheck", "", "Explain how to perform a check on the specified ability")
	flag.Parse()

	abilityCheckSummary, _ := dnd.AbilityCheckCli(*ability)
	view.ExplainAbilityCheck(abilityCheckSummary)
}
