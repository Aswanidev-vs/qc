# QuickCalc

A comprehensive Go math library providing arithmetic, algebra, trigonometry, statistics, number theory, geometry, linear algebra, probability, calculus, vector operations, and unit conversions.

## Installation

```sh
go get github.com/Aswanidev-vs/quickcalc/qc
```

## Modules

### Constants (`constants.go`)

Mathematical constants: `Pi`, `E`, `Phi`, `Sqrt2`, `SqrtPi`, `Ln2`, `Ln10`, `Log2E`, `Log10E`, `DegPerRad`, `RadPerDeg`.

### Arithmetic (`arithmetic.go`)

Basic operations with error handling for edge cases (division by zero, overflow).

| Function | Description |
|----------|-------------|
| `Add`, `AddFloat` | Addition |
| `Sub`, `SubFloat` | Subtraction |
| `Mul`, `MulFloat` | Multiplication |
| `Div(a, b)` | Float division (returns error on zero) |
| `DivInt(a, b)` | Integer quotient and remainder |
| `Mod`, `ModFloat` | Modulus |
| `Pow`, `PowInt` | Exponentiation |
| `Abs`, `AbsFloat` | Absolute value |
| `Sqrt`, `Cbrt` | Square/cube root |
| `Hypot(p, q)` | Hypotenuse |
| `Sum`, `SumFloat` | Sum of slices |
| `Product`, `ProductFloat` | Product of slices |
| `Max`, `Min`, `MaxFloat`, `MinFloat` | Two-value comparison |
| `MaxOf`, `MinOf` | Slice maximum/minimum |
| `Clamp`, `ClampFloat` | Constrain to range |
| `Sign`, `SignFloat` | Sign function |

### Algebra (`algebra.go`)

Polynomial solving, sequences, and interpolation.

| Function | Description |
|----------|-------------|
| `GCD`, `LCM` | Greatest common divisor, least common multiple |
| `GCDMulti`, `LCMMulti` | GCD/LCM of slices |
| `SolveQuadratic(a, b, c)` | Solve ax^2 + bx + c = 0 |
| `SolveLinear(a, b)` | Solve ax + b = 0 |
| `Fibonacci(n)` | nth Fibonacci number |
| `FibonacciSequence(n)` | First n Fibonacci numbers |
| `Factorial(n)` | n factorial |
| `DoubleFactorial`, `FallingFactorial`, `RisingFactorial` | Extended factorials |
| `IsEven`, `IsOdd` | Parity checks |
| `Lerp`, `InvLerp`, `Remap` | Linear interpolation |
| `ArithmeticMean`, `GeometricMean`, `HarmonicMean` | Mean types |
| `WeightedMean` | Weighted average |

### Trigonometry (`trigonometry.go`)

Trig, inverse trig, and hyperbolic functions. Both radian and degree variants.

| Function | Description |
|----------|-------------|
| `Sin`, `Cos`, `Tan` | Basic trig (radians) |
| `Cot`, `Sec`, `Csc` | Reciprocal trig |
| `ASin`, `ACos`, `ATan`, `ATan2` | Inverse trig |
| `Sinh`, `Cosh`, `Tanh` | Hyperbolic |
| `ASinh`, `ACosh`, `ATanh` | Inverse hyperbolic |
| `Degrees`, `Radians` | Angle conversion |
| `SinD`, `CosD`, `TanD` | Trig in degrees |
| `WrapAngle`, `WrapAngleDeg` | Angle normalization |

### Statistics (`statistics.go`)

Descriptive statistics, distribution analysis, and correlation.

| Function | Description |
|----------|-------------|
| `Mean`, `Median`, `Mode` | Central tendency |
| `Variance`, `SampleVariance` | Spread (population/sample) |
| `StdDev`, `SampleStdDev` | Standard deviation |
| `Range` | Max - min |
| `Percentile(data, p)` | p-th percentile |
| `Quartiles`, `IQR` | Q1, Q2, Q3 and IQR |
| `Covariance`, `Correlation` | Bivariate statistics |
| `ZScore` | Standardized value |
| `MAD` | Mean absolute deviation |
| `RMS` | Root mean square |
| `Entropy` | Shannon entropy |
| `Skewness`, `Kurtosis` | Distribution shape |

### Number Theory (`number_theory.go`)

Primes, factorization, divisibility, and digit operations.

