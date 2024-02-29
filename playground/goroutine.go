package main

import (
	"fmt"
	"time"
)


// goroutine example 

// 1. main goroutine starts
// 2. Invokes helloworld and helloworld goroutine starts
// 3. If there is no pause using the sleep method, the main will then invoke goodbye() and exit before the helloworld goroutine finishes its execution.
// func main(){
// 	go helloworld()
// 	time.Sleep(1 * time.Second)

// 	go goodbyee()
// }

// func helloworld(){
// 	fmt.Println("Hello World!")
// }
// func goodbyee(){
// 	fmt.Println("Goodbye!")
// }

//Another example

func CarRace(){
	car1 :=" BMW"
	car2 :=" Mercedes"

	//create a goroutine for each car
	go race (car1)
	go race (car2) 

    // Wait for the race to finish
	time.Sleep(1 * time.Second)
	
	fmt.Println("Race over!")
}

func race(car string){
	for i := 0; i < 5; i++ {
		fmt.Println(car, "is going", i, "meters")
		time.Sleep(1 * time.Second)
	}
}