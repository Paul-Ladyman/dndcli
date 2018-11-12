package dnd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Paul-Ladyman/dndcli/domain"
)

var reader = bufio.NewReader(os.Stdin)

func getBoolInput(response string) bool {
	if strings.TrimSpace(response) == "y" {
		return true
	}
	return false
}

func getSummaryAbility(ability domain.Ability, abilityErr error, skill domain.Skill, skillErr error) domain.Ability {
	if abilityErr == nil {
		return ability
	} else if skillErr == nil {
		return skill.Ability
	} else {
		return ""
	}
}

func getSummarySkill(skill domain.Skill, skillErr error) domain.Skill {
	if skillErr != nil {
		return domain.Skill{}
	}
	return skill
}

func getCircumstance(circumstanceString string, cooperation domain.Cooperation) (domain.Circumstance, error) {
	circumstance, _ := domain.CircumstanceFactory(circumstanceString)
	return circumstance.Resolve(cooperation), nil
}

func readCircumstance() (string, error) {
	fmt.Println("Does the player have Advantage (adv), Disadvantage (dis) or neither (press enter)?")
	return reader.ReadString('\n')
}

func readSkillProficiency(ability string) bool {
	skill, skillErr := domain.SkillFactory(ability)
	var skillProficiency = false
	if skillErr == nil {
		fmt.Printf("Is the player proficient in the %s skill?\n", skill)
		skillProficiencyResponse, _ := reader.ReadString('\n')
		skillProficiency = getBoolInput(skillProficiencyResponse)
	}
	return skillProficiency
}

func readToolProficiency() bool {
	fmt.Println("Is the player proficient with any relevant tools (y/n)?")
	toolProficiencyResponse, _ := reader.ReadString('\n')
	return getBoolInput(toolProficiencyResponse)
}

func readCooperation() (string, error) {
	fmt.Println("Does the player have help (h), are the party working as a group (g) or neither (press enter)?")
	return reader.ReadString('\n')
}

// AbilityCheck returns a summary of an ability check given an ability and an advantage circumstance
func AbilityCheck(abilityString string, circumstanceString string, toolProficiency bool, skillProficiency bool, cooperationString string) (domain.AbilityCheckSummary, error) {
	ability, abilityErr := domain.AbilityFactory(abilityString)
	skill, skillErr := domain.SkillFactory(abilityString)

	summaryAbility := getSummaryAbility(ability, abilityErr, skill, skillErr)
	summarySkill := getSummarySkill(skill, skillErr)

	cooperation, _ := domain.CooperationFactory(strings.TrimSpace(cooperationString))
	circumstance, _ := getCircumstance(circumstanceString, cooperation)

	return domain.AbilityCheckSummary{
		Ability:      summaryAbility,
		Modifier:     summaryAbility,
		Skill:        summarySkill,
		Circumstance: circumstance,
		Dice:         domain.D20,
		Roll:         domain.RollFactory(circumstance),
		Target:       domain.DC,
		Proficient:   toolProficiency || skillProficiency,
		Cooperation:  cooperation,
	}, nil
}

// AbilityCheckCli interactively retrieves information from the user relevant to ability checks
func AbilityCheckCli(ability string) (domain.AbilityCheckSummary, error) {
	circumstance, _ := readCircumstance()
	skillProficiency := readSkillProficiency(ability)
	toolProficiency := readToolProficiency()
	cooperation, _ := readCooperation()
	return AbilityCheck(ability, strings.TrimSpace(circumstance), toolProficiency, skillProficiency, cooperation)
}
