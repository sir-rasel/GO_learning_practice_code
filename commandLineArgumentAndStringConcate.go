package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s, sep := "", ""
	// first way
	for i:=0;i<len(os.Args);i++{
		s+=sep+os.Args[i];
		sep=" "
	}
	fmt.Println(s)

	// second way
	s=""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	// third way
	fmt.Println(strings.Join(os.Args[1:], " "))
}