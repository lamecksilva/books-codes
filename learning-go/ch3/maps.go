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
}
