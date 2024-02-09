package main
import "fmt"

// func main(){
//   for i:= 0; i < 10; i++ {
// 	fmt.Println("increaze", i)
//   }
// }

//range 
// func main(){
// 	rvariable := []string{"dodo", "rice", "stew","beans"}

// for i,j := range rvariable {
// 	fmt.Println(i,j)
// }
// }
// map

func main(){
	name := map[int] string{
		1 : "solly",
		2 : "dolly",
		3 : "sally",
		4 : "sully",
	}
	for key, value := range name{
		fmt.Println(key, value)
	}
}