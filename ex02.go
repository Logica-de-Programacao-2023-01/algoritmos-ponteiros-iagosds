package main

import "fmt"

func porcento(ptr *int) {
	if *ptr%2 == 0 {
		*ptr = 0
	} else {
		*ptr = 1
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	ptr := &n
	fmt.Println("ptr: ", *ptr)
	porcento(ptr)
	fmt.Println("ptr: ", *ptr)
	if *ptr == 0 {
		fmt.Print("Ponteiro é par")
	} else if *ptr == 1 {
		fmt.Print("Ponteiro é ímpar")
	}
}
