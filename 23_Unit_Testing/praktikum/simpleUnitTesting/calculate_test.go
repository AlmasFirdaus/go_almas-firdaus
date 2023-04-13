package simpleUnitTesting

import (
	"testing"
)

func TestAddition(t *testing.T) {
	if Additon(1, 2) != 3 {
		t.Error("Expected 1 (+) 2 to equal 3")
	}
	if Additon(5, 5) != 10 {
		t.Error("Expected 5 (+) 5 to equal 10")
	}
}

func TestSubtraction(t *testing.T) {
	if Subtraction(5, 2) != 3 {
		t.Error("Expected 5 (-) 2 to equal 3")
	}
	if Subtraction(10, 5) != 5 {
		t.Error("Expected 10 (-) 5 to equal 5")
	}
}

func TestDivision(t *testing.T) {
	if Division(10, 2) != 5 {
		t.Error("Expected 10 (/) 2 to equal 5")
	}
	if Division(6, 3) != 2 {
		t.Error("Expected 6 (/) 2 to equal 2")
	}
}

func TestMultiplication(t *testing.T) {
	if Multiplication(5, 2) != 10 {
		t.Error("Expected 5 (*) 2 to equal 10")
	}
	if Multiplication(10, 5) != 50 {
		t.Error("Expected 10 (*) 5 to equal 50")
	}
}
