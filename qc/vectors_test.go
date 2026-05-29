package qc

import (
	"math"
	"testing"
)

func TestVec2Operations(t *testing.T) {
	v1 := Vec2{3, 4}
	v2 := Vec2{1, 2}

	add := v1.Add(v2)
	if !floatEq(add.X, 4) || !floatEq(add.Y, 6) {
		t.Errorf("Vec2.Add = %v, want (4, 6)", add)
	}

	sub := v1.Sub(v2)
	if !floatEq(sub.X, 2) || !floatEq(sub.Y, 2) {
		t.Errorf("Vec2.Sub = %v, want (2, 2)", sub)
	}

	if !floatEq(v1.Magnitude(), 5) {
		t.Errorf("|v1| = %f, want 5", v1.Magnitude())
	}

	if !floatEq(v1.Dot(v2), 11) {
		t.Errorf("v1·v2 = %f, want 11", v1.Dot(v2))
	}
}

func TestVec2Normalize(t *testing.T) {
	v := Vec2{3, 4}
	n, err := v.Normalize()
	if err != nil {
		t.Fatalf("Normalize error: %v", err)
	}
	if !floatEq(n.Magnitude(), 1) {
		t.Errorf("|normalized| = %f, want 1", n.Magnitude())
	}

	zero := Vec2{0, 0}
	_, err = zero.Normalize()
	if err == nil {
		t.Error("normalizing zero vector should fail")
	}
}

func TestVec2Rotate(t *testing.T) {
	v := Vec2{1, 0}
	rotated := v.Rotate(Pi / 2)
	if !floatEq(rotated.X, 0) || !floatEq(rotated.Y, 1) {
		t.Errorf("Rotate(π/2) = (%f, %f), want (0, 1)", rotated.X, rotated.Y)
	}
}

func TestVec2Angle(t *testing.T) {
	v := Vec2{1, 1}
	if !floatEq(v.Angle(), Pi/4) {
		t.Errorf("Angle(1,1) = %f, want π/4", v.Angle())
	}
}

func TestVec3Operations(t *testing.T) {
	v1 := Vec3{1, 2, 3}
	v2 := Vec3{4, 5, 6}

	dot := v1.Dot(v2)
	if !floatEq(dot, 32) {
		t.Errorf("v1·v2 = %f, want 32", dot)
	}

	cross := v1.Cross(v2)
	if !floatEq(cross.X, -3) || !floatEq(cross.Y, 6) || !floatEq(cross.Z, -3) {
		t.Errorf("v1×v2 = %v, want (-3, 6, -3)", cross)
	}

	if !floatEq(v1.Magnitude(), math.Sqrt(14)) {
		t.Errorf("|v1| = %f, want √14", v1.Magnitude())
	}
}

func TestVec3Normalize(t *testing.T) {
	v := Vec3{1, 2, 2}
	n, err := v.Normalize()
	if err != nil {
		t.Fatalf("Normalize error: %v", err)
	}
	if !floatEq(n.Magnitude(), 1) {
		t.Errorf("|normalized| = %f, want 1", n.Magnitude())
	}
}

func TestVec2Cross(t *testing.T) {
	v1 := Vec2{1, 0}
	v2 := Vec2{0, 1}
	if !floatEq(v1.Cross(v2), 1) {
		t.Errorf("2D cross = %f, want 1", v1.Cross(v2))
	}
}

func TestVec3TripleScalar(t *testing.T) {
	v := Vec3{1, 0, 0}
	w := Vec3{0, 1, 0}
	u := Vec3{0, 0, 1}
	ts := v.TripleScalar(w, u)
	if !floatEq(ts, 1) {
		t.Errorf("TripleScalar = %f, want 1", ts)
	}
}

func TestVec2Project(t *testing.T) {
	v := Vec2{3, 4}
	onto := Vec2{1, 0}
	proj, err := v.Project(onto)
	if err != nil {
		t.Fatalf("Project error: %v", err)
	}
	if !floatEq(proj.X, 3) || !floatEq(proj.Y, 0) {
		t.Errorf("Project = %v, want (3, 0)", proj)
	}
}

func TestVec2Reflect(t *testing.T) {
	v := Vec2{1, -1}
	normal := Vec2{0, 1}
	reflected := v.Reflect(normal)
	if !floatEq(reflected.X, 1) || !floatEq(reflected.Y, 1) {
		t.Errorf("Reflect = %v, want (1, 1)", reflected)
	}
}

func TestLerpVec2(t *testing.T) {
	a := Vec2{0, 0}
	b := Vec2{10, 10}
	mid := LerpVec2(a, b, 0.5)
	if !floatEq(mid.X, 5) || !floatEq(mid.Y, 5) {
		t.Errorf("LerpVec2(0.5) = %v, want (5, 5)", mid)
	}
}
