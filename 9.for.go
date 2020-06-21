package main

import "fmt"

func main() {

	for i := 1; i < 10; i++ {
		fmt.Println("Hello World increment ", i, "count")
	}

	for i := 10; i > 0; i-- {
		fmt.Println("Hello World decrement", i, "count")
	}
}
