package main

import "fmt"

func main() {
	var grades [5]int = [5]int{10, 20, 50, 80, 100}
	fmt.Println(grades)

	fruits := [3]string{"Mangoe", "Kiwi", "Overcado"}
	fmt.Println(fruits)

	implicit := [...]float64{66.66, 77.77, 88.88}
	fmt.Println(implicit)
}
