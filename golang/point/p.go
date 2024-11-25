package main

import "log"

type Inter interface {
	A(p *Person)
}

type Person struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Code   int
	Detail Detail
}

type Detail struct {
	Country string
}

func main() {
	p := &Person{Name: "M"}
	var i Inter
	i = new(G)
	i.A(p)

	log.Printf("p:%+v", p)
	B(p)
}

type G struct {
}

func (g *G) A(p *Person) {
	p.Age = 18
	*p = Person{Name: "A"}
	C(p)
}

func C(p *Person) {
	p.Address.Code = 2
}

func B(p *Person) {
	log.Printf("B:%+v", p)
}
