package qc

import (
	"errors"
	"math"
)

// Point2D represents a 2D point.
type Point2D struct {
	X, Y float64
}

// Point3D represents a 3D point.
type Point3D struct {
	X, Y, Z float64
}

// Rectangle represents a rectangle with width and height.
type Rectangle struct {
	Width, Height float64
}

// Circle represents a circle with a given radius.
type Circle struct {
	Radius float64
}

// Triangle represents a triangle with three sides.
type Triangle struct {
	A, B, C float64
}

// Sphere represents a sphere with a given radius.
type Sphere struct {
	Radius float64
}

// Cylinder represents a cylinder with radius and height.
type Cylinder struct {
	Radius, Height float64
}

// Cone represents a cone with radius and height.
type Cone struct {
	Radius, Height float64
}

// --- 2D Distances ---

// Distance2D returns the Euclidean distance between two 2D points.
func Distance2D(p1, p2 Point2D) float64 {
	return math.Sqrt((p2.X-p1.X)*(p2.X-p1.X) + (p2.Y-p1.Y)*(p2.Y-p1.Y))
}

// ManhattanDistance2D returns the Manhattan distance between two 2D points.
func ManhattanDistance2D(p1, p2 Point2D) float64 {
	return math.Abs(p2.X-p1.X) + math.Abs(p2.Y-p1.Y)
}

// Distance3D returns the Euclidean distance between two 3D points.
func Distance3D(p1, p2 Point3D) float64 {
	return math.Sqrt((p2.X-p1.X)*(p2.X-p1.X) + (p2.Y-p1.Y)*(p2.Y-p1.Y) + (p2.Z-p1.Z)*(p2.Z-p1.Z))
}

// ManhattanDistance3D returns the Manhattan distance between two 3D points.
func ManhattanDistance3D(p1, p2 Point3D) float64 {
	return math.Abs(p2.X-p1.X) + math.Abs(p2.Y-p1.Y) + math.Abs(p2.Z-p1.Z)
}

// Midpoint2D returns the midpoint between two 2D points.
func Midpoint2D(p1, p2 Point2D) Point2D {
	return Point2D{(p1.X + p2.X) / 2, (p1.Y + p2.Y) / 2}
}

// Midpoint3D returns the midpoint between two 3D points.
func Midpoint3D(p1, p2 Point3D) Point3D {
	return Point3D{(p1.X + p2.X) / 2, (p1.Y + p2.Y) / 2, (p1.Z + p2.Z) / 2}
}

// --- Rectangle ---

// Area returns the area of a rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter returns the perimeter of a rectangle.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Diagonal returns the diagonal length of a rectangle.
func (r Rectangle) Diagonal() float64 {
	return math.Sqrt(r.Width*r.Width + r.Height*r.Height)
}

// IsSquare returns true if the rectangle is a square.
func (r Rectangle) IsSquare() bool {
	return r.Width == r.Height
}

// --- Circle ---

// Area returns the area of a circle.
func (c Circle) Area() float64 {
	return Pi * c.Radius * c.Radius
}

// Circumference returns the circumference of a circle.
func (c Circle) Circumference() float64 {
	return 2 * Pi * c.Radius
}

// Diameter returns the diameter of a circle.
func (c Circle) Diameter() float64 {
	return 2 * c.Radius
}

// SectorArea returns the area of a circular sector with the given angle in radians.
func (c Circle) SectorArea(angle float64) float64 {
	return 0.5 * c.Radius * c.Radius * angle
}

// ArcLength returns the arc length for the given angle in radians.
func (c Circle) ArcLength(angle float64) float64 {
	return c.Radius * angle
}

// --- Triangle ---

// Perimeter returns the perimeter of a triangle.
func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

// Area returns the area of a triangle using Heron's formula.
// Returns error if the sides don't form a valid triangle.
func (t Triangle) Area() (float64, error) {
	if !t.IsValid() {
		return 0, ErrInvalidInput
	}
	s := t.Perimeter() / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C)), nil
}

// IsValid returns true if the three sides can form a valid triangle.
func (t Triangle) IsValid() bool {
	return t.A > 0 && t.B > 0 && t.C > 0 &&
		t.A+t.B > t.C && t.A+t.C > t.B && t.B+t.C > t.A
}

// IsEquilateral returns true if all three sides are equal.
func (t Triangle) IsEquilateral() bool {
	return t.A == t.B && t.B == t.C
}

// IsIsosceles returns true if at least two sides are equal.
func (t Triangle) IsIsosceles() bool {
	return t.A == t.B || t.B == t.C || t.A == t.C
}

// IsRightAngled returns true if the triangle is right-angled (Pythagorean check with tolerance).
func (t Triangle) IsRightAngled() bool {
	sides := []float64{t.A, t.B, t.C}
	sortFloats(sides)
	return math.Abs(sides[0]*sides[0]+sides[1]*sides[1]-sides[2]*sides[2]) < 1e-9
}

// Angles returns the three angles (in radians) of the triangle using the law of cosines.
func (t Triangle) Angles() (float64, float64, float64, error) {
	if !t.IsValid() {
		return 0, 0, 0, ErrInvalidInput
	}
	alpha := math.Acos((t.B*t.B + t.C*t.C - t.A*t.A) / (2 * t.B * t.C))
	beta := math.Acos((t.A*t.A + t.C*t.C - t.B*t.B) / (2 * t.A * t.C))
	gamma := Pi - alpha - beta
	return alpha, beta, gamma, nil
}

