package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/tpldata.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpldata.gohtml", `Release self-focus; embrace other-focus.`)
	if err != nil {
		log.Fatalln(err)
	}
}
