package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	//HTTPレスポンスの出力
	fmt.Fprintf(w, "hello go web")
}

func main() {
	//URLとハンドラのマッピング
	http.HandleFunc("/", handler)
	//webサーバー起動
	port := ":8080"
	fmt.Println("起動" + port)
	http.ListenAndServe(port,nil)
}
