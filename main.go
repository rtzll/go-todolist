package main

import (
	"github.com/0xfoo/go-todolist/views"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", view.Hello) // setting router rule
	http.HandleFunc("/auth/login", view.Login)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
