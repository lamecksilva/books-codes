package main

import "fmt"

func example_4_13() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	// in for-range loop, the i is the position(index) of the data structure
	// the second variable (v) is the value of that position
	for _, v := range evenVals {
		fmt.Println(v)
	}
}
