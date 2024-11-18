package main

import "fmt"

func f1(a string) int {
	return len(a)
}

func f2(a string) int {
	total := 0
	for _, v := range a {
		intValue := int(v)
		fmt.Println(string(v), v, intValue)
		total += intValue
	}
	return total
}

func using_functions_values() {
	var myFuncVariable func(string) int

	myFuncVariable = f1
	result := myFuncVariable("Hello")
	fmt.Println(result)

	myFuncVariable = f2
	result = myFuncVariable("Hello")
	fmt.Println(result)
}
