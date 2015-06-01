package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type User struct {
	Name string
}

func hello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(getTemplateFile("hello.html"))
	if err != nil {
		log.Fatal("Error while parsing template: ", err)
	}
	user := new(User)
	user.Name = getNameFromUrl(r)
	t.Execute(w, user)
}

func getNameFromUrl(r *http.Request) string {
	path := r.URL.Path
	return
}

func getTemplateFile(filename string) string {
	dir, _ := os.Getwd()
	return filepath.Join(dir, "templates", filename)
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // print request method

	if r.Method == "GET" {
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(time.Now().Unix(), 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, err := template.ParseFiles(
			getTemplateFile("base.html"), getTemplateFile("login.html"))
		if err != nil {
			log.Fatal("Error while parsing template: ", err)
		}
		t.Execute(w, token)
	} else {
		// log in request
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			// check token validity
		} else {
			// give error if no token
		}
		fmt.Println("username length:", len(r.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) // print in server side
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) // respond to client
	}
}

func main() {
	http.HandleFunc("/", hello) // setting router rule
	http.HandleFunc("/auth/login", login)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
