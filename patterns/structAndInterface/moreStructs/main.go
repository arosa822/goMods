package main

import "fmt"

type person struct {
	name      string
	age, rank int
}

type admin struct {
	person
	level int
}

type desc interface {
	desc()
}

func (p person) desc() {
	fmt.Printf("Name: %s Age: %d\n", p.name, p.age)
	fmt.Printf("Rank: %d\n", p.rank)
}

func (a admin) desc() {
	fmt.Printf("this is an admin")
	fmt.Printf("\nAge: %d\n", a.age)
}

// constructor
//func (p *person) create(name string age int) {
//	p.name = name
//	p.age = age
//
//}

//func (a *admin) create(name string age int) {
//	a.level = 1

//}
func NewPerson(name string, age int) *person {
	var p person
	p.name = name
	p.age = age
	return &p
}

func main() {
	bob := NewPerson("bob", 55)
	bob.desc()
	fmt.Printf("%T\n", bob)
	fmt.Println("###################")

	var admin admin
	admin.desc()
	fmt.Printf("%T\n", admin)
	admin.name = "tom"
	admin.age = 45
	fmt.Printf("%s %d\n", admin.name, admin.age)

}
