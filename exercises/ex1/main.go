// Declare a struct type to contain information about a person.  Include a name and age.  Declare a function that pretty-prints values of your type.  Create a slice of people.  In `main` call your pretty-printer for each person.
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func PrettyPerson(p Person) string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

func main() {
	ppl := []Person{
		Person{
			Name: "Nino",
			Age:  33,
		},
		Person{
			Name: "Jonathan",
			Age:  38,
		},
	}

	for _, v := range ppl {
		fmt.Println(PrettyPerson(v))
	}
}
