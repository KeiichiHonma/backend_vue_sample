package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func handler(w http.ResponseWriter, r *http.Request){
	//HTTPレスポンスの出力
	fmt.Fprintf(w, "hello go ssl web")
}

func main() {
	//証明書
	certFile,_ := filepath.Abs("server.crt")
	keyFile,_ := filepath.Abs("server.key")
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	//webサーバー起動
	port := ":3000"
	fmt.Println("起動" + port)
	err := http.ListenAndServeTLS(port, certFile, keyFile, mux)
	if err != nil{
		log.Fatal("error", err)
	}
}