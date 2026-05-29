package qc

import "math"

// Sin returns the sine of x (in radians).
func Sin(x float64) float64 {
	return math.Sin(x)
}

// Cos returns the cosine of x (in radians).
func Cos(x float64) float64 {
	return math.Cos(x)
}

// Tan returns the tangent of x (in radians).
func Tan(x float64) float64 {
	return math.Tan(x)
}

// Cot returns the cotangent of x (in radians).
// Returns ErrDivByZero via error if sin(x) == 0.
func Cot(x float64) (float64, error) {
	s := math.Sin(x)
	if s == 0 {
		return 0, ErrDivByZero
	}
	return math.Cos(x) / s, nil
}

// Sec returns the secant of x (in radians).
// Returns ErrDivByZero if cos(x) == 0.
func Sec(x float64) (float64, error) {
	c := math.Cos(x)
	if c == 0 {
		return 0, ErrDivByZero
	}
	return 1.0 / c, nil
}

// Csc returns the cosecant of x (in radians).
// Returns ErrDivByZero if sin(x) == 0.
func Csc(x float64) (float64, error) {
	s := math.Sin(x)
	if s == 0 {
		return 0, ErrDivByZero
	}
	return 1.0 / s, nil
}

// ASin returns the arc sine (inverse sine) in radians.
// Input must be in [-1, 1].
func ASin(x float64) (float64, error) {
	if x < -1 || x > 1 {
		return 0, ErrInvalidInput
	}
	return math.Asin(x), nil
}

// ACos returns the arc cosine (inverse cosine) in radians.
// Input must be in [-1, 1].
func ACos(x float64) (float64, error) {
	if x < -1 || x > 1 {
		return 0, ErrInvalidInput
	}
	return math.Acos(x), nil
}

// ATan returns the arc tangent (inverse tangent) in radians.
func ATan(x float64) float64 {
	return math.Atan(x)
}

// ATan2 returns the arc tangent of y/x, using the signs to determine the quadrant.
func ATan2(y, x float64) float64 {
	return math.Atan2(y, x)
}

// Sinh returns the hyperbolic sine of x.
func Sinh(x float64) float64 {
	return math.Sinh(x)
}

// Cosh returns the hyperbolic cosine of x.
func Cosh(x float64) float64 {
	return math.Cosh(x)
}

// Tanh returns the hyperbolic tangent of x.
func Tanh(x float64) float64 {
	return math.Tanh(x)
}

// ASinh returns the inverse hyperbolic sine of x.
func ASinh(x float64) float64 {
	return math.Asinh(x)
}

// ACosh returns the inverse hyperbolic cosine of x.
// Returns ErrInvalidInput if x < 1.
func ACosh(x float64) (float64, error) {
	if x < 1 {
		return 0, ErrInvalidInput
	}
	return math.Acosh(x), nil
}

// ATanh returns the inverse hyperbolic tangent of x.
// Returns ErrInvalidInput if |x| >= 1.
func ATanh(x float64) (float64, error) {
	if x <= -1 || x >= 1 {
		return 0, ErrInvalidInput
	}
	return math.Atanh(x), nil
}

// Degrees converts radians to degrees.
func Degrees(radians float64) float64 {
	return radians * DegPerRad
}

// Radians converts degrees to radians.
func Radians(degrees float64) float64 {
	return degrees * RadPerDeg
}

// SinD returns the sine of x (given in degrees).
func SinD(degrees float64) float64 {
	return math.Sin(Radians(degrees))
}

// CosD returns the cosine of x (given in degrees).
func CosD(degrees float64) float64 {
	return math.Cos(Radians(degrees))
}

// TanD returns the tangent of x (given in degrees).
func TanD(degrees float64) float64 {
	return math.Tan(Radians(degrees))
}

// WrapAngle wraps an angle in radians to the range [-Pi, Pi].
func WrapAngle(rad float64) float64 {
	rad = math.Mod(rad, 2*Pi)
	if rad > Pi {
		rad -= 2 * Pi
	} else if rad < -Pi {
		rad += 2 * Pi
	}
	return rad
}

// WrapAngleDeg wraps an angle in degrees to the range [-180, 180].
func WrapAngleDeg(deg float64) float64 {
	deg = math.Mod(deg, 360)
	if deg > 180 {
		deg -= 360
	} else if deg < -180 {
		deg += 360
	}
	return deg
}
