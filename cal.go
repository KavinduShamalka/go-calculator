package main

import "fmt"

func main() {

	println()
	fmt.Println("This is a calculator.")
	userInput()

}

// Get user input
func userInput() {

	var (
		num1     int
		num2     int
		operator string
	)

	println()
	// Get user input 1
	fmt.Print("Enter number 1: ")
	fmt.Scanln(&num1)

	// Get user input 2
	fmt.Print("Enter number 2: ")
	fmt.Scanln(&num2)
	fmt.Println()

	// Get operand from the user
	fmt.Print("Enter operator +, -, *, /, %: ")
	fmt.Scanln(&operator)

	result := calculator(num1, num2, operator)

	continueCal(result)

}

// Calculation
func calculator(num1 int, num2 int, operator string) int {

	// Decalare result to store the result
	var result int = 0

	// Define switch case for match operand
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "/":
		result = num1 / num2
	case "*":
		result = num1 * num2
	case "%":
		result = num1 % num2
	default:
		fmt.Println("Invalid operand")
	}

	return result
}

// continue operaton
func continueCal(preVal int) {

	var (
		cont    string
		num3    int
		operand string
	)

	fmt.Println()
	fmt.Printf("Do you want to continue with the previous value %v?\nPress 'y' or 'n' key for exit: ", preVal)
	fmt.Scanln(&cont)

	// check user wants to exit
	if cont == "n" {
		fmt.Println()
		fmt.Println("Exit")
		fmt.Println()
	}

	// check user wants to continue
	if cont == "y" {

		fmt.Println()
		fmt.Println("Enter value you want to continue: ")
		fmt.Scanln(&num3)

		fmt.Print("Enter operator +, -, *, /, %: ")
		fmt.Scanln(&operand)

		// get calculated result
		result := calculator(preVal, num3, operand)

		// continue
		continueCal(result)
	}
}
