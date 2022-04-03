package main

import "fmt"

// constant error

type Error string

func (e Error) Error() string { return string(e) }

const FNE = Error("FNE")

func main() {
	len, err := open("not-exists.file")

	if err == FNE {
		fmt.Println(err)
		return
	}

	fmt.Printf("len=[%d], err=[%s]", len, err)
}

func open(file string) (int, error) {
	return 0, FNE
}
