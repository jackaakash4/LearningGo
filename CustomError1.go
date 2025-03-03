//divisible by 0 custom error

package main

import (
	"errors"
	"fmt"
)

type customError struct {
	arg     int
	message string
}

func (e *customError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func validateDivisior(a, b int) (int, error) {
	if b == 0 {
		return -1, &customError{b, "Divisible by zero"}
	}
	return a / b, nil
}

func main() {

	result, err := validateDivisior(20, 10)

	var ae *customError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("Your result is ", result)
	}

}
