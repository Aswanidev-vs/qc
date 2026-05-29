package qc

import (
	"math"
	"testing"
)

const floatTolerance = 1e-9

func floatEq(a, b float64) bool {
	return math.Abs(a-b) < floatTolerance
}

func TestAdd(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
		{-3, -7, -10},
		{100, 200, 300},
	}
	for _, tt := range tests {
		if got := Add(tt.a, tt.b); got != tt.want {
			t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestSub(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{5, 3, 2},
		{0, 0, 0},
		{1, 5, -4},
	}
	for _, tt := range tests {
		if got := Sub(tt.a, tt.b); got != tt.want {
			t.Errorf("Sub(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestMul(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{3, 4, 12},
		{0, 100, 0},
		{-2, 3, -6},
	}
	for _, tt := range tests {
		if got := Mul(tt.a, tt.b); got != tt.want {
			t.Errorf("Mul(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestDiv(t *testing.T) {
	result, err := Div(10, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !floatEq(result, 10.0/3.0) {
		t.Errorf("Div(10, 3) = %f, want %f", result, 10.0/3.0)
	}

	_, err = Div(1, 0)
	if err != ErrDivByZero {
		t.Errorf("Div(1, 0) should return ErrDivByZero, got %v", err)
	}
}

func TestDivInt(t *testing.T) {
	q, r, err := DivInt(17, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if q != 3 || r != 2 {
		t.Errorf("DivInt(17, 5) = (%d, %d), want (3, 2)", q, r)
	}

	_, _, err = DivInt(1, 0)
	if err != ErrDivByZero {
		t.Errorf("DivInt(1, 0) should return ErrDivByZero")
	}
}

func TestMod(t *testing.T) {
	result, err := Mod(17, 5)
	if err != nil || result != 2 {
		t.Errorf("Mod(17, 5) = %d, err=%v, want 2, nil", result, err)
	}
	_, err = Mod(1, 0)
	if err != ErrZeroModulus {
		t.Error("Mod(1, 0) should return ErrZeroModulus")
	}
}

func TestPow(t *testing.T) {
	if !floatEq(Pow(2, 10), 1024) {
		t.Errorf("Pow(2, 10) = %f, want 1024", Pow(2, 10))
	}
	if !floatEq(Pow(3, 0), 1) {
		t.Errorf("Pow(3, 0) = %f, want 1", Pow(3, 0))
	}
}

func TestPowInt(t *testing.T) {
	result, err := PowInt(2, 10)
	if err != nil || result != 1024 {
		t.Errorf("PowInt(2, 10) = %d, err=%v, want 1024", result, err)
	}
	_, err = PowInt(2, -1)
	if err != ErrNegativeInput {
		t.Error("PowInt(2, -1) should return ErrNegativeInput")
	}
}

func TestAbs(t *testing.T) {
	if Abs(-5) != 5 {
		t.Error("Abs(-5) should be 5")
	}
	if Abs(0) != 0 {
		t.Error("Abs(0) should be 0")
	}
	if Abs(3) != 3 {
		t.Error("Abs(3) should be 3")
	}
}

func TestSqrt(t *testing.T) {
	result, err := Sqrt(144)
	if err != nil || !floatEq(result, 12) {
		t.Errorf("Sqrt(144) = %f, err=%v, want 12", result, err)
	}
	_, err = Sqrt(-1)
	if err != ErrNegativeInput {
		t.Error("Sqrt(-1) should return ErrNegativeInput")
	}
}

func TestSum(t *testing.T) {
	if Sum([]int{1, 2, 3, 4, 5}) != 15 {
		t.Error("Sum([1,2,3,4,5]) should be 15")
	}
	if Sum([]int{}) != 0 {
		t.Error("Sum([]) should be 0")
	}
}

func TestProduct(t *testing.T) {
	if Product([]int{2, 3, 4}) != 24 {
		t.Error("Product([2,3,4]) should be 24")
	}
}

func TestClamp(t *testing.T) {
	if Clamp(15, 0, 10) != 10 {
		t.Error("Clamp(15, 0, 10) should be 10")
	}
	if Clamp(-5, 0, 10) != 0 {
		t.Error("Clamp(-5, 0, 10) should be 0")
	}
	if Clamp(5, 0, 10) != 5 {
		t.Error("Clamp(5, 0, 10) should be 5")
	}
}

func TestMaxMinOf(t *testing.T) {
	max, err := MaxOf([]int{3, 1, 4, 1, 5, 9})
	if err != nil || max != 9 {
		t.Errorf("MaxOf = %d, want 9", max)
	}
	min, err := MinOf([]int{3, 1, 4, 1, 5, 9})
	if err != nil || min != 1 {
		t.Errorf("MinOf = %d, want 1", min)
	}
	_, err = MaxOf([]int{})
	if err != ErrInvalidInput {
		t.Error("MaxOf([]) should return error")
	}
}

func TestSign(t *testing.T) {
	if Sign(-5) != -1 || Sign(0) != 0 || Sign(5) != 1 {
		t.Error("Sign function failed")
	}
}

func TestHypot(t *testing.T) {
	if !floatEq(Hypot(3, 4), 5) {
		t.Errorf("Hypot(3, 4) = %f, want 5", Hypot(3, 4))
	}
}

func TestCbrt(t *testing.T) {
	if !floatEq(Cbrt(27), 3) {
		t.Errorf("Cbrt(27) = %f, want 3", Cbrt(27))
	}
}
