package domain

import (
	"testing"
)

func TestCooperationFactoryHelp(t *testing.T) {
	expected := Help
	result, _ := CooperationFactory("h")
	if result != expected {
		t.Errorf("Expected result %q was not equal to %q", expected, result)
	}
}

func TestCooperationFactoryGroup(t *testing.T) {
	expected := Group
	result, _ := CooperationFactory("g")
	if result != expected {
		t.Errorf("Expected result %q was not equal to %q", expected, result)
	}
}

func TestCooperationFactoryIndividual(t *testing.T) {
	expected := Individual
	result, _ := CooperationFactory("")
	if result != expected {
		t.Errorf("Expected result %q was not equal to %q", expected, result)
	}
}

func TestCooperationFactoryUnknown(t *testing.T) {
	result, err := CooperationFactory("unknown")
	if err == nil {
		t.Errorf("Expected error, actually got %q", result)
	}
}
