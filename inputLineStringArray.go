package main 

import(
	"fmt"
	"bufio"
	"os"
)

func main() {
	var arr [10]string
	fmt.Println("Give string input:")
	// way one using bufio.NewReader() it uses to split on particular charecter
	/*
	reader := bufio.NewReader(os.Stdin)
	for i:=0;i<5;i++{
		arr[i],_=reader.ReadString(' ')
	}*/

	// way 2 using bufio.NewScanner() it works on new line
	input:=bufio.NewScanner(os.Stdin)
	for i:=0;i<5;i++{
		input.Scan()
		arr[i]=input.Text()
	}

	for i:=0;i<5;i++{
		fmt.Printf("%s\n",arr[i])
	}
}