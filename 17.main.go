package main

import (
	"fmt"
)

func main() {

	arr := [5]int{1, 2, 3, 4, 5}
	i := 2
	fmt.Println(arr)
	slice_one := arr[:i]
	slice_two := arr[i+1:]
	appended_slice := append(slice_one, slice_two...)
	fmt.Println(appended_slice)
}
