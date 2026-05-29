package qc

import "math"

// Permutation returns P(n, k) = n! / (n-k)! — the number of ways to arrange k items from n.
// Returns ErrInvalidInput if k > n or either is negative.
func Permutation(n, k int) (int, error) {
	if n < 0 || k < 0 || k > n {
		return 0, ErrInvalidInput
	}
	result := 1
	for i := 0; i < k; i++ {
		result *= (n - i)
	}
	return result, nil
}

// Combination returns C(n, k) = n! / (k! * (n-k)!) — the number of ways to choose k items from n.
// Returns ErrInvalidInput if k > n or either is negative.
func Combination(n, k int) (int, error) {
	if n < 0 || k < 0 || k > n {
		return 0, ErrInvalidInput
	}
	if k > n-k {
		k = n - k
	}
	result := 1
	for i := 0; i < k; i++ {
		result *= (n - i)
		result /= (i + 1)
	}
	return result, nil
}

// PermutationFloat returns P(n, k) as float64 for large values.
func PermutationFloat(n, k int) (float64, error) {
	if n < 0 || k < 0 || k > n {
		return 0, ErrInvalidInput
	}
	result := 1.0
	for i := 0; i < k; i++ {
		result *= float64(n - i)
	}
	return result, nil
}

// CombinationFloat returns C(n, k) as float64 for large values.
func CombinationFloat(n, k int) (float64, error) {
	if n < 0 || k < 0 || k > n {
		return 0, ErrInvalidInput
	}
	if k > n-k {
		k = n - k
	}
	result := 1.0
	for i := 0; i < k; i++ {
		result *= float64(n - i)
		result /= float64(i + 1)
	}
	return result, nil
}

// BinomialPMF returns the binomial probability P(X = k) for parameters n and p.
func BinomialPMF(n, k int, p float64) (float64, error) {
	if n < 0 || k < 0 || k > n {
		return 0, ErrInvalidInput
	}
	if p < 0 || p > 1 {
		return 0, ErrInvalidInput
	}
	comb, _ := CombinationFloat(n, k)
	return comb * math.Pow(p, float64(k)) * math.Pow(1-p, float64(n-k)), nil
}

// BinomialCDF returns P(X <= k) for binomial distribution.
func BinomialCDF(n int, k int, p float64) (float64, error) {
	if n < 0 || k < 0 {
		return 0, ErrInvalidInput
	}
	if p < 0 || p > 1 {
		return 0, ErrInvalidInput
	}
	if k >= n {
		return 1.0, nil
	}
	sum := 0.0
	for i := 0; i <= k; i++ {
		pmf, _ := BinomialPMF(n, i, p)
		sum += pmf
	}
	return sum, nil
}

// PoissonPMF returns the Poisson probability P(X = k) for parameter lambda.
func PoissonPMF(lambda float64, k int) (float64, error) {
	if lambda < 0 || k < 0 {
		return 0, ErrInvalidInput
	}
	factK, _ := Factorial(k)
	return math.Exp(-lambda) * math.Pow(lambda, float64(k)) / float64(factK), nil
}

// PoissonCDF returns P(X <= k) for Poisson distribution.
func PoissonCDF(lambda float64, k int) (float64, error) {
	if lambda < 0 || k < 0 {
		return 0, ErrInvalidInput
	}
	sum := 0.0
	for i := 0; i <= k; i++ {
		pmf, _ := PoissonPMF(lambda, i)
		sum += pmf
	}
	return sum, nil
}

// GeometricPMF returns P(X = k) for geometric distribution (trials until first success).
// k >= 1, p = probability of success.
func GeometricPMF(p float64, k int) (float64, error) {
	if p <= 0 || p > 1 || k < 1 {
		return 0, ErrInvalidInput
	}
	return math.Pow(1-p, float64(k-1)) * p, nil
}

// UniformPMF returns the probability of any single outcome in a discrete uniform distribution
// over {1, 2, ..., n}.
func UniformPMF(n int) (float64, error) {
	if n <= 0 {
		return 0, ErrInvalidInput
	}
	return 1.0 / float64(n), nil
}

// NormalPDF returns the probability density function of the normal distribution.
func NormalPDF(x, mu, sigma float64) (float64, error) {
	if sigma <= 0 {
		return 0, ErrInvalidInput
	}
	z := (x - mu) / sigma
	return math.Exp(-0.5*z*z) / (sigma * math.Sqrt(2*Pi)), nil
}

// NormalCDF returns the cumulative distribution function of the normal distribution
// using the error function approximation.
func NormalCDF(x, mu, sigma float64) (float64, error) {
	if sigma <= 0 {
		return 0, ErrInvalidInput
	}
	return 0.5 * (1 + math.Erf((x-mu)/(sigma*math.Sqrt2))), nil
}

// ExponentialPDF returns the PDF of the exponential distribution.
func ExponentialPDF(x, lambda float64) (float64, error) {
	if lambda <= 0 || x < 0 {
		return 0, ErrInvalidInput
	}
	return lambda * math.Exp(-lambda*x), nil
}

// ExponentialCDF returns the CDF of the exponential distribution.
func ExponentialCDF(x, lambda float64) (float64, error) {
	if lambda <= 0 || x < 0 {
		return 0, ErrInvalidInput
	}
	return 1 - math.Exp(-lambda*x), nil
}

// BernoulliTrial returns the probability of a Bernoulli trial: P(X = k) where k is 0 or 1.
func BernoulliTrial(p float64, k int) (float64, error) {
	if p < 0 || p > 1 {
		return 0, ErrInvalidInput
	}
	if k == 1 {
		return p, nil
	}
	if k == 0 {
		return 1 - p, nil
	}
	return 0, ErrInvalidInput
}

// HypergeometricPMF returns the hypergeometric distribution PMF.
// N = population size, K = success states in population, n = draws, k = observed successes.
func HypergeometricPMF(N, K, n, k int) (float64, error) {
	if N <= 0 || K < 0 || n <= 0 || k < 0 || K > N || n > N || k > K || k > n {
		return 0, ErrInvalidInput
	}
	num, _ := CombinationFloat(K, k)
	den1, _ := CombinationFloat(N-K, n-k)
	den2, _ := CombinationFloat(N, n)
	if den2 == 0 {
		return 0, ErrDivByZero
	}
	return num * den1 / den2, nil
}

// FactorialFloat returns n! as float64 for large values.
func FactorialFloat(n int) (float64, error) {
	if n < 0 {
		return 0, ErrNegativeInput
	}
	if n <= 1 {
		return 1, nil
	}
	result := 1.0
	for i := 2; i <= n; i++ {
		result *= float64(i)
	}
	return result, nil
}

// LogFactorial returns ln(n!) using Stirling's approximation for large n.
func LogFactorial(n int) (float64, error) {
	if n < 0 {
		return 0, ErrNegativeInput
	}
	if n <= 1 {
		return 0, nil
	}
	sum := 0.0
	for i := 2; i <= n; i++ {
		sum += math.Log(float64(i))
	}
	return sum, nil
}

// BinomialCoefficient returns C(n, k) as float64 (alias for CombinationFloat).
func BinomialCoefficient(n, k int) (float64, error) {
	return CombinationFloat(n, k)
}
