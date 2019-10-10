package main 

import "fmt"
import "time"

func main() {
	value:=50;
	// Single case value switch 
	switch value{
	case 10:
		fmt.Println("Value is : ",value)
	case 20:
		fmt.Println("Value is : ",value)
	case 30:
		fmt.Println("Value is : ",value)
	case 50:
		fmt.Println("Value is : ",value)
	case 60:
		fmt.Println("Value is : ",value)
	default:
		fmt.Println("Value is : ",value)
	}
	// Multiple case value switch 
	switch value{
	case 10,20:
		fmt.Println("Value is : ",value)
	case 30,40:
		fmt.Println("Value is : ",value)
	case 50,60:
		fmt.Println("Value is : ",value)
	case 70:
		fmt.Println("Value is : ",value)
	default:
		fmt.Println("Value is : ",value)
	}
	value = 15
	// switch with fallthrough statement
	switch value{
	case 10,15:
		fmt.Println("Value is : ",value)
		fallthrough
	case 20:
		fmt.Println("Value is : ",20)
	default:
		fmt.Println("Value is : ",value)
	}
	// switch with conditional case statement
	switch{
	case value>10:
		fmt.Println("Greater value is : ",value)
	case value==10:
		fmt.Println("Equal Value is : ",value)
	default:
		fmt.Println("Less value is : ",value)
	}
	// switch with initializer
	switch today := time.Now(); {
	case today.Day() < 5:
		fmt.Println("Clean your house.")
	case today.Day() <= 10:
		fmt.Println("Buy some wine.")
	case today.Day() > 15:
		fmt.Println("Visit a doctor.")
	case today.Day() == 25:
		fmt.Println("Buy some food.")
	default:
		fmt.Println("No information available for that day.")
	}
}