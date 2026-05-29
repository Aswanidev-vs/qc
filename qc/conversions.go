package qc

import "math"

// --- Angle Conversions ---

// DegreesToRadians converts degrees to radians.
func DegreesToRadians(deg float64) float64 {
	return deg * RadPerDeg
}

// RadiansToDegrees converts radians to degrees.
func RadiansToDegrees(rad float64) float64 {
	return rad * DegPerRad
}

// GradiansToDegrees converts gradians to degrees (400 gradians = 360 degrees).
func GradiansToDegrees(grad float64) float64 {
	return grad * 0.9
}

// DegreesToGradians converts degrees to gradians.
func DegreesToGradians(deg float64) float64 {
	return deg / 0.9
}

// --- Temperature Conversions ---

// CelsiusToFahrenheit converts Celsius to Fahrenheit.
func CelsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

// FahrenheitToCelsius converts Fahrenheit to Celsius.
func FahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

// CelsiusToKelvin converts Celsius to Kelvin.
func CelsiusToKelvin(c float64) float64 {
	return c + 273.15
}

// KelvinToCelsius converts Kelvin to Celsius.
func KelvinToCelsius(k float64) float64 {
	return k - 273.15
}

// FahrenheitToKelvin converts Fahrenheit to Kelvin.
func FahrenheitToKelvin(f float64) float64 {
	return CelsiusToKelvin(FahrenheitToCelsius(f))
}

// KelvinToFahrenheit converts Kelvin to Fahrenheit.
func KelvinToFahrenheit(k float64) float64 {
	return CelsiusToFahrenheit(KelvinToCelsius(k))
}

// --- Number Base Conversions ---

// DecimalToBinary converts a decimal integer to its binary representation.
func DecimalToBinary(n int) string {
	if n == 0 {
		return "0"
	}
	negative := false
	if n < 0 {
		negative = true
		n = -n
	}
	binary := ""
	for n > 0 {
		binary = string(rune('0'+n%2)) + binary
		n /= 2
	}
	if negative {
		binary = "-" + binary
	}
	return binary
}

// DecimalToHex converts a decimal integer to its hexadecimal representation.
func DecimalToHex(n int) string {
	if n == 0 {
		return "0"
	}
	negative := false
	if n < 0 {
		negative = true
		n = -n
	}
	hexChars := "0123456789ABCDEF"
	hex := ""
	for n > 0 {
		hex = string(hexChars[n%16]) + hex
		n /= 16
	}
	if negative {
		hex = "-" + hex
	}
	return hex
}

// DecimalToOctal converts a decimal integer to its octal representation.
func DecimalToOctal(n int) string {
	if n == 0 {
		return "0"
	}
	negative := false
	if n < 0 {
		negative = true
		n = -n
	}
	octal := ""
	for n > 0 {
		octal = string(rune('0'+n%8)) + octal
		n /= 8
	}
	if negative {
		octal = "-" + octal
	}
	return octal
}

// BinaryToDecimal converts a binary string to decimal.
// Returns ErrInvalidInput if the string contains non-binary characters.
func BinaryToDecimal(s string) (int, error) {
	if len(s) == 0 {
		return 0, ErrInvalidInput
	}
	negative := false
	if s[0] == '-' {
		negative = true
		s = s[1:]
	}
	result := 0
	for _, c := range s {
		if c != '0' && c != '1' {
			return 0, ErrInvalidInput
		}
		result = result*2 + int(c-'0')
	}
	if negative {
		result = -result
	}
	return result, nil
}

// HexToDecimal converts a hexadecimal string to decimal.
// Returns ErrInvalidInput for invalid hex characters.
func HexToDecimal(s string) (int, error) {
	if len(s) == 0 {
		return 0, ErrInvalidInput
	}
	negative := false
	if s[0] == '-' {
		negative = true
		s = s[1:]
	}
	result := 0
	for _, c := range s {
		result *= 16
		if c >= '0' && c <= '9' {
			result += int(c - '0')
		} else if c >= 'A' && c <= 'F' {
			result += int(c-'A') + 10
		} else if c >= 'a' && c <= 'f' {
			result += int(c-'a') + 10
		} else {
			return 0, ErrInvalidInput
		}
	}
	if negative {
		result = -result
	}
	return result, nil
}

// OctalToDecimal converts an octal string to decimal.
// Returns ErrInvalidInput for invalid octal characters.
func OctalToDecimal(s string) (int, error) {
	if len(s) == 0 {
		return 0, ErrInvalidInput
	}
	negative := false
	if s[0] == '-' {
		negative = true
		s = s[1:]
	}
	result := 0
	for _, c := range s {
		if c < '0' || c > '7' {
			return 0, ErrInvalidInput
		}
		result = result*8 + int(c-'0')
	}
	if negative {
		result = -result
	}
	return result, nil
}

// --- Coordinate Conversions ---

// PolarToCartesian converts polar coordinates (r, theta) to Cartesian (x, y).
func PolarToCartesian(r, theta float64) (x, y float64) {
	x = r * math.Cos(theta)
	y = r * math.Sin(theta)
	return
}

