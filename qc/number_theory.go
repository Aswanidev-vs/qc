package qc

import (
	"errors"
	"math"
)

// IsPrime checks if n is a prime number.
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// Prime returns n if n is prime, 0 otherwise (legacy compatibility).
func Prime(n int) int {
	if IsPrime(n) {
		return n
	}
	return 0
}

// Fact returns n! (legacy compatibility). Returns 0 for negative input.
func Fact(n int) int {
	if n < 0 {
		return 0
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// NextPrime returns the smallest prime number greater than n.
func NextPrime(n int) int {
	if n < 2 {
		return 2
	}
	candidate := n + 1
	if candidate%2 == 0 && candidate != 2 {
		candidate++
	}
	for !IsPrime(candidate) {
		candidate += 2
	}
	return candidate
}

// PrevPrime returns the largest prime number less than n.
// Returns 0 if no such prime exists (n <= 2).
func PrevPrime(n int) int {
	if n <= 2 {
		return 0
	}
	candidate := n - 1
	if candidate%2 == 0 {
		candidate--
	}
	for candidate >= 2 {
		if IsPrime(candidate) {
			return candidate
		}
		candidate -= 2
	}
	// Check 2
	if n > 2 {
		return 2
	}
	return 0
}

// PrimesUpTo returns all primes less than or equal to n using the Sieve of Eratosthenes.
func PrimesUpTo(n int) []int {
	if n < 2 {
		return []int{}
	}
	sieve := make([]bool, n+1)
	for i := range sieve {
		sieve[i] = true
	}
	sieve[0] = false
	sieve[1] = false
	for i := 2; i*i <= n; i++ {
		if sieve[i] {
			for j := i * i; j <= n; j += i {
				sieve[j] = false
			}
		}
	}
	primes := []int{}
	for i, isP := range sieve {
		if isP {
			primes = append(primes, i)
		}
	}
	return primes
}

// NthPrime returns the nth prime number (1-indexed).
// NthPrime(1) = 2, NthPrime(2) = 3, etc.
// Returns ErrInvalidInput if n < 1.
func NthPrime(n int) (int, error) {
	if n < 1 {
		return 0, ErrInvalidInput
	}
	count := 0
	candidate := 2
	for {
		if IsPrime(candidate) {
			count++
			if count == n {
				return candidate, nil
			}
		}
		candidate++
	}
}

// PrimeFactors returns the prime factorization of n as a slice of factors.
// Returns ErrNegativeInput if n < 2.
func PrimeFactors(n int) ([]int, error) {
	if n < 2 {
		return nil, ErrNegativeInput
	}
	factors := []int{}
	// Check for 2s
	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}
	// Check odd factors
	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	// If n is still > 1, it's a prime factor
	if n > 1 {
		factors = append(factors, n)
	}
	return factors, nil
}

// Totient returns Euler's totient function φ(n), the count of integers <= n that are coprime to n.
// Returns ErrNegativeInput if n < 1.
func Totient(n int) (int, error) {
	if n < 1 {
		return 0, ErrNegativeInput
	}
	if n == 1 {
		return 1, nil
	}
	result := n
	pf, _ := PrimeFactors(n)
	seen := make(map[int]bool)
	for _, p := range pf {
		if !seen[p] {
			result -= result / p
			seen[p] = true
		}
	}
	return result, nil
}

// IsCoprime returns true if a and b share no common factors other than 1.
func IsCoprime(a, b int) bool {
	return GCD(Abs(a), Abs(b)) == 1
}

// Divisors returns all positive divisors of n.
// Returns ErrNegativeInput if n < 1.
func Divisors(n int) ([]int, error) {
	if n < 1 {
		return nil, ErrNegativeInput
	}
	divs := []int{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			divs = append(divs, i)
			if i != n/i {
				divs = append(divs, n/i)
			}
		}
	}
	// Sort the divisors
	sortInts(divs)
	return divs, nil
}

