package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	Width  float64
	Height float64
}

type Number float64

func (n Number) Abs() float64 {
	if n < 0 {
		return float64(-n)
	} else {
		return float64(n)
	}
}

func (r Rectangle) Square() float64 {
	return math.Sqrt(r.Width*r.Width + r.Height*r.Height)
}

func main() {
	v := Rectangle{Width: 3, Height: 4}
	fmt.Println(v.Square())
	a := Number(-2.5464)
	fmt.Println(a.Abs())
}
