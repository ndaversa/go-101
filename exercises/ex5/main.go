package main

import (
	"fmt"
	"sync"

	"github.com/ndaversa/go-101/exercises/ex5/store"
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

func storeWorker(people chan Person, keys chan uint64, db *store.InMemory, wg *sync.WaitGroup) {
	for person := range people {
		key := db.Put(person)
		keys <- key
	}
	wg.Done()
}

func fetchWorker(keys chan uint64, db *store.InMemory, wg *sync.WaitGroup) {
	for key := range keys {
		person := db.Get(key)
		fmt.Println(person)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	db := store.New()

	people := make(chan Person)
	keys := make(chan uint64)

	wg.Add(5)
	go storeWorker(people, keys, db, &wg)
	go storeWorker(people, keys, db, &wg)
	go fetchWorker(keys, db, &wg)
	go fetchWorker(keys, db, &wg)
	go fetchWorker(keys, db, &wg)

	people <- Person{Name: "Nino", Age: 33}
	people <- Person{Name: "Jonathan", Age: 38}
	people <- Person{Name: "Yan", Age: 23}
	people <- Person{Name: "Nihal", Age: 28}
	people <- Person{Name: "Long", Age: 27}
	close(people)
	wg.Wait()
}
