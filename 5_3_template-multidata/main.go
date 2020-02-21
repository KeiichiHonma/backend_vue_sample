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
	t,err := template.ParseFiles("5_3_template-multidata/tmpl.html")//
	if err != nil{
		panic(err)
	}
	d := make(map[string]interface{})
	d["a"] = "hoge"
	d["b"] = true
	d["c"] = 10
	d["d"] = User{"honma",39}
	err2 := t.Execute(w, d)
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
