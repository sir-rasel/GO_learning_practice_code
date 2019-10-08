package main 

import "fmt"

// constant declaration both type and untype

const (
	a = 1
	b
	c float32 = 3.12
	d bool  = true
)

// constant declaration as Enumeration with iota
const (
	e = iota
	f
	g = iota + 5
	h
)

func main() {
	const i = a+e

	fmt.Println(a,b,c,d)
	fmt.Println(e,f,g,h)
	fmt.Println(i)
}