// CartesianToPolar converts Cartesian (x, y) to polar (r, theta).
func CartesianToPolar(x, y float64) (r, theta float64) {
	r = math.Sqrt(x*x + y*y)
	theta = math.Atan2(y, x)
	return
}

// SphericalToCartesian converts spherical (r, theta, phi) to Cartesian (x, y, z).
// theta = polar angle from z-axis, phi = azimuthal angle from x-axis.
func SphericalToCartesian(r, theta, phi float64) (x, y, z float64) {
	x = r * math.Sin(theta) * math.Cos(phi)
	y = r * math.Sin(theta) * math.Sin(phi)
	z = r * math.Cos(theta)
	return
}

// CartesianToSpherical converts Cartesian (x, y, z) to spherical (r, theta, phi).
func CartesianToSpherical(x, y, z float64) (r, theta, phi float64) {
	r = math.Sqrt(x*x + y*y + z*z)
	if r == 0 {
		return 0, 0, 0
	}
	theta = math.Acos(z / r)
	phi = math.Atan2(y, x)
	return
}

// CylindricalToCartesian converts cylindrical (r, theta, z) to Cartesian (x, y, z).
func CylindricalToCartesian(r, theta, z float64) (x, y, zOut float64) {
	x = r * math.Cos(theta)
	y = r * math.Sin(theta)
	zOut = z
	return
}

// CartesianToCylindrical converts Cartesian (x, y, z) to cylindrical (r, theta, z).
func CartesianToCylindrical(x, y, z float64) (r, theta, zOut float64) {
	r = math.Sqrt(x*x + y*y)
	theta = math.Atan2(y, x)
	zOut = z
	return
}

// --- Unit Conversions ---

// KilometersToMiles converts kilometers to miles.
func KilometersToMiles(km float64) float64 {
	return km * 0.621371
}

// MilesToKilometers converts miles to kilometers.
func MilesToKilometers(mi float64) float64 {
	return mi * 1.60934
}

// MetersToFeet converts meters to feet.
func MetersToFeet(m float64) float64 {
	return m * 3.28084
}

// FeetToMeters converts feet to meters.
func FeetToMeters(ft float64) float64 {
	return ft / 3.28084
}

// InchesToCentimeters converts inches to centimeters.
func InchesToCentimeters(in float64) float64 {
	return in * 2.54
}

// CentimetersToInches converts centimeters to inches.
func CentimetersToInches(cm float64) float64 {
	return cm / 2.54
}

// PoundsToKilograms converts pounds to kilograms.
func PoundsToKilograms(lb float64) float64 {
	return lb * 0.453592
}

// KilogramsToPounds converts kilograms to pounds.
func KilogramsToPounds(kg float64) float64 {
	return kg * 2.20462
}

// OuncesToGrams converts ounces to grams.
func OuncesToGrams(oz float64) float64 {
	return oz * 28.3495
}

// GramsToOunces converts grams to ounces.
func GramsToOunces(g float64) float64 {
	return g / 28.3495
}

// --- Logarithm Conversions ---

// LogBase returns log_base(x).
// Returns ErrInvalidInput if x <= 0 or base <= 0 or base == 1.
func LogBase(x, base float64) (float64, error) {
	if x <= 0 || base <= 0 || base == 1 {
		return 0, ErrInvalidInput
	}
	return math.Log(x) / math.Log(base), nil
}

// Log2 returns log base 2 of x.
// Returns ErrInvalidInput if x <= 0.
func Log2(x float64) (float64, error) {
	if x <= 0 {
		return 0, ErrInvalidInput
	}
	return math.Log2(x), nil
}

// Log10 returns log base 10 of x.
// Returns ErrInvalidInput if x <= 0.
func Log10(x float64) (float64, error) {
	if x <= 0 {
		return 0, ErrInvalidInput
	}
	return math.Log10(x), nil
}

// Ln returns the natural logarithm of x.
// Returns ErrInvalidInput if x <= 0.
func Ln(x float64) (float64, error) {
	if x <= 0 {
		return 0, ErrInvalidInput
	}
	return math.Log(x), nil
}

// Exp returns e^x.
func Exp(x float64) float64 {
	return math.Exp(x)
}

// Exp2 returns 2^x.
func Exp2(x float64) float64 {
	return math.Exp2(x)
}

// Exp10 returns 10^x.
func Exp10(x float64) float64 {
	return math.Pow(10, x)
}

// --- Complex Number Support ---

// Complex represents a complex number with real and imaginary parts.
type Complex struct {
	Real, Imag float64
}

// ComplexAdd adds two complex numbers.
func ComplexAdd(a, b Complex) Complex {
	return Complex{a.Real + b.Real, a.Imag + b.Imag}
}

// ComplexSub subtracts two complex numbers.
func ComplexSub(a, b Complex) Complex {
	return Complex{a.Real - b.Real, a.Imag - b.Imag}
}

