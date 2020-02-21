package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if r.Header.Get("Content-Type") != "application/json"{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//変数宣言している　後で指定サイズの配列を作るため
	length,err := strconv.Atoi(r.Header.Get("Content-Length"))

	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
	}
	body := make([]byte,length)
	//変数宣言が終わっているので、代入しているだけ
	//_,err = r.Body.Read(body) これでもOK
	length,err = r.Body.Read(body)//ここでの戻り値と元のlengthは基本同じ
	if err != nil && err != io.EOF{
		w.WriteHeader(http.StatusInternalServerError)
	}
	var jsonBody map[string]interface{}
	err = json.Unmarshal(body[:length2], &jsonBody)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Printf("%v\n",jsonBody)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w,jsonBody)

}
func main() {
	http.HandleFunc("/json", handler)
	//webサーバー起動
	port := ":8080"
	fmt.Println("起動" + port)
	http.ListenAndServe(port, nil)
}
