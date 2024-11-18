package main

import "fmt"

func anonFunc() {
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("Printing", j, "from inside anonymous function")
		}(i)
	}
}
