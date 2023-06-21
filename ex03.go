package main

import (
	"fmt"
)

func reverse(ptr string) {
	var s string
	for i := len(ptr); i >= 0; i-- {
		s += string(ptr[i])
	}
	ptr = s
}
func main() {
	str := "iago"
	fmt.Print("*ptr ", str)
	reverse(str)
	fmt.Print("*ptr ", str)
}
