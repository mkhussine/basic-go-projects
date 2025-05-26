// Calculator CLI project
// Arithmetic calculator via terminal input
package main

import "fmt"

func main() {
	fmt.Println("Welcome to our Calculator")
	fmt.Println("Enter the first Number: ")
	var first_number float64
	fmt.Scanln(&first_number)
	fmt.Println("Enter the second Number: ")
	var second_number float64
	fmt.Scanln(&second_number)
	fmt.Println("Enter the operation: ")
	fmt.Println("1 for +")
	fmt.Println("2 for -")
	fmt.Println("3 for *")
	fmt.Println("4 for /")
	var operation int
	fmt.Scanln(&operation)
	switch operation {
	case 1:
		fmt.Println(first_number, "+", second_number, "= ", (first_number + second_number))
	case 2:
		fmt.Println(first_number, "-", second_number, "= ", (first_number - second_number))
	case 3:
		fmt.Println(first_number, "*", second_number, "= ", (first_number * second_number))
	case 4:
		if second_number == 0 {
			fmt.Println("Second number MUST be positive or negative")
		} else {
			fmt.Println(first_number, "/", second_number, "= ", (first_number / second_number))
		}
	}
}
