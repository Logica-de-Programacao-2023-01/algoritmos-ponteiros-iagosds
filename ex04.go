package main

import (
	"fmt"
	"strconv"
)

func troca(ptr *int) {
	x := strconv.Itoa(*ptr)
	l := len(x)
	x1 := x[l]
	x2 := x[l-1]
	xx1, err := strconv.Atoi(string(x1))
	xx2, err2 := strconv.Atoi(string(x2))
	if err != nil {
		fmt.Print(err)
	}
	if err2 != nil {
		fmt.Print(err2)
	}
	*ptr = xx1 + xx2
}
func main() {
	num := 1234
	var ptr *int = &num
	fmt.Println("num: ", num)
	fmt.Println(*ptr)
	troca(ptr)
	fmt.Print(*ptr)
}
