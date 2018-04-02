// http project http.go
package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello axstaxie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		fmt.Println(r.Form)
	}
}

func main() {
	fmt.Println("hello world!")
	//	http.HandleFunc("/", sayHelloName)
	//	http.HandleFunc("/login", login)
	//	err := http.ListenAndServe(":9090", nil)
	//	if err != nil {
	//		log.Fatal("ListenAndServer:", err)
	//	}
	db := dbConn()
	rows, err := db.Query("SELECT * FROM t")
	fmt.Println(rows, err)

	for rows.Next() {
		var a string
		var b string

		err = rows.Scan(&a, &b)

		fmt.Println(a)
		fmt.Println(b)
	}

}

func dbConn() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	return db
}
