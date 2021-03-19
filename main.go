/*Reads a number from standard input and inserts separating commas. Supports floating-point 
numbers.*/

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var mainPart, decPart string

	for i := 1; i < len(os.Args); i++ {

		if strings.LastIndex(os.Args[i], ".") != -1 {
			//there is a decimal point
			mainPart, decPart = splitDecs(os.Args[i])
			fmt.Printf("  %s", comma(mainPart))
			fmt.Printf(".%s\n", decPart)

		} else {
			//there is no decimal point
			fmt.Printf("  %s\n", comma(os.Args[i]))
		}
	}
}


func comma(s string) string {
// From "The Go Programming Languag" by Alan A. A. Donovan and Brian W. Kernighan
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}

func splitDecs(s string) (string, string) {
	dot := strings.LastIndex(s, ".")
	s1 := s[:dot]
	s2 := s[dot+1:]

	return s1, s2
}