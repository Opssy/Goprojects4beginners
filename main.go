package main

import (
	"fmt"
	// "strings"
)

func main (){

	//strings
	var nameOne string = "sola"
	nameFour:= "solly"

	// int
	fmt.Println(nameOne, nameFour);

	// int
   var ageOne int = 20;
   ageTwo := 40
   fmt.Println(ageOne,ageTwo);

   //bits&memory
   var numOne int8 = 25
   var numFour int64 = 120

   fmt.Println(numOne, numFour);

   //float
   var scoreOne float32 =-1.5
   var scoreTwo float64  = 3344455566645432234566.7

   scoreThree := 1.5

   fmt.Println(scoreOne,scoreTwo,scoreThree)

   //Array &Slices

   var ages = [3]int{20,25,26}

   names := [4]string{"dada", "solly", "yoshi","doe"}
    fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	//slices
	//slice allows you to add - append and remove -
	var score = []int{100, 60, 50}
     score[2] = 25
	 score = append(score, 85)
	 fmt.Println(score, len(score))

	 //slice range
	 rangeOne := names[1:3]
	 rangeTwo := names[2:]
	 rangeThree := names[:3]

	 fmt.Println(rangeOne, rangeTwo, rangeThree)

	 //loop

	 for i :=0; i < 5; i++ {
		fmt.Println("value of i is:", i)
	 }

	 for index, value := range names {
		fmt.Printf("the value at index %v is %v \n", index, value);
	 }
	 for _, value := range names {
		fmt.Printf("the value is %v \n", value);
	 }

	 //boolean
	 age := 45

	 fmt.Println(age <= 50)
	 fmt.Println(age >= 50)
	 fmt.Println(age == 45)
	 fmt.Println(age != 50)

	 //if else statement
	 if age < 30 {
		fmt.Println("age is greater than 30")
	 } else if age < 40 {
       fmt.Println("age is less than 50")
	 } else{
		fmt.Println("age is not less than 45")
	 }

	 //maps
	 menu := map[string]float64{
		"soup": 4.99,
		"salad": 3.29,
		"okoro": 4.50,
		"dodo": 2.30,
	 }
	 fmt.Println(menu)
	 fmt.Println(menu["soup"])

	 //looping maps
	 for k, v := range menu {
		fmt.Println(k, "-", v)
	 }

	 //int as key  type
	 phonebook := map[int]string{
		23456756: "mario",
		54543355: "josh",
		12345677: "mars",
	 }
	 fmt.Println(phonebook)


	 //struct
	 mybill := newBill("mario's bill")
	 fmt.Println(mybill)
}

// *******************************

//function
//call or create function

// func sayGreeting(){
//   fmt.Println("Good morning boys")
// }

// func main(){
// 	sayGreeting();
// }

// parameters & argument

// func familyName(fname string, lname string, age int){
// 	fmt.Println("Hello", fname, lname, "and your age is", age)
// }

// func sayBye(n string){
// 	fmt.Printf("Goodbye", n)
// }

// func cycleNames(n []string, f func(string)){
// 	for _, v := range n {
// 		f(v)
// 	}
// }
  

// function that return multiple values

// func getInitials(n string)(string, string){
// 	s:= strings.ToUpper(n)
// 	names := strings.Split(s, "")

// 	var initials []string
// 	for _, v := range names{
// 		initials = append(initials, v[:1])
// 	}
// 	if len(initials) > 1 {
// 		return initials[0],initials[1]
// 	}
// 	return initials[0], "-"
// }
// func main(){
// 	familyName("Liam", "Sade", 23)

// 	cycleNames([]string{"cloud", "tifa", "clode"}, sayBye)

// 	fn1, sn1 := getInitials("tifa lockhart")
// 	fmt.Println(fn1, sn1)

// 	fn2, sn2 := getInitials("Ope Yemi")
// 	fmt.Println(fn2, sn2)

// 	sayHello("mario")
// }




