package main

import (
	"fmt"
)

func main() {
	codes := map[string]string{"en": "English", "fr": "French", "Ar": "Arabic", "hi": "Hindi"}
	codes = make(map[string]string)
	fmt.Println(codes)
}
