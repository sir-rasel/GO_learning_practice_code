package main 

import "fmt"

func main() {
	var arr [][]int
	arr=make([][]int,5)

	for i:=0;i<3;i++{
		arr[i] = make([]int,5)
		for j:=0;j<3;j++{
			fmt.Scan(&arr[i][j])
		}
	}

	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			fmt.Print(arr[i][j])
		}
		fmt.Println()
	}
}