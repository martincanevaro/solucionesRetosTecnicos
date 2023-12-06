package calc

import (
	"testing"
)

func TestAdd(t *testing.T) {
	// Arrange
	a := 1
	b := 2

	// Act
	result := Add(a, b)

	// Assert
	if result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
}

func TestSubtract(t *testing.T) {
	// Arrange
	a := 3
	b := 2

	// Act
	result := Subtract(a, b)

	// Assert
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}

func TestMultiply(t *testing.T) {
	// Arrange
	a := 2
	b := 3

	// Act
	result := Multiply(a, b)

	// Assert
	if result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}
}

func TestDivide(t *testing.T) {
	// Arrange
	a := 6
	b := 2

	// Act
	result := Divide(a, b)

	// Assert
	if result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
}
