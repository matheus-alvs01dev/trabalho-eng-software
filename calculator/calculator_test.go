package calculator

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b     float64
		expected float64
	}{
		{2, 3, 5},
		{-1, 1, 0},
		{0, 0, 0},
	}

	for _, tt := range tests {
		result := Add(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Add(%v, %v) = %v; esperado %v", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		a, b     float64
		expected float64
	}{
		{5, 2, 3},
		{2, 5, -3},
		{0, 0, 0},
	}

	for _, tt := range tests {
		result := Subtract(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Subtract(%v, %v) = %v; esperado %v", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		a, b     float64
		expected float64
	}{
		{3, 4, 12},
		{-2, 5, -10},
		{0, 99, 0},
	}

	for _, tt := range tests {
		result := Multiply(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Multiply(%v, %v) = %v; esperado %v", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		a, b     float64
		expected float64
		wantErr  bool
	}{
		{10, 2, 5, false},
		{5, 0, 0, true},
		{-10, 2, -5, false},
	}

	for _, tt := range tests {
		result, err := Divide(tt.a, tt.b)
		if (err != nil) != tt.wantErr {
			t.Errorf("Divide(%v, %v) erro = %v; esperado erro? %v", tt.a, tt.b, err, tt.wantErr)
		}
		if !tt.wantErr && result != tt.expected {
			t.Errorf("Divide(%v, %v) = %v; esperado %v", tt.a, tt.b, result, tt.expected)
		}
	}
}
