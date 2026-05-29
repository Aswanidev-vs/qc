package qc

import (
	"errors"
	"math"
)

var (
	// ErrDivByZero is returned when attempting to divide by zero.
	ErrDivByZero = errors.New("division by zero")

	// ErrOverflow is returned when an integer operation overflows.
	ErrOverflow = errors.New("integer overflow")

	// ErrNegativeInput is returned when a negative input is not allowed.
	ErrNegativeInput = errors.New("negative input not allowed")

	// ErrZeroModulus is returned when attempting modulo by zero.
	ErrZeroModulus = errors.New("modulo by zero")

	// ErrInvalidInput is returned for general invalid input.
	ErrInvalidInput = errors.New("invalid input")
)

// Add returns the sum of two integers.
func Add(a, b int) int {
	return a + b
}

// AddFloat returns the sum of two float64 values.
func AddFloat(a, b float64) float64 {
	return a + b
}

// Sub returns the difference of two integers (a - b).
func Sub(a, b int) int {
	return a - b
}

// SubFloat returns the difference of two float64 values (a - b).
func SubFloat(a, b float64) float64 {
	return a - b
}

// Mul returns the product of two integers.
func Mul(a, b int) int {
	return a * b
}

// MulFloat returns the product of two float64 values.
func MulFloat(a, b float64) float64 {
	return a * b
}

// Div returns the quotient of two float64 values.
// Returns ErrDivByZero if divisor is zero.
func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}
	return a / b, nil
}

// DivInt returns the integer quotient and remainder of a / b.
// Returns ErrDivByZero if b is zero.
func DivInt(a, b int) (quotient, remainder int, err error) {
	if b == 0 {
		return 0, 0, ErrDivByZero
	}
	return a / b, a % b, nil
}

// Mod returns the modulus (remainder) of a % b.
// Returns ErrZeroModulus if b is zero.
func Mod(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrZeroModulus
	}
	return a % b, nil
}

// ModFloat returns the floating-point remainder of a / b.
// Returns ErrDivByZero if b is zero.
func ModFloat(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}
	return math.Mod(a, b), nil
}

// Pow returns a raised to the power of b (a^b) as float64.
func Pow(a, b float64) float64 {
	return math.Pow(a, b)
}

// PowInt returns a raised to the power of b (a^b) as int.
// Returns ErrNegativeInput if b is negative.
func PowInt(a, b int) (int, error) {
	if b < 0 {
		return 0, ErrNegativeInput
	}
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result, nil
}

// Abs returns the absolute value of an integer.
func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// AbsFloat returns the absolute value of a float64.
func AbsFloat(a float64) float64 {
	return math.Abs(a)
}

// Negate returns the negation of an integer.
func Negate(a int) int {
	return -a
}

// NegateFloat returns the negation of a float64.
func NegateFloat(a float64) float64 {
	return -a
}

// Max returns the larger of two integers.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min returns the smaller of two integers.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MaxFloat returns the larger of two float64 values.
func MaxFloat(a, b float64) float64 {
	return math.Max(a, b)
}

// MinFloat returns the smaller of two float64 values.
func MinFloat(a, b float64) float64 {
	return math.Min(a, b)
}

// Clamp constrains value to the range [min, max].
func Clamp(value, minVal, maxVal int) int {
	if value < minVal {
		return minVal
	}
	if value > maxVal {
		return maxVal
	}
	return value
}

// ClampFloat constrains value to the range [min, max].
func ClampFloat(value, minVal, maxVal float64) float64 {
	return math.Max(minVal, math.Min(maxVal, value))
}

// Sqrt returns the square root of a non-negative float64.
// Returns ErrNegativeInput if a is negative.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, ErrNegativeInput
	}
	return math.Sqrt(a), nil
}

// Cbrt returns the cube root of a.
func Cbrt(a float64) float64 {
	return math.Cbrt(a)
}

// Hypot returns sqrt(p*p + q*q), the hypotenuse of a right triangle.
func Hypot(p, q float64) float64 {
	return math.Hypot(p, q)
}

// Sum returns the sum of a slice of integers.
func Sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// SumFloat returns the sum of a slice of float64 values.
func SumFloat(nums []float64) float64 {
	total := 0.0
	for _, n := range nums {
		total += n
	}
	return total
}

// Product returns the product of a slice of integers.
func Product(nums []int) int {
	result := 1
	for _, n := range nums {
		result *= n
	}
	return result
}

// ProductFloat returns the product of a slice of float64 values.
func ProductFloat(nums []float64) float64 {
	result := 1.0
	for _, n := range nums {
		result *= n
	}
	return result
}

// MaxOf returns the maximum value in a slice of integers.
// Returns ErrInvalidInput if the slice is empty.
func MaxOf(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, ErrInvalidInput
	}
	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}
	return max, nil
}

// MinOf returns the minimum value in a slice of integers.
// Returns ErrInvalidInput if the slice is empty.
func MinOf(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, ErrInvalidInput
	}
	min := nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
	}
	return min, nil
}

// MaxOfFloat returns the maximum value in a slice of float64.
// Returns ErrInvalidInput if the slice is empty.
func MaxOfFloat(nums []float64) (float64, error) {
	if len(nums) == 0 {
		return 0, ErrInvalidInput
	}
	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}
	return max, nil
}

// MinOfFloat returns the minimum value in a slice of float64.
// Returns ErrInvalidInput if the slice is empty.
func MinOfFloat(nums []float64) (float64, error) {
	if len(nums) == 0 {
		return 0, ErrInvalidInput
	}
	min := nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
	}
	return min, nil
}

// Sign returns -1 for negative, 0 for zero, 1 for positive.
func Sign(a int) int {
	if a < 0 {
		return -1
	}
	if a > 0 {
		return 1
	}
	return 0
}

// SignFloat returns -1.0 for negative, 0.0 for zero, 1.0 for positive.
func SignFloat(a float64) float64 {
	if a < 0 {
		return -1.0
	}
	if a > 0 {
		return 1.0
	}
	return 0.0
}
