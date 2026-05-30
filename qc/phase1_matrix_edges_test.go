package qc

import "testing"

func TestPhase1_NewMatrix_Errors(t *testing.T) {
	_, err := NewMatrix(0, 2)
	if err != ErrInvalidInput {
		t.Fatalf("NewMatrix(0,2) expected ErrInvalidInput, got %v", err)
	}

	_, err = NewMatrix(2, -1)
	if err != ErrInvalidInput {
		t.Fatalf("NewMatrix(2,-1) expected ErrInvalidInput, got %v", err)
	}
}

func TestPhase1_NewMatrixFromSlice_Errors(t *testing.T) {
	_, err := NewMatrixFromSlice(nil)
	if err != ErrEmptySlice {
		t.Fatalf("NewMatrixFromSlice(nil) expected ErrEmptySlice, got %v", err)
	}

	_, err = NewMatrixFromSlice([][]float64{})
	if err != ErrEmptySlice {
		t.Fatalf("NewMatrixFromSlice([]) expected ErrEmptySlice, got %v", err)
	}

	_, err = NewMatrixFromSlice([][]float64{
		{1, 2},
		{3},
	})
	if err == nil {
		t.Fatalf("NewMatrixFromSlice with mismatched row lengths expected error, got nil")
	}
}

func TestPhase1_Matrix_Set_Errors(t *testing.T) {
	m, err := NewMatrixFromSlice([][]float64{{1, 2}, {3, 4}})
	if err != nil {
		t.Fatalf("unexpected setup error: %v", err)
	}

	if err := m.Set(-1, 0, 10); err != ErrInvalidInput {
		t.Fatalf("Set(-1,0) expected ErrInvalidInput, got %v", err)
	}
	if err := m.Set(0, -1, 10); err != ErrInvalidInput {
		t.Fatalf("Set(0,-1) expected ErrInvalidInput, got %v", err)
	}
	if err := m.Set(2, 0, 10); err != ErrInvalidInput {
		t.Fatalf("Set(2,0) expected ErrInvalidInput, got %v", err)
	}
	if err := m.Set(0, 2, 10); err != ErrInvalidInput {
		t.Fatalf("Set(0,2) expected ErrInvalidInput, got %v", err)
	}
}