| Function | Description |
|----------|-------------|
| `IsPrime`, `Prime` | Primality test |
| `NextPrime`, `PrevPrime` | Adjacent primes |
| `PrimesUpTo(n)` | Sieve of Eratosthenes |
| `NthPrime(n)` | nth prime |
| `PrimeFactors(n)` | Prime factorization |
| `Totient(n)` | Euler's totient |
| `IsCoprime` | Coprimality |
| `Divisors`, `DivisorCount`, `DivisorSum` | Divisor functions |
| `IsPerfect`, `IsAbundant`, `IsDeficient` | Number classification |
| `CollatzSequence` | Collatz conjecture |
| `DigitSum`, `DigitCount`, `ReverseInt`, `IsPalindrome` | Digit operations |
| `ModularExp`, `ModInverse` | Modular arithmetic |
| `IsArmstrong`, `IsPowerOfTwo` | Special numbers |
| `TrailingZeros` | Trailing zeros in n! |

### Geometry (`geometry.go`)

2D/3D shapes, distances, and coordinate geometry.

**Types:** `Point2D`, `Point3D`, `Rectangle`, `Circle`, `Triangle`, `Sphere`, `Cylinder`, `Cone`

| Function | Description |
|----------|-------------|
| `Distance2D`, `Distance3D` | Euclidean distance |
| `ManhattanDistance2D/3D` | Manhattan distance |
| `Midpoint2D`, `Midpoint3D` | Midpoint |
| `Rectangle.Area/Perimeter/Diagonal/IsSquare` | Rectangle properties |
| `Circle.Area/Circumference/SectorArea/ArcLength` | Circle properties |
| `Triangle.Area/Perimeter/IsValid/IsRightAngled/Angles/Height` | Triangle properties |
| `Sphere.Volume/SurfaceArea` | Sphere properties |
| `Cylinder.Volume/SurfaceArea/LateralArea` | Cylinder properties |
| `Cone.Volume/SurfaceArea/LateralArea` | Cone properties |
| `RegularPolygonArea/Perimeter/InteriorAngle` | Regular polygon |
| `Slope`, `LineEquation`, `PointLineDistance` | Line operations |
| `CircleFromThreePoints` | Circumscribed circle |

### Matrix (`matrix.go`)

2D matrix operations with full linear algebra support.

**Type:** `Matrix` (2D float64 slice)

| Function | Description |
|----------|-------------|
| `NewMatrix`, `NewMatrixFromSlice`, `Identity` | Construction |
| `Add`, `Sub`, `Mul`, `ScalarMul` | Arithmetic |
| `Transpose` | Matrix transpose |
| `Determinant` | Determinant |
| `Inverse` | Matrix inverse (Gauss-Jordan) |
| `Trace` | Sum of diagonal |
| `Rank` | Matrix rank |
| `IsSquare`, `IsDiagonal`, `IsSymmetric` | Properties |

### Vectors (`vectors.go`)

2D and 3D vector algebra.

**Types:** `Vec2`, `Vec3`

| Function | Description |
|----------|-------------|
| `Add`, `Sub`, `Scale` | Arithmetic |
| `Dot`, `Cross` | Products |
| `Magnitude`, `MagnitudeSquared`, `Normalize` | Norms |
| `Angle`, `AngleBetween` | Angles |
| `Rotate` (Vec2) | Rotation |
| `Perpendicular` (Vec2) | Perpendicular vector |
| `Project`, `Reflect` | Projection, reflection |
| `Distance` | Distance between tips |
| `LerpVec2`, `LerpVec3` | Vector interpolation |
| `TripleScalar` (Vec3) | Scalar triple product |

### Probability (`probability.go`)

Combinatorics and probability distributions.

| Function | Description |
|----------|-------------|
| `Permutation`, `Combination` | Counting (int/float) |
| `BinomialPMF`, `BinomialCDF` | Binomial distribution |
| `PoissonPMF`, `PoissonCDF` | Poisson distribution |
| `NormalPDF`, `NormalCDF` | Normal distribution |
| `ExponentialPDF`, `ExponentialCDF` | Exponential distribution |
| `GeometricPMF` | Geometric distribution |
| `BernoulliTrial` | Bernoulli distribution |
| `HypergeometricPMF` | Hypergeometric distribution |
| `FactorialFloat`, `LogFactorial` | Large factorial support |

### Calculus (`calculus.go`)

Numerical differentiation, integration, root-finding, and ODE solvers.

