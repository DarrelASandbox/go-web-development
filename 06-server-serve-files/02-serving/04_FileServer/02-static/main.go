package main

import (
	"log"
	"net/http"
)

func main() {
	// using folder of main.go as root folder to run
	// log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("./02-serving/04_FileServer/02-static"))))
}
