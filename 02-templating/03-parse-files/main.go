package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("templates/three.gmao")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("template.ParseFiles:")
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\n\ntpl.ParseFiles:")
	tpl, err = tpl.ParseFiles("templates/two.gmao", "templates/vespa.gmao")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\n\ntpl.ExecuteTemplate:")
	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// err = tpl.ExecuteTemplate(os.Stdout, "one.gmao", nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	fmt.Println("\n\ntpl.Execute:")
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
