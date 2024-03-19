package shapes

import "math"

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Length  float64
	Breadth float64
}

type Triangle struct {
  Base float64 
  Height float64
}

func (r Rectangle) Area() float64 {
	return r.Breadth * r.Length
}

func (c Circle) Area() float64 {
	return math.Pow(c.Radius, 2) * math.Pi
}

func (t Triangle) Area() float64 {
  return t.Base * t.Height / 2;  
}
