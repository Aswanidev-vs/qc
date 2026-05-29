package qc

import (
	"math"
	"testing"
)

func TestSinCosTan(t *testing.T) {
	tests := []struct {
		x       float64
		sin, cos float64
	}{
		{0, 0, 1},
		{Pi / 6, 0.5, math.Sqrt(3) / 2},
		{Pi / 4, math.Sqrt(2) / 2, math.Sqrt(2) / 2},
		{Pi / 3, math.Sqrt(3) / 2, 0.5},
		{Pi / 2, 1, 0},
		{Pi, 0, -1},
	}
	for _, tt := range tests {
		if !floatEq(Sin(tt.x), tt.sin) {
			t.Errorf("Sin(%f) = %f, want %f", tt.x, Sin(tt.x), tt.sin)
		}
		if !floatEq(Cos(tt.x), tt.cos) {
			t.Errorf("Cos(%f) = %f, want %f", tt.x, Cos(tt.x), tt.cos)
		}
	}
	// tan(pi/4) = 1
	if !floatEq(Tan(Pi/4), 1) {
		t.Errorf("Tan(π/4) = %f, want 1", Tan(Pi/4))
	}
}

func TestASinACos(t *testing.T) {
	v, err := ASin(0.5)
	if err != nil {
		t.Fatalf("ASin(0.5) error: %v", err)
	}
	if !floatEq(v, Pi/6) {
		t.Errorf("ASin(0.5) = %f, want %f", v, Pi/6)
	}

	v, err = ACos(0.5)
	if err != nil {
		t.Fatalf("ACos(0.5) error: %v", err)
	}
	if !floatEq(v, Pi/3) {
		t.Errorf("ACos(0.5) = %f, want %f", v, Pi/3)
	}

	_, err = ASin(2)
	if err != ErrInvalidInput {
		t.Error("ASin(2) should return ErrInvalidInput")
	}
}

func TestATan(t *testing.T) {
	if !floatEq(ATan(1), Pi/4) {
		t.Errorf("ATan(1) = %f, want %f", ATan(1), Pi/4)
	}
}

func TestDegreesRadians(t *testing.T) {
	if !floatEq(Degrees(Pi), 180) {
		t.Errorf("Degrees(π) = %f, want 180", Degrees(Pi))
	}
	if !floatEq(Radians(180), Pi) {
		t.Errorf("Radians(180) = %f, want π", Radians(180))
	}
}

func TestSinDCosD(t *testing.T) {
	if !floatEq(SinD(30), 0.5) {
		t.Errorf("SinD(30) = %f, want 0.5", SinD(30))
	}
	if !floatEq(CosD(60), 0.5) {
		t.Errorf("CosD(60) = %f, want 0.5", CosD(60))
	}
}

func TestWrapAngle(t *testing.T) {
	tests := []struct{ in, want float64}{
		{0, 0},
		{Pi, Pi},
		{-Pi, -Pi},
		{3 * Pi, Pi},
		{-3 * Pi, -Pi},
	}
	for _, tt := range tests {
		got := WrapAngle(tt.in)
		if !floatEq(got, tt.want) {
			t.Errorf("WrapAngle(%f) = %f, want %f", tt.in, got, tt.want)
		}
	}
}

func TestHyperbolic(t *testing.T) {
	if !floatEq(Sinh(0), 0) {
		t.Error("Sinh(0) should be 0")
	}
	if !floatEq(Cosh(0), 1) {
		t.Error("Cosh(0) should be 1")
	}
	if !floatEq(Tanh(0), 0) {
		t.Error("Tanh(0) should be 0")
	}
}

func TestCotSecCsc(t *testing.T) {
	cot, err := Cot(Pi / 4)
	if err != nil || !floatEq(cot, 1) {
		t.Errorf("Cot(π/4) = %f, err=%v, want 1", cot, err)
	}
	sec, err := Sec(0)
	if err != nil || !floatEq(sec, 1) {
		t.Errorf("Sec(0) = %f, err=%v, want 1", sec, err)
	}
	csc, err := Csc(Pi / 2)
	if err != nil || !floatEq(csc, 1) {
		t.Errorf("Csc(π/2) = %f, err=%v, want 1", csc, err)
	}
}