// Height returns the height of the triangle relative to side a.
func (t Triangle) Height() (float64, error) {
	area, err := t.Area()
	if err != nil {
		return 0, err
	}
	return 2 * area / t.A, nil
}

// TriangleAreaFromCoords returns the area of a triangle defined by three 2D points.
func TriangleAreaFromCoords(p1, p2, p3 Point2D) float64 {
	return math.Abs((p1.X*(p2.Y-p3.Y)+p2.X*(p3.Y-p1.Y)+p3.X*(p1.Y-p2.Y))/2) / 1
}

// --- Sphere ---

// Volume returns the volume of a sphere.
func (s Sphere) Volume() float64 {
	return (4.0 / 3.0) * Pi * math.Pow(s.Radius, 3)
}

// SurfaceArea returns the surface area of a sphere.
func (s Sphere) SurfaceArea() float64 {
	return 4 * Pi * s.Radius * s.Radius
}

// --- Cylinder ---

// Volume returns the volume of a cylinder.
func (c Cylinder) Volume() float64 {
	return Pi * c.Radius * c.Radius * c.Height
}

// SurfaceArea returns the surface area of a cylinder.
func (c Cylinder) SurfaceArea() float64 {
	return 2*Pi*c.Radius*c.Height + 2*Pi*c.Radius*c.Radius
}

// LateralArea returns the lateral surface area of a cylinder.
func (c Cylinder) LateralArea() float64 {
	return 2 * Pi * c.Radius * c.Height
}

// --- Cone ---

// Volume returns the volume of a cone.
func (c Cone) Volume() float64 {
	return (1.0 / 3.0) * Pi * c.Radius * c.Radius * c.Height
}

// SurfaceArea returns the surface area of a cone.
func (c Cone) SurfaceArea() float64 {
	slant := math.Sqrt(c.Radius*c.Radius + c.Height*c.Height)
	return Pi*c.Radius*c.Radius + Pi*c.Radius*slant
}

// LateralArea returns the lateral surface area of a cone.
func (c Cone) LateralArea() float64 {
	slant := math.Sqrt(c.Radius*c.Radius + c.Height*c.Height)
	return Pi * c.Radius * slant
}

// --- Polygon ---

// RegularPolygonArea returns the area of a regular polygon with n sides and side length s.
// Returns ErrInvalidInput if n < 3 or s <= 0.
func RegularPolygonArea(n int, s float64) (float64, error) {
	if n < 3 || s <= 0 {
		return 0, ErrInvalidInput
	}
	return (float64(n) * s * s) / (4 * math.Tan(Pi/float64(n))), nil
}

// RegularPolygonPerimeter returns the perimeter of a regular polygon with n sides and side length s.
func RegularPolygonPerimeter(n int, s float64) (float64, error) {
	if n < 3 || s <= 0 {
		return 0, ErrInvalidInput
	}
	return float64(n) * s, nil
}

// RegularPolygonInteriorAngle returns the interior angle (in radians) of a regular polygon with n sides.
func RegularPolygonInteriorAngle(n int) (float64, error) {
	if n < 3 {
		return 0, ErrInvalidInput
	}
	return float64(n-2) * Pi / float64(n), nil
}

// --- Coordinate Geometry ---

// Slope returns the slope of a line through two points.
// Returns an error if the line is vertical (undefined slope).
func Slope(p1, p2 Point2D) (float64, error) {
	if p1.X == p2.X {
		return 0, errors.New("undefined slope (vertical line)")
	}
	return (p2.Y - p1.Y) / (p2.X - p1.X), nil
}

// LineEquation returns (m, b) for the line y = mx + b through two points.
func LineEquation(p1, p2 Point2D) (m, b float64, err error) {
	m, err = Slope(p1, p2)
	if err != nil {
		return 0, 0, err
	}
	b = p1.Y - m*p1.X
	return m, b, nil
}

// PointLineDistance returns the distance from a point to a line defined by ax + by + c = 0.
func PointLineDistance(p Point2D, a, b, c float64) (float64, error) {
	denom := math.Sqrt(a*a + b*b)
	if denom == 0 {
		return 0, ErrInvalidInput
	}
	return math.Abs(a*p.X+b*p.Y+c) / denom, nil
}

// CircleFromThreePoints returns the center and radius of a circle through three points.
func CircleFromThreePoints(p1, p2, p3 Point2D) (center Point2D, radius float64, err error) {
	d := 2 * (p1.X*(p2.Y-p3.Y) + p2.X*(p3.Y-p1.Y) + p3.X*(p1.Y-p2.Y))
	if math.Abs(d) < 1e-12 {
		return Point2D{}, 0, ErrInvalidInput
	}
	ux := ((p1.X*p1.X+p1.Y*p1.Y)*(p2.Y-p3.Y) + (p2.X*p2.X+p2.Y*p2.Y)*(p3.Y-p1.Y) + (p3.X*p3.X+p3.Y*p3.Y)*(p1.Y-p2.Y)) / d
	uy := ((p1.X*p1.X+p1.Y*p1.Y)*(p3.X-p2.X) + (p2.X*p2.X+p2.Y*p2.Y)*(p1.X-p3.X) + (p3.X*p3.X+p3.Y*p3.Y)*(p2.X-p1.X)) / d
	center = Point2D{ux, uy}
	radius = Distance2D(center, p1)
	return center, radius, nil
}

// sortFloats sorts a slice of float64 in ascending order.
func sortFloats(data []float64) {
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
