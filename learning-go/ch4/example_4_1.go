package main

import "fmt"

func example_4_1() {
	x := 10

	if x > 5 {
		fmt.Println(x)
		x := 5 // Shadowing variable x (Creating other variable with same name)

		fmt.Println(x)
	}
	fmt.Println(x)
}

func main() {
	fmt.Println("Func main")
	// example_4_1()
	// example_4_7()
	// example_4_13()
	// example_4_15()
	example_4_16()
}