// ComplexMul multiplies two complex numbers.
func ComplexMul(a, b Complex) Complex {
	return Complex{
		Real: a.Real*b.Real - a.Imag*b.Imag,
		Imag: a.Real*b.Imag + a.Imag*b.Real,
	}
}

// ComplexDiv divides two complex numbers.
// Returns ErrDivByZero if b is zero.
func ComplexDiv(a, b Complex) (Complex, error) {
	denom := b.Real*b.Real + b.Imag*b.Imag
	if denom == 0 {
		return Complex{}, ErrDivByZero
	}
	return Complex{
		Real: (a.Real*b.Real + a.Imag*b.Imag) / denom,
		Imag: (a.Imag*b.Real - a.Real*b.Imag) / denom,
	}, nil
}

// ComplexMagnitude returns |z| = sqrt(real^2 + imag^2).
func ComplexMagnitude(z Complex) float64 {
	return math.Sqrt(z.Real*z.Real + z.Imag*z.Imag)
}

// ComplexConjugate returns the conjugate of z.
func ComplexConjugate(z Complex) Complex {
	return Complex{z.Real, -z.Imag}
}

// ComplexPhase returns the phase (argument) of z in radians.
func ComplexPhase(z Complex) float64 {
	return math.Atan2(z.Imag, z.Real)
}

// ComplexFromPolar creates a complex number from polar form (r, theta).
func ComplexFromPolar(r, theta float64) Complex {
	return Complex{
		Real: r * math.Cos(theta),
		Imag: r * math.Sin(theta),
	}
}

// ComplexExp returns e^z.
func ComplexExp(z Complex) Complex {
	r := math.Exp(z.Real)
	return Complex{
		Real: r * math.Cos(z.Imag),
		Imag: r * math.Sin(z.Imag),
	}
}

// ComplexPow returns z^n for integer n.
func ComplexPow(z Complex, n int) Complex {
	if n == 0 {
		return Complex{1, 0}
	}
	if n < 0 {
		inv, _ := ComplexDiv(Complex{1, 0}, z)
		return ComplexPow(inv, -n)
	}
	result := Complex{1, 0}
	for i := 0; i < n; i++ {
		result = ComplexMul(result, z)
	}
	return result
}

// ComplexSqrt returns the principal square root of z.
func ComplexSqrt(z Complex) Complex {
	mag := ComplexMagnitude(z)
	r := math.Sqrt((mag + z.Real) / 2)
	imagPart := math.Sqrt((mag - z.Real) / 2)
	if math.Signbit(z.Imag) {
		imagPart = -imagPart
	}
	return Complex{r, imagPart}
}

// --- Error Function Approximation ---

// Erf returns the error function approximation of x.
func Erf(x float64) float64 {
	return math.Erf(x)
}

// Erfc returns the complementary error function (1 - erf(x)).
func Erfc(x float64) float64 {
	return math.Erfc(x)
}

// Gamma returns the gamma function Γ(x).
func Gamma(x float64) float64 {
	return math.Gamma(x)
}

// LogGamma returns ln(|Γ(x)|).
func LogGamma(x float64) (float64, int) {
	return math.Lgamma(x)
}

// --- Misc ---

// IsNaN reports whether f is an IEEE 754 "not-a-number" value.
func IsNaN(f float64) bool {
	return math.IsNaN(f)
}

// IsInf reports whether f is positive or negative infinity.
func IsInf(f float64) bool {
	return math.IsInf(f, 0)
}

// IsFinite reports whether f is a finite number (not NaN or Inf).
func IsFinite(f float64) bool {
	return !math.IsNaN(f) && !math.IsInf(f, 0)
}

// Copysign returns a value with the magnitude of x and the sign of y.
func Copysign(x, y float64) float64 {
	return math.Copysign(x, y)
}

// Round rounds x to the nearest integer.
func Round(x float64) float64 {
	return math.Round(x)
}

// Floor returns the greatest integer less than or equal to x.
func Floor(x float64) float64 {
	return math.Floor(x)
}

// Ceil returns the smallest integer greater than or equal to x.
func Ceil(x float64) float64 {
	return math.Ceil(x)
}

// Trunc returns the integer part of x.
func Trunc(x float64) float64 {
	return math.Trunc(x)
}

// RoundTo rounds x to the given number of decimal places.
func RoundTo(x float64, places int) float64 {
	factor := math.Pow(10, float64(places))
	return math.Round(x*factor) / factor
}

// ToFixed is an alias for RoundTo.
func ToFixed(x float64, places int) float64 {
	return RoundTo(x, places)
}

// SafeDiv performs division with a zero check, returns (result, true) or (0, false).
func SafeDiv(a, b float64) (float64, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

// SafeDivInt performs integer division with a zero check.
func SafeDivInt(a, b int) (int, int, bool) {
	if b == 0 {
		return 0, 0, false
	}
	return a / b, a % b, true
}

// DegToRad is an alias for DegreesToRadians.
func DegToRad(deg float64) float64 {
	return DegreesToRadians(deg)
}

// RadToDeg is an alias for RadiansToDegrees.
func RadToDeg(rad float64) float64 {
	return RadiansToDegrees(rad)
}
