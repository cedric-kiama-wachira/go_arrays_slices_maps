package main

import "fmt"

func main() {
	codes := map[string]string{"en": "English", "fr": "French", "hi": "Hindi"}
	for key := range codes {
		delete(codes, key)
	}
	fmt.Println(codes)
}
