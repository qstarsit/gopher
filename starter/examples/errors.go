package examples

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Go uses explicit error return values rather than exceptions.
// By convention, errors are the last return value and have type error (built-in interface).
// https://gobyexample.com/errors
func errorBasics() {

	// Functions return error as last value
	result, err := divide(10, 2)
	if err != nil {
		// Handle error
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result) // 5

	// Error on purpose
	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error: ", err) // "division by zero"
	}
}
