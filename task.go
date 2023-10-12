package main

import "fmt"

func sum(a int, b int) int {
	a += b
	return a
}
func main() {
	a := 10
	b := 15
	res := sum(a, b)
	fmt.Println(a, " + ", b, " = ", res)
}
