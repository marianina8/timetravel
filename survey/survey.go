package survey

import (
	"bufio"
	"fmt"
	"os"
)

func Run() {
	displayMenu()
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		choice := scanner.Text()
		printResponse(choice)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}

func displayMenu() {
	fmt.Println("    How was the journey?      ")
	fmt.Println("                             ")
	fmt.Println("    1. Breathtaking          ")
	fmt.Println("    2. Too fast!             ")
	fmt.Println("    3. Can we go again?      ")
	fmt.Print("Enter your choice (1/2/3): ")
}

func printResponse(choice string) {
	switch choice {
	case "1":
		fmt.Println("Glad you enjoyed the breathtaking view of time!")
	case "2":
		fmt.Println("88 miles per hour can indeed feel a bit quick!")
	case "3":
		fmt.Println("Hold on tight to that flux capacitor and let's do it again!")
	default:
		fmt.Println("Hmm, that wasn't an option on our time circuit.")
	}
}
