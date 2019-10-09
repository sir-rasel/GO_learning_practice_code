package main 

import "fmt"

func main() {
	i,j := 10,20
	// like c/c++ or others go doesn't has any () parentesis in its if--else condition statement
	// if..else if...else format is matter
	if i<j {
		fmt.Printf("%v is less than %v\n",i,j)
	} else if i==j {
		fmt.Printf("%v is equal to %v\n",i,j)
	}else {
		fmt.Printf("%v is greater than %v\n",i,j)
	}
}