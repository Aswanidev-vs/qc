package qc

import (
	"math"
	"testing"
)

func TestDegreesToRadians(t *testing.T) {
	if !floatEq(DegreesToRadians(180), Pi) {
		t.Errorf("180° = %f rad, want π", DegreesToRadians(180))
	}
	if !floatEq(DegreesToRadians(90), Pi/2) {
		t.Errorf("90° = %f rad, want π/2", DegreesToRadians(90))
	}
}

func TestRadiansToDegrees(t *testing.T) {
	if !floatEq(RadiansToDegrees(Pi), 180) {
		t.Errorf("π rad = %f°, want 180", RadiansToDegrees(Pi))
	}
}

func TestTemperatureConversions(t *testing.T) {
	if !floatEq(CelsiusToFahrenheit(100), 212) {
		t.Errorf("100°C = %f°F, want 212", CelsiusToFahrenheit(100))
	}
	if !floatEq(FahrenheitToCelsius(32), 0) {
		t.Errorf("32°F = %f°C, want 0", FahrenheitToCelsius(32))
	}
	if !floatEq(CelsiusToKelvin(0), 273.15) {
		t.Errorf("0°C = %f K, want 273.15", CelsiusToKelvin(0))
	}
	if !floatEq(KelvinToCelsius(273.15), 0) {
		t.Errorf("273.15 K = %f°C, want 0", KelvinToCelsius(273.15))
	}
}

func TestDecimalToBinary(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{0, "0"}, {1, "1"}, {5, "101"}, {42, "101010"}, {-7, "-111"},
	}
	for _, tt := range tests {
		if got := DecimalToBinary(tt.n); got != tt.want {
			t.Errorf("DecimalToBinary(%d) = %s, want %s", tt.n, got, tt.want)
		}
	}
}

func TestDecimalToHex(t *testing.T) {
	if got := DecimalToHex(255); got != "FF" {
		t.Errorf("DecimalToHex(255) = %s, want FF", got)
	}
	if got := DecimalToHex(42); got != "2A" {
		t.Errorf("DecimalToHex(42) = %s, want 2A", got)
	}
}

func TestDecimalToOctal(t *testing.T) {
	if got := DecimalToOctal(8); got != "10" {
		t.Errorf("DecimalToOctal(8) = %s, want 10", got)
	}
	if got := DecimalToOctal(64); got != "100" {
		t.Errorf("DecimalToOctal(64) = %s, want 100", got)
	}
}

func TestBinaryToDecimal(t *testing.T) {
	n, err := BinaryToDecimal("101010")
	if err != nil || n != 42 {
		t.Errorf("BinaryToDecimal('101010') = %d, want 42", n)
	}
	_, err = BinaryToDecimal("123")
	if err != ErrInvalidInput {
		t.Error("BinaryToDecimal('123') should fail")
	}
}

func TestHexToDecimal(t *testing.T) {
	n, err := HexToDecimal("FF")
	if err != nil || n != 255 {
		t.Errorf("HexToDecimal('FF') = %d, want 255", n)
	}
	n, err = HexToDecimal("2a")
	if err != nil || n != 42 {
		t.Errorf("HexToDecimal('2a') = %d, want 42", n)
	}
}

func TestOctalToDecimal(t *testing.T) {
	n, err := OctalToDecimal("10")
	if err != nil || n != 8 {
		t.Errorf("OctalToDecimal('10') = %d, want 8", n)
	}
}

func TestPolarCartesian(t *testing.T) {
	x, y := PolarToCartesian(1, Pi/2)
	if !floatEq(x, 0) || !floatEq(y, 1) {
		t.Errorf("PolarToCartesian(1, π/2) = (%f, %f), want (0, 1)", x, y)
	}
	r, theta := CartesianToPolar(0, 1)
	if !floatEq(r, 1) || !floatEq(theta, Pi/2) {
		t.Errorf("CartesianToPolar(0, 1) = (%f, %f), want (1, π/2)", r, theta)
	}
}

func TestSphericalCartesian(t *testing.T) {
	x, y, z := SphericalToCartesian(1, Pi/2, 0)
	if !floatEq(x, 1) || !floatEq(y, 0) || !floatEq(z, 0) {
		t.Errorf("SphericalToCartesian(1, π/2, 0) = (%f, %f, %f), want (1, 0, 0)", x, y, z)
	}
}

func TestUnitConversions(t *testing.T) {
	if !floatEq(KilometersToMiles(1), 0.621371) {
		t.Errorf("1 km = %f mi, want 0.621371", KilometersToMiles(1))
	}
	if !floatEq(MilesToKilometers(1), 1.60934) {
		t.Errorf("1 mi = %f km, want 1.60934", MilesToKilometers(1))
	}
	if !floatEq(MetersToFeet(1), 3.28084) {
		t.Errorf("1 m = %f ft, want 3.28084", MetersToFeet(1))
	}
	if !floatEq(PoundsToKilograms(1), 0.453592) {
		t.Errorf("1 lb = %f kg, want 0.453592", PoundsToKilograms(1))
	}
}

