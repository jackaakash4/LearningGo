package main

import "fmt"

//learning generic in go
//sum
//input: 1, 2 or 1.5, 1.5
//output: 3 or 3.00

func sum(a, b int) int {
	return a + b
}

func main() {
	result := sum(1, 2)

	fmt.Println("The sum is ", result)
}
