package main

import (
	"fmt"
	"maps"
)

func maps_examples() {
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

	n := map[string]int{
		"hello": 5,
		"world": 0,
	}

	fmt.Println("m is equal to n: ", maps.Equal(m, n))

	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok2 := m["world"]
	fmt.Println(v, ok2)

	v, ok3 := m["goodbye"]
	fmt.Println(v, ok3)

	delete(m, "hello")

	example_3_11()
}

func example_3_11() {
	fmt.Println("example_3_11")

	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}
}

func using_strucsts() {
	type person struct {
		name string
		age  int
		pet  string
	}
}