func TestLogBase(t *testing.T) {
	result, err := LogBase(8, 2)
	if err != nil || !floatEq(result, 3) {
		t.Errorf("LogBase(8, 2) = %f, want 3", result)
	}
	result, err = LogBase(100, 10)
	if err != nil || !floatEq(result, 2) {
		t.Errorf("LogBase(100, 10) = %f, want 2", result)
	}
}

func TestLogLn(t *testing.T) {
	result, err := Ln(math.E)
	if err != nil || !floatEq(result, 1) {
		t.Errorf("Ln(e) = %f, want 1", result)
	}
	_, err = Ln(-1)
	if err != ErrInvalidInput {
		t.Error("Ln(-1) should return ErrInvalidInput")
	}
}

func TestExp(t *testing.T) {
	if !floatEq(Exp(0), 1) {
		t.Errorf("Exp(0) = %f, want 1", Exp(0))
	}
	if !floatEq(Exp(1), math.E) {
		t.Errorf("Exp(1) = %f, want e", Exp(1))
	}
}

func TestComplexArithmetic(t *testing.T) {
	z1 := Complex{3, 4}
	z2 := Complex{1, -2}

	sum := ComplexAdd(z1, z2)
	if !floatEq(sum.Real, 4) || !floatEq(sum.Imag, 2) {
		t.Errorf("ComplexAdd = %v, want (4, 2)", sum)
	}

	diff := ComplexSub(z1, z2)
	if !floatEq(diff.Real, 2) || !floatEq(diff.Imag, 6) {
		t.Errorf("ComplexSub = %v, want (2, 6)", diff)
	}

	prod := ComplexMul(z1, z2)
	// (3+4i)(1-2i) = 3-6i+4i-8i² = 3-2i+8 = 11-2i
	if !floatEq(prod.Real, 11) || !floatEq(prod.Imag, -2) {
		t.Errorf("ComplexMul = %v, want (11, -2)", prod)
	}

	quot, err := ComplexDiv(z1, z2)
	if err != nil {
		t.Fatalf("ComplexDiv error: %v", err)
	}
	// (3+4i)/(1-2i) = (3+4i)(1+2i)/5 = (3+6i+4i+8i²)/5 = (3+10i-8)/5 = (-5+10i)/5 = -1+2i
	if !floatEq(quot.Real, -1) || !floatEq(quot.Imag, 2) {
		t.Errorf("ComplexDiv = %v, want (-1, 2)", quot)
	}
}

func TestComplexMagnitude(t *testing.T) {
	z := Complex{3, 4}
	if !floatEq(ComplexMagnitude(z), 5) {
		t.Errorf("|3+4i| = %f, want 5", ComplexMagnitude(z))
	}
}

func TestComplexConjugate(t *testing.T) {
	z := Complex{3, 4}
	c := ComplexConjugate(z)
	if !floatEq(c.Real, 3) || !floatEq(c.Imag, -4) {
		t.Errorf("Conjugate(3+4i) = %v, want (3, -4)", c)
	}
}

func TestComplexPow(t *testing.T) {
	z := Complex{0, 1} // i
	z2 := ComplexPow(z, 2)
	if !floatEq(z2.Real, -1) || !floatEq(z2.Imag, 0) {
		t.Errorf("i^2 = %v, want (-1, 0)", z2)
	}
}

func TestRoundTo(t *testing.T) {
	if !floatEq(RoundTo(math.Pi, 2), 3.14) {
		t.Errorf("RoundTo(π, 2) = %f, want 3.14", RoundTo(math.Pi, 2))
	}
	if !floatEq(RoundTo(math.Pi, 4), 3.1416) {
		t.Errorf("RoundTo(π, 4) = %f, want 3.1416", RoundTo(math.Pi, 4))
	}
}

func TestFloorCeilRound(t *testing.T) {
	if Floor(3.7) != 3 {
		t.Error("Floor(3.7) should be 3")
	}
	if Ceil(3.2) != 4 {
		t.Error("Ceil(3.2) should be 4")
	}
	if Round(3.5) != 4 {
		t.Error("Round(3.5) should be 4")
	}
	if Round(3.4) != 3 {
		t.Error("Round(3.4) should be 3")
	}
}

func TestErfGamma(t *testing.T) {
	if !floatEq(Erf(0), 0) {
		t.Errorf("Erf(0) = %f, want 0", Erf(0))
	}
	if !floatEq(Gamma(5), 24) { // Γ(5) = 4! = 24
		t.Errorf("Gamma(5) = %f, want 24", Gamma(5))
	}
	if !floatEq(Gamma(0.5), math.Sqrt(Pi)) {
		t.Errorf("Gamma(0.5) = %f, want √π", Gamma(0.5))
	}
}

func TestGradians(t *testing.T) {
	if !floatEq(GradiansToDegrees(200), 180) {
		t.Errorf("200 grad = %f°, want 180", GradiansToDegrees(200))
	}
	if !floatEq(DegreesToGradians(180), 200) {
		t.Errorf("180° = %f grad, want 200", DegreesToGradians(180))
	}
}
