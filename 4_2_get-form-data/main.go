package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
	key1 := r.FormValue("key1")
	fmt.Println(key1)
}
func main() {
	http.HandleFunc("/", handler)
	//webサーバー起動
	port := ":8080"
	fmt.Println("起動" + port)
	http.ListenAndServe(port, nil)
}
