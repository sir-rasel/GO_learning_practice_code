package main 

import "fmt"

func main() {
	var str string = "Hello"
	// range for loop over Hello string
	// range for loop return a pair of index and value
	for ind,val:=range str{
		fmt.Printf("Charecter at index %d is : %c\n",ind,val)
	}
}