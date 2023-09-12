package main

import "fmt"

func main() {
	a := Divide(8, 4)
	fmt.Println(a)
	b := Divide(8, 8)
	fmt.Println(b)
	c := Divide(8, 0)
	fmt.Println(c)
	d := Divide(8, 1)
	fmt.Println(d)
}

func HandlePanic() {
	a := recover() //Tiếp tục chương trình khi có panic

	if a != nil {
		fmt.Println("Recover", a)
	}
}

func Divide(num1, num2 int) int {
	//ngoai le

	defer HandlePanic() //HandlePanic được thêm vào hàng chờ xử lý

	var result int

	if num2 == 0 {
		panic("Cannot divide a number by zero") //Chương trình tạm ngưng, chạy HandlePanic
	} else {
		result = num1 / num2
	}
	return result
}
