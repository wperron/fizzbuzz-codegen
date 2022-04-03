package main

//go:generate go run ./codegen -o lib.go -p main 3 Fizz 5 Buzz

import (
	"fmt"
)

func main() {
	for i := 1; i < 120; i++ {
		fmt.Println(fmt.Sprintf("%d:", i), FizzBuzz(i))
	}
}
