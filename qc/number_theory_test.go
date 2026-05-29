package qc

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 97, 101}
	for _, p := range primes {
		if !IsPrime(p) {
			t.Errorf("IsPrime(%d) = false, want true", p)
		}
	}
	notPrimes := []int{0, 1, 4, 6, 8, 9, 10, 15, 100}
	for _, n := range notPrimes {
		if IsPrime(n) {
			t.Errorf("IsPrime(%d) = true, want false", n)
		}
	}
}

func TestNextPrevPrime(t *testing.T) {
	if NextPrime(10) != 11 {
		t.Errorf("NextPrime(10) = %d, want 11", NextPrime(10))
	}
	if PrevPrime(10) != 7 {
		t.Errorf("PrevPrime(10) = %d, want 7", PrevPrime(10))
	}
}

func TestPrimesUpTo(t *testing.T) {
	primes := PrimesUpTo(30)
	expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	if len(primes) != len(expected) {
		t.Fatalf("PrimesUpTo(30) len = %d, want %d", len(primes), len(expected))
	}
	for i, p := range primes {
		if p != expected[i] {
			t.Errorf("PrimesUpTo(30)[%d] = %d, want %d", i, p, expected[i])
		}
	}
}

func TestNthPrime(t *testing.T) {
	p, err := NthPrime(1)
	if err != nil || p != 2 {
		t.Errorf("NthPrime(1) = %d, want 2", p)
	}
	p, err = NthPrime(10)
	if err != nil || p != 29 {
		t.Errorf("NthPrime(10) = %d, want 29", p)
	}
}

func TestPrimeFactors(t *testing.T) {
	factors, err := PrimeFactors(360)
	if err != nil {
		t.Fatalf("PrimeFactors(360) error: %v", err)
	}
	// 360 = 2^3 * 3^2 * 5
	expected := []int{2, 2, 2, 3, 3, 5}
	if len(factors) != len(expected) {
		t.Fatalf("PrimeFactors(360) = %v, want %v", factors, expected)
	}
	for i, f := range factors {
		if f != expected[i] {
			t.Errorf("PrimeFactors(360)[%d] = %d, want %d", i, f, expected[i])
		}
	}
}

func TestTotient(t *testing.T) {
	// φ(9) = 6 (1,2,4,5,7,8 are coprime to 9)
	tot, err := Totient(9)
	if err != nil || tot != 6 {
		t.Errorf("Totient(9) = %d, want 6", tot)
	}
	// φ(1) = 1
	tot, err = Totient(1)
	if err != nil || tot != 1 {
		t.Errorf("Totient(1) = %d, want 1", tot)
	}
}

func TestIsCoprime(t *testing.T) {
	if !IsCoprime(7, 13) {
		t.Error("7 and 13 should be coprime")
	}
	if IsCoprime(6, 9) {
		t.Error("6 and 9 should not be coprime")
	}
}

func TestDivisors(t *testing.T) {
	divs, err := Divisors(28)
	if err != nil {
		t.Fatalf("Divisors(28) error: %v", err)
	}
	expected := []int{1, 2, 4, 7, 14, 28}
	if len(divs) != len(expected) {
		t.Fatalf("Divisors(28) = %v, want %v", divs, expected)
	}
	for i, d := range divs {
		if d != expected[i] {
			t.Errorf("Divisors(28)[%d] = %d, want %d", i, d, expected[i])
		}
	}
}

func TestIsPerfect(t *testing.T) {
	if !IsPerfect(6) {
		t.Error("6 should be perfect")
	}
	if !IsPerfect(28) {
		t.Error("28 should be perfect")
	}
	if IsPerfect(12) {
		t.Error("12 should not be perfect")
	}
}

func TestCollatzSequence(t *testing.T) {
	seq, err := CollatzSequence(6)
	if err != nil {
		t.Fatalf("CollatzSequence(6) error: %v", err)
	}
	if seq[0] != 6 || seq[len(seq)-1] != 1 {
		t.Errorf("CollatzSequence(6) = %v, should start with 6 and end with 1", seq)
	}
}

func TestDigitSum(t *testing.T) {
	if DigitSum(12345) != 15 {
		t.Errorf("DigitSum(12345) = %d, want 15", DigitSum(12345))
	}
	if DigitSum(-123) != 6 {
		t.Errorf("DigitSum(-123) = %d, want 6", DigitSum(-123))
	}
}

func TestDigitCount(t *testing.T) {
	if DigitCount(12345) != 5 {
		t.Errorf("DigitCount(12345) = %d, want 5", DigitCount(12345))
	}
	if DigitCount(0) != 1 {
		t.Errorf("DigitCount(0) = %d, want 1", DigitCount(0))
	}
}

func TestReverseInt(t *testing.T) {
	if ReverseInt(12345) != 54321 {
		t.Errorf("ReverseInt(12345) = %d, want 54321", ReverseInt(12345))
	}
	if ReverseInt(-123) != -321 {
		t.Errorf("ReverseInt(-123) = %d, want -321", ReverseInt(-123))
	}
}

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome(121) {
		t.Error("121 should be palindrome")
	}
	if IsPalindrome(123) {
		t.Error("123 should not be palindrome")
	}
}

func TestModularExp(t *testing.T) {
	result, err := ModularExp(2, 10, 1000)
	if err != nil || result != 24 {
		t.Errorf("ModularExp(2, 10, 1000) = %d, want 24", result)
	}
}

func TestIsPowerOfTwo(t *testing.T) {
	if !IsPowerOfTwo(1024) {
		t.Error("1024 should be power of 2")
	}
	if IsPowerOfTwo(100) {
		t.Error("100 should not be power of 2")
	}
	if IsPowerOfTwo(0) {
		t.Error("0 should not be power of 2")
	}
}

func TestIsArmstrong(t *testing.T) {
	if !IsArmstrong(153) {
		t.Error("153 should be Armstrong")
	}
	if !IsArmstrong(370) {
		t.Error("370 should be Armstrong")
	}
	if IsArmstrong(123) {
		t.Error("123 should not be Armstrong")
	}
}

func TestTrailingZeros(t *testing.T) {
	// 10! = 3628800 → 2 trailing zeros
	if TrailingZeros(10) != 2 {
		t.Errorf("TrailingZeros(10) = %d, want 2", TrailingZeros(10))
	}
	// 25! has 6 trailing zeros
	if TrailingZeros(25) != 6 {
		t.Errorf("TrailingZeros(25) = %d, want 6", TrailingZeros(25))
	}
}
