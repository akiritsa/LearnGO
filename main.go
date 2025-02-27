package main

import (
	"fmt"
)

func main() {
	s := "gopher"
	fmt.Println("Hello and welcome, %s!", s)
	for i := 1; i <= 5; i++ {
		fmt.Println("i =", 100/i)
	}

	inc := increment()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())

	fmt.Println(increment2())
	fmt.Println(increment2())
	fmt.Println(increment2())
	fmt.Println(increment2())
}
func increment() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
func increment2() int {
	count := 0
	count++
	return count
}