// DivisorCount returns the number of positive divisors of n.
// Returns ErrNegativeInput if n < 1.
func DivisorCount(n int) (int, error) {
	if n < 1 {
		return 0, ErrNegativeInput
	}
	count := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			if i == n/i {
				count++
			} else {
				count += 2
			}
		}
	}
	return count, nil
}

// DivisorSum returns the sum of all positive divisors of n.
// Returns ErrNegativeInput if n < 1.
func DivisorSum(n int) (int, error) {
	divs, err := Divisors(n)
	if err != nil {
		return 0, err
	}
	return Sum(divs), nil
}

// IsPerfect returns true if n equals the sum of its proper divisors.
func IsPerfect(n int) bool {
	if n < 2 {
		return false
	}
	sum, _ := DivisorSum(n)
	return sum-n == n
}

// IsAbundant returns true if the sum of proper divisors exceeds n.
func IsAbundant(n int) bool {
	if n < 2 {
		return false
	}
	sum, _ := DivisorSum(n)
	return sum-n > n
}

// IsDeficient returns true if the sum of proper divisors is less than n.
func IsDeficient(n int) bool {
	if n < 2 {
		return false
	}
	sum, _ := DivisorSum(n)
	return sum-n < n
}

// CollatzSequence returns the Collatz sequence starting from n.
// Returns ErrNegativeInput if n < 1.
func CollatzSequence(n int) ([]int, error) {
	if n < 1 {
		return nil, ErrNegativeInput
	}
	seq := []int{n}
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		seq = append(seq, n)
	}
	return seq, nil
}

// DigitSum returns the sum of digits of n.
func DigitSum(n int) int {
	n = Abs(n)
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// DigitCount returns the number of digits in n.
func DigitCount(n int) int {
	n = Abs(n)
	if n == 0 {
		return 1
	}
	return int(math.Log10(float64(n))) + 1
}

// ReverseInt reverses the digits of n.
func ReverseInt(n int) int {
	sign := 1
	if n < 0 {
		sign = -1
		n = -n
	}
	reversed := 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return sign * reversed
}

// IsPalindrome returns true if n reads the same forwards and backwards.
func IsPalindrome(n int) bool {
	return n == ReverseInt(n)
}

// ModularExp computes (base^exp) % mod efficiently.
// Returns ErrZeroModulus if mod is 0.
func ModularExp(base, exp, mod int) (int, error) {
	if mod == 0 {
		return 0, ErrZeroModulus
	}
	if mod == 1 {
		return 0, nil
	}
	result := 1
	base = base % mod
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		exp /= 2
		base = (base * base) % mod
	}
	return result, nil
}

// ModInverse computes the modular multiplicative inverse of a mod m.
// Returns an error if the inverse doesn't exist (a and m not coprime).
func ModInverse(a, m int) (int, error) {
	if m == 0 {
		return 0, ErrZeroModulus
	}
	g, x, _ := extendedGCD(a, m)
	if g != 1 {
		return 0, errors.New("modular inverse does not exist")
	}
	return ((x % m) + m) % m, nil
}

// extendedGCD returns (gcd, x, y) such that a*x + b*y = gcd.
func extendedGCD(a, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	}
	g, x, y := extendedGCD(b%a, a)
	return g, y - (b/a)*x, x
}

// IsArmstrong returns true if n is an Armstrong (narcissistic) number.
func IsArmstrong(n int) bool {
	if n < 0 {
		return false
	}
	original := n
	digits := DigitCount(n)
	sum := 0
	for n > 0 {
		d := n % 10
		sum += int(math.Pow(float64(d), float64(digits)))
		n /= 10
	}
	return sum == original
}

// IsPowerOfTwo returns true if n is a power of 2.
func IsPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

// TrailingZeros returns the number of trailing zeros in n!.
func TrailingZeros(n int) int {
	if n < 0 {
		return 0
	}
	count := 0
	for n >= 5 {
		n /= 5
		count += n
	}
	return count
}

// sortInts sorts a slice of integers in ascending order.
func sortInts(data []int) {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j] > key {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}
