package qc

import "math"

// Vec2 represents a 2D vector.
type Vec2 struct {
	X, Y float64
}

// Vec3 represents a 3D vector.
type Vec3 struct {
	X, Y, Z float64
}

// --- Vec2 Operations ---

// Add adds two 2D vectors.
func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{v.X + other.X, v.Y + other.Y}
}

// Sub subtracts another 2D vector.
func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{v.X - other.X, v.Y - other.Y}
}

// Scale multiplies the vector by a scalar.
func (v Vec2) Scale(s float64) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}

// Dot returns the dot product of two 2D vectors.
func (v Vec2) Dot(other Vec2) float64 {
	return v.X*other.X + v.Y*other.Y
}

// Cross returns the cross product magnitude (scalar for 2D).
func (v Vec2) Cross(other Vec2) float64 {
	return v.X*other.Y - v.Y*other.X
}

// Magnitude returns the length (magnitude) of the vector.
func (v Vec2) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// MagnitudeSquared returns the squared magnitude (avoids sqrt).
func (v Vec2) MagnitudeSquared() float64 {
	return v.X*v.X + v.Y*v.Y
}

// Normalize returns a unit vector in the same direction.
func (v Vec2) Normalize() (Vec2, error) {
	mag := v.Magnitude()
	if mag == 0 {
		return Vec2{}, ErrInvalidInput
	}
	return Vec2{v.X / mag, v.Y / mag}, nil
}

// Angle returns the angle (in radians) between the vector and the positive x-axis.
func (v Vec2) Angle() float64 {
	return math.Atan2(v.Y, v.X)
}

// AngleBetween returns the angle (in radians) between two 2D vectors.
func (v Vec2) AngleBetween(other Vec2) (float64, error) {
	dot := v.Dot(other)
	mag1 := v.Magnitude()
	mag2 := other.Magnitude()
	if mag1 == 0 || mag2 == 0 {
		return 0, ErrInvalidInput
	}
	cosAngle := dot / (mag1 * mag2)
	// Clamp for floating-point safety
	cosAngle = math.Max(-1, math.Min(1, cosAngle))
	return math.Acos(cosAngle), nil
}

// Rotate rotates the vector by angle (in radians) around the origin.
func (v Vec2) Rotate(angle float64) Vec2 {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return Vec2{
		X: v.X*cos - v.Y*sin,
		Y: v.X*sin + v.Y*cos,
	}
}

// Perpendicular returns a vector perpendicular to v (rotated 90 degrees counter-clockwise).
func (v Vec2) Perpendicular() Vec2 {
	return Vec2{-v.Y, v.X}
}

// Project projects this vector onto another vector.
func (v Vec2) Project(onto Vec2) (Vec2, error) {
	denom := onto.MagnitudeSquared()
	if denom == 0 {
		return Vec2{}, ErrInvalidInput
	}
	scale := v.Dot(onto) / denom
	return onto.Scale(scale), nil
}

// Reflect reflects this vector across a normal vector.
func (v Vec2) Reflect(normal Vec2) Vec2 {
	dot := v.Dot(normal)
	return Vec2{
		X: v.X - 2*dot*normal.X,
		Y: v.Y - 2*dot*normal.Y,
	}
}

// Distance returns the distance between the tips of two vectors.
func (v Vec2) Distance(other Vec2) float64 {
	return v.Sub(other).Magnitude()
}

// LerpVec2 linearly interpolates between two 2D vectors.
func LerpVec2(a, b Vec2, t float64) Vec2 {
	return Vec2{
		X: a.X + t*(b.X-a.X),
		Y: a.Y + t*(b.Y-a.Y),
	}
}

// --- Vec3 Operations ---

// Add adds two 3D vectors.
func (v Vec3) Add(other Vec3) Vec3 {
	return Vec3{v.X + other.X, v.Y + other.Y, v.Z + other.Z}
}

// Sub subtracts another 3D vector.
func (v Vec3) Sub(other Vec3) Vec3 {
	return Vec3{v.X - other.X, v.Y - other.Y, v.Z - other.Z}
}

// Scale multiplies the vector by a scalar.
func (v Vec3) Scale(s float64) Vec3 {
	return Vec3{v.X * s, v.Y * s, v.Z * s}
}

// Dot returns the dot product of two 3D vectors.
func (v Vec3) Dot(other Vec3) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

// Cross returns the cross product of two 3D vectors.
func (v Vec3) Cross(other Vec3) Vec3 {
	return Vec3{
		X: v.Y*other.Z - v.Z*other.Y,
		Y: v.Z*other.X - v.X*other.Z,
		Z: v.X*other.Y - v.Y*other.X,
	}
}

// Magnitude returns the length (magnitude) of the vector.
func (v Vec3) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// MagnitudeSquared returns the squared magnitude (avoids sqrt).
func (v Vec3) MagnitudeSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Normalize returns a unit vector in the same direction.
func (v Vec3) Normalize() (Vec3, error) {
	mag := v.Magnitude()
	if mag == 0 {
		return Vec3{}, ErrInvalidInput
	}
	return Vec3{v.X / mag, v.Y / mag, v.Z / mag}, nil
}

// AngleBetween returns the angle (in radians) between two 3D vectors.
func (v Vec3) AngleBetween(other Vec3) (float64, error) {
	dot := v.Dot(other)
	mag1 := v.Magnitude()
	mag2 := other.Magnitude()
	if mag1 == 0 || mag2 == 0 {
		return 0, ErrInvalidInput
	}
	cosAngle := dot / (mag1 * mag2)
	cosAngle = math.Max(-1, math.Min(1, cosAngle))
	return math.Acos(cosAngle), nil
}

// Project projects this vector onto another vector.
func (v Vec3) Project(onto Vec3) (Vec3, error) {
	denom := onto.MagnitudeSquared()
	if denom == 0 {
		return Vec3{}, ErrInvalidInput
	}
	scale := v.Dot(onto) / denom
	return onto.Scale(scale), nil
}

// Reflect reflects this vector across a normal vector.
func (v Vec3) Reflect(normal Vec3) Vec3 {
	dot := v.Dot(normal)
	return Vec3{
		X: v.X - 2*dot*normal.X,
		Y: v.Y - 2*dot*normal.Y,
		Z: v.Z - 2*dot*normal.Z,
	}
}

// Distance returns the distance between the tips of two vectors.
func (v Vec3) Distance(other Vec3) float64 {
	return v.Sub(other).Magnitude()
}

// LerpVec3 linearly interpolates between two 3D vectors.
func LerpVec3(a, b Vec3, t float64) Vec3 {
	return Vec3{
		X: a.X + t*(b.X-a.X),
		Y: a.Y + t*(b.Y-a.Y),
		Z: a.Z + t*(b.Z-a.Z),
	}
}

// TripleScalar returns the scalar triple product: v · (w × u).
func (v Vec3) TripleScalar(w, u Vec3) float64 {
	return v.Dot(w.Cross(u))
}
