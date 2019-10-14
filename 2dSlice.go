package main

import "fmt"

func main() {
	var primes [][]int
	primes = make([][]int,5,100) // make a dynamic array using make(type,len,cap)
	for i:=0;i<5;i++{
		primes[i] = make([]int,3,100)
		for j:=0;j<3;j++{
			primes[i][j]=i+j
		}
	}
	
	for i:=0;i<5;i++{
		for j:=0;j<3;j++{
			fmt.Printf("%v ",primes[i][j])
		}
		fmt.Println()
	}
	
	primes = append(primes,make([]int,0,100))// Append a 2d index last
	primes[5] = append(primes[5],7,8,9) // append 1d in index 5
	fmt.Println(len(primes)) // len() gives length
	fmt.Println(len(primes[5]))
	
	p := primes // Reference same memory address
	p[3][2] = 50 // Also change main primes slice
	fmt.Println()
	
	
	for i:=0;i<6;i++{
		for j:=0;j<3;j++{
			fmt.Printf("%v ",primes[i][j])
		}
		fmt.Println()
	}

	fmt.Println(p[1:])
	fmt.Println(p[1:3])
	fmt.Println(p[:])
	fmt.Println(p[3:5])
	fmt.Println(p[:4])
	fmt.Println(p[1:3],p[3:6])

}
