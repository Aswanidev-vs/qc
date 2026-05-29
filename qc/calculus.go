package qc

import (
	"errors"
	"math"
)

// Derivative computes the numerical derivative of f at x using the central difference method.
// h is the step size (use 1e-8 or similar for good accuracy).
func Derivative(f func(float64) float64, x, h float64) float64 {
	return (f(x+h) - f(x-h)) / (2 * h)
}

// SecondDerivative computes the second numerical derivative of f at x.
func SecondDerivative(f func(float64) float64, x, h float64) float64 {
	return (f(x+h) - 2*f(x) + f(x-h)) / (h * h)
}

// NthDerivative computes the nth numerical derivative using recursive central differences.
// Returns ErrInvalidInput if n < 1.
func NthDerivative(f func(float64) float64, x, h float64, n int) (float64, error) {
	if n < 1 {
		return 0, ErrInvalidInput
	}
	if n == 1 {
		return Derivative(f, x, h), nil
	}
	if n == 2 {
		return SecondDerivative(f, x, h), nil
	}
	// Recursive approach: f^(n)(x) ≈ (f^(n-1)(x+h) - f^(n-1)(x-h)) / (2h)
	// For practical purposes, use central difference of order n
	return centralDifferenceN(f, x, h, n), nil
}

func centralDifferenceN(f func(float64) float64, x, h float64, n int) float64 {
	coeffs := binomialCoeffs(n)
	sum := 0.0
	for i := 0; i <= n; i++ {
		sign := 1.0
		if (n-i)%2 == 1 {
			sign = -1.0
		}
		offset := (2*float64(i) - float64(n)) * h / 2
		sum += sign * float64(coeffs[i]) * f(x+offset)
	}
	return sum / math.Pow(h, float64(n))
}

func binomialCoeffs(n int) []int {
	coeffs := make([]int, n+1)
	coeffs[0] = 1
	for i := 1; i <= n; i++ {
		coeffs[i] = coeffs[i-1] * (n - i + 1) / i
	}
	return coeffs
}

// SimpsonIntegrate computes the definite integral of f from a to b using Simpson's 1/3 rule.
// n is the number of subintervals (must be even).
func SimpsonIntegrate(f func(float64) float64, a, b float64, n int) (float64, error) {
	if n <= 0 || n%2 != 0 {
		return 0, ErrInvalidInput
	}
	h := (b - a) / float64(n)
	sum := f(a) + f(b)

	for i := 1; i < n; i++ {
		x := a + float64(i)*h
		if i%2 == 0 {
			sum += 2 * f(x)
		} else {
			sum += 4 * f(x)
		}
	}
	return sum * h / 3, nil
}

// TrapezoidalIntegrate computes the definite integral of f from a to b using the trapezoidal rule.
func TrapezoidalIntegrate(f func(float64) float64, a, b float64, n int) (float64, error) {
	if n <= 0 {
		return 0, ErrInvalidInput
	}
	h := (b - a) / float64(n)
	sum := (f(a) + f(b)) / 2

	for i := 1; i < n; i++ {
		x := a + float64(i)*h
		sum += f(x)
	}
	return sum * h, nil
}

// MidpointIntegrate computes the definite integral of f from a to b using the midpoint rule.
func MidpointIntegrate(f func(float64) float64, a, b float64, n int) (float64, error) {
	if n <= 0 {
		return 0, ErrInvalidInput
	}
	h := (b - a) / float64(n)
	sum := 0.0

	for i := 0; i < n; i++ {
		x := a + (float64(i)+0.5)*h
		sum += f(x)
	}
	return sum * h, nil
}

// RombergIntegrate computes the integral using Romberg's method (Richardson extrapolation on trapezoidal).
// iterations controls the accuracy (higher = more accurate, 4-6 is usually sufficient).
func RombergIntegrate(f func(float64) float64, a, b float64, iterations int) (float64, error) {
	if iterations <= 0 {
		return 0, ErrInvalidInput
	}
	R := make([][]float64, iterations)
	for i := range R {
		R[i] = make([]float64, iterations)
	}

	// First approximation using trapezoidal rule
	h := b - a
	R[0][0] = h * (f(a) + f(b)) / 2

	for i := 1; i < iterations; i++ {
		// Composite trapezoidal with 2^i subintervals
		h /= 2
		sum := 0.0
		for k := 1; k <= (1 << (i - 1)); k++ {
			sum += f(a + float64(2*k-1)*h)
		}
		R[i][0] = R[i-1][0]/2 + h*sum

		// Richardson extrapolation
		for j := 1; j <= i; j++ {
			R[i][j] = R[i][j-1] + (R[i][j-1]-R[i-1][j-1])/(math.Pow(4, float64(j))-1)
		}
	}
	return R[iterations-1][iterations-1], nil
}

// Limit computes the limit of f as x approaches target from both sides.
// Returns the average of left and right limits.
func Limit(f func(float64) float64, target float64) float64 {
	h := 1e-8
	left := f(target - h)
	right := f(target + h)
	return (left + right) / 2
}

// LimitFromLeft computes the limit of f as x approaches target from the left.
func LimitFromLeft(f func(float64) float64, target, h float64) float64 {
	return f(target - h)
}

// LimitFromRight computes the limit of f as x approaches target from the right.
func LimitFromRight(f func(float64) float64, target, h float64) float64 {
	return f(target + h)
}

