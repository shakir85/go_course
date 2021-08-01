package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

func main() {
	// This will only take place if we can't open the
	// keyboard package:
	// keyboard.Open() will return either an 'error' or 'nil'
	err := keyboard.Open()

	// If the input is not nil, then it means we have error. Log it
	// as fatal & close the application.
	if err != nil {
		log.Fatal(err)
	}
	// Anonymous function:
	// A function without a name. This function will
	// close the keyboard as soon as the func main() stops running.
	// And this is the meaning of 'defer' keyword (to defer the execution up until
	// the parent function stops).
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press any key on the keyboard. Press ESC to quit...")

	for {
		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if key != 0 {
			fmt.Println("You pressed", char, key)
		} else {
			fmt.Println("You pressed", char)
		}

		if key == keyboard.KeyEsc {
			break
		}
	}

	fmt.Println("Program Exiting")
}
