package qc

import (
	"math"
	"testing"
)

func TestDistance2D(t *testing.T) {
	p1 := Point2D{0, 0}
	p2 := Point2D{3, 4}
	if !floatEq(Distance2D(p1, p2), 5) {
		t.Errorf("Distance2D = %f, want 5", Distance2D(p1, p2))
	}
}

func TestManhattanDistance2D(t *testing.T) {
	p1 := Point2D{0, 0}
	p2 := Point2D{3, 4}
	if !floatEq(ManhattanDistance2D(p1, p2), 7) {
		t.Errorf("ManhattanDistance2D = %f, want 7", ManhattanDistance2D(p1, p2))
	}
}

func TestDistance3D(t *testing.T) {
	p1 := Point3D{0, 0, 0}
	p2 := Point3D{1, 2, 2}
	if !floatEq(Distance3D(p1, p2), 3) {
		t.Errorf("Distance3D = %f, want 3", Distance3D(p1, p2))
	}
}

func TestMidpoint2D(t *testing.T) {
	m := Midpoint2D(Point2D{0, 0}, Point2D{4, 6})
	if !floatEq(m.X, 2) || !floatEq(m.Y, 3) {
		t.Errorf("Midpoint2D = (%f, %f), want (2, 3)", m.X, m.Y)
	}
}

func TestRectangle(t *testing.T) {
	r := Rectangle{Width: 4, Height: 6}
	if !floatEq(r.Area(), 24) {
		t.Errorf("Rect area = %f, want 24", r.Area())
	}
	if !floatEq(r.Perimeter(), 20) {
		t.Errorf("Rect perimeter = %f, want 20", r.Perimeter())
	}
	if !floatEq(r.Diagonal(), math.Sqrt(52)) {
		t.Errorf("Rect diagonal = %f, want √52", r.Diagonal())
	}
	if r.IsSquare() {
		t.Error("4x6 rectangle should not be square")
	}
	sq := Rectangle{5, 5}
	if !sq.IsSquare() {
		t.Error("5x5 should be square")
	}
}

func TestCircle(t *testing.T) {
	c := Circle{Radius: 5}
	if !floatEq(c.Area(), Pi*25) {
		t.Errorf("Circle area = %f, want %f", c.Area(), Pi*25)
	}
	if !floatEq(c.Circumference(), 2*Pi*5) {
		t.Errorf("Circle circumference = %f, want %f", c.Circumference(), 10*Pi)
	}
	if !floatEq(c.Diameter(), 10) {
		t.Errorf("Circle diameter = %f, want 10", c.Diameter())
	}
	// Sector area for full circle
	if !floatEq(c.SectorArea(2*Pi), c.Area()) {
		t.Error("Full sector area should equal circle area")
	}
}

func TestTriangle(t *testing.T) {
	tri := Triangle{A: 3, B: 4, C: 5}
	area, err := tri.Area()
	if err != nil {
		t.Fatalf("Triangle area error: %v", err)
	}
	if !floatEq(area, 6) {
		t.Errorf("Triangle(3,4,5) area = %f, want 6", area)
	}
	if !tri.IsRightAngled() {
		t.Error("3-4-5 triangle should be right-angled")
	}
	if !tri.IsValid() {
		t.Error("3-4-5 triangle should be valid")
	}

	bad := Triangle{A: 1, B: 2, C: 10}
	if bad.IsValid() {
		t.Error("1-2-10 should not be valid triangle")
	}

	eq := Triangle{A: 5, B: 5, C: 5}
	if !eq.IsEquilateral() {
		t.Error("5-5-5 should be equilateral")
	}
	if !eq.IsIsosceles() {
		t.Error("5-5-5 should be isosceles")
	}
}

func TestTriangleAngles(t *testing.T) {
	tri := Triangle{A: 3, B: 4, C: 5}
	alpha, beta, gamma, err := tri.Angles()
	if err != nil {
		t.Fatalf("Triangle.Angles error: %v", err)
	}
	// Sum of angles should be Pi
	if !floatEq(alpha+beta+gamma, Pi) {
		t.Errorf("Sum of angles = %f, want π", alpha+beta+gamma)
	}
}

func TestSphere(t *testing.T) {
	s := Sphere{Radius: 3}
	expectedVol := (4.0 / 3.0) * Pi * 27
	if !floatEq(s.Volume(), expectedVol) {
		t.Errorf("Sphere volume = %f, want %f", s.Volume(), expectedVol)
	}
	expectedArea := 4 * Pi * 9
	if !floatEq(s.SurfaceArea(), expectedArea) {
		t.Errorf("Sphere surface area = %f, want %f", s.SurfaceArea(), expectedArea)
	}
}

func TestCylinder(t *testing.T) {
	c := Cylinder{Radius: 2, Height: 5}
	if !floatEq(c.Volume(), Pi*4*5) {
		t.Errorf("Cylinder volume = %f, want %f", c.Volume(), Pi*4*5)
	}
}

func TestCone(t *testing.T) {
	c := Cone{Radius: 3, Height: 4}
	expectedVol := Pi * 9 * 4.0 / 3.0
	if !floatEq(c.Volume(), expectedVol) {
		t.Errorf("Cone volume = %f, want %f", c.Volume(), expectedVol)
	}
}

func TestRegularPolygonArea(t *testing.T) {
	// Hexagon with side 2
	area, err := RegularPolygonArea(6, 2)
	if err != nil {
		t.Fatalf("RegularPolygonArea error: %v", err)
	}
	expected := (6.0 * 4.0) / (4.0 * math.Tan(Pi/6))
	if !floatEq(area, expected) {
		t.Errorf("Hexagon area = %f, want %f", area, expected)
	}
}

func TestCircleFromThreePoints(t *testing.T) {
	p1 := Point2D{0, 1}
	p2 := Point2D{1, 0}
	p3 := Point2D{-1, 0}
	center, radius, err := CircleFromThreePoints(p1, p2, p3)
	if err != nil {
		t.Fatalf("CircleFromThreePoints error: %v", err)
	}
	if !floatEq(center.X, 0) || !floatEq(center.Y, 0) {
		t.Errorf("center = (%f, %f), want (0, 0)", center.X, center.Y)
	}
	if !floatEq(radius, 1) {
		t.Errorf("radius = %f, want 1", radius)
	}
}

func TestSlope(t *testing.T) {
	m, err := Slope(Point2D{0, 0}, Point2D{2, 4})
	if err != nil || !floatEq(m, 2) {
		t.Errorf("Slope = %f, want 2, err=%v", m, err)
	}
	_, err = Slope(Point2D{1, 0}, Point2D{1, 5})
	if err == nil {
		t.Error("Vertical line should return error")
	}
}
