package main 

import "fmt"

func main() {
	var n int
	fmt.Printf("Enter array size: \n")
	fmt.Scan(&n)

	var arr  = make([]int,n+5) // make a runtime dynamic array
	fmt.Println("Give array element: ")
	for i:=0;i<n;i++{
		fmt.Scan(&arr[i])
	}

	fmt.Printf("Array numbers: \n")
	for i:=0;i<n;i++{
		fmt.Printf("%d -> %v\n",i,arr[i])
	}
}