package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Paul-Ladyman/dndcli/dnd"
	"github.com/Paul-Ladyman/dndcli/view"
)

func main() {
	var ability = flag.String("abcheck", "", "Explain how to perform a check on the specified ability")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Does the player have Advantage (adv), Disadvantage (dis) or neither (press enter)?")
	circumstance, _ := reader.ReadString('\n')

	abilityCheckSummary, _ := dnd.AbilityCheck(*ability, strings.TrimSpace(circumstance))

	view.ExplainAbilityCheck(abilityCheckSummary)
}
