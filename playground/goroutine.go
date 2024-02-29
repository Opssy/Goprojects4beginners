package main

import (
	"fmt"
	"math/rand"
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

//channels
// Syntax to declare a channel
// ch := make(chan Type)
// ch <- value // send value to the channel
// value := <- ch // receive value from the channel


func simpleChannel(){
	// Create an unbuffered channel of type int
	c := make(chan int)

	// Create a goroutine that sends a value to the channel
	go func() {
		c <- 42
	}()

	// Receive the value from the channel
	value := <-c

	fmt.Println("Received value:", value)
}

//buffered channel 

func BufferedChannel(){
	// // Create a buffered channel of type int with capacity of 5
	// c := make(chan int, 5)

	// // Send 5 values to the channel
	// for i := 0; i < 5; i++ {
	// 	c <- i
	// }

	// // Receive and print the values from the channel
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(<-c)
	// }

	//create a buffered channel of type string with capacity of 5

	cars := make(chan string, 5)

	// Start five Goroutines that add cars to the channel
	go addCar("BMW", cars)
	go addCar("Mercedes", cars)
	go addCar("Ferrari", cars)
	go addCar("Toyota", cars)
	go addCar("Honda", cars)

	// Wait for the cars to be added to the channel
	time.Sleep(2 * time.Second)
	close(cars)

	//Start a Goroutine that simulates the race
	go startRace(cars)

	// Wait for the race to finish
	time.Sleep(5 * time.Second)

	fmt.Println("Race over!")
}

// to add a car to the channel 
func addCar(name string, cars chan string) {
	cars <- name
	fmt.Println(name, "added to the race!")
} 

func startRace(cars chan string) {
	for {
	  // Receive a car from the channel
    car, open := <-cars
    if !open {
        break
    }

    fmt.Println(car, "is racing...")
	
     // Simulate the race by waiting for a random duration
    time.Sleep(time.Duration(1+rand.Intn(5)) * time.Second)
    }
   fmt.Println("All cars have finished the race!")

}