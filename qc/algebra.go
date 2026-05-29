package qc

import (
	"errors"
	"math"
)

var (
	// ErrNoRealRoots is returned when a quadratic equation has no real roots.
	ErrNoRealRoots = errors.New("no real roots")

	// ErrNotPositive is returned when input must be positive.
	ErrNotPositive = errors.New("input must be positive")

	// ErrEmptySlice is returned when a slice is empty but shouldn't be.
	ErrEmptySlice = errors.New("empty slice")
)

// GCD returns the greatest common divisor of two integers using the Euclidean algorithm.
func GCD(a, b int) int {
	a = Abs(a)
	b = Abs(b)
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM returns the least common multiple of two integers.
func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return Abs(a*b) / GCD(a, b)
}

// GCDMulti returns the GCD of a slice of integers.
// Returns ErrEmptySlice if the slice is empty.
func GCDMulti(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, ErrEmptySlice
	}
	result := nums[0]
	for _, n := range nums[1:] {
		result = GCD(result, n)
	}
	return result, nil
}

// LCMMulti returns the LCM of a slice of integers.
// Returns ErrEmptySlice if the slice is empty.
func LCMMulti(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, ErrEmptySlice
	}
	result := nums[0]
	for _, n := range nums[1:] {
		result = LCM(result, n)
	}
	return result, nil
}

// SolveQuadratic solves ax^2 + bx + c = 0 and returns the real roots.
// Returns ErrNoRealRoots if the discriminant is negative.
func SolveQuadratic(a, b, c float64) (x1, x2 float64, err error) {
	if a == 0 {
		return 0, 0, errors.New("not a quadratic equation (a=0)")
	}
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return 0, 0, ErrNoRealRoots
	}
	sqrtDisc := math.Sqrt(discriminant)
	x1 = (-b + sqrtDisc) / (2 * a)
	x2 = (-b - sqrtDisc) / (2 * a)
	return x1, x2, nil
}

// SolveLinear solves ax + b = 0 and returns x.
// Returns error if a is zero.
func SolveLinear(a, b float64) (float64, error) {
	if a == 0 {
		if b == 0 {
			return 0, errors.New("infinite solutions (0 = 0)")
		}
		return 0, errors.New("no solution (0 = non-zero)")
	}
	return -b / a, nil
}

// Fibonacci returns the nth Fibonacci number (0-indexed).
// Fibonacci(0) = 0, Fibonacci(1) = 1, Fibonacci(2) = 1, etc.
// Returns ErrNegativeInput if n is negative.
func Fibonacci(n int) (int, error) {
	if n < 0 {
		return 0, ErrNegativeInput
	}
	if n <= 1 {
		return n, nil
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b, nil
}

// FibonacciSequence returns the first n Fibonacci numbers.
// Returns ErrNegativeInput if n is negative.
func FibonacciSequence(n int) ([]int, error) {
	if n < 0 {
		return nil, ErrNegativeInput
	}
	if n == 0 {
		return []int{}, nil
	}
	seq := make([]int, n)
	if n >= 1 {
		seq[0] = 0
	}
	if n >= 2 {
		seq[1] = 1
	}
	for i := 2; i < n; i++ {
		seq[i] = seq[i-1] + seq[i-2]
	}
	return seq, nil
}

// IsEven returns true if n is even.
func IsEven(n int) bool {
	return n%2 == 0
}

// IsOdd returns true if n is odd.
func IsOdd(n int) bool {
	return n%2 != 0
}

// Factorial returns n! (n factorial).
// Returns ErrNegativeInput if n is negative.
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, ErrNegativeInput
	}
	if n <= 1 {
		return 1, nil
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result, nil
}

// DoubleFactorial returns n!! (n double factorial).
// For even n: n * (n-2) * (n-4) * ... * 2
// For odd n:  n * (n-2) * (n-4) * ... * 1
// Returns ErrNegativeInput if n is negative.
func DoubleFactorial(n int) (int, error) {
	if n < 0 {
		return 0, ErrNegativeInput
	}
	if n <= 1 {
		return 1, nil
	}
	result := 1
	for i := n; i > 0; i -= 2 {
		result *= i
	}
	return result, nil
}

// FallingFactorial returns n_(k) = n * (n-1) * ... * (n-k+1).
// Returns ErrNegativeInput if k is negative.
func FallingFactorial(n, k int) (int, error) {
	if k < 0 {
		return 0, ErrNegativeInput
	}
	if k == 0 {
		return 1, nil
	}
	result := 1
	for i := 0; i < k; i++ {
		result *= (n - i)
	}
	return result, nil
}

// RisingFactorial returns n^(k) = n * (n+1) * ... * (n+k-1).
// Returns ErrNegativeInput if k is negative.
func RisingFactorial(n, k int) (int, error) {
	if k < 0 {
		return 0, ErrNegativeInput
	}
	if k == 0 {
		return 1, nil
	}
	result := 1
	for i := 0; i < k; i++ {
		result *= (n + i)
	}
	return result, nil
}

// Lerp performs linear interpolation between a and b by factor t (0 to 1).
func Lerp(a, b, t float64) float64 {
	return a + t*(b-a)
}

// InvLerp finds the inverse: given value between a and b, returns t (0 to 1).
func InvLerp(a, b, value float64) float64 {
	if a == b {
		return 0
	}
	return (value - a) / (b - a)
}

// Remap maps value from range [inMin, inMax] to [outMin, outMax].
func Remap(value, inMin, inMax, outMin, outMax float64) float64 {
	t := InvLerp(inMin, inMax, value)
	return Lerp(outMin, outMax, t)
}

// ArithmeticMean returns the arithmetic mean of a slice of float64.
func ArithmeticMean(nums []float64) (float64, error) {
	if len(nums) == 0 {
		return 0, ErrEmptySlice
	}
	return SumFloat(nums) / float64(len(nums)), nil
}

// GeometricMean returns the geometric mean of a slice of positive float64.
// Returns ErrInvalidInput if any value is negative or the slice is empty.
func GeometricMean(nums []float64) (float64, error) {
	if len(nums) == 0 {
		return 0, ErrEmptySlice
	}
	product := 1.0
	for _, n := range nums {
		if n < 0 {
			return 0, ErrInvalidInput
		}
		product *= n
	}
	return math.Pow(product, 1.0/float64(len(nums))), nil
}

// HarmonicMean returns the harmonic mean of a slice of positive float64.
// Returns ErrInvalidInput if any value is zero or negative.
func HarmonicMean(nums []float64) (float64, error) {
	if len(nums) == 0 {
		return 0, ErrEmptySlice
	}
	reciprocalSum := 0.0
	for _, n := range nums {
		if n <= 0 {
			return 0, ErrInvalidInput
		}
		reciprocalSum += 1.0 / n
	}
	return float64(len(nums)) / reciprocalSum, nil
}

// WeightedMean returns the weighted mean given values and their weights.
// Returns ErrInvalidInput if slices have different lengths or are empty.
func WeightedMean(values, weights []float64) (float64, error) {
	if len(values) == 0 || len(weights) == 0 {
		return 0, ErrEmptySlice
	}
	if len(values) != len(weights) {
		return 0, ErrInvalidInput
	}
	var weightedSum, weightSum float64
	for i := range values {
		weightedSum += values[i] * weights[i]
		weightSum += weights[i]
	}
	if weightSum == 0 {
		return 0, ErrDivByZero
	}
	return weightedSum / weightSum, nil
}
