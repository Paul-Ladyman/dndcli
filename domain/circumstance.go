package domain

import (
	"errors"
)

// Circumstance is a generic type of the advantage circumstances
type Circumstance string

// Advantage circumstances
const (
	Advantage    Circumstance = "adv"
	Disadvantage Circumstance = "dis"
	Neutral      Circumstance = ""
)

func (circumstance Circumstance) String() string {
	switch circumstance {
	case Advantage:
		return "advantage"
	case Disadvantage:
		return "disadvantage"
	default:
		return ""
	}
}

func increment(circumstance Circumstance) Circumstance {
	switch circumstance {
	case Disadvantage:
		return Neutral
	case Neutral:
		return Advantage
	default:
		return circumstance
	}
}

func (circumstance Circumstance) Resolve(cooperation Cooperation) Circumstance {
	if cooperation == Help {
		return increment(circumstance)
	}
	return circumstance
}

// CircumstanceFactory takes a string and returns the appropriate advantage circumstance
func CircumstanceFactory(circumstance string) (Circumstance, error) {
	switch circumstance {
	case string(Advantage):
		return Advantage, nil
	case string(Disadvantage):
		return Disadvantage, nil
	case string(Neutral):
		return Neutral, nil
	default:
		return "", errors.New("Invalid circumstance")
	}
}
