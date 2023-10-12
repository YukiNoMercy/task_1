package main

import "fmt"

func division(a float64, b float64) float64 {
	a /= b
	return a
}
func multiplication(a int, b int) int {
	a *= b
	return a
}
func subtraction(a int, b int) int {
	a -= b
	return a
}
func sum(a int, b int) int {
	a += b
	return a
}
func main() {
	a := 10
	b := 15
	div_a := 10.0
	div_b := 15.0
	div := division(div_a, div_b)
	summa := sum(a, b)
	sub := subtraction(a, b)
	multi := multiplication(a, b)
	fmt.Println(a, " + ", b, " = ", summa)
	fmt.Println(a, " - ", b, " = ", sub)
	fmt.Println(a, " * ", b, " = ", multi)
	fmt.Println(div_a, " / ", div_b, " = ", div)
}
