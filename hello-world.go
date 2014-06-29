package main

import (
	"fmt"
	"github.com/lokinell/newmath"
)

func main() {
	fmt.Println("hello, world. Sqrt(2) = %v", newmath.Sqrt(2))
	fmt.Println(newmath.GetOwner())

	defer myDefer()

	newmath.SetOwner("sholloy")
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

	fmt.Println(newmath.GetOwner())
}

func myDefer() {
	fmt.Println("run from my defer...")
}
