package main

import (
	"fmt"
)

func main() {

	var grades [5]int = [5]int{60, 70, 80, 90, 100}
	for index, element := range grades {
		fmt.Println(index, "=>", element)
	}
}
