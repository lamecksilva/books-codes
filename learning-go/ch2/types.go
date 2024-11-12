package main

import "fmt"

const value = 10

func main() {
	// Exercise 2.1
	i := 20
	f := float32(i)

	fmt.Println(i)
	fmt.Println(f)

	// Exercise 2.2
	i2 := value
	f2 := float32(value)

	fmt.Println(i2)
	fmt.Println(f2)

	// Exercise 2.3
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	fmt.Println(b + 1)
	fmt.Println(smallI + 1)
	fmt.Println(bigI + 1)
}
