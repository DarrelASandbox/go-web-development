package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/02-serving/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/02-serving/toby.jpg">`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("02-serving/toby.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}
