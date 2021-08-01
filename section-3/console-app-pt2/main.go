package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	// Creating a map for the MENU
	coffees := make(map[int]string)
	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("-----")
	fmt.Println("1 - Cappucino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Espresso")
	fmt.Println("Q - Quit the program")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		// Terminate here before printing: "You chose q"
		if char == 'q' || char == 'Q' {
			break
		}

		// convert string to int, wrapping 'char' in string()
		// because keyboard.GetSingleKey() returns it as rune.
		i, _ := strconv.Atoi(string(char))
		// Get the number index 'i' and call the corresponding map value:
		fmt.Println(fmt.Sprintf("You chose %s", coffees[i]))

	}
}
