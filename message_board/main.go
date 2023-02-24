package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"

	// "github.com/gorilla/mux"
	"github.com/go-sql-driver/mysql"
)

// var router = mux.NewRouter()
var db *sql.DB

func initDB() {

	var err error
	config := mysql.Config{
		User:                 "homestead",
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

func createFunc(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "留言板界面")
	tmpl, err := template.ParseFiles("views/create.gohtml")

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, nil)

	// storeURL, _ := router.Get("articles.store").URL()
	// fmt.Fprintf(w, tmpl, storeURL)
}

func main() {

	initDB()

	http.HandleFunc("/", indexFunc)
	//留言板界面
	http.HandleFunc("/articles/create", createFunc)

	http.ListenAndServe(":3001", nil)
}
