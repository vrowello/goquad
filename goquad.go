package main

import (
	"fmt"
	"math"
)

var (
	a float64
	b float64
	c float64
)

func main() {
	solutions := make(chan float64, 2)

	fmt.Printf("a: ")
	fmt.Scanln(&a)
	fmt.Printf("b: ")
	fmt.Scanln(&b)
	fmt.Printf("c: ")
	fmt.Scanln(&c)

	go s1(a, b, c, solutions)
	go s2(a, b, c, solutions)

	fmt.Println("X = ", <-solutions)
	fmt.Println("X = ", <-solutions)
}

func s1(a float64, b float64, c float64, solutions chan float64) {
	var div float64
	var quo float64

	div = (b * -1) + math.Sqrt((math.Pow(b, 2))-(4*a*c))
	quo = div / (2 * a)

	solutions <- quo
}

func s2(a float64, b float64, c float64, solutions chan float64) {
	var div float64
	var quo float64

	div = (b * -1) - math.Sqrt((math.Pow(b, 2))-(4*a*c))
	quo = div / (2 * a)

	solutions <- quo
}
