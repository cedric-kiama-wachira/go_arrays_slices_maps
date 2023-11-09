package main

import (
	"fmt"
)

func main() {
	var grades [5]int = [5]int{30, 20, 10, 50, 40}

	for i := 0; i < len(grades); i++ {
		fmt.Println(grades[i])

	}

}
