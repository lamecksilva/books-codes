package main

import "fmt"

func maps() {
	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2

	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittens"])
	totalWins["Kittens"]++

	fmt.Println(totalWins["Kittens"])
	totalWins["Lions"]++

	fmt.Println(totalWins["Lions"])

	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok2 := m["world"]
	fmt.Println(v, ok2)

	v, ok3 := m["goodbye"]
	fmt.Println(v, ok3)
}
