package main

import (
	"fmt"
)

func main() {
	main_slice := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	slice_of_main_slice := main_slice[3:10]
	fmt.Println(slice_of_main_slice)
}
