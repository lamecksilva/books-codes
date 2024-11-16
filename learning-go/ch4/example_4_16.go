package main

import "fmt"

func example_4_16() {
	samples := []string{"Hello", "apple_π!"}

	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}
}
