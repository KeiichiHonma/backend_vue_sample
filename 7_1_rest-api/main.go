package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)
var Db *gorm.DB
type PostForm struct{
	gorm.Model
	Content string `"form:content"binding:"required"`
	Author string `"form:author"binding:"required"`
}

func InitDB(){
	var err error
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "blog3"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true"
	Db,err = gorm.Open(DBMS,CONNECT)
	if err != nil{
		panic(err.Error())
	}
	fmt.Println("接続OK")
	Db.AutoMigrate(&PostForm{})
}
func InitRestApi() *gin.Engine{
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET","POST","DELETE"}
	config.AllowHeaders = []string{"Authorization"}
	r.Use(cors.New(config))
	r.GET("/",IndexHandler)
	r.GET("/post",GetPostsHandler)
	r.GET("/post/:id",GetPostHandler)
	r.POST("/post",CreatePostHandler)
	r.PUT("/post/:id",UpdatePostHandler)
	r.DELETE("/post/:id",DeletePostHandler)
	return r
}


type IndexResponse struct{
	Message string `json:"message"`
}
func IndexHandler(c *gin.Context){
	c.JSON(200, IndexResponse{Message:"投稿API"})
}

func GetPostsHandler(c *gin.Context){
	posts := []PostForm{}
	Db.Find(&posts)
	if len(posts) > 0{
		c.JSON(200,posts)
	}else{
		c.JSON(200,posts)
	}
}
func GetPostHandler(c *gin.Context){
	postId := c.Param("id")
	post := PostForm{}
	postIdInt,err := strconv.Atoi(postId)
	if err != nil{
		fmt.Println("IDが不正")
		c.JSON(http.StatusBadRequest,gin.H{"status":"Bad req"})
		return
	}
	post.ID = uint(postIdInt)
	Db.First(&post)
	c.JSON(200,post)
}

func CreatePostHandler(c *gin.Context){
	var form PostForm
	if err := c.Bind(&form);err != nil{
		fmt.Println(err)
		c.JSON(http.StatusBadRequest,gin.H{"status":"Bad req"})
		return
	}
	post := PostForm{Content:form.Content,Author:form.Author}
	Db.NewRecord(post)
	Db.Create(&post)
	Db.Save(&post)
	c.JSON(200,post)
	return
}
func UpdatePostHandler(c *gin.Context){
	var form PostForm
	if err := c.Bind(&form);err != nil{
		fmt.Println(err)
		c.JSON(http.StatusBadRequest,gin.H{"status":"Bad req"})
		return
	}
	postId := c.Param("id")
	post := PostForm{}
	postIdInt,err := strconv.Atoi(postId)
	if err != nil{
		fmt.Println("IDが不正")
		c.JSON(http.StatusBadRequest,gin.H{"status":"Bad req"})
		return
	}
	post.ID = uint(postIdInt)
	Db.First(&post)
	post = PostForm{Content:form.Content,Author:form.Author}
	Db.Save(&post)
	c.JSON(200,gin.H{"message":"更新","post":post})
	return
}
func DeletePostHandler(c *gin.Context){
	postId := c.Param("id")
	post := PostForm{}
	postIdInt,err := strconv.Atoi(postId)
	if err != nil{
		fmt.Println("IDが不正")
		c.JSON(http.StatusBadRequest,gin.H{"status":"Bad req"})
		return
	}
	post.ID = uint(postIdInt)
	Db.First(&post)
	Db.Delete(&post)
	c.JSON(200,gin.H{"message":"削除","post":post})
	return
}
func main() {
	InitDB()
	r := InitRestApi()
	fmt.Println("OK")
	r.Run(":8080")
}