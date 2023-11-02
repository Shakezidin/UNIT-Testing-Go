package math

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 5)
	if result != 6 {
		t.Errorf("Add 1 and 5 failed, expected 6, but got %d", result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(7, 3)
	if result != 4 {
		t.Errorf("Subtract 7 and 3 failed, expected 4, but got %d", result)
	}
}

func TestDivision(t *testing.T) {
	result := Division(10, 2)
	if result != 5.0 {
		t.Errorf("Division 10 by 2 failed, expected 5.0, but got %f", result)
	}
}

func TestMultiplication(t *testing.T) {
	result := Multiplication(3, 4)
	if result != 12 {
		t.Errorf("Multiplication of 3 and 4 failed, expected 12, but got %d", result)
	}
}
