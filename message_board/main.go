package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// var router = mux.NewRouter()
var db *sql.DB

func initDB() {

	var err error
	config := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		DBName:               "goblog",
		AllowNativePasswords: true,
	}

	db, err = sql.Open("mysql", config.FormatDSN())

	//设置最大连接数
	db.SetMaxOpenConns(25)
	//设置最大空闲连接数
	db.SetMaxIdleConns(25)
	// 设置每个链接过期时间
	db.SetConnMaxLifetime(5 * time.Minute)

	//尝试连接，失败会报错
	err = db.Ping()

	checkError(err)
}

func checkError(err error) {

}

func indexFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("first demo")
	fmt.Print("first demo 2")
	fmt.Fprint(w, "first demo")
}

// 中间件
func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//1.设置标头
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		//2.继续处理请求
		next.ServeHTTP(w, r)

	})
}

type ArticlesFormData struct {
	Title, Body string
	URL         *url.URL
	Errors      map[string]string
}

var router = mux.NewRouter()

// 添加数据界面
func createFunc(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "留言板界面")

	storeURL, _ := router.Get("articles.store").URL()
	data := ArticlesFormData{
		Title:  "",
		Body:   "",
		URL:    storeURL,
		Errors: nil,
	}

	tmpl, err := template.ParseFiles("views/create2.gohtml")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		panic(err)
	}

	// storeURL, _ := router.Get("articles.store").URL()
	// fmt.Fprintf(w, tmpl, storeURL)
}

// Article 对应定义一条文章数据结构体
type Article struct {
	Title, Body string
	ID          int64
}

// func (a Article) Link() string {
// 	showURL, _ := router.Get("articles.show").URL("id", strconv.FormatInt(a.ID, 10))

// 	return showURL.string()

// }

// 文章列表
func createlistFunc(w http.ResponseWriter, r *http.Request) {

	// 查询数据
	rows, err := db.Query("SELECT * from articles")

	defer rows.Close()
	var articles []Article

	//循环读取结果
	for rows.Next() {
		var article Article
		err := rows.Scan(&article.ID, &article.Title, &article.Body)

		checkError(err)

		articles = append(articles, article)
	}

	tmpl, err := template.ParseFiles("views/list.gohtml")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, articles)
}

// 添加数据
func createdataFunc(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	body := r.PostFormValue("body")

	LastInsertId, _ := saveArticleTODB(title, body)

	if LastInsertId > 0 {
		fmt.Fprintf(w, "插入成功, id为"+strconv.FormatInt(LastInsertId, 10))
	} else {
		fmt.Fprintf(w, "500 服务器内部错误")
	}

}

// 插入数据
func saveArticleTODB(title string, body string) (int64, error) {
	//变量初始化
	var (
		id   int64
		err  error
		rs   sql.Result
		stmt *sql.Stmt
	)

	//获取一个prepare 申明语句
	stmt, err = db.Prepare("INSERT INTO articles(title, body) values(?, ?)")

	//例行错误检查
	if err != nil {
		return 0, err
	}

	//2.在此函数运行结束后关闭此语句， 防止占用sql连接
	defer stmt.Close()

	//3. 执行请求， 传参进入绑定的内容
	rs, err = stmt.Exec(title, body)
	if err != nil {
		return 0, err
	}

	//4.插入成功的话， 会返回自增ID
	if id, err = rs.LastInsertId(); id > 0 {
		return id, nil
	}

	return 0, err

}

// 展示数据
func showFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	article := Article{}
	query := "SELECT * FROM articles where id = ?"
	err := db.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Body)

	tmpl, err := template.ParseFiles("views/show.gohtml")
	err = tmpl.Execute(w, article)
	fmt.Fprintf(w, "文章id"+id)
	checkError(err)
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//1.除首页外，移除所有请求路径的斜杠
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}

		//2. 将请求传递下去
		next.ServeHTTP(w, r)
	})
}

func main() {

	initDB()

	router.HandleFunc("/", indexFunc)
	// //留言板界面
	// router.HandleFunc("/articles/create", createFunc).Methods("GET").Name("articles.create")
	// //添加数据
	// router.HandleFunc("/articles", createdataFunc).Methods("POST").Name("articles.store")
	router.HandleFunc("/articles/create", createFunc).Methods("GET").Name("articles.create")
	router.HandleFunc("/articles", createdataFunc).Methods("POST").Name("articles.store")

	router.HandleFunc("/articles/list", createlistFunc).Methods("GET").Name("articles.list")
	router.HandleFunc("/articles/{id:[0-9]+}", showFunc).Methods("GET").Name("articles.show")

	//中间件
	router.Use(forceHTMLMiddleware)

	http.ListenAndServe(":3001", router)
}
