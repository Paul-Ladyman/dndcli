package domain

import "errors"

// Skill is a generic type for skills
type Skill struct {
	Ability Ability
	Name    string
}

var (
	Athletics Skill = Skill{Ability: Strength, Name: "ath"}
)

func (skill Skill) String() string {
	switch skill {
	case Athletics:
		return "athletics"
	default:
		return ""
	}
}

// SkillFactory takes a string and returns the appropriate Skill
func SkillFactory(skill string) (Skill, error) {
	switch skill {
	case Athletics.Name:
		return Athletics, nil
	default:
		return Skill{}, errors.New("Invalid skill")
	}
}
