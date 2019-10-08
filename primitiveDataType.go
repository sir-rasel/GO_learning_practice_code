package main 

import "fmt"

/* There are 3 primitive type exist ig Go:
	1. bool
	2. Numeric
		- int - int  int8  int16  int32(also called rune)  int64 uint uint8(also called byte) uint16 uint32 uint64 uintptr
		- float - float32 float64
		- complex - complex64 complex128
	3. text (string) 
*/

var(
	integer int = 100;
	floatingNumber float32 = 10.50
	flag bool = true
	name string = "Saiful Islam Rasel"
	complexNumber1 complex64 = complex(5,2)
	complexNumber2 complex64 = 2 + 3i
)

func main() {
	fmt.Printf("integer number %v, type %T\n",integer,integer);
	fmt.Printf("Floating number %v, type %T\n",floatingNumber,floatingNumber);
	fmt.Printf("Bool number %v, type %T\n",flag,flag);
	fmt.Printf("Complex number %v, type %T\n",complexNumber1,complexNumber1);
	fmt.Printf("Complex number real %v, type %T\n",real(complexNumber2),real(complexNumber2))
	fmt.Printf("Complex number imaginary %v, type %T\n",imag(complexNumber2),imag(complexNumber2))
}