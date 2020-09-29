package main

import (
	"fmt"
)

func main() {
	var ciclos int
	var e float64 = 0.0

	fmt.Scan(&ciclos)

	for i := 0; i <= ciclos; i++ {
		e += 1.0 / float64(factorial(uint(i)))
	}

	fmt.Println(e)
}

func factorial(n uint) uint {
	if n == 0 {
		return 1
	}
	var f uint = 1
	var i uint
	for i = 1; i <= n; i++ {
		f *= i
	}
	return f
}