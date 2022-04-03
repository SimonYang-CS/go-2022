package main

import (
	"errors"
	"fmt"
)

//Use Strings for Simple Errors

func main() {
	err1 := BasicErr1("err1")
	if err1 != nil {
		fmt.Println(err1)
	}

	err2 := BasicErr2("err2")
	if err2 != nil {
		fmt.Println(err2)
	}
}

func BasicErr1(msg string) error {
	return errors.New(msg)
}

func BasicErr2(msg string) error {
	return fmt.Errorf("error: %s", msg)
}
