package main

import (
	"fmt"

	"github.com/ndaversa/go-101/exercises/ex3/store"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

func (p Person) Key() uint64 {
	return uint64(p.Age)
}

func main() {
	db := store.New()

	db.Put(Person{Name: "Nino", Age: 33})
	db.Put(Person{Name: "Jonathan", Age: 38})

	for _, i := range []uint64{33, 38} {
		fmt.Println(db.Get(i))
	}
}
