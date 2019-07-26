package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	counter := 0

	zn := 1.0
	znBefore := 0.0

	// for n := 1; n <= 10; n++ {
	// 	zn = zn - (math.Pow(zn, 2)-x)/(2*zn)
	// }

	for math.Abs(zn-znBefore) > 0 {
		znBefore = zn
		zn = zn - (math.Pow(zn, 2)-x)/(2*zn)
		counter++
	}

	fmt.Println(counter)
	return zn
}

func main() {
	fmt.Println(sqrt(4))
}
