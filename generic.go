package main

import (
	"fmt"
)

//learning generic in go
//sum
//input: 1, 2 or 1.5, 1.5
//output: 3 or 3.00

//creating a custom type
// type customType2 interface{
// 	constraints.Ordered | []byte | []rune
// }

type customType interface {
	int | int32 | float32 | float64
}

func sum[T customType](a, b T) T {
	return a + b
}

func main() {
	result := sum(1.5, 1.5)

	fmt.Println("The sum is ", result)
}
