package structsmethodsinterfaces

import "math"

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle struct
type Circle struct {
	Radius float64
}

// Triangle struct
type Triangle struct {
	Width  float64
	Height float64
}

// Perimeters func
func Perimeters(rec Rectangle) float64 {
	return 2 * (rec.Width + rec.Height)
}

// Area func
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area func
func (t Triangle) Area() float64 {
	return 0.5 * t.Width * t.Height
}

// Area func
func (c Circle) Area() float64 {
	return math.Pi * float64(c.Radius) * float64(c.Radius)
}
