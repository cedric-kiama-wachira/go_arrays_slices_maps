package main

import "fmt"

func main() {

	// slice_name := make_function([]data_type, length, capacity)
	slice := make([]int, 5, 10)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

}