| Function | Description |
|----------|-------------|
| `Derivative(f, x, h)` | First derivative (central difference) |
| `SecondDerivative` | Second derivative |
| `NthDerivative` | nth derivative |
| `SimpsonIntegrate` | Simpson's 1/3 rule |
| `TrapezoidalIntegrate` | Trapezoidal rule |
| `MidpointIntegrate` | Midpoint rule |
| `RombergIntegrate` | Romberg's method |
| `Limit`, `LimitFromLeft/Right` | Numerical limits |
| `NewtonRaphson` | Root-finding |
| `BisectionMethod` | Root-finding |
| `SecantMethod` | Root-finding |
| `TaylorPolynomial` | Taylor series evaluation |
| `EulerMethod` | ODE solver |
| `RungeKutta4` | ODE solver (RK4) |

### Conversions (`conversions.go`)

Unit conversions, number bases, coordinate systems, and complex numbers.

| Category | Functions |
|----------|-----------|
| **Angles** | `DegreesToRadians`, `RadiansToDegrees`, `GradiansToDegrees`, `DegreesToGradians` |
| **Temperature** | `CelsiusToFahrenheit`, `FahrenheitToCelsius`, `CelsiusToKelvin`, `KelvinToCelsius`, etc. |
| **Number Bases** | `DecimalToBinary/Hex/Octal`, `BinaryToDecimal`, `HexToDecimal`, `OctalToDecimal` |
| **Coordinates** | `PolarToCartesian`, `CartesianToPolar`, `SphericalToCartesian`, `CylindricalToCartesian` |
| **Distance** | `KilometersToMiles`, `MetersToFeet`, `InchesToCentimeters` |
| **Mass** | `PoundsToKilograms`, `OuncesToGrams` |
| **Logarithms** | `LogBase`, `Log2`, `Log10`, `Ln`, `Exp`, `Exp2`, `Exp10` |
| **Complex** | `ComplexAdd/Sub/Mul/Div`, `ComplexMagnitude/Conjugate/Phase`, `ComplexFromPolar`, `ComplexExp/Pow/Sqrt` |
| **Special** | `Erf`, `Erfc`, `Gamma`, `LogGamma` |
| **Rounding** | `Round`, `Floor`, `Ceil`, `Trunc`, `RoundTo`, `ToFixed` |

## Example Usage

```go
package main

import (
    "fmt"
    qc "github.com/Aswanidev-vs/quickcalc/qc"
)

func main() {
    // Arithmetic
    sum := qc.Add(10, 20)
    fmt.Println("Sum:", sum)

    // Solve quadratic
    x1, x2, _ := qc.SolveQuadratic(1, -5, 6)
    fmt.Printf("Roots: %.1f, %.1f\n", x1, x2)

    // Statistics
    data := []float64{2, 4, 4, 4, 5, 5, 7, 9}
    mean, _ := qc.Mean(data)
    std, _ := qc.StdDev(data)
    fmt.Printf("Mean=%.2f, StdDev=%.2f\n", mean, std)

    // Geometry
    c := qc.Circle{Radius: 5}
    fmt.Printf("Area: %.2f\n", c.Area())

    // Matrix
    m, _ := qc.NewMatrixFromSlice([][]float64{{1, 2}, {3, 4}})
    det, _ := m.Determinant()
    fmt.Printf("Determinant: %.0f\n", det)

    // Calculus
    f := func(x float64) float64 { return x * x }
    deriv := qc.Derivative(f, 3, 1e-8)
    fmt.Printf("d/dx(x^2) at x=3: %.2f\n", deriv)

    // Number Theory
    fmt.Println("Is 97 prime?", qc.IsPrime(97))
    factors, _ := qc.PrimeFactors(360)
    fmt.Println("Factors of 360:", factors)
}
```

## Testing

```sh
go test ./qc/ -v
```

## Project Structure

```
quickcalc/
├── main.go              # Demo application
├── go.mod
├── README.md
└── qc/                  # Library package
    ├── constants.go     # Mathematical constants
    ├── arithmetic.go    # Basic operations
    ├── algebra.go       # Polynomials, GCD, sequences
    ├── trigonometry.go  # Trig and inverse trig
    ├── statistics.go    # Descriptive statistics
    ├── number_theory.go # Primes, factorization
    ├── geometry.go      # Shapes and coordinate geometry
    ├── matrix.go        # Matrix operations
    ├── vectors.go       # 2D/3D vector algebra
    ├── probability.go   # Distributions and combinatorics
    ├── calculus.go      # Numerical calculus
    ├── conversions.go   # Units, bases, complex numbers
    └── *_test.go        # Comprehensive tests
```
