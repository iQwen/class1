package main

import (
	"fmt"
	"hello-world/demo1"
	"strconv"
)

func main() {
	var str string = "true"
	var b bool
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v\n", b, b)

	var i string = "123456789"
	var n1 int64
	n1, _ = strconv.ParseInt(i, 10, 64)
	fmt.Printf("n1 type %T  n1=%v", n1, n1)
	demo1.Hi()
}
