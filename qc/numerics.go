package qc

// Numerical tolerance / epsilon values used across the package.
//
// These constants exist to keep numerical behavior consistent across modules
// and to avoid scattering hard-coded thresholds throughout the code.
const (
	// Eps is the default tolerance used for near-zero comparisons.
	Eps = 1e-12

	// EpsTiny is a stricter tolerance for checks where we want to detect
	// near-singularity / division by near-zero.
	EpsTiny = 1e-15

	// EpsSingular is used to decide whether a value/pivot is effectively zero.
	// Chosen to be robust against floating point rounding error.
	EpsSingular = 1e-12

	// DefaultDiffStep is the step size used in finite-difference derivatives
	// when callers don't provide one.
	DefaultDiffStep = 1e-8

	// DefaultNewtonStep is an internal step size used by Newton-Raphson when
	// approximating derivatives.
	DefaultNewtonStep = 1e-10
)
