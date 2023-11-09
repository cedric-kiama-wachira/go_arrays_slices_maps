package main

import (
	"fmt"
)

func main() {
	my_array_one := [6]int{0, 1, 2, 3, 4, 5}
	my_slice_one := my_array_one[:2]
	my_arry_two := [6]int{7, 8, 9, 10, 11, 12}
	my_slice_two := my_arry_two[:2]
	my_joined_slice := append(my_slice_one, my_slice_two...)

	fmt.Println(my_joined_slice)

}
