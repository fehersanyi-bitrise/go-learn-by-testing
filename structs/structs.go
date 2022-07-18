package structs

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	area() float64
}

func (r Rectangle) perimeter() float64 {
	return (r.Width + r.Height) * 2
}

func (r Rectangle) area() float64 {
	return r.Width * r.Height
}

func (c Circle) area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (t Triangle) area() float64 {
	return (t.Base * t.Height) / 2
}
