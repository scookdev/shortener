package main

import (
	"flag"
	"fmt"
)

func main() {
	url := flag.String("url", "", "URL to shorten")
	flag.Parse()

	if *url == "" {
		fmt.Println("Please provide a URL with --url")
		return
	}

	shortened, err := ShortenURL(*url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Shortened URL:", shortened)
}
