package main

import "fmt"

func main() {
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Printf("n1=%v, n2=%v", a, b)
}

func swap(n1 *int, n2 *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
}
