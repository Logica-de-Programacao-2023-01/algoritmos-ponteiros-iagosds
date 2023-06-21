package main

import "fmt"

func soma(ptr *int, x int) {
	for i := 0; i <= x; i++ {
		*ptr += i
	}
}
func main() {
	pt := 10
	var ptr *int = &pt
	var n int = 5
	fmt.Println(*ptr)
	soma(ptr, n)
	fmt.Print(*ptr)
	num := 1234
}
