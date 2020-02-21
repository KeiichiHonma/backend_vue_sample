package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request){
	r.ParseMultipartForm(1024)
	fileHeader := r.MultipartForm.File["upload"][0]
	file,err := fileHeader.Open()
	if err != nil {
		panic("file error")
	}
	data,err := ioutil.ReadAll(file)
	if err != nil {
		panic("file error")
	}
	wfile,err := os.Create(fileHeader.Filename)
	if err != nil {
		panic("file error")
	}
	defer wfile.Close()
	wfile.Write(([]byte)(string(data)))
}
func main() {
	http.HandleFunc("/upload", handler)
	//webサーバー起動
	port := ":8080"
	fmt.Println("起動" + port)
	http.ListenAndServe(port, nil)
}
