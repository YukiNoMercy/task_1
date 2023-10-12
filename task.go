package main

import "fmt"

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
	summa := sum(a, b)
	sub := subtraction(a, b)
	multi := multiplication(a, b)
	fmt.Println(a, " + ", b, " = ", summa)
	fmt.Println(a, " - ", b, " = ", sub)
	fmt.Println(a, " - ", b, " = ", multi)
}
