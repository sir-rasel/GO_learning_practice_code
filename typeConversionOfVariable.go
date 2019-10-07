package main 

import(
	"fmt"
	"strconv" // for string conversion
)

// for type conversion : T(variable) where T indicates Type

func main() {
	var a int = 5;
	var b float32 = float32(a);
	fmt.Println(a,b)

	var str string = strconv.Itoa(a)
	fmt.Println(str)
}