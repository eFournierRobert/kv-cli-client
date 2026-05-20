package main

import (
	"fmt"
	"kv-cli-client/internal/api"
	"kv-cli-client/internal/input"
	"os"
)

func main() {
	for {
		userInput := input.GetUserInput()
		fmt.Println(userInput)

		if userInput != nil && len(userInput) > 0 {
			switch userInput[0] {
			case "exit":
				os.Exit(0)
			case "set":
				if len(userInput) < 3 {
					fmt.Println("Error: set needs a key and a value")
					break
				}
				api.SendSet(userInput[1], userInput[2])
			}
		}
	}
}
