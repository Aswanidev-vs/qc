package qc

import (
	"math"
	"testing"
)

func TestMean(t *testing.T) {
	data := []float64{2, 4, 4, 4, 5, 5, 7, 9}
	mean, err := Mean(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !floatEq(mean, 5.0) {
		t.Errorf("Mean = %f, want 5.0", mean)
	}
	_, err = Mean([]float64{})
	if err != ErrEmptySlice {
		t.Error("Mean([]) should return ErrEmptySlice")
	}
}

func TestMedian(t *testing.T) {
	// Odd
	med, err := Median([]float64{1, 3, 5, 7, 9})
	if err != nil || !floatEq(med, 5) {
		t.Errorf("Median(odd) = %f, want 5", med)
	}
	// Even
	med, err = Median([]float64{1, 2, 3, 4})
	if err != nil || !floatEq(med, 2.5) {
		t.Errorf("Median(even) = %f, want 2.5", med)
	}
}

func TestMode(t *testing.T) {
	mode, err := Mode([]float64{1, 2, 2, 3, 3, 3, 4})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(mode) != 1 || mode[0] != 3 {
		t.Errorf("Mode = %v, want [3]", mode)
	}
}

func TestVarianceStdDev(t *testing.T) {
	data := []float64{2, 4, 4, 4, 5, 5, 7, 9}
	v, err := Variance(data)
	if err != nil {
		t.Fatalf("Variance error: %v", err)
	}
	if !floatEq(v, 4.0) {
		t.Errorf("Variance = %f, want 4.0", v)
	}
	sd, err := StdDev(data)
	if err != nil {
		t.Fatalf("StdDev error: %v", err)
	}
	if !floatEq(sd, 2.0) {
		t.Errorf("StdDev = %f, want 2.0", sd)
	}
}

func TestPercentile(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5}
	p50, err := Percentile(data, 50)
	if err != nil || !floatEq(p50, 3) {
		t.Errorf("Percentile(50) = %f, want 3", p50)
	}
}

func TestQuartiles(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8}
	q1, q2, q3, err := Quartiles(data)
	if err != nil {
		t.Fatalf("Quartiles error: %v", err)
	}
	if !floatEq(q2, 4.5) {
		t.Errorf("Q2 = %f, want 4.5", q2)
	}
	_ = q1
	_ = q3
}

func TestCovarianceCorrelation(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 5, 4, 5}
	corr, err := Correlation(x, y)
	if err != nil {
		t.Fatalf("Correlation error: %v", err)
	}
	if corr < 0.7 || corr > 1.0 {
		t.Errorf("Correlation = %f, expected ~0.83", corr)
	}
}

func TestZScore(t *testing.T) {
	z, err := ZScore(85, 70, 10)
	if err != nil || !floatEq(z, 1.5) {
		t.Errorf("ZScore(85, 70, 10) = %f, want 1.5", z)
	}
}

func TestMAD(t *testing.T) {
	data := []float64{2, 4, 4, 4, 5, 5, 7, 9}
	mad, err := MAD(data)
	if err != nil {
		t.Fatalf("MAD error: %v", err)
	}
	if !floatEq(mad, 1.5) {
		t.Errorf("MAD = %f, want 1.5", mad)
	}
}

func TestEntropy(t *testing.T) {
	// Uniform distribution over 2 outcomes: H = 1 bit
	e, err := Entropy([]float64{0.5, 0.5})
	if err != nil || !floatEq(e, 1.0) {
		t.Errorf("Entropy([0.5, 0.5]) = %f, want 1.0", e)
	}
	// Certain outcome: H = 0
	e, err = Entropy([]float64{1.0, 0.0})
	if err != nil || !floatEq(e, 0.0) {
		t.Errorf("Entropy([1, 0]) = %f, want 0.0", e)
	}
}

func TestRMS(t *testing.T) {
	rms, err := RMS([]float64{3, 4})
	if err != nil {
		t.Fatalf("RMS error: %v", err)
	}
	expected := math.Sqrt((9 + 16) / 2.0)
	if !floatEq(rms, expected) {
		t.Errorf("RMS([3,4]) = %f, want %f", rms, expected)
	}
}

func TestSkewness(t *testing.T) {
	// Symmetric data should have skewness ~ 0
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	skew, err := Skewness(data)
	if err != nil {
		t.Fatalf("Skewness error: %v", err)
	}
	if math.Abs(skew) > 0.01 {
		t.Errorf("Skewness of symmetric data = %f, expected ~0", skew)
	}
}
