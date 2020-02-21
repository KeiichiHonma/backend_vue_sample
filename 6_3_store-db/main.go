package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
	"strconv"
)

var Db *sql.DB

//postsテーブル
type Posts struct{
	Id string
	Content string
	Author string
}
func (value Posts) ToString(){
	fmt.Printf("Id=%s,Content=%s,Author=%s\n", value.Id,value.Content,value.Author)
}
func InitDB(){
	var err error
	Db,err = sql.Open("mysql", "root:@tcp(localhost:3306)/blog1")
	if err != nil{
		panic(err.Error())
	}
	fmt.Println("接続OK")
}

func GetPost() map[string]*Posts{
	rows,err := Db.Query("select * from posts")
	if err != nil{
		panic(err.Error())
	}
	defer rows.Close()
	var posts = map[string]*Posts{}
	for rows.Next(){
		post := Posts{}
		err := rows.Scan(&post.Id,&post.Content,&post.Author)//ここでデータ入れている、なので参照渡し
		if err != nil{
			panic(err.Error())
			return posts
		}
		posts[post.Id] = &post
	}
	return posts
}

func StorePost(post Posts){
	stmt,err := Db.Prepare(fmt.Sprintf("insert into posts (content, author) values(?,?)"))
	if err != nil{
		 panic(err.Error())
	}
	defer stmt.Close()
	_,err = stmt.Exec(post.Content,post.Author)
	if err != nil{
		panic(err.Error())
	}
	return
}

func DeletePost(postId string){
	stmt,err := Db.Prepare(fmt.Sprintf("delete from posts where id = ?"))
	if err != nil{
		panic(err.Error())
	}
	defer  stmt.Close()
	id,_ := strconv.Atoi(postId)
	_,err = stmt.Exec(id)
	if err != nil{
		panic(err.Error())
	}
	return

}

func ShowPOst(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		//id := r.FormValue("id")
		content := r.FormValue("content")
		author := r.FormValue("author")
		StorePost(Posts{Content:content, Author:author})
	}
	posts := GetPost()
	t,_ := template.ParseFiles("6_3_store-db/tmpl.html")
	err := t.Execute(w, posts)
	if err != nil{
		panic(err)
	}
}

func main() {
	InitDB()
	//var post Posts
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
	http.HandleFunc("/",ShowPOst)
	//webサーバー起動
	port := ":8080"
	fmt.Println("起動" + port)
	http.ListenAndServe(port, nil)
}
