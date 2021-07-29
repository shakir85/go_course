package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	// Declare a reader, and specify where do you want to read from
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')
		// Replace the appended '\n' from user input (remove it)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {
			break
		} else {
			fmt.Println(doctor.Response(userInput))
		}

	}

}

// Notes:
// Variables declaration:
// 		var myString string
// 		var myString = ""
// 		myString := ""
