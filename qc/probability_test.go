package qc

import (
	"math"
	"testing"
)

func TestPermutation(t *testing.T) {
	// P(10, 3) = 720
	p, err := Permutation(10, 3)
	if err != nil || p != 720 {
		t.Errorf("Permutation(10, 3) = %d, want 720", p)
	}
	// P(5, 0) = 1
	p, err = Permutation(5, 0)
	if err != nil || p != 1 {
		t.Errorf("Permutation(5, 0) = %d, want 1", p)
	}
}

func TestCombination(t *testing.T) {
	// C(10, 3) = 120
	c, err := Combination(10, 3)
	if err != nil || c != 120 {
		t.Errorf("Combination(10, 3) = %d, want 120", c)
	}
	// C(5, 5) = 1
	c, err = Combination(5, 5)
	if err != nil || c != 1 {
		t.Errorf("Combination(5, 5) = %d, want 1", c)
	}
	// C(n, 0) = 1
	c, err = Combination(100, 0)
	if err != nil || c != 1 {
		t.Errorf("Combination(100, 0) = %d, want 1", c)
	}
}

func TestBinomialPMF(t *testing.T) {
	// P(X=5) for Binomial(10, 0.5) = C(10,5) * 0.5^10 ≈ 0.2461
	pmf, err := BinomialPMF(10, 5, 0.5)
	if err != nil {
		t.Fatalf("BinomialPMF error: %v", err)
	}
	if math.Abs(pmf-0.24609375) > 1e-6 {
		t.Errorf("BinomialPMF(10,5,0.5) = %f, want ~0.2461", pmf)
	}
}

func TestBinomialCDF(t *testing.T) {
	cdf, err := BinomialCDF(10, 5, 0.5)
	if err != nil {
		t.Fatalf("BinomialCDF error: %v", err)
	}
	if cdf < 0.5 || cdf > 1.0 {
		t.Errorf("BinomialCDF(10,5,0.5) = %f, expected in [0.5, 1.0]", cdf)
	}
}

func TestPoissonPMF(t *testing.T) {
	// P(X=2) for Poisson(3) = e^-3 * 3^2 / 2! ≈ 0.2240
	pmf, err := PoissonPMF(3, 2)
	if err != nil {
		t.Fatalf("PoissonPMF error: %v", err)
	}
	expected := math.Exp(-3) * 9 / 2
	if !floatEq(pmf, expected) {
		t.Errorf("PoissonPMF(3,2) = %f, want %f", pmf, expected)
	}
}

func TestNormalPDF(t *testing.T) {
	// Standard normal at 0: 1/sqrt(2π) ≈ 0.3989
	pdf, err := NormalPDF(0, 0, 1)
	if err != nil {
		t.Fatalf("NormalPDF error: %v", err)
	}
	expected := 1.0 / math.Sqrt(2*Pi)
	if !floatEq(pdf, expected) {
		t.Errorf("NormalPDF(0,0,1) = %f, want %f", pdf, expected)
	}
}

func TestNormalCDF(t *testing.T) {
	cdf, err := NormalCDF(0, 0, 1)
	if err != nil {
		t.Fatalf("NormalCDF error: %v", err)
	}
	if !floatEq(cdf, 0.5) {
		t.Errorf("NormalCDF(0,0,1) = %f, want 0.5", cdf)
	}
}

func TestExponentialPDF(t *testing.T) {
	pdf, err := ExponentialPDF(1, 1) // lambda=1, x=1
	if err != nil {
		t.Fatalf("ExponentialPDF error: %v", err)
	}
	expected := math.Exp(-1)
	if !floatEq(pdf, expected) {
		t.Errorf("ExponentialPDF(1,1) = %f, want %f", pdf, expected)
	}
}

func TestGeometricPMF(t *testing.T) {
	// P(X=1) with p=0.5 = 0.5
	pmf, err := GeometricPMF(0.5, 1)
	if err != nil || !floatEq(pmf, 0.5) {
		t.Errorf("GeometricPMF(0.5, 1) = %f, want 0.5", pmf)
	}
	// P(X=3) with p=0.5 = (0.5)^2 * 0.5 = 0.125
	pmf, err = GeometricPMF(0.5, 3)
	if err != nil || !floatEq(pmf, 0.125) {
		t.Errorf("GeometricPMF(0.5, 3) = %f, want 0.125", pmf)
	}
}

func TestBernoulliTrial(t *testing.T) {
	p, _ := BernoulliTrial(0.7, 1)
	if !floatEq(p, 0.7) {
		t.Errorf("BernoulliTrial(0.7, 1) = %f, want 0.7", p)
	}
	p, _ = BernoulliTrial(0.7, 0)
	if !floatEq(p, 0.3) {
		t.Errorf("BernoulliTrial(0.7, 0) = %f, want 0.3", p)
	}
}

func TestHypergeometricPMF(t *testing.T) {
	// N=50, K=5, n=10, k=1
	pmf, err := HypergeometricPMF(50, 5, 10, 1)
	if err != nil {
		t.Fatalf("HypergeometricPMF error: %v", err)
	}
	if pmf < 0 || pmf > 1 {
		t.Errorf("HypergeometricPMF = %f, should be in [0,1]", pmf)
	}
}

func TestLogFactorial(t *testing.T) {
	lf, err := LogFactorial(10)
	if err != nil {
		t.Fatalf("LogFactorial error: %v", err)
	}
	expected := math.Log(3628800.0)
	if !floatEq(lf, expected) {
		t.Errorf("LogFactorial(10) = %f, want %f", lf, expected)
	}
}
