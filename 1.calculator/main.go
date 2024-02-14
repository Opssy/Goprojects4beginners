package main

import "fmt"

func calculator(){
	 
	var operator string;
	fmt.Println("Enter operator: +, -, *, /")
	fmt.Scan(&operator)


	var a,b float64

	fmt.Println("Enter your two digits: ")
	fmt.Scan(&a, &b)

	switch choice{
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
    case "/":
		fmt.Println(a / b)
	 default:
        fmt.Println("Invalid operator")
	}
}