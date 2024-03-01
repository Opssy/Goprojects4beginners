package main

import (
	"fmt"
	"time"
)





func main(){
	start := time.Now()
	userName := fetchUser()
     
	respChan := make(chan any, 2)

    go fetchUserLikes(userName, respChan)
	go fetchUserMatch(userName, respChan)
   
	
	fmt.Println("Time taken: ", time.Since(start))
}

func fetchUser() string {
   time.Sleep(time.Second * 3)
   return "Ope"
}
 
func fetchUserLikes(userName string, respChan chan any){
  time.Sleep(time.Second * 3)

   return <- 12

}
func fetchUserMatch( userName string, respChan chan any){
  time.Sleep(time.Second * 3)
  return <- "Dayo"
}