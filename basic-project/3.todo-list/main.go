// Todo list in golang
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct{
   Name string
   Completed bool 
} 

var tasks []Task

func addTask(task string){
	newTask := Task{Name:task, Completed:false}
	tasks = append(tasks, newTask)

	fmt.Println("Task added successfully")
}
func listTask(){
    for _, task := range tasks{
		status := "Not Completed"
		if task.Completed{
			status = "Completed"
		}
     fmt.Printf("%s - %s\n", task.Name, status)
    } 
}

func completedTask(index int){
   if index >= 1 || index <= len(tasks){
	  tasks[index - 1].Completed = true
	  fmt.Println("Task completed successfully")
   } else {
	fmt.Println("Invalid index")
   }
}
func editTask(index int, newName string){
   if index >= 1 || index <= len(tasks){
	  tasks[index - 1].Name = newName
	  fmt.Println("Task edited successfully")
   } else {
	fmt.Println("Invalid index")
   }
}
func deleteTask(index int){
	if index >= 1 || index <= len(tasks){
	  tasks = append(tasks[:index - 1], tasks[index:]...)
	  fmt.Println("Task deleted successfully")
   } else {
	fmt.Println("Invalid index")
   }
}
func main() {
	var indexInput int
	var taskInput, newTaskInput string  

	fmt.Println("1. Add task")
	fmt.Println("2. List task")
	fmt.Println("3. Mark as Complete")
	fmt.Println("4. Edit task")
	fmt.Println("5. Delete task")
	fmt.Println("6. Exit")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter choice(1, 2, 3, 4, 5, 6): ")
		scanner.Scan()
		input := scanner.Text()

		choice, err := strconv.Atoi(input)
		if err != nil{
			fmt.Println("Invalid choice")
			continue
		}
		switch choice {
			case 1: 
				fmt.Print("Enter task: ")
				scanner.Scan()
		      taskInput = scanner.Text()
			    addTask(taskInput)
		   case 2:
			   listTask()	

			case 3:	
				fmt.Print("Enter task index: ")	
				scanner.Scan()
				indexInput, err = strconv.Atoi(scanner.Text())
				completedTask(indexInput)	

			case 4:
				fmt.Print("Enter task index: ")
				scanner.Scan()
				indexInput, err = strconv.Atoi(scanner.Text())
				fmt.Print("Enter new task name: ")
				
				scanner.Scan()
				newTaskInput = scanner.Text()
				editTask(indexInput, newTaskInput)

			case 5:
				fmt.Print("Enter task index: ")
				scanner.Scan()
				indexInput, err = strconv.Atoi(scanner.Text())
				deleteTask(indexInput)

			case 6:
				os.Exit(0)
			
			default:
				fmt.Println("Invalid choice")

		}
	}
}