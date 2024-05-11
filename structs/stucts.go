package structs

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Triangle struct {
	Base  float64
	Left  float64
	Right float64
}

func (t Triangle) Area() float64 {
	semiP := t.Perimeter() / 2
	return math.Sqrt(semiP * (semiP - t.Base) * (semiP - t.Left) * (semiP - t.Right))
}

func (t Triangle) Perimeter() float64 {
	return t.Base + t.Left + t.Right
}
