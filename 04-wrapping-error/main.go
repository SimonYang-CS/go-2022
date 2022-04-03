package main

import (
	"errors"
	"fmt"
	"os"
)

//When an error is passed back through your code, you often want to add additional context to it

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

func main() {
	err := fileChecker("not_here.txt")
	if err != nil {
		fmt.Println(err)

		//os.IsNotExist does not support wrapped error
		if os.IsNotExist(err) {
			fmt.Println("IsNotExist true")
		} else {
			fmt.Println("IsNotExist false")
		}

		//errors.Is support wrapped error
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("ErrNotExist true")
		} else {
			fmt.Println("ErrNotExist false")
		}

		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println(wrappedErr)
		}
	}
}
