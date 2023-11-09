package main

import (
	"fmt"
)

func main() {
	my_array := [4]int{10, 20, 30, 40}
	my_slice := my_array[1:3]

	fmt.Println(my_slice)
	fmt.Println(len(my_slice))
	fmt.Println(cap(my_slice))

	my_slice = append(my_slice, 900, 800, 700)
	fmt.Println(my_slice)
	fmt.Println(len(my_slice))
	fmt.Println(cap(my_slice))

}
