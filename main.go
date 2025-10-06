package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var numberOfNumbers = 100

func getUserInput() int {
	var input int
	_, err := fmt.Scanln(&input)

	if input < 0 || input > numberOfNumbers {
		fmt.Println("Please enter a number between 0 and", numberOfNumbers)
		return getUserInput()
	}

	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")

		var discard string
		fmt.Scanln(&discard)

		return getUserInput()
	}
	return input
}

func playAgain() bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Play again? (y/n): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "y", "Y":
			return true
		case "n", "N":
			return false
		default:
			fmt.Println("Invalid input. Please enter 'y' or 'n'.")
		}
	}
}

func main() {
	number := rand.Intn(numberOfNumbers + 1)

	fmt.Println("I am thinking of a number between 0 and", numberOfNumbers, "can you guess it?: ")

	guess := getUserInput()

	if guess == number {
		fmt.Println("You guessed it!")
	} else {
		fmt.Println("Sorry, the number was:", number)
	}

	restart := playAgain()

	if restart {
		main()
	} else {
		fmt.Println("Thanks for playing!")
	}

}
