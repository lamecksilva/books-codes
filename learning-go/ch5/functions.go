package main

import (
	"errors"
)

func divAndRemainderWithNamedReturnVariables(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err := errors.New(`cannot divide by zero`)
		return result, remainder, err
	}
	result, remainder = num/denom, num%denom

	return result, remainder, err
}

func divAndRemainder(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New(`cannot divide by zero`)
	}

	return num / denom, num % denom, nil
}

func main() {
	// result, remainder, err := divAndRemainderWithNamedReturnVariables(5, 2)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	//
	// fmt.Println(result, remainder)

	// using_functions_values()
	anonFunc()
}
