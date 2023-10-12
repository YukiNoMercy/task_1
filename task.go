package main

import "fmt"

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
	fmt.Println(a, " + ", b, " = ", summa)
	fmt.Println(a, " - ", b, " = ", sub)
}
