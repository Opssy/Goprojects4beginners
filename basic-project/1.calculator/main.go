package main

import "fmt"

func calculator(){
	 
	var operator string;
	fmt.Println("Enter operator: +, -, *, /")
	_, err  := fmt.Scan(&operator)
	if err != nil {
		fmt.Println("Error reading operator:", err)
		return
	}
	var a,b float64
	fmt.Println("Enter your two digits: ")
	_, err = fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("Error reading digits:", err)
		return
	}
	
	switch operator {
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

func main(){
	calculator()
}