package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	src, err := readInput()
	if err != nil {
		fail(err)
	}
	words := strings.Fields(src)
	fmt.Println(len(words))
}

// readInput reads pattern and source string
// from command line arguments and returns them.
func readInput() (src string, err error) {
	src = os.Args[1]
	if src == "" {
		return src, errors.New("missing string to match")
	}
	return src, nil
}

// fail prints the error and exits.
func fail(err error) {
	fmt.Println("match:", err)
	os.Exit(1)
}
