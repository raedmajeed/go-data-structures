package basics

import "fmt"

type Animal struct{}

func (a *Animal) Eat(animal string) {
	fmt.Println(animal, " is eating")
}

func (a *Animal) Drink(animal string) {
	fmt.Println(animal, " is drinking")
}

type Cow struct {
	Animal
}

func Inheritence() {
	c := Cow{}
	c.Eat("cow")
	c.Drink("cow")
}
