package main

import (
	"fmt"
)

func main() {

	my_array := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	my_slice := my_array[:3]

	fmt.Println(my_array)
	fmt.Println(my_slice)

	my_slice[1] = 999

	fmt.Println("After Modification")
	fmt.Println(my_array)
	fmt.Println(my_slice)
}