// NewtonRaphson finds a root of f near x0 using the Newton-Raphson method.
// maxIter is the maximum number of iterations, tol is the convergence tolerance.
func NewtonRaphson(f func(float64) float64, x0, tol float64, maxIter int) (float64, error) {
	if maxIter <= 0 {
		return 0, ErrInvalidInput
	}
	h := 1e-10
	x := x0
	for i := 0; i < maxIter; i++ {
		fx := f(x)
		if math.Abs(fx) < tol {
			return x, nil
		}
		df := Derivative(f, x, h)
		if math.Abs(df) < 1e-15 {
			return 0, errors.New("derivative is zero; Newton-Raphson failed")
		}
		x = x - fx/df
	}
	return x, errors.New("Newton-Raphson did not converge within max iterations")
}

// BisectionMethod finds a root of f in the interval [a, b].
// Requires f(a) and f(b) to have opposite signs.
func BisectionMethod(f func(float64) float64, a, b, tol float64, maxIter int) (float64, error) {
	if maxIter <= 0 {
		return 0, ErrInvalidInput
	}
	fa := f(a)
	fb := f(b)
	if fa*fb > 0 {
		return 0, errors.New("f(a) and f(b) must have opposite signs")
	}
	for i := 0; i < maxIter; i++ {
		mid := (a + b) / 2
		fmid := f(mid)
		if math.Abs(fmid) < tol || (b-a)/2 < tol {
			return mid, nil
		}
		if fa*fmid < 0 {
			b = mid
			fb = fmid
		} else {
			a = mid
			fa = fmid
		}
	}
	return (a + b) / 2, errors.New("bisection did not converge within max iterations")
}

// SecantMethod finds a root of f using the secant method.
func SecantMethod(f func(float64) float64, x0, x1, tol float64, maxIter int) (float64, error) {
	if maxIter <= 0 {
		return 0, ErrInvalidInput
	}
	for i := 0; i < maxIter; i++ {
		fx0 := f(x0)
		fx1 := f(x1)
		if math.Abs(fx1) < tol {
			return x1, nil
		}
		denom := fx1 - fx0
		if math.Abs(denom) < 1e-15 {
			return 0, errors.New("secant method failed: division by near-zero")
		}
		x2 := x1 - fx1*(x1-x0)/denom
		x0 = x1
		x1 = x2
	}
	return x1, errors.New("secant method did not converge within max iterations")
}

// TaylorPolynomial evaluates the Taylor polynomial of f centered at a, at point x.
// n is the order of the polynomial. Derivatives are computed iteratively using central differences.
func TaylorPolynomial(f func(float64) float64, a, x float64, n int) (float64, error) {
	if n < 0 {
		return 0, ErrInvalidInput
	}
	dx := x - a
	sum := 0.0

	// Compute k-th derivative at point a by applying Derivative k times
	derivs := make([]float64, n+1)
	derivs[0] = f(a)

	// For each successive derivative, we need the function values
	// Compute using progressively smaller step sizes for stability
	for k := 1; k <= n; k++ {
		// Use the original function and compute k-th derivative directly
		// using the binomial coefficient central difference formula
		h := math.Pow(1e-16, 1.0/float64(2*k+1)) // optimal step size
		coeffs := binomialCoeffs(k)
		dk := 0.0
		for j := 0; j <= k; j++ {
			sign := 1.0
			if (k-j)%2 == 1 {
				sign = -1.0
			}
			offset := (2*float64(j) - float64(k)) * h / 2
			dk += sign * float64(coeffs[j]) * f(a+offset)
		}
		derivs[k] = dk / math.Pow(h, float64(k))
	}

	for k := 0; k <= n; k++ {
		factK, _ := FactorialFloat(k)
		sum += derivs[k] * math.Pow(dx, float64(k)) / factK
	}
	return sum, nil
}

// EulerMethod solves the ODE y' = f(x, y) with initial condition y(x0) = y0.
// Returns a slice of (x, y) pairs from x0 to xEnd with step size h.
func EulerMethod(f func(x, y float64) float64, x0, y0, xEnd, h float64) ([][2]float64, error) {
	if h <= 0 || xEnd <= x0 {
		return nil, ErrInvalidInput
	}
	var result [][2]float64
	x, y := x0, y0
	result = append(result, [2]float64{x, y})
	for x < xEnd {
		y += h * f(x, y)
		x += h
		if x > xEnd {
			x = xEnd
		}
		result = append(result, [2]float64{x, y})
	}
	return result, nil
}

// RungeKutta4 solves the ODE y' = f(x, y) using the classic 4th-order Runge-Kutta method.
// Returns a slice of (x, y) pairs from x0 to xEnd with step size h.
func RungeKutta4(f func(x, y float64) float64, x0, y0, xEnd, h float64) ([][2]float64, error) {
	if h <= 0 || xEnd <= x0 {
		return nil, ErrInvalidInput
	}
	var result [][2]float64
	x, y := x0, y0
	result = append(result, [2]float64{x, y})
	for x < xEnd {
		k1 := h * f(x, y)
		k2 := h * f(x+h/2, y+k1/2)
		k3 := h * f(x+h/2, y+k2/2)
		k4 := h * f(x+h, y+k3)
		y += (k1 + 2*k2 + 2*k3 + k4) / 6
		x += h
		if x > xEnd {
			x = xEnd
		}
		result = append(result, [2]float64{x, y})
	}
	return result, nil
}
