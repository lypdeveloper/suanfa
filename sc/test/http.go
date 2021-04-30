package main

import (
	"log"
	"net/http"
)

func main() {
	// 设置 路由
	http.HandleFunc("/", IndexAction)

	// 开启监听
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func IndexAction(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<h1 align="center">来自小韩说课的问候</h1>`))
}