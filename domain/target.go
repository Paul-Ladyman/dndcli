package domain

type Target string

const (
	DC Target = "DC"
	AC Target = "AC"
)

func (target Target) Summarise() string {
	switch target {
	case DC:
		return "the Difficulty Class (DC): Easy - 10, Medium - 15, Hard - 20"
	default:
		return "the character's Armor Class (AC on character sheet)"
	}
}
