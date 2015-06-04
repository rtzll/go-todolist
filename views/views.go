package view

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
	"strings"
	"time"
)

type User struct {
	Name string
}

func Hello(w http.ResponseWriter, r *http.Request) {
	log.Println("Calling hello func")
	t, err := template.ParseFiles(getTemplateFile("hello.html"))
	if err != nil {
		log.Fatal("Error while parsing template: ", err)
	}
	user := User{Name: getNameFromUrl(r)}
	t.Execute(w, user)
}

func getNameFromUrl(r *http.Request) string {
	return strings.Replace(r.URL.Path, "/", "", -1)
}

func getTemplateFile(filename string) string {
	dir, _ := os.Getwd()
	return filepath.Join(dir, "templates", filename)
}

func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Calling login func")

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
		log.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) // print in server side
		log.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) // respond to client
	}
}
