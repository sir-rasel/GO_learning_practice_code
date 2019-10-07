package main 

import "fmt"

// Basic format : func functionName (perameter with type) (return type list) { }

func add(x int,y int) int { // natural declaration
	return x+y
}

func sub(x,y int) int { // perametar type short declaration
	return x-y
}

func mul(x,y int) (multyplyResult int){ // variable name return declaration, also called naked return
	multyplyResult = x*y
	return;
}

func swap(x,y int) (int, int){ // multyple return declaration
	return y,x
}

func main() {
	x := 50
	y := 20
	
	fmt.Println(add(x,y))
	fmt.Println(sub(x,y))
	fmt.Println(mul(x,y))
	fmt.Println(swap(x,y))
}