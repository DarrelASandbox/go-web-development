package main

import (
	"log"
	"os"
	"text/template"
)

// string-to-html
func main() {
	tpl, err := template.ParseFiles("templates/tpl.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// Do not use the above code in production
// Can have efficiency improvements
