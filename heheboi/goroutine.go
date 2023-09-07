package main

import (
	"fmt"
	"time"
)

func Status(id int) {
	fmt.Println("Task %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Println("Task %d done\n", id)
}

func main() {

}
