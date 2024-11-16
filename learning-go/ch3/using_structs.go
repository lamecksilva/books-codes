package main

import "fmt"

func using_structs() {
	fmt.Println("Using structs")

	type person struct {
		name string
		age  int
		pet  string
	}

	julia := person{
		"Julia",
		40,
		"cat",
	}

	fmt.Println(julia.name)

	beth := person{
		age:  30,
		name: "Beth",
	}

	fmt.Println(beth.name)

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}

	fmt.Println(pet.name)
}
