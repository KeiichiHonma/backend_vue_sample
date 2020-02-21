package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	t,err := template.ParseFiles("5_1_template-hello/tmpl.html")//
	if err != nil{
		panic(err)
	}
	err2 := t.Execute(w, "hello tmpl")
	if err2 != nil{
		panic(err2)
	}
}
func main() {
	http.HandleFunc("/", handler)
	//webサーバー起動
	port := ":8080"
	fmt.Println("起動" + port)
	http.ListenAndServe(port, nil)
}
