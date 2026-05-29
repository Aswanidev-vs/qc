package qc

import "errors"

// Matrix is a 2D slice of float64 values.
type Matrix [][]float64

// NewMatrix creates a new matrix with the given number of rows and columns,
// initialized to zero.
func NewMatrix(rows, cols int) (Matrix, error) {
	if rows <= 0 || cols <= 0 {
		return nil, ErrInvalidInput
	}
	m := make(Matrix, rows)
	for i := range m {
		m[i] = make([]float64, cols)
	}
	return m, nil
}

// NewMatrixFromSlice creates a matrix from a 2D slice.
func NewMatrixFromSlice(data [][]float64) (Matrix, error) {
	if len(data) == 0 {
		return nil, ErrEmptySlice
	}
	cols := len(data[0])
	if cols == 0 {
		return nil, ErrInvalidInput
	}
	for _, row := range data {
		if len(row) != cols {
			return nil, errors.New("all rows must have the same number of columns")
		}
	}
	m := make(Matrix, len(data))
	for i, row := range data {
		m[i] = make([]float64, cols)
		copy(m[i], row)
	}
	return m, nil
}

// Identity returns an identity matrix of size n x n.
func Identity(n int) (Matrix, error) {
	if n <= 0 {
		return nil, ErrInvalidInput
	}
	m, _ := NewMatrix(n, n)
	for i := 0; i < n; i++ {
		m[i][i] = 1
	}
	return m, nil
}

// Rows returns the number of rows.
func (m Matrix) Rows() int {
	return len(m)
}

// Cols returns the number of columns.
func (m Matrix) Cols() int {
	if len(m) == 0 {
		return 0
	}
	return len(m[0])
}

// Get returns the value at position (row, col).
func (m Matrix) Get(row, col int) (float64, error) {
	if row < 0 || row >= m.Rows() || col < 0 || col >= m.Cols() {
		return 0, ErrInvalidInput
	}
	return m[row][col], nil
}

// Set sets the value at position (row, col).
func (m Matrix) Set(row, col int, val float64) error {
	if row < 0 || row >= m.Rows() || col < 0 || col >= m.Cols() {
		return ErrInvalidInput
	}
	m[row][col] = val
	return nil
}

// Add adds two matrices element-wise.
func (m Matrix) Add(other Matrix) (Matrix, error) {
	if m.Rows() != other.Rows() || m.Cols() != other.Cols() {
		return nil, errors.New("matrices must have the same dimensions")
	}
	result, _ := NewMatrix(m.Rows(), m.Cols())
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			result[i][j] = m[i][j] + other[i][j]
		}
	}
	return result, nil
}

// Sub subtracts another matrix element-wise.
func (m Matrix) Sub(other Matrix) (Matrix, error) {
	if m.Rows() != other.Rows() || m.Cols() != other.Cols() {
		return nil, errors.New("matrices must have the same dimensions")
	}
	result, _ := NewMatrix(m.Rows(), m.Cols())
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			result[i][j] = m[i][j] - other[i][j]
		}
	}
	return result, nil
}

// ScalarMul multiplies every element by a scalar.
func (m Matrix) ScalarMul(scalar float64) Matrix {
	result, _ := NewMatrix(m.Rows(), m.Cols())
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			result[i][j] = m[i][j] * scalar
		}
	}
	return result
}

// Mul multiplies two matrices.
func (m Matrix) Mul(other Matrix) (Matrix, error) {
	if m.Cols() != other.Rows() {
		return nil, errors.New("incompatible dimensions for multiplication")
	}
	result, _ := NewMatrix(m.Rows(), other.Cols())
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < other.Cols(); j++ {
			sum := 0.0
			for k := 0; k < m.Cols(); k++ {
				sum += m[i][k] * other[k][j]
			}
			result[i][j] = sum
		}
	}
	return result, nil
}

// Transpose returns the transpose of the matrix.
func (m Matrix) Transpose() Matrix {
	result, _ := NewMatrix(m.Cols(), m.Rows())
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			result[j][i] = m[i][j]
		}
	}
	return result
}

// Determinant returns the determinant of a square matrix.
func (m Matrix) Determinant() (float64, error) {
	if m.Rows() != m.Cols() {
		return 0, errors.New("matrix must be square")
	}
	n := m.Rows()
	if n == 1 {
		return m[0][0], nil
	}
	if n == 2 {
		return m[0][0]*m[1][1] - m[0][1]*m[1][0], nil
	}
	// LU decomposition via cofactor expansion
	det := 0.0
	for j := 0; j < n; j++ {
		minor, _ := m.minor(0, j)
		minorDet, _ := minor.Determinant()
		sign := 1.0
		if j%2 == 1 {
			sign = -1.0
		}
		det += sign * m[0][j] * minorDet
	}
	return det, nil
}

