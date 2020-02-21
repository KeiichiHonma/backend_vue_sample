package main

import (
	"fmt"
	"html/template"
	"net/http"
)
type User struct{
	Name string
	Age int
}
func handler(w http.ResponseWriter, r *http.Request){
	t,err := template.ParseFiles("5_5_template-loop-map/tmpl.html")//
	if err != nil{
		panic(err)
	}
	data := map[string]string{"a":"hoge","b":"fuga"}
	err2 := t.Execute(w, data)
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
