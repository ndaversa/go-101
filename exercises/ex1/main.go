package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func Pretty(p Person) string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

func main() {
	ppl := []Person{
		Person{Name: "Nino", Age: 33},
		Person{Name: "Jonathan", Age: 38},
	}

	for _, p := range ppl {
		fmt.Println(Pretty(p))
	}
}
