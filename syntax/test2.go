package main

import (
	"fmt"
	"net/http"
)

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Trang chủ"))
}

func aboutHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Giới thiệu"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homePage) //đăng ký router "/" với hàm xử lý homeHandle
	//http.HandleFunc("/about", aboutHandle) //đăng ký router "/about" với hàm xử lý aboutHandle
	fmt.Println("Server listenning on port 3000 ...")
	fmt.Println(http.ListenAndServe(":3000", mux)) //tạo server, hàm ListenAnhServe nhận yêu cầu
}
