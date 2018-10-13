package domain

// Dice is a generic class for the various D&D dice
type Dice string

// D&D dice
const (
	D20 Dice = "d20"
	D12 Dice = "d12"
	D10 Dice = "d10"
	D8  Dice = "d8"
	D6  Dice = "d6"
	D4  Dice = "d4"
)
