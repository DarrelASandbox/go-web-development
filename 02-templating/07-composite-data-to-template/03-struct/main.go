package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template
var tplSlice *template.Template
var tplSliceStruct *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("07-composite-data-to-template/03-struct/tpl.gohtml"))
	tplSlice = template.Must(template.ParseFiles("07-composite-data-to-template/03-struct/tplslice.gohtml"))
	tplSliceStruct = template.Must(template.ParseFiles("07-composite-data-to-template/03-struct/tplslicestruct.gohtml"))
}

func main() {
	structOnly()
	structSlice()
	structSliceStruct()
}

func structOnly() {
	fmt.Println("structOnly:")

	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	err := tpl.Execute(os.Stdout, buddha)
	if err != nil {
		log.Fatalln(err)
	}
}

func structSlice() {
	fmt.Println("\n\nstructSlice:")

	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	mlk := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}

	muhammad := sage{
		Name:  "Muhammad",
		Motto: "To overcome evil with good is good, to resist evil by evil is evil.",
	}

	sages := []sage{buddha, gandhi, mlk, jesus, muhammad}

	err := tplSlice.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}

func structSliceStruct() {
	fmt.Println("\n\nstructSliceStruct:")

	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	m := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{b, g, m}
	cars := []car{f, c}

	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tplSliceStruct.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
