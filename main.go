package main

import (
	"fmt"
)

func main () {
	counter := 0
	fmt.Println("Welcome to the go Counter App!")
	
	for {
		fmt.Printf("\nThe current count is: %d\n", counter)
		fmt.Println("What would you like to do?")
		fmt.Println("1. Increment")
		fmt.Println("2. Decrement")
		fmt.Println("3. Exit")

		var choice int
		fmt.Print("Enter your choice (1-3): ")

		//fmt.Scanf reads user input
		_, err := fmt.Scanf("%d", &choice)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			counter++
			fmt.Println("Counter incremented.")
		case 2:
			counter--
			fmt.Println("Counter decremented")
		case 3:
			fmt.Println("Thank you for using the go counter App. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please enter 1,2, or 3.")
		}
	}
}