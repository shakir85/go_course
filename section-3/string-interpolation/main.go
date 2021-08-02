package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

func main() {
	reader = bufio.NewReader(os.Stdin)
	userName := readString("What is your name?")
	age := readInt("How old are you?")

	fmt.Printf("Your name is %s. You are %d years old.\n", userName, age)

}

func promot() {
	fmt.Print("-> ")
}

func readString(s string) string {
	//for loop to force input to always be correct. If the input
	// is not an integer, the loop will continue. If input is integer
	// loop will break.
	for {
		fmt.Println(s)
		promot()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)
		if userInput == "" {
			fmt.Println("Please enter a value")
		} else {
			return userInput
		}
	}
}

func readInt(s string) int {
	//for loop to force input to always be correct. If the input
	// is not an integer, the loop will continue. If input is integer
	// loop will break.
	for {
		fmt.Println(s)
		promot()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}
	}

}
