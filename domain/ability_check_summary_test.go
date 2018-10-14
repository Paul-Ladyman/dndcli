package domain

import (
	"testing"
)

func TestSummariseTarget(t *testing.T) {
	summary := AbilityCheckSummary{
		Strength,
		Strength,
		Neutral,
		D20,
		Throw{1, Higher},
		DC,
	}
	expected := "Compare against the Difficulty Class (DC): Easy - 10, Medium - 15, Hard - 20"
	result := summary.SummariseTarget()
	if result != expected {
		t.Errorf(" %q was not equal to %q", result, expected)
	}
}
