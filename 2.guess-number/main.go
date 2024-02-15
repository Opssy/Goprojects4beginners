//Todo using Rand golang package build a simple golang guessing game

package main
import "fmt"
import "math/rand"
import "time"

func main(){
	fmt.Println("Guessing game")


	//random generator
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	secretNumber := r.Intn(10)

	var guess int
    
	for{
	   fmt.Println("Enter your guess number")
		fmt.Scan(&guess)

		if guess < secretNumber{
			fmt.Println("too big")
		} else if guess > secretNumber{
		fmt.Println("Too small ")
		} else{
		fmt.Println("you win!!!!!!")
		break
		}
	}
}