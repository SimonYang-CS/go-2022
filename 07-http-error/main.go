package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.github.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The status code we got is:", resp.StatusCode)
	fmt.Println("The status code text we got is:", http.StatusText(resp.StatusCode))
}
