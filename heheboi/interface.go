package main

import (
	"fmt"
	"math"
)

type Number float64

func (n Number) Abs() float64 {
	if n < 0 {
		return float64(-n)
	} else {
		return float64(n)
	}
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Abser interface {
	Abs() float64
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var a Abser
	f := Number(-5)
	v := Vertex{2, 3}
	a = f
	a = &v
	fmt.Println(a.Abs())

	text := T{"Hello NNS"}

	text.M()
}
