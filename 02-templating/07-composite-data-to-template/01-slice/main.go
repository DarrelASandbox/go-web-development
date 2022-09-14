package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template  // tplRange
var tpl2 *template.Template // tplRangeVariable

func init() {
	tpl = template.Must(template.ParseFiles("templates/tplrange.gohtml"))
	tpl2 = template.Must(template.ParseFiles("templates/tplrangevariable.gohtml"))
}

func main() {
	tplRange()
	tplRangeVariable()
}

func tplRange() {
	sages := []string{"Gandhi", "MLK", "Buddha", "jesus", "Muhammad"}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}

func tplRangeVariable() {
	sages := []string{"Gandhi", "MLK", "Buddha", "jesus", "Muhammad"}

	err := tpl2.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
