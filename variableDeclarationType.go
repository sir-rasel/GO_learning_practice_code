package main

import "fmt"

// global or package level declaration type 1 factored declaration
var(
	integer int
	floating float32
	boool bool
)

// global or package level declaration type 2
// var integer1, floating1, boool1 int,float,bool = 10, 20.2, true // invalid because many type in one declaration
 var integer1 int = 10
 var floating1 float32  = 10.20
 var boool1 bool  = true

// x,y:=10,10.20 // invalid short declaration outside function

func main() {
	// short variable declaration (local variable)
	x := 50
	y := 10.20 

	integer = 15
	floating = 20.5
	boool = false

	fmt.Println(x,y,integer,integer1,floating,floating1,boool,boool1)
}
