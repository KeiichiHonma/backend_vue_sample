package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct{
	Id int `json:"id"`
	Content string `json:"content"`
	Author Author `json:"author"`
	Comments []Comment `json:"comments"`
}
type Author struct{
	Id int `json:"id"`
	Name string `json:"name"`
}
type Comment struct{
	Id int `json:"id"`
	Content string `json:"content"`
}
func Decode(filename string) (post Post,err error) {
	jsonFile,err := os.Open(filename)
	if err != nil{
		fmt.Println("jsonファイルエラー")
	}
	defer jsonFile.Close()
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&post)
	if err != nil{
		fmt.Println("jsonでコードエラー")
	}
	return
}
func Unmarshal(filename string) (post Post,err error) {
	jsonFile,err := os.Open(filename)
	if err != nil{
		fmt.Println("jsonファイルエラー")
	}
	defer jsonFile.Close()
	jsonData,err := ioutil.ReadAll(jsonFile)
	if err != nil{
		fmt.Println("readエラー")
	}
	json.Unmarshal(jsonData,&post)
	return
}

func main(){
	_,err := Decode("test.json")
	if err != nil{
		fmt.Println("readエラー")
	}
}