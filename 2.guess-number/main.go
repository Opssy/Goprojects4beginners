//Todo using Rand golang package build a simple golang guessing game

package main
import "fmt"
import "math/rand"
import "time"
func main(){

	min, max := 1, 1000
	var guess int 
	rand.Seed(time.Now().UnixNano())
	
	guess = rand.Intn(max - min) + min

	fmt.Println("Enter your guess")
	fmt.Scan(&guess)

	if guess == guess {
		fmt.Println("You win")
	}  
	if guess != guess{
		fmt.Println("You losen")
	}
	if guess > guess{
	  fmt.Println("Too high")
	}
	if guess < guess{
	  fmt.Println("Too low")
	}

}