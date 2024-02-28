package main

import (
	"fmt"
	"time"
)


// goroutine example 

// 1. main goroutine starts
// 2. Invokes helloworld and helloworld goroutine starts
// 3. If there is no pause using the sleep method, the main will then invoke goodbye() and exit before the helloworld goroutine finishes its execution.
func main(){

	go helloworld()
	time.Sleep(1 * time.Second)

	go goodbyee()
}

func helloworld(){
	fmt.Println("Hello World!")
}
func goodbyee(){
	fmt.Println("Goodbye!")
}

