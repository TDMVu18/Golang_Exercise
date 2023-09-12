package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Point Point  `json:"point"`
}

type Point struct {
	Math string `json:"math"`
	Phys string `json:"phys"`
	Chem string `json:"chem"`
}

func main() {
	hk1 := Point{Math: "5", Phys: "7", Chem: "8"}
	Vu := Student{Id: "123", Name: "MinhVu", Age: 18, Point: hk1}
	sample_json, err := json.Marshal(Vu) //json.Marshal() return về mảng kiểu byte và lỗi (nếu có)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(string(sample_json))
}
