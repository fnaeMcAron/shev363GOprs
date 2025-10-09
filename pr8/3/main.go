package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64 // площадь
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	// стороны
	A float64
	B float64
	C float64
}

// --------------------------------------------
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (t *Triangle) Area() float64 {
	//герон
	p := t.Perimeter() / 2
	return math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))
}

func (t *Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

// --------------------------------------------
func main() {
	rec := Rectangle{Width: 5, Height: 4}
	cir := Circle{Radius: 3}
	tri := Triangle{A: 3, B: 4, C: 5}

	fmt.Println(rec.Perimeter())
	fmt.Println(cir.Perimeter())
	fmt.Println(tri)
	fmt.Println("Le fin. 😞")
}
