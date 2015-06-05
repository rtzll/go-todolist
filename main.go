package main

import (
	"github.com/0xfoo/go-todolist/views"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", view.Hello)
	http.HandleFunc("/auth/login", view.Login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
