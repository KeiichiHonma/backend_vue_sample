package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"html/template"
	"net/http"
	"strconv"
)

var Db *gorm.DB

//postsテーブル
type Post struct{
	gorm.Model
	Content string
	Author string
}
type PostRepository interface {
	SelectAll()
	SelectById(id int)
	SelectByIdOrName(id int,name string)
}
func (value Post) SelectById(id int){

}
func InitDB(){
	var err error
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "blog2"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true"
	Db,err = gorm.Open(DBMS,CONNECT)
	if err != nil{
		panic(err.Error())
	}
	fmt.Println("接続OK")
	Db.AutoMigrate(&Post{})
}

func GetPost() []Post{
	posts := []Post{}
	Db.Find(&posts)
	return posts
}
func DeletePost(postId string){
	post := Post{}
	id,_ := strconv.Atoi(postId)
	post.ID = uint(id)
	Db.First(&post)
	Db.Delete(&post)
}
func StorePost(post Post){
	Db.NewRecord(post)
	Db.Create(&post)
	Db.Save(&post)
	return
}

func (value Post) ToString(){
	fmt.Printf("Id=%d,Content=%s,Author=%s\n", value.ID,value.Content,value.Author)
}
func ShowPost(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		//id := r.FormValue("id")
		content := r.FormValue("content")
		author := r.FormValue("author")
		StorePost(Post{Content:content, Author:author})
	}
	posts := GetPost()
	t,_ := template.ParseFiles("6_4_store-db-gorm/tmpl.html")
	err := t.Execute(w, posts)
	if err != nil{
		panic(err)
	}
}

func main() {
	InitDB()
	//var post Post
	//post.Content = "con"
	//post.Author = "au"
	//StorePost(post)
	//var post Posts
	//post.Id = "con"

	//DeletePost("2")
	//posts := GetPost()
	//for _,value := range posts{
		//fmt.Println(key)
		//fmt.Println(value)
		//value.ToString()
	//}
	http.HandleFunc("/",ShowPost)
	//webサーバー起動
	port := ":8080"
	fmt.Println("起動" + port)
	http.ListenAndServe(port, nil)
}
