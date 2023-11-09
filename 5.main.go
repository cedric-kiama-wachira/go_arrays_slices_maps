package main

import (
	"fmt"
)

func main() {
	var grades [5]int = [5]int{100, 55, 90, 45}
	fmt.Println(grades)
	grades[0] = 65
	fmt.Println(grades)
}
