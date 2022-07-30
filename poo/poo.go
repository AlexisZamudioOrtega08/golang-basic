package main

import "fmt"

type Animal struct {
	name string
	age  int
}

type Cat struct {
	sound  string
	animal *Animal
}

type Dog struct {
	sound  string
	animal *Animal
}

/*func (a *Animal) dog() {
	// * indicates that we are changing the value of the Animal struct itself
	// if * is not declared, we are working "locally" on the Animal struct and not on the value of the Animal struct
	a.type_ = "dog"
}c
func (a *Animal) cat() {
	a.type_ = "cat"
}
*/

func main() {
	a_c := &Animal{"Nami", 1}
	a_d := &Animal{"Barry", 3}
	c := Cat{sound: "meow", animal: a_c}
	d := Dog{sound: "woof", animal: a_d}
	fmt.Println(c.sound)
	fmt.Println(d.sound)

}
