package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("3_4_web-static/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	//webサーバー起動
	port := ":8080"
	fmt.Println("起動" + port)
	err := http.ListenAndServe(port, mux)
	if err != nil{
		log.Fatal("error", err)
	}
}