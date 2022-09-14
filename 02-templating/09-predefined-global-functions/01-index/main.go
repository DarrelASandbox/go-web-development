package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("09-predefined-global-functions/01-index/tpl.gohtml"))
}
func main() {
	xs := []string{"zero", "one", "two", "three", "four", "five"}

	err := tpl.Execute(os.Stdout, xs)
	if err != nil {
		log.Fatalln(err)
	}
}
