package main

import (
	"log"
	"math"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("08-functions/03-pipeline/tpl.gohtml"))
}

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

var fm = template.FuncMap{
	"fdbl":     double,
	"fsq":      square,
	"fsqrt":    sqRoot,
	"fdateMDY": monthDayYear,
}

func main() {

	data := struct {
		Num  int
		Time time.Time
	}{
		3,
		time.Now(),
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

}