// minor returns the minor matrix obtained by removing row r and column c.
func (m Matrix) minor(r, c int) (Matrix, error) {
	if r < 0 || r >= m.Rows() || c < 0 || c >= m.Cols() {
		return nil, ErrInvalidInput
	}
	result, _ := NewMatrix(m.Rows()-1, m.Cols()-1)
	ri := 0
	for i := 0; i < m.Rows(); i++ {
		if i == r {
			continue
		}
		ci := 0
		for j := 0; j < m.Cols(); j++ {
			if j == c {
				continue
			}
			result[ri][ci] = m[i][j]
			ci++
		}
		ri++
	}
	return result, nil
}

// Trace returns the trace (sum of diagonal elements) of a square matrix.
func (m Matrix) Trace() (float64, error) {
	if m.Rows() != m.Cols() {
		return 0, errors.New("matrix must be square")
	}
	trace := 0.0
	for i := 0; i < m.Rows(); i++ {
		trace += m[i][i]
	}
	return trace, nil
}

// Inverse returns the inverse of a square matrix using Gauss-Jordan elimination.
func (m Matrix) Inverse() (Matrix, error) {
	if m.Rows() != m.Cols() {
		return nil, errors.New("matrix must be square")
	}
	n := m.Rows()
	// Create augmented matrix [m | I]
	aug, _ := NewMatrix(n, 2*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			aug[i][j] = m[i][j]
		}
		aug[i][n+i] = 1
	}
	// Forward elimination
	for i := 0; i < n; i++ {
		// Find pivot
		maxVal := aug[i][i]
		maxRow := i
		for k := i + 1; k < n; k++ {
			if abs(aug[k][i]) > abs(maxVal) {
				maxVal = aug[k][i]
				maxRow = k
			}
		}
		if abs(maxVal) < 1e-12 {
			return nil, errors.New("matrix is singular (not invertible)")
		}
		aug[i], aug[maxRow] = aug[maxRow], aug[i]
		// Scale pivot row
		pivot := aug[i][i]
		for j := 0; j < 2*n; j++ {
			aug[i][j] /= pivot
		}
		// Eliminate column
		for k := 0; k < n; k++ {
			if k == i {
				continue
			}
			factor := aug[k][i]
			for j := 0; j < 2*n; j++ {
				aug[k][j] -= factor * aug[i][j]
			}
		}
	}
	// Extract inverse
	inv, _ := NewMatrix(n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			inv[i][j] = aug[i][n+j]
		}
	}
	return inv, nil
}

// Rank returns the rank of the matrix.
func (m Matrix) Rank() (int, error) {
	if m.Rows() == 0 || m.Cols() == 0 {
		return 0, ErrEmptySlice
	}
	// Copy the matrix
	a := make(Matrix, m.Rows())
	for i := range m {
		a[i] = make([]float64, m.Cols())
		copy(a[i], m[i])
	}
	rank := 0
	for col := 0; col < m.Cols(); col++ {
		// Find pivot
		pivotRow := -1
		for row := rank; row < m.Rows(); row++ {
			if abs(a[row][col]) > 1e-12 {
				pivotRow = row
				break
			}
		}
		if pivotRow == -1 {
			continue
		}
		a[rank], a[pivotRow] = a[pivotRow], a[rank]
		pivot := a[rank][col]
		for j := 0; j < m.Cols(); j++ {
			a[rank][j] /= pivot
		}
		for i := 0; i < m.Rows(); i++ {
			if i == rank {
				continue
			}
			factor := a[i][col]
			for j := 0; j < m.Cols(); j++ {
				a[i][j] -= factor * a[rank][j]
			}
		}
		rank++
	}
	return rank, nil
}

// IsSquare returns true if the matrix is square.
func (m Matrix) IsSquare() bool {
	return m.Rows() == m.Cols()
}

// IsDiagonal returns true if the matrix is diagonal.
func (m Matrix) IsDiagonal() bool {
	if !m.IsSquare() {
		return false
	}
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			if i != j && abs(m[i][j]) > 1e-12 {
				return false
			}
		}
	}
	return true
}

// IsSymmetric returns true if the matrix equals its transpose.
func (m Matrix) IsSymmetric() bool {
	if !m.IsSquare() {
		return false
	}
	for i := 0; i < m.Rows(); i++ {
		for j := i + 1; j < m.Cols(); j++ {
			if abs(m[i][j]-m[j][i]) > 1e-12 {
				return false
			}
		}
	}
	return true
}

// abs returns the absolute value of a float64 (unexported helper).
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
