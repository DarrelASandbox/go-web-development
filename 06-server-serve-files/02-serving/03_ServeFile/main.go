package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/02-serving/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="02-serving/toby.jpg">	`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "02-serving/toby.jpg")
}
