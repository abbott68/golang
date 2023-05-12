package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 导入MySQL数据库驱动
	"html/template"
	"log"
	"net/http"
)

// 博客文章结构体
type BlogPost struct {
	ID      int
	Title   string
	Content string
}

// 用户结构体
type User struct {
	ID       int
	Username string
	Password string
}

var db *sql.DB // 全局数据库连接对象

func initDB() {
	// 连接数据库
	database, err := sql.Open("mysql", "dev:dev123456@tcp(81.70.196.2:3306)/blog_db")
	if err != nil {
		log.Fatal(err)
	}

	// 设置全局数据库连接对象
	db = database

	// 创建博客文章表
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS blog_posts (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		content TEXT
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// 创建用户表
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(50) NOT NULL,
		password VARCHAR(255) NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}
}

// 博客文章列表
var blogPosts = []BlogPost{
	{Title: "Blog Post 1", Content: "This is the content of Blog Post 1."},
	{Title: "Blog Post 2", Content: "This is the content of Blog Post 2."},
	{Title: "Blog Post 3", Content: "This is the content of Blog Post 3."},
}

// 主页处理函数
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// 使用模板渲染主页
	rows, err := db.Query("SELECT * FROM blog_posts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 构建博客文章列表
	blogPosts := []BlogPost{}
	for rows.Next() {
		post := BlogPost{}
		err := rows.Scan(&post.ID, &post.Title, &post.Content)
		if err != nil {
			log.Fatal(err)
		}
		blogPosts = append(blogPosts, post)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// 使用模板渲染主页
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(w, blogPosts)
	if err != nil {
		log.Fatal(err)
	}
}

// 用户注册处理函数
func registerHandler(w http.ResponseWriter, r *http.Request) {
	// 获取注册表单提交的数据
	username := r.FormValue("username")
	password := r.FormValue("password")

	// 将用户信息插入数据库
	_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	if err != nil {
		log.Fatal(err)
	}

	// 重定向到主页
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	// 初始化数据库连接
	initDB()
	defer db.Close()

	// 设置路由和处理函数
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/register", registerHandler)

	// 启动Web服务器
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
