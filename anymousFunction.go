package main

import "fmt"

// anymous function 3
var(
	area = func(l,b int) int {return l*b}
)

func main() {
	// anymous function 1
	fmt.Printf("Area 1: %d\n",func(l int, b int) int {return l * b}(2, 3))
	// anymous function 2
	func(l int, b int) {
		fmt.Printf("Area 2: %d\n",l * b)
	}(20, 30)

	fmt.Printf("Area 3: %d\n",area(5,6))	
}