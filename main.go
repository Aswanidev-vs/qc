package main

import (
	"fmt"
	"math"

	q "github.com/Aswanidev-vs/quickcalc/qc"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("  QuickCalc Math Library — Full Demo")
	fmt.Println("========================================")

	// --- Constants ---
	fmt.Println("\n--- Constants ---")
	fmt.Printf("Pi       = %.15f\n", q.Pi)
	fmt.Printf("E        = %.15f\n", q.E)
	fmt.Printf("Phi      = %.15f\n", q.Phi)
	fmt.Printf("Sqrt2    = %.15f\n", q.Sqrt2)

	// --- Arithmetic ---
	fmt.Println("\n--- Arithmetic ---")
	a, b := 15, 4
	fmt.Printf("Add(%d, %d)       = %d\n", a, b, q.Add(a, b))
	fmt.Printf("Sub(%d, %d)       = %d\n", a, b, q.Sub(a, b))
	fmt.Printf("Mul(%d, %d)       = %d\n", a, b, q.Mul(a, b))
	div, _ := q.Div(float64(a), float64(b))
	fmt.Printf("Div(15, 4)       = %.4f\n", div)
	quot, rem, _ := q.DivInt(a, b)
	fmt.Printf("DivInt(15, 4)    = quotient=%d, remainder=%d\n", quot, rem)
	fmt.Printf("Pow(2, 10)       = %.0f\n", q.Pow(2, 10))
	fmt.Printf("Abs(-42)         = %d\n", q.Abs(-42))
	sqrt, _ := q.Sqrt(144)
	fmt.Printf("Sqrt(144)        = %.0f\n", sqrt)
	fmt.Printf("Sum([1..10])     = %d\n", q.Sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Printf("Clamp(15, 0, 10) = %d\n", q.Clamp(15, 0, 10))

	// --- Algebra ---
	fmt.Println("\n--- Algebra ---")
	fmt.Printf("GCD(48, 18)      = %d\n", q.GCD(48, 18))
	fmt.Printf("LCM(4, 6)        = %d\n", q.LCM(4, 6))
	x1, x2, _ := q.SolveQuadratic(1, -5, 6)
	fmt.Printf("x²-5x+6=0 roots = %.1f, %.1f\n", x1, x2)
	fib, _ := q.Fibonacci(20)
	fmt.Printf("Fibonacci(20)    = %d\n", fib)
	fact, _ := q.Factorial(10)
	fmt.Printf("10!              = %d\n", fact)
	fmt.Printf("IsEven(42)       = %t\n", q.IsEven(42))
	fmt.Printf("Lerp(0, 100, 0.3)= %.1f\n", q.Lerp(0, 100, 0.3))

	// --- Number Theory ---
	fmt.Println("\n--- Number Theory ---")
	fmt.Printf("IsPrime(97)      = %t\n", q.IsPrime(97))
	fmt.Printf("NextPrime(100)   = %d\n", q.NextPrime(100))
	primes := q.PrimesUpTo(50)
	fmt.Printf("Primes ≤ 50      = %v\n", primes)
	factors, _ := q.PrimeFactors(360)
	fmt.Printf("PrimeFactors(360)= %v\n", factors)
	tot, _ := q.Totient(12)
	fmt.Printf("φ(12)            = %d\n", tot)
	divs, _ := q.Divisors(28)
	fmt.Printf("Divisors(28)     = %v\n", divs)
	fmt.Printf("IsPerfect(28)    = %t\n", q.IsPerfect(28))
	collatz, _ := q.CollatzSequence(7)
	fmt.Printf("Collatz(7)       = %v\n", collatz)
	fmt.Printf("IsPalindrome(121)= %t\n", q.IsPalindrome(121))
	fmt.Printf("DigitSum(12345)  = %d\n", q.DigitSum(12345))

	// --- Trigonometry ---
	fmt.Println("\n--- Trigonometry ---")
	angle := q.Pi / 4 // 45 degrees
	fmt.Printf("sin(π/4)         = %.10f\n", q.Sin(angle))
	fmt.Printf("cos(π/4)         = %.10f\n", q.Cos(angle))
	fmt.Printf("tan(π/4)         = %.10f\n", q.Tan(angle))
	fmt.Printf("Degrees(π)       = %.1f°\n", q.Degrees(q.Pi))
	fmt.Printf("Radians(180)     = %.5f rad\n", q.Radians(180))
	fmt.Printf("sin(45°)         = %.10f\n", q.SinD(45))
	asin, _ := q.ASin(0.5)
	fmt.Printf("asin(0.5)        = %.5f rad (%.1f°)\n", asin, q.Degrees(asin))

	// --- Statistics ---
	fmt.Println("\n--- Statistics ---")
	data := []float64{2, 4, 4, 4, 5, 5, 7, 9}
	mean, _ := q.Mean(data)
	median, _ := q.Median(data)
	mode, _ := q.Mode(data)
	variance, _ := q.Variance(data)
	stddev, _ := q.StdDev(data)
	fmt.Printf("Data             = %v\n", data)
	fmt.Printf("Mean             = %.2f\n", mean)
	fmt.Printf("Median           = %.2f\n", median)
	fmt.Printf("Mode             = %v\n", mode)
	fmt.Printf("Variance         = %.2f\n", variance)
	fmt.Printf("StdDev           = %.2f\n", stddev)
	q1, q2, q3, _ := q.Quartiles(data)
	fmt.Printf("Q1=%.1f, Q2=%.1f, Q3=%.1f\n", q1, q2, q3)
	skew, _ := q.Skewness(data)
	fmt.Printf("Skewness         = %.4f\n", skew)

	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 5, 4, 5}
	corr, _ := q.Correlation(x, y)
	fmt.Printf("Correlation(x,y) = %.4f\n", corr)

	// --- Geometry ---
	fmt.Println("\n--- Geometry ---")
	circ := q.Circle{Radius: 5}
	fmt.Printf("Circle(r=5) area = %.2f\n", circ.Area())
	fmt.Printf("Circle(r=5) circ = %.2f\n", circ.Circumference())

	rect := q.Rectangle{Width: 4, Height: 6}
	fmt.Printf("Rect(4x6) area   = %.2f\n", rect.Area())
	fmt.Printf("Rect(4x6) diag   = %.2f\n", rect.Diagonal())

	tri := q.Triangle{A: 3, B: 4, C: 5}
	triArea, _ := tri.Area()
	fmt.Printf("Triangle(3,4,5)   area = %.2f\n", triArea)
	fmt.Printf("Triangle(3,4,5)   right = %t\n", tri.IsRightAngled())

	sphere := q.Sphere{Radius: 3}
	fmt.Printf("Sphere(r=3) vol  = %.2f\n", sphere.Volume())
	fmt.Printf("Sphere(r=3) area = %.2f\n", sphere.SurfaceArea())

	p1 := q.Point2D{X: 0, Y: 0}
	p2 := q.Point2D{X: 3, Y: 4}
	fmt.Printf("Distance(0,0)-(3,4) = %.1f\n", q.Distance2D(p1, p2))

	// --- Vectors ---
	fmt.Println("\n--- Vectors ---")
	v1 := q.Vec3{X: 1, Y: 2, Z: 3}
	v2 := q.Vec3{X: 4, Y: 5, Z: 6}
	fmt.Printf("v1·v2            = %.1f\n", v1.Dot(v2))
	cross := v1.Cross(v2)
	fmt.Printf("v1×v2            = (%.1f, %.1f, %.1f)\n", cross.X, cross.Y, cross.Z)
	fmt.Printf("|v1|             = %.4f\n", v1.Magnitude())

	v2d := q.Vec2{X: 3, Y: 4}
	fmt.Printf("|(3,4)|          = %.1f\n", v2d.Magnitude())
	fmt.Printf("angle(3,4)       = %.4f rad (%.1f°)\n", v2d.Angle(), q.Degrees(v2d.Angle()))

	// --- Matrix ---
	fmt.Println("\n--- Matrix ---")
	m1, _ := q.NewMatrixFromSlice([][]float64{{1, 2}, {3, 4}})
	m2, _ := q.NewMatrixFromSlice([][]float64{{5, 6}, {7, 8}})
	product, _ := m1.Mul(m2)
	fmt.Printf("[1,2;3,4] × [5,6;7,8] =\n")
	for _, row := range product {
		fmt.Printf("  [%.0f, %.0f]\n", row[0], row[1])
	}
	det, _ := m1.Determinant()
	fmt.Printf("det([1,2;3,4])   = %.0f\n", det)
	fmt.Printf("transpose(m1)    =\n")
	for _, row := range m1.Transpose() {
		fmt.Printf("  [%.0f, %.0f]\n", row[0], row[1])
	}
	inv, _ := m1.Inverse()
	fmt.Printf("inverse(m1)      =\n")
	for _, row := range inv {
		fmt.Printf("  [%.4f, %.4f]\n", row[0], row[1])
	}

	// --- Probability ---
	fmt.Println("\n--- Probability ---")
	perm, _ := q.Permutation(10, 3)
	fmt.Printf("P(10,3)          = %d\n", perm)
	comb, _ := q.Combination(10, 3)
	fmt.Printf("C(10,3)          = %d\n", comb)
	binom, _ := q.BinomialPMF(10, 5, 0.5)
	fmt.Printf("BinomPMF(10,5,.5)= %.4f\n", binom)
	pois, _ := q.PoissonPMF(3, 2)
	fmt.Printf("PoissonPMF(3,2)  = %.4f\n", pois)
	norm, _ := q.NormalPDF(0, 0, 1)
	fmt.Printf("NormalPDF(0,0,1) = %.4f\n", norm)

	// --- Calculus ---
	fmt.Println("\n--- Calculus ---")
	f := func(x float64) float64 { return x * x }
	d := q.Derivative(f, 3, 1e-8)
	fmt.Printf("d/dx(x²) at x=3  = %.6f (expected 6)\n", d)

	area, _ := q.SimpsonIntegrate(func(x float64) float64 { return x * x }, 0, 1, 100)
	fmt.Printf("∫₀¹ x² dx        = %.6f (expected 0.333333)\n", area)

	root, _ := q.NewtonRaphson(func(x float64) float64 { return x*x - 2 }, 1.5, 1e-10, 100)
	fmt.Printf("√2 via Newton     = %.10f\n", root)
	fmt.Printf("math.Sqrt(2)      = %.10f\n", math.Sqrt(2))

	// --- Conversions ---
	fmt.Println("\n--- Conversions ---")
	fmt.Printf("180° → %.5f rad\n", q.DegreesToRadians(180))
	fmt.Printf("π rad → %.1f°\n", q.RadiansToDegrees(q.Pi))
	fmt.Printf("100°C → %.1f°F\n", q.CelsiusToFahrenheit(100))
	fmt.Printf("212°F → %.1f°C\n", q.FahrenheitToCelsius(212))
	fmt.Printf("42 → binary: %s, hex: %s, octal: %s\n",
		q.DecimalToBinary(42), q.DecimalToHex(42), q.DecimalToOctal(42))
	fmt.Printf("10110 (bin) → %d\n", must(q.BinaryToDecimal("10110")))
	fmt.Printf("2A (hex) → %d\n", must(q.HexToDecimal("2A")))

	// Polar/Cartesian
	x_c, y_c := q.PolarToCartesian(5, q.Pi/3)
	fmt.Printf("Polar(5, π/3) → Cartesian(%.4f, %.4f)\n", x_c, y_c)

	// Complex numbers
	z1 := q.Complex{Real: 3, Imag: 4}
	z2 := q.Complex{Real: 1, Imag: 2}
	zSum := q.ComplexAdd(z1, z2)
	zProd := q.ComplexMul(z1, z2)
	fmt.Printf("(3+4i)+(1+2i) = %.0f+%.0fi\n", zSum.Real, zSum.Imag)
	fmt.Printf("(3+4i)*(1+2i) = %.0f+%.0fi\n", zProd.Real, zProd.Imag)
	fmt.Printf("|3+4i|        = %.0f\n", q.ComplexMagnitude(z1))

	// --- Rounding ---
	fmt.Println("\n--- Rounding ---")
	fmt.Printf("Round(3.7)       = %.0f\n", q.Round(3.7))
	fmt.Printf("Floor(3.7)       = %.0f\n", q.Floor(3.7))
	fmt.Printf("Ceil(3.2)        = %.0f\n", q.Ceil(3.2))
	fmt.Printf("RoundTo(π, 4)    = %.4f\n", q.RoundTo(math.Pi, 4))

	fmt.Println("\n========================================")
	fmt.Println("  Demo complete!")
	fmt.Println("========================================")
}

func must(v int, err error) int {
	if err != nil {
		panic(err)
	}
	return v
}
