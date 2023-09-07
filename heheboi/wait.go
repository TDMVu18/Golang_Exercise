package main

import (
	"fmt"
	"sync"
)

func main() {
	array := []int{} //tạo mảng chứa phần tử từ 1 - 20

	//tính 2 tổng phần tử index 1-10 và index 11-20
	chunk1 := 0
	chunk2 := 0

	for i := 0; i < 20; i++ {
		array = append(array, i)
	}

	var wg sync.WaitGroup

	wg.Add(2)

	fmt.Printf("%v\n", array)

	go func() {
		defer wg.Done() //giảm wg đi 1
		for i := 0; i < 10; i++ {
			chunk1 += array[i]
		}
		fmt.Printf("chunk1 = %v\n", chunk1)
	}()

	go func() {
		defer wg.Done() //giảm wg đi 1
		for i := 10; i < 20; i++ {
			chunk2 += array[i]
		}
		fmt.Printf("chunk2 = %v\n", chunk2)
	}()
	wg.Wait() // lock hàm main, tính tổng chỉ thực thi khi 2 goroutine được thực thi (wg giảm về 0)
	fmt.Printf("chunk1 + chunk2 = %v\n", chunk1+chunk2)

}
