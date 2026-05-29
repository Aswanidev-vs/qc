package qc

import (
	"testing"
)

func TestMatrixAddSub(t *testing.T) {
	m1, _ := NewMatrixFromSlice([][]float64{{1, 2}, {3, 4}})
	m2, _ := NewMatrixFromSlice([][]float64{{5, 6}, {7, 8}})
	sum, err := m1.Add(m2)
	if err != nil {
		t.Fatalf("Add error: %v", err)
	}
	if !floatEq(sum[0][0], 6) || !floatEq(sum[1][1], 12) {
		t.Errorf("Add result incorrect: %v", sum)
	}
	diff, _ := m1.Sub(m2)
	if !floatEq(diff[0][0], -4) || !floatEq(diff[1][1], -4) {
		t.Errorf("Sub result incorrect: %v", diff)
	}
}

func TestMatrixMul(t *testing.T) {
	m1, _ := NewMatrixFromSlice([][]float64{{1, 2}, {3, 4}})
	m2, _ := NewMatrixFromSlice([][]float64{{5, 6}, {7, 8}})
	product, err := m1.Mul(m2)
	if err != nil {
		t.Fatalf("Mul error: %v", err)
	}
	// [1,2;3,4] * [5,6;7,8] = [19,22;43,50]
	if !floatEq(product[0][0], 19) || !floatEq(product[0][1], 22) ||
		!floatEq(product[1][0], 43) || !floatEq(product[1][1], 50) {
		t.Errorf("Mul result = %v, want [[19,22],[43,50]]", product)
	}
}

func TestMatrixDeterminant(t *testing.T) {
	m, _ := NewMatrixFromSlice([][]float64{{1, 2}, {3, 4}})
	det, err := m.Determinant()
	if err != nil {
		t.Fatalf("Determinant error: %v", err)
	}
	if !floatEq(det, -2) {
		t.Errorf("det = %f, want -2", det)
	}

	// 3x3
	m3, _ := NewMatrixFromSlice([][]float64{{6, 1, 1}, {4, -2, 5}, {2, 8, 7}})
	det3, _ := m3.Determinant()
	if !floatEq(det3, -306) {
		t.Errorf("det(3x3) = %f, want -306", det3)
	}
}

func TestMatrixTranspose(t *testing.T) {
	m, _ := NewMatrixFromSlice([][]float64{{1, 2, 3}, {4, 5, 6}})
	tm := m.Transpose()
	if tm.Rows() != 3 || tm.Cols() != 2 {
		t.Errorf("Transpose dimensions = %dx%d, want 3x2", tm.Rows(), tm.Cols())
	}
	if !floatEq(tm[0][0], 1) || !floatEq(tm[0][1], 4) {
		t.Errorf("Transpose values incorrect")
	}
}

func TestMatrixInverse(t *testing.T) {
	m, _ := NewMatrixFromSlice([][]float64{{4, 7}, {2, 6}})
	inv, err := m.Inverse()
	if err != nil {
		t.Fatalf("Inverse error: %v", err)
	}
	// Verify M * M^-1 ≈ I
	product, _ := m.Mul(inv)
	if !floatEq(product[0][0], 1) || !floatEq(product[0][1], 0) ||
		!floatEq(product[1][0], 0) || !floatEq(product[1][1], 1) {
		t.Errorf("M * M^-1 != I: %v", product)
	}

	// Singular matrix
	singular, _ := NewMatrixFromSlice([][]float64{{1, 2}, {2, 4}})
	_, err = singular.Inverse()
	if err == nil {
		t.Error("Singular matrix inverse should fail")
	}
}

func TestMatrixIdentity(t *testing.T) {
	id, _ := Identity(3)
	if !floatEq(id[0][0], 1) || !floatEq(id[1][1], 1) || !floatEq(id[2][2], 1) {
		t.Error("Identity diagonal should be 1")
	}
	if !floatEq(id[0][1], 0) || !floatEq(id[1][0], 0) {
		t.Error("Identity off-diagonal should be 0")
	}
}

func TestMatrixTrace(t *testing.T) {
	m, _ := NewMatrixFromSlice([][]float64{{1, 2}, {3, 4}})
	trace, _ := m.Trace()
	if !floatEq(trace, 5) {
		t.Errorf("Trace = %f, want 5", trace)
	}
}

func TestMatrixRank(t *testing.T) {
	m, _ := NewMatrixFromSlice([][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 0}})
	rank, _ := m.Rank()
	if rank != 2 {
		t.Errorf("Rank = %d, want 2", rank)
	}
}

func TestMatrixIsSymmetric(t *testing.T) {
	sym, _ := NewMatrixFromSlice([][]float64{{1, 2}, {2, 3}})
	if !sym.IsSymmetric() {
		t.Error("symmetric matrix should return true")
	}
	nonsym, _ := NewMatrixFromSlice([][]float64{{1, 2}, {3, 4}})
	if nonsym.IsSymmetric() {
		t.Error("non-symmetric matrix should return false")
	}
}

func TestMatrixScalarMul(t *testing.T) {
	m, _ := NewMatrixFromSlice([][]float64{{1, 2}, {3, 4}})
	result := m.ScalarMul(3)
	if !floatEq(result[0][0], 3) || !floatEq(result[1][1], 12) {
		t.Errorf("ScalarMul result incorrect: %v", result)
	}
}
