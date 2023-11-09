package main

import (
	"fmt"
)

func main() {
	source_slice := []int{10, 20, 30, 40, 50}
	dest_slice := make([]int, 3)
	num := copy(dest_slice, source_slice)

	fmt.Println(dest_slice)
	fmt.Println("Number of elements copied from source slice are:", num)

}
