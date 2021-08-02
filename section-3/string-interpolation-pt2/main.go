package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	UserName       string
	Age            int
	FavoriteNumber float64
}

func main() {
	var user User
	reader = bufio.NewReader(os.Stdin)
	user.UserName = readString("What is your name?")
	user.Age = readInt("How old are you?")
	user.FavoriteNumber = readFloat("What is your favorite number?")

	fmt.Printf("Your name is %s. You are %d years old and your favorite number is %.2f.\n",
		user.UserName,
		user.Age,
		user.FavoriteNumber)
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

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		promot()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}
	}

}
