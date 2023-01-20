package main

import (
	"fmt"
	"strings"
)

const refString = "Mary*had+a/little-lamb"

func main() {

	// splitFunc вызывается для каждой
	// руны в строке. Если руна
	// равна любому символу в "*%,_"
	// refString разделяется.
	splitFunc := func(r rune) bool {
		return strings.ContainsRune("+-/*", r)
	}

	words := strings.FieldsFunc(refString, splitFunc)
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}

}
