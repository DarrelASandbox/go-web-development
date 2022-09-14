package main

import "fmt"

var y int

type hotdog int

type person struct {
	fname string
	lname string
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good morning, James."`)
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `Shaken, not stirred."`)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	x := 7
	fmt.Println(x)
	fmt.Printf("%T", x)

	xi := []int{2, 4, 5, 6}
	fmt.Println(xi)

	m := map[string]int{
		"Todd": 45,
		"Job":  44,
	}
	fmt.Println(m)

	var hd hotdog
	hd = 7
	fmt.Println(hd)

	p1 := person{
		"Miss",
		"Moneypenny",
	}
	fmt.Println(p1)
	p1.speak()

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}

	sa1.speak()
	sa1.person.speak()

	fmt.Println("\n\npolymorphism")
	saySomething(p1)
	saySomething(sa1)
}
