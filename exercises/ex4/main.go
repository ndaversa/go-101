package main

import (
	"fmt"
	"sync"

	"github.com/ndaversa/go-101/exercises/ex4/store"
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

func Log(p Person, db *store.InMemory, wg *sync.WaitGroup) {
	key := db.Put(p)
	person := db.Get(key)
	fmt.Println(person)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	db := store.New()

	wg.Add(5)
	go Log(Person{Name: "Nino", Age: 33}, db, &wg)
	go Log(Person{Name: "Jonathan", Age: 38}, db, &wg)
	go Log(Person{Name: "Yan", Age: 23}, db, &wg)
	go Log(Person{Name: "Nihal", Age: 28}, db, &wg)
	go Log(Person{Name: "Long", Age: 27}, db, &wg)
	wg.Wait()
}
