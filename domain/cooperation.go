package domain

import "errors"

// Cooperation is a generic type for describing whether a player has help in their activity
type Cooperation string

// Cooperation by being helped by individual party members, by working as a group, or just working alone
const (
	Help       Cooperation = "help"
	Group      Cooperation = "group"
	Individual Cooperation = "individual"
)

func CooperationFactory(cooperation string) (Cooperation, error) {
	switch cooperation {
	case "h":
		return Help, nil
	case "g":
		return Group, nil
	case "":
		return Individual, nil
	default:
		return "", errors.New("Invalid cooperation value")
	}
}
