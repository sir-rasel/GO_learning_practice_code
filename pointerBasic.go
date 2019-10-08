package main 

import "fmt"

// Go doesn't support any pointer arithmatic like C

func main() {
	var num int = 10
	increment(&num)
	fmt.Println(num)
	fmt.Println(increment(&num))
}

func increment(x *int) int {
	*x++;
	return *x;
} 