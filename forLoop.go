package main 

import "fmt"

func main() {
	// like c/c++ or others go doesn't has any () parenthesis in for loop dexlaration
	for i:=0;i<10;i++{
		fmt.Println("Now at",i,"iteration")
	}

	// for loop as while
	i:=0;
	for i<10{
		fmt.Println("Now at",i,"iteration")
		i++;
	}
}