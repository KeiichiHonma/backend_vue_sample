package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request){
	cookie := http.Cookie{
		Name:       "cookie_test_name",
		Value:      "cookie_test_value",
		HttpOnly:   true,//残りpathとかあってカンマがないと怒られる
	}
	w.Header().Set("Set-Cookie", cookie.String())
	fmt.Println("セットしました")
}
func getCookie(w http.ResponseWriter, r *http.Request){
	cookie, err := r.Cookie("cookie_test_name")
	if err != nil{
		fmt.Printf("qwe")
	}
	cookies := r.Cookies()

	fmt.Fprintln(w, cookie.Value)
	fmt.Fprintln(w, cookies)
}
func main() {
	http.HandleFunc("/set", setCookie)
	http.HandleFunc("/get", getCookie)
	//webサーバー起動
	port := ":8080"
	fmt.Println("起動" + port)
	http.ListenAndServe(port, nil)
}
