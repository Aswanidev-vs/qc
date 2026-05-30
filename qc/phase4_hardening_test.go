package qc

import "testing"

func TestLeastSquaresSolve_InvalidInputs(t *testing.T) {
	// A has m=0 via empty matrix
	_, err := LeastSquaresSolve(Matrix{}, []float64{})
	if err == nil {
		t.Fatalf("expected error for empty A")
	}

	// b length mismatch
	A, _ := NewMatrixFromSlice([][]float64{
		{1, 0},
		{2, 0},
	})
	_, err = LeastSquaresSolve(A, []float64{1})
	if err == nil {
		t.Fatalf("expected error for b length mismatch")
	}

	// A with m < n (underdetermined for our least-squares solver)
	A2, _ := NewMatrixFromSlice([][]float64{
		{1, 2, 3},
	})
	_, err = LeastSquaresSolve(A2, []float64{1})
	if err == nil {
		t.Fatalf("expected error for m < n")
	}
}
