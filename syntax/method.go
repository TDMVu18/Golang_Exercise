package main

import (
	"fmt"
	"math"
)

// Khởi tạo struct Rectangle (dài, rộng)

type Rectangle struct {
	Width  float64
	Height float64
}

type Number float64

// Tạo method Abs cho Number

func (n Number) Abs() float64 {
	if n < 0 {
		return float64(-n)
	} else {
		return float64(n)
	}
}

// Tạo method Square cho Struct Rectangle

func (r Rectangle) Square() float64 {
	return math.Sqrt(r.Width*r.Width + r.Height*r.Height)
}

// Tạo method Scale cho Struct Rectangle

func (r *Rectangle) Scale(f float64) {
	r.Width = r.Width * f
	r.Height = r.Height * f
}

func main() {
	v := Rectangle{Width: 3, Height: 4}
	fmt.Println("Đường chéo của hcn là", v.Square())
	a := Number(-5)
	fmt.Println("Trị tuyệt đối của", a, "là:", a.Abs())
	v.Scale(5)
	fmt.Println("Đường chéo sau khi scale là:", v.Square())
}
