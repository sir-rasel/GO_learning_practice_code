package main 

import "fmt"

func update(age *int, name *string){
	*age+=10
	*name+=" Islam"
}

func main() {
	age:=20
	name:="Saiful"
	fmt.Printf("Before : %s %d\n",name,age)
	update(&age,&name)
	fmt.Printf("After : %s %d\n",name,age)
}