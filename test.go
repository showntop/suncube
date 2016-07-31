package main

import (
	"fmt"
)

func main() {
	array := []int{1, 2, 2, 3, 4, 4}
	slice := array[1:3]
	fmt.Printf("%v", slice)
}
