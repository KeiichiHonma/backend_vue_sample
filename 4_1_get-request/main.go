package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	header := r.Header
	fmt.Fprintln(w, "header")
	fmt.Fprintln(w, header)

	len := r.ContentLength//サイズ確認
	body := make([]byte, len)//サイズ分のbyte配列
	r.Body.Read(body)
	fmt.Fprintln(w, "body")
	fmt.Fprintln(w, body)
}
func main() {
	http.HandleFunc("/", handler)
	//webサーバー起動
	port := ":8080"
	fmt.Println("起動" + port)
	http.ListenAndServe(port, nil)
}
