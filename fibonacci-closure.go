package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var a, b int
	b = 1
	return func() int {
		ret := a
		a, b = b, a+b
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 2; i++ {
		fmt.Println(f())
	}
}
