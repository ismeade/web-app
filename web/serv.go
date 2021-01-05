package web

import (
	"fmt"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		io.WriteString(w, "<html><head></head><body><h1>Welcome!</h1><img src=\"/static/test.jpg\"></body></html>")
		return
	}
}

func static(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func StartUp() {
	http.Handle("/static/", http.FileServer(http.Dir("static")))
	//http.HandleFunc("/static/", static)
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Listen And Server", err.Error())
	}
}
