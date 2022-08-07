package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	src := readInput()
	words := strings.Fields(src)
	fmt.Println(len(words))
}

// readInput reads pattern and source string
// from command line arguments and returns them.
func readInput() (src string) {
	src = os.Args[1]
	return src
}

