package qc

import (
	"math"
	"testing"
)

func TestGCD(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{48, 18, 6},
		{0, 5, 5},
		{5, 0, 5},
		{7, 13, 1},
		{100, 75, 25},
	}
	for _, tt := range tests {
		if got := GCD(tt.a, tt.b); got != tt.want {
			t.Errorf("GCD(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestLCM(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{4, 6, 12},
		{3, 7, 21},
		{0, 5, 0},
	}
	for _, tt := range tests {
		if got := LCM(tt.a, tt.b); got != tt.want {
			t.Errorf("LCM(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestSolveQuadratic(t *testing.T) {
	x1, x2, err := SolveQuadratic(1, -5, 6) // x^2 - 5x + 6 = 0 → roots 3, 2
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	roots := map[float64]bool{x1: true, x2: true}
	if !roots[3] || !roots[2] {
		t.Errorf("SolveQuadratic(1,-5,6) roots = %f, %f, want 3 and 2", x1, x2)
	}

	_, _, err = SolveQuadratic(1, 0, 1) // x^2 + 1 = 0 → no real roots
	if err != ErrNoRealRoots {
		t.Error("should return ErrNoRealRoots for x^2+1=0")
	}

	_, _, err = SolveQuadratic(0, 1, 1)
	if err == nil {
		t.Error("should return error for a=0")
	}
}

func TestSolveLinear(t *testing.T) {
	x, err := SolveLinear(2, -6) // 2x - 6 = 0 → x = 3
	if err != nil || !floatEq(x, 3) {
		t.Errorf("SolveLinear(2, -6) = %f, want 3", x)
	}
	_, err = SolveLinear(0, 1)
	if err == nil {
		t.Error("should return error for 0x+1=0")
	}
}

func TestFibonacci(t *testing.T) {
	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	for i, want := range expected {
		got, err := Fibonacci(i)
		if err != nil || got != want {
			t.Errorf("Fibonacci(%d) = %d, want %d", i, got, want)
		}
	}
	_, err := Fibonacci(-1)
	if err != ErrNegativeInput {
		t.Error("Fibonacci(-1) should return ErrNegativeInput")
	}
}

func TestFibonacciSequence(t *testing.T) {
	seq, err := FibonacciSequence(10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	for i, v := range seq {
		if v != expected[i] {
			t.Errorf("FibonacciSequence[%d] = %d, want %d", i, v, expected[i])
		}
	}
}

func TestIsEvenOdd(t *testing.T) {
	if !IsEven(4) || IsEven(3) {
		t.Error("IsEven failed")
	}
	if !IsOdd(3) || IsOdd(4) {
		t.Error("IsOdd failed")
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct{ n, want int }{
		{0, 1}, {1, 1}, {5, 120}, {10, 3628800},
	}
	for _, tt := range tests {
		got, err := Factorial(tt.n)
		if err != nil || got != tt.want {
			t.Errorf("Factorial(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
	_, err := Factorial(-1)
	if err != ErrNegativeInput {
		t.Error("Factorial(-1) should return ErrNegativeInput")
	}
}

func TestLerp(t *testing.T) {
	if !floatEq(Lerp(0, 100, 0.5), 50) {
		t.Error("Lerp(0, 100, 0.5) should be 50")
	}
	if !floatEq(Lerp(0, 100, 0), 0) {
		t.Error("Lerp(0, 100, 0) should be 0")
	}
	if !floatEq(Lerp(0, 100, 1), 100) {
		t.Error("Lerp(0, 100, 1) should be 100")
	}
}

func TestWeightedMean(t *testing.T) {
	vals := []float64{3, 5, 7}
	weights := []float64{2, 3, 1}
	result, err := WeightedMean(vals, weights)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := (3*2 + 5*3 + 7*1) / float64(2+3+1)
	if !floatEq(result, expected) {
		t.Errorf("WeightedMean = %f, want %f", result, expected)
	}
}

func TestGeometricMean(t *testing.T) {
	result, err := GeometricMean([]float64{4, 16, 64})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := math.Pow(4*16*64, 1.0/3.0)
	if !floatEq(result, expected) {
		t.Errorf("GeometricMean = %f, want %f", result, expected)
	}
}

func TestHarmonicMean(t *testing.T) {
	result, err := HarmonicMean([]float64{1, 2, 4})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := 3.0 / (1.0/1 + 1.0/2 + 1.0/4)
	if !floatEq(result, expected) {
		t.Errorf("HarmonicMean = %f, want %f", result, expected)
	}
}
