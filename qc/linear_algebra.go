package qc

import "math"

// NormL1Vec2 returns the L1 norm of a 2D vector: |x| + |y|.
func NormL1Vec2(v Vec2) float64 {
	return math.Abs(v.X) + math.Abs(v.Y)
}

// NormLinfVec2 returns the L-infinity norm of a 2D vector: max(|x|, |y|).
func NormLinfVec2(v Vec2) float64 {
	ax := math.Abs(v.X)
	ay := math.Abs(v.Y)
	if ax > ay {
		return ax
	}
	return ay
}

// NormL2Vec2 returns the L2 norm of a 2D vector (magnitude).
func NormL2Vec2(v Vec2) float64 {
	return v.Magnitude()
}

// NormL1Vec3 returns the L1 norm of a 3D vector: |x| + |y| + |z|.
func NormL1Vec3(v Vec3) float64 {
	return math.Abs(v.X) + math.Abs(v.Y) + math.Abs(v.Z)
}

// NormLinfVec3 returns the L-infinity norm of a 3D vector: max(|x|, |y|, |z|).
func NormLinfVec3(v Vec3) float64 {
	ax := math.Abs(v.X)
	ay := math.Abs(v.Y)
	az := math.Abs(v.Z)

	m := ax
	if ay > m {
		m = ay
	}
	if az > m {
		m = az
	}
	return m
}

// NormL2Vec3 returns the L2 norm of a 3D vector (magnitude).
func NormL2Vec3(v Vec3) float64 {
	return v.Magnitude()
}

// GramSchmidtOrthonormalizeVec3 takes 3 input vectors and returns an orthonormal basis
// using the Gram-Schmidt process.
// It returns ErrInvalidInput if vectors are linearly dependent / cannot form a basis.
func GramSchmidtOrthonormalizeVec3(a, b, c Vec3) ([3]Vec3, error) {
	u1 := a
	n1 := u1.Magnitude()
	if n1 < Eps {
		return [3]Vec3{}, ErrInvalidInput
	}
	e1 := u1.Scale(1 / n1)

	// u2 = b - proj_{e1}(b)
	proj2, err := b.Project(e1)
	if err != nil {
		return [3]Vec3{}, err
	}
	u2 := b.Sub(proj2)
	n2 := u2.Magnitude()
	if n2 < Eps {
		return [3]Vec3{}, ErrInvalidInput
	}
	e2 := u2.Scale(1 / n2)

	// u3 = c - proj_{e1}(c) - proj_{e2}(c)
	proj3e1, err := c.Project(e1)
	if err != nil {
		return [3]Vec3{}, err
	}
	proj3e2, err := c.Project(e2)
	if err != nil {
		return [3]Vec3{}, err
	}
	u3 := c.Sub(proj3e1).Sub(proj3e2)
	n3 := u3.Magnitude()
	if n3 < Eps {
		return [3]Vec3{}, ErrInvalidInput
	}
	e3 := u3.Scale(1 / n3)

	return [3]Vec3{e1, e2, e3}, nil
}

// LeastSquaresSolve solves an overdetermined linear system A x ≈ b in the
// least-squares sense using normal equations:
// x = (A^T A)^(-1) A^T b
//
// It expects A to have shape (m x n) where m >= n and b length m.
// Returns ErrInvalidInput for mismatched dimensions or empty inputs.
func LeastSquaresSolve(A Matrix, b []float64) ([]float64, error) {
	m := A.Rows()
	if m == 0 {
		return nil, ErrInvalidInput
	}
	n := A.Cols()
	if n == 0 {
		return nil, ErrInvalidInput
	}
	if len(b) != m {
		return nil, ErrInvalidInput
	}
	if m < n {
		return nil, ErrInvalidInput
	}

	AT := A.Transpose()   // (n x m)
	ATA, err := AT.Mul(A) // (n x n)
	if err != nil {
		return nil, err
	}

	// Compute A^T b as an (n x 1) matrix.
	ATbData := make([][]float64, n)
	for i := 0; i < n; i++ {
		ATbData[i] = []float64{0}
		for j := 0; j < m; j++ {
			ATbData[i][0] += AT[i][j] * b[j]
		}
	}
	ATb, err := NewMatrixFromSlice(ATbData)
	if err != nil {
		return nil, err
	}

	ATAinv, err := ATA.Inverse()
	if err != nil {
		return nil, err
	}

	xMat, err := ATAinv.Mul(ATb) // (n x 1)
	if err != nil {
		return nil, err
	}

	x := make([]float64, n)
	for i := 0; i < n; i++ {
		x[i] = xMat[i][0]
	}
	return x, nil
}
