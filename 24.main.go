package main

import "fmt"

func main() {

	codes := map[string]string{"en": "English", "fr": "French", "hi": "Hindi"}
	codes["it"] = "italian"
	fmt.Println(codes)

}
