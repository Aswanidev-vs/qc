package qc

import (
	"math"
	"testing"
)

func TestDerivative(t *testing.T) {
	// d/dx(x^2) at x=3 should be 6
	f := func(x float64) float64 { return x * x }
	d := Derivative(f, 3, 1e-8)
	if math.Abs(d-6) > 1e-4 {
		t.Errorf("d/dx(x²) at 3 = %f, want ~6", d)
	}

	// d/dx(sin(x)) at x=0 should be 1
	d = Derivative(math.Sin, 0, 1e-8)
	if math.Abs(d-1) > 1e-4 {
		t.Errorf("d/dx(sin) at 0 = %f, want ~1", d)
	}
}

func TestSecondDerivative(t *testing.T) {
	// d²/dx²(x^3) at x=2 should be 12
	f := func(x float64) float64 { return x * x * x }
	d2 := SecondDerivative(f, 2, 1e-5)
	if math.Abs(d2-12) > 1e-2 {
		t.Errorf("d²/dx²(x³) at 2 = %f, want ~12", d2)
	}
}

func TestSimpsonIntegrate(t *testing.T) {
	// ∫₀¹ x² dx = 1/3
	f := func(x float64) float64 { return x * x }
	result, err := SimpsonIntegrate(f, 0, 1, 100)
	if err != nil {
		t.Fatalf("SimpsonIntegrate error: %v", err)
	}
	if math.Abs(result-1.0/3.0) > 1e-8 {
		t.Errorf("∫₀¹ x² dx = %f, want ~0.3333", result)
	}

	// ∫₀^π sin(x) dx = 2
	result, err = SimpsonIntegrate(math.Sin, 0, Pi, 100)
	if err != nil {
		t.Fatalf("SimpsonIntegrate error: %v", err)
	}
	if math.Abs(result-2) > 1e-6 {
		t.Errorf("∫₀^π sin(x) dx = %f, want ~2", result)
	}
}

func TestTrapezoidalIntegrate(t *testing.T) {
	f := func(x float64) float64 { return x * x }
	result, err := TrapezoidalIntegrate(f, 0, 1, 10000)
	if err != nil {
		t.Fatalf("TrapezoidalIntegrate error: %v", err)
	}
	if math.Abs(result-1.0/3.0) > 1e-4 {
		t.Errorf("Trapezoidal ∫₀¹ x² dx = %f, want ~0.3333", result)
	}
}

func TestRombergIntegrate(t *testing.T) {
	f := func(x float64) float64 { return x * x }
	result, err := RombergIntegrate(f, 0, 1, 6)
	if err != nil {
		t.Fatalf("RombergIntegrate error: %v", err)
	}
	if math.Abs(result-1.0/3.0) > 1e-10 {
		t.Errorf("Romberg ∫₀¹ x² dx = %f, want ~0.3333333333", result)
	}
}

func TestNewtonRaphson(t *testing.T) {
	// Find sqrt(2): solve x^2 - 2 = 0
	f := func(x float64) float64 { return x*x - 2 }
	root, err := NewtonRaphson(f, 1.5, 1e-10, 100)
	if err != nil {
		t.Fatalf("NewtonRaphson error: %v", err)
	}
	if math.Abs(root-math.Sqrt(2)) > 1e-8 {
		t.Errorf("Newton-Raphson √2 = %f, want %f", root, math.Sqrt(2))
	}
}

func TestBisectionMethod(t *testing.T) {
	f := func(x float64) float64 { return x*x - 2 }
	root, err := BisectionMethod(f, 0, 2, 1e-10, 100)
	if err != nil {
		t.Fatalf("BisectionMethod error: %v", err)
	}
	if math.Abs(root-math.Sqrt(2)) > 1e-8 {
		t.Errorf("Bisection √2 = %f, want %f", root, math.Sqrt(2))
	}
}

func TestSecantMethod(t *testing.T) {
	f := func(x float64) float64 { return x*x - 2 }
	root, err := SecantMethod(f, 0, 2, 1e-10, 100)
	if err != nil {
		t.Fatalf("SecantMethod error: %v", err)
	}
	if math.Abs(root-math.Sqrt(2)) > 1e-8 {
		t.Errorf("Secant √2 = %f, want %f", root, math.Sqrt(2))
	}
}

func TestMidpointIntegrate(t *testing.T) {
	f := func(x float64) float64 { return x }
	result, err := MidpointIntegrate(f, 0, 1, 1000)
	if err != nil {
		t.Fatalf("MidpointIntegrate error: %v", err)
	}
	if math.Abs(result-0.5) > 1e-4 {
		t.Errorf("Midpoint ∫₀¹ x dx = %f, want 0.5", result)
	}
}

func TestLimit(t *testing.T) {
	// lim(x→0) sin(x)/x = 1
	f := func(x float64) float64 {
		if x == 0 {
			return 1 // handle singularity
		}
		return math.Sin(x) / x
	}
	lim := Limit(f, 0)
	if math.Abs(lim-1) > 1e-4 {
		t.Errorf("lim sin(x)/x = %f, want 1", lim)
	}
}

func TestEulerMethod(t *testing.T) {
	// y' = y, y(0) = 1 → y = e^x
	f := func(x, y float64) float64 { return y }
	result, err := EulerMethod(f, 0, 1, 1, 0.01)
	if err != nil {
		t.Fatalf("EulerMethod error: %v", err)
	}
	last := result[len(result)-1]
	expected := math.E
	if math.Abs(last[1]-expected) > 0.1 {
		t.Errorf("Euler y(1) = %f, want ~%f", last[1], expected)
	}
}

func TestRungeKutta4(t *testing.T) {
	// y' = y, y(0) = 1 → y = e^x
	f := func(x, y float64) float64 { return y }
	result, err := RungeKutta4(f, 0, 1, 1, 0.01)
	if err != nil {
		t.Fatalf("RungeKutta4 error: %v", err)
	}
	last := result[len(result)-1]
	expected := math.E
	if math.Abs(last[1]-expected) > 1e-6 {
		t.Errorf("RK4 y(1) = %f, want ~%f", last[1], expected)
	}
}

func TestTaylorPolynomial(t *testing.T) {
	// Taylor expansion of e^x at x=0, evaluated at x=1 with order 10 ≈ e
	f := math.Exp
	result, err := TaylorPolynomial(f, 0, 1, 10)
	if err != nil {
		t.Fatalf("TaylorPolynomial error: %v", err)
	}
	if math.Abs(result-math.E) > 0.01 {
		t.Errorf("Taylor(e^x, 0, 1, 10) = %f, want ~%f", result, math.E)
	}
}
