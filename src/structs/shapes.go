package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Height, Width float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
